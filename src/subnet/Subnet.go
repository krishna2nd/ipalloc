package subnet

import (
	"errors"
	"net"
)

type Subnet struct {
	ip          string
	hosts       []string
	broadcastIp string
	CIDR        string

	available []string
}

func New(cidr string) (*Subnet, error) {
	return new(Subnet).init(cidr)
}

func (s *Subnet) init(cidr string) (*Subnet, error) {

	s.CIDR = cidr

	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return s, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); next(ip) {
		ips = append(ips, ip.String())
	}

	if len(ips) > 0 {
		s.ip = ips[0]
	}

	if len(ips) > 1 {
		s.broadcastIp = ips[len(ips)-1]
	}

	if len(ips) > 2 {
		s.hosts = ips[1 : len(ips)-1]
		s.available = s.hosts[0:]
	}

	return s, nil
}

func next(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func (s *Subnet) HostIps() []string {
	return s.hosts
}

func (s *Subnet) LoadAvailableIps(allocatedIps []string) {
	allocated := false
	s.available = make([]string, 0, len(s.hosts)-len(allocatedIps))
	for ip := range s.hosts {
		allocated = false
		for aip := range allocatedIps {
			if allocatedIps[aip] == s.hosts[ip] {
				allocated = true
				break
			}
		}
		if !allocated {
			s.available = append(s.available, s.hosts[ip])
		}
	}
	//fmt.Println(s.available[0])
}
func (s *Subnet) NextAvailableIp() (string, error) {

	if len(s.available) > 0 {
		var ip = s.available[0]
		s.available = s.available[1:]
		return ip, nil
	}

	return "", errors.New("No Free IP available.")
}

func (s *Subnet) FreeIp(ip *string) {
	s.available = append(s.available, *ip)
}

func (s *Subnet) Reload(ip string) {
	var index int
	for index = 0; index < len(s.available); index++ {
		if s.available[index] == ip {
			break
		}
	}
	if index < len(s.available) {
		s.available = append(s.available[:index], s.available[index+1:]...)
	}

}

func (s *Subnet) IsValid(ip string) bool {

	in := net.ParseIP(ip)
	if in.To4() == nil {
		return false
	}

	_, ipnet, err := net.ParseCIDR(s.CIDR)

	if err != nil {
		return false
	}
	//fmt.Println(ip, in, ipnet.Contains(in))
	return ipnet.Contains(in)
}
