package main

type Controller struct {
	ValidatePayload func(p *Payload) error
}

func NewController() *Controller {
	return &Controller{
		ValidatePayload: ValidatePayload,
	}
}
