package models

import "database/sql"

type UserModel struct {
	ID         int    `json:"ID"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	WalletUUID string `json:"walletUUID"`
	AuthToken  string `json:"AuthToken"`
	CreatedAt  int    `json:"created-at"`
}

func (c UserModel) ParseModel(resp *sql.Rows) UserModel {
	var mdl UserModel
	err := resp.Scan(&mdl.ID, &mdl.Username, &mdl.Password, &mdl.WalletUUID, &mdl.AuthToken, &mdl.CreatedAt)
	if err != nil {
		panic(err.Error())
	}
	return mdl
}
