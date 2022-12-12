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

type EventFilterObservation struct {
}

type EventFilterParameters struct {

	// The source of the events.
	// +kubebuilder:validation:Required
	Source *string `json:"source" tf:"source,omitempty"`
}

type EventIntegrationObservation struct {

	// The Amazon Resource Name (ARN) of the Event Integration.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The identifier of the Event Integration which is the name of the Event Integration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type EventIntegrationParameters struct {

	// Specifies the description of the Event Integration.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A block that defines the configuration information for the event filter. The Event Filter block is documented below.
	// +kubebuilder:validation:Required
	EventFilter []EventFilterParameters `json:"eventFilter" tf:"event_filter,omitempty"`

	// Specifies the EventBridge bus.
	// +kubebuilder:validation:Required
	EventbridgeBus *string `json:"eventbridgeBus" tf:"eventbridge_bus,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// EventIntegrationSpec defines the desired state of EventIntegration
type EventIntegrationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventIntegrationParameters `json:"forProvider"`
}

// EventIntegrationStatus defines the observed state of EventIntegration.
type EventIntegrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventIntegrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventIntegration is the Schema for the EventIntegrations API. Provides details about a specific Amazon AppIntegrations Event Integration
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EventIntegration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventIntegrationSpec   `json:"spec"`
	Status            EventIntegrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventIntegrationList contains a list of EventIntegrations
type EventIntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventIntegration `json:"items"`
}

// Repository type metadata.
var (
	EventIntegration_Kind             = "EventIntegration"
	EventIntegration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventIntegration_Kind}.String()
	EventIntegration_KindAPIVersion   = EventIntegration_Kind + "." + CRDGroupVersion.String()
	EventIntegration_GroupVersionKind = CRDGroupVersion.WithKind(EventIntegration_Kind)
)

func init() {
	SchemeBuilder.Register(&EventIntegration{}, &EventIntegrationList{})
}
