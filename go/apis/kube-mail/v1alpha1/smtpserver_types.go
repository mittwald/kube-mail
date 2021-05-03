package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type SMTPServerSpec struct {
	Server   string `json:"server"`
	Port     int    `json:"port"`
	TLS      bool   `json:"tls"`
	AuthType string `json:"authType"`
}

type SMTPServerStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=smtpservers,scope=Namespaced,shortName=smtpserver

// SMTPServer is the Schema for the smtpserver API
type SMTPServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SMTPServerSpec   `json:"spec,omitempty"`
	Status SMTPServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SMTPServerList contains a list of Project
type SMTPServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SMTPServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SMTPServer{}, &SMTPServerList{})
}
