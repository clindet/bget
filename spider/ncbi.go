package spider

import (
	"errors"
	"fmt"
	neturl "net/url"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/openbiox/ligo/archive"
	cnet "github.com/openbiox/ligo/net"
	"github.com/openbiox/ligo/stringo"
)

// Geofetch get GEO files
func Geofetch(geo string, outDir string, gpl bool,
	uncompress bool, opt *cnet.Params) (err error) {
	if geo == "" {
		return errors.New("at least one of geo is required")
	}
	var spiderOpt = QuerySpiderOpt{
		Query:   geo,
		Proxy:   opt.Proxy,
		Timeout: opt.Timeout,
	}
	urls, gplURLs, sraLink := GeoSpider(&spiderOpt, gpl)
	if gpl {
		urls = append(urls, gplURLs...)
	}
	if sraLink != "" {
		u, _ := neturl.Parse(sraLink)
		uQ := u.Query()
		accAll := fmt.Sprintf(`https://www.ncbi.nlm.nih.gov/Traces/study/backends/solr_proxy/solr_proxy.cgi?core=run_sel_index&action=acc_all&fl=acc_s&rs=(primary_search_ids:"%s")`, uQ["acc"][0])
		rtAll := `https://www.ncbi.nlm.nih.gov/Traces/study/backends/solr_proxy/solr_proxy.cgi?core=run_sel_index&action=rt_all&fl=acc_s%2Cantibody_sam_ss%2Cassay_type_s%2Cavgspotlen_l%2Cbioproject_s%2Cbiosample_s%2Ccell_line_sam_ss_dpl110_ss%2Ccenter_name_s%2Cconsent_s%2Cdatastore_filetype_ss%2Cdatastore_provider_ss%2Cdatastore_region_ss%2Cexperiment_s%2Cgeo_accession_exp_ss%2Cinstrument_s%2Clibrary_name_s%2Clibrarylayout_s%2Clibraryselection_s%2Clibrarysource_s%2Cmbases_l%2Cmbytes_l%2Corganism_s%2Cplatform_s%2Creleasedate_dt%2Csample_acc_s%2Csample_name_s%2Csource_name_sam_ss%2Csra_study_s%2Ctreatment_sam_ss%2Cquality_book_char_run_ss&ft=Run%2CAntibody%2CAssay+Type%2CAvgSpotLen%2CBioProject%2CBioSample%2Ccell_line%2CCenter+Name%2CConsent%2CDATASTORE+filetype%2CDATASTORE+provider%2CDATASTORE+region%2CExperiment%2CGEO_Accession%2CInstrument%2CLibrary+Name%2CLibraryLayout%2CLibrarySelection%2CLibrarySource%2CMBases%2CMBytes%2COrganism%2CPlatform%2CReleaseDate%2Csample_acc%2CSample+Name%2Csource_name%2CSRA+Study%2Ctreatment%2Cquality_book_char&rs=%28primary_search_ids%3A%22` + uQ["acc"][0] + `%22%29`
		sraLinks := []string{accAll, rtAll}
		urls = append(urls, sraLinks...)
	}
	destDirArray := []string{}
	for range urls {
		destDirArray = append(destDirArray, outDir)
	}
	done := cnet.HTTPGetURLs(urls, destDirArray, opt)
	for _, dest := range done {
		if uncompress {
			if err := archive.UnarchiveLog(dest, path.Dir(dest)); err != nil {
				log.Warn(err)
			}
		}
	}
	return nil
}

// GeoSpider access https://www.ncbi.nlm.nih.gov/geo files via spider
func GeoSpider(opt *QuerySpiderOpt, gpl bool) (gseURLs []string, gplURLs []string, sraLink string) {
	c := initQueryColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.ncbi.nlm.nih.gov"}...)
	c.OnHTML("table td a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "/geo/download/?acc=GS") {
			gseURLs = append(gseURLs, "https://www.ncbi.nlm.nih.gov"+link)
		} else if gpl && strings.Contains(link, "/geo/download/?acc=GPL") {
			gplURLs = append(gplURLs, "https://www.ncbi.nlm.nih.gov"+link)
		}
	})
	if gpl {
		c.OnHTML("input[name=fulltable]", func(e *colly.HTMLElement) {
			link := e.Attr("onclick")
			if strings.Contains(link, "OpenLink") {
				link = "https://www.ncbi.nlm.nih.gov" + stringo.StrReplaceAll(link, "(OpenLink[(])|(')", "")
				link = stringo.StrReplaceAll(link, ",.*", "")
				gplURLs = append(gplURLs, link)
			}
		})
		c.OnHTML("table td a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "geo/query/acc.cgi?acc=GPL") && !strings.Contains(link, "targ=self") {
				Visit(c, "https://www.ncbi.nlm.nih.gov"+link)
			}
		})
	}
	c.OnHTML("tr td a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "/Traces/study/?acc=") {
			link = "https://www.ncbi.nlm.nih.gov" + e.Attr("href")
			sraLink = link
		}
	})
	Visit(c, fmt.Sprintf("https://www.ncbi.nlm.nih.gov/geo/query/acc.cgi?acc=%s", opt.Query))
	return gseURLs, gplURLs, sraLink
}
func PmcSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.ncbi.nlm.nih.gov"}...)
	if opt.FullText {
		c.OnHTML(".links a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "pdf") {
				link = "https://www.ncbi.nlm.nih.gov" + link
				urls = append(urls, link)
			}
		})
		c.OnHTML(fmt.Sprintf(".doi b:contains('%s')", opt.Doi), func(e *colly.HTMLElement) {
			e.DOM.Parents().Filter(".rslt").Find(".aux .links a.view").Each(func(i int, s *goquery.Selection) {
				link, _ := s.Attr("href")
				if strings.Contains(link, "pdf") {
					link = "https://www.ncbi.nlm.nih.gov" + link
					urls = append(urls, link)
				}
			})
		})
	}
	Visit(c, fmt.Sprintf("https://www.ncbi.nlm.nih.gov/pmc/?term=%s", opt.Doi))
	return urls
}
