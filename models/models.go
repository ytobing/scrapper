package models

import "github.com/julienschmidt/httprouter"

type App struct {
	Router *httprouter.Router
}

type Data struct {
	URL                string         `json:"url"`
	HTMLVersion        string         `json:"html_version"`
	PageTitle          string         `json:"page_title"`
	HeadingsCount      map[string]int `json:"headings_count"`
	LinkDetails        LinkDetails    `json:"link_details"`
	IsContainLoginForm bool           `json:"is_contain_login_form"`
}

type LinkDetails struct {
	InternalLink     int `json:"internal_link"`
	ExternalLink     int `json:"external_link"`
	InaccessibleLink int `json:"inaccessible_link"`
}

type ParseURL struct {
	URL string `json:"url"`
}
