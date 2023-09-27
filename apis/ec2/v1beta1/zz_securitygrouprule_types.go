/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SecurityGroupRuleInitParameters_2 struct {

	// List of CIDR blocks. Cannot be specified with source_security_group_id or self.
	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`

	// Description of the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Start port (or ICMP type number if protocol is "icmp" or "icmpv6").
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// List of IPv6 CIDR blocks. Cannot be specified with source_security_group_id or self.
	IPv6CidrBlocks []*string `json:"ipv6CidrBlocks,omitempty" tf:"ipv6_cidr_blocks,omitempty"`

	// Protocol. If not icmp, icmpv6, tcp, udp, or all use the protocol number
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Whether the security group itself will be added as a source to this ingress rule. Cannot be specified with cidr_blocks, ipv6_cidr_blocks, or source_security_group_id.
	Self *bool `json:"self,omitempty" tf:"self,omitempty"`

	// End port (or ICMP code if protocol is "icmp").
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// Type of rule being created. Valid options are ingress (inbound)
	// or egress (outbound).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SecurityGroupRuleObservation_2 struct {

	// List of CIDR blocks. Cannot be specified with source_security_group_id or self.
	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`

	// Description of the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Start port (or ICMP type number if protocol is "icmp" or "icmpv6").
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// ID of the security group rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of IPv6 CIDR blocks. Cannot be specified with source_security_group_id or self.
	IPv6CidrBlocks []*string `json:"ipv6CidrBlocks,omitempty" tf:"ipv6_cidr_blocks,omitempty"`

	// List of Prefix List IDs.
	PrefixListIds []*string `json:"prefixListIds,omitempty" tf:"prefix_list_ids,omitempty"`

	// Protocol. If not icmp, icmpv6, tcp, udp, or all use the protocol number
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Security group to apply this rule to.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// If the aws_security_group_rule resource has a single source or destination then this is the AWS Security Group Rule resource ID. Otherwise it is empty.
	SecurityGroupRuleID *string `json:"securityGroupRuleId,omitempty" tf:"security_group_rule_id,omitempty"`

	// Whether the security group itself will be added as a source to this ingress rule. Cannot be specified with cidr_blocks, ipv6_cidr_blocks, or source_security_group_id.
	Self *bool `json:"self,omitempty" tf:"self,omitempty"`

	// Security group id to allow access to/from, depending on the type. Cannot be specified with cidr_blocks, ipv6_cidr_blocks, or self.
	SourceSecurityGroupID *string `json:"sourceSecurityGroupId,omitempty" tf:"source_security_group_id,omitempty"`

	// End port (or ICMP code if protocol is "icmp").
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// Type of rule being created. Valid options are ingress (inbound)
	// or egress (outbound).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SecurityGroupRuleParameters_2 struct {

	// List of CIDR blocks. Cannot be specified with source_security_group_id or self.
	// +kubebuilder:validation:Optional
	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`

	// Description of the rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Start port (or ICMP type number if protocol is "icmp" or "icmpv6").
	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// List of IPv6 CIDR blocks. Cannot be specified with source_security_group_id or self.
	// +kubebuilder:validation:Optional
	IPv6CidrBlocks []*string `json:"ipv6CidrBlocks,omitempty" tf:"ipv6_cidr_blocks,omitempty"`

	// References to ManagedPrefixList in ec2 to populate prefixListIds.
	// +kubebuilder:validation:Optional
	PrefixListIDRefs []v1.Reference `json:"prefixListIdRefs,omitempty" tf:"-"`

	// Selector for a list of ManagedPrefixList in ec2 to populate prefixListIds.
	// +kubebuilder:validation:Optional
	PrefixListIDSelector *v1.Selector `json:"prefixListIdSelector,omitempty" tf:"-"`

	// List of Prefix List IDs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.ManagedPrefixList
	// +crossplane:generate:reference:refFieldName=PrefixListIDRefs
	// +crossplane:generate:reference:selectorFieldName=PrefixListIDSelector
	// +kubebuilder:validation:Optional
	PrefixListIds []*string `json:"prefixListIds,omitempty" tf:"prefix_list_ids,omitempty"`

	// Protocol. If not icmp, icmpv6, tcp, udp, or all use the protocol number
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Security group to apply this rule to.
	// +crossplane:generate:reference:type=SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Reference to a SecurityGroup to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// Whether the security group itself will be added as a source to this ingress rule. Cannot be specified with cidr_blocks, ipv6_cidr_blocks, or source_security_group_id.
	// +kubebuilder:validation:Optional
	Self *bool `json:"self,omitempty" tf:"self,omitempty"`

	// Security group id to allow access to/from, depending on the type. Cannot be specified with cidr_blocks, ipv6_cidr_blocks, or self.
	// +crossplane:generate:reference:type=SecurityGroup
	// +kubebuilder:validation:Optional
	SourceSecurityGroupID *string `json:"sourceSecurityGroupId,omitempty" tf:"source_security_group_id,omitempty"`

	// Reference to a SecurityGroup to populate sourceSecurityGroupId.
	// +kubebuilder:validation:Optional
	SourceSecurityGroupIDRef *v1.Reference `json:"sourceSecurityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup to populate sourceSecurityGroupId.
	// +kubebuilder:validation:Optional
	SourceSecurityGroupIDSelector *v1.Selector `json:"sourceSecurityGroupIdSelector,omitempty" tf:"-"`

	// End port (or ICMP code if protocol is "icmp").
	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// Type of rule being created. Valid options are ingress (inbound)
	// or egress (outbound).
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// SecurityGroupRuleSpec defines the desired state of SecurityGroupRule
type SecurityGroupRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityGroupRuleParameters_2 `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SecurityGroupRuleInitParameters_2 `json:"initProvider,omitempty"`
}

// SecurityGroupRuleStatus defines the observed state of SecurityGroupRule.
type SecurityGroupRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityGroupRuleObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupRule is the Schema for the SecurityGroupRules API. Provides an security group rule resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SecurityGroupRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.fromPort) || (has(self.initProvider) && has(self.initProvider.fromPort))",message="spec.forProvider.fromPort is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.toPort) || (has(self.initProvider) && has(self.initProvider.toPort))",message="spec.forProvider.toPort is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   SecurityGroupRuleSpec   `json:"spec"`
	Status SecurityGroupRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupRuleList contains a list of SecurityGroupRules
type SecurityGroupRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroupRule `json:"items"`
}

// Repository type metadata.
var (
	SecurityGroupRule_Kind             = "SecurityGroupRule"
	SecurityGroupRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityGroupRule_Kind}.String()
	SecurityGroupRule_KindAPIVersion   = SecurityGroupRule_Kind + "." + CRDGroupVersion.String()
	SecurityGroupRule_GroupVersionKind = CRDGroupVersion.WithKind(SecurityGroupRule_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityGroupRule{}, &SecurityGroupRuleList{})
}
