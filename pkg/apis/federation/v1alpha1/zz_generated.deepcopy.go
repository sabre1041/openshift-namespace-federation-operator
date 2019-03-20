// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRegistrationStatus) DeepCopyInto(out *ClusterRegistrationStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRegistrationStatus.
func (in *ClusterRegistrationStatus) DeepCopy() *ClusterRegistrationStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterRegistrationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultipleNamespaceFederation) DeepCopyInto(out *MultipleNamespaceFederation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultipleNamespaceFederation.
func (in *MultipleNamespaceFederation) DeepCopy() *MultipleNamespaceFederation {
	if in == nil {
		return nil
	}
	out := new(MultipleNamespaceFederation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultipleNamespaceFederation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultipleNamespaceFederationList) DeepCopyInto(out *MultipleNamespaceFederationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MultipleNamespaceFederation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultipleNamespaceFederationList.
func (in *MultipleNamespaceFederationList) DeepCopy() *MultipleNamespaceFederationList {
	if in == nil {
		return nil
	}
	out := new(MultipleNamespaceFederationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultipleNamespaceFederationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultipleNamespaceFederationSpec) DeepCopyInto(out *MultipleNamespaceFederationSpec) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultipleNamespaceFederationSpec.
func (in *MultipleNamespaceFederationSpec) DeepCopy() *MultipleNamespaceFederationSpec {
	if in == nil {
		return nil
	}
	out := new(MultipleNamespaceFederationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultipleNamespaceFederationStatus) DeepCopyInto(out *MultipleNamespaceFederationStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultipleNamespaceFederationStatus.
func (in *MultipleNamespaceFederationStatus) DeepCopy() *MultipleNamespaceFederationStatus {
	if in == nil {
		return nil
	}
	out := new(MultipleNamespaceFederationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceFederation) DeepCopyInto(out *NamespaceFederation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceFederation.
func (in *NamespaceFederation) DeepCopy() *NamespaceFederation {
	if in == nil {
		return nil
	}
	out := new(NamespaceFederation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceFederation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceFederationList) DeepCopyInto(out *NamespaceFederationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NamespaceFederation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceFederationList.
func (in *NamespaceFederationList) DeepCopy() *NamespaceFederationList {
	if in == nil {
		return nil
	}
	out := new(NamespaceFederationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceFederationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceFederationSpec) DeepCopyInto(out *NamespaceFederationSpec) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FederatedTypes != nil {
		in, out := &in.FederatedTypes, &out.FederatedTypes
		*out = make([]v1.TypeMeta, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceFederationSpec.
func (in *NamespaceFederationSpec) DeepCopy() *NamespaceFederationSpec {
	if in == nil {
		return nil
	}
	out := new(NamespaceFederationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceFederationStatus) DeepCopyInto(out *NamespaceFederationStatus) {
	*out = *in
	if in.ClusterRegistrationStatuses != nil {
		in, out := &in.ClusterRegistrationStatuses, &out.ClusterRegistrationStatuses
		*out = make([]ClusterRegistrationStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceFederationStatus.
func (in *NamespaceFederationStatus) DeepCopy() *NamespaceFederationStatus {
	if in == nil {
		return nil
	}
	out := new(NamespaceFederationStatus)
	in.DeepCopyInto(out)
	return out
}
