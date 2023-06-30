package models

import (
	"github.com/2O23/crawler/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Site struct{
	gorm.Model
	Hostip string
	Domain string
	// Files []File
}

type File struct{
	gorm.Model
	Name string
	URL string
	// SiteId Site
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Site{})
}

func Sites() []Site {
	var Sites []Site
	db.Find(&Sites)
	return Sites
}

func (site *Site) CreateSite() *Site {
	db.NewRecord(site)
	db.Create(&site)
	return site
}