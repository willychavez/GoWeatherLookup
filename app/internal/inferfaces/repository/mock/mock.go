package mock

import "net/http"

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func (m *MockClient) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}
