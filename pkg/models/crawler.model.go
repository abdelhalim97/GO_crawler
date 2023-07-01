package models

import (
	"github.com/2O23/crawler/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Site struct{
	ID        int    `json:"id" gorm:"autoIncrement; primaryKey"`
	Hostip string
	Domain string
	Files []File
	LastSeen int64 `gorm:"autoUpdateTime:milli"`
}

type ReponseSites struct{
	Sites []Site
	Status string
}

type ReponseSite struct{
	Site Site
	Status string
}

type FileReponse struct{
	Files []File
	Status string
	Filter string
	Query string
}

type File struct{
	ID        int    `json:"id" gorm:"autoIncrement; primaryKey"`
	Name string
	URL string
	SiteId int
	LastSeen int64  `gorm:"autoUpdateTime:milli; default:nill"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Site{})
	db.AutoMigrate(&File{})
}

func SiteById(Id int64) (*Site, *gorm.DB) {
	var getSite Site
	db := db.Where("ID=?", Id).Find(&getSite)
	return &getSite, db
}

func GetSites(query string,filter string) []Site {
	var Sites []Site
	if query=="by-domain"{
		db.Where("domain LIKE ?",filter).Find(&Sites)
	} else if query=="by-date" {
		db.Order("last_seen asc").Find(&Sites)
	} else if query=="by-visit" {
		db.Order("last_seen asc").Find(&Sites)
	} else{
 		panic("query can only have by-domain or by-date or by-visit")
	}
	return Sites
}

func (site *Site) CreateSite() *Site {
	db.NewRecord(site)
	db.Create(&site)
	return site
}

func FileById(Id int64) (*File, *gorm.DB) {
	var getFile File
	db := db.Where("ID=?", Id).Find(&getFile)
	return &getFile, db
}

func GetFiles(query string,filter string) []File {
	var Files []File
	if query=="by-domain"{
		db.Where("url LIKE ?",filter).Find(&Files)
	} else if query=="by-date" {
		db.Order("last_seen asc").Find(&Files)
	} else if query=="by-visit" {
		db.Order("last_seen asc").Find(&Files)
	} else{
 		panic("query can only have by-domain or by-date or by-visit")
	}
	return Files
}