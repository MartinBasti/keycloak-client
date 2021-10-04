// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package common

import (
	"github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	"sync"
)

// Ensure, that KeycloakClientFactoryMock does implement KeycloakClientFactory.
// If this is not the case, regenerate this file with moq.
var _ KeycloakClientFactory = &KeycloakClientFactoryMock{}

// KeycloakClientFactoryMock is a mock implementation of KeycloakClientFactory.
//
// 	func TestSomethingThatUsesKeycloakClientFactory(t *testing.T) {
//
// 		// make and configure a mocked KeycloakClientFactory
// 		mockedKeycloakClientFactory := &KeycloakClientFactoryMock{
// 			AuthenticatedClientFunc: func(kc v1alpha1.Keycloak) (KeycloakInterface, error) {
// 				panic("mock out the AuthenticatedClient method")
// 			},
// 		}
//
// 		// use mockedKeycloakClientFactory in code that requires KeycloakClientFactory
// 		// and then make assertions.
//
// 	}
type KeycloakClientFactoryMock struct {
	// AuthenticatedClientFunc mocks the AuthenticatedClient method.
	AuthenticatedClientFunc func(kc v1alpha1.Keycloak) (KeycloakInterface, error)

	// calls tracks calls to the methods.
	calls struct {
		// AuthenticatedClient holds details about calls to the AuthenticatedClient method.
		AuthenticatedClient []struct {
			// Kc is the kc argument value.
			Kc v1alpha1.Keycloak
		}
	}
	lockAuthenticatedClient sync.RWMutex
}

// AuthenticatedClient calls AuthenticatedClientFunc.
func (mock *KeycloakClientFactoryMock) AuthenticatedClient(kc v1alpha1.Keycloak) (KeycloakInterface, error) {
	if mock.AuthenticatedClientFunc == nil {
		panic("KeycloakClientFactoryMock.AuthenticatedClientFunc: method is nil but KeycloakClientFactory.AuthenticatedClient was just called")
	}
	callInfo := struct {
		Kc v1alpha1.Keycloak
	}{
		Kc: kc,
	}
	mock.lockAuthenticatedClient.Lock()
	mock.calls.AuthenticatedClient = append(mock.calls.AuthenticatedClient, callInfo)
	mock.lockAuthenticatedClient.Unlock()
	return mock.AuthenticatedClientFunc(kc)
}

// AuthenticatedClientCalls gets all the calls that were made to AuthenticatedClient.
// Check the length with:
//     len(mockedKeycloakClientFactory.AuthenticatedClientCalls())
func (mock *KeycloakClientFactoryMock) AuthenticatedClientCalls() []struct {
	Kc v1alpha1.Keycloak
} {
	var calls []struct {
		Kc v1alpha1.Keycloak
	}
	mock.lockAuthenticatedClient.RLock()
	calls = mock.calls.AuthenticatedClient
	mock.lockAuthenticatedClient.RUnlock()
	return calls
}
