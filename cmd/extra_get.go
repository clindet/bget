package cmd

import (
	"errors"
	"fmt"
	neturl "net/url"
	"path"

	"github.com/Miachol/bget/spider"
	"github.com/openbiox/butils/archive"
	log "github.com/openbiox/butils/log"
	cnet "github.com/openbiox/butils/net"
)

// Geofetch get GEO files
func Geofetch(geo string, outDir string, opt *cnet.BnetParams) (err error) {
	if geo == "" {
		return errors.New("at least one of geo is required")
	}
	gseURLs, gplURLs, sraLink := spider.GeoSpider(geo, bgetClis.proxy, opt.Timeout)
	u, _ := neturl.Parse(sraLink)
	uQ := u.Query()
	accAll := fmt.Sprintf(`https://www.ncbi.nlm.nih.gov/Traces/study/backends/solr_proxy/solr_proxy.cgi?core=run_sel_index&action=acc_all&fl=acc_s&rs=(primary_search_ids:"%s")`, uQ["acc"][0])
	rtAll := `https://www.ncbi.nlm.nih.gov/Traces/study/backends/solr_proxy/solr_proxy.cgi?core=run_sel_index&action=rt_all&fl=acc_s%2Cantibody_sam_ss%2Cassay_type_s%2Cavgspotlen_l%2Cbioproject_s%2Cbiosample_s%2Ccell_line_sam_ss_dpl110_ss%2Ccenter_name_s%2Cconsent_s%2Cdatastore_filetype_ss%2Cdatastore_provider_ss%2Cdatastore_region_ss%2Cexperiment_s%2Cgeo_accession_exp_ss%2Cinstrument_s%2Clibrary_name_s%2Clibrarylayout_s%2Clibraryselection_s%2Clibrarysource_s%2Cmbases_l%2Cmbytes_l%2Corganism_s%2Cplatform_s%2Creleasedate_dt%2Csample_acc_s%2Csample_name_s%2Csource_name_sam_ss%2Csra_study_s%2Ctreatment_sam_ss%2Cquality_book_char_run_ss&ft=Run%2CAntibody%2CAssay+Type%2CAvgSpotLen%2CBioProject%2CBioSample%2Ccell_line%2CCenter+Name%2CConsent%2CDATASTORE+filetype%2CDATASTORE+provider%2CDATASTORE+region%2CExperiment%2CGEO_Accession%2CInstrument%2CLibrary+Name%2CLibraryLayout%2CLibrarySelection%2CLibrarySource%2CMBases%2CMBytes%2COrganism%2CPlatform%2CReleaseDate%2Csample_acc%2CSample+Name%2Csource_name%2CSRA+Study%2Ctreatment%2Cquality_book_char&rs=%28primary_search_ids%3A%22` + uQ["acc"][0] + `%22%29`
	sraLinks := []string{accAll, rtAll}
	urls := append(gseURLs, gplURLs...)
	urls = append(urls, sraLinks...)
	destDirArray := []string{}
	for range urls {
		destDirArray = append(destDirArray, outDir)
	}
	done := cnet.HttpGetURLs(urls, destDirArray, opt)
	for _, dest := range done {
		if bgetClis.uncompress {
			if err := archive.UnarchiveLog(dest, path.Dir(dest)); err != nil {
				log.Warn(err)
			}
		}
	}
	return nil
}
