package main

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type WhProduct struct {
	Position         int     `json:"position"`
	ID               string  `json:"id"`
	ItemID           string  `json:"item_id"`
	Name             string  `json:"name"`
	Type             string  `json:"type"`
	Brand            string  `json:"brand"`
	ShortDescription string  `json:"short_description"`
	AverageRating    float64 `json:"average_rating"`
	NumberOfReviews  int     `json:"number_of_reviews"`
	SalesUnit        string  `json:"sales_unit"`
	SellerName       string  `json:"seller_name"`
	Image            string  `json:"image"`
	ImageSize        string  `json:"image_size"`
	Price            float64 `json:"price"`
	PriceInfo        struct {
		LinePrice  string `json:"line_price"`
		WasPrice   string `json:"was_price"`
		UnitPrice  string `json:"unit_price"`
		ItemPrice  string `json:"item_price"`
		ShipPrice  string `json:"ship_price"`
		PriceRange string `json:"price_range"`
		Savings    string `json:"savings"`
	} `json:"price_info"`
	IsEligible      bool      `json:"is_eligible"`
	IsShowATC       bool      `json:"is_show_atc"`
	IsShowOptions   bool      `json:"is_show_options"`
	IsBuyNow        bool      `json:"is_buy_now"`
	IsSponsored     bool      `json:"is_sponsored"`
	IsOutofStock    bool      `json:"is_outofstock"`
	Availability    string    `json:"availability"`
	ProductLocation string    `json:"product_location"`
	Flag            string    `json:"flag"`
	Fulfillment     []string  `json:"fulfillment"`
	URL             string    `json:"url"`
	Variants        []WStrMap `json:"variants"`
}

type WhrProduct struct {
	TypeName         string `json:"__typename"`
	Type             string `json:"type"`
	Id               string `json:"id,omitempty"`
	ItemId           string `json:"usItemId,omitempty"`
	Name             string `json:"name,omitempty"`
	Brand            string `json:"brand,omitempty"`
	ShortDescription string `json:"shortDescription,omitempty"`
	ImageInfo        struct {
		ThumbnailUrl string `json:"thumbnailUrl"`
		Size         string `json:"size"`
	} `json:"imageInfo,omitempty"`
	CanonicalUrl string `json:"canonicalUrl,omitempty"`
	Badges       struct {
		Flags []struct {
			Key  string `json:"key"`
			Text string `json:"text"`
			Type string `json:"type"`
		} `json:"flags"`
		Tags []struct {
			Key  string `json:"key"`
			Text string `json:"text"`
			Type string `json:"type"`
		} `json:"tags"`
		Groups []struct {
			Name    string `json:"name"`
			Members []struct {
				Text    string `json:"text"`
				SlaText string `json:"slaText"`
				Type    string `json:"memberType"`
				Key     string `json:"key"`
			} `json:"members"`
		} `json:"groups"`
	} `json:"badges,omitempty"`
	SnapEligible    bool    `json:"snapEligible,omitempty"`
	BuyNowEligible  bool    `json:"buyNowEligible,omitempty"`
	AverageRating   float64 `json:"averageRating,omitempty"`
	NumberOfReviews int     `json:"numberOfReviews,omitempty"`
	SalesUnitType   string  `json:"salesUnitType,omitempty"`
	SellerName      string  `json:"sellerName,omitempty"`
	PriceInfo       struct {
		ItemPrice  string `json:"itemPrice"`
		LinePrice  string `json:"linePrice"`
		UnitPrice  string `json:"unitPrice"`
		WasPrice   string `json:"wasPrice"`
		ShipPrice  string `json:"shipPrice"`
		Savings    string `json:"savings"`
		PriceRange string `json:"priceRangeString"`
	} `json:"priceInfo,omitempty"`
	ManufacturerName   string `json:"manufacturerName,omitempty"`
	ShowAtc            bool   `json:"showAtc,omitempty"`
	ShowOptions        bool   `json:"showOptions,omitempty"`
	ShowBuyNow         bool   `json:"showBuyNow,omitempty"`
	AvailabilityStatus string `json:"availabilityStatusDisplayValue,omitempty"`
	ProductLocation    string `json:"productLocationDisplayValue,omitempty"`
	CanAddToCart       bool   `json:"canAddToCart,omitempty"`
	Flag               string `json:"flag,omitempty"`
	Badge              struct {
		Key  string `json:"key"`
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"badge,omitempty"`
	FulfillmentBadgeGroups []struct {
		Text    string `json:"text"`
		SlaText string `json:"slaText"`
	} `json:"fulfillmentBadgeGroups,omitempty"`
	SpecialBuy   bool    `json:"specialBuy,omitempty"`
	PriceFlip    bool    `json:"priceFlip,omitempty"`
	Image        string  `json:"image,omitempty"`
	ImageSize    string  `json:"imageSize,omitempty"`
	IsOutOfStock bool    `json:"isOutOfStock,omitempty"`
	Price        float64 `json:"price,omitempty"`
	Rating       struct {
		AverageRating   float64 `json:"averageRating"`
		NumberOfReviews int     `json:"numberOfReviews"`
	} `json:"rating,omitempty"`
	SalesUnit   string `json:"salesUnit,omitempty"`
	VariantList []struct {
		Name         string `json:"name"`
		Image        string `json:"image"`
		CanonicalUrl string `json:"canonicalUrl"`
	} `json:"variantList"`
	IsVariantTypeSwatch bool `json:"isVariantTypeSwatch,omitempty"`
	IsSponsoredFlag     bool `json:"isSponsoredFlag,omitempty"`
	TileTakeOverTile    struct {
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
		Image    struct {
			Src string `json:"src"`
			Alt string `json:"alt"`
		} `json:"image"`
		TileCta []struct {
			CtaLink struct {
				ClickThrough struct {
					Value string `json:"value"`
				} `json:"clickThrough"`
				LinkText string `json:"linkText"`
				Title    string `json:"title"`
			} `json:"ctaLink"`
			CtaType string `json:"ctaType"`
		} `json:"tileCta"`
		AdsEnabled string `json:"adsEnabled"`
	} `json:"tileTakeOverTile,omitempty"`
	ProductIndex int `json:"productIndex"`
}

type WhrData struct {
	SearchResult struct {
		Title           string `json:"title"`
		AggregatedCount int    `json:"aggregatedCount"`
		BreadCrumb      []struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"breadCrumb"`
		ItemStacks []struct {
			Title string       `json:"title"`
			Count int          `json:"count"`
			Items []WbrProduct `json:"items"`
		} `json:"itemStacks"`
		Pagination struct {
			MaxPage    int `json:"maxPage"`
			Properties struct {
				Page int `json:"page"`
			} `json:"pageProperties"`
		} `json:"paginationV2"`
	} `json:"searchResult"`
	ContentLayout struct {
		Modules []WmrModule `json:"modules"`
	} `json:"contentLayout"`
	SeoMetaData struct {
		MetaDesc  string `json:"metaDesc"`
		MetaTitle string `json:"metaTitle"`
		MetaCanon string `json:"metaCanon"`
	} `json:"seoMetaData"`
}

type WhrResult struct {
	Props struct {
		PageProps struct {
			InitialData WbrData `json:"initialData"`
		} `json:"pageProps"`
	} `json:"props"`
	Page          string            `json:"page"`
	Query         map[string]string `json:"query"`
	RuntimeConfig struct {
		Host struct {
			Wmt string `json:"wmt"`
		} `json:"host"`
	} `json:"runtimeConfig"`
}

type WhProductStack struct {
	Title    string      `json:"title"`
	Count    int         `json:"count"`
	Products []WbProduct `json:"products"`
}

type WhData struct {
	Title         string           `json:"title"`
	Description   string           `json:"description"`
	CopyBlock     string           `json:"copy_block"`
	TotalCount    int              `json:"total_count"`
	BreadCrumbs   []WStrMap        `json:"bread_crumbs"`
	Query         WStrMap          `json:"query"`
	ProductStacks []WbProductStack `json:"product_stacks"`
	Pills         []WStrMap        `json:"pills"`
	Faqs          []WStrMap        `json:"faqs"`
	Populars      []WStrMap        `json:"populars"`
	SkinnyBanners []WmSkinnyBanner `json:"skinny_banners"`
	Pagination    struct {
		PageCount   int      `json:"page_count"`
		CurrentPage int      `json:"current_page"`
		PageLinks   []string `json:"page_links"`
	} `json:"pagination"`
}

type WhResult struct {
	Data   WbData `json:"data"`
	Status string `json:"status"`
	URL    string `json:"url"`
}

func WH_MakeUrl(path string, params WStrMap) string {
	values := url.Values{}
	for key, value := range params {
		if key != "seo" {
			values.Add(key, value)
		}
	}
	// Construct the query string
	queryString := values.Encode()
	if queryString == "" {
		return path
	}
	return path + "?" + queryString
}

func WH_MakePageUrl(mainUrl string, page int) string {
	pageUrl := mainUrl
	if page != 1 {
		if strings.Contains(mainUrl, "?") {
			pageUrl += "&page=" + strconv.Itoa(page)
		} else {
			pageUrl += "?page=" + strconv.Itoa(page)
		}
	}
	return pageUrl
}

func WH_ParseProduct(item WbrProduct, position int, baseUrl string) WbProduct {
	var product WbProduct
	product.Position = position
	product.ID = item.Id
	product.ItemID = item.ItemId
	product.Availability = item.AvailabilityStatus
	product.AverageRating = item.AverageRating
	product.Brand = item.Brand
	product.Flag = item.Flag
	for _, value := range item.FulfillmentBadgeGroups {
		product.Fulfillment = append(product.Fulfillment, value.Text+value.SlaText)
	}
	// product.Image = WH_NormalizeImage(item.Image)
	product.ImageSize = item.ImageSize
	product.IsBuyNow = item.ShowBuyNow
	product.IsEligible = item.SnapEligible
	product.IsOutofStock = item.IsOutOfStock
	product.IsShowATC = item.ShowAtc
	product.IsShowOptions = item.ShowOptions
	product.IsSponsored = item.IsSponsoredFlag
	product.Name = item.Name
	product.NumberOfReviews = item.NumberOfReviews
	product.Price = item.Price
	product.PriceInfo.PriceRange = item.PriceInfo.PriceRange
	product.PriceInfo.ItemPrice = item.PriceInfo.ItemPrice
	product.PriceInfo.LinePrice = item.PriceInfo.LinePrice
	product.PriceInfo.UnitPrice = item.PriceInfo.UnitPrice
	product.PriceInfo.WasPrice = item.PriceInfo.WasPrice
	product.PriceInfo.ShipPrice = item.PriceInfo.ShipPrice
	product.PriceInfo.Savings = item.PriceInfo.Savings
	product.ProductLocation = item.ProductLocation
	product.SalesUnit = item.SalesUnit
	product.SellerName = item.SellerName
	product.ShortDescription = item.ShortDescription
	product.Type = item.Type
	product.URL = baseUrl + item.CanonicalUrl
	for _, elem := range item.VariantList {
		variant := make(WStrMap)
		variant["name"] = elem.Name
		variant["image"] = elem.Image
		variant["url"] = baseUrl + elem.CanonicalUrl
		product.Variants = append(product.Variants, variant)
	}
	return product
}

func Walmart_ShopPageScraper(jsonTag *goquery.Selection) WbResult {
	var raw WbrResult
	var result WbResult
	var data WbData
	baseUrl := "https://www.walmart.com"
	json.Unmarshal([]byte(jsonTag.Text()), &raw)
	baseUrl = raw.RuntimeConfig.Host.Wmt
	rawData := raw.Props.PageProps.InitialData
	mainUrl := WB_MakeUrl(rawData.SeoMetaData.MetaCanon, raw.Query)
	data.Title = rawData.SeoMetaData.MetaTitle
	data.Description = rawData.SeoMetaData.MetaDesc
	data.TotalCount = rawData.SearchResult.AggregatedCount
	for _, item := range rawData.SearchResult.BreadCrumb {
		breadcrumb := make(WStrMap)
		breadcrumb["name"] = item.Name
		breadcrumb["url"] = baseUrl + item.Url
		data.BreadCrumbs = append(data.BreadCrumbs, breadcrumb)
	}
	for _, stack := range rawData.SearchResult.ItemStacks {
		var productStack WbProductStack
		productStack.Title = stack.Title
		productStack.Count = stack.Count
		position := 1
		for _, item := range stack.Items {
			if item.TypeName == "Product" {
				product := WB_ParseProduct(item, position, baseUrl)
				position += 1
				productStack.Products = append(productStack.Products, product)
			}
		}
		data.ProductStacks = append(data.ProductStacks, productStack)
	}
	data.Query = raw.Query
	data.Pagination.CurrentPage = rawData.SearchResult.Pagination.Properties.Page
	data.Pagination.PageCount = rawData.SearchResult.Pagination.MaxPage
	for i := 1; i <= data.Pagination.PageCount; i++ {
		pageUrl := WB_MakePageUrl(mainUrl, i)
		data.Pagination.PageLinks = append(data.Pagination.PageLinks, pageUrl)
	}
	for _, module := range rawData.ContentLayout.Modules {
		if module.Type == "PillsModule" {
			data.Pills = WM_ParsePills(module, baseUrl)
		} else if module.Type == "CopyBlock" {
			data.CopyBlock = WM_ParseCopyBlock(module, baseUrl)
		} else if module.Type == "GenericSEOFAQModule" {
			data.Faqs = WM_ParseFaqs(module, baseUrl)
		} else if module.Type == "PopularInBrowse" {
			data.Populars = WM_ParsePopularInBrowser(module, baseUrl)
		}
	}
	result.Data = data
	result.URL = mainUrl
	result.Status = "parse_successful"
	return result
}

func Walmart_IsShopPage(jsonTag *goquery.Selection) bool {
	var page WRawResult
	json.Unmarshal([]byte(jsonTag.Text()), &page)
	return page.Page != "" && strings.Split(page.Page, "/")[1] == "shop"
}
