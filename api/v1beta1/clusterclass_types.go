package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=clusterclasses,shortName=cc,scope=Namespaced,categories=cluster-api

type ClusterClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterClassSpec `json:"spec,omitempty"`
}

type ClusterClassSpec struct {
	OpenAPIV3Schema JSONSchemaProps `json:"openAPIV3Schema"`
}

type JSONSchemaProps struct {
	Items      *JSONSchemaProps           `json:"items,omitempty"`
	Properties map[string]JSONSchemaProps `json:"properties,omitempty"`
	AllOf      []JSONSchemaProps          `json:"allOf,omitempty"`

	ExclusiveMaximum bool `json:"exclusiveMaximum,omitempty"`
}

// +kubebuilder:object:root=true

type ClusterClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterClass `json:"items"`
}

func init() {
	objectTypes = append(objectTypes, &ClusterClass{}, &ClusterClassList{})
}
