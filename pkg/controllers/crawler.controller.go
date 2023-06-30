package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/2O23/crawler/internal/utils"
	"github.com/2O23/crawler/pkg/models"
)

func GetSites(w http.ResponseWriter,r *http.Request) {
		getSites := models.Sites()
		res, _ := json.Marshal(getSites)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
}

func CreateSite(w http.ResponseWriter,r *http.Request) {
	SiteData := &models.Site{}
	utils.ParseBody(r, SiteData)
	b := SiteData.CreateSite()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	}
