/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-datadog-api/apis/logs/v1alpha1"
	"kubeform.dev/provider-datadog-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type LogsV1alpha1Interface interface {
	RESTClient() rest.Interface
	ArchivesGetter
	ArchiveOrdersGetter
	CustomPipelinesGetter
	IndexesGetter
	IndexOrdersGetter
	IntegrationPipelinesGetter
	MetricsGetter
	PipelineOrdersGetter
}

// LogsV1alpha1Client is used to interact with features provided by the logs.datadog.kubeform.com group.
type LogsV1alpha1Client struct {
	restClient rest.Interface
}

func (c *LogsV1alpha1Client) Archives(namespace string) ArchiveInterface {
	return newArchives(c, namespace)
}

func (c *LogsV1alpha1Client) ArchiveOrders(namespace string) ArchiveOrderInterface {
	return newArchiveOrders(c, namespace)
}

func (c *LogsV1alpha1Client) CustomPipelines(namespace string) CustomPipelineInterface {
	return newCustomPipelines(c, namespace)
}

func (c *LogsV1alpha1Client) Indexes(namespace string) IndexInterface {
	return newIndexes(c, namespace)
}

func (c *LogsV1alpha1Client) IndexOrders(namespace string) IndexOrderInterface {
	return newIndexOrders(c, namespace)
}

func (c *LogsV1alpha1Client) IntegrationPipelines(namespace string) IntegrationPipelineInterface {
	return newIntegrationPipelines(c, namespace)
}

func (c *LogsV1alpha1Client) Metrics(namespace string) MetricInterface {
	return newMetrics(c, namespace)
}

func (c *LogsV1alpha1Client) PipelineOrders(namespace string) PipelineOrderInterface {
	return newPipelineOrders(c, namespace)
}

// NewForConfig creates a new LogsV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*LogsV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &LogsV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new LogsV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *LogsV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new LogsV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *LogsV1alpha1Client {
	return &LogsV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *LogsV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}