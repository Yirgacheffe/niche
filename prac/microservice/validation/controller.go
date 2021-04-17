package validation

type Controller struct {
	ValidatePayload func(p *Payload) error
}

func New() *Controller {
	return &Controller{
		ValidatePayload: ValidatePayload,
	}
}
