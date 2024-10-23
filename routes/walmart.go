package routes

import (
	"errors"
	"go-parser/helper"
	"go-parser/parsers"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func Walmart_ParseHtml(filename string) bool {
	content, err := os.ReadFile("./data/" + filename + ".html")
	if err != nil {
		println(err)
		return false
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
	if err != nil {
		println(err)
		return false
	}
	json, err := Walmart_ExtractJson(doc)
	if err != nil {
		println(err)
		return false
	}
	var result interface{}
	if parsers.Walmart_IsSearchPage(json) {
		result = parsers.Walmart_BrowsePageScraper(json)
	} else if parsers.Walmart_IsProductPage(json) {
		result = parsers.Walmart_ProductPageScraper(json)
	} else if parsers.Walmart_IsReviewPage(json) {
		result = parsers.Walmart_ReviewPageScraper(json)
	} else if parsers.Walmart_IsCategoryPage(json) {
		result = parsers.Walmart_CategoryPageScraper(json)
	} else if parsers.Walmart_IsBrowsePage(json) {
		result = parsers.Walmart_BrowsePageScraper(json)
	} else if parsers.Walmart_IsShopPage(json) {
		result = parsers.Walmart_ShopPageScraper(json)
	} else {
		println("Not found")
		return false
	}
	// return true
	return helper.SaveJsonFile(result, filename)
}

func Walmart_ExtractJson(doc *goquery.Document) (*goquery.Selection, error) {
	jsonTag := doc.Find("script#__NEXT_DATA__").First()
	if jsonTag.Length() > 0 {
		return jsonTag, nil
	}
	return nil, errors.New("parsing failed")
}

func Walmart_PostRequest(c *gin.Context) {
	var postData helper.RequestData
	// Get post data
	if err := c.BindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// parse html using goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(postData.Html))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	json, err := Walmart_ExtractJson(doc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var result interface{}
	if parsers.Walmart_IsSearchPage(json) {
		result = parsers.Walmart_BrowsePageScraper(json)
	} else if parsers.Walmart_IsProductPage(json) {
		result = parsers.Walmart_ProductPageScraper(json)
	} else if parsers.Walmart_IsReviewPage(json) {
		result = parsers.Walmart_ReviewPageScraper(json)
	} else if parsers.Walmart_IsCategoryPage(json) {
		result = parsers.Walmart_CategoryPageScraper(json)
	} else if parsers.Walmart_IsBrowsePage(json) {
		result = parsers.Walmart_BrowsePageScraper(json)
	} else if parsers.Walmart_IsShopPage(json) {
		result = parsers.Walmart_ShopPageScraper(json)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported page"})
		return
	}
	// id := uuid.New()
	// SaveJsonFile(result, id.String())
	c.JSON(http.StatusOK, result)
}
