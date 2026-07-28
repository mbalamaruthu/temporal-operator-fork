package main

import "testing"

func TestCacheOptionsForNamespacesEmpty(t *testing.T) {
	cacheOptions, namespaces := cacheOptionsForNamespaces("")

	if len(namespaces) != 0 {
		t.Fatalf("expected no watched namespaces, got %v", namespaces)
	}
	if cacheOptions.DefaultNamespaces != nil {
		t.Fatalf("expected default all-namespace cache, got %v", cacheOptions.DefaultNamespaces)
	}
}

func TestCacheOptionsForNamespacesCSV(t *testing.T) {
	cacheOptions, namespaces := cacheOptionsForNamespaces(" temporal,temporal-dev1,,temporal ")

	expectedNamespaces := []string{"temporal", "temporal-dev1"}
	if len(namespaces) != len(expectedNamespaces) {
		t.Fatalf("expected namespaces %v, got %v", expectedNamespaces, namespaces)
	}
	for i := range expectedNamespaces {
		if namespaces[i] != expectedNamespaces[i] {
			t.Fatalf("expected namespaces %v, got %v", expectedNamespaces, namespaces)
		}
	}

	if len(cacheOptions.DefaultNamespaces) != len(expectedNamespaces) {
		t.Fatalf("expected cache namespaces %v, got %v", expectedNamespaces, cacheOptions.DefaultNamespaces)
	}
	for _, namespace := range expectedNamespaces {
		if _, exists := cacheOptions.DefaultNamespaces[namespace]; !exists {
			t.Fatalf("expected cache namespace %q to be configured", namespace)
		}
	}
}
