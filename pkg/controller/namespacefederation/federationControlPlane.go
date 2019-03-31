package namespacefederation

import (
	federationv1alpha1 "github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1"
)

func (r *ReconcileNamespaceFederation) createOrUpdateFederationControlPlane(instance *federationv1alpha1.NamespaceFederation) error {

	objs, err := processTemplateArray(instance, federationControllerTemplate)
	if err != nil {
		log.Error(err, "error creating manifest from template")
		return err
	}
	for _, obj := range *objs {
		err = r.CreateOrUpdateResource(instance, &obj)
		if err != nil {
			log.Error(err, "unable to create object", "object", &obj)
			return err
		}
	}
	return nil
}
