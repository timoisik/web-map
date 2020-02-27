package models

import "github.com/jinzhu/gorm"

type Ip struct {
	gorm.Model
	Address string
	Domains []*Domain `gorm:"many2many:domain_ip_address;"`
}