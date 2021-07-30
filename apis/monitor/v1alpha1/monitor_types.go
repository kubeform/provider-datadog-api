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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Monitor struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorSpec   `json:"spec,omitempty"`
	Status            MonitorStatus `json:"status,omitempty"`
}

type MonitorSpecMonitorThresholdWindows struct {
	// Describes how long an anomalous metric must be normal before the alert recovers.
	// +optional
	RecoveryWindow *string `json:"recoveryWindow,omitempty" tf:"recovery_window"`
	// Describes how long a metric must be anomalous before an alert triggers.
	// +optional
	TriggerWindow *string `json:"triggerWindow,omitempty" tf:"trigger_window"`
}

type MonitorSpecMonitorThresholds struct {
	// The monitor `CRITICAL` threshold. Must be a number.
	// +optional
	Critical *string `json:"critical,omitempty" tf:"critical"`
	// The monitor `CRITICAL` recovery threshold. Must be a number.
	// +optional
	CriticalRecovery *string `json:"criticalRecovery,omitempty" tf:"critical_recovery"`
	// The monitor `OK` threshold. Must be a number.
	// +optional
	Ok *string `json:"ok,omitempty" tf:"ok"`
	// The monitor `UNKNOWN` threshold. Must be a number.
	// +optional
	Unknown *string `json:"unknown,omitempty" tf:"unknown"`
	// The monitor `WARNING` threshold. Must be a number.
	// +optional
	Warning *string `json:"warning,omitempty" tf:"warning"`
	// The monitor `WARNING` recovery threshold. Must be a number.
	// +optional
	WarningRecovery *string `json:"warningRecovery,omitempty" tf:"warning_recovery"`
}

type MonitorSpec struct {
	State *MonitorSpecResource `json:"state,omitempty" tf:"-"`

	Resource MonitorSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type MonitorSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// A boolean indicating whether or not to include a list of log values which triggered the alert. This is only used by log monitors. Defaults to `false`.
	// +optional
	EnableLogsSample *bool `json:"enableLogsSample,omitempty" tf:"enable_logs_sample"`
	// A message to include with a re-notification. Supports the `@username` notification allowed elsewhere.
	// +optional
	EscalationMessage *string `json:"escalationMessage,omitempty" tf:"escalation_message"`
	// (Only applies to metric alert) Time (in seconds) to delay evaluation, as a non-negative integer.
	//
	// For example, if the value is set to `300` (5min), the `timeframe` is set to `last_5m` and the time is 7:00, the monitor will evaluate data from 6:50 to 6:55. This is useful for AWS CloudWatch and other backfilled metrics to ensure the monitor will always have data during evaluation.
	// +optional
	EvaluationDelay *int64 `json:"evaluationDelay,omitempty" tf:"evaluation_delay"`
	// A boolean indicating whether this monitor can be deleted even if it’s referenced by other resources (e.g. SLO, composite monitor).
	// +optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete"`
	// Whether or not to trigger one alert if any source breaches a threshold. This is only used by log monitors. Defaults to `false`.
	// +optional
	GroupbySimpleMonitor *bool `json:"groupbySimpleMonitor,omitempty" tf:"groupby_simple_monitor"`
	// A boolean indicating whether notifications from this monitor automatically insert its triggering tags into the title. Defaults to `true`.
	// +optional
	IncludeTags *bool `json:"includeTags,omitempty" tf:"include_tags"`
	// A boolean indicating whether changes to to this monitor should be restricted to the creator or admins. Defaults to `false`.
	// +optional
	Locked *bool `json:"locked,omitempty" tf:"locked"`
	// A message to include with notifications for this monitor.
	//
	// Email notifications can be sent to specific users by using the same `@username` notation as events.
	Message *string `json:"message" tf:"message"`
	// A mapping containing `recovery_window` and `trigger_window` values, e.g. `last_15m` . Can only be used for, and are required for, anomaly monitors.
	// +optional
	MonitorThresholdWindows *MonitorSpecMonitorThresholdWindows `json:"monitorThresholdWindows,omitempty" tf:"monitor_threshold_windows"`
	// Alert thresholds of the monitor.
	// +optional
	MonitorThresholds *MonitorSpecMonitorThresholds `json:"monitorThresholds,omitempty" tf:"monitor_thresholds"`
	// Name of Datadog monitor.
	Name *string `json:"name" tf:"name"`
	// Time (in seconds) to allow a host to boot and applications to fully start before starting the evaluation of monitor results. Should be a non negative integer. Defaults to `300`.
	// +optional
	NewHostDelay *int64 `json:"newHostDelay,omitempty" tf:"new_host_delay"`
	// The number of minutes before a monitor will notify when data stops reporting. Provider defaults to 10 minutes.
	//
	// We recommend at least 2x the monitor timeframe for metric alerts or 2 minutes for service checks.
	// +optional
	NoDataTimeframe *int64 `json:"noDataTimeframe,omitempty" tf:"no_data_timeframe"`
	// A boolean indicating whether tagged users will be notified on changes to this monitor. Defaults to `false`.
	// +optional
	NotifyAudit *bool `json:"notifyAudit,omitempty" tf:"notify_audit"`
	// A boolean indicating whether this monitor will notify when data stops reporting. Defaults to `false`.
	// +optional
	NotifyNoData *bool `json:"notifyNoData,omitempty" tf:"notify_no_data"`
	// Integer from 1 (high) to 5 (low) indicating alert severity.
	// +optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`
	// The monitor query to notify on. Note this is not the same query you see in the UI and the syntax is different depending on the monitor type, please see the [API Reference](https://docs.datadoghq.com/api/v1/monitors/#create-a-monitor) for details. `terraform plan` will validate query contents unless `validate` is set to `false`.
	//
	// **Note:** APM latency data is now available as Distribution Metrics. Existing monitors have been migrated automatically but all terraformed monitors can still use the existing metrics. We strongly recommend updating monitor definitions to query the new metrics. To learn more, or to see examples of how to update your terraform definitions to utilize the new distribution metrics, see the [detailed doc](https://docs.datadoghq.com/tracing/guide/ddsketch_trace_metrics/).
	Query *string `json:"query" tf:"query"`
	// The number of minutes after the last notification before a monitor will re-notify on the current status. It will only re-notify if it's not resolved.
	// +optional
	RenotifyInterval *int64 `json:"renotifyInterval,omitempty" tf:"renotify_interval"`
	// A boolean indicating whether this monitor needs a full window of data before it's evaluated.
	//
	// We highly recommend you set this to `false` for sparse metrics, otherwise some evaluations will be skipped. Default: `true` for `on average`, `at all times` and `in total` aggregation. `false` otherwise.
	// +optional
	RequireFullWindow *bool `json:"requireFullWindow,omitempty" tf:"require_full_window"`
	// +optional
	RestrictedRoles []string `json:"restrictedRoles,omitempty" tf:"restricted_roles"`
	// A list of tags to associate with your monitor. This can help you categorize and filter monitors in the manage monitors page of the UI. Note: it's not currently possible to filter by these tags when querying via the API
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// The number of hours of the monitor not reporting data before it will automatically resolve from a triggered state.
	// +optional
	TimeoutH *int64 `json:"timeoutH,omitempty" tf:"timeout_h"`
	// The type of the monitor. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API [documentation page](https://docs.datadoghq.com/api/v1/monitors/#create-a-monitor). Note: The monitor type cannot be changed after a monitor is created.
	Type *string `json:"type" tf:"type"`
	// If set to `false`, skip the validation call done during plan.
	// +optional
	Validate *bool `json:"validate,omitempty" tf:"validate"`
}

type MonitorStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitorList is a list of Monitors
type MonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Monitor CRD objects
	Items []Monitor `json:"items,omitempty"`
}
