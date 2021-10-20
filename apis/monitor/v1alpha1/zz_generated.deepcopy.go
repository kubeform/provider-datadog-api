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
func (in *Monitor) DeepCopyInto(out *Monitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Monitor.
func (in *Monitor) DeepCopy() *Monitor {
	if in == nil {
		return nil
	}
	out := new(Monitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Monitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorList) DeepCopyInto(out *MonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Monitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorList.
func (in *MonitorList) DeepCopy() *MonitorList {
	if in == nil {
		return nil
	}
	out := new(MonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorSpec) DeepCopyInto(out *MonitorSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(MonitorSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorSpec.
func (in *MonitorSpec) DeepCopy() *MonitorSpec {
	if in == nil {
		return nil
	}
	out := new(MonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorSpecMonitorThresholdWindows) DeepCopyInto(out *MonitorSpecMonitorThresholdWindows) {
	*out = *in
	if in.RecoveryWindow != nil {
		in, out := &in.RecoveryWindow, &out.RecoveryWindow
		*out = new(string)
		**out = **in
	}
	if in.TriggerWindow != nil {
		in, out := &in.TriggerWindow, &out.TriggerWindow
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorSpecMonitorThresholdWindows.
func (in *MonitorSpecMonitorThresholdWindows) DeepCopy() *MonitorSpecMonitorThresholdWindows {
	if in == nil {
		return nil
	}
	out := new(MonitorSpecMonitorThresholdWindows)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorSpecMonitorThresholds) DeepCopyInto(out *MonitorSpecMonitorThresholds) {
	*out = *in
	if in.Critical != nil {
		in, out := &in.Critical, &out.Critical
		*out = new(string)
		**out = **in
	}
	if in.CriticalRecovery != nil {
		in, out := &in.CriticalRecovery, &out.CriticalRecovery
		*out = new(string)
		**out = **in
	}
	if in.Ok != nil {
		in, out := &in.Ok, &out.Ok
		*out = new(string)
		**out = **in
	}
	if in.Unknown != nil {
		in, out := &in.Unknown, &out.Unknown
		*out = new(string)
		**out = **in
	}
	if in.Warning != nil {
		in, out := &in.Warning, &out.Warning
		*out = new(string)
		**out = **in
	}
	if in.WarningRecovery != nil {
		in, out := &in.WarningRecovery, &out.WarningRecovery
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorSpecMonitorThresholds.
func (in *MonitorSpecMonitorThresholds) DeepCopy() *MonitorSpecMonitorThresholds {
	if in == nil {
		return nil
	}
	out := new(MonitorSpecMonitorThresholds)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorSpecResource) DeepCopyInto(out *MonitorSpecResource) {
	*out = *in
	if in.EnableLogsSample != nil {
		in, out := &in.EnableLogsSample, &out.EnableLogsSample
		*out = new(bool)
		**out = **in
	}
	if in.EscalationMessage != nil {
		in, out := &in.EscalationMessage, &out.EscalationMessage
		*out = new(string)
		**out = **in
	}
	if in.EvaluationDelay != nil {
		in, out := &in.EvaluationDelay, &out.EvaluationDelay
		*out = new(int64)
		**out = **in
	}
	if in.ForceDelete != nil {
		in, out := &in.ForceDelete, &out.ForceDelete
		*out = new(bool)
		**out = **in
	}
	if in.GroupbySimpleMonitor != nil {
		in, out := &in.GroupbySimpleMonitor, &out.GroupbySimpleMonitor
		*out = new(bool)
		**out = **in
	}
	if in.IncludeTags != nil {
		in, out := &in.IncludeTags, &out.IncludeTags
		*out = new(bool)
		**out = **in
	}
	if in.Locked != nil {
		in, out := &in.Locked, &out.Locked
		*out = new(bool)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.MonitorThresholdWindows != nil {
		in, out := &in.MonitorThresholdWindows, &out.MonitorThresholdWindows
		*out = new(MonitorSpecMonitorThresholdWindows)
		(*in).DeepCopyInto(*out)
	}
	if in.MonitorThresholds != nil {
		in, out := &in.MonitorThresholds, &out.MonitorThresholds
		*out = new(MonitorSpecMonitorThresholds)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NewHostDelay != nil {
		in, out := &in.NewHostDelay, &out.NewHostDelay
		*out = new(int64)
		**out = **in
	}
	if in.NoDataTimeframe != nil {
		in, out := &in.NoDataTimeframe, &out.NoDataTimeframe
		*out = new(int64)
		**out = **in
	}
	if in.NotifyAudit != nil {
		in, out := &in.NotifyAudit, &out.NotifyAudit
		*out = new(bool)
		**out = **in
	}
	if in.NotifyNoData != nil {
		in, out := &in.NotifyNoData, &out.NotifyNoData
		*out = new(bool)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int64)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(string)
		**out = **in
	}
	if in.RenotifyInterval != nil {
		in, out := &in.RenotifyInterval, &out.RenotifyInterval
		*out = new(int64)
		**out = **in
	}
	if in.RequireFullWindow != nil {
		in, out := &in.RequireFullWindow, &out.RequireFullWindow
		*out = new(bool)
		**out = **in
	}
	if in.RestrictedRoles != nil {
		in, out := &in.RestrictedRoles, &out.RestrictedRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TimeoutH != nil {
		in, out := &in.TimeoutH, &out.TimeoutH
		*out = new(int64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Validate != nil {
		in, out := &in.Validate, &out.Validate
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorSpecResource.
func (in *MonitorSpecResource) DeepCopy() *MonitorSpecResource {
	if in == nil {
		return nil
	}
	out := new(MonitorSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorStatus) DeepCopyInto(out *MonitorStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorStatus.
func (in *MonitorStatus) DeepCopy() *MonitorStatus {
	if in == nil {
		return nil
	}
	out := new(MonitorStatus)
	in.DeepCopyInto(out)
	return out
}
