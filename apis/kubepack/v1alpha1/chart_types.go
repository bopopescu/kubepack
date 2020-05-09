/*
Copyright The Kubepack Authors.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindChart = "Chart"
	ResourceChart     = "chart"
	ResourceCharts    = "charts"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:path=charts,singular=chart,shortName=app,categories={kubepack,appscode}
// +kubebuilder:subresource:status

// Chart is the Schema for the charts API
type Chart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   ChartSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status ChartStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// ChartSpec defines the specification for an Chart.
// Matches ChartMetadata: https://github.com/helm/helm/blob/v3.2.1/pkg/chart/metadata.go#L28-L65
type ChartSpec struct {
	// Description regroups information and metadata about an chart.
	Description Descriptor `json:"description,omitempty" protobuf:"bytes,1,opt,name=description"`

	// Info contains human readable key,value pairs for the Chart.
	Info []InfoItem `json:"info,omitempty" protobuf:"bytes,2,rep,name=info"`

	Steps []Step `json:"steps" protobuf:"bytes,3,rep,name=steps"`

	Resources *ResourceDefinitions `json:"resources,omitempty" protobuf:"bytes,4,opt,name=resources"`
}

type Step struct {
	/*
		Templates can be listed by name or via glob. Templates will be automatically searched inside the templates directory
		Examples:
		 templates:
		   - service.yaml
		   - step/*
	*/
	Templates []string `json:"templates" protobuf:"bytes,1,rep,name=templates"`
	// Resources     *ResourceDefinitions `json:"resources,omitempty"`

	WaitFors []WaitFlags `json:"waitFors,omitempty" protobuf:"bytes,2,rep,name=waitFors"`

	ExportValues []ResourceValueMap `json:"exportValues" protobuf:"bytes,3,rep,name=exportValues"`
}

type ResourceValueMap struct {
	metav1.TypeMeta `json:",inline"`
	Name            string     `json:"name" protobuf:"bytes,1,opt,name=name"`
	Namespace       *string    `json:"namespace" protobuf:"bytes,2,opt,name=namespace"` // If not provided, use release namespace
	Values          []ValueDef `json:"value" protobuf:"bytes,3,rep,name=value"`
}

type ValueDef struct {
	Key       string `json:"key" protobuf:"bytes,1,opt,name=key"`
	ValuePath string `json:"valuePath" protobuf:"bytes,2,opt,name=valuePath"` // json path for the Value
}

// ChartStatus defines controller's the observed state of Chart
type ChartStatus struct {
	// ObservedGeneration is the most recent generation observed. It corresponds to the
	// Object's generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ChartList contains a list of Chart
type ChartList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Chart `json:"items" protobuf:"bytes,2,rep,name=items"`
}
