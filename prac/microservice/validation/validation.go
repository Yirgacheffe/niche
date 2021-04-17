package validation

import "errors"

type Payload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type VError struct {
	error
}

func ValidatePayload(p *Payload) error {
	if p.Name == "" {
		return VError{errors.New("name is required")}
	}
	if p.Age <= 0 || p.Age >= 120 {
		return VError{
			errors.New(
				"age is required and must be a value greeter than 0 and less then 120"),
		}
	}
	return nil
}
