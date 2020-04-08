// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha3

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSCluster) DeepCopyInto(out *AKSCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSCluster.
func (in *AKSCluster) DeepCopy() *AKSCluster {
	if in == nil {
		return nil
	}
	out := new(AKSCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AKSCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSClusterClass) DeepCopyInto(out *AKSClusterClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.SpecTemplate.DeepCopyInto(&out.SpecTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSClusterClass.
func (in *AKSClusterClass) DeepCopy() *AKSClusterClass {
	if in == nil {
		return nil
	}
	out := new(AKSClusterClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AKSClusterClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSClusterClassList) DeepCopyInto(out *AKSClusterClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AKSClusterClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSClusterClassList.
func (in *AKSClusterClassList) DeepCopy() *AKSClusterClassList {
	if in == nil {
		return nil
	}
	out := new(AKSClusterClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AKSClusterClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSClusterClassSpecTemplate) DeepCopyInto(out *AKSClusterClassSpecTemplate) {
	*out = *in
	in.ClassSpecTemplate.DeepCopyInto(&out.ClassSpecTemplate)
	in.AKSClusterParameters.DeepCopyInto(&out.AKSClusterParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSClusterClassSpecTemplate.
func (in *AKSClusterClassSpecTemplate) DeepCopy() *AKSClusterClassSpecTemplate {
	if in == nil {
		return nil
	}
	out := new(AKSClusterClassSpecTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSClusterList) DeepCopyInto(out *AKSClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AKSCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSClusterList.
func (in *AKSClusterList) DeepCopy() *AKSClusterList {
	if in == nil {
		return nil
	}
	out := new(AKSClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AKSClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSClusterParameters) DeepCopyInto(out *AKSClusterParameters) {
	*out = *in
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(ResourceGroupNameReferencerForAKSCluster)
		**out = **in
	}
	if in.VnetSubnetIDRef != nil {
		in, out := &in.VnetSubnetIDRef, &out.VnetSubnetIDRef
		*out = new(SubnetIDReferencerForAKSCluster)
		**out = **in
	}
	if in.NodeCount != nil {
		in, out := &in.NodeCount, &out.NodeCount
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSClusterParameters.
func (in *AKSClusterParameters) DeepCopy() *AKSClusterParameters {
	if in == nil {
		return nil
	}
	out := new(AKSClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSClusterSpec) DeepCopyInto(out *AKSClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.AKSClusterParameters.DeepCopyInto(&out.AKSClusterParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSClusterSpec.
func (in *AKSClusterSpec) DeepCopy() *AKSClusterSpec {
	if in == nil {
		return nil
	}
	out := new(AKSClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSClusterStatus) DeepCopyInto(out *AKSClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSClusterStatus.
func (in *AKSClusterStatus) DeepCopy() *AKSClusterStatus {
	if in == nil {
		return nil
	}
	out := new(AKSClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupNameReferencerForAKSCluster) DeepCopyInto(out *ResourceGroupNameReferencerForAKSCluster) {
	*out = *in
	out.ResourceGroupNameReferencer = in.ResourceGroupNameReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupNameReferencerForAKSCluster.
func (in *ResourceGroupNameReferencerForAKSCluster) DeepCopy() *ResourceGroupNameReferencerForAKSCluster {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupNameReferencerForAKSCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetIDReferencerForAKSCluster) DeepCopyInto(out *SubnetIDReferencerForAKSCluster) {
	*out = *in
	out.SubnetIDReferencer = in.SubnetIDReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetIDReferencerForAKSCluster.
func (in *SubnetIDReferencerForAKSCluster) DeepCopy() *SubnetIDReferencerForAKSCluster {
	if in == nil {
		return nil
	}
	out := new(SubnetIDReferencerForAKSCluster)
	in.DeepCopyInto(out)
	return out
}
