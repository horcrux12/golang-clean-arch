package in

type UserRequest struct {
	ID        int64  `json:"id" validate:"required_if=IsUpdate true"`
	Username  string `json:"username" validate:"required,gte=6,max=256,is-username"`
	Password  string `json:"password" validate:"required,gte=8,max=256"`
	FirstName string `json:"first_name" validate:"required_if=IsUpdate false,min=1,max=256"`
	LastName  string `json:"last_name" validate:"omitempty,min=1,max=256"`
	IsUpdate  bool
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required,min=1,max=256"`
	Password string `json:"password" validate:"required,min=1,max=256"`
}
