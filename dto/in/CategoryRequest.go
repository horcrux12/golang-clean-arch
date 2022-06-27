package in

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required,min=1,max=100"`
}

type CategoryUpdateRequest struct {
	ID   int64  `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,max=200,min=1"`
}

type CategoryRequest struct {
	ID       int64  `json:"id" validate:"required_if=IsUpdate true"`
	Name     string `json:"name" validate:"required,max=200,min=1"`
	IsUpdate bool
}
