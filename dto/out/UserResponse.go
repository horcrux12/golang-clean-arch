package out

type UserLoginResponse struct {
	AuthorizationToken string `json:"authorization_token"`
}

type UserDetailResponse struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
