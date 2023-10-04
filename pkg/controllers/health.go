package controllers

type Heather interface {
	Check() Health
}

type Health struct{}

func NewHealth() *Health {
	return &Health{}
}

func (h Health) Check() (string, error) {
	return "The API is up and running", nil
}
