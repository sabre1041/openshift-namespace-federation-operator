package namespacefederation

import (
	federationv1alpha1 "github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1"
)

func (r *ReconcileNamespaceFederation) createOrUpdateFederationControlPlane(instance *federationv1alpha1.NamespaceFederation) error {

	return r.CreateOrUpdateTemplatedResources(instance, "", instance, federationControllerTemplate)

}
