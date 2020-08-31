package main

import (
	"io/ioutil"
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/yaml"
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

// LoadTestWorkflow returns a workflow relative to the test file
func LoadTestWorkflow(path string) *wfv1.Workflow {
	yamlBytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return LoadWorkflowFromBytes(yamlBytes)
}

// LoadWorkflowFromBytes returns a workflow unmarshalled from an yaml byte array
func LoadWorkflowFromBytes(yamlBytes []byte) *wfv1.Workflow {
	v := &wfv1.Workflow{}
	MustUnmarshallYAML(string(yamlBytes), v)
	return v
}

func MustUnmarshallYAML(text string, v interface{}) {
	err := yaml.UnmarshalStrict([]byte(text), v)
	if err != nil {
		log.Warnf("invalid YAML: %v", err)
		err = yaml.Unmarshal([]byte(text), v)
	}
	if err != nil {
		panic(err)
	}
}