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

func (d *Domain) Exists() bool {
	var domain Domain
	Db.Where("name = ? and tld = ?", d.Name, d.Tld).Find(&domain)
	return domain.Name != ""
}

func (d *Domain) HasBeenCrawled() bool {
	var domain Domain
	Db.Where("name = ? and tld = ?", d.Name, d.Tld).Find(&domain)
	return !domain.CrawledAt.IsZero()
}

func (d *Domain) MarkAsCrawled() {
	var domain Domain
	Db.Where("name = ? and tld = ?", d.Name, d.Tld).Find(&domain)
	domain.CrawledAt = time.Now()
	Db.Save(&domain)
}