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
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LevelObjective) DeepCopyInto(out *LevelObjective) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LevelObjective.
func (in *LevelObjective) DeepCopy() *LevelObjective {
	if in == nil {
		return nil
	}
	out := new(LevelObjective)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LevelObjective) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LevelObjectiveList) DeepCopyInto(out *LevelObjectiveList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LevelObjective, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LevelObjectiveList.
func (in *LevelObjectiveList) DeepCopy() *LevelObjectiveList {
	if in == nil {
		return nil
	}
	out := new(LevelObjectiveList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LevelObjectiveList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LevelObjectiveSpec) DeepCopyInto(out *LevelObjectiveSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(LevelObjectiveSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LevelObjectiveSpec.
func (in *LevelObjectiveSpec) DeepCopy() *LevelObjectiveSpec {
	if in == nil {
		return nil
	}
	out := new(LevelObjectiveSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LevelObjectiveSpecQuery) DeepCopyInto(out *LevelObjectiveSpecQuery) {
	*out = *in
	if in.Denominator != nil {
		in, out := &in.Denominator, &out.Denominator
		*out = new(string)
		**out = **in
	}
	if in.Numerator != nil {
		in, out := &in.Numerator, &out.Numerator
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LevelObjectiveSpecQuery.
func (in *LevelObjectiveSpecQuery) DeepCopy() *LevelObjectiveSpecQuery {
	if in == nil {
		return nil
	}
	out := new(LevelObjectiveSpecQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LevelObjectiveSpecResource) DeepCopyInto(out *LevelObjectiveSpecResource) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ForceDelete != nil {
		in, out := &in.ForceDelete, &out.ForceDelete
		*out = new(bool)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MonitorIDS != nil {
		in, out := &in.MonitorIDS, &out.MonitorIDS
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(LevelObjectiveSpecQuery)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Thresholds != nil {
		in, out := &in.Thresholds, &out.Thresholds
		*out = make([]LevelObjectiveSpecThresholds, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LevelObjectiveSpecResource.
func (in *LevelObjectiveSpecResource) DeepCopy() *LevelObjectiveSpecResource {
	if in == nil {
		return nil
	}
	out := new(LevelObjectiveSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LevelObjectiveSpecThresholds) DeepCopyInto(out *LevelObjectiveSpecThresholds) {
	*out = *in
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(float64)
		**out = **in
	}
	if in.TargetDisplay != nil {
		in, out := &in.TargetDisplay, &out.TargetDisplay
		*out = new(string)
		**out = **in
	}
	if in.Timeframe != nil {
		in, out := &in.Timeframe, &out.Timeframe
		*out = new(string)
		**out = **in
	}
	if in.Warning != nil {
		in, out := &in.Warning, &out.Warning
		*out = new(float64)
		**out = **in
	}
	if in.WarningDisplay != nil {
		in, out := &in.WarningDisplay, &out.WarningDisplay
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LevelObjectiveSpecThresholds.
func (in *LevelObjectiveSpecThresholds) DeepCopy() *LevelObjectiveSpecThresholds {
	if in == nil {
		return nil
	}
	out := new(LevelObjectiveSpecThresholds)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LevelObjectiveStatus) DeepCopyInto(out *LevelObjectiveStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LevelObjectiveStatus.
func (in *LevelObjectiveStatus) DeepCopy() *LevelObjectiveStatus {
	if in == nil {
		return nil
	}
	out := new(LevelObjectiveStatus)
	in.DeepCopyInto(out)
	return out
}
