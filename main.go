package main

import (
	"flag"
	"github.com/ui-kreinhard/go-grafcli/dashboards/exporter"
	"github.com/ui-kreinhard/go-grafcli/dashboards/importer"
	datasourceExport "github.com/ui-kreinhard/go-grafcli/datasources/exporter"
	datasourceImport "github.com/ui-kreinhard/go-grafcli/datasources/importer"
	"github.com/ui-kreinhard/go-grafcli/http"
	"github.com/ui-kreinhard/go-grafcli/user"
	"log"
)

func main() {
	modeOfOperation := flag.String("mode", "import", "mode of operation")
	filename := flag.String("filename", "", "Filename to read/write")
	baseUrl := flag.String("baseUrl", "", "URL of grafana")
	username := flag.String("username", "", "Username")
	password := flag.String("password", "", "Password")
	newPassword := flag.String("newPassword", "", "new password")
	flag.Parse()

	httpClient := http.NewHttpClient(*baseUrl, *username, *password)

	var err error

	switch *modeOfOperation {
	case "import-dashboards":
		i := importer.Importer{ httpClient}
		err = i.Import(*filename)
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
	case "export-datasources":
		log.Println("exporting ds")
		e := datasourceExport.Exporter{httpClient}
		err = e.Export(*filename)
		break
	case "changePassword":
		log.Println("Changing password")
		c := user.ChangePassword{httpClient}
		err = c.Change(*password, *newPassword)
		break;
	default:
		log.Println("Unknown cmd, doing nothing")
		break
	}
	if err != nil {
		log.Println(err)
	}
}
