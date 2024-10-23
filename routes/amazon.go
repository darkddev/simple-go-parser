package routes

import (
	"go-parser/helper"
	"go-parser/parsers"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func Amazon_ParseHtml(filename string) bool {
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
	var result interface{}
	if parsers.Amazon_IsSearchPage(doc) {
		result = parsers.Amazon_SearchPagesScraper(doc)
	} else if parsers.Amazon_IsReviewPage(doc) {
		result = parsers.Amazon_ReviewPagesScraper(doc)
	} else {
		result = parsers.Amazon_ProductPagesScraper(doc)
	}
	return helper.SaveJsonFile(result, filename)

}

func Amazon_PostRequest(c *gin.Context) {
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
	var result interface{}
	if parsers.Amazon_IsSearchPage(doc) {
		result = parsers.Amazon_SearchPagesScraper(doc)
	} else if parsers.Amazon_IsReviewPage(doc) {
		result = parsers.Amazon_ReviewPagesScraper(doc)
	} else {
		result = parsers.Amazon_ProductPagesScraper(doc)
	}
	// SaveJsonFile(result, "result")
	c.JSON(http.StatusOK, result)
}
