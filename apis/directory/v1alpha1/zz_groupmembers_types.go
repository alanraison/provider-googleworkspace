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

type GroupMembersObservation struct {

	// ETag of the resource.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group ID.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The members of the group
	Members []MembersObservation `json:"members,omitempty" tf:"members,omitempty"`
}

type GroupMembersParameters struct {

	// Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group ID.
	// +crossplane:generate:reference:type=Group
	// +crossplane:generate:reference:refFieldName=GroupRef
	// +crossplane:generate:reference:selectorFieldName=GroupSelector
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a Group to populate groupId.
	// +kubebuilder:validation:Optional
	GroupRef *v1.Reference `json:"groupRef,omitempty" tf:"-"`

	// Selector for a Group to populate groupId.
	// +kubebuilder:validation:Optional
	GroupSelector *v1.Selector `json:"groupSelector,omitempty" tf:"-"`

	// The members of the group
	// +kubebuilder:validation:Optional
	Members []MembersParameters `json:"members,omitempty" tf:"members,omitempty"`
}

type MembersObservation struct {

	// Defaults to `ALL_MAIL`. Defines mail delivery preferences of member. Acceptable values are:
	// - `ALL_MAIL`: All messages, delivered as soon as they arrive.
	// - `DAILY`: No more than one message a day.
	// - `DIGEST`: Up to 25 messages bundled into a single message.
	// - `DISABLED`: Remove subscription.
	// - `NONE`: No messages.
	DeliverySettings *string `json:"deliverySettings,omitempty" tf:"delivery_settings,omitempty"`

	// The member's email address. A member can be a user or another group. This property isrequired when adding a member to a group. The email must be unique and cannot be an alias of another group. If the email address is changed, the API automatically reflects the email address changes.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The unique ID of the group member. A member id can be used as a member request URI's memberKey.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Defaults to `MEMBER`. The member's role in a group. The API returns an error for cycles in group memberships. For example, if group1 is a member of group2, group2 cannot be a member of group1. Acceptable values are:
	// - `MANAGER`: This role is only available if the Google Groups for Business is enabled using the Admin Console. A `MANAGER` role can do everything done by an `OWNER` role except make a member an `OWNER` or delete the group. A group can have multiple `MANAGER` members.
	// - `MEMBER`: This role can subscribe to a group, view discussion archives, and view the group's membership list.
	// - `OWNER`: This role can send messages to the group, add or remove members, change member roles, change group's settings, and delete the group. An OWNER must be a member of the group. A group can have more than one OWNER.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Status of member.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Defaults to `USER`. The type of group member. Acceptable values are:
	// - `CUSTOMER`: The member represents all users in a domain. An email address is not returned and the ID returned is the customer ID.
	// - `GROUP`: The member is another group.
	// - `USER`: The member is a user.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type MembersParameters struct {

	// Defaults to `ALL_MAIL`. Defines mail delivery preferences of member. Acceptable values are:
	// - `ALL_MAIL`: All messages, delivered as soon as they arrive.
	// - `DAILY`: No more than one message a day.
	// - `DIGEST`: Up to 25 messages bundled into a single message.
	// - `DISABLED`: Remove subscription.
	// - `NONE`: No messages.
	// +kubebuilder:validation:Optional
	DeliverySettings *string `json:"deliverySettings,omitempty" tf:"delivery_settings,omitempty"`

	// The member's email address. A member can be a user or another group. This property isrequired when adding a member to a group. The email must be unique and cannot be an alias of another group. If the email address is changed, the API automatically reflects the email address changes.
	// +kubebuilder:validation:Required
	Email *string `json:"email" tf:"email,omitempty"`

	// Defaults to `MEMBER`. The member's role in a group. The API returns an error for cycles in group memberships. For example, if group1 is a member of group2, group2 cannot be a member of group1. Acceptable values are:
	// - `MANAGER`: This role is only available if the Google Groups for Business is enabled using the Admin Console. A `MANAGER` role can do everything done by an `OWNER` role except make a member an `OWNER` or delete the group. A group can have multiple `MANAGER` members.
	// - `MEMBER`: This role can subscribe to a group, view discussion archives, and view the group's membership list.
	// - `OWNER`: This role can send messages to the group, add or remove members, change member roles, change group's settings, and delete the group. An OWNER must be a member of the group. A group can have more than one OWNER.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Defaults to `USER`. The type of group member. Acceptable values are:
	// - `CUSTOMER`: The member represents all users in a domain. An email address is not returned and the ID returned is the customer ID.
	// - `GROUP`: The member is another group.
	// - `USER`: The member is a user.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// GroupMembersSpec defines the desired state of GroupMembers
type GroupMembersSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupMembersParameters `json:"forProvider"`
}

// GroupMembersStatus defines the observed state of GroupMembers.
type GroupMembersStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupMembersObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GroupMembers is the Schema for the GroupMemberss API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,googleworkspace}
type GroupMembers struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupMembersSpec   `json:"spec"`
	Status            GroupMembersStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupMembersList contains a list of GroupMemberss
type GroupMembersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupMembers `json:"items"`
}

// Repository type metadata.
var (
	GroupMembers_Kind             = "GroupMembers"
	GroupMembers_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupMembers_Kind}.String()
	GroupMembers_KindAPIVersion   = GroupMembers_Kind + "." + CRDGroupVersion.String()
	GroupMembers_GroupVersionKind = CRDGroupVersion.WithKind(GroupMembers_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupMembers{}, &GroupMembersList{})
}