package utils

import (
	"testing"
)

func TestPath(t *testing.T) {

	id := ".projects-e544b82a-f5ca-4dee-9722-09f80e50d4e1.collections-b9416363-eaf6-4e41-9ece-9ffc0026b289"

	result := FirestorePath(id)
	expected := "projects/.projects-e544b82a-f5ca-4dee-9722-09f80e50d4e1/collections/.projects-e544b82a-f5ca-4dee-9722-09f80e50d4e1.collections-b9416363-eaf6-4e41-9ece-9ffc0026b289"
	if result != expected {
		t.Fatal()
	}
}
