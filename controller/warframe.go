package controller

import (
	"github.com/gin-gonic/gin"
	"warframes-tools-api/model"
)

func GetWarframes(c *gin.Context) {
	warframeType := c.Query("type")
	var warframes []*model.Warframe

	if warframeType == "normal" {
		warframes = model.FindWarframeNormals()
	}

	if warframeType == "prime" {
		warframes = model.FindWarframePrimes()
	}

	if warframes == nil {
		c.JSON(404, "not found")
		return
	}

	c.JSON(200, warframes)
}
