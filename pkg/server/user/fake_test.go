package user

import (
	"testing"

	"github.com/luizalabs/teresa-api/pkg/common"
)

func TestFakeOperationsLogin(t *testing.T) {
	fake := NewFakeOperations()

	expectedEmail := "teresa@luizalabs.com"
	expectedPassword := "123456"
	fake.(*FakeOperations).Storage[expectedEmail] = expectedPassword

	token, err := fake.Login(expectedEmail, expectedPassword)
	if err != nil {
		t.Fatal("Error on perform Login in FakeUserOperations: ", err)
	}
	expectedToken := "good token"
	if token != expectedToken {
		t.Errorf("expected %s, got %s", expectedToken, token)
	}
}

func TestFakeOperationsBadLogin(t *testing.T) {
	fake := NewFakeOperations()

	if _, err := fake.Login("invalid@luizalabs.com", "foo"); err != common.ErrPermissionDenied {
		t.Errorf("expected ErrPermissionDenied, got %s", err)
	}
}
