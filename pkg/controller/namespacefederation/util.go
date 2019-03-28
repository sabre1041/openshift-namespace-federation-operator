package namespacefederation

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"

	federationv1alpha1 "github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1"
	extensionv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/yaml"
)

var federatedClusterTemplate *template.Template
var remoteFederatedClusterTemplate *template.Template
var federationControllerTemplate *template.Template
var federatedTypesTemplate *template.Template

func InitializeFederatedClusterTemplates(federatedClusterTemplateFileName string, remoteFederatedClusterTemplateFileName string) error {
	text, err := ioutil.ReadFile(federatedClusterTemplateFileName)
	if err != nil {
		log.Error(err, "Error reading rolebinding template file", "filename", federatedClusterTemplateFileName)
		return err
	}

	federatedClusterTemplate = template.New("FederatedCluster").Funcs(template.FuncMap{
		"parseNewLines": func(value string) string {
			return strings.Replace(value, "\n", "\n\n", -1)
		},
	})

	federatedClusterTemplate, err = federatedClusterTemplate.Parse(string(text))
	if err != nil {
		log.Error(err, "Error parsing template", "template", text)
		return err
	}

	text, err = ioutil.ReadFile(remoteFederatedClusterTemplateFileName)
	if err != nil {
		log.Error(err, "Error reading rolebinding template file", "filename", federatedClusterTemplateFileName)
		return err
	}

	remoteFederatedClusterTemplate, err = template.New("RemoteFederatedCluster").Parse(string(text))
	if err != nil {
		log.Error(err, "Error parsing template", "template", text)
		return err
	}

	return nil
}

// InitializeTemplates initializes the temolates needed by this controller, it must be called at controller boot time
func InitializeFederationControlPlaneTemplates(federationControllerTemplateFileName string) error {

	text, err := ioutil.ReadFile(federationControllerTemplateFileName)
	if err != nil {
		log.Error(err, "Error reading statefulset template file", "filename", federationControllerTemplateFileName)
		return err
	}

	federationControllerTemplate, err = template.New("Job").Parse(string(text))
	if err != nil {
		log.Error(err, "Error parsing template", "template", text)
		return err
	}

	return nil
}

func InitializeFederatedTypesTemplates(federatedTypesTemplateFileName string) error {
	text, err := ioutil.ReadFile(federatedTypesTemplateFileName)
	if err != nil {
		log.Error(err, "Error reading rolebinding template file", "filename", federatedTypesTemplateFileName)
		return err
	}

	federatedTypesTemplate = template.New("FederatedTypes").Funcs(template.FuncMap{
		"getName": func(simpleType metav1.TypeMeta) string {
			lowerCasePlural := strings.ToLower(simpleType.Kind) + "s"
			group := simpleType.APIVersion[:strings.Index(simpleType.APIVersion, "/")]
			return lowerCasePlural + "." + group
		},
		"getGroup": func(apiVersion string) string {
			return apiVersion[:strings.Index(apiVersion, "/")]
		},
		"getVersion": func(apiVersion string) string {
			return apiVersion[strings.Index(apiVersion, "/"):]
		},
		"namespaced": func(crd extensionv1beta1.CustomResourceDefinition) bool {
			return crd.Spec.Scope == "namespace"
		},
	})

	federatedTypesTemplate, err = federatedTypesTemplate.Parse(string(text))
	if err != nil {
		log.Error(err, "Error parsing template", "template", text)
		return err
	}

	return nil
}

func processTemplate(data interface{}, template *template.Template) (*unstructured.Unstructured, error) {
	obj := unstructured.Unstructured{}
	var b bytes.Buffer
	err := template.Execute(&b, data)
	if err != nil {
		log.Error(err, "Error executing template", "template", template)
		return &obj, err
	}
	bb, err := yaml.YAMLToJSON(b.Bytes())
	if err != nil {
		log.Error(err, "Error trasnfoming yaml to json", "manifest", string(b.Bytes()))
		return &obj, err
	}

	err = json.Unmarshal(bb, &obj)
	if err != nil {
		log.Error(err, "Error unmarshalling json manifest", "manifest", string(bb))
		return &obj, err
	}

	return &obj, err
}

func processTemplateArray(data interface{}, template *template.Template) (*[]unstructured.Unstructured, error) {
	obj := []unstructured.Unstructured{}
	var b bytes.Buffer
	err := template.Execute(&b, data)
	if err != nil {
		log.Error(err, "Error executing template", "template", template)
		return &obj, err
	}
	bb, err := yaml.YAMLToJSON(b.Bytes())
	if err != nil {
		log.Error(err, "Error trasnfoming yaml to json", "manifest", string(b.Bytes()))
		return &obj, err
	}

	err = json.Unmarshal(bb, &obj)
	if err != nil {
		log.Error(err, "Error unmarshalling json manifest", "manifest", string(bb))
		return &obj, err
	}

	return &obj, err
}

func createOrUpdateResource(r RuntimeClient, instance *federationv1alpha1.NamespaceFederation, obj metav1.Object) error {
	runtimeObj, ok := (obj).(runtime.Object)
	if !ok {
		return fmt.Errorf("is not a %T a runtime.Object", obj)
	}

	if instance != nil {
		_ = controllerutil.SetControllerReference(instance, obj, r.GetScheme())
	}

	obj2 := unstructured.Unstructured{}
	obj2.SetKind(runtimeObj.GetObjectKind().GroupVersionKind().Kind)
	if runtimeObj.GetObjectKind().GroupVersionKind().Group != "" {
		obj2.SetAPIVersion(runtimeObj.GetObjectKind().GroupVersionKind().Group + "/" + runtimeObj.GetObjectKind().GroupVersionKind().Version)
	} else {
		obj2.SetAPIVersion(runtimeObj.GetObjectKind().GroupVersionKind().Version)
	}

	err := r.GetClient().Get(context.TODO(), types.NamespacedName{
		Namespace: obj.GetNamespace(),
		Name:      obj.GetName(),
	}, &obj2)

	if err != nil {
		return r.GetClient().Create(context.TODO(), runtimeObj)
	}
	obj.SetResourceVersion(obj2.GetResourceVersion())
	return r.GetClient().Update(context.TODO(), runtimeObj)
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

func createIfNotExists(r RuntimeClient, obj metav1.Object) error {
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
