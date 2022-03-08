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

type Downtime struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DowntimeSpec   `json:"spec,omitempty"`
	Status            DowntimeStatus `json:"status,omitempty"`
}

type DowntimeSpecRecurrence struct {
	// How often to repeat as an integer. For example to repeat every 3 days, select a `type` of `days` and a `period` of `3`.
	// +optional
	Period *int64 `json:"period,omitempty" tf:"period"`
	// The RRULE standard for defining recurring events. For example, to have a recurring event on the first day of each month, use `FREQ=MONTHLY;INTERVAL=1`. Most common rrule options from the iCalendar Spec are supported. Attributes specifying the duration in RRULE are not supported (for example, `DTSTART`, `DTEND`, `DURATION`).
	// +optional
	Rrule *string `json:"rrule,omitempty" tf:"rrule"`
	// One of `days`, `weeks`, `months`, or `years`
	Type *string `json:"type" tf:"type"`
	// The date at which the recurrence should end as a POSIX timestamp. `until_occurrences` and `until_date` are mutually exclusive.
	// +optional
	UntilDate *int64 `json:"untilDate,omitempty" tf:"until_date"`
	// How many times the downtime will be rescheduled. `until_occurrences` and `until_date` are mutually exclusive.
	// +optional
	UntilOccurrences *int64 `json:"untilOccurrences,omitempty" tf:"until_occurrences"`
	// A list of week days to repeat on. Choose from: `Mon`, `Tue`, `Wed`, `Thu`, `Fri`, `Sat` or `Sun`. Only applicable when `type` is `weeks`. First letter must be capitalized.
	// +optional
	WeekDays []string `json:"weekDays,omitempty" tf:"week_days"`
}

type DowntimeSpec struct {
	State *DowntimeSpecResource `json:"state,omitempty" tf:"-"`

	Resource DowntimeSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type DowntimeSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// When true indicates this downtime is being actively applied
	// +optional
	Active *bool `json:"active,omitempty" tf:"active"`
	// The id corresponding to the downtime object definition of the active child for the original parent recurring downtime. This field will only exist on recurring downtimes.
	// +optional
	ActiveChildID *int64 `json:"activeChildID,omitempty" tf:"active_child_id"`
	// When true indicates this downtime is not being applied
	// +optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled"`
	// Optionally specify an end date when this downtime should expire
	// +optional
	End *int64 `json:"end,omitempty" tf:"end"`
	// String representing date and time to end the downtime in RFC3339 format.
	// +optional
	EndDate *string `json:"endDate,omitempty" tf:"end_date"`
	// An optional message to provide when creating the downtime, can include notification handles
	// +optional
	Message *string `json:"message,omitempty" tf:"message"`
	// When specified, this downtime will only apply to this monitor
	// +optional
	MonitorID *int64 `json:"monitorID,omitempty" tf:"monitor_id"`
	// A list of monitor tags (up to 32) to base the scheduled downtime on. Only monitors that have all selected tags are silenced
	// +optional
	MonitorTags []string `json:"monitorTags,omitempty" tf:"monitor_tags"`
	// Optional recurring schedule for this downtime
	// +optional
	Recurrence *DowntimeSpecRecurrence `json:"recurrence,omitempty" tf:"recurrence"`
	// specify the group scope to which this downtime applies. For everything use '*'
	Scope []string `json:"scope" tf:"scope"`
	// Specify when this downtime should start
	// +optional
	Start *int64 `json:"start,omitempty" tf:"start"`
	// String representing date and time to start the downtime in RFC3339 format.
	// +optional
	StartDate *string `json:"startDate,omitempty" tf:"start_date"`
	// The timezone for the downtime, default UTC
	// +optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone"`
}

type DowntimeStatus struct {
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

// DowntimeList is a list of Downtimes
type DowntimeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Downtime CRD objects
	Items []Downtime `json:"items,omitempty"`
}
