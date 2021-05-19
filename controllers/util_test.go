package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestSetObjectMeta(t *testing.T) {
	owner := appv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-deployment",
			Namespace: "fake-namespace",
		},
	}
	pod := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "fake-pod",
		},
	}

	err := SetObjectMeta(&owner, &pod, owner.GroupVersionKind())
	assert.Nil(t, err)
	assert.Equal(t, "fake-namespace", pod.Namespace)
	assert.Equal(t, owner.GroupVersionKind().Kind, pod.OwnerReferences[0].Kind)
	assert.NotEmpty(t, pod.Annotations[AnnotationResourceSpecHash])
}
