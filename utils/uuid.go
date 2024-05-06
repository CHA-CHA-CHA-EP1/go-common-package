package utils

import (
	"errors"
	"github.com/google/uuid"
)

type IGenerateUUID interface {
    GenerateUUID() (string, error)
}

type GenerateUUID struct {}

func NewGenerateUUID() IGenerateUUID {
    return &GenerateUUID{}
}

func (g *GenerateUUID) GenerateUUID() (string, error) {
    id, err := uuid.NewRandom()
    if err != nil {
        return "", errors.New("Failed to generate UUID")
    }
    return id.String(), nil
}

