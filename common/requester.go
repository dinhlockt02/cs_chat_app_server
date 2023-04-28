package common

type Requester interface {
	GetId() string
	GetDeviceId() string
}
