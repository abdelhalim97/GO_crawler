package models

import (
	"github.com/2O23/crawler/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Site struct{
	Hostip string
	Domain string
	Files []File
	LastSeen int64 `gorm:"autoUpdateTime:milli"`
}

type Reponse struct{
	Site Site
	Status string
}

type File struct{
	Name string
	URL string
	SiteId Site
	LastSeen int64 `gorm:"autoUpdateTime:milli"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Site{})
	db.AutoMigrate(&File{})
}

func Sites(byDomain string) []Site {
	var Sites []Site
	db.Where("domain LIKE ?",byDomain).Find(&Sites)
	return Sites
}

func (site *Site) CreateSite() *Site {
	db.NewRecord(site)
	db.Create(&site)
	return site
}