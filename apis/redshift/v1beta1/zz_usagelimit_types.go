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

type UsageLimitInitParameters struct {

	// The limit amount. If time-based, this amount is in minutes. If data-based, this amount is in terabytes (TB). The value must be a positive number.
	Amount *float64 `json:"amount,omitempty" tf:"amount,omitempty"`

	// The action that Amazon Redshift takes when the limit is reached. The default is log. Valid values are log, emit-metric, and disable.
	BreachAction *string `json:"breachAction,omitempty" tf:"breach_action,omitempty"`

	// The Amazon Redshift feature that you want to limit. Valid values are spectrum, concurrency-scaling, and cross-region-datasharing.
	FeatureType *string `json:"featureType,omitempty" tf:"feature_type,omitempty"`

	// The type of limit. Depending on the feature type, this can be based on a time duration or data size. If FeatureType is spectrum, then LimitType must be data-scanned. If FeatureType is concurrency-scaling, then LimitType must be time. If FeatureType is cross-region-datasharing, then LimitType must be data-scanned. Valid values are data-scanned, and time.
	LimitType *string `json:"limitType,omitempty" tf:"limit_type,omitempty"`

	// The time period that the amount applies to. A weekly period begins on Sunday. The default is monthly. Valid values are daily, weekly, and monthly.
	Period *string `json:"period,omitempty" tf:"period,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type UsageLimitObservation struct {

	// The limit amount. If time-based, this amount is in minutes. If data-based, this amount is in terabytes (TB). The value must be a positive number.
	Amount *float64 `json:"amount,omitempty" tf:"amount,omitempty"`

	// Amazon Resource Name (ARN) of the Redshift Usage Limit.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The action that Amazon Redshift takes when the limit is reached. The default is log. Valid values are log, emit-metric, and disable.
	BreachAction *string `json:"breachAction,omitempty" tf:"breach_action,omitempty"`

	// The identifier of the cluster that you want to limit usage.
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// The Amazon Redshift feature that you want to limit. Valid values are spectrum, concurrency-scaling, and cross-region-datasharing.
	FeatureType *string `json:"featureType,omitempty" tf:"feature_type,omitempty"`

	// The Redshift Usage Limit ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type of limit. Depending on the feature type, this can be based on a time duration or data size. If FeatureType is spectrum, then LimitType must be data-scanned. If FeatureType is concurrency-scaling, then LimitType must be time. If FeatureType is cross-region-datasharing, then LimitType must be data-scanned. Valid values are data-scanned, and time.
	LimitType *string `json:"limitType,omitempty" tf:"limit_type,omitempty"`

	// The time period that the amount applies to. A weekly period begins on Sunday. The default is monthly. Valid values are daily, weekly, and monthly.
	Period *string `json:"period,omitempty" tf:"period,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type UsageLimitParameters struct {

	// The limit amount. If time-based, this amount is in minutes. If data-based, this amount is in terabytes (TB). The value must be a positive number.
	// +kubebuilder:validation:Optional
	Amount *float64 `json:"amount,omitempty" tf:"amount,omitempty"`

	// The action that Amazon Redshift takes when the limit is reached. The default is log. Valid values are log, emit-metric, and disable.
	// +kubebuilder:validation:Optional
	BreachAction *string `json:"breachAction,omitempty" tf:"breach_action,omitempty"`

	// The identifier of the cluster that you want to limit usage.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/redshift/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// Reference to a Cluster in redshift to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierRef *v1.Reference `json:"clusterIdentifierRef,omitempty" tf:"-"`

	// Selector for a Cluster in redshift to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierSelector *v1.Selector `json:"clusterIdentifierSelector,omitempty" tf:"-"`

	// The Amazon Redshift feature that you want to limit. Valid values are spectrum, concurrency-scaling, and cross-region-datasharing.
	// +kubebuilder:validation:Optional
	FeatureType *string `json:"featureType,omitempty" tf:"feature_type,omitempty"`

	// The type of limit. Depending on the feature type, this can be based on a time duration or data size. If FeatureType is spectrum, then LimitType must be data-scanned. If FeatureType is concurrency-scaling, then LimitType must be time. If FeatureType is cross-region-datasharing, then LimitType must be data-scanned. Valid values are data-scanned, and time.
	// +kubebuilder:validation:Optional
	LimitType *string `json:"limitType,omitempty" tf:"limit_type,omitempty"`

	// The time period that the amount applies to. A weekly period begins on Sunday. The default is monthly. Valid values are daily, weekly, and monthly.
	// +kubebuilder:validation:Optional
	Period *string `json:"period,omitempty" tf:"period,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// UsageLimitSpec defines the desired state of UsageLimit
type UsageLimitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UsageLimitParameters `json:"forProvider"`
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
	InitProvider UsageLimitInitParameters `json:"initProvider,omitempty"`
}

// UsageLimitStatus defines the observed state of UsageLimit.
type UsageLimitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UsageLimitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UsageLimit is the Schema for the UsageLimits API. Provides a Redshift Usage Limit resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UsageLimit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.amount) || (has(self.initProvider) && has(self.initProvider.amount))",message="spec.forProvider.amount is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.featureType) || (has(self.initProvider) && has(self.initProvider.featureType))",message="spec.forProvider.featureType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.limitType) || (has(self.initProvider) && has(self.initProvider.limitType))",message="spec.forProvider.limitType is a required parameter"
	Spec   UsageLimitSpec   `json:"spec"`
	Status UsageLimitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UsageLimitList contains a list of UsageLimits
type UsageLimitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UsageLimit `json:"items"`
}

// Repository type metadata.
var (
	UsageLimit_Kind             = "UsageLimit"
	UsageLimit_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UsageLimit_Kind}.String()
	UsageLimit_KindAPIVersion   = UsageLimit_Kind + "." + CRDGroupVersion.String()
	UsageLimit_GroupVersionKind = CRDGroupVersion.WithKind(UsageLimit_Kind)
)

func init() {
	SchemeBuilder.Register(&UsageLimit{}, &UsageLimitList{})
}
