package fetchers

import (
	"github.com/timoisik/web-map/models"
	"net"
)

func FetchDomainIp(domain models.Domain) ([]net.IP, error) {
	ip, err := net.LookupIP(domain.GetUrl())

	if err != nil {
		return nil, err
	}

	return ip, nil
}