package namespacefederation

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"text/template"

	"github.com/ghodss/yaml"

	federationv1alpha1 "github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func (r *ReconcileNamespaceFederation) createOrUpdateFederationControlPlane(instance *federationv1alpha1.NamespaceFederation) error {

	objs, err := processTemplateArray(instance, federationControllerTemplate)
	if err != nil {
		log.Error(err, "error creating manifest from template")
		return err
	}
	for _, obj := range *objs {
		err = createOrUpdateResource(r, instance, &obj)
		if err != nil {
			log.Error(err, "unable to create object","object", &obj)
			return err
		}
	}

	// for _, template := range templates {
	// 	obj, err := processTemplate(instance, template)
	// 	if err != nil {
	// 		log.Error(err, "error creating manifest from template")
	// 		return err
	// 	}
	// 	err = r.createOrUpdateResource(instance, obj)
	// 	if err != nil {
	// 		log.Error(err, "unable to create object")
	// 		return err
	// 	}
	// }
	return nil
}

var roleBindingTemplate *template.Template
var roleTemplate *template.Template
var serviceTemplate *template.Template
var statefulsetTemplate *template.Template
var federationControllerTemplate *template.Template

var templates = [4]*template.Template{roleTemplate, roleBindingTemplate, serviceTemplate, statefulsetTemplate}

// InitializeTemplates initializes the temolates needed by this controller, it must be called at controller boot time
func InitializeFederationControlPlaneTemplates(federationControllerTemplateFileName string, roleBindingTempateFileName string, roleTemplateFilename string, serviceTempateFileName string, statefulsetTemplateFilename string) error {
	text, err := ioutil.ReadFile(roleBindingTempateFileName)
	if err != nil {
		log.Error(err, "Error reading rolebinding template file", "filename", roleBindingTempateFileName)
		return err
	}

	roleBindingTemplate, err = template.New("RoleBinding").Parse(string(text))
	if err != nil {
		log.Error(err, "Error parsing template", "template", text)
		return err
	}

	text, err = ioutil.ReadFile(roleTemplateFilename)
	if err != nil {
		log.Error(err, "Error reading role template file", "filename", roleTemplateFilename)
		return err
	}

	roleTemplate, err = template.New("Job").Parse(string(text))
	if err != nil {
		log.Error(err, "Error parsing template", "template", text)
		return err
	}

	text, err = ioutil.ReadFile(serviceTempateFileName)
	if err != nil {
		log.Error(err, "Error reading service template file", "filename", serviceTempateFileName)
		return err
	}

	serviceTemplate, err = template.New("Job").Parse(string(text))
	if err != nil {
		log.Error(err, "Error parsing template", "template", text)
		return err
	}

	text, err = ioutil.ReadFile(statefulsetTemplateFilename)
	if err != nil {
		log.Error(err, "Error reading statefulset template file", "filename", statefulsetTemplateFilename)
		return err
	}

	statefulsetTemplate, err = template.New("Job").Parse(string(text))
	if err != nil {
		log.Error(err, "Error parsing template", "template", text)
		return err
	}

	text, err = ioutil.ReadFile(federationControllerTemplateFileName)
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
