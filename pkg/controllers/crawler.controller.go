package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "github.com/2O23/crawler/internal/utils"
	"github.com/2O23/crawler/internal/utils"
	"github.com/2O23/crawler/pkg/models"
	"github.com/PuerkitoBio/goquery"
)

func GetSites(w http.ResponseWriter,r *http.Request) {
	byDomain := r.URL.Query().Get("by-domain")

	getSites := models.Sites(byDomain)
	res, _ := json.Marshal(getSites)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateSite(w http.ResponseWriter,r *http.Request) {
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
		  // For each item found, get the title
			content,isHrefExists := s.Attr("href")
			if isHrefExists {
				fmt.Printf("Review %d: %s\n", i, content)
			}
		  
	  })
		
	// b := SiteData.CreateSite()
response:=&models.Reponse{Site: *SiteData.CreateSite(),  Status: "ok"}
	res, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	}
