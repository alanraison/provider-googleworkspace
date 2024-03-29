/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GroupObservation struct {

	// Value is true if this group was created by an administrator rather than a user.
	AdminCreated *bool `json:"adminCreated,omitempty" tf:"admin_created,omitempty"`

	// asps.list of group's email addresses.
	Aliases []*string `json:"aliases,omitempty" tf:"aliases,omitempty"`

	// An extended description to help users determine the purpose of a group.For example, you can include information about who should join the group,the types of messages to send to the group, links to FAQs about the group, or related groups.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The number of users that are direct members of the group.If a group is a member (child) of this group (the parent),members of the child group are not counted in the directMembersCount property of the parent group.
	DirectMembersCount *float64 `json:"directMembersCount,omitempty" tf:"direct_members_count,omitempty"`

	// The group's email address. If your account has multiple domains,select the appropriate domain for the email address. The email must be unique.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// ETag of the resource.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The group's display name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// asps.list of the group's non-editable alias email addresses that are outside of the account's primary domain or subdomains. These are functioning email addresses used by the group.
	NonEditableAliases []*string `json:"nonEditableAliases,omitempty" tf:"non_editable_aliases,omitempty"`
}

type GroupParameters struct {

	// asps.list of group's email addresses.
	// +kubebuilder:validation:Optional
	Aliases []*string `json:"aliases,omitempty" tf:"aliases,omitempty"`

	// An extended description to help users determine the purpose of a group.For example, you can include information about who should join the group,the types of messages to send to the group, links to FAQs about the group, or related groups.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The group's email address. If your account has multiple domains,select the appropriate domain for the email address. The email must be unique.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The group's display name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupParameters `json:"forProvider"`
}

// GroupStatus defines the observed state of Group.
type GroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Group is the Schema for the Groups API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,googleworkspace}
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.email)",message="email is a required parameter"
	Spec   GroupSpec   `json:"spec"`
	Status GroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupList contains a list of Groups
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

// Repository type metadata.
var (
	Group_Kind             = "Group"
	Group_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Group_Kind}.String()
	Group_KindAPIVersion   = Group_Kind + "." + CRDGroupVersion.String()
	Group_GroupVersionKind = CRDGroupVersion.WithKind(Group_Kind)
)

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}
