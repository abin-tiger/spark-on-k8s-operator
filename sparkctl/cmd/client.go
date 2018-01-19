/*
Copyright 2017 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	crdclientset "k8s.io/spark-on-k8s-operator/pkg/client/clientset/versioned"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/rest"
	clientset "k8s.io/client-go/kubernetes"
)

func buildConfig(kubeConfig string) (*rest.Config, error) {
	return clientcmd.BuildConfigFromFlags("", kubeConfig)
}

func getKubeClient() (clientset.Interface, error) {
	config, err := buildConfig(KubeConfig)
	if err != nil {
		return nil, err
	}
	return clientset.NewForConfig(config)
}

func getSparkApplicationClient() (crdclientset.Interface, error) {
	config, err := buildConfig(KubeConfig)
	if err != nil {
		return nil, err
	}
	return crdclientset.NewForConfig(config)
}