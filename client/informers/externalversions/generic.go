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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "kubeform.dev/provider-datadog-api/apis/apikey/v1alpha1"
	applicationv1alpha1 "kubeform.dev/provider-datadog-api/apis/application/v1alpha1"
	childv1alpha1 "kubeform.dev/provider-datadog-api/apis/child/v1alpha1"
	dashboardv1alpha1 "kubeform.dev/provider-datadog-api/apis/dashboard/v1alpha1"
	downtimev1alpha1 "kubeform.dev/provider-datadog-api/apis/downtime/v1alpha1"
	integrationv1alpha1 "kubeform.dev/provider-datadog-api/apis/integration/v1alpha1"
	logsv1alpha1 "kubeform.dev/provider-datadog-api/apis/logs/v1alpha1"
	metricv1alpha1 "kubeform.dev/provider-datadog-api/apis/metric/v1alpha1"
	monitorv1alpha1 "kubeform.dev/provider-datadog-api/apis/monitor/v1alpha1"
	organizationv1alpha1 "kubeform.dev/provider-datadog-api/apis/organization/v1alpha1"
	rolev1alpha1 "kubeform.dev/provider-datadog-api/apis/role/v1alpha1"
	securityv1alpha1 "kubeform.dev/provider-datadog-api/apis/security/v1alpha1"
	servicev1alpha1 "kubeform.dev/provider-datadog-api/apis/service/v1alpha1"
	slov1alpha1 "kubeform.dev/provider-datadog-api/apis/slo/v1alpha1"
	syntheticsv1alpha1 "kubeform.dev/provider-datadog-api/apis/synthetics/v1alpha1"
	userv1alpha1 "kubeform.dev/provider-datadog-api/apis/user/v1alpha1"
	webhookv1alpha1 "kubeform.dev/provider-datadog-api/apis/webhook/v1alpha1"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=apikey.datadog.kubeform.com, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("apikeys"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apikey().V1alpha1().ApiKeys().Informer()}, nil

		// Group=application.datadog.kubeform.com, Version=v1alpha1
	case applicationv1alpha1.SchemeGroupVersion.WithResource("keys"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Application().V1alpha1().Keys().Informer()}, nil

		// Group=child.datadog.kubeform.com, Version=v1alpha1
	case childv1alpha1.SchemeGroupVersion.WithResource("organizations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Child().V1alpha1().Organizations().Informer()}, nil

		// Group=dashboard.datadog.kubeform.com, Version=v1alpha1
	case dashboardv1alpha1.SchemeGroupVersion.WithResource("dashboards"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dashboard().V1alpha1().Dashboards().Informer()}, nil
	case dashboardv1alpha1.SchemeGroupVersion.WithResource("jsons"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dashboard().V1alpha1().Jsons().Informer()}, nil
	case dashboardv1alpha1.SchemeGroupVersion.WithResource("lists"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dashboard().V1alpha1().Lists().Informer()}, nil

		// Group=downtime.datadog.kubeform.com, Version=v1alpha1
	case downtimev1alpha1.SchemeGroupVersion.WithResource("downtimes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Downtime().V1alpha1().Downtimes().Informer()}, nil

		// Group=integration.datadog.kubeform.com, Version=v1alpha1
	case integrationv1alpha1.SchemeGroupVersion.WithResource("awses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Integration().V1alpha1().Awses().Informer()}, nil
	case integrationv1alpha1.SchemeGroupVersion.WithResource("awslambdaarns"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Integration().V1alpha1().AwsLambdaArns().Informer()}, nil
	case integrationv1alpha1.SchemeGroupVersion.WithResource("awslogcollections"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Integration().V1alpha1().AwsLogCollections().Informer()}, nil
	case integrationv1alpha1.SchemeGroupVersion.WithResource("awstagfilters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Integration().V1alpha1().AwsTagFilters().Informer()}, nil
	case integrationv1alpha1.SchemeGroupVersion.WithResource("azures"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Integration().V1alpha1().Azures().Informer()}, nil
	case integrationv1alpha1.SchemeGroupVersion.WithResource("gcps"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Integration().V1alpha1().Gcps().Informer()}, nil
	case integrationv1alpha1.SchemeGroupVersion.WithResource("pagerduties"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Integration().V1alpha1().Pagerduties().Informer()}, nil
	case integrationv1alpha1.SchemeGroupVersion.WithResource("pagerdutyserviceobjects"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Integration().V1alpha1().PagerdutyServiceObjects().Informer()}, nil
	case integrationv1alpha1.SchemeGroupVersion.WithResource("slackchannels"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Integration().V1alpha1().SlackChannels().Informer()}, nil

		// Group=logs.datadog.kubeform.com, Version=v1alpha1
	case logsv1alpha1.SchemeGroupVersion.WithResource("archives"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Logs().V1alpha1().Archives().Informer()}, nil
	case logsv1alpha1.SchemeGroupVersion.WithResource("archiveorders"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Logs().V1alpha1().ArchiveOrders().Informer()}, nil
	case logsv1alpha1.SchemeGroupVersion.WithResource("custompipelines"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Logs().V1alpha1().CustomPipelines().Informer()}, nil
	case logsv1alpha1.SchemeGroupVersion.WithResource("indexes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Logs().V1alpha1().Indexes().Informer()}, nil
	case logsv1alpha1.SchemeGroupVersion.WithResource("indexorders"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Logs().V1alpha1().IndexOrders().Informer()}, nil
	case logsv1alpha1.SchemeGroupVersion.WithResource("integrationpipelines"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Logs().V1alpha1().IntegrationPipelines().Informer()}, nil
	case logsv1alpha1.SchemeGroupVersion.WithResource("metrics"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Logs().V1alpha1().Metrics().Informer()}, nil
	case logsv1alpha1.SchemeGroupVersion.WithResource("pipelineorders"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Logs().V1alpha1().PipelineOrders().Informer()}, nil

		// Group=metric.datadog.kubeform.com, Version=v1alpha1
	case metricv1alpha1.SchemeGroupVersion.WithResource("metadatas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Metric().V1alpha1().Metadatas().Informer()}, nil
	case metricv1alpha1.SchemeGroupVersion.WithResource("tagconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Metric().V1alpha1().TagConfigurations().Informer()}, nil

		// Group=monitor.datadog.kubeform.com, Version=v1alpha1
	case monitorv1alpha1.SchemeGroupVersion.WithResource("jsons"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Monitor().V1alpha1().Jsons().Informer()}, nil
	case monitorv1alpha1.SchemeGroupVersion.WithResource("monitors"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Monitor().V1alpha1().Monitors().Informer()}, nil

		// Group=organization.datadog.kubeform.com, Version=v1alpha1
	case organizationv1alpha1.SchemeGroupVersion.WithResource("settingses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Organization().V1alpha1().Settingses().Informer()}, nil

		// Group=role.datadog.kubeform.com, Version=v1alpha1
	case rolev1alpha1.SchemeGroupVersion.WithResource("roles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Role().V1alpha1().Roles().Informer()}, nil

		// Group=security.datadog.kubeform.com, Version=v1alpha1
	case securityv1alpha1.SchemeGroupVersion.WithResource("monitoringdefaultrules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Security().V1alpha1().MonitoringDefaultRules().Informer()}, nil
	case securityv1alpha1.SchemeGroupVersion.WithResource("monitoringfilters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Security().V1alpha1().MonitoringFilters().Informer()}, nil
	case securityv1alpha1.SchemeGroupVersion.WithResource("monitoringrules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Security().V1alpha1().MonitoringRules().Informer()}, nil

		// Group=service.datadog.kubeform.com, Version=v1alpha1
	case servicev1alpha1.SchemeGroupVersion.WithResource("levelobjectives"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Service().V1alpha1().LevelObjectives().Informer()}, nil

		// Group=slo.datadog.kubeform.com, Version=v1alpha1
	case slov1alpha1.SchemeGroupVersion.WithResource("corrections"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Slo().V1alpha1().Corrections().Informer()}, nil

		// Group=synthetics.datadog.kubeform.com, Version=v1alpha1
	case syntheticsv1alpha1.SchemeGroupVersion.WithResource("globalvariables"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Synthetics().V1alpha1().GlobalVariables().Informer()}, nil
	case syntheticsv1alpha1.SchemeGroupVersion.WithResource("privatelocations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Synthetics().V1alpha1().PrivateLocations().Informer()}, nil
	case syntheticsv1alpha1.SchemeGroupVersion.WithResource("syntheticstests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Synthetics().V1alpha1().Syntheticstests().Informer()}, nil

		// Group=user.datadog.kubeform.com, Version=v1alpha1
	case userv1alpha1.SchemeGroupVersion.WithResource("users"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.User().V1alpha1().Users().Informer()}, nil

		// Group=webhook.datadog.kubeform.com, Version=v1alpha1
	case webhookv1alpha1.SchemeGroupVersion.WithResource("customvariables"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Webhook().V1alpha1().CustomVariables().Informer()}, nil
	case webhookv1alpha1.SchemeGroupVersion.WithResource("webhooks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Webhook().V1alpha1().Webhooks().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
