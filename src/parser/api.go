package parser

import (
	m "github.com/ytobing/scrapper/models"
)

func RegisterRoutes(c m.App) {

	parserSvc := InitHandler()
	c.Router.POST("/parse", parserSvc.ParseURL)

}
