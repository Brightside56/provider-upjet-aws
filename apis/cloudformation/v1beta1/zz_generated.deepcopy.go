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
func (in *AutoDeploymentInitParameters) DeepCopyInto(out *AutoDeploymentInitParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.RetainStacksOnAccountRemoval != nil {
		in, out := &in.RetainStacksOnAccountRemoval, &out.RetainStacksOnAccountRemoval
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoDeploymentInitParameters.
func (in *AutoDeploymentInitParameters) DeepCopy() *AutoDeploymentInitParameters {
	if in == nil {
		return nil
	}
	out := new(AutoDeploymentInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoDeploymentObservation) DeepCopyInto(out *AutoDeploymentObservation) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.RetainStacksOnAccountRemoval != nil {
		in, out := &in.RetainStacksOnAccountRemoval, &out.RetainStacksOnAccountRemoval
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoDeploymentObservation.
func (in *AutoDeploymentObservation) DeepCopy() *AutoDeploymentObservation {
	if in == nil {
		return nil
	}
	out := new(AutoDeploymentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoDeploymentParameters) DeepCopyInto(out *AutoDeploymentParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.RetainStacksOnAccountRemoval != nil {
		in, out := &in.RetainStacksOnAccountRemoval, &out.RetainStacksOnAccountRemoval
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoDeploymentParameters.
func (in *AutoDeploymentParameters) DeepCopy() *AutoDeploymentParameters {
	if in == nil {
		return nil
	}
	out := new(AutoDeploymentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationPreferencesInitParameters) DeepCopyInto(out *OperationPreferencesInitParameters) {
	*out = *in
	if in.FailureToleranceCount != nil {
		in, out := &in.FailureToleranceCount, &out.FailureToleranceCount
		*out = new(float64)
		**out = **in
	}
	if in.FailureTolerancePercentage != nil {
		in, out := &in.FailureTolerancePercentage, &out.FailureTolerancePercentage
		*out = new(float64)
		**out = **in
	}
	if in.MaxConcurrentCount != nil {
		in, out := &in.MaxConcurrentCount, &out.MaxConcurrentCount
		*out = new(float64)
		**out = **in
	}
	if in.MaxConcurrentPercentage != nil {
		in, out := &in.MaxConcurrentPercentage, &out.MaxConcurrentPercentage
		*out = new(float64)
		**out = **in
	}
	if in.RegionConcurrencyType != nil {
		in, out := &in.RegionConcurrencyType, &out.RegionConcurrencyType
		*out = new(string)
		**out = **in
	}
	if in.RegionOrder != nil {
		in, out := &in.RegionOrder, &out.RegionOrder
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationPreferencesInitParameters.
func (in *OperationPreferencesInitParameters) DeepCopy() *OperationPreferencesInitParameters {
	if in == nil {
		return nil
	}
	out := new(OperationPreferencesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationPreferencesObservation) DeepCopyInto(out *OperationPreferencesObservation) {
	*out = *in
	if in.FailureToleranceCount != nil {
		in, out := &in.FailureToleranceCount, &out.FailureToleranceCount
		*out = new(float64)
		**out = **in
	}
	if in.FailureTolerancePercentage != nil {
		in, out := &in.FailureTolerancePercentage, &out.FailureTolerancePercentage
		*out = new(float64)
		**out = **in
	}
	if in.MaxConcurrentCount != nil {
		in, out := &in.MaxConcurrentCount, &out.MaxConcurrentCount
		*out = new(float64)
		**out = **in
	}
	if in.MaxConcurrentPercentage != nil {
		in, out := &in.MaxConcurrentPercentage, &out.MaxConcurrentPercentage
		*out = new(float64)
		**out = **in
	}
	if in.RegionConcurrencyType != nil {
		in, out := &in.RegionConcurrencyType, &out.RegionConcurrencyType
		*out = new(string)
		**out = **in
	}
	if in.RegionOrder != nil {
		in, out := &in.RegionOrder, &out.RegionOrder
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationPreferencesObservation.
func (in *OperationPreferencesObservation) DeepCopy() *OperationPreferencesObservation {
	if in == nil {
		return nil
	}
	out := new(OperationPreferencesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationPreferencesParameters) DeepCopyInto(out *OperationPreferencesParameters) {
	*out = *in
	if in.FailureToleranceCount != nil {
		in, out := &in.FailureToleranceCount, &out.FailureToleranceCount
		*out = new(float64)
		**out = **in
	}
	if in.FailureTolerancePercentage != nil {
		in, out := &in.FailureTolerancePercentage, &out.FailureTolerancePercentage
		*out = new(float64)
		**out = **in
	}
	if in.MaxConcurrentCount != nil {
		in, out := &in.MaxConcurrentCount, &out.MaxConcurrentCount
		*out = new(float64)
		**out = **in
	}
	if in.MaxConcurrentPercentage != nil {
		in, out := &in.MaxConcurrentPercentage, &out.MaxConcurrentPercentage
		*out = new(float64)
		**out = **in
	}
	if in.RegionConcurrencyType != nil {
		in, out := &in.RegionConcurrencyType, &out.RegionConcurrencyType
		*out = new(string)
		**out = **in
	}
	if in.RegionOrder != nil {
		in, out := &in.RegionOrder, &out.RegionOrder
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationPreferencesParameters.
func (in *OperationPreferencesParameters) DeepCopy() *OperationPreferencesParameters {
	if in == nil {
		return nil
	}
	out := new(OperationPreferencesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stack) DeepCopyInto(out *Stack) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stack.
func (in *Stack) DeepCopy() *Stack {
	if in == nil {
		return nil
	}
	out := new(Stack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Stack) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackInitParameters) DeepCopyInto(out *StackInitParameters) {
	*out = *in
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DisableRollback != nil {
		in, out := &in.DisableRollback, &out.DisableRollback
		*out = new(bool)
		**out = **in
	}
	if in.NotificationArns != nil {
		in, out := &in.NotificationArns, &out.NotificationArns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OnFailure != nil {
		in, out := &in.OnFailure, &out.OnFailure
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
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
	if in.PolicyBody != nil {
		in, out := &in.PolicyBody, &out.PolicyBody
		*out = new(string)
		**out = **in
	}
	if in.PolicyURL != nil {
		in, out := &in.PolicyURL, &out.PolicyURL
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
	if in.TemplateBody != nil {
		in, out := &in.TemplateBody, &out.TemplateBody
		*out = new(string)
		**out = **in
	}
	if in.TemplateURL != nil {
		in, out := &in.TemplateURL, &out.TemplateURL
		*out = new(string)
		**out = **in
	}
	if in.TimeoutInMinutes != nil {
		in, out := &in.TimeoutInMinutes, &out.TimeoutInMinutes
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackInitParameters.
func (in *StackInitParameters) DeepCopy() *StackInitParameters {
	if in == nil {
		return nil
	}
	out := new(StackInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackList) DeepCopyInto(out *StackList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Stack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackList.
func (in *StackList) DeepCopy() *StackList {
	if in == nil {
		return nil
	}
	out := new(StackList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StackList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackObservation) DeepCopyInto(out *StackObservation) {
	*out = *in
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DisableRollback != nil {
		in, out := &in.DisableRollback, &out.DisableRollback
		*out = new(bool)
		**out = **in
	}
	if in.IAMRoleArn != nil {
		in, out := &in.IAMRoleArn, &out.IAMRoleArn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NotificationArns != nil {
		in, out := &in.NotificationArns, &out.NotificationArns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OnFailure != nil {
		in, out := &in.OnFailure, &out.OnFailure
		*out = new(string)
		**out = **in
	}
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
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
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
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
	if in.PolicyBody != nil {
		in, out := &in.PolicyBody, &out.PolicyBody
		*out = new(string)
		**out = **in
	}
	if in.PolicyURL != nil {
		in, out := &in.PolicyURL, &out.PolicyURL
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
	if in.TemplateBody != nil {
		in, out := &in.TemplateBody, &out.TemplateBody
		*out = new(string)
		**out = **in
	}
	if in.TemplateURL != nil {
		in, out := &in.TemplateURL, &out.TemplateURL
		*out = new(string)
		**out = **in
	}
	if in.TimeoutInMinutes != nil {
		in, out := &in.TimeoutInMinutes, &out.TimeoutInMinutes
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackObservation.
func (in *StackObservation) DeepCopy() *StackObservation {
	if in == nil {
		return nil
	}
	out := new(StackObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackParameters) DeepCopyInto(out *StackParameters) {
	*out = *in
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DisableRollback != nil {
		in, out := &in.DisableRollback, &out.DisableRollback
		*out = new(bool)
		**out = **in
	}
	if in.IAMRoleArn != nil {
		in, out := &in.IAMRoleArn, &out.IAMRoleArn
		*out = new(string)
		**out = **in
	}
	if in.IAMRoleArnRef != nil {
		in, out := &in.IAMRoleArnRef, &out.IAMRoleArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.IAMRoleArnSelector != nil {
		in, out := &in.IAMRoleArnSelector, &out.IAMRoleArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NotificationArns != nil {
		in, out := &in.NotificationArns, &out.NotificationArns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OnFailure != nil {
		in, out := &in.OnFailure, &out.OnFailure
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
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
	if in.PolicyBody != nil {
		in, out := &in.PolicyBody, &out.PolicyBody
		*out = new(string)
		**out = **in
	}
	if in.PolicyURL != nil {
		in, out := &in.PolicyURL, &out.PolicyURL
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
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
	if in.TemplateBody != nil {
		in, out := &in.TemplateBody, &out.TemplateBody
		*out = new(string)
		**out = **in
	}
	if in.TemplateURL != nil {
		in, out := &in.TemplateURL, &out.TemplateURL
		*out = new(string)
		**out = **in
	}
	if in.TimeoutInMinutes != nil {
		in, out := &in.TimeoutInMinutes, &out.TimeoutInMinutes
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackParameters.
func (in *StackParameters) DeepCopy() *StackParameters {
	if in == nil {
		return nil
	}
	out := new(StackParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackSet) DeepCopyInto(out *StackSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackSet.
func (in *StackSet) DeepCopy() *StackSet {
	if in == nil {
		return nil
	}
	out := new(StackSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StackSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackSetInitParameters) DeepCopyInto(out *StackSetInitParameters) {
	*out = *in
	if in.AutoDeployment != nil {
		in, out := &in.AutoDeployment, &out.AutoDeployment
		*out = make([]AutoDeploymentInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CallAs != nil {
		in, out := &in.CallAs, &out.CallAs
		*out = new(string)
		**out = **in
	}
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ExecutionRoleName != nil {
		in, out := &in.ExecutionRoleName, &out.ExecutionRoleName
		*out = new(string)
		**out = **in
	}
	if in.OperationPreferences != nil {
		in, out := &in.OperationPreferences, &out.OperationPreferences
		*out = make([]OperationPreferencesInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
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
	if in.PermissionModel != nil {
		in, out := &in.PermissionModel, &out.PermissionModel
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
	if in.TemplateBody != nil {
		in, out := &in.TemplateBody, &out.TemplateBody
		*out = new(string)
		**out = **in
	}
	if in.TemplateURL != nil {
		in, out := &in.TemplateURL, &out.TemplateURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackSetInitParameters.
func (in *StackSetInitParameters) DeepCopy() *StackSetInitParameters {
	if in == nil {
		return nil
	}
	out := new(StackSetInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackSetList) DeepCopyInto(out *StackSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StackSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackSetList.
func (in *StackSetList) DeepCopy() *StackSetList {
	if in == nil {
		return nil
	}
	out := new(StackSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StackSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackSetObservation) DeepCopyInto(out *StackSetObservation) {
	*out = *in
	if in.AdministrationRoleArn != nil {
		in, out := &in.AdministrationRoleArn, &out.AdministrationRoleArn
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AutoDeployment != nil {
		in, out := &in.AutoDeployment, &out.AutoDeployment
		*out = make([]AutoDeploymentObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CallAs != nil {
		in, out := &in.CallAs, &out.CallAs
		*out = new(string)
		**out = **in
	}
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ExecutionRoleName != nil {
		in, out := &in.ExecutionRoleName, &out.ExecutionRoleName
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.OperationPreferences != nil {
		in, out := &in.OperationPreferences, &out.OperationPreferences
		*out = make([]OperationPreferencesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
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
	if in.PermissionModel != nil {
		in, out := &in.PermissionModel, &out.PermissionModel
		*out = new(string)
		**out = **in
	}
	if in.StackSetID != nil {
		in, out := &in.StackSetID, &out.StackSetID
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
	if in.TemplateBody != nil {
		in, out := &in.TemplateBody, &out.TemplateBody
		*out = new(string)
		**out = **in
	}
	if in.TemplateURL != nil {
		in, out := &in.TemplateURL, &out.TemplateURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackSetObservation.
func (in *StackSetObservation) DeepCopy() *StackSetObservation {
	if in == nil {
		return nil
	}
	out := new(StackSetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackSetParameters) DeepCopyInto(out *StackSetParameters) {
	*out = *in
	if in.AdministrationRoleArn != nil {
		in, out := &in.AdministrationRoleArn, &out.AdministrationRoleArn
		*out = new(string)
		**out = **in
	}
	if in.AdministrationRoleArnRef != nil {
		in, out := &in.AdministrationRoleArnRef, &out.AdministrationRoleArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AdministrationRoleArnSelector != nil {
		in, out := &in.AdministrationRoleArnSelector, &out.AdministrationRoleArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoDeployment != nil {
		in, out := &in.AutoDeployment, &out.AutoDeployment
		*out = make([]AutoDeploymentParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CallAs != nil {
		in, out := &in.CallAs, &out.CallAs
		*out = new(string)
		**out = **in
	}
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ExecutionRoleName != nil {
		in, out := &in.ExecutionRoleName, &out.ExecutionRoleName
		*out = new(string)
		**out = **in
	}
	if in.OperationPreferences != nil {
		in, out := &in.OperationPreferences, &out.OperationPreferences
		*out = make([]OperationPreferencesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
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
	if in.PermissionModel != nil {
		in, out := &in.PermissionModel, &out.PermissionModel
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
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
	if in.TemplateBody != nil {
		in, out := &in.TemplateBody, &out.TemplateBody
		*out = new(string)
		**out = **in
	}
	if in.TemplateURL != nil {
		in, out := &in.TemplateURL, &out.TemplateURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackSetParameters.
func (in *StackSetParameters) DeepCopy() *StackSetParameters {
	if in == nil {
		return nil
	}
	out := new(StackSetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackSetSpec) DeepCopyInto(out *StackSetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackSetSpec.
func (in *StackSetSpec) DeepCopy() *StackSetSpec {
	if in == nil {
		return nil
	}
	out := new(StackSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackSetStatus) DeepCopyInto(out *StackSetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackSetStatus.
func (in *StackSetStatus) DeepCopy() *StackSetStatus {
	if in == nil {
		return nil
	}
	out := new(StackSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackSpec) DeepCopyInto(out *StackSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackSpec.
func (in *StackSpec) DeepCopy() *StackSpec {
	if in == nil {
		return nil
	}
	out := new(StackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackStatus) DeepCopyInto(out *StackStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackStatus.
func (in *StackStatus) DeepCopy() *StackStatus {
	if in == nil {
		return nil
	}
	out := new(StackStatus)
	in.DeepCopyInto(out)
	return out
}
