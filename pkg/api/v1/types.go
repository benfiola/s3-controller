package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +groupName=s3-controller.benfiola.com
// +versionName=v1

var (
	SchemeGroupVersion = schema.GroupVersion{Group: "s3-controller.benfiola.com", Version: "v1"}
	SchemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
	localSchemeBuilder.Register(addKnownTypes)
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion, &Bucket{}, &BucketList{})

	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)

	return nil
}
