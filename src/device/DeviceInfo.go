package device

import (
	. "bytes"
	. "common"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Device struct {
	// Requirement insists to store IP Block,IP Address,Device
	Name string `json:"device"`
	CIDR string `json:"ip_block"`
	Ip   string `json:"ip_address"`
}

func New(cidr, ip, name string) *Device {
	return new(Device).init(name, ip, cidr)
}

func (this *Device) init(name, ip, cidr string) *Device {
	this.Name = name
	this.Ip = ip
	this.CIDR = cidr
	return this
}

func (d *Device) Json() (string, error) {
	data, err := json.Marshal(d)
	return string(data), err
}

func FromJson(data io.Reader) (*Device, error) {

	decoder := json.NewDecoder(data)
	device := &Device{}
	err := decoder.Decode(device)
	if err != nil {
		return nil, NewError(UNEXPECTED_ERROR, err)
	}

	return device, nil
}

func (this *Device) Serialize() *Buffer {
	return NewBufferString(fmt.Sprintf("%s,%s,%s \n", this.CIDR, this.Ip, this.Name))
}

func Deserialize(record string) *Device {

	records := strings.Split(record, ",")
	if 3 == len(records) {
		return New(records[0], records[1], records[2])
	}
	return nil
}
