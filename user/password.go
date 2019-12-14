package user

import (
	"encoding/json"
	"github.com/ui-kreinhard/go-grafcli/http"
	"github.com/ui-kreinhard/go-grafcli/types"
)

type ChangePassword struct {
	HttpClient *http.HttpClient
}

func (c *ChangePassword) Change(oldPassword, newPassword string) error {
	newPasswordChange := types.NewPasswordChange{newPassword, newPassword, oldPassword}
	json, err := json.Marshal(newPasswordChange)
	if err != nil {
		return err
	}
	_, _, err = c.HttpClient.HttpPut("/api/user/password", string(json))
	return err
}

func changePassword() {
	/*

	http://localhost:3000/api/user/password
	payloda:
	{"newPassword":"newPw","confirmNew":"newPw","oldPassword":"admin"}
	 */
}
