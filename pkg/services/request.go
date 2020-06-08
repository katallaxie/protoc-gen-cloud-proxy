package services

type Request interface {
	Do() error
}
