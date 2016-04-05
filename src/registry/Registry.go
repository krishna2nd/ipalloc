package registry

import (
	. "bytes"
	. "common"
	"device"
	//"fmt"
	"sync"
)

type Registry struct {
	store   string
	Subnets map[string]*SubnetRegistry
	sync.Mutex
}

func New() *Registry {
	return new(Registry).init()
}

func (this *Registry) init() *Registry {

	this.Subnets = make(map[string]*SubnetRegistry)
	return this
}

func (this *Registry) AddSubnet(cidr string) error {

	_, ok := this.Subnets[cidr]
	if ok {
		return NewError(SUBNET_EXISTS, cidr)
	}

	subnetRegistry, err := NewSubnetRegistry(cidr)
	if nil != err {
		return NewError(UNEXPECTED_ERROR, err)
	}

	this.Subnets[cidr] = subnetRegistry

	return nil
}

func (this *Registry) Find(ip string) (*device.Device, error) {

	for _, subnet := range this.Subnets {
		device := subnet.Find(ip)

		if nil != device {
			return device, nil
		}
	}
	return nil, NewError(IP_NOT_ASSIGNED, ip)
}

func (this *Registry) Register(device *device.Device) error {

	if "" == device.Name {
		return NewError(DEVICE_NAME_REQUIRED)
	}

	if "" == device.CIDR {
		return NewError(IP_BLOCK_REQUIRED)
	}

	err := this.addDevice(device)

	if nil != err {
		return err
	}

	return this.Save()
}

func (this *Registry) UnRegister(ip string) error {
	var device *device.Device
	for _, subnet := range this.Subnets {
		device = subnet.Find(ip)
	}

	if nil != device {
		this.Subnets[device.CIDR].UnRegister(device.Name)
	}

	return this.Save()
}

func (this *Registry) addDevice(device *device.Device) error {
	if nil != device {
		_, ok := this.Subnets[device.CIDR]
		if !ok {
			this.AddSubnet(device.CIDR)
		}

		return this.Subnets[device.CIDR].Register(device)
	}
	return nil
}

func (this *Registry) TextBytes() *Buffer {
	var data = NewBuffer([]byte(""))
	for index := range this.Subnets {
		//fmt.Println(index)
		//fmt.Println(this.Subnets[index].Data)
		data.Write(this.Subnets[index].TextBytes().Bytes())
	}
	return data
}
