package types

type IImporter interface {
	Run(exportFileName string)
}

type IExporter interface {
	Export(importFileName string)
}