package controllers

import (
	"net/http"

	"github.com/labstack/echo"

	"lawencon.com/imamfarisi/models"
	"lawencon.com/imamfarisi/services"
)

var candidatProfileService services.CandidatProfileService = services.CandidatProfileServiceImpl{}

func SetCandidatProfile(c *echo.Group) {
	c.POST("/candidate", createCandidatProfile)
}

func createCandidatProfile(c echo.Context) error {
	data := new(models.CandidateProfiles)

	if err := c.Bind(data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	var err = candidatProfileService.CreateCandidat(data)
	if err == nil {
		return res(c, data)
	}
	return resErr(c, err)
}
