package exporter

import (
	"encoding/json"
	"errors"
	"github.com/ui-kreinhard/go-grafcli/http"
	"github.com/ui-kreinhard/go-grafcli/types"
	"log"
	"io/ioutil"
	"strconv"
)

type Exporter struct {
	HttpClient *http.HttpClient
}

func (e *Exporter) getRegisteredNotificationChannels() ([]types.NotificationChannel, error){
	notificationChannels := []types.NotificationChannel{}
	output, err := e.HttpClient.HttpGet("/api/alert-notifications")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(output, &notificationChannels)
	if err != nil {
		return nil, err
	}
	return notificationChannels, nil
}

func contains(registeredNotificationChannels []types.NotificationChannel, notificationChannel types.NotificationChannel) bool {
	for _, registereDataSource := range registeredNotificationChannels {
		if registereDataSource.Uid== notificationChannel.Uid {
			return true
		}
	}
	return false
}

func (e *Exporter) createNotificationChannel(notfificationChannel types.NotificationChannel) error {
	dataToBeSent, err := json.Marshal(notfificationChannel)
	if err != nil {
		return err
	}
	_, _, err = e.HttpClient.HttpPost("/api/alert-notifications", string(dataToBeSent))
	return err
}

func (e *Exporter) updateNotificationChannel(notificationChannel types.NotificationChannel) error {
	dataToBeSent, err := json.Marshal(notificationChannel)
	if err != nil {
		return err
	}
	_, _, err = e.HttpClient.HttpPut("/api/alert-notifications/" + strconv.FormatUint(notificationChannel.Id, 10), string(dataToBeSent))
	return err
}

func (e *Exporter) decideToUpdateOrCreate(registeredNotificationChannels, notificationChannelsToExport []types.NotificationChannel) error{
	err := errors.New("")
	for _, notificationChannelToExport := range notificationChannelsToExport {
		if contains(registeredNotificationChannels, notificationChannelToExport) {
			log.Println("updating")
			err = e.updateNotificationChannel(notificationChannelToExport)
		} else {
			log.Println("creating")
			err = e.createNotificationChannel(notificationChannelToExport)
		}
	}
	return err
}

func (e *Exporter) Export(filename string) error {
	notificationChannels := []types.NotificationChannel{}
	registeredNotificationChannels, err := e.getRegisteredNotificationChannels()
	if err != nil {
		return err
	}

	contentByte, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(contentByte, &notificationChannels)
	if err != nil {
		return err
	}
	return e.decideToUpdateOrCreate(registeredNotificationChannels, notificationChannels)
}
