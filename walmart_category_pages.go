package main

import (
	"encoding/json"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type WcData struct {
	Title              string           `json:"title"`
	Desc               string           `json:"description"`
	CopyBlock          string           `json:"copy_block"`
	LeftNavs           []WmLeftNav      `json:"left_navs"`
	HubSpokes          []WmHubSpoke     `json:"hub_spokes"`
	PovCards           []WmPovCard      `json:"pov_cards"`
	RelatedPages       []WStrMap        `json:"related_pages"`
	HeroPovs           []WStrMap        `json:"hero_povs"`
	ItemCarousels      []WmItemCarousel `json:"item_carousels"`
	SkinnyBanners      []WmSkinnyBanner `json:"skinny_banners"`
	Faqs               []WStrMap        `json:"faqs"`
	NavPills           []WStrMap        `json:"nav_pills"`
	InspirationModules []WmInspiration  `json:"inspiration_modules"`
	BreadCrumbs        []WStrMap        `json:"bread_crumbs"`
}

type WcResult struct {
	Data   WcData `json:"data"`
	Status string `json:"status"`
	URL    string `json:"url"`
}

type WcrData struct {
	ContentLayout struct {
		Modules []WmrModule `json:"modules"`
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
			leftNav := WM_ParseLeftNav(module, baseUrl)
			data.LeftNavs = append(data.LeftNavs, leftNav)
		} else if module.Type == "HubSpokesNxM" {
			hubspoke := WM_ParseHubSpokeNxM(module, baseUrl)
			data.HubSpokes = append(data.HubSpokes, hubspoke)
		} else if module.Type == "HubSpokes4x1" {
			hubspoke := WM_ParseHubSpoke4x1(module, baseUrl)
			data.HubSpokes = append(data.HubSpokes, hubspoke)
		} else if module.Type == "POVCards" {
			povCard := WM_ParsePovCard(module, baseUrl)
			data.PovCards = append(data.PovCards, povCard)
		} else if module.Type == "GenericCopyBlock" {
			data.CopyBlock = module.Configs.CatCopyBlock
		} else if module.Type == "StaticNavigationPills" {
			data.NavPills = WM_ParseNavPills(module, baseUrl)
		} else if module.Type == "HWBreadcrumb" {
			data.BreadCrumbs = WM_ParseBreadCrumbs(module, baseUrl)
		} else if module.Type == "CategoryRelatedShelves" {
			data.RelatedPages = WM_ParseRelatedPages(module, baseUrl)
		} else if module.Type == "InspirationModule" {
			insps := WM_ParseInspirationModule(module, baseUrl)
			data.InspirationModules = append(data.InspirationModules, insps...)
		} else if module.Type == "HeroPov" {
			data.HeroPovs = WM_ParseHeroPov(module, baseUrl)
		} else if module.Type == "GenericSEOFAQModule" {
			data.Faqs = WM_ParseFaqs(module, baseUrl)
		} else if module.Type == "ItemCarousel" {
			carousel := WM_ParseItemCarousel(module, baseUrl)
			data.ItemCarousels = append(data.ItemCarousels, carousel)
		} else if module.Type == "SkinnyBanner" {
			banner := WM_ParseSkinnyBanner(module, baseUrl)
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
