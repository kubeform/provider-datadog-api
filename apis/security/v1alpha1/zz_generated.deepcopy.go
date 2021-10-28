//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringDefaultRule) DeepCopyInto(out *MonitoringDefaultRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringDefaultRule.
func (in *MonitoringDefaultRule) DeepCopy() *MonitoringDefaultRule {
	if in == nil {
		return nil
	}
	out := new(MonitoringDefaultRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringDefaultRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringDefaultRuleList) DeepCopyInto(out *MonitoringDefaultRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MonitoringDefaultRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringDefaultRuleList.
func (in *MonitoringDefaultRuleList) DeepCopy() *MonitoringDefaultRuleList {
	if in == nil {
		return nil
	}
	out := new(MonitoringDefaultRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringDefaultRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringDefaultRuleSpec) DeepCopyInto(out *MonitoringDefaultRuleSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(MonitoringDefaultRuleSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringDefaultRuleSpec.
func (in *MonitoringDefaultRuleSpec) DeepCopy() *MonitoringDefaultRuleSpec {
	if in == nil {
		return nil
	}
	out := new(MonitoringDefaultRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringDefaultRuleSpecCase) DeepCopyInto(out *MonitoringDefaultRuleSpecCase) {
	*out = *in
	if in.Notifications != nil {
		in, out := &in.Notifications, &out.Notifications
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringDefaultRuleSpecCase.
func (in *MonitoringDefaultRuleSpecCase) DeepCopy() *MonitoringDefaultRuleSpecCase {
	if in == nil {
		return nil
	}
	out := new(MonitoringDefaultRuleSpecCase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringDefaultRuleSpecFilter) DeepCopyInto(out *MonitoringDefaultRuleSpecFilter) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringDefaultRuleSpecFilter.
func (in *MonitoringDefaultRuleSpecFilter) DeepCopy() *MonitoringDefaultRuleSpecFilter {
	if in == nil {
		return nil
	}
	out := new(MonitoringDefaultRuleSpecFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringDefaultRuleSpecResource) DeepCopyInto(out *MonitoringDefaultRuleSpecResource) {
	*out = *in
	if in.Case != nil {
		in, out := &in.Case, &out.Case
		*out = make([]MonitoringDefaultRuleSpecCase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = make([]MonitoringDefaultRuleSpecFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringDefaultRuleSpecResource.
func (in *MonitoringDefaultRuleSpecResource) DeepCopy() *MonitoringDefaultRuleSpecResource {
	if in == nil {
		return nil
	}
	out := new(MonitoringDefaultRuleSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringDefaultRuleStatus) DeepCopyInto(out *MonitoringDefaultRuleStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringDefaultRuleStatus.
func (in *MonitoringDefaultRuleStatus) DeepCopy() *MonitoringDefaultRuleStatus {
	if in == nil {
		return nil
	}
	out := new(MonitoringDefaultRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringFilter) DeepCopyInto(out *MonitoringFilter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringFilter.
func (in *MonitoringFilter) DeepCopy() *MonitoringFilter {
	if in == nil {
		return nil
	}
	out := new(MonitoringFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringFilter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringFilterList) DeepCopyInto(out *MonitoringFilterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MonitoringFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringFilterList.
func (in *MonitoringFilterList) DeepCopy() *MonitoringFilterList {
	if in == nil {
		return nil
	}
	out := new(MonitoringFilterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringFilterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringFilterSpec) DeepCopyInto(out *MonitoringFilterSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(MonitoringFilterSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringFilterSpec.
func (in *MonitoringFilterSpec) DeepCopy() *MonitoringFilterSpec {
	if in == nil {
		return nil
	}
	out := new(MonitoringFilterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringFilterSpecExclusionFilter) DeepCopyInto(out *MonitoringFilterSpecExclusionFilter) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringFilterSpecExclusionFilter.
func (in *MonitoringFilterSpecExclusionFilter) DeepCopy() *MonitoringFilterSpecExclusionFilter {
	if in == nil {
		return nil
	}
	out := new(MonitoringFilterSpecExclusionFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringFilterSpecResource) DeepCopyInto(out *MonitoringFilterSpecResource) {
	*out = *in
	if in.ExclusionFilter != nil {
		in, out := &in.ExclusionFilter, &out.ExclusionFilter
		*out = make([]MonitoringFilterSpecExclusionFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FilteredDataType != nil {
		in, out := &in.FilteredDataType, &out.FilteredDataType
		*out = new(string)
		**out = **in
	}
	if in.IsEnabled != nil {
		in, out := &in.IsEnabled, &out.IsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringFilterSpecResource.
func (in *MonitoringFilterSpecResource) DeepCopy() *MonitoringFilterSpecResource {
	if in == nil {
		return nil
	}
	out := new(MonitoringFilterSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringFilterStatus) DeepCopyInto(out *MonitoringFilterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringFilterStatus.
func (in *MonitoringFilterStatus) DeepCopy() *MonitoringFilterStatus {
	if in == nil {
		return nil
	}
	out := new(MonitoringFilterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringRule) DeepCopyInto(out *MonitoringRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringRule.
func (in *MonitoringRule) DeepCopy() *MonitoringRule {
	if in == nil {
		return nil
	}
	out := new(MonitoringRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringRuleList) DeepCopyInto(out *MonitoringRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MonitoringRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringRuleList.
func (in *MonitoringRuleList) DeepCopy() *MonitoringRuleList {
	if in == nil {
		return nil
	}
	out := new(MonitoringRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringRuleSpec) DeepCopyInto(out *MonitoringRuleSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(MonitoringRuleSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringRuleSpec.
func (in *MonitoringRuleSpec) DeepCopy() *MonitoringRuleSpec {
	if in == nil {
		return nil
	}
	out := new(MonitoringRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringRuleSpecCase) DeepCopyInto(out *MonitoringRuleSpecCase) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Notifications != nil {
		in, out := &in.Notifications, &out.Notifications
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringRuleSpecCase.
func (in *MonitoringRuleSpecCase) DeepCopy() *MonitoringRuleSpecCase {
	if in == nil {
		return nil
	}
	out := new(MonitoringRuleSpecCase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringRuleSpecFilter) DeepCopyInto(out *MonitoringRuleSpecFilter) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringRuleSpecFilter.
func (in *MonitoringRuleSpecFilter) DeepCopy() *MonitoringRuleSpecFilter {
	if in == nil {
		return nil
	}
	out := new(MonitoringRuleSpecFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringRuleSpecOptions) DeepCopyInto(out *MonitoringRuleSpecOptions) {
	*out = *in
	if in.DetectionMethod != nil {
		in, out := &in.DetectionMethod, &out.DetectionMethod
		*out = new(string)
		**out = **in
	}
	if in.EvaluationWindow != nil {
		in, out := &in.EvaluationWindow, &out.EvaluationWindow
		*out = new(int64)
		**out = **in
	}
	if in.KeepAlive != nil {
		in, out := &in.KeepAlive, &out.KeepAlive
		*out = new(int64)
		**out = **in
	}
	if in.MaxSignalDuration != nil {
		in, out := &in.MaxSignalDuration, &out.MaxSignalDuration
		*out = new(int64)
		**out = **in
	}
	if in.NewValueOptions != nil {
		in, out := &in.NewValueOptions, &out.NewValueOptions
		*out = new(MonitoringRuleSpecOptionsNewValueOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringRuleSpecOptions.
func (in *MonitoringRuleSpecOptions) DeepCopy() *MonitoringRuleSpecOptions {
	if in == nil {
		return nil
	}
	out := new(MonitoringRuleSpecOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringRuleSpecOptionsNewValueOptions) DeepCopyInto(out *MonitoringRuleSpecOptionsNewValueOptions) {
	*out = *in
	if in.ForgetAfter != nil {
		in, out := &in.ForgetAfter, &out.ForgetAfter
		*out = new(int64)
		**out = **in
	}
	if in.LearningDuration != nil {
		in, out := &in.LearningDuration, &out.LearningDuration
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringRuleSpecOptionsNewValueOptions.
func (in *MonitoringRuleSpecOptionsNewValueOptions) DeepCopy() *MonitoringRuleSpecOptionsNewValueOptions {
	if in == nil {
		return nil
	}
	out := new(MonitoringRuleSpecOptionsNewValueOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringRuleSpecQuery) DeepCopyInto(out *MonitoringRuleSpecQuery) {
	*out = *in
	if in.AgentRule != nil {
		in, out := &in.AgentRule, &out.AgentRule
		*out = make([]MonitoringRuleSpecQueryAgentRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Aggregation != nil {
		in, out := &in.Aggregation, &out.Aggregation
		*out = new(string)
		**out = **in
	}
	if in.DistinctFields != nil {
		in, out := &in.DistinctFields, &out.DistinctFields
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.GroupByFields != nil {
		in, out := &in.GroupByFields, &out.GroupByFields
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Metric != nil {
		in, out := &in.Metric, &out.Metric
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringRuleSpecQuery.
func (in *MonitoringRuleSpecQuery) DeepCopy() *MonitoringRuleSpecQuery {
	if in == nil {
		return nil
	}
	out := new(MonitoringRuleSpecQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringRuleSpecQueryAgentRule) DeepCopyInto(out *MonitoringRuleSpecQueryAgentRule) {
	*out = *in
	if in.AgentRuleID != nil {
		in, out := &in.AgentRuleID, &out.AgentRuleID
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringRuleSpecQueryAgentRule.
func (in *MonitoringRuleSpecQueryAgentRule) DeepCopy() *MonitoringRuleSpecQueryAgentRule {
	if in == nil {
		return nil
	}
	out := new(MonitoringRuleSpecQueryAgentRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringRuleSpecResource) DeepCopyInto(out *MonitoringRuleSpecResource) {
	*out = *in
	if in.Case != nil {
		in, out := &in.Case, &out.Case
		*out = make([]MonitoringRuleSpecCase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = make([]MonitoringRuleSpecFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HasExtendedTitle != nil {
		in, out := &in.HasExtendedTitle, &out.HasExtendedTitle
		*out = new(bool)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(MonitoringRuleSpecOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = make([]MonitoringRuleSpecQuery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringRuleSpecResource.
func (in *MonitoringRuleSpecResource) DeepCopy() *MonitoringRuleSpecResource {
	if in == nil {
		return nil
	}
	out := new(MonitoringRuleSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringRuleStatus) DeepCopyInto(out *MonitoringRuleStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringRuleStatus.
func (in *MonitoringRuleStatus) DeepCopy() *MonitoringRuleStatus {
	if in == nil {
		return nil
	}
	out := new(MonitoringRuleStatus)
	in.DeepCopyInto(out)
	return out
}
