/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"flag"
	"fmt"
	"time"

	argoclient "github.com/argoproj/argo/pkg/client/clientset/versioned"
	argoprojv1alpha1 "github.com/argoproj/argo/pkg/client/clientset/versioned/typed/workflow/v1alpha1"
	"github.com/cenkalti/backoff"
	"github.com/golang/glog"
	"github.com/pkg/errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// ArgoClientInterface ...
type ArgoClientInterface interface {
	Workflow(namespace string) argoprojv1alpha1.WorkflowInterface
}

// ArgoClient ...
type ArgoClient struct {
	argoProjClient argoprojv1alpha1.ArgoprojV1alpha1Interface
}

// Workflow ...
func (argoClient *ArgoClient) Workflow(namespace string) argoprojv1alpha1.WorkflowInterface {
	return argoClient.argoProjClient.Workflows(namespace)
}

// NewArgoClientOrFatal ...
func NewArgoClientOrFatal(initConnectionTimeout time.Duration) *ArgoClient {
	var argoProjClient argoprojv1alpha1.ArgoprojV1alpha1Interface
	var operation = func() error {
		restConfig, err := rest.InClusterConfig()
		if err != nil {
			return errors.Wrap(err, "Failed to initialize the RestConfig")
		}
		argoProjClient = argoclient.NewForConfigOrDie(restConfig).ArgoprojV1alpha1()
		return nil
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = initConnectionTimeout
	err := backoff.Retry(operation, b)

	if err != nil {
		glog.Fatalf("Failed to create ArgoClient. Error: %v", err)
	}
	return &ArgoClient{argoProjClient}
}

func main() {
	// flags
	var sleepTime int
	flag.IntVar(&sleepTime, "sleepTime", 10, "Time in minutes to sleep between the first and second call.")
	flag.Parse()
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	argoClient := NewArgoClientOrFatal(60 * time.Second)
	namespace := "default"
	argoClient.Workflow(namespace)
	worflows, err := argoClient.Workflow(namespace).List(v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d workflows in the namespace %s \n", len(worflows.Items), namespace)
	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the namespace %s \n", len(pods.Items), namespace)
	// sleep for x minutes
	fmt.Printf("Sleeping for %d minutes\n", sleepTime)
	time.Sleep(time.Duration(sleepTime*60) * time.Second)

	fmt.Printf("Calling the api a second time \n")
	worflows, err = argoClient.Workflow(namespace).List(v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d workflows in the namespace %s \n", len(worflows.Items), namespace)
	pods, err = clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the namespace %s \n", len(pods.Items), namespace)
}
