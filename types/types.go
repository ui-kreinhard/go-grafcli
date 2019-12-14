package types

type SearchReslt struct {
	Uri   string `json:"Uri"`
	Title string `json:"Title"`
}

type Export struct {
	Title     string
	Dashboard map[string]interface{}
}

type Mode struct {
	HttpMode string `json:"httpMode"`
}

type Datasource struct {
	Id          uint64 `json:"id"`
	OrgId       uint64 `json:"orgId"`
	Name        string `json:"name"`
	TypeLogoUrl string `json:"typeLogoUrl"`
	Access      string `json:"access"`
	Url         string `json:"url"`
	Password    string `json:"password"`
	User        string `json:"user"`
	Database    string `json:"database"`
	BasicAuth   bool   `json:"basicAuth"`
	IsDefault   bool   `json:"isDefault"`
	JsonData    Mode   `json:"jsonData"`
	ReadOnly    bool   `json:"readOnly"`
	Type        string `json:"type"`
}
