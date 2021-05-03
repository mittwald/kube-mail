package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EmailPolicySMTPSink struct {
	Server      v1.ObjectReference `json:"server"`
	Credentials v1.ObjectReference `json:"credentials"`
}

type EmailPolicyRateLimiting struct {
	Maximum int    `json:"maximum"`
	Period  string `json:"period"`
}

type EmailPolicySink struct {
	SMTP EmailPolicySMTPSink `json:"smtp"`
}

type EmailPolicySpec struct {
	Default      bool                    `json:"default"`
	PodSelector  metav1.LabelSelector    `json:"podSelector"`
	RateLimiting EmailPolicyRateLimiting `json:"ratelimiting"`
	Sink         EmailPolicySink         `json:"sink"`
}

type EmailPolicyStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=emailpolicies,scope=Namespaced,shortName=emailpolicy

// EmailPolicy is the Schema for the emailpolicy API
type EmailPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EmailPolicySpec   `json:"spec,omitempty"`
	Status EmailPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmailPolicyList contains a list of Project
type EmailPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmailPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EmailPolicy{}, &EmailPolicyList{})
}
