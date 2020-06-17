package mercedes

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

var (
	apiTryoutBase     = "https://api.mercedes-benz.com/experimental/connectedvehicle_tryout/v1"
	apiProductionBase = "https://api.mercedes-benz.com/experimental/connectedvehicle/v1"
)

// Provide a default HTTP transport mechanism that wraps requests with the API key / Bearer token.
type apiKeyHttpTransport struct {
	apiKey string
}

func (a *apiKeyHttpTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("Authorization", "Bearer "+a.apiKey)
	return http.DefaultTransport.RoundTrip(r)
}

type ConnectedVehicleClient struct {
	apiBase    string
	httpClient *http.Client
}

func (m *ConnectedVehicleClient) getJson(ctx context.Context, path string, target interface{}) error {
	req, _ := http.NewRequestWithContext(ctx, "GET", m.apiBase+path, nil)
	r, err := m.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode == http.StatusUnauthorized {
		println("unauthorized")
		return err
	}
	return json.NewDecoder(r.Body).Decode(target)
}

func (m *ConnectedVehicleClient) postJson(ctx context.Context, path string, payload interface{}) error {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(payload)
	if err != nil {
		return err
	}

	req, _ := http.NewRequestWithContext(ctx, "POST", m.apiBase+path, b)
	req.Header.Add("Content-Type", "application/json")

	_, err = m.httpClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func NewClient(apiKey string, productionMode bool) *ConnectedVehicleClient {
	m := &ConnectedVehicleClient{
		httpClient: &http.Client{Transport: &apiKeyHttpTransport{apiKey: apiKey}},
	}

	if productionMode == true {
		m.apiBase = apiProductionBase
	} else {
		m.apiBase = apiTryoutBase
	}

	return m
}

func (m *ConnectedVehicleClient) NewVehicle(vehicleId string) *ConnectedVehicle {
	return &ConnectedVehicle{
		VehicleID: vehicleId,
		client:    m,
	}
}

func (m *ConnectedVehicleClient) GetVehicles(ctx context.Context) ([]Vehicle, error) {
	var vehicles []Vehicle
	ret := m.getJson(ctx, "/vehicles", &vehicles)
	return vehicles, ret
}
