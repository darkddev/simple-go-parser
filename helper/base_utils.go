package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func NormalizeText(text string) string {
	text = strings.ReplaceAll(text, "‎", "")
	re := regexp.MustCompile(`\s+`)
	replacedStr := re.ReplaceAllString(strings.TrimSpace(text), " ")
	newStr := strings.ReplaceAll(replacedStr, "\u0026", "&")
	return newStr
}

func NormalizeTextWithReturn(text string) string {
	text = strings.ReplaceAll(text, "‎", "")
	re := regexp.MustCompile(`[^\S\r\n]+`)
	replacedStr := re.ReplaceAllString(strings.TrimSpace(text), " ")
	newStr := strings.ReplaceAll(replacedStr, "\u0026", "&")
	return newStr
}

func NormalizeImage(url string) string {
	if url[:15] != "https://m.media" {
		return ""
	}
	segs := strings.Split(url, "/")
	last := len(segs) - 1
	imageName := segs[last]
	nameSegs := strings.Split(imageName, ".")
	imageName = nameSegs[0] + "." + nameSegs[len(nameSegs)-1]
	segs[last] = imageName
	return strings.Join(segs, "/")
}

func ExtractPrice(text string) string {
	for _, price := range strings.Split(text, " ") {
		pattern := `(.\d+(\.\d+)?)`
		regex := regexp.MustCompile(pattern)
		matches := regex.FindAllString(price, -1)
		if len(matches) > 0 {
			return matches[0]
		}
	}
	return ""
}

func ExtractBaseUrl(url string) string {
	parts := strings.Split(url, "/")[:3]
	return strings.Join(parts, "/")
}

func NormalizeUrl(baseurl string, path string) string {

	if path[:4] == "http" {
		return strings.ReplaceAll(path, "\u0026", "&")
	} else if path != "" {
		return strings.ReplaceAll(baseurl+path, "\u0026", "&")
	}
	return ""
}

func ExtractText(s *goquery.Selection, splitter string) string {
	var allText string = ""
	if s.Get(0).Data != "span" && s.Children().Length() > 0 {
		s.Children().Each(func(index int, item *goquery.Selection) {
			allText += ExtractText(item, splitter)
		})
	} else {
		if s.Get(0).Type == html.TextNode || s.Get(0).Data == "p" || s.Get(0).Data == "span" || s.Get(0).Data == "h2" || s.Get(0).Data == "h3" || s.Get(0).Data == "h4" {
			allText += NormalizeText(s.Text()) + splitter
		}
	}
	return strings.TrimSpace(allText)
}

func ExtractText1(s *goquery.Selection, splitter string) string {
	var allText string = ""
	println(s.Get(0).Type, s.Get(0).Data, s.Text())
	if s.Get(0).Data != "span" && s.Children().Length() > 0 {
		s.Children().Each(func(index int, item *goquery.Selection) {
			allText += ExtractText(item, splitter)
		})
	} else {
		if s.Get(0).Type == html.TextNode || s.Get(0).Data == "p" || s.Get(0).Data == "span" || s.Get(0).Data == "h2" || s.Get(0).Data == "h3" || s.Get(0).Data == "h4" {
			allText += NormalizeText(s.Text()) + splitter
		}
	}
	return allText
}

func ConvertTableToMap(s *goquery.Selection) map[string]string {
	dataMap := make(map[string]string)

	// Iterate over each table row
	s.Find("tbody tr").Each(func(i int, row *goquery.Selection) {
		// Extract the key and value from each row
		key := NormalizeText(row.Find("th").Text())
		value := NormalizeText(row.Find("td").Text())
		// Add the key-value pair to the map
		dataMap[key] = value
	})
	// Print the map
	return dataMap
}

func customMarshal(v interface{}) ([]byte, error) {
	var buf []byte
	buf, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	// Replace \u0026 with &
	buf = bytes.Replace(buf, []byte("\\u0026"), []byte("&"), -1)
	buf = bytes.Replace(buf, []byte("\\u003c"), []byte("<"), -1)
	buf = bytes.Replace(buf, []byte("\\u003e"), []byte(">"), -1)
	return buf, nil
}

func SaveJsonFile(result interface{}, filename string) bool {
	// Marshal the struct into JSON
	jsonData, err := customMarshal(result)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return false
	}
	// Write the JSON data to a file
	file, err := os.Create("./data/" + filename + ".json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return false
	}
	defer file.Close()
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return false
	}
	return true
}

func ExtractIntFromPattern(text string, pattern string) (int, error) {
	// Compile the regular expression pattern
	re := regexp.MustCompile(pattern)
	// Find all matches of the pattern in the input string
	matches := re.FindAllStringSubmatch(text, -1)
	// Extract the first match (which should be the number)
	var result string
	if len(matches) > 0 {
		result = matches[0][1]
	} else {
		return 0, errors.New("Pattern not found")
	}
	num, err := strconv.Atoi(strings.ReplaceAll(result, ",", ""))
	return num, err
}

func ExtractFloatFromPattern(text string, pattern string) (float64, error) {
	// Compile the regular expression pattern
	re := regexp.MustCompile(pattern)
	// Find all matches of the pattern in the input string
	matches := re.FindAllStringSubmatch(text, -1)
	// Extract the first match (which should be the number)
	var result string
	if len(matches) > 0 {
		result = matches[0][1]
	} else {
		return 0, errors.New("Pattern not found")
	}
	num, err := strconv.ParseFloat(strings.ReplaceAll(result, ",", ""), 64)
	return num, err
}
