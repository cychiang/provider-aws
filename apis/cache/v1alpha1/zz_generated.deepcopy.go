//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheCluster) DeepCopyInto(out *CacheCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheCluster.
func (in *CacheCluster) DeepCopy() *CacheCluster {
	if in == nil {
		return nil
	}
	out := new(CacheCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CacheCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheClusterList) DeepCopyInto(out *CacheClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CacheCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheClusterList.
func (in *CacheClusterList) DeepCopy() *CacheClusterList {
	if in == nil {
		return nil
	}
	out := new(CacheClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CacheClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheClusterObservation) DeepCopyInto(out *CacheClusterObservation) {
	*out = *in
	if in.CacheNodes != nil {
		in, out := &in.CacheNodes, &out.CacheNodes
		*out = make([]CacheNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.CacheParameterGroup.DeepCopyInto(&out.CacheParameterGroup)
	out.ConfigurationEndpoint = in.ConfigurationEndpoint
	in.NotificationConfiguration.DeepCopyInto(&out.NotificationConfiguration)
	in.PendingModifiedValues.DeepCopyInto(&out.PendingModifiedValues)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheClusterObservation.
func (in *CacheClusterObservation) DeepCopy() *CacheClusterObservation {
	if in == nil {
		return nil
	}
	out := new(CacheClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheClusterParameters) DeepCopyInto(out *CacheClusterParameters) {
	*out = *in
	if in.ApplyImmediately != nil {
		in, out := &in.ApplyImmediately, &out.ApplyImmediately
		*out = new(bool)
		**out = **in
	}
	if in.AZMode != nil {
		in, out := &in.AZMode, &out.AZMode
		*out = new(string)
		**out = **in
	}
	if in.AuthToken != nil {
		in, out := &in.AuthToken, &out.AuthToken
		*out = new(string)
		**out = **in
	}
	if in.AuthTokenUpdateStrategy != nil {
		in, out := &in.AuthTokenUpdateStrategy, &out.AuthTokenUpdateStrategy
		*out = new(string)
		**out = **in
	}
	if in.CacheNodeIDsToRemove != nil {
		in, out := &in.CacheNodeIDsToRemove, &out.CacheNodeIDsToRemove
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CacheParameterGroupName != nil {
		in, out := &in.CacheParameterGroupName, &out.CacheParameterGroupName
		*out = new(string)
		**out = **in
	}
	if in.CacheSecurityGroupNames != nil {
		in, out := &in.CacheSecurityGroupNames, &out.CacheSecurityGroupNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CacheSubnetGroupName != nil {
		in, out := &in.CacheSubnetGroupName, &out.CacheSubnetGroupName
		*out = new(string)
		**out = **in
	}
	if in.CacheSubnetGroupNameRef != nil {
		in, out := &in.CacheSubnetGroupNameRef, &out.CacheSubnetGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.CacheSubnetGroupNameSelector != nil {
		in, out := &in.CacheSubnetGroupNameSelector, &out.CacheSubnetGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Engine != nil {
		in, out := &in.Engine, &out.Engine
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.NotificationTopicARN != nil {
		in, out := &in.NotificationTopicARN, &out.NotificationTopicARN
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.PreferredAvailabilityZone != nil {
		in, out := &in.PreferredAvailabilityZone, &out.PreferredAvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.PreferredAvailabilityZones != nil {
		in, out := &in.PreferredAvailabilityZones, &out.PreferredAvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PreferredMaintenanceWindow != nil {
		in, out := &in.PreferredMaintenanceWindow, &out.PreferredMaintenanceWindow
		*out = new(string)
		**out = **in
	}
	if in.ReplicationGroupID != nil {
		in, out := &in.ReplicationGroupID, &out.ReplicationGroupID
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupIDRefs != nil {
		in, out := &in.SecurityGroupIDRefs, &out.SecurityGroupIDRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityGroupIDSelector != nil {
		in, out := &in.SecurityGroupIDSelector, &out.SecurityGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SnapshotARNs != nil {
		in, out := &in.SnapshotARNs, &out.SnapshotARNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SnapshotName != nil {
		in, out := &in.SnapshotName, &out.SnapshotName
		*out = new(string)
		**out = **in
	}
	if in.SnapshotRetentionLimit != nil {
		in, out := &in.SnapshotRetentionLimit, &out.SnapshotRetentionLimit
		*out = new(int32)
		**out = **in
	}
	if in.SnapshotWindow != nil {
		in, out := &in.SnapshotWindow, &out.SnapshotWindow
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheClusterParameters.
func (in *CacheClusterParameters) DeepCopy() *CacheClusterParameters {
	if in == nil {
		return nil
	}
	out := new(CacheClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheClusterSpec) DeepCopyInto(out *CacheClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheClusterSpec.
func (in *CacheClusterSpec) DeepCopy() *CacheClusterSpec {
	if in == nil {
		return nil
	}
	out := new(CacheClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheClusterStatus) DeepCopyInto(out *CacheClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheClusterStatus.
func (in *CacheClusterStatus) DeepCopy() *CacheClusterStatus {
	if in == nil {
		return nil
	}
	out := new(CacheClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheNode) DeepCopyInto(out *CacheNode) {
	*out = *in
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(Endpoint)
		**out = **in
	}
	if in.SourceCacheNodeID != nil {
		in, out := &in.SourceCacheNodeID, &out.SourceCacheNodeID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheNode.
func (in *CacheNode) DeepCopy() *CacheNode {
	if in == nil {
		return nil
	}
	out := new(CacheNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheParameterGroupStatus) DeepCopyInto(out *CacheParameterGroupStatus) {
	*out = *in
	if in.CacheNodeIDsToReboot != nil {
		in, out := &in.CacheNodeIDsToReboot, &out.CacheNodeIDsToReboot
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheParameterGroupStatus.
func (in *CacheParameterGroupStatus) DeepCopy() *CacheParameterGroupStatus {
	if in == nil {
		return nil
	}
	out := new(CacheParameterGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheSubnetGroup) DeepCopyInto(out *CacheSubnetGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheSubnetGroup.
func (in *CacheSubnetGroup) DeepCopy() *CacheSubnetGroup {
	if in == nil {
		return nil
	}
	out := new(CacheSubnetGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CacheSubnetGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheSubnetGroupExternalStatus) DeepCopyInto(out *CacheSubnetGroupExternalStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheSubnetGroupExternalStatus.
func (in *CacheSubnetGroupExternalStatus) DeepCopy() *CacheSubnetGroupExternalStatus {
	if in == nil {
		return nil
	}
	out := new(CacheSubnetGroupExternalStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheSubnetGroupList) DeepCopyInto(out *CacheSubnetGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CacheSubnetGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheSubnetGroupList.
func (in *CacheSubnetGroupList) DeepCopy() *CacheSubnetGroupList {
	if in == nil {
		return nil
	}
	out := new(CacheSubnetGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CacheSubnetGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheSubnetGroupParameters) DeepCopyInto(out *CacheSubnetGroupParameters) {
	*out = *in
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetIDRefs != nil {
		in, out := &in.SubnetIDRefs, &out.SubnetIDRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheSubnetGroupParameters.
func (in *CacheSubnetGroupParameters) DeepCopy() *CacheSubnetGroupParameters {
	if in == nil {
		return nil
	}
	out := new(CacheSubnetGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheSubnetGroupSpec) DeepCopyInto(out *CacheSubnetGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheSubnetGroupSpec.
func (in *CacheSubnetGroupSpec) DeepCopy() *CacheSubnetGroupSpec {
	if in == nil {
		return nil
	}
	out := new(CacheSubnetGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheSubnetGroupStatus) DeepCopyInto(out *CacheSubnetGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheSubnetGroupStatus.
func (in *CacheSubnetGroupStatus) DeepCopy() *CacheSubnetGroupStatus {
	if in == nil {
		return nil
	}
	out := new(CacheSubnetGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationConfiguration) DeepCopyInto(out *NotificationConfiguration) {
	*out = *in
	if in.TopicStatus != nil {
		in, out := &in.TopicStatus, &out.TopicStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationConfiguration.
func (in *NotificationConfiguration) DeepCopy() *NotificationConfiguration {
	if in == nil {
		return nil
	}
	out := new(NotificationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PendingModifiedValues) DeepCopyInto(out *PendingModifiedValues) {
	*out = *in
	if in.CacheNodeIDsToRemove != nil {
		in, out := &in.CacheNodeIDsToRemove, &out.CacheNodeIDsToRemove
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.NumCacheNodes != nil {
		in, out := &in.NumCacheNodes, &out.NumCacheNodes
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PendingModifiedValues.
func (in *PendingModifiedValues) DeepCopy() *PendingModifiedValues {
	if in == nil {
		return nil
	}
	out := new(PendingModifiedValues)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}
