package namespacefederation

import (
	"io/ioutil"
	"testing"
	"text/template"

	federationv1alpha1 "github.com/raffaelespazzoli/openshift-namespace-federation-operator/pkg/apis/federation/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var instance = federationv1alpha1.NamespaceFederation{
	ObjectMeta: metav1.ObjectMeta{
		Namespace: "ciao",
	},
}

const templateFile string = "/home/rspazzol/go/src/github.com/raffaelespazzoli/openshift-namespace-federation-operator/templates/federation-controller/federation-controller.yaml"

func TestFullConfig(t *testing.T) {
	text, err := ioutil.ReadFile(templateFile)
	if err != nil {
		t.Errorf("Error reading template file: %v", err)
		t.Fail()
	}
	template, err := template.New("template").Parse(string(text))

	objs, err := processTemplateArray(&instance, template)
	if err != nil {
		t.Errorf("Error processing the template: %v", err)
		t.Fail()
	}
	t.Logf("array length %d", len(*objs))
	t.Logf("resulting manifest: %+v", *objs)
}