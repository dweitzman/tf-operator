package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"

	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

// Install registers the API group and adds type to a schema
func Install(scheme *runtime.Scheme) {
	utilruntime.Must(AddToScheme(scheme))
	utilruntime.Must(scheme.SetVersionPriority(SchemeGroupVersion))
}
