package importer

import (
	"encoding/json"
	"github.com/ui-kreinhard/go-grafcli/http"
	"github.com/ui-kreinhard/go-grafcli/types"
	"io/ioutil"
)

type Importer struct {
	HttpClient *http.HttpClient
}

func writeFolders(filename string, folders []types.Folder) error {
	jsonData, err := json.Marshal(folders)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, jsonData, 0655)
}

func (i *Importer) Import(filename string) error {
	var folders []types.Folder
	output, err := i.HttpClient.HttpGet("/api/search?type=dash-folder")
	if err != nil {
		return err
	}
	err = json.Unmarshal(output, &folders)
	if err != nil {
		return err
	}
	return writeFolders(filename, folders)
}
