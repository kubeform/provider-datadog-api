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

package versioned

import (
	"fmt"

	dashboardv1alpha1 "kubeform.dev/provider-datadog-api/client/clientset/versioned/typed/dashboard/v1alpha1"
	downtimev1alpha1 "kubeform.dev/provider-datadog-api/client/clientset/versioned/typed/downtime/v1alpha1"
	integrationv1alpha1 "kubeform.dev/provider-datadog-api/client/clientset/versioned/typed/integration/v1alpha1"
	logsv1alpha1 "kubeform.dev/provider-datadog-api/client/clientset/versioned/typed/logs/v1alpha1"
	metricv1alpha1 "kubeform.dev/provider-datadog-api/client/clientset/versioned/typed/metric/v1alpha1"
	monitorv1alpha1 "kubeform.dev/provider-datadog-api/client/clientset/versioned/typed/monitor/v1alpha1"
	rolev1alpha1 "kubeform.dev/provider-datadog-api/client/clientset/versioned/typed/role/v1alpha1"
	securityv1alpha1 "kubeform.dev/provider-datadog-api/client/clientset/versioned/typed/security/v1alpha1"
	servicev1alpha1 "kubeform.dev/provider-datadog-api/client/clientset/versioned/typed/service/v1alpha1"
	slov1alpha1 "kubeform.dev/provider-datadog-api/client/clientset/versioned/typed/slo/v1alpha1"
	syntheticsv1alpha1 "kubeform.dev/provider-datadog-api/client/clientset/versioned/typed/synthetics/v1alpha1"
	userv1alpha1 "kubeform.dev/provider-datadog-api/client/clientset/versioned/typed/user/v1alpha1"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	DashboardV1alpha1() dashboardv1alpha1.DashboardV1alpha1Interface
	DowntimeV1alpha1() downtimev1alpha1.DowntimeV1alpha1Interface
	IntegrationV1alpha1() integrationv1alpha1.IntegrationV1alpha1Interface
	LogsV1alpha1() logsv1alpha1.LogsV1alpha1Interface
	MetricV1alpha1() metricv1alpha1.MetricV1alpha1Interface
	MonitorV1alpha1() monitorv1alpha1.MonitorV1alpha1Interface
	RoleV1alpha1() rolev1alpha1.RoleV1alpha1Interface
	SecurityV1alpha1() securityv1alpha1.SecurityV1alpha1Interface
	ServiceV1alpha1() servicev1alpha1.ServiceV1alpha1Interface
	SloV1alpha1() slov1alpha1.SloV1alpha1Interface
	SyntheticsV1alpha1() syntheticsv1alpha1.SyntheticsV1alpha1Interface
	UserV1alpha1() userv1alpha1.UserV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	dashboardV1alpha1   *dashboardv1alpha1.DashboardV1alpha1Client
	downtimeV1alpha1    *downtimev1alpha1.DowntimeV1alpha1Client
	integrationV1alpha1 *integrationv1alpha1.IntegrationV1alpha1Client
	logsV1alpha1        *logsv1alpha1.LogsV1alpha1Client
	metricV1alpha1      *metricv1alpha1.MetricV1alpha1Client
	monitorV1alpha1     *monitorv1alpha1.MonitorV1alpha1Client
	roleV1alpha1        *rolev1alpha1.RoleV1alpha1Client
	securityV1alpha1    *securityv1alpha1.SecurityV1alpha1Client
	serviceV1alpha1     *servicev1alpha1.ServiceV1alpha1Client
	sloV1alpha1         *slov1alpha1.SloV1alpha1Client
	syntheticsV1alpha1  *syntheticsv1alpha1.SyntheticsV1alpha1Client
	userV1alpha1        *userv1alpha1.UserV1alpha1Client
}

// DashboardV1alpha1 retrieves the DashboardV1alpha1Client
func (c *Clientset) DashboardV1alpha1() dashboardv1alpha1.DashboardV1alpha1Interface {
	return c.dashboardV1alpha1
}

// DowntimeV1alpha1 retrieves the DowntimeV1alpha1Client
func (c *Clientset) DowntimeV1alpha1() downtimev1alpha1.DowntimeV1alpha1Interface {
	return c.downtimeV1alpha1
}

// IntegrationV1alpha1 retrieves the IntegrationV1alpha1Client
func (c *Clientset) IntegrationV1alpha1() integrationv1alpha1.IntegrationV1alpha1Interface {
	return c.integrationV1alpha1
}

// LogsV1alpha1 retrieves the LogsV1alpha1Client
func (c *Clientset) LogsV1alpha1() logsv1alpha1.LogsV1alpha1Interface {
	return c.logsV1alpha1
}

// MetricV1alpha1 retrieves the MetricV1alpha1Client
func (c *Clientset) MetricV1alpha1() metricv1alpha1.MetricV1alpha1Interface {
	return c.metricV1alpha1
}

// MonitorV1alpha1 retrieves the MonitorV1alpha1Client
func (c *Clientset) MonitorV1alpha1() monitorv1alpha1.MonitorV1alpha1Interface {
	return c.monitorV1alpha1
}

// RoleV1alpha1 retrieves the RoleV1alpha1Client
func (c *Clientset) RoleV1alpha1() rolev1alpha1.RoleV1alpha1Interface {
	return c.roleV1alpha1
}

// SecurityV1alpha1 retrieves the SecurityV1alpha1Client
func (c *Clientset) SecurityV1alpha1() securityv1alpha1.SecurityV1alpha1Interface {
	return c.securityV1alpha1
}

// ServiceV1alpha1 retrieves the ServiceV1alpha1Client
func (c *Clientset) ServiceV1alpha1() servicev1alpha1.ServiceV1alpha1Interface {
	return c.serviceV1alpha1
}

// SloV1alpha1 retrieves the SloV1alpha1Client
func (c *Clientset) SloV1alpha1() slov1alpha1.SloV1alpha1Interface {
	return c.sloV1alpha1
}

// SyntheticsV1alpha1 retrieves the SyntheticsV1alpha1Client
func (c *Clientset) SyntheticsV1alpha1() syntheticsv1alpha1.SyntheticsV1alpha1Interface {
	return c.syntheticsV1alpha1
}

// UserV1alpha1 retrieves the UserV1alpha1Client
func (c *Clientset) UserV1alpha1() userv1alpha1.UserV1alpha1Interface {
	return c.userV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.dashboardV1alpha1, err = dashboardv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.downtimeV1alpha1, err = downtimev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.integrationV1alpha1, err = integrationv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.logsV1alpha1, err = logsv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.metricV1alpha1, err = metricv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.monitorV1alpha1, err = monitorv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.roleV1alpha1, err = rolev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.securityV1alpha1, err = securityv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.serviceV1alpha1, err = servicev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.sloV1alpha1, err = slov1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.syntheticsV1alpha1, err = syntheticsv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.userV1alpha1, err = userv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.dashboardV1alpha1 = dashboardv1alpha1.NewForConfigOrDie(c)
	cs.downtimeV1alpha1 = downtimev1alpha1.NewForConfigOrDie(c)
	cs.integrationV1alpha1 = integrationv1alpha1.NewForConfigOrDie(c)
	cs.logsV1alpha1 = logsv1alpha1.NewForConfigOrDie(c)
	cs.metricV1alpha1 = metricv1alpha1.NewForConfigOrDie(c)
	cs.monitorV1alpha1 = monitorv1alpha1.NewForConfigOrDie(c)
	cs.roleV1alpha1 = rolev1alpha1.NewForConfigOrDie(c)
	cs.securityV1alpha1 = securityv1alpha1.NewForConfigOrDie(c)
	cs.serviceV1alpha1 = servicev1alpha1.NewForConfigOrDie(c)
	cs.sloV1alpha1 = slov1alpha1.NewForConfigOrDie(c)
	cs.syntheticsV1alpha1 = syntheticsv1alpha1.NewForConfigOrDie(c)
	cs.userV1alpha1 = userv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.dashboardV1alpha1 = dashboardv1alpha1.New(c)
	cs.downtimeV1alpha1 = downtimev1alpha1.New(c)
	cs.integrationV1alpha1 = integrationv1alpha1.New(c)
	cs.logsV1alpha1 = logsv1alpha1.New(c)
	cs.metricV1alpha1 = metricv1alpha1.New(c)
	cs.monitorV1alpha1 = monitorv1alpha1.New(c)
	cs.roleV1alpha1 = rolev1alpha1.New(c)
	cs.securityV1alpha1 = securityv1alpha1.New(c)
	cs.serviceV1alpha1 = servicev1alpha1.New(c)
	cs.sloV1alpha1 = slov1alpha1.New(c)
	cs.syntheticsV1alpha1 = syntheticsv1alpha1.New(c)
	cs.userV1alpha1 = userv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
