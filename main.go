package main

import (
	"flag"
	"github.com/ui-kreinhard/go-grafcli/dashboards/exporter"
	"github.com/ui-kreinhard/go-grafcli/dashboards/importer"
	datasourceExport "github.com/ui-kreinhard/go-grafcli/datasources/exporter"
	datasourceImport "github.com/ui-kreinhard/go-grafcli/datasources/importer"
	notificationChannelsImport "github.com/ui-kreinhard/go-grafcli/notificationChannels/importer"
	notificationChannelsExport "github.com/ui-kreinhard/go-grafcli/notificationChannels/exporter"
	folderImport "github.com/ui-kreinhard/go-grafcli/folders/importer"
	folderExport "github.com/ui-kreinhard/go-grafcli/folders/exporter"
	"github.com/ui-kreinhard/go-grafcli/http"
	"github.com/ui-kreinhard/go-grafcli/user"
	"log"
)

func main() {
	modeOfOperation := flag.String("mode", "", "mode of operation")
	filename := flag.String("filename", "", "Filename to read/write")
	baseUrl := flag.String("baseUrl", "", "URL of grafana")
	username := flag.String("username", "", "Username")
	password := flag.String("password", "", "Password")
	newPassword := flag.String("newPassword", "", "new password")
	dashboardUri := flag.String("dashboardUri", "", "single dashboard parameter")
	flag.Parse()

	httpClient := http.NewHttpClient(*baseUrl, *username, *password)

	var err error

	switch *modeOfOperation {
	case "import-dashboards":
		i := importer.Importer{ httpClient}
		err = i.Import(*filename)
		break
	case "import-single-dashboard":
		i := importer.Importer{ httpClient}
		err = i.ImportSingleDashboard(*filename, *dashboardUri)
		break
	case "export-dashboards":
		e := exporter.Exporter{httpClient}
		err = e.ExportDashboard(*filename)
		break
	case "import-datasources":
		log.Println("import ds", *filename)
		i := datasourceImport.Importer{httpClient}
		err = i.Import(*filename)
		break
	case "import-notificationchannels":
		log.Println("importing notification channels", *filename)
		i:= notificationChannelsImport.Importer{httpClient}
		err = i.Import(*filename)
		break
	case "export-notificationchannels":
		log.Println("exporting notification channels", *filename)
		e := notificationChannelsExport.Exporter{httpClient}
		err = e.Export(*filename)
		break
	case "export-datasources":
		log.Println("exporting ds")
		e := datasourceExport.Exporter{httpClient}
		err = e.Export(*filename)
		break
	case "import-folders":
		log.Println("importing folders")
		i := folderImport.Importer{httpClient}
		err = i.Import(*filename)
		break
	case "export-folders":
		log.Println("exporting folders")
		e := folderExport.Export{httpClient}
		err = e.Export(*filename)
	case "changePassword":
		log.Println("Changing password")
		c := user.ChangePassword{httpClient}
		err = c.Change(*password, *newPassword)
		break
	default:
		log.Println("Unknown cmd, doing nothing")
		break
	}
	if err != nil {
		log.Println(err)
	}
}
