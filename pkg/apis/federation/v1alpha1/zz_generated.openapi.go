// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.MultipleNamespaceFederation":       schema_pkg_apis_federation_v1alpha1_MultipleNamespaceFederation(ref),
		"github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.MultipleNamespaceFederationSpec":   schema_pkg_apis_federation_v1alpha1_MultipleNamespaceFederationSpec(ref),
		"github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.MultipleNamespaceFederationStatus": schema_pkg_apis_federation_v1alpha1_MultipleNamespaceFederationStatus(ref),
		"github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.NamespaceFederation":               schema_pkg_apis_federation_v1alpha1_NamespaceFederation(ref),
		"github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.NamespaceFederationSpec":           schema_pkg_apis_federation_v1alpha1_NamespaceFederationSpec(ref),
		"github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.NamespaceFederationStatus":         schema_pkg_apis_federation_v1alpha1_NamespaceFederationStatus(ref),
	}
}

func schema_pkg_apis_federation_v1alpha1_MultipleNamespaceFederation(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MultipleNamespaceFederation is the Schema for the multiplenamespacefederations API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.MultipleNamespaceFederationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.MultipleNamespaceFederationStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.MultipleNamespaceFederationSpec", "github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.MultipleNamespaceFederationStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_federation_v1alpha1_MultipleNamespaceFederationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MultipleNamespaceFederationSpec defines the desired state of MultipleNamespaceFederation",
				Properties: map[string]spec.Schema{
					"namespaceFederationSpec": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Ref:         ref("github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.NamespaceFederationSpec"),
						},
					},
					"namespaceSelector": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"),
						},
					},
				},
				Required: []string{"namespaceSelector"},
			},
		},
		Dependencies: []string{
			"github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.NamespaceFederationSpec", "k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"},
	}
}

func schema_pkg_apis_federation_v1alpha1_MultipleNamespaceFederationStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MultipleNamespaceFederationStatus defines the observed state of MultipleNamespaceFederation",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_federation_v1alpha1_NamespaceFederation(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NamespaceFederation is the Schema for the namespacefederations API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.NamespaceFederationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.NamespaceFederationStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.NamespaceFederationSpec", "github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.NamespaceFederationStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_federation_v1alpha1_NamespaceFederationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NamespaceFederationSpec defines the desired state of NamespaceFederation",
				Properties: map[string]spec.Schema{
					"clusters": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html These are cluster name ref to cluster defined in the cluster registry",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.Cluster"),
									},
								},
							},
						},
					},
					"federatedTypes": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.TypeMeta"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.Cluster", "k8s.io/apimachinery/pkg/apis/meta/v1.TypeMeta"},
	}
}

func schema_pkg_apis_federation_v1alpha1_NamespaceFederationStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NamespaceFederationStatus defines the observed state of NamespaceFederation",
				Properties: map[string]spec.Schema{
					"clusterRegistrationStatuses": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.ClusterRegistrationStatus"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1.ClusterRegistrationStatus"},
	}
}
