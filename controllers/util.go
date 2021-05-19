package controllers

import (
	"encoding/json"
	"fmt"
	"hash/fnv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	// AnnotationResourceSpecHash is the annotation of a K8s resource spec hash
	AnnotationResourceSpecHash = "resource-spec-hash"
)

// SetObjectMeta sets ObjectMeta of child resource
func SetObjectMeta(owner, obj metav1.Object, gvk schema.GroupVersionKind) error {
	references := obj.GetOwnerReferences()
	references = append(references,
		*metav1.NewControllerRef(owner, gvk),
	)
	obj.SetOwnerReferences(references)

	if obj.GetName() == "" && obj.GetGenerateName() == "" {
		obj.SetName(owner.GetName())
	}
	if obj.GetNamespace() == "" {
		obj.SetNamespace(owner.GetNamespace())
	}

	hash, err := GetObjectHash(obj)
	if err != nil {
		return err
	}
	annotations := obj.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations[AnnotationResourceSpecHash] = hash
	obj.SetAnnotations(annotations)

	return nil
}

// Hasher hashes a string
func Hasher(value string) string {
	h := fnv.New32a()
	_, _ = h.Write([]byte(value))
	return fmt.Sprintf("%v", h.Sum32())
}

// GetObjectHash returns hash of a given object
func GetObjectHash(obj metav1.Object) (string, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return "", fmt.Errorf("failed to marshal resource")
	}
	return Hasher(string(b)), nil
}
