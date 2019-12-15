package exporter

import (
	"encoding/json"
	"errors"
	"github.com/ui-kreinhard/go-grafcli/http"
	"github.com/ui-kreinhard/go-grafcli/types"
	"io/ioutil"
	"log"
)

type Export struct {
	HttpClient *http.HttpClient
}

func (e *Export) getRegisteredFolders() ([]types.Folder, error) {
	var folders []types.Folder
	output, err := e.HttpClient.HttpGet("/api/search?type=dash-folder")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(output, &folders)
	if err != nil {
		return nil, err
	}
	return folders, nil
}

func contains(registeredFolders []types.Folder, folderToExport types.Folder) bool {
	for _, registeredFolder := range registeredFolders {
		if registeredFolder.Uid == folderToExport.Uid {
			return true
		}
	}
	return false
}

func (e *Export) getFolder(folder types.Folder) (*types.Folder, error) {
	var folderDetails types.Folder
	output, err := e.HttpClient.HttpGet("/api/folders/" + folder.Uid)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(output, &folderDetails)
	if err != nil {
		return nil, err
	}
	folder.Version = folderDetails.Version 

	return &folder, nil
}

func (e *Export) updateFolder(folderToExport types.Folder) error {
	folderWithVersion, err := e.getFolder(folderToExport)
	log.Println(folderWithVersion)
	if err != nil {
		return err
	}
	dataToBeSent, err := json.Marshal(folderWithVersion)
	if err != nil {
		return err
	}
	_,_, err = e.HttpClient.HttpPut("/api/folders/" + folderWithVersion.Uid, string(dataToBeSent))
	return err
}

func (e *Export) createFolder(folderToExport types.Folder) error {
	dataToBeSent, err := json.Marshal(folderToExport)
	if err != nil {
		return err
	}
	_,_, err = e.HttpClient.HttpPost("/api/folders/", string(dataToBeSent))
	return err
}

func (e *Export) decideToUpdateOrCreate(registeredFolders, foldersToExport []types.Folder) error {
	err := errors.New("")
	for _, folderToExport := range foldersToExport {
		if contains(registeredFolders, folderToExport) {
			log.Print("Updating")
			err = e.updateFolder(folderToExport)
		} else {
			log.Println("creating")
			err = e.createFolder(folderToExport)
		}
	}
	return err
}

func (e *Export) Export(filename string) error {
	var folders []types.Folder
	registeredFolders, err := e.getRegisteredFolders()
	if err != nil {
		return err
	}
	contentByte, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(contentByte, &folders)
	if err != nil {
		return err
	}
	return e.decideToUpdateOrCreate(registeredFolders, folders)
}
