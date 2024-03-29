//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Settings) DeepCopyInto(out *Settings) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Settings.
func (in *Settings) DeepCopy() *Settings {
	if in == nil {
		return nil
	}
	out := new(Settings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Settings) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingsList) DeepCopyInto(out *SettingsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Settings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingsList.
func (in *SettingsList) DeepCopy() *SettingsList {
	if in == nil {
		return nil
	}
	out := new(SettingsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SettingsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingsObservation) DeepCopyInto(out *SettingsObservation) {
	*out = *in
	if in.AllowExternalMembers != nil {
		in, out := &in.AllowExternalMembers, &out.AllowExternalMembers
		*out = new(bool)
		**out = **in
	}
	if in.AllowWebPosting != nil {
		in, out := &in.AllowWebPosting, &out.AllowWebPosting
		*out = new(bool)
		**out = **in
	}
	if in.ArchiveOnly != nil {
		in, out := &in.ArchiveOnly, &out.ArchiveOnly
		*out = new(bool)
		**out = **in
	}
	if in.CustomFooterText != nil {
		in, out := &in.CustomFooterText, &out.CustomFooterText
		*out = new(string)
		**out = **in
	}
	if in.CustomReplyTo != nil {
		in, out := &in.CustomReplyTo, &out.CustomReplyTo
		*out = new(string)
		**out = **in
	}
	if in.CustomRolesEnabledForSettingsToBeMerged != nil {
		in, out := &in.CustomRolesEnabledForSettingsToBeMerged, &out.CustomRolesEnabledForSettingsToBeMerged
		*out = new(bool)
		**out = **in
	}
	if in.DefaultMessageDenyNotificationText != nil {
		in, out := &in.DefaultMessageDenyNotificationText, &out.DefaultMessageDenyNotificationText
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.EnableCollaborativeInbox != nil {
		in, out := &in.EnableCollaborativeInbox, &out.EnableCollaborativeInbox
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IncludeCustomFooter != nil {
		in, out := &in.IncludeCustomFooter, &out.IncludeCustomFooter
		*out = new(bool)
		**out = **in
	}
	if in.IncludeInGlobalAddressList != nil {
		in, out := &in.IncludeInGlobalAddressList, &out.IncludeInGlobalAddressList
		*out = new(bool)
		**out = **in
	}
	if in.IsArchived != nil {
		in, out := &in.IsArchived, &out.IsArchived
		*out = new(bool)
		**out = **in
	}
	if in.MembersCanPostAsTheGroup != nil {
		in, out := &in.MembersCanPostAsTheGroup, &out.MembersCanPostAsTheGroup
		*out = new(bool)
		**out = **in
	}
	if in.MessageModerationLevel != nil {
		in, out := &in.MessageModerationLevel, &out.MessageModerationLevel
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrimaryLanguage != nil {
		in, out := &in.PrimaryLanguage, &out.PrimaryLanguage
		*out = new(string)
		**out = **in
	}
	if in.ReplyTo != nil {
		in, out := &in.ReplyTo, &out.ReplyTo
		*out = new(string)
		**out = **in
	}
	if in.SendMessageDenyNotification != nil {
		in, out := &in.SendMessageDenyNotification, &out.SendMessageDenyNotification
		*out = new(bool)
		**out = **in
	}
	if in.SpamModerationLevel != nil {
		in, out := &in.SpamModerationLevel, &out.SpamModerationLevel
		*out = new(string)
		**out = **in
	}
	if in.WhoCanAssistContent != nil {
		in, out := &in.WhoCanAssistContent, &out.WhoCanAssistContent
		*out = new(string)
		**out = **in
	}
	if in.WhoCanContactOwner != nil {
		in, out := &in.WhoCanContactOwner, &out.WhoCanContactOwner
		*out = new(string)
		**out = **in
	}
	if in.WhoCanDiscoverGroup != nil {
		in, out := &in.WhoCanDiscoverGroup, &out.WhoCanDiscoverGroup
		*out = new(string)
		**out = **in
	}
	if in.WhoCanJoin != nil {
		in, out := &in.WhoCanJoin, &out.WhoCanJoin
		*out = new(string)
		**out = **in
	}
	if in.WhoCanLeaveGroup != nil {
		in, out := &in.WhoCanLeaveGroup, &out.WhoCanLeaveGroup
		*out = new(string)
		**out = **in
	}
	if in.WhoCanModerateContent != nil {
		in, out := &in.WhoCanModerateContent, &out.WhoCanModerateContent
		*out = new(string)
		**out = **in
	}
	if in.WhoCanModerateMembers != nil {
		in, out := &in.WhoCanModerateMembers, &out.WhoCanModerateMembers
		*out = new(string)
		**out = **in
	}
	if in.WhoCanPostMessage != nil {
		in, out := &in.WhoCanPostMessage, &out.WhoCanPostMessage
		*out = new(string)
		**out = **in
	}
	if in.WhoCanViewGroup != nil {
		in, out := &in.WhoCanViewGroup, &out.WhoCanViewGroup
		*out = new(string)
		**out = **in
	}
	if in.WhoCanViewMembership != nil {
		in, out := &in.WhoCanViewMembership, &out.WhoCanViewMembership
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingsObservation.
func (in *SettingsObservation) DeepCopy() *SettingsObservation {
	if in == nil {
		return nil
	}
	out := new(SettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingsParameters) DeepCopyInto(out *SettingsParameters) {
	*out = *in
	if in.AllowExternalMembers != nil {
		in, out := &in.AllowExternalMembers, &out.AllowExternalMembers
		*out = new(bool)
		**out = **in
	}
	if in.AllowWebPosting != nil {
		in, out := &in.AllowWebPosting, &out.AllowWebPosting
		*out = new(bool)
		**out = **in
	}
	if in.ArchiveOnly != nil {
		in, out := &in.ArchiveOnly, &out.ArchiveOnly
		*out = new(bool)
		**out = **in
	}
	if in.CustomFooterText != nil {
		in, out := &in.CustomFooterText, &out.CustomFooterText
		*out = new(string)
		**out = **in
	}
	if in.CustomReplyTo != nil {
		in, out := &in.CustomReplyTo, &out.CustomReplyTo
		*out = new(string)
		**out = **in
	}
	if in.DefaultMessageDenyNotificationText != nil {
		in, out := &in.DefaultMessageDenyNotificationText, &out.DefaultMessageDenyNotificationText
		*out = new(string)
		**out = **in
	}
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.EnableCollaborativeInbox != nil {
		in, out := &in.EnableCollaborativeInbox, &out.EnableCollaborativeInbox
		*out = new(bool)
		**out = **in
	}
	if in.GroupRef != nil {
		in, out := &in.GroupRef, &out.GroupRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.GroupSelector != nil {
		in, out := &in.GroupSelector, &out.GroupSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.IncludeCustomFooter != nil {
		in, out := &in.IncludeCustomFooter, &out.IncludeCustomFooter
		*out = new(bool)
		**out = **in
	}
	if in.IncludeInGlobalAddressList != nil {
		in, out := &in.IncludeInGlobalAddressList, &out.IncludeInGlobalAddressList
		*out = new(bool)
		**out = **in
	}
	if in.IsArchived != nil {
		in, out := &in.IsArchived, &out.IsArchived
		*out = new(bool)
		**out = **in
	}
	if in.MembersCanPostAsTheGroup != nil {
		in, out := &in.MembersCanPostAsTheGroup, &out.MembersCanPostAsTheGroup
		*out = new(bool)
		**out = **in
	}
	if in.MessageModerationLevel != nil {
		in, out := &in.MessageModerationLevel, &out.MessageModerationLevel
		*out = new(string)
		**out = **in
	}
	if in.PrimaryLanguage != nil {
		in, out := &in.PrimaryLanguage, &out.PrimaryLanguage
		*out = new(string)
		**out = **in
	}
	if in.ReplyTo != nil {
		in, out := &in.ReplyTo, &out.ReplyTo
		*out = new(string)
		**out = **in
	}
	if in.SendMessageDenyNotification != nil {
		in, out := &in.SendMessageDenyNotification, &out.SendMessageDenyNotification
		*out = new(bool)
		**out = **in
	}
	if in.SpamModerationLevel != nil {
		in, out := &in.SpamModerationLevel, &out.SpamModerationLevel
		*out = new(string)
		**out = **in
	}
	if in.WhoCanAssistContent != nil {
		in, out := &in.WhoCanAssistContent, &out.WhoCanAssistContent
		*out = new(string)
		**out = **in
	}
	if in.WhoCanContactOwner != nil {
		in, out := &in.WhoCanContactOwner, &out.WhoCanContactOwner
		*out = new(string)
		**out = **in
	}
	if in.WhoCanDiscoverGroup != nil {
		in, out := &in.WhoCanDiscoverGroup, &out.WhoCanDiscoverGroup
		*out = new(string)
		**out = **in
	}
	if in.WhoCanJoin != nil {
		in, out := &in.WhoCanJoin, &out.WhoCanJoin
		*out = new(string)
		**out = **in
	}
	if in.WhoCanLeaveGroup != nil {
		in, out := &in.WhoCanLeaveGroup, &out.WhoCanLeaveGroup
		*out = new(string)
		**out = **in
	}
	if in.WhoCanModerateContent != nil {
		in, out := &in.WhoCanModerateContent, &out.WhoCanModerateContent
		*out = new(string)
		**out = **in
	}
	if in.WhoCanModerateMembers != nil {
		in, out := &in.WhoCanModerateMembers, &out.WhoCanModerateMembers
		*out = new(string)
		**out = **in
	}
	if in.WhoCanPostMessage != nil {
		in, out := &in.WhoCanPostMessage, &out.WhoCanPostMessage
		*out = new(string)
		**out = **in
	}
	if in.WhoCanViewGroup != nil {
		in, out := &in.WhoCanViewGroup, &out.WhoCanViewGroup
		*out = new(string)
		**out = **in
	}
	if in.WhoCanViewMembership != nil {
		in, out := &in.WhoCanViewMembership, &out.WhoCanViewMembership
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingsParameters.
func (in *SettingsParameters) DeepCopy() *SettingsParameters {
	if in == nil {
		return nil
	}
	out := new(SettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingsSpec) DeepCopyInto(out *SettingsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingsSpec.
func (in *SettingsSpec) DeepCopy() *SettingsSpec {
	if in == nil {
		return nil
	}
	out := new(SettingsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingsStatus) DeepCopyInto(out *SettingsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingsStatus.
func (in *SettingsStatus) DeepCopy() *SettingsStatus {
	if in == nil {
		return nil
	}
	out := new(SettingsStatus)
	in.DeepCopyInto(out)
	return out
}
