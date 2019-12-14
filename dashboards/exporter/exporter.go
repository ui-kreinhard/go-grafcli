package exporter

import (
	"encoding/json"
	"github.com/ui-kreinhard/go-grafcli/http"
	"github.com/ui-kreinhard/go-grafcli/types"
	"io/ioutil"
)

type Exporter struct {
	HttpClient *http.HttpClient
}

type SearchResultElement struct {
	Id    uint64 `json:"id"`
	Uid   string `json:"uid"`
	Title string `json:"title"`
	Url   string `json:"url"`
	Slug  string `json:"slug"`
	Type  string `json:"type"`
}

func (e *Exporter) getRegisteredDashboards() ([]SearchResultElement, error) {
	searchResults := []SearchResultElement{}
	output, err := e.HttpClient.HttpGet("/api/search")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(output, &searchResults)
	if err != nil {
		return nil, err
	}

	return searchResults, nil
}

func (e *Exporter) dashboardExists(id uint64) (bool, error) {
	registeredDashboards, err := e.getRegisteredDashboards()
	if err != nil {
		return false, err
	}
	for _, toCompare := range registeredDashboards {
		if toCompare.Id == id {
			return true, nil
		}
	}
	return false, nil
}

func (e *Exporter) exportDashboard(exports []types.Export) error {

	for _, export := range exports {
		id, _ := export.Dashboard["id"].(float64)
		dataToBeSent := make(map[string]interface{})
		dataToBeSent["dashboard"] = export.Dashboard
		exists, err := e.dashboardExists(uint64(id))
		if err != nil {
			return err
		}
		dataToBeSent["overwrite"] = exists

		if !exists {
			export.Dashboard["id"] = nil
		}
		data, err := json.Marshal(dataToBeSent)
		if err != nil {
			return err
		}
		_, _, err = e.HttpClient.HttpPost("/api/dashboards/db", string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Exporter) ExportDashboard(filename string) error {
	dashboards := []types.Export{}
	rawJson, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawJson, &dashboards)
	if err != nil {
		return err
	}
	err = e.exportDashboard(dashboards)
	if err != nil {
		return err
	}

	return nil
}
