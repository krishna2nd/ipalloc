package registry

import (
	. "bytes"
	. "common"
	"device"
	//"fmt"
	"subnet"
)

type SubnetRegistry struct {
	Data   map[string]*device.Device
	Subnet *subnet.Subnet
}

func NewSubnetRegistry(cidr string) (*SubnetRegistry, error) {
	return new(SubnetRegistry).init(cidr)
}

func (this *SubnetRegistry) init(cidr string) (*SubnetRegistry, error) {
	this.Data = make(map[string]*device.Device)

	subnet, err := subnet.New(cidr)
	if nil != err {
		return nil, NewError(UNEXPECTED_ERROR, err)
	}

	this.Subnet = subnet
	return this, err

}

func (this *SubnetRegistry) TextBytes() *Buffer {
	data := NewBuffer([]byte(""))
	for _, device := range this.Data {
		data.Write(device.Serialize().Bytes())
	}

	return data
}

func (this *SubnetRegistry) Find(ip string) *device.Device {
	device, ok := this.Data[ip]

	if ok {
		return device
	}
	return nil
}

func (this *SubnetRegistry) Register(device *device.Device) error {
	var err error

	reload := false

	if "" == device.Ip {
		device.Ip, err = this.Subnet.NextAvailableIp()
		device.CIDR = this.Subnet.CIDR
	} else {
		_, ok := this.Data[device.Ip]

		if ok {
			return NewError(IP_ALREADY_USED, device.Ip)
		}

		if !this.Subnet.IsValid(device.Ip) {
			return NewError(IP_INVALID, device.Ip)
		}
		reload = true
	}

	if err != nil {
		return err
	}

	//fmt.Println(device.Ip)
	//fmt.Println(this.Subnet.CIDR)
	//fmt.Println(this.Data)
	this.Data[device.Ip] = device

	if reload {
		this.Subnet.Reload(device.Ip)
	}

	return nil
}

func (this *SubnetRegistry) UnRegister(name string) error {
	var ip *string
	for k, v := range this.Data {
		if name == (v.Name) {
			ip = &k
			delete(this.Data, k)
			this.Subnet.FreeIp(ip)
		}
	}

	return nil
}

func keys(mapData map[string]*device.Device) []string {

	keys := make([]string, 0, len(mapData))
	for k := range mapData {
		keys = append(keys, k)
	}
	return keys
}
