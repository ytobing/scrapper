package html

import (
	m "github.com/ytobing/scrapper/models"
)

func RegisterRoutes(c m.App) {

	htmlSvc := InitHandler()
	c.Router.GET("/", htmlSvc.ServeHTML)

}
