package applicationModel

import "database/sql"

type ContextModel struct {
	Permission string
	AuthAccessModel
	LoggerModel
	ConnectionModel
}

type AuthAccessModel struct {
	UserID               string `json:"uid"`
	AuthenticationUserID int64  `json:"auid"`
	Locale               string `json:"locale"`
}

type ConnectionModel struct {
	Tx *sql.Tx
}
