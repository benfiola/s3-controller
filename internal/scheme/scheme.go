package scheme

import (
	apiv1 "github.com/benfiola/s3-controller/pkg/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func Build() (*runtime.Scheme, error) {
	scheme := runtime.NewScheme()

	err := apiv1.AddToScheme(scheme)
	if err != nil {
		return nil, err
	}

	return scheme, nil
}
