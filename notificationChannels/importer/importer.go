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

func writeNotificationChannels(filename string, notificationChannels []types.NotificationChannel) error {
	json, err := json.Marshal(notificationChannels)
	if err != nil {
		return nil
	}
	return ioutil.WriteFile(filename,json, 0655)
}

func (i *Importer) Import(filename string) error {
	notificationChannels := []types.NotificationChannel{}
	output, err := i.HttpClient.HttpGet("/api/alert-notifications")
	if err != nil {
		return err
	}
	err = json.Unmarshal(output, &notificationChannels)
	if err != nil {
		return err
	}
	return writeNotificationChannels(filename, notificationChannels)
}