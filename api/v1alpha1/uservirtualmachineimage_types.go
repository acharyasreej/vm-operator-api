// Copyright (c) 2018-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// UserVirtualMachineImageSpec defines the desired state of VirtualMachineImage
type UserVirtualMachineImageSpec struct {
	// Type describes the type of the VirtualMachineImage. Currently, the only supported image is "OVF"
	Type string `json:"type"`

	// ImageSourceType describes the type of content source of the VirtualMachineImage.  The only Content Source
	// supported currently is the vSphere Content Library.
	// +optional
	ImageSourceType string `json:"imageSourceType,omitempty"`

	// ImageID is a unique identifier exposed by the provider of this VirtualMachineImage.
	ImageID string `json:"imageID"`

	// ProviderRef is a reference to a content provider object that describes a provider.
	ProviderRef ContentProviderReference `json:"providerRef"`

	// ProductInfo describes the attributes of the VirtualMachineImage relating to the product contained in the
	// image.
	// +optional
	ProductInfo VirtualMachineImageProductInfo `json:"productInfo,omitempty"`

	// OSInfo describes the attributes of the VirtualMachineImage relating to the Operating System contained in the
	// image.
	// +optional
	OSInfo VirtualMachineImageOSInfo `json:"osInfo,omitempty"`

	// OVFEnv describes the user configurable customization parameters of the VirtualMachineImage.
	// +optional
	OVFEnv map[string]OvfProperty `json:"ovfEnv,omitempty"`

	// HardwareVersion describes the virtual hardware version of the image
	// +optional
	HardwareVersion int32 `json:"hwVersion,omitempty"`

	// AnnotationData includes the annotation data
	// +optional
	AnnotationData string `json:"annotationData"`
}

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

	Spec   UserVirtualMachineImageSpec   `json:"spec,omitempty"`
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
