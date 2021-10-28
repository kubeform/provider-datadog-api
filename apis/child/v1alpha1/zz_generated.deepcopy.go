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
func (in *Organization) DeepCopyInto(out *Organization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Organization.
func (in *Organization) DeepCopy() *Organization {
	if in == nil {
		return nil
	}
	out := new(Organization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Organization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationList) DeepCopyInto(out *OrganizationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Organization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationList.
func (in *OrganizationList) DeepCopy() *OrganizationList {
	if in == nil {
		return nil
	}
	out := new(OrganizationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSpec) DeepCopyInto(out *OrganizationSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(OrganizationSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSpec.
func (in *OrganizationSpec) DeepCopy() *OrganizationSpec {
	if in == nil {
		return nil
	}
	out := new(OrganizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSpecApiKey) DeepCopyInto(out *OrganizationSpecApiKey) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSpecApiKey.
func (in *OrganizationSpecApiKey) DeepCopy() *OrganizationSpecApiKey {
	if in == nil {
		return nil
	}
	out := new(OrganizationSpecApiKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSpecApplicationKey) DeepCopyInto(out *OrganizationSpecApplicationKey) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSpecApplicationKey.
func (in *OrganizationSpecApplicationKey) DeepCopy() *OrganizationSpecApplicationKey {
	if in == nil {
		return nil
	}
	out := new(OrganizationSpecApplicationKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSpecResource) DeepCopyInto(out *OrganizationSpecResource) {
	*out = *in
	if in.ApiKey != nil {
		in, out := &in.ApiKey, &out.ApiKey
		*out = make([]OrganizationSpecApiKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ApplicationKey != nil {
		in, out := &in.ApplicationKey, &out.ApplicationKey
		*out = make([]OrganizationSpecApplicationKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PublicID != nil {
		in, out := &in.PublicID, &out.PublicID
		*out = new(string)
		**out = **in
	}
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = make([]OrganizationSpecSettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = make([]OrganizationSpecUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSpecResource.
func (in *OrganizationSpecResource) DeepCopy() *OrganizationSpecResource {
	if in == nil {
		return nil
	}
	out := new(OrganizationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSpecSettings) DeepCopyInto(out *OrganizationSpecSettings) {
	*out = *in
	if in.PrivateWidgetShare != nil {
		in, out := &in.PrivateWidgetShare, &out.PrivateWidgetShare
		*out = new(bool)
		**out = **in
	}
	if in.Saml != nil {
		in, out := &in.Saml, &out.Saml
		*out = make([]OrganizationSpecSettingsSaml, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SamlAutocreateAccessRole != nil {
		in, out := &in.SamlAutocreateAccessRole, &out.SamlAutocreateAccessRole
		*out = new(string)
		**out = **in
	}
	if in.SamlAutocreateUsersDomains != nil {
		in, out := &in.SamlAutocreateUsersDomains, &out.SamlAutocreateUsersDomains
		*out = make([]OrganizationSpecSettingsSamlAutocreateUsersDomains, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SamlCanBeEnabled != nil {
		in, out := &in.SamlCanBeEnabled, &out.SamlCanBeEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SamlIdpEndpoint != nil {
		in, out := &in.SamlIdpEndpoint, &out.SamlIdpEndpoint
		*out = new(string)
		**out = **in
	}
	if in.SamlIdpInitiatedLogin != nil {
		in, out := &in.SamlIdpInitiatedLogin, &out.SamlIdpInitiatedLogin
		*out = make([]OrganizationSpecSettingsSamlIdpInitiatedLogin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SamlIdpMetadataUploaded != nil {
		in, out := &in.SamlIdpMetadataUploaded, &out.SamlIdpMetadataUploaded
		*out = new(bool)
		**out = **in
	}
	if in.SamlLoginURL != nil {
		in, out := &in.SamlLoginURL, &out.SamlLoginURL
		*out = new(string)
		**out = **in
	}
	if in.SamlStrictMode != nil {
		in, out := &in.SamlStrictMode, &out.SamlStrictMode
		*out = make([]OrganizationSpecSettingsSamlStrictMode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSpecSettings.
func (in *OrganizationSpecSettings) DeepCopy() *OrganizationSpecSettings {
	if in == nil {
		return nil
	}
	out := new(OrganizationSpecSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSpecSettingsSaml) DeepCopyInto(out *OrganizationSpecSettingsSaml) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSpecSettingsSaml.
func (in *OrganizationSpecSettingsSaml) DeepCopy() *OrganizationSpecSettingsSaml {
	if in == nil {
		return nil
	}
	out := new(OrganizationSpecSettingsSaml)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSpecSettingsSamlAutocreateUsersDomains) DeepCopyInto(out *OrganizationSpecSettingsSamlAutocreateUsersDomains) {
	*out = *in
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSpecSettingsSamlAutocreateUsersDomains.
func (in *OrganizationSpecSettingsSamlAutocreateUsersDomains) DeepCopy() *OrganizationSpecSettingsSamlAutocreateUsersDomains {
	if in == nil {
		return nil
	}
	out := new(OrganizationSpecSettingsSamlAutocreateUsersDomains)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSpecSettingsSamlIdpInitiatedLogin) DeepCopyInto(out *OrganizationSpecSettingsSamlIdpInitiatedLogin) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSpecSettingsSamlIdpInitiatedLogin.
func (in *OrganizationSpecSettingsSamlIdpInitiatedLogin) DeepCopy() *OrganizationSpecSettingsSamlIdpInitiatedLogin {
	if in == nil {
		return nil
	}
	out := new(OrganizationSpecSettingsSamlIdpInitiatedLogin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSpecSettingsSamlStrictMode) DeepCopyInto(out *OrganizationSpecSettingsSamlStrictMode) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSpecSettingsSamlStrictMode.
func (in *OrganizationSpecSettingsSamlStrictMode) DeepCopy() *OrganizationSpecSettingsSamlStrictMode {
	if in == nil {
		return nil
	}
	out := new(OrganizationSpecSettingsSamlStrictMode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSpecUser) DeepCopyInto(out *OrganizationSpecUser) {
	*out = *in
	if in.AccessRole != nil {
		in, out := &in.AccessRole, &out.AccessRole
		*out = new(string)
		**out = **in
	}
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSpecUser.
func (in *OrganizationSpecUser) DeepCopy() *OrganizationSpecUser {
	if in == nil {
		return nil
	}
	out := new(OrganizationSpecUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationStatus) DeepCopyInto(out *OrganizationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationStatus.
func (in *OrganizationStatus) DeepCopy() *OrganizationStatus {
	if in == nil {
		return nil
	}
	out := new(OrganizationStatus)
	in.DeepCopyInto(out)
	return out
}