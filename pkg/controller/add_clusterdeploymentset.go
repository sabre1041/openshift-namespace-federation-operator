package controller

import (
	"github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/controller/clusterdeploymentset"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, clusterdeploymentset.Add)
}
