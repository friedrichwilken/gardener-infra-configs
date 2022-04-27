//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package azure_infra_config

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfrastructureConfig) DeepCopyInto(out *InfrastructureConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.ResourceGroup != nil {
		in, out := &in.ResourceGroup, &out.ResourceGroup
		*out = new(ResourceGroup)
		**out = **in
	}
	in.Networks.DeepCopyInto(&out.Networks)
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(IdentityConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfrastructureConfig.
func (in *InfrastructureConfig) DeepCopy() *InfrastructureConfig {
	if in == nil {
		return nil
	}
	out := new(InfrastructureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfrastructureConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfrastructureStatus) DeepCopyInto(out *InfrastructureStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Networks.DeepCopyInto(&out.Networks)
	out.ResourceGroup = in.ResourceGroup
	if in.AvailabilitySets != nil {
		in, out := &in.AvailabilitySets, &out.AvailabilitySets
		*out = make([]AvailabilitySet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RouteTables != nil {
		in, out := &in.RouteTables, &out.RouteTables
		*out = make([]RouteTable, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]SecurityGroup, len(*in))
		copy(*out, *in)
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(IdentityStatus)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfrastructureStatus.
func (in *InfrastructureStatus) DeepCopy() *InfrastructureStatus {
	if in == nil {
		return nil
	}
	out := new(InfrastructureStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfrastructureStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
