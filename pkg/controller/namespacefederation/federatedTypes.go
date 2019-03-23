package namespacefederation

import (
	"context"
	"strings"

	federationv2v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	extensionv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

	federationv1alpha1 "github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1"
)

type SimpleTyepCRDPair struct {
	SimpleType metav1.TypeMeta
	CRD        extensionv1beta1.CustomResourceDefinition
}

func (r *ReconcileNamespaceFederation) createOrUpdateFederatedTypes(instance *federationv1alpha1.NamespaceFederation) error {
	addTypes, deleteDeleteTypes, err := r.getAddAndDeleteTypes(instance)

	if err != nil {
		log.Error(err, "Error calculating add and delete types for instance", "instance", *instance)
		return err
	}

	//for delete types we delete only the federated type and not the CRD, because the CRD may be in use by some other namespaces
	err = r.deleteFederatedTypes(instance, deleteDeleteTypes)
	if err != nil {
		log.Error(err, "Error deleting federated types for instance", "instance", *instance)
		return err
	}

	//for add clusters we generate the crd and the federated types and we create them
	pairs, err := generateCRDS(addTypes)
	if err != nil {
		log.Error(err, "Error generating crd for addTypes", "addTyoes", addTypes)
		return err
	}
	for _, pair := range pairs {
		err = createOrUpdateResource(r, nil, &pair.CRD)
		if err != nil {
			log.Error(err, "unable to create/update object", "object", &pair.CRD)
			return err
		}
	}

	objs, err := processTemplateArray(pairs, federatedTypesTemplate)
	if err != nil {
		log.Error(err, "error creating manifest from template")
		return err
	}
	for _, obj := range *objs {
		obj.SetNamespace(instance.GetNamespace())
		err = createOrUpdateResource(r, instance, &obj)
		if err != nil {
			log.Error(err, "unable to create/update object", "object", &obj)
			return err
		}
	}

	return nil
}

func (r *ReconcileNamespaceFederation) deleteFederatedTypes(instance *federationv1alpha1.NamespaceFederation, deleteTypes []metav1.TypeMeta) error {
	for _, simpleType := range deleteTypes {
		federatedType := getFederatedType(simpleType)
		federatedType.ObjectMeta.Namespace = instance.GetNamespace()
		err := deleteResource(r, &federatedType)
		if err != nil {
			log.Error(err, "Unable to delete federated type", "type", federatedType)
		}
	}
	return nil
}

func getFederatedType(simpleType metav1.TypeMeta) federationv2v1alpha1.FederatedTypeConfig {
	return federationv2v1alpha1.FederatedTypeConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name: strings.ToLower(simpleType.Kind) + "s",
		},
	}
}

func (r *ReconcileNamespaceFederation) getAddAndDeleteTypes(instance *federationv1alpha1.NamespaceFederation) ([]metav1.TypeMeta, []metav1.TypeMeta, error) {
	federatedTypeList := &federationv2v1alpha1.FederatedTypeConfigList{}
	err := r.client.List(context.TODO(), &client.ListOptions{Namespace: instance.GetNamespace()}, federatedTypeList)
	if err != nil {
		log.Error(err, "Error listing federated types in namespace", "namespace", instance.GetNamespace())
		return nil, nil, err
	}
	// let's calculate the add federatedType
	addTypes := make([]metav1.TypeMeta, len(instance.Spec.FederatedTypes))
	for _, simpleType := range instance.Spec.FederatedTypes {
		if !containsSimpleType(federatedTypeList, simpleType) {
			addTypes = append(addTypes, simpleType)
		}
	}

	//let's calculate the delete federatedType
	deleteTypes := make([]metav1.TypeMeta, len(federatedTypeList.Items))
	for _, federatedType := range federatedTypeList.Items {
		if !containsFederatedType(instance.Spec.FederatedTypes, &federatedType) {
			deleteTypes = append(deleteTypes, federatedType.TypeMeta)
		}
	}

	return addTypes, deleteTypes, nil
}

func containsSimpleType(federatedTypeList *federationv2v1alpha1.FederatedTypeConfigList, simpleType metav1.TypeMeta) bool {
	for _, federatedType := range federatedTypeList.Items {
		if simpleType == getAncestorType(&federatedType) {
			return true
		}
	}
	return false
}
func containsFederatedType(simpleTypes []metav1.TypeMeta, federatedType *federationv2v1alpha1.FederatedTypeConfig) bool {
	ancestorType := getAncestorType(federatedType)
	for _, simpleType := range simpleTypes {
		if simpleType == ancestorType {
			return true
		}
	}
	return false
}

func getAncestorType(federatedType *federationv2v1alpha1.FederatedTypeConfig) metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: federatedType.Spec.Target.Group + "/" + federatedType.Spec.Target.Version,
		Kind:       federatedType.Spec.Target.Kind,
	}
}

func generateCRDS(types []metav1.TypeMeta) ([]SimpleTyepCRDPair, error) {
	return nil, nil
}
