package tests

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"server-skeleton/api/page/dto"
	"server-skeleton/api/page/model"
	"server-skeleton/api/page/route"
	"server-skeleton/api_init"
	"testing"
)

func TestPage_ValidationInput(t *testing.T) {

	fmt.Println("TestPage_ValidationInput")
	fmt.Println("Init Fixtures")
	page, err := fixtures(api_init.InitGlobal.Dbh)

	if err != nil {
		t.Fatal(err)
	}

	route.InitPageRoutes(Router)
	uri := fmt.Sprintf(route.UriPageGetByIdS, page.ID.String())

	req, err := http.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	var resultDto dto.ResponsePageDTO

	err = json.NewDecoder(w.Body).Decode(&resultDto)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, resultDto.ID.String(), page.ID.String())
}

func fixtures(db *gorm.DB) (*model.Page, error) {
	err := db.Exec(`truncate table public.pages restart identity cascade;`).Error
	if err != nil {
		return nil, err
	}

	page := model.Page{Name: "test page"}

	err = db.Create(&page).Error
	if err != nil {
		return nil, err
	}

	return &page, nil
}
