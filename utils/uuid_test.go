package utils

import (
    "testing"
    "github.com/google/uuid"
)

func TestGenerateUUID(t *testing.T) {
    g := NewGenerateUUID()
    id, err := g.GenerateUUID()
    if err != nil {
        t.Errorf("Failed to generate UUID")
    }
    _, err = uuid.Parse(id)
    if err != nil {
        t.Errorf("Invalid UUID")
    }
}
