package service

import (
	repository2 "aivo-code-challenge/repository"
	"testing"
)

func TestNewIntegrationService_success(t *testing.T) {
	repository, err := repository2.NewIntegrationRepository()
	assertEqual(t, err, nil)
	_, err = NewIntegrationService(repository)
	assertEqual(t, err, nil)
}

func TestNewIntegrationService_fails(t *testing.T) {
	_, err := NewIntegrationService(nil)
	assertEqual(t, err.Error(), "repository is required but was empty")
}

func assertEqual(t *testing.T, got interface{}, want interface{}) {
	if got != want {
		t.Fatalf("%s != %s", got, want)
	}
}