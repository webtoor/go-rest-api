package web

type UserCreateRequest struct {
	Id       int    `json:"id"`
	Email    string `json:"email" validate:"required,min=1,max=100"`
	Password string `json:"password" validate:"required,min=1,max=100"`
}

type UserCreateResponse struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type UserUpdateRequest struct {
	Email string `json:"email" validate:"required,min=1,max=100"`
}

type UserUpdateResponse struct {
	Email string `json:"email"`
}
type GetUserResponse struct {
	Id    int    `json:"id"`
	Email string `json:"Email"`
}
