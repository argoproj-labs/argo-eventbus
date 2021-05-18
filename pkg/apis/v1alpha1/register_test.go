package v1alpha1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/argoproj-labs/argo-eventbus/pkg/apis"
)

func TestResource(t *testing.T) {
	expect := schema.GroupResource{
		Group:    apis.Group,
		Resource: "hello",
	}

	got := Resource("hello")

	if diff := cmp.Diff(expect, got); diff != "" {
		t.Errorf("unexpected resource (-expects, +got) = %v", diff)
	}
}
