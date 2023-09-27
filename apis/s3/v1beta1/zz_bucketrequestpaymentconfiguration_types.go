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

type BucketRequestPaymentConfigurationInitParameters struct {

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Specifies who pays for the download and request fees. Valid values: BucketOwner, Requester.
	Payer *string `json:"payer,omitempty" tf:"payer,omitempty"`
}

type BucketRequestPaymentConfigurationObservation struct {

	// Name of the bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// The bucket or bucket and expected_bucket_owner separated by a comma (,) if the latter is provided.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies who pays for the download and request fees. Valid values: BucketOwner, Requester.
	Payer *string `json:"payer,omitempty" tf:"payer,omitempty"`
}

type BucketRequestPaymentConfigurationParameters struct {

	// Name of the bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Account ID of the expected bucket owner.
	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Specifies who pays for the download and request fees. Valid values: BucketOwner, Requester.
	// +kubebuilder:validation:Optional
	Payer *string `json:"payer,omitempty" tf:"payer,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// BucketRequestPaymentConfigurationSpec defines the desired state of BucketRequestPaymentConfiguration
type BucketRequestPaymentConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketRequestPaymentConfigurationParameters `json:"forProvider"`
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
	InitProvider BucketRequestPaymentConfigurationInitParameters `json:"initProvider,omitempty"`
}

// BucketRequestPaymentConfigurationStatus defines the observed state of BucketRequestPaymentConfiguration.
type BucketRequestPaymentConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketRequestPaymentConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketRequestPaymentConfiguration is the Schema for the BucketRequestPaymentConfigurations API. Provides an S3 bucket request payment configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketRequestPaymentConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.payer) || (has(self.initProvider) && has(self.initProvider.payer))",message="spec.forProvider.payer is a required parameter"
	Spec   BucketRequestPaymentConfigurationSpec   `json:"spec"`
	Status BucketRequestPaymentConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketRequestPaymentConfigurationList contains a list of BucketRequestPaymentConfigurations
type BucketRequestPaymentConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketRequestPaymentConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketRequestPaymentConfiguration_Kind             = "BucketRequestPaymentConfiguration"
	BucketRequestPaymentConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketRequestPaymentConfiguration_Kind}.String()
	BucketRequestPaymentConfiguration_KindAPIVersion   = BucketRequestPaymentConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketRequestPaymentConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketRequestPaymentConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketRequestPaymentConfiguration{}, &BucketRequestPaymentConfigurationList{})
}
