package main

import (
	"encoding/json"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type WcCategory struct {
	Name          string    `json:"name"`
	Image         string    `json:"image"`
	Subcategories []WStrMap `json:"subcategories"`
}

type WcLeftNav struct {
	Title      string       `json:"title"`
	Categories []WcCategory `json:"categories"`
}

type WcHubSpoke struct {
	Title      string    `json:"title"`
	Categories []WStrMap `json:"categories"`
}

type WcPovCard struct {
	Title string    `json:"title"`
	Cards []WStrMap `json:"cards"`
}

type WcSkinnyBanner struct {
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle"`
	Image    string    `json:"image"`
	Links    []WStrMap `json:"links"`
}

type WcItemCarousel struct {
	Title    string      `json:"title"`
	Subtitle string      `json:"subtitle"`
	Products []WcProduct `json:"products"`
}

type WcInspirationModule struct {
	Title    string      `json:"title"`
	Subtitle string      `json:"subtitle"`
	Image    string      `json:"image"`
	Link     string      `json:"link"`
	Products []WcProduct `json:"products"`
}

type WcData struct {
	Title              string                `json:"title"`
	Desc               string                `json:"description"`
	CopyBlock          string                `json:"copy_block"`
	LeftNavs           []WcLeftNav           `json:"left_navs"`
	HubSpokes          []WcHubSpoke          `json:"hub_spokes"`
	PovCards           []WcPovCard           `json:"pov_cards"`
	RelatedPages       []WStrMap             `json:"related_pages"`
	HeroPovs           []WStrMap             `json:"hero_povs"`
	ItemCarousels      []WcItemCarousel      `json:"item_carousels"`
	SkinnyBanners      []WcSkinnyBanner      `json:"skinny_banners"`
	Faqs               []WStrMap             `json:"faqs"`
	NavPills           []WStrMap             `json:"nav_pills"`
	InspirationModules []WcInspirationModule `json:"inspiration_modules"`
}

type WcResult struct {
	Data   WcData `json:"data"`
	Status string `json:"status"`
	URL    string `json:"url"`
}

type WcProduct struct {
	Name               string  `json:"name"`
	Image              string  `json:"image"`
	Brand              string  `json:"brand"`
	Url                string  `json:"url"`
	AverageRating      float64 `json:"average_rating"`
	NumberOfReviews    int     `json:"total_reviews"`
	Type               string  `json:"type"`
	ItemId             string  `json:"item_id"`
	AvailabilityStatus string  `json:"availability_status"`
	CanAddToCart       bool    `json:"can_add_to_cart"`
	ShowOptions        bool    `json:"show_options"`
	IsOutOfStock       bool    `json:"is_out_of_stock"`
	IsSponsoredFlag    bool    `json:"is_sponsored_flag"`
	SellerName         string  `json:"seller_name"`
	SalesUnit          string  `json:"sales_unit"`
	Price              float64 `json:"price"`
	PriceInfo          struct {
		ItemPrice  string `json:"item_price"`
		LinePrice  string `json:"line_price"`
		UnitPrice  string `json:"unit_price"`
		WasPrice   string `json:"was_price"`
		ShipPrice  string `json:"ship_price"`
		PriceRange string `json:"price_range"`
		Savings    string `json:"savings"`
	} `json:"price_info"`
	Flags        []string  `json:"flags"`
	Labels       []string  `json:"labels"`
	Tags         []string  `json:"tags"`
	Fulfillments []string  `json:"fulfillments"`
	Variants     []WStrMap `json:"variants"`
}

type WcrProduct struct {
	CanAddToCart       bool    `json:"canAddToCart"`
	ShowOptions        bool    `json:"showOptions"`
	Flag               string  `json:"flag"`
	Image              string  `json:"image"`
	AvailabilityStatus string  `json:"availabilityStatus"`
	IsOutOfStock       bool    `json:"isOutOfStock"`
	Name               string  `json:"name"`
	SellerName         string  `json:"sellerName"`
	SalesUnitType      string  `json:"salesUnitType"`
	Price              float64 `json:"price"`
	PriceInfo          struct {
		ItemPrice        string `json:"itemPrice"`
		LinePrice        string `json:"linePrice"`
		UnitPrice        string `json:"unitPrice"`
		WasPrice         string `json:"wasPrice"`
		ShipPrice        string `json:"shipPrice"`
		Savings          string `json:"savings"`
		PriceRangeString string `json:"priceRangeString"`
	} `json:"priceInfo"`
	Badges struct {
		Flags []struct {
			Text string `json:"text"`
			Type string `json:"type"`
			Key  string `json:"key"`
		} `json:"flags"`
		Labels []struct {
			Text string `json:"text"`
			Type string `json:"type"`
			Key  string `json:"key"`
		} `json:"labels"`
		Tags []struct {
			Text string `json:"text"`
			Key  string `json:"key"`
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
	} `json:"badges"`
	VariantList []struct {
		Name         string `json:"name"`
		Image        string `json:"image"`
		CanonicalUrl string `json:"canonicalUrl"`
	} `json:"variantList"`
	Type   string `json:"type"`
	ItemId string `json:"usItemId"`
	Rating struct {
		AverageRating   float64 `json:"averageRating"`
		NumberOfReviews int     `json:"numberOfReviews"`
	} `json:"rating"`
	IsSponsoredFlag bool   `json:"isSponsoredFlag"`
	Brand           string `json:"brand"`
	CanonicalUrl    string `json:"canonicalUrl"`
}

type WcrModule struct {
	Name    string `json:"name"`
	Type    string `json:"type"` // CategoryLeftHandNav, ItemCarousel, HubSpokesNxM, POVCards, GenericCopyBlock, CategoryRelatedShelves, HeroPov, SkinnyBanner, CategoryTopNav
	Version int    `json:"version"`
	// for CategoryLeftHandNav
	Configs struct {
		Title               string `json:"title,omitempty"`        // for CategoryLeftHandNav, ItemCarousel,
		SubTitle            string `json:"subTitle,omitempty"`     // for ItemCarousel
		CatCopyBlock        string `json:"catCopyBlock,omitempty"` // for GenericCopyBlock
		SkinnyBannerHeading struct {
			Title string `json:"title"`
		} `json:"skinnyBannerHeading,omitempty"` // for SkinnyBanner
		SubHeading struct {
			Title string `json:"title"`
		} `json:"subHeading,omitempty"` // for SkinnyBanner
		BannerImage struct {
			Src string `json:"src"`
		} `json:"bannerImage,omitempty"` // for SkinnyBanner
		BannerCta []struct {
			CtaLink struct {
				Title        string `json:"title"`
				ClickThrough struct {
					Value string `json:"value"`
				} `json:"clickThrough"`
			} `json:"ctaLink"`
		} `json:"bannerCta,omitempty"` // for SkinnyBanner
		SeoFaqList []struct {
			SeoFaqQuestion string `json:"seoFaqQuestion"`
			SeoFaqAnswer   string `json:"seoFaqAnswer"`
		} `json:"seoFaqList,omitempty"` // for GenericSEOFAQModule
		NavHeaders []struct {
			CtaLink struct {
				Title        string `json:"title"`
				ClickThrough struct {
					Value string `json:"value"`
				} `json:"clickThrough"`
			} `json:"ctaLink"`
		} `json:"navHeaders,omitempty"` // for CategoryTopNav
		NavPills []struct {
			Title string `json:"title"`
			Url   struct {
				ClickThrough struct {
					Value string `json:"value"`
				} `json:"clickThrough"`
			} `json:"url"`
		} `json:"NavPills,omitempty"` // for CategoryTopNav
		SeoCategoryRelmData struct {
			Relm []struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"relm"`
		} `json:"seoCategoryRelmData,omitempty"` // for GenericCopyBlock
		ViewAllLink struct {
			LinkText     string `json:"linkText"`
			ClickThrough struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"clickThrough"`
		} `json:"viewAllLink,omitempty"` // for ItemCarousel
		ProductsConfig struct {
			Products []WcrProduct `json:"products"`
		} `json:"productsConfig,omitempty"` // for ItemCarousel
		Categories []struct {
			Name          string `json:"name"`
			Image         string `json:"image"`
			Subcategories []struct {
				SubCategoryLink struct {
					Title        string `json:"title"`
					ClickThrough struct {
						Type  string `json:"type"`
						Value string `json:"value"`
					} `json:"clickThrough"`
				} `json:"subCategoryLink"`
			} `json:"subcategories"`
		} `json:"categories,omitempty"` // for CategoryLeftHandNav
		HeadingText  string `json:"headingText,omitempty"` // for HubSpokesNxM, POVCards, HubSpokes4x1
		Categories41 []struct {
			Name  string `json:"name"`
			Image struct {
				Src          string `json:"src"`
				ClickThrough struct {
					Value string `json:"value"`
				} `json:"clickThrough"`
			} `json:"image"`
		} `json:"categories4x1"`
		InspirationModule []struct {
			CardTitle    string `json:"cardTitle"`
			CardSubTitle string `json:"cardSubTitle"`
			CardImage    struct {
				Src          string `json:"src"`
				ClickThrough struct {
					Value string `json:"value"`
				} `json:"clickThrough"`
			} `json:"cardImage"`
			ProductsConfig struct {
				Products []WcrProduct `json:"products"`
			} `json:"productsConfig,omitempty"` // for ItemCarousel
		} `json:"inspirationModule"`
		Rows6 []struct {
			Categories []struct {
				Name  string `json:"name"`
				Image struct {
					Src          string `json:"src"`
					ClickThrough struct {
						Value string `json:"value"`
					} `json:"clickThrough"`
				} `json:"image"`
			} `json:"categories"`
		} `json:"rows6,omitempty"` // for HubSpokesNxM
		CardsV1 []struct {
			HeadingText     string `json:"headingText"`
			DescriptionText string `json:"descriptionText"`
			Image           struct {
				Src string `json:"src"`
			} `json:"image"`
			Link struct {
				LinkText     string `json:"linkText"`
				ClickThrough struct {
					Value string `json:"value"`
				} `json:"clickThrough"`
			} `json:"link"`
		} `json:"cardsV1,omitempty"` // for POVCards
		PovCards []struct {
			Image struct {
				Url string `json:"url"`
			} `json:"image"`
			Heading struct {
				Text string `json:"text"`
			} `json:"heading"`
			SubHeading struct {
				Text string `json:"text"`
			} `json:"subHeading"`
			ButtonCTA struct {
				Text string `json:"text"`
				Link string `json:"link"`
			} `json:"buttonCTA"`
		} `json:"povCards,omitempty"` // for HeroPov
	} `json:"configs"`
}

type WcrData struct {
	ContentLayout struct {
		Modules []WcrModule `json:"modules"`
	} `json:"contentLayout"`
	SeoCategoryMetaData struct {
		MetaDesc  string `json:"metaDesc"`
		MetaTitle string `json:"metaTitle"`
		MetaCanon string `json:"metaCanon"`
	} `json:"seoCategoryMetaData"`
}

type WcrResult struct {
	Props struct {
		PageProps struct {
			InitialTempoData struct {
				Data WcrData `json:"data"`
			} `json:"initialTempoData"`
		} `json:"pageProps"`
	} `json:"props"`
	Page  string `json:"page"`
	Query struct {
		ItemParams []string `json:"itemParams"`
	} `json:"query"`
	RuntimeConfig struct {
		Host struct {
			Wmt string `json:"wmt"`
		} `json:"host"`
	} `json:"runtimeConfig"`
}

func WC_ParsePovCard(module WcrModule, baseUrl string) WcPovCard {
	var povcard WcPovCard
	povcard.Title = module.Configs.HeadingText
	for _, item := range module.Configs.CardsV1 {
		card := make(WStrMap)
		card["name"] = item.HeadingText
		card["description"] = item.DescriptionText
		card["image"] = item.Image.Src
		card["url"] = item.Link.ClickThrough.Value
		card["link_text"] = item.Link.LinkText
		povcard.Cards = append(povcard.Cards, card)
	}
	return povcard
}

func WC_ParseSkinnyBanner(module WcrModule, baseUrl string) WcSkinnyBanner {
	var banner WcSkinnyBanner
	banner.Title = module.Configs.SkinnyBannerHeading.Title
	banner.Subtitle = module.Configs.SubHeading.Title
	banner.Image = module.Configs.BannerImage.Src
	for _, item := range module.Configs.BannerCta {
		link := make(WStrMap)
		link["url"] = item.CtaLink.ClickThrough.Value
		link["link_text"] = item.CtaLink.Title
		banner.Links = append(banner.Links, link)
	}
	return banner
}

func WC_ParseFaqs(module WcrModule, baseUrl string) []WStrMap {
	var faqs []WStrMap
	for _, item := range module.Configs.SeoFaqList {
		faq := make(WStrMap)
		faq["question"] = item.SeoFaqQuestion
		faq["answer"] = item.SeoFaqAnswer
		faqs = append(faqs, faq)
	}
	return faqs
}

func WC_ParseNavPills(module WcrModule, baseUrl string) []WStrMap {
	var pills []WStrMap
	for _, item := range module.Configs.NavPills {
		pill := make(WStrMap)
		pill["title"] = item.Title
		pill["url"] = baseUrl + item.Url.ClickThrough.Value
		pills = append(pills, pill)
	}
	return pills
}

func WC_ParseProduct(item WcrProduct, baseUrl string) WcProduct {
	var product WcProduct
	product.AverageRating = item.Rating.AverageRating
	product.AvailabilityStatus = item.AvailabilityStatus
	product.Brand = item.Brand
	product.CanAddToCart = item.CanAddToCart
	for _, elem := range item.Badges.Flags {
		product.Flags = append(product.Flags, elem.Text)
	}
	for _, elem := range item.Badges.Labels {
		product.Labels = append(product.Labels, elem.Text)
	}
	for _, elem := range item.Badges.Tags {
		product.Tags = append(product.Tags, elem.Text)
	}
	for _, group := range item.Badges.Groups {
		if group.Name == "fulfillment" {
			for _, member := range group.Members {
				product.Fulfillments = append(product.Fulfillments, member.Text+member.SlaText)
			}
		}
	}
	product.Image = item.Image
	product.IsOutOfStock = item.IsOutOfStock
	product.IsSponsoredFlag = item.IsSponsoredFlag
	product.ItemId = item.ItemId
	product.Name = item.Name
	product.NumberOfReviews = item.Rating.NumberOfReviews
	product.Price = item.Price
	product.PriceInfo.ItemPrice = item.PriceInfo.ItemPrice
	product.PriceInfo.LinePrice = item.PriceInfo.LinePrice
	product.PriceInfo.PriceRange = item.PriceInfo.PriceRangeString
	product.PriceInfo.Savings = item.PriceInfo.Savings
	product.PriceInfo.ShipPrice = item.PriceInfo.ShipPrice
	product.PriceInfo.UnitPrice = item.PriceInfo.UnitPrice
	product.PriceInfo.WasPrice = item.PriceInfo.WasPrice
	product.SalesUnit = item.SalesUnitType
	product.SellerName = item.SellerName
	product.ShowOptions = item.ShowOptions
	product.Type = item.Type
	product.Url = baseUrl + item.CanonicalUrl
	for _, elem := range item.VariantList {
		variant := make(WStrMap)
		variant["name"] = elem.Name
		variant["image"] = elem.Image
		product.Variants = append(product.Variants, variant)
	}
	return product
}

func WC_ParseItemCarousel(module WcrModule, baseUrl string) WcItemCarousel {
	var carousel WcItemCarousel
	carousel.Title = module.Configs.Title
	carousel.Subtitle = module.Configs.SubTitle
	for _, item := range module.Configs.ProductsConfig.Products {
		product := WC_ParseProduct(item, baseUrl)
		carousel.Products = append(carousel.Products, product)
	}
	return carousel
}

func WC_ParseInspirationModule(module WcrModule, baseUrl string) []WcInspirationModule {
	var insps []WcInspirationModule
	for _, item := range module.Configs.InspirationModule {
		var insp WcInspirationModule
		insp.Title = item.CardTitle
		insp.Subtitle = item.CardSubTitle
		insp.Image = item.CardImage.Src
		insp.Link = item.CardImage.ClickThrough.Value
		for _, elem := range item.ProductsConfig.Products {
			product := WC_ParseProduct(elem, baseUrl)
			insp.Products = append(insp.Products, product)
		}
		insps = append(insps, insp)
	}
	return insps
}

func WC_ParseHubSpokeNxM(module WcrModule, baseUrl string) WcHubSpoke {
	var hubspoke WcHubSpoke
	hubspoke.Title = module.Configs.HeadingText
	for _, row := range module.Configs.Rows6 {
		for _, item := range row.Categories {
			cat := make(WStrMap)
			cat["name"] = item.Name
			cat["image"] = item.Image.Src
			cat["url"] = baseUrl + item.Image.ClickThrough.Value
			hubspoke.Categories = append(hubspoke.Categories, cat)
		}
	}
	return hubspoke
}

func WC_ParseHubSpoke4x1(module WcrModule, baseUrl string) WcHubSpoke {
	var hubspoke WcHubSpoke
	hubspoke.Title = module.Configs.HeadingText
	for _, item := range module.Configs.Categories41 {
		cat := make(WStrMap)
		cat["name"] = item.Name
		cat["image"] = item.Image.Src
		cat["url"] = baseUrl + item.Image.ClickThrough.Value
		hubspoke.Categories = append(hubspoke.Categories, cat)
	}
	return hubspoke
}

func WC_ParseRelatedPages(module WcrModule, baseUrl string) []WStrMap {
	var pages []WStrMap
	for _, item := range module.Configs.SeoCategoryRelmData.Relm {
		page := make(WStrMap)
		page["name"] = item.Name
		page["url"] = baseUrl + item.Url
		pages = append(pages, page)
	}
	return pages
}

func WC_ParseLeftNav(module WcrModule, baseUrl string) WcLeftNav {
	var nav WcLeftNav
	nav.Title = module.Configs.Title
	for _, item := range module.Configs.Categories {
		var category WcCategory
		category.Name = item.Name
		category.Image = item.Image
		for _, subitem := range item.Subcategories {
			subcat := make(WStrMap)
			subcat["title"] = subitem.SubCategoryLink.Title
			subcat["url"] = baseUrl + subitem.SubCategoryLink.ClickThrough.Value
			category.Subcategories = append(category.Subcategories, subcat)
		}
		nav.Categories = append(nav.Categories, category)
	}
	return nav
}

func WC_ParseHeroPov(module WcrModule, baseUrl string) []WStrMap {
	var povs []WStrMap
	println(module.Type, len(module.Configs.PovCards))
	for _, item := range module.Configs.PovCards {
		pov := make(WStrMap)
		pov["image"] = item.Image.Url
		pov["title"] = item.Heading.Text
		pov["description"] = item.SubHeading.Text
		pov["link_text"] = item.ButtonCTA.Text
		pov["link"] = baseUrl + item.ButtonCTA.Link
		povs = append(povs, pov)
	}
	return povs
}

func Walmart_CategoryPageScraper(jsonTag *goquery.Selection) WcResult {
	var raw WcrResult
	var result WcResult
	var data WcData
	baseUrl := "https://www.walmart.com"
	json.Unmarshal([]byte(jsonTag.Text()), &raw)
	baseUrl = raw.RuntimeConfig.Host.Wmt
	result.Status = "parse_successful"
	rawData := raw.Props.PageProps.InitialTempoData.Data
	data.Desc = rawData.SeoCategoryMetaData.MetaDesc
	data.Title = rawData.SeoCategoryMetaData.MetaTitle
	for _, module := range rawData.ContentLayout.Modules {
		if module.Type == "CategoryLeftHandNav" {
			leftNav := WC_ParseLeftNav(module, baseUrl)
			data.LeftNavs = append(data.LeftNavs, leftNav)
		} else if module.Type == "HubSpokesNxM" {
			hubspoke := WC_ParseHubSpokeNxM(module, baseUrl)
			data.HubSpokes = append(data.HubSpokes, hubspoke)
		} else if module.Type == "HubSpokes4x1" {
			hubspoke := WC_ParseHubSpoke4x1(module, baseUrl)
			data.HubSpokes = append(data.HubSpokes, hubspoke)
		} else if module.Type == "POVCards" {
			povCard := WC_ParsePovCard(module, baseUrl)
			data.PovCards = append(data.PovCards, povCard)
		} else if module.Type == "GenericCopyBlock" {
			data.CopyBlock = module.Configs.CatCopyBlock
		} else if module.Type == "StaticNavigationPills" {
			data.NavPills = WC_ParseNavPills(module, baseUrl)
		} else if module.Type == "CategoryRelatedShelves" {
			data.RelatedPages = WC_ParseRelatedPages(module, baseUrl)
		} else if module.Type == "InspirationModule" {
			insps := WC_ParseInspirationModule(module, baseUrl)
			data.InspirationModules = append(data.InspirationModules, insps...)
		} else if module.Type == "HeroPov" {
			data.HeroPovs = WC_ParseHeroPov(module, baseUrl)
		} else if module.Type == "GenericSEOFAQModule" {
			data.Faqs = WC_ParseFaqs(module, baseUrl)
		} else if module.Type == "ItemCarousel" {
			carousel := WC_ParseItemCarousel(module, baseUrl)
			data.ItemCarousels = append(data.ItemCarousels, carousel)
		} else if module.Type == "SkinnyBanner" {
			banner := WC_ParseSkinnyBanner(module, baseUrl)
			data.SkinnyBanners = append(data.SkinnyBanners, banner)
		}
	}
	result.Data = data
	result.Status = "parse_successful"
	println(rawData.SeoCategoryMetaData.MetaCanon)
	result.URL = rawData.SeoCategoryMetaData.MetaCanon
	return result
}

func Walmart_IsCategoryPage(jsonTag *goquery.Selection) bool {
	var page WRawResult
	json.Unmarshal([]byte(jsonTag.Text()), &page)
	return page.Page != "" && strings.Split(page.Page, "/")[1] == "content"
}
