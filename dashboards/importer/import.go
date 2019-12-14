package importer

import (
	"encoding/json"
	"github.com/ui-kreinhard/go-grafcli/http"
	"github.com/ui-kreinhard/go-grafcli/types"
	"io/ioutil"
)

type Importer struct{
	HttpClient *http.HttpClient
}

type IImporter interface {
	Download()
}

type SearchResultElement struct {
	Id    uint64 `json:"id"`
	Uid   string `json:"uid"`
	Title string `json:"title"`
	Url   string `json:"url"`
	Slug  string `json:"slug"`
	Type  string `json:"type"`
}

func (i *Importer) downloadSingleDashboard(searchResult types.SearchReslt) (*types.Export, error) {
	jsonMap := make(map[string]interface{})

	//output, err := utils.HttpGet("http://localhost:3000/api/dashboards/"+searchResult.Uri, "admin", "admin")
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
