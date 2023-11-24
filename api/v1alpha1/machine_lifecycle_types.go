// Copyright 2023 T-Systems International GmbH, SAP SE or an SAP affiliate company. All right reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:generate=true

// MachineLifecycleSpec contains desired configuration of machine lifecycle.
type MachineLifecycleSpec struct {
}

// MachineLifecycleStatus contains observed state of MachineLifecycle object.
type MachineLifecycleStatus struct {
}

// +kubebuilder:object:root=true

// MachineLifecycle is the schema for MachineLifecycle API object.
type MachineLifecycle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MachineLifecycleSpec   `json:"spec,omitempty"`
	Status MachineLifecycleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MachineLifecycleList contains a list of MachineLifecycle objects.
type MachineLifecycleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MachineLifecycle `json:"items"`
}
