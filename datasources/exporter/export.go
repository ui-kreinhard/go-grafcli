package exporter

import (
	//"encoding/json"
	//"io/ioutil"
	"encoding/json"
	"errors"
	"github.com/ui-kreinhard/go-grafcli/http"
	"github.com/ui-kreinhard/go-grafcli/types"
	"io/ioutil"
	"log"
	"strconv"
)

type Exporter struct{
	HttpClient *http.HttpClient
}

func (e *Exporter) getRegisteredDatasources() ([]types.Datasource, error) {
	registeredDatasources := []types.Datasource{}

	output, err := e.HttpClient.HttpGet("/api/datasources")
	if err != nil {
		return nil, err
	}
	log.Println(string(output))
	err = json.Unmarshal(output, &registeredDatasources)
	if err != nil {
		return nil, err
	}
	return registeredDatasources, nil
}

func (e *Exporter) export(datasources []types.Datasource) error {
	for _, datasource := range datasources {
		dataToBeSent, err := json.Marshal(datasource)
		if err != nil {
			return err
		}

		_, _, err = e.HttpClient.HttpPost("/api/datasources/", string(dataToBeSent))
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Exporter) createDatasource(datasource types.Datasource) error {
	dataToBeSent, err := json.Marshal(datasource)
	if err != nil {
		return err
	}

	_, _, err = e.HttpClient.HttpPost("http://localhost:3000/api/datasources/", string(dataToBeSent))
	return err
}

func (e *Exporter) updateDatasource(datasource types.Datasource) error {
	dataToBeSent, err := json.Marshal(datasource)
	if err != nil {
		return err
	}

	_, _, err = e.HttpClient.HttpPut("http://localhost:3000/api/datasources/"+strconv.FormatUint(datasource.Id, 10), string(dataToBeSent))
	return err
}

func contains(registereDataSources []types.Datasource, datasource types.Datasource) bool {
	for _, registereDataSource := range registereDataSources {
		if registereDataSource.Id == datasource.Id {
			return true
		}
	}
	return false
}

func (e *Exporter) decideToUpdateOrCreate(registeredDataSources, dataSourcesToExport []types.Datasource) error {
	for _, dataSourceToExport := range dataSourcesToExport {
		err := errors.New("")
		if contains(registeredDataSources, dataSourceToExport) {
			log.Println("updating")
			err = e.updateDatasource(dataSourceToExport)

		} else {
			log.Println("creating")
			err = e.createDatasource(dataSourceToExport)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Exporter) Export(filename string) error {
	datasources := []types.Datasource{}
	registeredDataSources, err := e.getRegisteredDatasources()
	if err != nil {
		return err
	}
	contentByte, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(contentByte, &datasources)
	if err != nil {
		return err
	}
	return e.decideToUpdateOrCreate(registeredDataSources, datasources)
}
