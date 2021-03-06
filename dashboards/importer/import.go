package importer

import (
	"encoding/json"
	"github.com/ui-kreinhard/go-grafcli/http"
	"github.com/ui-kreinhard/go-grafcli/types"
	"io/ioutil"
	"strings"
)

type Importer struct{
	HttpClient *http.HttpClient
}

type IImporter interface {
	Download()
}

func (i *Importer) downloadSingleDashboard(searchResult types.SearchReslt) (*types.Export, error) {
	jsonMap := make(map[string]interface{})
	output, err := i.HttpClient.HttpGet("/api/dashboards/" + searchResult.Uri)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(output, &jsonMap)
	if err != nil {
		return nil, err
	}
	newMap := make(map[string]interface{})
	dashb := jsonMap["dashboard"].(map[string]interface{})
	// dashb["id"] = nil
	delete(dashb, "version")
	delete(dashb, "schemaVersion")
	newMap["dashboard"] = dashb
	newMap["overwrite"] = false
	newMap["folderId"] = nil
	return &types.Export{searchResult.Title, dashb}, nil
}

func writeDashboards(filename string, dashboards []types.Export) error {
	json, err := json.Marshal(dashboards)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, json, 0655)
}

func (i *Importer) Import(filename string) ( error) {
	var dashboards []types.Export
	var searchResults []types.SearchReslt

	output, err := i.HttpClient.HttpGet("/api/search?type=dash-db")
	err = json.Unmarshal(output, &searchResults)
	if err != nil {
		return err
	}
	for _, searchResult := range searchResults {
		result, err := i.downloadSingleDashboard(searchResult)
		if err != nil {
			return err
		}
		dashboards = append(dashboards, *result)
	}
	return writeDashboards(filename, dashboards)
}

func correctUri(uri string) string {
	if strings.HasPrefix(uri, "db/") {
		return uri
	}
	return "db/" + uri
}

func (i *Importer) ImportSingleDashboard(filename, uri string) error {
	toDownload := types.SearchReslt{
		Uri:   correctUri(uri),
		Title: "",
	}
	export, err := i.downloadSingleDashboard(toDownload)
	if err != nil {
		return err
	}
	exports := []types.Export{*export}
	return writeDashboards(filename, exports)
}