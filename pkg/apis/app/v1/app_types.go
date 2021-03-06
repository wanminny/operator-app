package v1

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AppSpec defines the desired state of App
type AppSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	Replicas  *int32                      `json:"replicas"`
	Image     string                      `json:"image"`
	Resources corev1.ResourceRequirements `json:"resouces,omitempty"`
	Envs      []corev1.EnvVar             `json:"envs,omitempty"`
	Ports     []corev1.ServicePort        `json:"ports,omitempty"`
}

// AppStatus defines the observed state of App
type AppStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	appsv1.DeploymentStatus `json:",inline"`
}

// type AppConditions struct{
// 	Type string
// 	Message string
// 	Reason string
// 	Ready bool
// 	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty" protobuf:"bytes,6,opt,name=lastUpdateTime"`
// 	// Last time the condition transitioned from one status to another.
// 	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,7,opt,name=lastTransitionTime"`
// }

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// App is the Schema for the apps API
// +k8s:openapi-gen=true
type App struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppSpec   `json:"spec,omitempty"`
	Status AppStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AppList contains a list of App
type AppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []App `json:"items"`
}

func init() {
	SchemeBuilder.Register(&App{}, &AppList{})
}
