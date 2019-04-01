package multiplenamespacefederation

import (
	"io/ioutil"

	"text/template"
)

var remoteGlobalLoadBalancerTemplate *template.Template
var localLoadBalancerServiceAccountTemplate *template.Template

func InitializeRemoteGlobaLoadBalancerTemplate(remoteGlobalLoadBalancerTemplateFileName string) error {

	text, err := ioutil.ReadFile(remoteGlobalLoadBalancerTemplateFileName)
	if err != nil {
		log.Error(err, "Error reading statefulset template file", "filename", remoteGlobalLoadBalancerTemplateFileName)
		return err
	}

	remoteGlobalLoadBalancerTemplate, err = template.New("RemoteGlobalLoadBalancer").Parse(string(text))
	if err != nil {
		log.Error(err, "Error parsing template", "template", text)
		return err
	}

	return nil
}

func InitializeLocalLoadBalancerServiceAccountTemplate(localLoadBalancerServiceAccountTemplateFileName string) error {

	text, err := ioutil.ReadFile(localLoadBalancerServiceAccountTemplateFileName)
	if err != nil {
		log.Error(err, "Error reading statefulset template file", "filename", localLoadBalancerServiceAccountTemplateFileName)
		return err
	}

	localLoadBalancerServiceAccountTemplate, err = template.New("RemoteGlobalLoadBalancer").Parse(string(text))
	if err != nil {
		log.Error(err, "Error parsing template", "template", text)
		return err
	}

	return nil
}
