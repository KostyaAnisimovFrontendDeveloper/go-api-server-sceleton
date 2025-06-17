package page

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"server-skeleton/api_init"
	"server-skeleton/dictionary"
	"server-skeleton/utils"
	"strconv"
	"testing"
	"time"
)

var db *gorm.DB

func init() {
	api_init.TestInit("../../")
	db = api_init.InitGlobal.Dbh
}

func TestGetPagesList(t *testing.T) {
	clearDbTablePage(t)

	_, err := createPages(14)

	u, err := url.Parse(UriPage)
	if err != nil {
		panic(err)
	}

	const limit = 3

	q := u.Query()
	q.Set("limit", strconv.Itoa(limit))

	u.RawQuery = q.Encode()

	var result ResultListDTO
	w := sendRequest(t, u.String(), "GET", nil, &result)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, limit, len(result.List))
	assert.Equal(t, "Test page 14", result.List[0].Name)
}

func TestGetPagesListWithFilterName(t *testing.T) {
	clearDbTablePage(t)

	_, err := createPages(14)

	u, err := url.Parse(UriPage)
	if err != nil {
		panic(err)
	}

	const limit = 5

	q := u.Query()
	q.Set("limit", strconv.Itoa(limit))
	q.Set("names", "Test page 10,Test page 9,Test page 8")

	u.RawQuery = q.Encode()

	var result ResultListDTO
	w := sendRequest(t, u.String(), "GET", nil, &result)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(result.List))
	assert.Equal(t, "Test page 10", result.List[0].Name)
	assert.Equal(t, "Test page 9", result.List[1].Name)
	assert.Equal(t, "Test page 8", result.List[2].Name)
}

func TestGetPagesListPagination(t *testing.T) {
	clearDbTablePage(t)
	pages, err := createPages(14)

	u, err := url.Parse(UriPage)
	if err != nil {
		panic(err)
	}

	const limit = 3

	q := u.Query()
	q.Set("limit", strconv.Itoa(limit))
	q.Set("cursor", pages[11].ID.String())
	q.Set("lastTimestamp", pages[11].CreatedAt.Format(time.RFC3339Nano))

	u.RawQuery = q.Encode()

	var result ResultListDTO
	w := sendRequest(t, u.String(), "GET", nil, &result)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, limit, len(result.List))
	assert.Equal(t, "Test page 11", result.List[0].Name)
}

func TestGetPageById_NotFoundResult(t *testing.T) {
	clearDbTablePage(t)
	fakeId := "987fbc97-4bed-5078-9f07-9141ba07c9f3"
	var result ErrorResponseDto
	w := sendRequest(t, fmt.Sprintf(UriPageGetByIdS, fakeId), "GET", nil, &result)

	assert.Equal(t, http.StatusNotFound, w.Code)

	message := fmt.Sprintf(dictionary.PageByIdNotFound, fakeId)
	assert.Equal(t, message, result.Message)
}

func TestGetPageById_WrongIdFormat(t *testing.T) {
	clearDbTablePage(t)
	fakeId := "987fbc97"
	var result ErrorResponseDto
	w := sendRequest(t, fmt.Sprintf(UriPageGetByIdS, fakeId), "GET", nil, &result)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

	assert.Equal(t, "Key: 'RequestPageIdDTO.ID' Error:Field validation for 'ID' failed on the 'uuid' tag", result.Message)
}

func TestGetPageById_SuccessfulResult(t *testing.T) {
	clearDbTablePage(t)
	page := Page{Name: "Test page 1"}
	if err := db.Create(&page).Error; err != nil {
		t.Fatal(err)
	}

	var result PageItemResultDto
	w := sendRequest(t, fmt.Sprintf(UriPageGetByIdS, page.ID.String()), "GET", nil, &result)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, page.ID, result.ID)

}

func TestCreatePage_SuccessfulResult(t *testing.T) {
	clearDbTablePage(t)

	var result SuccessResponseDto
	post := map[string]string{
		"name": "Test page 1",
	}

	jsonData, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}

	w := sendRequest(t, UriPage, "POST", bytes.NewBuffer(jsonData), &result)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPutPageItem_SuccessfulResult(t *testing.T) {
	clearDbTablePage(t)

	now := time.Now()
	page := Page{Name: "Test page 1"}

	if err := db.Create(&page).Error; err != nil {
		t.Fatal(err)
	}

	page.DeletedAt = now
	jsonData, err := json.Marshal(page)
	if err != nil {
		panic(err)
	}
	var result SuccessResponseDto
	w := sendRequest(t, fmt.Sprintf(UriPageGetByIdS, page.ID.String()), "PUT", bytes.NewBuffer(jsonData), &result)
	assert.Equal(t, http.StatusOK, w.Code)
	updatedPage, err := GetOneById(RequestPageIdDTO{
		ID: page.ID.String(),
	})

	if err != nil {
		panic(err)
	}

	assert.Equal(t, page.ID, updatedPage.ID)
}

func TestPatchPageItem_SuccessfulResult(t *testing.T) {
	clearDbTablePage(t)

	page := Page{Name: "Test page 1"}

	if err := db.Create(&page).Error; err != nil {
		t.Fatal(err)
	}

	page.Name = "Test page 2"

	jsonData, err := json.Marshal(page)
	if err != nil {
		panic(err)
	}

	var result SuccessResponseDto
	w := sendRequest(t, fmt.Sprintf(UriPageGetByIdS, page.ID.String()), "PATCH", bytes.NewBuffer(jsonData), &result)
	assert.Equal(t, http.StatusOK, w.Code)
	updatedPage, err := GetOneById(RequestPageIdDTO{
		ID: page.ID.String(),
	})

	if err != nil {
		panic(err)
	}

	assert.Equal(t, page.Name, updatedPage.Name)
}

func TestDeletePageItem_SuccessfulResult(t *testing.T) {
	clearDbTablePage(t)

	page := Page{Name: "Test page 1"}

	if err := db.Create(&page).Error; err != nil {
		t.Fatal(err)
	}

	var result SuccessResponseDto
	w := sendRequest(t, fmt.Sprintf(UriPageGetByIdS, page.ID.String()), "DELETE", nil, &result)
	assert.Equal(t, http.StatusOK, w.Code)
	deletedPage, err := GetOneById(RequestPageIdDTO{
		ID: page.ID.String(),
	})

	if err != nil {
		panic(err)
	}

	assert.Equal(t, uuid.Nil, deletedPage.ID)
}

// === Sys
func clearDbTablePage(t *testing.T) {
	if err := db.Exec("truncate table pages restart identity cascade").Error; err != nil {
		utils.Dump(err)
		t.Fatal(err)
	}
}

func sendRequest(
	t *testing.T,
	uri string,
	method string,
	body io.Reader,
	result any,
) *httptest.ResponseRecorder {

	//Init

	router := gin.Default()
	InitPageRoutes(router)

	// Creating test request
	req, err := http.NewRequest(method, uri, body)

	if err != nil {
		t.Fatal(err)
	}

	//Sending test request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//Parsing result

	err = json.NewDecoder(w.Body).Decode(&result)
	if err != nil {
		t.Fatal(err)
	}

	return w
}

func createPages(length int) ([]Page, error) {
	var pages []Page
	for i := 1; i <= length; i++ {
		page := Page{Name: fmt.Sprintf("Test page %d", i)}
		if err := db.Create(&page).Error; err != nil {
			return pages, err
		}
		pages = append(pages, page)
	}
	return pages, nil
}
