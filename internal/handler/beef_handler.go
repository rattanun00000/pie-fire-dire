package handler

import (
	"net/http"

	"pie-fire-dire/internal/service"

	"github.com/gin-gonic/gin"
)

type BeefHandler struct {
	beefService  *service.BeefService
	baconService *service.MeatIpsumService
}

func NewBeefHandler(beefService *service.BeefService, baconService *service.MeatIpsumService) *BeefHandler {
	return &BeefHandler{
		beefService:  beefService,
		baconService: baconService,
	}
}

func (h *BeefHandler) GetBeefSummary(c *gin.Context) {
	if cachedResult, exists := h.beefService.GetCachedResult(); exists {
		c.JSON(http.StatusOK, gin.H{
			"beef": cachedResult,
		})
		return
	}

	text, err := h.baconService.FetchMeatIpsum()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถดึงข้อมูลจาก BaconIpsum ได้"})
		return
	}

	beefCounts := h.beefService.CountBeefTypes(text)

	h.beefService.SetCachedResult(beefCounts)

	c.JSON(http.StatusOK, gin.H{
		"beef": beefCounts,
	})
}
