package applicationModel

type ContextModel struct {
	Permission string
	AuthAccessModel
	LoggerModel
}

type AuthAccessModel struct {
	UserID               string `json:"uid"`
	AuthenticationUserID int64  `json:"auid"`
	Locale               string `json:"locale"`
}
