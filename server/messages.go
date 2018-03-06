package server

import (
	"encoding/json"
	_ "log"

	"github.com/dutchcoders/marija/server/datasources"
)

type Datasource struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type InitialStateMessage struct {
	Datasources []Datasource
	Version     string
	CommitID    string
}

func (em *InitialStateMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type        string       `json:"type"`
		Datasources []Datasource `json:"datasources"`
		Version     string       `json:"version"`
		CommitID    string       `json:"commit-id"`
	}{
		Type:        InitialStateReceive,
		Datasources: em.Datasources,
		Version:     em.Version,
		CommitID:    em.CommitID,
	})
}

type Request struct {
	RequestID string `json:"request-id"`
	Type      string `json:"type"`
}

type Response struct {
	RequestID string `json:"request-id"`
}

type ItemsRequest struct {
	Request

	Items []string `json:"items"`
}

type ItemsResponse struct {
	RequestID string

	Items []datasources.Item `json:"items"`
}

func (em *ItemsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type      string             `json:"type"`
		RequestID string             `json:"request-id"`
		Items     []datasources.Item `json:"items"`
	}{
		Type:      ActionTypeItemsReceive,
		RequestID: em.RequestID,
		Items:     em.Items,
	})
}

type SearchRequest struct {
	Request

	Datasources []string `json:"datasources"`
	Fields      []string `json:"fields"`
	Query       string   `json:"query"`
}

type RequestCanceled struct {
	RequestID string
}

func (em *RequestCanceled) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type      string `json:"type"`
		RequestID string `json:"request-id"`
	}{
		Type:      ActionTypeRequestCanceled,
		RequestID: em.RequestID,
	})
}

type RequestCompleted struct {
	RequestID string
}

func (em *RequestCompleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type      string `json:"type"`
		RequestID string `json:"request-id"`
	}{
		Type:      ActionTypeRequestCompleted,
		RequestID: em.RequestID,
	})
}

type LiveResponse struct {
	Datasource string
	Graphs     []datasources.Graph
}

func (em *LiveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type       string              `json:"type"`
		Datasource string              `json:"datasource"`
		Graphs     []datasources.Graph `json:"graphs"`
	}{
		Type:       ActionTypeLiveReceive,
		Datasource: em.Datasource,
		Graphs:     em.Graphs,
	})
}

type SearchResponse struct {
	RequestID string

	Server string
	Query  string
	Nodes  []datasources.Node
}

func (em *SearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type      string             `json:"type"`
		RequestID string             `json:"request-id"`
		Server    string             `json:"server,omitempty"`
		Query     string             `json:"query"`
		Nodes     []datasources.Node `json:"results"`
	}{
		Type:      ActionTypeSearchReceive,
		RequestID: em.RequestID,
		Server:    em.Server,
		Query:     em.Query,
		Nodes:     em.Nodes,
	})
}

type GetFieldsRequest struct {
	Request

	Datasources []string `json:"datasources"`
}

type GetFieldsResponse struct {
	RequestID string

	Server string
	Index  string
	Fields interface{}
}

func (em *GetFieldsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type      string      `json:"type"`
		RequestID string      `json:"request-id"`
		Server    string      `json:"server,omitempty"`
		Index     string      `json:"index,omitempty"`
		Fields    interface{} `json:"fields"`
	}{
		Type:      ActionTypeGetFieldsReceive,
		RequestID: em.RequestID,
		Server:    em.Server,
		Index:     em.Index,
		Fields:    em.Fields,
	})
}

type ErrorMessage struct {
	RequestID string
	Message   string
}

func (em *ErrorMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type      string `json:"type"`
		RequestID string `json:"request-id"`
		Message   string `json:"message"`
	}{
		Type:      ERROR,
		RequestID: em.RequestID,
		Message:   em.Message,
	})
}
