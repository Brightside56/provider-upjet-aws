//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExcludesInitParameters) DeepCopyInto(out *ExcludesInitParameters) {
	*out = *in
	if in.FilterType != nil {
		in, out := &in.FilterType, &out.FilterType
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExcludesInitParameters.
func (in *ExcludesInitParameters) DeepCopy() *ExcludesInitParameters {
	if in == nil {
		return nil
	}
	out := new(ExcludesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExcludesObservation) DeepCopyInto(out *ExcludesObservation) {
	*out = *in
	if in.FilterType != nil {
		in, out := &in.FilterType, &out.FilterType
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExcludesObservation.
func (in *ExcludesObservation) DeepCopy() *ExcludesObservation {
	if in == nil {
		return nil
	}
	out := new(ExcludesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExcludesParameters) DeepCopyInto(out *ExcludesParameters) {
	*out = *in
	if in.FilterType != nil {
		in, out := &in.FilterType, &out.FilterType
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExcludesParameters.
func (in *ExcludesParameters) DeepCopy() *ExcludesParameters {
	if in == nil {
		return nil
	}
	out := new(ExcludesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IncludesInitParameters) DeepCopyInto(out *IncludesInitParameters) {
	*out = *in
	if in.FilterType != nil {
		in, out := &in.FilterType, &out.FilterType
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IncludesInitParameters.
func (in *IncludesInitParameters) DeepCopy() *IncludesInitParameters {
	if in == nil {
		return nil
	}
	out := new(IncludesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IncludesObservation) DeepCopyInto(out *IncludesObservation) {
	*out = *in
	if in.FilterType != nil {
		in, out := &in.FilterType, &out.FilterType
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IncludesObservation.
func (in *IncludesObservation) DeepCopy() *IncludesObservation {
	if in == nil {
		return nil
	}
	out := new(IncludesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IncludesParameters) DeepCopyInto(out *IncludesParameters) {
	*out = *in
	if in.FilterType != nil {
		in, out := &in.FilterType, &out.FilterType
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IncludesParameters.
func (in *IncludesParameters) DeepCopy() *IncludesParameters {
	if in == nil {
		return nil
	}
	out := new(IncludesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationS3) DeepCopyInto(out *LocationS3) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationS3.
func (in *LocationS3) DeepCopy() *LocationS3 {
	if in == nil {
		return nil
	}
	out := new(LocationS3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocationS3) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationS3InitParameters) DeepCopyInto(out *LocationS3InitParameters) {
	*out = *in
	if in.AgentArns != nil {
		in, out := &in.AgentArns, &out.AgentArns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.S3Config != nil {
		in, out := &in.S3Config, &out.S3Config
		*out = make([]S3ConfigInitParameters, len(*in))
		copy(*out, *in)
	}
	if in.S3StorageClass != nil {
		in, out := &in.S3StorageClass, &out.S3StorageClass
		*out = new(string)
		**out = **in
	}
	if in.Subdirectory != nil {
		in, out := &in.Subdirectory, &out.Subdirectory
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationS3InitParameters.
func (in *LocationS3InitParameters) DeepCopy() *LocationS3InitParameters {
	if in == nil {
		return nil
	}
	out := new(LocationS3InitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationS3List) DeepCopyInto(out *LocationS3List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocationS3, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationS3List.
func (in *LocationS3List) DeepCopy() *LocationS3List {
	if in == nil {
		return nil
	}
	out := new(LocationS3List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocationS3List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationS3Observation) DeepCopyInto(out *LocationS3Observation) {
	*out = *in
	if in.AgentArns != nil {
		in, out := &in.AgentArns, &out.AgentArns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.S3BucketArn != nil {
		in, out := &in.S3BucketArn, &out.S3BucketArn
		*out = new(string)
		**out = **in
	}
	if in.S3Config != nil {
		in, out := &in.S3Config, &out.S3Config
		*out = make([]S3ConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.S3StorageClass != nil {
		in, out := &in.S3StorageClass, &out.S3StorageClass
		*out = new(string)
		**out = **in
	}
	if in.Subdirectory != nil {
		in, out := &in.Subdirectory, &out.Subdirectory
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationS3Observation.
func (in *LocationS3Observation) DeepCopy() *LocationS3Observation {
	if in == nil {
		return nil
	}
	out := new(LocationS3Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationS3Parameters) DeepCopyInto(out *LocationS3Parameters) {
	*out = *in
	if in.AgentArns != nil {
		in, out := &in.AgentArns, &out.AgentArns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.S3BucketArn != nil {
		in, out := &in.S3BucketArn, &out.S3BucketArn
		*out = new(string)
		**out = **in
	}
	if in.S3BucketArnRef != nil {
		in, out := &in.S3BucketArnRef, &out.S3BucketArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.S3BucketArnSelector != nil {
		in, out := &in.S3BucketArnSelector, &out.S3BucketArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.S3Config != nil {
		in, out := &in.S3Config, &out.S3Config
		*out = make([]S3ConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.S3StorageClass != nil {
		in, out := &in.S3StorageClass, &out.S3StorageClass
		*out = new(string)
		**out = **in
	}
	if in.Subdirectory != nil {
		in, out := &in.Subdirectory, &out.Subdirectory
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationS3Parameters.
func (in *LocationS3Parameters) DeepCopy() *LocationS3Parameters {
	if in == nil {
		return nil
	}
	out := new(LocationS3Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationS3Spec) DeepCopyInto(out *LocationS3Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationS3Spec.
func (in *LocationS3Spec) DeepCopy() *LocationS3Spec {
	if in == nil {
		return nil
	}
	out := new(LocationS3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationS3Status) DeepCopyInto(out *LocationS3Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationS3Status.
func (in *LocationS3Status) DeepCopy() *LocationS3Status {
	if in == nil {
		return nil
	}
	out := new(LocationS3Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionsInitParameters) DeepCopyInto(out *OptionsInitParameters) {
	*out = *in
	if in.Atime != nil {
		in, out := &in.Atime, &out.Atime
		*out = new(string)
		**out = **in
	}
	if in.BytesPerSecond != nil {
		in, out := &in.BytesPerSecond, &out.BytesPerSecond
		*out = new(float64)
		**out = **in
	}
	if in.GID != nil {
		in, out := &in.GID, &out.GID
		*out = new(string)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.Mtime != nil {
		in, out := &in.Mtime, &out.Mtime
		*out = new(string)
		**out = **in
	}
	if in.OverwriteMode != nil {
		in, out := &in.OverwriteMode, &out.OverwriteMode
		*out = new(string)
		**out = **in
	}
	if in.PosixPermissions != nil {
		in, out := &in.PosixPermissions, &out.PosixPermissions
		*out = new(string)
		**out = **in
	}
	if in.PreserveDeletedFiles != nil {
		in, out := &in.PreserveDeletedFiles, &out.PreserveDeletedFiles
		*out = new(string)
		**out = **in
	}
	if in.PreserveDevices != nil {
		in, out := &in.PreserveDevices, &out.PreserveDevices
		*out = new(string)
		**out = **in
	}
	if in.SecurityDescriptorCopyFlags != nil {
		in, out := &in.SecurityDescriptorCopyFlags, &out.SecurityDescriptorCopyFlags
		*out = new(string)
		**out = **in
	}
	if in.TaskQueueing != nil {
		in, out := &in.TaskQueueing, &out.TaskQueueing
		*out = new(string)
		**out = **in
	}
	if in.TransferMode != nil {
		in, out := &in.TransferMode, &out.TransferMode
		*out = new(string)
		**out = **in
	}
	if in.UID != nil {
		in, out := &in.UID, &out.UID
		*out = new(string)
		**out = **in
	}
	if in.VerifyMode != nil {
		in, out := &in.VerifyMode, &out.VerifyMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionsInitParameters.
func (in *OptionsInitParameters) DeepCopy() *OptionsInitParameters {
	if in == nil {
		return nil
	}
	out := new(OptionsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionsObservation) DeepCopyInto(out *OptionsObservation) {
	*out = *in
	if in.Atime != nil {
		in, out := &in.Atime, &out.Atime
		*out = new(string)
		**out = **in
	}
	if in.BytesPerSecond != nil {
		in, out := &in.BytesPerSecond, &out.BytesPerSecond
		*out = new(float64)
		**out = **in
	}
	if in.GID != nil {
		in, out := &in.GID, &out.GID
		*out = new(string)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.Mtime != nil {
		in, out := &in.Mtime, &out.Mtime
		*out = new(string)
		**out = **in
	}
	if in.OverwriteMode != nil {
		in, out := &in.OverwriteMode, &out.OverwriteMode
		*out = new(string)
		**out = **in
	}
	if in.PosixPermissions != nil {
		in, out := &in.PosixPermissions, &out.PosixPermissions
		*out = new(string)
		**out = **in
	}
	if in.PreserveDeletedFiles != nil {
		in, out := &in.PreserveDeletedFiles, &out.PreserveDeletedFiles
		*out = new(string)
		**out = **in
	}
	if in.PreserveDevices != nil {
		in, out := &in.PreserveDevices, &out.PreserveDevices
		*out = new(string)
		**out = **in
	}
	if in.SecurityDescriptorCopyFlags != nil {
		in, out := &in.SecurityDescriptorCopyFlags, &out.SecurityDescriptorCopyFlags
		*out = new(string)
		**out = **in
	}
	if in.TaskQueueing != nil {
		in, out := &in.TaskQueueing, &out.TaskQueueing
		*out = new(string)
		**out = **in
	}
	if in.TransferMode != nil {
		in, out := &in.TransferMode, &out.TransferMode
		*out = new(string)
		**out = **in
	}
	if in.UID != nil {
		in, out := &in.UID, &out.UID
		*out = new(string)
		**out = **in
	}
	if in.VerifyMode != nil {
		in, out := &in.VerifyMode, &out.VerifyMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionsObservation.
func (in *OptionsObservation) DeepCopy() *OptionsObservation {
	if in == nil {
		return nil
	}
	out := new(OptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionsParameters) DeepCopyInto(out *OptionsParameters) {
	*out = *in
	if in.Atime != nil {
		in, out := &in.Atime, &out.Atime
		*out = new(string)
		**out = **in
	}
	if in.BytesPerSecond != nil {
		in, out := &in.BytesPerSecond, &out.BytesPerSecond
		*out = new(float64)
		**out = **in
	}
	if in.GID != nil {
		in, out := &in.GID, &out.GID
		*out = new(string)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.Mtime != nil {
		in, out := &in.Mtime, &out.Mtime
		*out = new(string)
		**out = **in
	}
	if in.OverwriteMode != nil {
		in, out := &in.OverwriteMode, &out.OverwriteMode
		*out = new(string)
		**out = **in
	}
	if in.PosixPermissions != nil {
		in, out := &in.PosixPermissions, &out.PosixPermissions
		*out = new(string)
		**out = **in
	}
	if in.PreserveDeletedFiles != nil {
		in, out := &in.PreserveDeletedFiles, &out.PreserveDeletedFiles
		*out = new(string)
		**out = **in
	}
	if in.PreserveDevices != nil {
		in, out := &in.PreserveDevices, &out.PreserveDevices
		*out = new(string)
		**out = **in
	}
	if in.SecurityDescriptorCopyFlags != nil {
		in, out := &in.SecurityDescriptorCopyFlags, &out.SecurityDescriptorCopyFlags
		*out = new(string)
		**out = **in
	}
	if in.TaskQueueing != nil {
		in, out := &in.TaskQueueing, &out.TaskQueueing
		*out = new(string)
		**out = **in
	}
	if in.TransferMode != nil {
		in, out := &in.TransferMode, &out.TransferMode
		*out = new(string)
		**out = **in
	}
	if in.UID != nil {
		in, out := &in.UID, &out.UID
		*out = new(string)
		**out = **in
	}
	if in.VerifyMode != nil {
		in, out := &in.VerifyMode, &out.VerifyMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionsParameters.
func (in *OptionsParameters) DeepCopy() *OptionsParameters {
	if in == nil {
		return nil
	}
	out := new(OptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ConfigInitParameters) DeepCopyInto(out *S3ConfigInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ConfigInitParameters.
func (in *S3ConfigInitParameters) DeepCopy() *S3ConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(S3ConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ConfigObservation) DeepCopyInto(out *S3ConfigObservation) {
	*out = *in
	if in.BucketAccessRoleArn != nil {
		in, out := &in.BucketAccessRoleArn, &out.BucketAccessRoleArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ConfigObservation.
func (in *S3ConfigObservation) DeepCopy() *S3ConfigObservation {
	if in == nil {
		return nil
	}
	out := new(S3ConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ConfigParameters) DeepCopyInto(out *S3ConfigParameters) {
	*out = *in
	if in.BucketAccessRoleArn != nil {
		in, out := &in.BucketAccessRoleArn, &out.BucketAccessRoleArn
		*out = new(string)
		**out = **in
	}
	if in.BucketAccessRoleArnRef != nil {
		in, out := &in.BucketAccessRoleArnRef, &out.BucketAccessRoleArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.BucketAccessRoleArnSelector != nil {
		in, out := &in.BucketAccessRoleArnSelector, &out.BucketAccessRoleArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ConfigParameters.
func (in *S3ConfigParameters) DeepCopy() *S3ConfigParameters {
	if in == nil {
		return nil
	}
	out := new(S3ConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleInitParameters) DeepCopyInto(out *ScheduleInitParameters) {
	*out = *in
	if in.ScheduleExpression != nil {
		in, out := &in.ScheduleExpression, &out.ScheduleExpression
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleInitParameters.
func (in *ScheduleInitParameters) DeepCopy() *ScheduleInitParameters {
	if in == nil {
		return nil
	}
	out := new(ScheduleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleObservation) DeepCopyInto(out *ScheduleObservation) {
	*out = *in
	if in.ScheduleExpression != nil {
		in, out := &in.ScheduleExpression, &out.ScheduleExpression
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleObservation.
func (in *ScheduleObservation) DeepCopy() *ScheduleObservation {
	if in == nil {
		return nil
	}
	out := new(ScheduleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleParameters) DeepCopyInto(out *ScheduleParameters) {
	*out = *in
	if in.ScheduleExpression != nil {
		in, out := &in.ScheduleExpression, &out.ScheduleExpression
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleParameters.
func (in *ScheduleParameters) DeepCopy() *ScheduleParameters {
	if in == nil {
		return nil
	}
	out := new(ScheduleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Task) DeepCopyInto(out *Task) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Task.
func (in *Task) DeepCopy() *Task {
	if in == nil {
		return nil
	}
	out := new(Task)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Task) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskInitParameters) DeepCopyInto(out *TaskInitParameters) {
	*out = *in
	if in.Excludes != nil {
		in, out := &in.Excludes, &out.Excludes
		*out = make([]ExcludesInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Includes != nil {
		in, out := &in.Includes, &out.Includes
		*out = make([]IncludesInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]OptionsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = make([]ScheduleInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskInitParameters.
func (in *TaskInitParameters) DeepCopy() *TaskInitParameters {
	if in == nil {
		return nil
	}
	out := new(TaskInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskList) DeepCopyInto(out *TaskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Task, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskList.
func (in *TaskList) DeepCopy() *TaskList {
	if in == nil {
		return nil
	}
	out := new(TaskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TaskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskObservation) DeepCopyInto(out *TaskObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.CloudwatchLogGroupArn != nil {
		in, out := &in.CloudwatchLogGroupArn, &out.CloudwatchLogGroupArn
		*out = new(string)
		**out = **in
	}
	if in.DestinationLocationArn != nil {
		in, out := &in.DestinationLocationArn, &out.DestinationLocationArn
		*out = new(string)
		**out = **in
	}
	if in.Excludes != nil {
		in, out := &in.Excludes, &out.Excludes
		*out = make([]ExcludesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Includes != nil {
		in, out := &in.Includes, &out.Includes
		*out = make([]IncludesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]OptionsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = make([]ScheduleObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SourceLocationArn != nil {
		in, out := &in.SourceLocationArn, &out.SourceLocationArn
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskObservation.
func (in *TaskObservation) DeepCopy() *TaskObservation {
	if in == nil {
		return nil
	}
	out := new(TaskObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskParameters) DeepCopyInto(out *TaskParameters) {
	*out = *in
	if in.CloudwatchLogGroupArn != nil {
		in, out := &in.CloudwatchLogGroupArn, &out.CloudwatchLogGroupArn
		*out = new(string)
		**out = **in
	}
	if in.CloudwatchLogGroupArnRef != nil {
		in, out := &in.CloudwatchLogGroupArnRef, &out.CloudwatchLogGroupArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.CloudwatchLogGroupArnSelector != nil {
		in, out := &in.CloudwatchLogGroupArnSelector, &out.CloudwatchLogGroupArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.DestinationLocationArn != nil {
		in, out := &in.DestinationLocationArn, &out.DestinationLocationArn
		*out = new(string)
		**out = **in
	}
	if in.DestinationLocationArnRef != nil {
		in, out := &in.DestinationLocationArnRef, &out.DestinationLocationArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DestinationLocationArnSelector != nil {
		in, out := &in.DestinationLocationArnSelector, &out.DestinationLocationArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Excludes != nil {
		in, out := &in.Excludes, &out.Excludes
		*out = make([]ExcludesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Includes != nil {
		in, out := &in.Includes, &out.Includes
		*out = make([]IncludesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]OptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = make([]ScheduleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SourceLocationArn != nil {
		in, out := &in.SourceLocationArn, &out.SourceLocationArn
		*out = new(string)
		**out = **in
	}
	if in.SourceLocationArnRef != nil {
		in, out := &in.SourceLocationArnRef, &out.SourceLocationArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SourceLocationArnSelector != nil {
		in, out := &in.SourceLocationArnSelector, &out.SourceLocationArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskParameters.
func (in *TaskParameters) DeepCopy() *TaskParameters {
	if in == nil {
		return nil
	}
	out := new(TaskParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskSpec) DeepCopyInto(out *TaskSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskSpec.
func (in *TaskSpec) DeepCopy() *TaskSpec {
	if in == nil {
		return nil
	}
	out := new(TaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskStatus) DeepCopyInto(out *TaskStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskStatus.
func (in *TaskStatus) DeepCopy() *TaskStatus {
	if in == nil {
		return nil
	}
	out := new(TaskStatus)
	in.DeepCopyInto(out)
	return out
}
