package mcmaclient

import (
	"../model"
	"net/http"
)

type ServiceClient struct {
	authProvider    *AuthProvider
	httpClient      *http.Client
	service         model.Service
	tracker         model.McmaTracker
	resources       []*ResourceEndpointClient
	resourcesByType map[string]*ResourceEndpointClient
}

func (serviceClient *ServiceClient) loadResources() {
	if serviceClient.resourcesByType != nil {
		return
	}
	serviceClient.resourcesByType = make(map[string]*ResourceEndpointClient)
	for _, r := range serviceClient.service.Resources {
		resourceEndpointClient := &ResourceEndpointClient{
			authProvider:       serviceClient.authProvider,
			httpClient:         serviceClient.httpClient,
			resourceEndpoint:   r,
			serviceAuthType:    serviceClient.service.AuthType,
			serviceAuthContext: serviceClient.service.AuthContext,
			tracker:            serviceClient.tracker,
		}
		serviceClient.resources = append(serviceClient.resources, resourceEndpointClient)
		serviceClient.resourcesByType[r.ResourceType] = resourceEndpointClient
	}
}

func (serviceClient *ServiceClient) GetResourceEndpointClient(resourceType string) (*ResourceEndpointClient, bool) {
	serviceClient.loadResources()
	resourceEndpoint, found := serviceClient.resourcesByType[resourceType]
	return resourceEndpoint, found
}
