package importer

import (
	"encoding/json"
	"github.com/ui-kreinhard/go-grafcli/http"
	"github.com/ui-kreinhard/go-grafcli/types"
	"io/ioutil"
	"log"
)

type Importer struct{
	HttpCLient *http.HttpClient
}

func writeDaSources(filename string, datasources []types.Datasource) (error) {
	json, err := json.Marshal(datasources)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, json, 0655)
}

func (i *Importer) Import(filename string) (error) {
	dataSources := []types.Datasource{}

	output, err := i.HttpCLient.HttpGet("/api/datasources")
	if err != nil {
		log.Println(err)
		return err
	}
	err = json.Unmarshal(output, &dataSources)
	if err != nil {
		return err
	}
	return writeDaSources(filename, dataSources)
}
