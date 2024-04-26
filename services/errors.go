package services

import "fmt"

// ServiceError is a custom error type for services.
type ServiceError struct {
	Code int
	Msg  string
}

func (e ServiceError) Error() string {
	return e.Msg
}

// LogMsg returns the error and service that the error happened in.
func (e ServiceError) LogMsg(serviceName string) string {
	return fmt.Sprintf("[%s]: [%s]", serviceName, e.Msg)
}
