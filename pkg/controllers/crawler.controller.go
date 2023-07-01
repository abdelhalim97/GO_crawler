package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	// "github.com/2O23/crawler/internal/utils"
	"github.com/2O23/crawler/internal/utils"
	"github.com/2O23/crawler/pkg/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
)

func GetSites(w http.ResponseWriter,r *http.Request) {
	query := r.URL.Query().Get("query")
	filter := r.URL.Query().Get("filter")

	if(query!="by-domain" && query!="by-date" && query!="by-visit"){
		response:=&models.ReponseSites{Sites:[]models.Site{},Status: "nok" }
		res, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
	} else{
		getSites := models.GetSites(query,filter)

		response:=&models.ReponseSites{Sites:getSites,Status: "ok" }

		res, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
}
}

func CreateSite(w http.ResponseWriter,r *http.Request) {
	var files []models.File
	
	SiteData := &models.Site{}
	utils.ParseBody(r, SiteData)

	resDomain, err := http.Get(SiteData.Domain)
	if err != nil {
		log.Fatal(err)
		w.Write([]byte("nok"))
	}
	defer resDomain.Body.Close()
	if resDomain.StatusCode != 200 {
	  log.Fatalf("status code error: %d %s", resDomain.StatusCode, resDomain.Status)
	}
  
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resDomain.Body)
	if err != nil {
	  log.Fatal(err)
	  w.Write([]byte("nok"))
	}
  
	// Find the review items
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		name:=s.Text()
		content,isHrefExists := s.Attr("href")
		if isHrefExists {
			fmt.Printf("Review %d: %s\n", i, content)
			file := models.File{
			Name:name,URL:content,
			}
			files=append(files,file)
			}
		  
	  })

	siteWithFiles:=models.Site{
	Hostip: SiteData.Hostip,Domain: SiteData.Domain,Files: files,
  	}
  	cereatedSite:=siteWithFiles.CreateSite()
	response:=&models.ReponseSite{Site: *cereatedSite,  Status: "ok"}
  
	res, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetFiles(w http.ResponseWriter,r *http.Request) {
	query := r.URL.Query().Get("query")
	filter := r.URL.Query().Get("filter")
	if(query!="by-domain" && query!="by-date" && query!="by-visit"){
		response:=&models.FileReponse{Files:[]models.File{},Status: "nok",Filter: filter,Query: query }
		res, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
	} else{
		getFiles := models.GetFiles(query,filter)
		response:=&models.FileReponse{Files: getFiles,  Status: "ok",Query:query,Filter: filter}

		res, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func UpdateSite(w http.ResponseWriter,r *http.Request){
	var updateSite = &models.File{}
	utils.ParseBody(r, updateSite)
	vars := mux.Vars(r)
	siteId := vars["id"]
	ID, err := strconv.ParseInt(siteId, 0, 0)
	if err != nil {
		fmt.Println("error parsing")
	}
	SiteDetails, db := models.SiteById(ID)
	SiteDetails.LastSeen = time.Now().UnixMilli()
	db.Save(&SiteDetails)
	res, _ := json.Marshal(SiteDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateFile(w http.ResponseWriter,r *http.Request){
	var updateFile = &models.File{}
	utils.ParseBody(r, updateFile)
	vars := mux.Vars(r)
	fileId := vars["id"]
	ID, err := strconv.ParseInt(fileId, 0, 0)
	if err != nil {
		fmt.Println("error parsing")
	}
	FileDetails, db := models.FileById(ID)
	FileDetails.LastSeen = time.Now().UnixMilli()
	db.Save(&FileDetails)
	res, _ := json.Marshal(FileDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}