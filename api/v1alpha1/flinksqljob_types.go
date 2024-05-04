/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FlinkSqlJobSpec defines the desired state of FlinkSqlJob
type FlinkSqlJobSpec struct {
	// Configurations will be passed to the Flink SQL Gateway to configure the session for this jobs.
	//+kubebuilder:validation:Optional
	Configurations []string `json:"configurations,omitempty"`

	// SQL statement of the job.
	//+kubebuilder:validation:Required
	Statement string `json:"statement"`
}

// FlinkSqlJobStatus defines the observed state of FlinkSqlJob
type FlinkSqlJobStatus struct {
	//+kubebuilder:validation:Optional
	Conditions []StatusCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
	// Represents time when the job controller started processing a job.
	//+kubebuilder:validation:Optional
	StartTime *metav1.Time `json:"startTime,omitempty"`
	// Represents time when the job was completed.
	//+kubebuilder:validation:Optional
	CompletionTime *metav1.Time `json:"completionTime,omitempty"`
	// Represents the global status of the job.
	GlobalStatus JobConditionType `json:"globalStatus"`
	// Session handle of the job.
	//+kubebuilder:validation:Optional
	SessionHandle string `json:"sessionHandle,omitempty"`
	// Operations of the job and their statuses.
	//+kubebuilder:validation:Optional
	Operations []OperationStatus `json:"operations,omitempty"`
}

// OperationStatus represents the status of an operation.
type OperationStatus struct {
	// Represents the operation ID.
	//+kubebuilder:validation:Required
	OperationID int `json:"operationID"`
	// The operation handle.
	//+kubebuilder:validation:Required
	OperationHandle string `json:"operationHandle"`
	//+kubebuilder:validation:Optional
	Conditions []StatusCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// JobConditionType is a type of job condition.
type JobConditionType string

// These are built-in conditions of a job.
const (
	// JobComplete means the job has completed its execution.
	JobComplete JobConditionType = "Complete"
	// JobRunning means the job is still running.
	JobRunning JobConditionType = "Running"
	// JobFailed means the job has failed its execution.
	JobFailed JobConditionType = "Failed"
)

// StatusCondition describes the state a job or operation went through at a certain point.
type StatusCondition struct {
	// Type of job condition, Complete, Running or Failed.
	Type JobConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`
	// Last time the condition was checked.
	//+kubebuilder:validation:Optional
	LastProbeTime metav1.Time `json:"lastProbeTime,omitempty"`
	// Last time the condition transit from one status to another.
	//+kubebuilder:validation:Optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	//+kubebuilder:validation:Optional
	Reason string `json:"reason,omitempty"`
	// Human-readable message indicating details about last transition.
	//+kubebuilder:validation:Optional
	Message string `json:"message,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// FlinkSqlJob is the Schema for the flinksqljobs API
// +kubebuilder:printcolumn:name="SessionHandle",type=string,JSONPath=`.status.sessionHandle`
// +kubebuilder:printcolumn:name="StartTime",type=date,JSONPath=`.status.startTime`
// +kubebuilder:printcolumn:name="CompletionTime",type=date,JSONPath=`.status.completionTime`
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.globalStatus`
type FlinkSqlJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FlinkSqlJobSpec   `json:"spec,omitempty"`
	Status FlinkSqlJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// FlinkSqlJobList contains a list of FlinkSqlJob
type FlinkSqlJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlinkSqlJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FlinkSqlJob{}, &FlinkSqlJobList{})
}
