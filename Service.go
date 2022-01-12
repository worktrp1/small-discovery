package smalldiscovery

import (
	"net"
)

type Service struct {
	Name        string
	IP          string
	IPAddr      net.IP
	LanIP       string
	LanIPAddr   net.IP
	Port        string
	Domain      string
	Description string
	Hostname    string
	Status      string
}

func InitService(name, description, port string) (s Service) {
	s.Name = name
	s.Description = description
	s.Port = port
	s.Hostname = GetOSHostname()

	s.SetIP(GetWANIP())
	s.LanIPAddr = GetLANIP()
	s.LanIP = string(s.LanIPAddr)

	return
}

func (s *Service) SetIP(ip string) {
	s.IP = ip
	s.IPAddr = []byte(ip)
}
