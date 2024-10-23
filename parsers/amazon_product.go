package parsers

import (
	"go-parser/helper"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type AmazonProductData struct {
	AplusPresent           bool                   `json:"aplus_present"`
	AvailabilityStatus     string                 `json:"availability_status"`
	AverageRating          float64                `json:"average_rating"`
	Brand                  string                 `json:"brand"`
	BrandURL               string                 `json:"brand_url"`
	CustomizationOptions   map[string]interface{} `json:"customization_options"`
	FeatureBullets         []string               `json:"feature_bullets"`
	FullDescription        string                 `json:"full_description"`
	Images                 []string               `json:"images"`
	IsCouponExists         bool                   `json:"is_coupon_exists"`
	Model                  string                 `json:"model"`
	Name                   string                 `json:"name"`
	Price                  float64                `json:"price"`
	PriceString            string                 `json:"price_string"`
	PriceCurrency          string                 `json:"price_currency"`
	ProductCategory        string                 `json:"product_category"`
	ProductInformation     map[string]interface{} `json:"product_information"`
	ShippingPrice          string                 `json:"shipping_price"`
	ShipsFrom              string                 `json:"ships_from"`
	SoldBy                 string                 `json:"sold_by"`
	TotalReviews           int                    `json:"total_reviews"`
	TotalAnsweredQuestions interface{}            `json:"total_answered_questions"`
}

type AmazonProductResult struct {
	Data   AmazonProductData `json:"data"`
	Status string            `json:"status"`
	URL    string            `json:"url"`
}

func parseBtn(btn *goquery.Selection, first bool) map[string]interface{} {
	record := make(map[string]interface{})
	textDiv := btn.Find(".twisterTextDiv").First()
	if textDiv.Length() > 0 {
		record["value"] = helper.NormalizeText(textDiv.Text())
	}
	imageDiv := btn.Find(".twisterImageDiv").First()
	if imageDiv.Length() > 0 {
		imgTag := imageDiv.Find("img").First()
		record["image"] = helper.NormalizeImage(imgTag.AttrOr("src", ""))
		record["value"] = helper.NormalizeText(imgTag.AttrOr("alt", ""))
	}
	slotDiv := btn.Find(".twisterSlotDiv").First()
	if slotDiv.Length() > 0 {
		record["price"] = helper.ExtractPrice(helper.NormalizeText(slotDiv.Text()))
	}
	record["is_selected"] = first
	return record
}

func Amazon_ProductPagesScraper(doc *goquery.Document) AmazonProductResult {
	var result AmazonProductResult
	var data AmazonProductData
	baseUrl := "https://www.amazon.com"
	// url
	href, exists := doc.Find("link[rel='canonical']").Attr("href")
	if exists {
		result.URL = href
		baseUrl = helper.ExtractBaseUrl(href)
	}
	data.ProductInformation = make(map[string]interface{})
	data.CustomizationOptions = make(map[string]interface{})
	// aplus_present
	data.AplusPresent = doc.Find("#aplus").Length() > 0
	// availability_status
	availTag := doc.Find("#availability").Find("span").First()
	if availTag.Length() > 0 {
		data.AvailabilityStatus = helper.NormalizeText(availTag.Text())
	}
	// avarage_rating, total_reviews
	reviewTag := doc.Find("#cm_cr_dp_d_rating_histogram").First()
	if reviewTag.Length() > 0 {
		averageReviewTag := reviewTag.Find(".AverageCustomerReviews").First()
		averageReviewText := strings.Split(helper.NormalizeText(averageReviewTag.Find("span").First().Text()), " ")[0]
		rate, err := strconv.ParseFloat(averageReviewText, 64)
		if err == nil {
			data.AverageRating = rate
		}
		totalReviewsText := helper.NormalizeText(reviewTag.Find(".averageStarRatingNumerical").First().Text())
		totalReviewsText = strings.ReplaceAll(strings.Split(totalReviewsText, " ")[0], ",", "")
		num, err := strconv.Atoi(totalReviewsText)
		if err == nil {
			data.TotalReviews = num
		}
	}
	// total_answered_questions ???

	// brand and brand_url
	brandTag := doc.Find("#bylineInfo").First()
	if brandTag.Length() > 0 {
		data.Brand = helper.NormalizeText(brandTag.Text())
		data.BrandURL = helper.NormalizeUrl(baseUrl, brandTag.AttrOr("href", ""))
	}
	// customization_options
	twisterTag := doc.Find("#twister").First()
	if twisterTag.Length() > 0 {
		twisterTag.ChildrenFiltered("div").Each(func(i int, s *goquery.Selection) {
			var values []interface{}
			varrow := helper.NormalizeText(s.Find("div.a-row").First().Text())
			varname := strings.Split(varrow, ":")[0]
			varvalue := strings.Split(varrow, ":")[1]
			ulTag := s.Find("ul").First()
			if ulTag.Length() > 0 {
				ulTag.Find("button").Each(func(i int, btn *goquery.Selection) {
					record := parseBtn(btn, i == 0)
					values = append(values, record)
				})
			} else {
				value := make(map[string]interface{})
				value["value"] = varvalue
				values = append(values, value)
			}
			data.CustomizationOptions[varname] = values
		})

	}

	// feature_bullets
	doc.Find("li.a-spacing-mini").Each(func(i int, s *goquery.Selection) {
		data.FeatureBullets = append(data.FeatureBullets, helper.NormalizeText(s.Text()))
	})

	// full_description
	aplusDesc := doc.Find("#aplus_feature_div").First()
	if aplusDesc.Length() > 0 {
		data.FullDescription = helper.ExtractText(aplusDesc, "\n")
	}
	prodDesc := doc.Find("#productDescription_feature_div").First()
	if prodDesc.Length() > 0 {
		data.FullDescription += helper.ExtractText(prodDesc, "\n")
	}
	// name
	var titleTag = doc.Find("#productTitle").First()
	data.Name = helper.NormalizeText(titleTag.Text())

	// price, price_string, price_currency
	priceTag := doc.Find("#apex_desktop").First()
	if priceTag.Length() > 0 {
		priceText := helper.ExtractPrice(helper.ExtractText(priceTag, " "))
		data.PriceString = priceText
		if len(priceText) > 0 {
			price, err := strconv.ParseFloat(priceText[1:], 64)
			if err == nil {
				data.PriceCurrency = priceText[:1]
				data.Price = price
			} else {
				price, err := strconv.ParseFloat(priceText[2:], 64)
				if err == nil {
					data.PriceCurrency = priceText[:2]
					data.Price = price
				}
			}
		}
	}

	// product_category
	var categoryTag = doc.Find("#wayfinding-breadcrumbs_feature_div").First()
	categoryTag.Find("span.a-list-item").Each(func(i int, s *goquery.Selection) {
		data.ProductCategory += helper.NormalizeText(s.Text())
	})

	// product_information
	detailTag := doc.Find("#prodDetails").First()
	if detailTag.Length() > 0 {
		detailTag.Find("table").Each(func(i int, s *goquery.Selection) {
			detailMap := helper.ConvertTableToMap(s)
			for key, value := range detailMap {
				data.ProductInformation[key] = value
			}
		})
		data.ProductInformation["Customer Reviews"] = map[string]interface{}{
			"ratings_count": data.TotalReviews,
			"stars":         data.AverageRating,
		}
		if v, ok := data.ProductInformation["Best Sellers Rank"].(string); ok {
			data.ProductInformation["Best Sellers Rank"] = strings.Split(v, ") ")
		}
		if v, ok := data.ProductInformation["Item model number"].(string); ok {
			data.Model = v
		}
	}

	buyBoxTag := doc.Find("#buybox").First()
	if buyBoxTag.Length() > 0 {
		corePriceTag := buyBoxTag.Find("#corePrice_feature_div").Find("span.a-offscreen").First()
		if corePriceTag.Length() > 0 {
			data.ShippingPrice = helper.NormalizeText(corePriceTag.Text())
		}
		fulfillerTag := buyBoxTag.Find("#fulfillerInfoFeature_feature_div").Find(".offer-display-feature-text").First()
		if fulfillerTag.Length() > 0 {
			data.ShipsFrom = helper.NormalizeText(fulfillerTag.Text())
		}
		merchantTag := buyBoxTag.Find("#merchantInfoFeature_feature_div").Find(".offer-display-feature-text").First()
		if fulfillerTag.Length() > 0 {
			data.SoldBy = helper.NormalizeText(merchantTag.Text())
		}
	}
	// images
	imageTag := doc.Find("#altImages").First()
	if imageTag.Length() > 0 {
		imageTag.Find("img").Each(func(i int, s *goquery.Selection) {
			imageUrl := helper.NormalizeImage(s.AttrOr("src", ""))
			if imageUrl != "" {
				data.Images = append(data.Images, imageUrl)
			}
		})
	}
	result.Data = data
	result.Status = "parse_successful"
	return result
}
