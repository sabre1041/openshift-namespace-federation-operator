package namespacefederation

import (
	"context"
	"fmt"

	federationv1alpha1 "github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_namespacefederation")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new NamespaceFederation Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileNamespaceFederation{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("namespacefederation-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource NamespaceFederation
	err = c.Watch(&source.Kind{Type: &federationv1alpha1.NamespaceFederation{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner NamespaceFederation
	// err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
	// 	IsController: true,
	// 	OwnerType:    &federationv1alpha1.NamespaceFederation{},
	// })
	// if err != nil {
	// 	return err
	// }

	return nil
}

var _ reconcile.Reconciler = &ReconcileNamespaceFederation{}

type RuntimeClient interface {
	GetClient() client.Client
	GetScheme() *runtime.Scheme
}

// ReconcileNamespaceFederation reconciles a NamespaceFederation object
type ReconcileNamespaceFederation struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

func (r *ReconcileNamespaceFederation) GetClient() client.Client {
	return r.client
}

func (r *ReconcileNamespaceFederation) GetScheme() *runtime.Scheme {
	return r.scheme
}

// Reconcile reads that state of the cluster for a NamespaceFederation object and makes changes based on the state read
// and what is in the NamespaceFederation.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileNamespaceFederation) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling NamespaceFederation")

	// Fetch the NamespaceFederation instance
	instance := &federationv1alpha1.NamespaceFederation{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	r.createOrUpdateFederationControlPlane(instance)

	r.createOrUpdateFederatedClusters(instance)

	r.createOrUpdateFederatedTypes(instance)

	return reconcile.Result{}, nil
}

func createOrUpdateResource(r RuntimeClient, instance *federationv1alpha1.NamespaceFederation, obj metav1.Object) error {
	runtimeObj, ok := (obj).(runtime.Object)
	if !ok {
		return fmt.Errorf("is not a %T a runtime.Object", obj)
	}

	if instance != nil {
		_ = controllerutil.SetControllerReference(instance, obj, r.GetScheme())
	}

	err := r.GetClient().Create(context.TODO(), runtimeObj)
	if err != nil && apierrors.IsAlreadyExists(err) {
		return r.GetClient().Update(context.TODO(), runtimeObj)
	} else if err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	}
	return nil
}

func deleteResource(r RuntimeClient, obj metav1.Object) error {
	runtimeObj, ok := (obj).(runtime.Object)
	if !ok {
		return fmt.Errorf("is not a %T a runtime.Object", obj)
	}

	err := r.GetClient().Delete(context.TODO(), runtimeObj, nil)
	if err != nil && !apierrors.IsNotFound(err) {
		log.Error(err, "unable to delete object ", "object", runtimeObj)
		return err
	}
	return nil
}

func createIFNotExistsResource(r RuntimeClient, obj metav1.Object) error {
	runtimeObj, ok := (obj).(runtime.Object)
	if !ok {
		return fmt.Errorf("is not a %T a runtime.Object", obj)
	}

	err := r.GetClient().Create(context.TODO(), runtimeObj)
	if err != nil && !apierrors.IsAlreadyExists(err) {
		log.Error(err, "unable to create object ", "object", runtimeObj)
		return err
	}
	return nil
}
