package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Domain struct {
	gorm.Model
	FetchedAt time.Time
	CrawledAt time.Time
	Name string
	Tld string
	IpAddresses []Ip `gorm:"many2many:domain_ip_address;"`
}

func (d *Domain) GetUrl() string {
	return d.Name + "." + d.Tld
}

func (d *Domain) HasBeenCrawled() bool {
	return !d.CrawledAt.IsZero()
}

func (d *Domain) MarkAsCrawled() {
	d.CrawledAt = time.Now()
	Db.Save(d)
}