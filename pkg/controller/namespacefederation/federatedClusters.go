package namespacefederation

import (
	"context"
	"errors"
	"io/ioutil"
	"strings"
	"text/template"

	federationv2v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/core/v1alpha1"
	federationv1alpha1 "github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var federatedClusterTemplate *template.Template
var remoteFederatedClusterTemplate *template.Template
const remoteServiceAccountName string = "federation-controllee"

func (r *ReconcileNamespaceFederation) createOrUpdateFederatedClusters(instance *federationv1alpha1.NamespaceFederation) []error {
	addClusters, deleteClusters, err := r.getAddAndDeleteCluster(instance)
	if err != nil {
		log.Error(err, "Error calculating add and delete clusters for instance", "instance", *instance)
		return []error{err}
	}
	errs := make([]error, len(addClusters)+len(deleteClusters))
	// first we take care of deleting the deleteclusetr

	for _, cluster := range deleteClusters {
		err = r.manageDeleteCluster(cluster, instance)
		log.Error(err, "Unable to successfully delete cluster", "cluster", cluster)
		errs = append(errs, err)
	}

	//then we add new clusters.

	for _, cluster := range addClusters {
		err = r.manageAddCluster(cluster, instance)
		log.Error(err, "Unable to successfully add cluster", "cluster", cluster)
		errs = append(errs, err)
	}

	return errs

}

//adding a cluster consist of creating the namespace in the target cluster and populating it with the service account and then creating the federatedcluster and the secret in the same namespace instance
func (r *ReconcileNamespaceFederation) manageAddCluster(cluster string, instance *federationv1alpha1.NamespaceFederation) error {
	// create new namespace in remote cluster
	remoteClusterClient, err := r.getAdminClientForCluster(cluster, instance)
	if err != nil {
		log.Error(err, "Error creating remote client for cluster", "cluster", cluster)
		return err
	}
	namespace := corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: instance.GetNamespace(),
		},
	}

	err = createIFNotExistsResource(remoteClusterClient, &namespace)
	if err != nil {
		log.Error(err, "Error creating remote namespace for cluster", "cluster", cluster, "namespace", namespace.GetName())
		return err
	}

	// apply template in remote cluster
	objs, err := processTemplateArray(nil, remoteFederatedClusterTemplate)
	if err != nil {
		log.Error(err, "error creating manifest from template")
		return err
	}
	for _, obj := range *objs {
		err = createOrUpdateResource(remoteClusterClient, instance, &obj)
		if err != nil {
			log.Error(err, "unable to create/update object", "object", &obj)
			return err
		}
	}

	//apply template in local cluster
	remoteSecret, err := getSecretForRemoteServiceAccount(remoteClusterClient, cluster, instance)
	if err != nil {
		log.Error(err, "unable to retrive remote secret for cluster", "cluster", cluster, "namespace", instance.GetNamespace())
		return err
	}
	federatedClusterMerge := federatedClusterMerge{
		Namespace:    instance.GetNamespace(),
		Cluster:      cluster,
		CaCRT:        string(remoteSecret.Data["ca.crt"]),
		ServiceCaCRT: string(remoteSecret.Data["service-ca.crt"]),
		Token:        string(remoteSecret.Data["token.crt"]),
		SecretName:   cluster + "-remote",
	}

	objs, err = processTemplateArray(federatedClusterMerge, federatedClusterTemplate)
	if err != nil {
		log.Error(err, "error creating manifest from template")
		return err
	}
	for _, obj := range *objs {
		err = createOrUpdateResource(remoteClusterClient, instance, &obj)
		if err != nil {
			log.Error(err, "unable to create/update object", "object", &obj)
			return err
		}
	}

	return nil
}



func getSecretForRemoteServiceAccount(remoteClusterClient *RemoteClusterClient, cluster string, instance *federationv1alpha1.NamespaceFederation) (*corev1.Secret, error) {
	remoteServiceAccount := &corev1.ServiceAccount{}
	err := remoteClusterClient.client.Get(context.TODO(), types.NamespacedName{
		Namespace: instance.GetNamespace(),
		Name:      remoteServiceAccountName,
	}, remoteServiceAccount)
	if err != nil {
		log.Error(err, "unable to retrieve remote service account", "service account", remoteServiceAccountName, "cluster", cluster)
		return nil, err
	}

	var remoteSecretName string
	for _, secret := range remoteServiceAccount.Secrets {
		if strings.Contains(secret.Name, "token") {
			remoteSecretName = secret.Name
			break
		}
	}
	if remoteSecretName == "" {
		err = errors.New("unable to find remote token secret")
		log.Error(err, "unable to find remote token secret", "service account", remoteServiceAccountName, "cluster", cluster)
		return nil, err
	}
	remoteTokenSecret := &corev1.Secret{}
	err = remoteClusterClient.client.Get(context.TODO(), types.NamespacedName{
		Namespace: instance.GetNamespace(),
		Name:      remoteSecretName,
	}, remoteTokenSecret)
	if err != nil {
		log.Error(err, "unable to retrieve remote token secret", "token secret", remoteSecretName, "cluster", cluster)
		return nil, err
	}
	return remoteTokenSecret, nil
}



func InitializeFederatedClusterTemplates(federatedClusterTemplateFileName string, remoteFederatedClusterTemplateFileName string) error {
	text, err := ioutil.ReadFile(federatedClusterTemplateFileName)
	if err != nil {
		log.Error(err, "Error reading rolebinding template file", "filename", federatedClusterTemplateFileName)
		return err
	}

	federatedClusterTemplate, err = template.New("FederatedCluster").Parse(string(text))
	if err != nil {
		log.Error(err, "Error parsing template", "template", text)
		return err
	}

	text, err = ioutil.ReadFile(remoteFederatedClusterTemplateFileName)
	if err != nil {
		log.Error(err, "Error reading rolebinding template file", "filename", federatedClusterTemplateFileName)
		return err
	}

	remoteFederatedClusterTemplate, err = template.New("FederatedCluster").Parse(string(text))
	if err != nil {
		log.Error(err, "Error parsing template", "template", text)
		return err
	}

	return nil
}

type federatedClusterMerge struct {
	Namespace    string
	Cluster      string
	CaCRT        string
	ServiceCaCRT string
	Token        string
	SecretName   string
}

// deleting a cluste consist of deleting the federated namespace on the target cluster and deleting the federatedcluster and secret object in the same namespace as the instance
func (r *ReconcileNamespaceFederation) manageDeleteCluster(cluster string, instance *federationv1alpha1.NamespaceFederation) error {
	// delete the remote namespace
	remoteClusterClient, err := r.getAdminClientForCluster(cluster, instance)
	if err != nil {
		log.Error(err, "Error creating remote client for cluster", "cluster", cluster)
		return err
	}
	namespace := corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: instance.GetNamespace(),
		},
	}
	err = deleteResource(remoteClusterClient, &namespace)
	if err != nil {
		log.Error(err, "Error deleting the namespace in the remote cluster", "cluster", cluster, "namespace", namespace)
		return err
	}
	// retrieve the federated cluster to know what the associated secret is
	federatedCluster := &federationv2v1alpha1.FederatedCluster{}
	err = r.client.Get(context.TODO(), types.NamespacedName{
		Namespace: instance.GetNamespace(),
		Name:      cluster,
	}, federatedCluster)
	if err != nil {
		if apierrors.IsNotFound(err) {
			// If we don't find the cluster we assume it has already been correclty deleted
			return nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Error redetrieving federatedcluster", "name", cluster, "namespace", instance.GetNamespace())
		return err
	}

	//delete the local federatedcluster and secret
	federatedClusterMerge := federatedClusterMerge{
		Namespace:    instance.GetNamespace(),
		Cluster:      cluster,
		CaCRT:        "N/A",
		ServiceCaCRT: "N/A",
		Token:        "N/A",
		SecretName:   federatedCluster.Spec.SecretRef.Name,
	}

	objs, err := processTemplateArray(federatedClusterMerge, federatedClusterTemplate)
	if err != nil {
		log.Error(err, "error creating manifest from template")
		return err
	}
	for _, obj := range *objs {
		err = deleteResource(r, &obj)
		if err != nil {
			log.Error(err, "unable to delete object", "object", &obj)
			return err
		}
	}
	return nil
}

func (r *ReconcileNamespaceFederation) getAddAndDeleteCluster(instance *federationv1alpha1.NamespaceFederation) ([]string, []string, error) {
	//we assume that if a federated cluster object exists it means that it has been correclty federated

	federatedClusterList := &federationv2v1alpha1.FederatedClusterList{}
	err := r.client.List(context.TODO(), &client.ListOptions{Namespace: instance.GetNamespace()}, federatedClusterList)
	if err != nil {
		log.Error(err, "Error listing federatedclusters in namespace", "namespace", instance.GetNamespace())
		return nil, nil, err
	}
	// let's calculate the add clusters
	addClusters := make([]string, len(instance.Spec.Clusters))
	for _, cluster := range instance.Spec.Clusters {
		if !contains(federatedClusterList, cluster) {
			addClusters = append(addClusters, cluster)
		}
	}

	//let's calculate the delete clusters
	deleteClusters := make([]string, len(federatedClusterList.Items))
	for _, federatedCluster := range federatedClusterList.Items {
		if !contains2(instance.Spec.Clusters, &federatedCluster) {
			deleteClusters = append(deleteClusters, federatedCluster.Spec.ClusterRef.Name)
		}
	}

	return addClusters, deleteClusters, nil

}

func contains(federatedClusterList *federationv2v1alpha1.FederatedClusterList, cluster string) bool {
	for _, federatedCluster := range federatedClusterList.Items {
		if cluster == federatedCluster.Spec.ClusterRef.Name {
			return true
		}
	}
	return false
}

func contains2(clusters []string, federatedCluster *federationv2v1alpha1.FederatedCluster) bool {
	for _, cluster := range clusters {
		if cluster == federatedCluster.Spec.ClusterRef.Name {
			return true
		}
	}
	return false
}

type RemoteClusterClient struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

func (r *RemoteClusterClient) GetClient() client.Client {
	return r.client
}

func (r *RemoteClusterClient) GetScheme() *runtime.Scheme {
	return r.scheme
}

func (r *ReconcileNamespaceFederation) getAdminClientForCluster(cluster string, instance *federationv1alpha1.NamespaceFederation) (*RemoteClusterClient, error) {
	return nil, nil
}
