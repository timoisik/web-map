package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Domain struct {
	gorm.Model
	FetchedAt *time.Time
	Name string
	Tld string
	// Ips []net.IP
	IpAddresses []Ip `gorm:"many2many:domain_ip_address;"`
}

func (d *Domain) HasBeenChecked() bool {
	return d.FetchedAt.IsZero()
}

func (d *Domain) GetUrl() string {
	return d.Name + "." + d.Tld
}