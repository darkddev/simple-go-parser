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

func Google_ParseHtml(filename string) bool {
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
	result := parsers.Google_SearchPagesScraper(doc)
	return helper.SaveJsonFile(result, filename)
}

func Google_PostRequest(c *gin.Context) {
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
	result = parsers.Google_SearchPagesScraper(doc)
	// SaveJsonFile(result, "result")
	c.JSON(http.StatusOK, result)
}
