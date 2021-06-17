package user

import (
	"encoding/json"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

type createUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,min=2,max=50"`
	Password string `json:"name" validate:"required,min=8"`
}

type UserHandler struct {
	validate    *validator.Validate
	userService Service
	// all other dependencies
}

func NewUserHandler(s Service, v *validator.Validate /* all other dependencies */) UserHandler {
	return UserHandler{userService: s, validate: v}
}

// Different frameworks do this differently. For simplicity let's assume
// we have built-in http server
func (h UserHandler) Create(req *http.Request) error {
	input := createUserRequest{}

	// unmarshal incoming input into pre-defined structure
	defer req.Body.Close()
	if err := json.NewEncoder(req.Body).Encode(&input); err != nil {
		return err
	}

	// Just terminate the request if the input is not valid
	if err := h.validate.Struct(input); err != nil {
		return err
	}

	// Map to domain entity
	user := User{
		Email:    input.Email,
		Name:     input.Name,
		Password: input.Password,
	}

	// Finally pass the entity into service to get the job done!
	return h.userService.Create(user)
}
