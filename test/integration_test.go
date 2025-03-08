package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"pie-fire-dire/internal/handler"
	"pie-fire-dire/internal/model"
	"pie-fire-dire/internal/service"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCountBeefTypes(t *testing.T) {
	beefService := service.NewBeefService()

	text := "T-bone ribeye sirloin, pork belly chicken turkey. Brisket filet mignon corned beef. Fatback pork hamburger."

	result := beefService.CountBeefTypes(text)

	expectedBeef := map[string]int{
		"t-bone":       1,
		"ribeye":       1,
		"sirloin":      1,
		"brisket":      1,
		"filet mignon": 1,
		"corned beef":  1,
	}

	fmt.Println("ผลลัพธ์จากการนับ:", result)

	for beefType, expectedCount := range expectedBeef {
		count, exists := result[beefType]
		if !exists {
			t.Errorf("ไม่พบชนิดเนื้อวัว '%s' ในผลลัพธ์", beefType)
		} else if count != expectedCount {
			t.Errorf("จำนวนของ '%s' ควรเป็น %d แต่ได้ %d", beefType, expectedCount, count)
		}
	}

	nonBeefMeats := []string{"pork", "chicken", "turkey", "pork belly", "hamburger", "fatback"}
	for _, meatType := range nonBeefMeats {
		if _, exists := result[meatType]; exists {
			t.Errorf("พบเนื้อที่ไม่ใช่เนื้อวัว '%s' ในผลลัพธ์", meatType)
		}
	}
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	beefService := service.NewBeefService()
	baconService := service.NewMeatIpsumService()
	beefHandler := handler.NewBeefHandler(beefService, baconService)

	r.GET("/beef/summary", beefHandler.GetBeefSummary)

	return r
}

func TestBeefSummaryEndpoint(t *testing.T) {
	router := setupTestRouter()

	req := httptest.NewRequest("GET", "/beef/summary", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status OK but got %v", w.Code)
	}

	var response map[string]map[string]int
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("error unmarshaling response: %v", err)
	}

	beef, exists := response["beef"]
	if !exists {
		t.Error("expected response to have key 'beef'")
	}

	jsonBytes, _ := json.Marshal(response)
	fmt.Println("Response JSON:", string(jsonBytes))

	beefTypes := model.NewBeefTypes()
	for beefType := range beef {
		if !beefTypes.TypesMap[beefType] {
			t.Errorf("พบเนื้อที่ไม่ใช่เนื้อวัว '%s' ในผลลัพธ์", beefType)
		}
	}
}
