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

type NewPasswordChange struct {
	NewPassword string `json:"newPassword"`
	ConfirmNew  string `json:"confirmNew"`
	OldPassword string `json:"oldPassword"`
}

type NotificationChannelSettings struct {
	Addresses   string `json:"addresses"`
	AutoResolve bool   `json:"autoResolve"`
	HttpMethod  string `json:"httpMethod"`
	UploadImage bool   `json:"uploadImage"`
}
type NotificationChannel struct {
	SendReminder          bool                        `json:"sendReminder"`
	Frequency             string                      `json:"frequency"`
	Created               string                      `json:"created"`
	Settings              NotificationChannelSettings `json:"settings"`
	Name                  string                      `json:"name"`
	Type                  string                      `json:"type"`
	IsDefault             bool                        `json:"isDefault"`
	Updated               string                      `json:"updated"`
	Id                    int                         `json:"id"`
	Uid                   string                      `json:"uid"`
	DisableResolveMessage bool                        `json:"disableResolveMessage"`
}
