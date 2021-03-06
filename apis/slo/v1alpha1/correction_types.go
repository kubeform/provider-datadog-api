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

type Correction struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CorrectionSpec   `json:"spec,omitempty"`
	Status            CorrectionStatus `json:"status,omitempty"`
}

type CorrectionSpec struct {
	State *CorrectionSpecResource `json:"state,omitempty" tf:"-"`

	Resource CorrectionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type CorrectionSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Category the SLO correction belongs to.
	Category *string `json:"category" tf:"category"`
	// Description of the correction being made.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Length of time in seconds for a specified `rrule` recurring SLO correction (required if specifying `rrule`)
	// +optional
	Duration *int64 `json:"duration,omitempty" tf:"duration"`
	// Ending time of the correction in epoch seconds. Required for one time corrections, but optional if `rrule` is specified
	// +optional
	End *int64 `json:"end,omitempty" tf:"end"`
	// Recurrence rules as defined in the iCalendar RFC 5545. Supported rules for SLO corrections are `FREQ`, `INTERVAL`, `COUNT` and `UNTIL`.
	// +optional
	Rrule *string `json:"rrule,omitempty" tf:"rrule"`
	// ID of the SLO that this correction will be applied to.
	SloID *string `json:"sloID" tf:"slo_id"`
	// Starting time of the correction in epoch seconds.
	Start *int64 `json:"start" tf:"start"`
	// The timezone to display in the UI for the correction times (defaults to "UTC")
	// +optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone"`
}

type CorrectionStatus struct {
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

// CorrectionList is a list of Corrections
type CorrectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Correction CRD objects
	Items []Correction `json:"items,omitempty"`
}
