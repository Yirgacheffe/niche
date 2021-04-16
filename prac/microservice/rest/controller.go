package main

type Controller struct {
	storage Storage
}

func New(storage Storage) *Controller {
	return &Controller{
		storage: storage,
	}
}

type Payload struct {
	Value string `json:"value"`
}
