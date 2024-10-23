package main

type WmCategory struct {
	Name          string    `json:"name"`
	Image         string    `json:"image"`
	Subcategories []WStrMap `json:"subcategories"`
}

type WmLeftNav struct {
	Title      string       `json:"title"`
	Categories []WmCategory `json:"categories"`
}

type WmHubSpoke struct {
	Title      string    `json:"title"`
	Categories []WStrMap `json:"categories"`
}

type WmPovCard struct {
	Title string    `json:"title"`
	Cards []WStrMap `json:"cards"`
}

type WmSkinnyBanner struct {
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle"`
	Image    string    `json:"image"`
	Links    []WStrMap `json:"links"`
}

type WmProduct struct {
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

type WmrProduct struct {
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

type WmItemCarousel struct {
	Title    string      `json:"title"`
	Subtitle string      `json:"subtitle"`
	Products []WmProduct `json:"products"`
}

type WmInspiration struct {
	Title    string      `json:"title"`
	Subtitle string      `json:"subtitle"`
	Image    string      `json:"image"`
	Link     string      `json:"link"`
	Products []WmProduct `json:"products"`
}

type WmrModule struct {
	Name string `json:"name"`
	// CategoryLeftHandNav, ItemCarousel, HubSpokesNxM, POVCards, GenericCopyBlock, CopyBlock
	// CategoryRelatedShelves, HeroPov, SkinnyBanner, CategoryTopNav, PillsModule, PopularInBrowse
	Type    string `json:"type"`
	Version int    `json:"version"`
	Configs struct {
		// for CategoryLeftHandNav, ItemCarousel
		Title string `json:"title,omitempty"`
		// for ItemCarousel
		SubTitle string `json:"subTitle,omitempty"`
		// for GenericCopyBlock
		CatCopyBlock string `json:"catCopyBlock,omitempty"`
		// for CopyBlock
		CopyBlock struct {
			Cwc string `json:"cwc"`
		} `json:"copyBlock,omitempty"`
		// for SkinnyBanner
		SkinnyBannerHeading struct {
			Title string `json:"title"`
		} `json:"skinnyBannerHeading,omitempty"`
		// for SkinnyBanner
		SubHeading struct {
			Title string `json:"title"`
		} `json:"subHeading,omitempty"`
		// for SkinnyBanner
		BannerImage struct {
			Src string `json:"src"`
		} `json:"bannerImage,omitempty"`
		// for SkinnyBanner
		BannerCta []struct {
			CtaLink struct {
				Title        string `json:"title"`
				ClickThrough struct {
					Value string `json:"value"`
				} `json:"clickThrough"`
			} `json:"ctaLink"`
		} `json:"bannerCta,omitempty"`
		// for HWBreadcrumb
		BreadcrumbList []struct {
			Label string `json:"label"`
			Link  string `json:"link"`
		} `json:"BreadcrumbList,omitempty"`
		// for GenericSEOFAQModule
		SeoFaqList []struct {
			SeoFaqQuestion string `json:"seoFaqQuestion"`
			SeoFaqAnswer   string `json:"seoFaqAnswer"`
		} `json:"seoFaqList,omitempty"`
		// for CategoryTopNav
		NavHeaders []struct {
			CtaLink struct {
				Title        string `json:"title"`
				ClickThrough struct {
					Value string `json:"value"`
				} `json:"clickThrough"`
			} `json:"ctaLink"`
		} `json:"navHeaders,omitempty"`
		// for CategoryTopNav
		NavPills []struct {
			Title string `json:"title"`
			Url   struct {
				ClickThrough struct {
					Value string `json:"value"`
				} `json:"clickThrough"`
			} `json:"url"`
		} `json:"NavPills,omitempty"`
		// for PillsModule
		PillsV2 []struct {
			Title string `json:"title"`
			Url   string `json:"url"`
			Image struct {
				Src string `json:"src"`
			} `json:"image"`
		} `json:"pillsV2,omitempty"`
		// for GenericCopyBlock
		SeoCategoryRelmData struct {
			Relm []struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"relm"`
		} `json:"seoCategoryRelmData,omitempty"`
		// for PopularInBrowse
		SeoBrowseRelmData struct {
			Relm []struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"relm"`
		} `json:"seoBrowseRelmData,omitempty"`
		// for ItemCarousel
		ViewAllLink struct {
			LinkText     string `json:"linkText"`
			ClickThrough struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"clickThrough"`
		} `json:"viewAllLink,omitempty"`
		// for ItemCarousel
		ProductsConfig struct {
			Products []WmrProduct `json:"products"`
		} `json:"productsConfig,omitempty"`
		// for CategoryLeftHandNav
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
		} `json:"categories,omitempty"`
		// for HubSpokesNxM, POVCards, HubSpokes4x1
		HeadingText string `json:"headingText,omitempty"`
		// HubSpokes4x1
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
				Products []WmrProduct `json:"products"`
			} `json:"productsConfig,omitempty"`
		} `json:"inspirationModule"`
		// for HubSpokesNxM
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
		} `json:"rows6,omitempty"`
		// for POVCards
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
		} `json:"cardsV1,omitempty"`
		// for HeroPov
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
		} `json:"povCards,omitempty"`
	} `json:"configs"`
}

func WM_ParsePovCard(module WmrModule, baseUrl string) WmPovCard {
	var povcard WmPovCard
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

func WM_ParseSkinnyBanner(module WmrModule, baseUrl string) WmSkinnyBanner {
	var banner WmSkinnyBanner
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

func WM_ParseFaqs(module WmrModule, baseUrl string) []WStrMap {
	var faqs []WStrMap
	for _, item := range module.Configs.SeoFaqList {
		faq := make(WStrMap)
		faq["question"] = item.SeoFaqQuestion
		faq["answer"] = item.SeoFaqAnswer
		faqs = append(faqs, faq)
	}
	return faqs
}

func WM_ParsePills(module WmrModule, baseUrl string) []WStrMap {
	var pills []WStrMap
	for _, item := range module.Configs.PillsV2 {
		pill := make(WStrMap)
		pill["title"] = item.Title
		pill["image"] = item.Image.Src
		pill["url"] = item.Url
		pills = append(pills, pill)
	}
	return pills
}

func WM_ParsePopularInBrowser(module WmrModule, baseUrl string) []WStrMap {
	var populars []WStrMap
	for _, item := range module.Configs.SeoBrowseRelmData.Relm {
		popular := make(WStrMap)
		popular["name"] = item.Name
		popular["url"] = baseUrl + item.Url
		populars = append(populars, popular)
	}
	return populars
}

func WM_ParseBreadCrumbs(module WmrModule, baseUrl string) []WStrMap {
	var breadcrumbs []WStrMap
	for _, item := range module.Configs.BreadcrumbList {
		breadcrumb := make(WStrMap)
		breadcrumb["label"] = item.Label
		breadcrumb["link"] = item.Link
		breadcrumbs = append(breadcrumbs, breadcrumb)
	}
	return breadcrumbs
}

func WM_ParseNavPills(module WmrModule, baseUrl string) []WStrMap {
	var pills []WStrMap
	for _, item := range module.Configs.NavPills {
		pill := make(WStrMap)
		pill["title"] = item.Title
		pill["url"] = baseUrl + item.Url.ClickThrough.Value
		pills = append(pills, pill)
	}
	return pills
}

func WM_ParseProduct(item WmrProduct, baseUrl string) WmProduct {
	var product WmProduct
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

func WM_ParseItemCarousel(module WmrModule, baseUrl string) WmItemCarousel {
	var carousel WmItemCarousel
	carousel.Title = module.Configs.Title
	carousel.Subtitle = module.Configs.SubTitle
	for _, item := range module.Configs.ProductsConfig.Products {
		product := WM_ParseProduct(item, baseUrl)
		carousel.Products = append(carousel.Products, product)
	}
	return carousel
}

func WM_ParseInspirationModule(module WmrModule, baseUrl string) []WmInspiration {
	var insps []WmInspiration
	for _, item := range module.Configs.InspirationModule {
		var insp WmInspiration
		insp.Title = item.CardTitle
		insp.Subtitle = item.CardSubTitle
		insp.Image = item.CardImage.Src
		insp.Link = item.CardImage.ClickThrough.Value
		for _, elem := range item.ProductsConfig.Products {
			product := WM_ParseProduct(elem, baseUrl)
			insp.Products = append(insp.Products, product)
		}
		insps = append(insps, insp)
	}
	return insps
}

func WM_ParseHubSpokeNxM(module WmrModule, baseUrl string) WmHubSpoke {
	var hubspoke WmHubSpoke
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

func WM_ParseHubSpoke4x1(module WmrModule, baseUrl string) WmHubSpoke {
	var hubspoke WmHubSpoke
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

func WM_ParseRelatedPages(module WmrModule, baseUrl string) []WStrMap {
	var pages []WStrMap
	for _, item := range module.Configs.SeoCategoryRelmData.Relm {
		page := make(WStrMap)
		page["name"] = item.Name
		page["url"] = baseUrl + item.Url
		pages = append(pages, page)
	}
	return pages
}

func WM_ParseCopyBlock(module WmrModule, baseUrl string) string {
	return module.Configs.CopyBlock.Cwc
}

func WM_ParseLeftNav(module WmrModule, baseUrl string) WmLeftNav {
	var nav WmLeftNav
	nav.Title = module.Configs.Title
	for _, item := range module.Configs.Categories {
		var category WmCategory
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

func WM_ParseHeroPov(module WmrModule, baseUrl string) []WStrMap {
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
