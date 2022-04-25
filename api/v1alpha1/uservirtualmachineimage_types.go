// Copyright (c) 2018-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=uservmimage
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="ContentSourceName",type="string",JSONPath=".spec.providerRef.name"
// +kubebuilder:printcolumn:name="Version",type="string",JSONPath=".spec.productInfo.version"
// +kubebuilder:printcolumn:name="OsType",type="string",JSONPath=".spec.osInfo.type"
// +kubebuilder:printcolumn:name="Format",type="string",JSONPath=".spec.type"
// +kubebuilder:printcolumn:name="ImageSupported",type="boolean",priority=1,JSONPath=".status.imageSupported"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// UserVirtualMachineImage is the Schema for the uservirtualmachineimages API
// A UserVirtualMachineImage represents a UserVirtualMachine image (e.g. VM template) that can be used as the base image
// for creating a VirtualMachine instance.  The UserVirtualMachineImage is a required field of the VirtualMachine
// spec.  Currently, UserVirtualMachineImages are immutable to end users.
type UserVirtualMachineImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineImageSpec   `json:"spec,omitempty"`
	Status VirtualMachineImageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineImageList contains a list of VirtualMachineImage
type UserVirtualMachineImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserVirtualMachineImage `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&UserVirtualMachineImage{}, &UserVirtualMachineImageList{})
}
