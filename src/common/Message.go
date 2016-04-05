package common

import (
	"encoding/json"
	"fmt"
)

const (
	IP_INVALID           = "%s is invalid ip."
	IP_ALREADY_USED      = "%s ip is already used."
	IP_NOT_ASSIGNED      = "%s not assigned to any device."
	IP_BLOCK_REQUIRED    = "IP Block CIDR is required."
	SUBNET_EXISTS        = "%s subnet exists."
	DEVICE_NOT_FOUND     = "%s device not foud."
	DEVICE_NAME_REQUIRED = "Device name required."
	ERROR                = "error"
	SUCCESS              = "succes"
	UNEXPECTED_ERROR     = "{\"status\": \"error\", \"message\":\"Unexpected error %s\"}"
	ERR_OPEN_DATAFILE    = "Error in Opening data file (%s), error %s"
	ERR_READ_DATAFILE    = "Error in reading data file (%s), error %s"
)

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (this *Message) Error() string {
	return this.Json()
}

func (this *Message) Json() string {
	data, err := json.Marshal(this)
	if err != nil {
		return fmt.Sprintf(UNEXPECTED_ERROR, err)
	}
	return string(data)
}

func NewMessage(format, status string, e ...interface{}) error {
	message := fmt.Sprintf(format, e...)
	return &Message{Status: ERROR, Message: message}
}

func NewError(format string, e ...interface{}) error {
	message := fmt.Sprintf(format, e...)
	return &Message{Status: "error", Message: message}
}
