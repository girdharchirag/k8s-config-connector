// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ResourcepolicyDailySchedule struct {
	/* Immutable. The number of days between snapshots. */
	DaysInCycle int `json:"daysInCycle"`

	/* Immutable. This must be in UTC format that resolves to one of
	00:00, 04:00, 08:00, 12:00, 16:00, or 20:00. For example,
	both 13:00-5 and 08:00 are valid. */
	StartTime string `json:"startTime"`
}

type ResourcepolicyDayOfWeeks struct {
	/* Immutable. The day of the week to create the snapshot. e.g. MONDAY Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]. */
	Day string `json:"day"`

	/* Immutable. Time within the window to start the operations.
	It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT. */
	StartTime string `json:"startTime"`
}

type ResourcepolicyGroupPlacementPolicy struct {
	/* Immutable. The number of availability domains instances will be spread across. If two instances are in different
	availability domain, they will not be put in the same low latency network. */
	// +optional
	AvailabilityDomainCount *int `json:"availabilityDomainCount,omitempty"`

	/* Immutable. Collocation specifies whether to place VMs inside the same availability domain on the same low-latency network.
	Specify 'COLLOCATED' to enable collocation. Can only be specified with 'vm_count'. If compute instances are created
	with a COLLOCATED policy, then exactly 'vm_count' instances must be created at the same time with the resource policy
	attached. Possible values: ["COLLOCATED"]. */
	// +optional
	Collocation *string `json:"collocation,omitempty"`

	/* Immutable. Number of vms in this placement group. */
	// +optional
	VmCount *int `json:"vmCount,omitempty"`
}

type ResourcepolicyHourlySchedule struct {
	/* Immutable. The number of hours between snapshots. */
	HoursInCycle int `json:"hoursInCycle"`

	/* Immutable. Time within the window to start the operations.
	It must be in an hourly format "HH:MM",
	where HH : [00-23] and MM : [00] GMT.
	eg: 21:00. */
	StartTime string `json:"startTime"`
}

type ResourcepolicyInstanceSchedulePolicy struct {
	/* Immutable. The expiration time of the schedule. The timestamp is an RFC3339 string. */
	// +optional
	ExpirationTime *string `json:"expirationTime,omitempty"`

	/* Immutable. The start time of the schedule. The timestamp is an RFC3339 string. */
	// +optional
	StartTime *string `json:"startTime,omitempty"`

	/* Immutable. Specifies the time zone to be used in interpreting the schedule. The value of this field must be a time zone name
	from the tz database: http://en.wikipedia.org/wiki/Tz_database. */
	TimeZone string `json:"timeZone"`

	/* Immutable. Specifies the schedule for starting instances. */
	// +optional
	VmStartSchedule *ResourcepolicyVmStartSchedule `json:"vmStartSchedule,omitempty"`

	/* Immutable. Specifies the schedule for stopping instances. */
	// +optional
	VmStopSchedule *ResourcepolicyVmStopSchedule `json:"vmStopSchedule,omitempty"`
}

type ResourcepolicyRetentionPolicy struct {
	/* Immutable. Maximum age of the snapshot that is allowed to be kept. */
	MaxRetentionDays int `json:"maxRetentionDays"`

	/* Immutable. Specifies the behavior to apply to scheduled snapshots when
	the source disk is deleted. Default value: "KEEP_AUTO_SNAPSHOTS" Possible values: ["KEEP_AUTO_SNAPSHOTS", "APPLY_RETENTION_POLICY"]. */
	// +optional
	OnSourceDiskDelete *string `json:"onSourceDiskDelete,omitempty"`
}

type ResourcepolicySchedule struct {
	/* Immutable. The policy will execute every nth day at the specified time. */
	// +optional
	DailySchedule *ResourcepolicyDailySchedule `json:"dailySchedule,omitempty"`

	/* Immutable. The policy will execute every nth hour starting at the specified time. */
	// +optional
	HourlySchedule *ResourcepolicyHourlySchedule `json:"hourlySchedule,omitempty"`

	/* Immutable. Allows specifying a snapshot time for each day of the week. */
	// +optional
	WeeklySchedule *ResourcepolicyWeeklySchedule `json:"weeklySchedule,omitempty"`
}

type ResourcepolicySnapshotProperties struct {
	/* Immutable. Whether to perform a 'guest aware' snapshot. */
	// +optional
	GuestFlush *bool `json:"guestFlush,omitempty"`

	/* Immutable. A set of key-value pairs. */
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	/* Immutable. Cloud Storage bucket location to store the auto snapshot
	(regional or multi-regional). */
	// +optional
	StorageLocations []string `json:"storageLocations,omitempty"`
}

type ResourcepolicySnapshotSchedulePolicy struct {
	/* Immutable. Retention policy applied to snapshots created by this resource policy. */
	// +optional
	RetentionPolicy *ResourcepolicyRetentionPolicy `json:"retentionPolicy,omitempty"`

	/* Immutable. Contains one of an 'hourlySchedule', 'dailySchedule', or 'weeklySchedule'. */
	Schedule ResourcepolicySchedule `json:"schedule"`

	/* Immutable. Properties with which the snapshots are created, such as labels. */
	// +optional
	SnapshotProperties *ResourcepolicySnapshotProperties `json:"snapshotProperties,omitempty"`
}

type ResourcepolicyVmStartSchedule struct {
	/* Immutable. Specifies the frequency for the operation, using the unix-cron format. */
	Schedule string `json:"schedule"`
}

type ResourcepolicyVmStopSchedule struct {
	/* Immutable. Specifies the frequency for the operation, using the unix-cron format. */
	Schedule string `json:"schedule"`
}

type ResourcepolicyWeeklySchedule struct {
	/* Immutable. May contain up to seven (one for each day of the week) snapshot times. */
	DayOfWeeks []ResourcepolicyDayOfWeeks `json:"dayOfWeeks"`
}

type ComputeResourcePolicySpec struct {
	/* Immutable. An optional description of this resource. Provide this property when you create the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. Resource policy for instances used for placement configuration. */
	// +optional
	GroupPlacementPolicy *ResourcepolicyGroupPlacementPolicy `json:"groupPlacementPolicy,omitempty"`

	/* Immutable. Resource policy for scheduling instance operations. */
	// +optional
	InstanceSchedulePolicy *ResourcepolicyInstanceSchedulePolicy `json:"instanceSchedulePolicy,omitempty"`

	/* Immutable. Region where resource policy resides. */
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. Policy for creating snapshots of persistent disks. */
	// +optional
	SnapshotSchedulePolicy *ResourcepolicySnapshotSchedulePolicy `json:"snapshotSchedulePolicy,omitempty"`
}

type ComputeResourcePolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeResourcePolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/*  */
	SelfLink string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeResourcePolicy is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeResourcePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeResourcePolicySpec   `json:"spec,omitempty"`
	Status ComputeResourcePolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeResourcePolicyList contains a list of ComputeResourcePolicy
type ComputeResourcePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeResourcePolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeResourcePolicy{}, &ComputeResourcePolicyList{})
}