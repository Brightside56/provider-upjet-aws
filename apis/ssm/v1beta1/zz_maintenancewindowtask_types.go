// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AutomationParametersInitParameters struct {

	// The version of an Automation document to use during task execution.
	DocumentVersion *string `json:"documentVersion,omitempty" tf:"document_version,omitempty"`

	// The parameters for the RUN_COMMAND task execution. Documented below.
	Parameter []AutomationParametersParameterInitParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`
}

type AutomationParametersObservation struct {

	// The version of an Automation document to use during task execution.
	DocumentVersion *string `json:"documentVersion,omitempty" tf:"document_version,omitempty"`

	// The parameters for the RUN_COMMAND task execution. Documented below.
	Parameter []AutomationParametersParameterObservation `json:"parameter,omitempty" tf:"parameter,omitempty"`
}

type AutomationParametersParameterInitParameters struct {

	// The name of the maintenance window task.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The array of strings.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type AutomationParametersParameterObservation struct {

	// The name of the maintenance window task.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The array of strings.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type AutomationParametersParameterParameters struct {

	// The name of the maintenance window task.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The array of strings.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type AutomationParametersParameters struct {

	// The version of an Automation document to use during task execution.
	// +kubebuilder:validation:Optional
	DocumentVersion *string `json:"documentVersion,omitempty" tf:"document_version,omitempty"`

	// The parameters for the RUN_COMMAND task execution. Documented below.
	// +kubebuilder:validation:Optional
	Parameter []AutomationParametersParameterParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`
}

type CloudwatchConfigInitParameters struct {

	// The name of the CloudWatch log group where you want to send command output. If you don't specify a group name, Systems Manager automatically creates a log group for you. The log group uses the following naming format: aws/ssm/SystemsManagerDocumentName.
	CloudwatchLogGroupName *string `json:"cloudwatchLogGroupName,omitempty" tf:"cloudwatch_log_group_name,omitempty"`

	// Enables Systems Manager to send command output to CloudWatch Logs.
	CloudwatchOutputEnabled *bool `json:"cloudwatchOutputEnabled,omitempty" tf:"cloudwatch_output_enabled,omitempty"`
}

type CloudwatchConfigObservation struct {

	// The name of the CloudWatch log group where you want to send command output. If you don't specify a group name, Systems Manager automatically creates a log group for you. The log group uses the following naming format: aws/ssm/SystemsManagerDocumentName.
	CloudwatchLogGroupName *string `json:"cloudwatchLogGroupName,omitempty" tf:"cloudwatch_log_group_name,omitempty"`

	// Enables Systems Manager to send command output to CloudWatch Logs.
	CloudwatchOutputEnabled *bool `json:"cloudwatchOutputEnabled,omitempty" tf:"cloudwatch_output_enabled,omitempty"`
}

type CloudwatchConfigParameters struct {

	// The name of the CloudWatch log group where you want to send command output. If you don't specify a group name, Systems Manager automatically creates a log group for you. The log group uses the following naming format: aws/ssm/SystemsManagerDocumentName.
	// +kubebuilder:validation:Optional
	CloudwatchLogGroupName *string `json:"cloudwatchLogGroupName,omitempty" tf:"cloudwatch_log_group_name,omitempty"`

	// Enables Systems Manager to send command output to CloudWatch Logs.
	// +kubebuilder:validation:Optional
	CloudwatchOutputEnabled *bool `json:"cloudwatchOutputEnabled,omitempty" tf:"cloudwatch_output_enabled,omitempty"`
}

type LambdaParametersInitParameters struct {

	// Pass client-specific information to the Lambda function that you are invoking.
	ClientContext *string `json:"clientContext,omitempty" tf:"client_context,omitempty"`

	// JSON to provide to your Lambda function as input.
	PayloadSecretRef *v1.SecretKeySelector `json:"payloadSecretRef,omitempty" tf:"-"`

	// Specify a Lambda function version or alias name.
	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier,omitempty"`
}

type LambdaParametersObservation struct {

	// Pass client-specific information to the Lambda function that you are invoking.
	ClientContext *string `json:"clientContext,omitempty" tf:"client_context,omitempty"`

	// Specify a Lambda function version or alias name.
	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier,omitempty"`
}

type LambdaParametersParameters struct {

	// Pass client-specific information to the Lambda function that you are invoking.
	// +kubebuilder:validation:Optional
	ClientContext *string `json:"clientContext,omitempty" tf:"client_context,omitempty"`

	// JSON to provide to your Lambda function as input.
	// +kubebuilder:validation:Optional
	PayloadSecretRef *v1.SecretKeySelector `json:"payloadSecretRef,omitempty" tf:"-"`

	// Specify a Lambda function version or alias name.
	// +kubebuilder:validation:Optional
	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier,omitempty"`
}

type MaintenanceWindowTaskInitParameters struct {

	// Indicates whether tasks should continue to run after the cutoff time specified in the maintenance windows is reached. Valid values are CONTINUE_TASK and CANCEL_TASK.
	CutoffBehavior *string `json:"cutoffBehavior,omitempty" tf:"cutoff_behavior,omitempty"`

	// The description of the maintenance window task.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The maximum number of targets this task can be run for in parallel.
	MaxConcurrency *string `json:"maxConcurrency,omitempty" tf:"max_concurrency,omitempty"`

	// The maximum number of errors allowed before this task stops being scheduled.
	MaxErrors *string `json:"maxErrors,omitempty" tf:"max_errors,omitempty"`

	// The name of the maintenance window task.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The role that should be assumed when executing the task. If a role is not provided, Systems Manager uses your account's service-linked role. If no service-linked role for Systems Manager exists in your account, it is created for you.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	ServiceRoleArn *string `json:"serviceRoleArn,omitempty" tf:"service_role_arn,omitempty"`

	// Reference to a Role in iam to populate serviceRoleArn.
	// +kubebuilder:validation:Optional
	ServiceRoleArnRef *v1.Reference `json:"serviceRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate serviceRoleArn.
	// +kubebuilder:validation:Optional
	ServiceRoleArnSelector *v1.Selector `json:"serviceRoleArnSelector,omitempty" tf:"-"`

	// The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
	Targets []MaintenanceWindowTaskTargetsInitParameters `json:"targets,omitempty" tf:"targets,omitempty"`

	// The ARN of the task to execute.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lambda/v1beta1.Function
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	TaskArn *string `json:"taskArn,omitempty" tf:"task_arn,omitempty"`

	// Reference to a Function in lambda to populate taskArn.
	// +kubebuilder:validation:Optional
	TaskArnRef *v1.Reference `json:"taskArnRef,omitempty" tf:"-"`

	// Selector for a Function in lambda to populate taskArn.
	// +kubebuilder:validation:Optional
	TaskArnSelector *v1.Selector `json:"taskArnSelector,omitempty" tf:"-"`

	// Configuration block with parameters for task execution.
	TaskInvocationParameters []TaskInvocationParametersInitParameters `json:"taskInvocationParameters,omitempty" tf:"task_invocation_parameters,omitempty"`

	// The type of task being registered. Valid values: AUTOMATION, LAMBDA, RUN_COMMAND or STEP_FUNCTIONS.
	TaskType *string `json:"taskType,omitempty" tf:"task_type,omitempty"`

	// The Id of the maintenance window to register the task with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ssm/v1beta1.MaintenanceWindow
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	WindowID *string `json:"windowId,omitempty" tf:"window_id,omitempty"`

	// Reference to a MaintenanceWindow in ssm to populate windowId.
	// +kubebuilder:validation:Optional
	WindowIDRef *v1.Reference `json:"windowIdRef,omitempty" tf:"-"`

	// Selector for a MaintenanceWindow in ssm to populate windowId.
	// +kubebuilder:validation:Optional
	WindowIDSelector *v1.Selector `json:"windowIdSelector,omitempty" tf:"-"`
}

type MaintenanceWindowTaskObservation struct {

	// The ARN of the maintenance window task.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Indicates whether tasks should continue to run after the cutoff time specified in the maintenance windows is reached. Valid values are CONTINUE_TASK and CANCEL_TASK.
	CutoffBehavior *string `json:"cutoffBehavior,omitempty" tf:"cutoff_behavior,omitempty"`

	// The description of the maintenance window task.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the maintenance window task.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The maximum number of targets this task can be run for in parallel.
	MaxConcurrency *string `json:"maxConcurrency,omitempty" tf:"max_concurrency,omitempty"`

	// The maximum number of errors allowed before this task stops being scheduled.
	MaxErrors *string `json:"maxErrors,omitempty" tf:"max_errors,omitempty"`

	// The name of the maintenance window task.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The role that should be assumed when executing the task. If a role is not provided, Systems Manager uses your account's service-linked role. If no service-linked role for Systems Manager exists in your account, it is created for you.
	ServiceRoleArn *string `json:"serviceRoleArn,omitempty" tf:"service_role_arn,omitempty"`

	// The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
	Targets []MaintenanceWindowTaskTargetsObservation `json:"targets,omitempty" tf:"targets,omitempty"`

	// The ARN of the task to execute.
	TaskArn *string `json:"taskArn,omitempty" tf:"task_arn,omitempty"`

	// Configuration block with parameters for task execution.
	TaskInvocationParameters []TaskInvocationParametersObservation `json:"taskInvocationParameters,omitempty" tf:"task_invocation_parameters,omitempty"`

	// The type of task being registered. Valid values: AUTOMATION, LAMBDA, RUN_COMMAND or STEP_FUNCTIONS.
	TaskType *string `json:"taskType,omitempty" tf:"task_type,omitempty"`

	// The Id of the maintenance window to register the task with.
	WindowID *string `json:"windowId,omitempty" tf:"window_id,omitempty"`

	// The ID of the maintenance window task.
	WindowTaskID *string `json:"windowTaskId,omitempty" tf:"window_task_id,omitempty"`
}

type MaintenanceWindowTaskParameters struct {

	// Indicates whether tasks should continue to run after the cutoff time specified in the maintenance windows is reached. Valid values are CONTINUE_TASK and CANCEL_TASK.
	// +kubebuilder:validation:Optional
	CutoffBehavior *string `json:"cutoffBehavior,omitempty" tf:"cutoff_behavior,omitempty"`

	// The description of the maintenance window task.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The maximum number of targets this task can be run for in parallel.
	// +kubebuilder:validation:Optional
	MaxConcurrency *string `json:"maxConcurrency,omitempty" tf:"max_concurrency,omitempty"`

	// The maximum number of errors allowed before this task stops being scheduled.
	// +kubebuilder:validation:Optional
	MaxErrors *string `json:"maxErrors,omitempty" tf:"max_errors,omitempty"`

	// The name of the maintenance window task.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The role that should be assumed when executing the task. If a role is not provided, Systems Manager uses your account's service-linked role. If no service-linked role for Systems Manager exists in your account, it is created for you.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ServiceRoleArn *string `json:"serviceRoleArn,omitempty" tf:"service_role_arn,omitempty"`

	// Reference to a Role in iam to populate serviceRoleArn.
	// +kubebuilder:validation:Optional
	ServiceRoleArnRef *v1.Reference `json:"serviceRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate serviceRoleArn.
	// +kubebuilder:validation:Optional
	ServiceRoleArnSelector *v1.Selector `json:"serviceRoleArnSelector,omitempty" tf:"-"`

	// The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
	// +kubebuilder:validation:Optional
	Targets []MaintenanceWindowTaskTargetsParameters `json:"targets,omitempty" tf:"targets,omitempty"`

	// The ARN of the task to execute.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lambda/v1beta1.Function
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	TaskArn *string `json:"taskArn,omitempty" tf:"task_arn,omitempty"`

	// Reference to a Function in lambda to populate taskArn.
	// +kubebuilder:validation:Optional
	TaskArnRef *v1.Reference `json:"taskArnRef,omitempty" tf:"-"`

	// Selector for a Function in lambda to populate taskArn.
	// +kubebuilder:validation:Optional
	TaskArnSelector *v1.Selector `json:"taskArnSelector,omitempty" tf:"-"`

	// Configuration block with parameters for task execution.
	// +kubebuilder:validation:Optional
	TaskInvocationParameters []TaskInvocationParametersParameters `json:"taskInvocationParameters,omitempty" tf:"task_invocation_parameters,omitempty"`

	// The type of task being registered. Valid values: AUTOMATION, LAMBDA, RUN_COMMAND or STEP_FUNCTIONS.
	// +kubebuilder:validation:Optional
	TaskType *string `json:"taskType,omitempty" tf:"task_type,omitempty"`

	// The Id of the maintenance window to register the task with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ssm/v1beta1.MaintenanceWindow
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WindowID *string `json:"windowId,omitempty" tf:"window_id,omitempty"`

	// Reference to a MaintenanceWindow in ssm to populate windowId.
	// +kubebuilder:validation:Optional
	WindowIDRef *v1.Reference `json:"windowIdRef,omitempty" tf:"-"`

	// Selector for a MaintenanceWindow in ssm to populate windowId.
	// +kubebuilder:validation:Optional
	WindowIDSelector *v1.Selector `json:"windowIdSelector,omitempty" tf:"-"`
}

type MaintenanceWindowTaskTargetsInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The array of strings.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta2.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`

	// References to Instance in ec2 to populate values.
	// +kubebuilder:validation:Optional
	ValuesRefs []v1.Reference `json:"valuesRefs,omitempty" tf:"-"`

	// Selector for a list of Instance in ec2 to populate values.
	// +kubebuilder:validation:Optional
	ValuesSelector *v1.Selector `json:"valuesSelector,omitempty" tf:"-"`
}

type MaintenanceWindowTaskTargetsObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The array of strings.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MaintenanceWindowTaskTargetsParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// The array of strings.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta2.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`

	// References to Instance in ec2 to populate values.
	// +kubebuilder:validation:Optional
	ValuesRefs []v1.Reference `json:"valuesRefs,omitempty" tf:"-"`

	// Selector for a list of Instance in ec2 to populate values.
	// +kubebuilder:validation:Optional
	ValuesSelector *v1.Selector `json:"valuesSelector,omitempty" tf:"-"`
}

type NotificationConfigInitParameters struct {

	// An Amazon Resource Name (ARN) for a Simple Notification Service (SNS) topic. Run Command pushes notifications about command status changes to this topic.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	NotificationArn *string `json:"notificationArn,omitempty" tf:"notification_arn,omitempty"`

	// Reference to a Topic in sns to populate notificationArn.
	// +kubebuilder:validation:Optional
	NotificationArnRef *v1.Reference `json:"notificationArnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate notificationArn.
	// +kubebuilder:validation:Optional
	NotificationArnSelector *v1.Selector `json:"notificationArnSelector,omitempty" tf:"-"`

	// The different events for which you can receive notifications. Valid values: All, InProgress, Success, TimedOut, Cancelled, and Failed
	NotificationEvents []*string `json:"notificationEvents,omitempty" tf:"notification_events,omitempty"`

	// When specified with Command, receive notification when the status of a command changes. When specified with Invocation, for commands sent to multiple instances, receive notification on a per-instance basis when the status of a command changes. Valid values: Command and Invocation
	NotificationType *string `json:"notificationType,omitempty" tf:"notification_type,omitempty"`
}

type NotificationConfigObservation struct {

	// An Amazon Resource Name (ARN) for a Simple Notification Service (SNS) topic. Run Command pushes notifications about command status changes to this topic.
	NotificationArn *string `json:"notificationArn,omitempty" tf:"notification_arn,omitempty"`

	// The different events for which you can receive notifications. Valid values: All, InProgress, Success, TimedOut, Cancelled, and Failed
	NotificationEvents []*string `json:"notificationEvents,omitempty" tf:"notification_events,omitempty"`

	// When specified with Command, receive notification when the status of a command changes. When specified with Invocation, for commands sent to multiple instances, receive notification on a per-instance basis when the status of a command changes. Valid values: Command and Invocation
	NotificationType *string `json:"notificationType,omitempty" tf:"notification_type,omitempty"`
}

type NotificationConfigParameters struct {

	// An Amazon Resource Name (ARN) for a Simple Notification Service (SNS) topic. Run Command pushes notifications about command status changes to this topic.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	NotificationArn *string `json:"notificationArn,omitempty" tf:"notification_arn,omitempty"`

	// Reference to a Topic in sns to populate notificationArn.
	// +kubebuilder:validation:Optional
	NotificationArnRef *v1.Reference `json:"notificationArnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate notificationArn.
	// +kubebuilder:validation:Optional
	NotificationArnSelector *v1.Selector `json:"notificationArnSelector,omitempty" tf:"-"`

	// The different events for which you can receive notifications. Valid values: All, InProgress, Success, TimedOut, Cancelled, and Failed
	// +kubebuilder:validation:Optional
	NotificationEvents []*string `json:"notificationEvents,omitempty" tf:"notification_events,omitempty"`

	// When specified with Command, receive notification when the status of a command changes. When specified with Invocation, for commands sent to multiple instances, receive notification on a per-instance basis when the status of a command changes. Valid values: Command and Invocation
	// +kubebuilder:validation:Optional
	NotificationType *string `json:"notificationType,omitempty" tf:"notification_type,omitempty"`
}

type RunCommandParametersInitParameters struct {

	// Configuration options for sending command output to CloudWatch Logs. Documented below.
	CloudwatchConfig []CloudwatchConfigInitParameters `json:"cloudwatchConfig,omitempty" tf:"cloudwatch_config,omitempty"`

	// Information about the command(s) to execute.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The SHA-256 or SHA-1 hash created by the system when the document was created. SHA-1 hashes have been deprecated.
	DocumentHash *string `json:"documentHash,omitempty" tf:"document_hash,omitempty"`

	// SHA-256 or SHA-1. SHA-1 hashes have been deprecated. Valid values: Sha256 and Sha1
	DocumentHashType *string `json:"documentHashType,omitempty" tf:"document_hash_type,omitempty"`

	// The version of an Automation document to use during task execution.
	DocumentVersion *string `json:"documentVersion,omitempty" tf:"document_version,omitempty"`

	// Configurations for sending notifications about command status changes on a per-instance basis. Documented below.
	NotificationConfig []NotificationConfigInitParameters `json:"notificationConfig,omitempty" tf:"notification_config,omitempty"`

	// The name of the Amazon S3 bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	OutputS3Bucket *string `json:"outputS3Bucket,omitempty" tf:"output_s3_bucket,omitempty"`

	// Reference to a Bucket in s3 to populate outputS3Bucket.
	// +kubebuilder:validation:Optional
	OutputS3BucketRef *v1.Reference `json:"outputS3BucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate outputS3Bucket.
	// +kubebuilder:validation:Optional
	OutputS3BucketSelector *v1.Selector `json:"outputS3BucketSelector,omitempty" tf:"-"`

	// The Amazon S3 bucket subfolder.
	OutputS3KeyPrefix *string `json:"outputS3KeyPrefix,omitempty" tf:"output_s3_key_prefix,omitempty"`

	// The parameters for the RUN_COMMAND task execution. Documented below.
	Parameter []RunCommandParametersParameterInitParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// The role that should be assumed when executing the task. If a role is not provided, Systems Manager uses your account's service-linked role. If no service-linked role for Systems Manager exists in your account, it is created for you.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	ServiceRoleArn *string `json:"serviceRoleArn,omitempty" tf:"service_role_arn,omitempty"`

	// Reference to a Role in iam to populate serviceRoleArn.
	// +kubebuilder:validation:Optional
	ServiceRoleArnRef *v1.Reference `json:"serviceRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate serviceRoleArn.
	// +kubebuilder:validation:Optional
	ServiceRoleArnSelector *v1.Selector `json:"serviceRoleArnSelector,omitempty" tf:"-"`

	// If this time is reached and the command has not already started executing, it doesn't run.
	TimeoutSeconds *float64 `json:"timeoutSeconds,omitempty" tf:"timeout_seconds,omitempty"`
}

type RunCommandParametersObservation struct {

	// Configuration options for sending command output to CloudWatch Logs. Documented below.
	CloudwatchConfig []CloudwatchConfigObservation `json:"cloudwatchConfig,omitempty" tf:"cloudwatch_config,omitempty"`

	// Information about the command(s) to execute.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The SHA-256 or SHA-1 hash created by the system when the document was created. SHA-1 hashes have been deprecated.
	DocumentHash *string `json:"documentHash,omitempty" tf:"document_hash,omitempty"`

	// SHA-256 or SHA-1. SHA-1 hashes have been deprecated. Valid values: Sha256 and Sha1
	DocumentHashType *string `json:"documentHashType,omitempty" tf:"document_hash_type,omitempty"`

	// The version of an Automation document to use during task execution.
	DocumentVersion *string `json:"documentVersion,omitempty" tf:"document_version,omitempty"`

	// Configurations for sending notifications about command status changes on a per-instance basis. Documented below.
	NotificationConfig []NotificationConfigObservation `json:"notificationConfig,omitempty" tf:"notification_config,omitempty"`

	// The name of the Amazon S3 bucket.
	OutputS3Bucket *string `json:"outputS3Bucket,omitempty" tf:"output_s3_bucket,omitempty"`

	// The Amazon S3 bucket subfolder.
	OutputS3KeyPrefix *string `json:"outputS3KeyPrefix,omitempty" tf:"output_s3_key_prefix,omitempty"`

	// The parameters for the RUN_COMMAND task execution. Documented below.
	Parameter []RunCommandParametersParameterObservation `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// The role that should be assumed when executing the task. If a role is not provided, Systems Manager uses your account's service-linked role. If no service-linked role for Systems Manager exists in your account, it is created for you.
	ServiceRoleArn *string `json:"serviceRoleArn,omitempty" tf:"service_role_arn,omitempty"`

	// If this time is reached and the command has not already started executing, it doesn't run.
	TimeoutSeconds *float64 `json:"timeoutSeconds,omitempty" tf:"timeout_seconds,omitempty"`
}

type RunCommandParametersParameterInitParameters struct {

	// The name of the maintenance window task.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The array of strings.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type RunCommandParametersParameterObservation struct {

	// The name of the maintenance window task.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The array of strings.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type RunCommandParametersParameterParameters struct {

	// The name of the maintenance window task.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The array of strings.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type RunCommandParametersParameters struct {

	// Configuration options for sending command output to CloudWatch Logs. Documented below.
	// +kubebuilder:validation:Optional
	CloudwatchConfig []CloudwatchConfigParameters `json:"cloudwatchConfig,omitempty" tf:"cloudwatch_config,omitempty"`

	// Information about the command(s) to execute.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The SHA-256 or SHA-1 hash created by the system when the document was created. SHA-1 hashes have been deprecated.
	// +kubebuilder:validation:Optional
	DocumentHash *string `json:"documentHash,omitempty" tf:"document_hash,omitempty"`

	// SHA-256 or SHA-1. SHA-1 hashes have been deprecated. Valid values: Sha256 and Sha1
	// +kubebuilder:validation:Optional
	DocumentHashType *string `json:"documentHashType,omitempty" tf:"document_hash_type,omitempty"`

	// The version of an Automation document to use during task execution.
	// +kubebuilder:validation:Optional
	DocumentVersion *string `json:"documentVersion,omitempty" tf:"document_version,omitempty"`

	// Configurations for sending notifications about command status changes on a per-instance basis. Documented below.
	// +kubebuilder:validation:Optional
	NotificationConfig []NotificationConfigParameters `json:"notificationConfig,omitempty" tf:"notification_config,omitempty"`

	// The name of the Amazon S3 bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	OutputS3Bucket *string `json:"outputS3Bucket,omitempty" tf:"output_s3_bucket,omitempty"`

	// Reference to a Bucket in s3 to populate outputS3Bucket.
	// +kubebuilder:validation:Optional
	OutputS3BucketRef *v1.Reference `json:"outputS3BucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate outputS3Bucket.
	// +kubebuilder:validation:Optional
	OutputS3BucketSelector *v1.Selector `json:"outputS3BucketSelector,omitempty" tf:"-"`

	// The Amazon S3 bucket subfolder.
	// +kubebuilder:validation:Optional
	OutputS3KeyPrefix *string `json:"outputS3KeyPrefix,omitempty" tf:"output_s3_key_prefix,omitempty"`

	// The parameters for the RUN_COMMAND task execution. Documented below.
	// +kubebuilder:validation:Optional
	Parameter []RunCommandParametersParameterParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// The role that should be assumed when executing the task. If a role is not provided, Systems Manager uses your account's service-linked role. If no service-linked role for Systems Manager exists in your account, it is created for you.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ServiceRoleArn *string `json:"serviceRoleArn,omitempty" tf:"service_role_arn,omitempty"`

	// Reference to a Role in iam to populate serviceRoleArn.
	// +kubebuilder:validation:Optional
	ServiceRoleArnRef *v1.Reference `json:"serviceRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate serviceRoleArn.
	// +kubebuilder:validation:Optional
	ServiceRoleArnSelector *v1.Selector `json:"serviceRoleArnSelector,omitempty" tf:"-"`

	// If this time is reached and the command has not already started executing, it doesn't run.
	// +kubebuilder:validation:Optional
	TimeoutSeconds *float64 `json:"timeoutSeconds,omitempty" tf:"timeout_seconds,omitempty"`
}

type StepFunctionsParametersInitParameters struct {

	// The inputs for the STEP_FUNCTION task.
	InputSecretRef *v1.SecretKeySelector `json:"inputSecretRef,omitempty" tf:"-"`

	// The name of the maintenance window task.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type StepFunctionsParametersObservation struct {

	// The name of the maintenance window task.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type StepFunctionsParametersParameters struct {

	// The inputs for the STEP_FUNCTION task.
	// +kubebuilder:validation:Optional
	InputSecretRef *v1.SecretKeySelector `json:"inputSecretRef,omitempty" tf:"-"`

	// The name of the maintenance window task.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type TaskInvocationParametersInitParameters struct {

	// The parameters for an AUTOMATION task type. Documented below.
	AutomationParameters []AutomationParametersInitParameters `json:"automationParameters,omitempty" tf:"automation_parameters,omitempty"`

	// The parameters for a LAMBDA task type. Documented below.
	LambdaParameters []LambdaParametersInitParameters `json:"lambdaParameters,omitempty" tf:"lambda_parameters,omitempty"`

	// The parameters for a RUN_COMMAND task type. Documented below.
	RunCommandParameters []RunCommandParametersInitParameters `json:"runCommandParameters,omitempty" tf:"run_command_parameters,omitempty"`

	// The parameters for a STEP_FUNCTIONS task type. Documented below.
	StepFunctionsParameters []StepFunctionsParametersInitParameters `json:"stepFunctionsParameters,omitempty" tf:"step_functions_parameters,omitempty"`
}

type TaskInvocationParametersObservation struct {

	// The parameters for an AUTOMATION task type. Documented below.
	AutomationParameters []AutomationParametersObservation `json:"automationParameters,omitempty" tf:"automation_parameters,omitempty"`

	// The parameters for a LAMBDA task type. Documented below.
	LambdaParameters []LambdaParametersObservation `json:"lambdaParameters,omitempty" tf:"lambda_parameters,omitempty"`

	// The parameters for a RUN_COMMAND task type. Documented below.
	RunCommandParameters []RunCommandParametersObservation `json:"runCommandParameters,omitempty" tf:"run_command_parameters,omitempty"`

	// The parameters for a STEP_FUNCTIONS task type. Documented below.
	StepFunctionsParameters []StepFunctionsParametersObservation `json:"stepFunctionsParameters,omitempty" tf:"step_functions_parameters,omitempty"`
}

type TaskInvocationParametersParameters struct {

	// The parameters for an AUTOMATION task type. Documented below.
	// +kubebuilder:validation:Optional
	AutomationParameters []AutomationParametersParameters `json:"automationParameters,omitempty" tf:"automation_parameters,omitempty"`

	// The parameters for a LAMBDA task type. Documented below.
	// +kubebuilder:validation:Optional
	LambdaParameters []LambdaParametersParameters `json:"lambdaParameters,omitempty" tf:"lambda_parameters,omitempty"`

	// The parameters for a RUN_COMMAND task type. Documented below.
	// +kubebuilder:validation:Optional
	RunCommandParameters []RunCommandParametersParameters `json:"runCommandParameters,omitempty" tf:"run_command_parameters,omitempty"`

	// The parameters for a STEP_FUNCTIONS task type. Documented below.
	// +kubebuilder:validation:Optional
	StepFunctionsParameters []StepFunctionsParametersParameters `json:"stepFunctionsParameters,omitempty" tf:"step_functions_parameters,omitempty"`
}

// MaintenanceWindowTaskSpec defines the desired state of MaintenanceWindowTask
type MaintenanceWindowTaskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MaintenanceWindowTaskParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider MaintenanceWindowTaskInitParameters `json:"initProvider,omitempty"`
}

// MaintenanceWindowTaskStatus defines the observed state of MaintenanceWindowTask.
type MaintenanceWindowTaskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MaintenanceWindowTaskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MaintenanceWindowTask is the Schema for the MaintenanceWindowTasks API. Provides an SSM Maintenance Window Task resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MaintenanceWindowTask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.taskType) || (has(self.initProvider) && has(self.initProvider.taskType))",message="spec.forProvider.taskType is a required parameter"
	Spec   MaintenanceWindowTaskSpec   `json:"spec"`
	Status MaintenanceWindowTaskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MaintenanceWindowTaskList contains a list of MaintenanceWindowTasks
type MaintenanceWindowTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MaintenanceWindowTask `json:"items"`
}

// Repository type metadata.
var (
	MaintenanceWindowTask_Kind             = "MaintenanceWindowTask"
	MaintenanceWindowTask_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MaintenanceWindowTask_Kind}.String()
	MaintenanceWindowTask_KindAPIVersion   = MaintenanceWindowTask_Kind + "." + CRDGroupVersion.String()
	MaintenanceWindowTask_GroupVersionKind = CRDGroupVersion.WithKind(MaintenanceWindowTask_Kind)
)

func init() {
	SchemeBuilder.Register(&MaintenanceWindowTask{}, &MaintenanceWindowTaskList{})
}
