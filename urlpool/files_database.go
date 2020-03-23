package urlpool

var annovarLinks = []bgetFilesURLType{
	{
		Name: "db/annovar",
		URL: []string{"http://www.openbioinformatics.org/annovar/download/{{builder}}_{{version}}.txt.gz",
			"http://www.openbioinformatics.org/annovar/download/{{builder}}_{{version}}.txt.idx.gz"},
		Versions: []string{"clinvar_20131105", "clinvar_20140211", "clinvar_20140303",
			"clinvar_20140702", "clinvar_20140902", "clinvar_20140929", "clinvar_20150330",
			"clinvar_20150629", "clinvar_20151201", "clinvar_20160302", "clinvar_20161128",
			"clinvar_20170130", "clinvar_20170501", "clinvar_20170905", "clinvar_20180603",
			"avsnp150", "avsnp147", "avsnp144", "avsnp142", "avsnp138",
			"cadd", "caddgt10", "caddgt20", "cadd13", "cadd13gt10", "cadd13gt20", "cg69", "cg46",
			"cosmic70", "cosmic68wgs", "cosmic68", "cosmic67wgs", "cosmic67", "cosmic65", "cosmic64",
			"dbnsfp35a", "dbnsfp33a", "dbnsfp31a_interpro", "dbnsfp30a", "dbscsnv11", "eigen",
			"esp6500siv2_ea", "esp6500siv2_aa", "esp6500siv2_all", "exac03nontcga", "exac03nonpsych",
			"exac03", "fathmm", "gerp++elem", "gerp++gt2", "gme", "gnomad_exome", "gnomad_genome",
			"gnomad30_genome", "gwava", "hrcr1", "icgc21", "intervar_20170202",
			"kaviar_20150923", "ljb26_all", "mcap", "mitimpact2", "mitimpact24", "nci60",
			"popfreq_max_20150413", "popfreq_all_20150413", "revel", "regsnpintron",
			"gene4denovo201907"},
	},
	{
		Name:     "db/annovar-1000g",
		URL:      []string{"http://www.openbioinformatics.org/annovar/download/{{builder}}_{{version}}.zip"},
		Versions: []string{"1000g2015aug", "1000g2014oct", "1000g2014sep", "1000g2014aug", "1000g2012apr", "1000g2012feb", "1000g2011may", "1000g2010nov", "1000g2012apr", "1000g2010jul", "1000g2010", "1000g"},
	},
	{
		Name: "db/annovar-noidx",
		URL: []string{"http://www.openbioinformatics.org/annovar/download/{{version}}.txt.gz",
			"http://www.openbioinformatics.org/annovar/download/{{builder}}_{{version}}.txt.gz"},
		Versions: []string{"GDI_full_10282015", "RVIS_ExAC_4KW", "LoFtool_scores", "tmcsnpdb"},
	},
	{
		Name: "db/annovar-knowngene",
		URL: []string{"http://www.openbioinformatics.org/annovar/download/{{builder}}_knownGene.txt.gz",
			"http://www.openbioinformatics.org/annovar/download/{{builder}}_knownGeneMrna.fa.gz",
			"http://www.openbioinformatics.org/annovar/download/{{builder}}_kgXref.txt.gz"},
	},
	{
		Name: "db/annovar-ensgene",
		URL: []string{"http://www.openbioinformatics.org/annovar/download/{{builder}}_ensGene.txt.gz",
			"http://www.openbioinformatics.org/annovar/download/{{builder}}_ensGeneMrna.fa.gz"},
	},
	{
		Name: "db/annovar-refgene",
		URL: []string{"http://www.openbioinformatics.org/annovar/download/{{builder}}_refGene.txt.gz",
			"http://www.openbioinformatics.org/annovar/download/{{builder}}_refGeneMrna.fa.gz",
			"http://www.openbioinformatics.org/annovar/download/{{builder}}_refGeneVersion.txt.gz"},
	},
	{
		Name: "db/ucsc-cytoband",
		URL:  []string{"http://hgdownload.cse.ucsc.edu/goldenPath/{{builder}}/database/cytoBand.txt.gz"},
	},
	{
		Name: "db/ucsc-dnase-clustered",
		URL:  []string{"http://hgdownload.cse.ucsc.edu/goldenPath/{{builder}}/encodeDCC/wgEncodeRegDnaseClustered/wgEncodeRegDnaseClustered{{version}}.bed.gz"},
	},
	{
		Name: "db/ucsc-ensgene",
		URL:  []string{"http://hgdownload.cse.ucsc.edu/goldenPath/{{builder}}/database/ensGene.txt.gz"},
	},
	{
		Name: "db/ucsc-knowngene",
		URL: []string{"http://hgdownload.cse.ucsc.edu/goldenPath/{{builder}}/database/knownGene.txt.gz",
			"http://hgdownload.cse.ucsc.edu/goldenPath/{{builder}}/database/kgXref.txt.gz"},
	},
	{
		Name: "db/ucsc-refgene",
		URL:  []string{"http://hgdownload.cse.ucsc.edu/goldenPath/{{builder}}/database/refGene.txt.gz"},
	},
	{
		Name: "db/ucsc-tfbs-clustered",
		URL:  []string{"http://hgdownload.cse.ucsc.edu/goldenPath/{{builder}}/encodeDCC/wgEncodeRegTfbsClustered/wgEncodeRegTfbsClustered{{version}}.bed.gz"},
	},
	{
		Name:     "db/appris",
		URL:      []string{"http://apprisws.bioinfo.cnio.es/pub/current_release/datafiles/homo_sapiens/{{version}}/appris_data.principal.txt"},
		Versions: []string{"GRCh38", "rs108v26", "up201703v26", "a1v26", "GRCh37", "rs105v24", "g12v24"},
	},
	{
		Name:     "db/atcircdb",
		URL:      []string{"http://genome.sdau.edu.cn/circRNAdata/{{version}}"},
		Versions: []string{"sw_final_circ_web_final.xlsx", "enrich_circ.csv", "SRRList_622samples_website.csv", "SRRList_150samples_CIRI2_validation.xlsx", "psTarget.csv", "TAPIR.csv", "Sun.xls", "Ye.csv"},
	},
	{
		Name:     "db/awesome",
		URL:      []string{"http://www.awesome-hust.com/downloads/{{version}}.zip"},
		Versions: []string{"awesomeAll"},
	},
	{
		Name:     "db/biosystems",
		URL:      []string{"ftp://ftp.ncbi.nih.gov/pub/biosystems/biosystems.{{version}}/"},
		Versions: []string{"20170414", "20170421"},
	},
	{
		Name:     "db/cancer_hotspots",
		URL:      []string{"http://bioinfo.rjh.com.cn/download/bioinstaller/cancer_hotspots/cancer_hotspots_{{version}}.txt.gz"},
		Versions: []string{"v1_sheet1", "v1_sheet2", "v2_sheet1", "v2_sheet2"},
	},
	{
		Name:     "db/cancersplicingqtl",
		URL:      []string{"http://www.cancersplicingqtl-hust.com/downloads/{{version}}.xlsx"},
		Versions: []string{"ACC_sQTLs", "BLCA_sQTLs", "BRCA_sQTLs", "CESC_sQTLs", "CHOL_sQTLs", "COAD_sQTLs", "DLBC_sQTLs", "ESCA_sQTLs", "GBM_sQTLs", "HNSC_sQTLs", "KICH_sQTLs", "KIRC_sQTLs", "KIRP_sQTLs", "LAML_sQTLs", "LGG_sQTLs", "LIHC_sQTLs", "LUAD_sQTLs", "LUSC_sQTLs", "MESO_sQTLs", "OV_sQTLs", "PAAD_sQTLs", "PCPG_sQTLs", "PRAD_sQTLs", "READ_sQTLs", "SARC_sQTLs", "SKCM_sQTLs", "STAD_sQTLs", "TGCT_sQTLs", "THCA_sQTLs", "THYM_sQTLs", "UCEC_sQTLs", "UCS_sQTLs", "UVM_sQTLs", "ACC_Survival_sQTLs", "BLCA_Survival_sQTLs", "BRCA_Survival_sQTLs", "CESC_Survival_sQTLs", "CHOL_Survival_sQTLs", "COAD_Survival_sQTLs", "DLBC_Survival_sQTLs", "ESCA_Survival_sQTLs", "GBM_Survival_sQTLs", "HNSC_Survival_sQTLs", "KICH_Survival_sQTLs", "KIRC_Survival_sQTLs", "KIRP_Survival_sQTLs", "LAML_Survival_sQTLs", "LGG_Survival_sQTLs", "LIHC_Survival_sQTLs", "LUAD_Survival_sQTLs", "LUSC_Survival_sQTLs", "MESO_Survival_sQTLs", "OV_Survival_sQTLs", "PAAD_Survival_sQTLs", "PCPG_Survival_sQTLs", "PRAD_Survival_sQTLs", "READ_Survival_sQTLs", "SARC_Survival_sQTLs", "SKCM_Survival_sQTLs", "STAD_Survival_sQTLs", "TGCT_Survival_sQTLs", "THCA_Survival_sQTLs", "THYM_Survival_sQTLs", "UCEC_Survival_sQTLs", "UCS_Survival_sQTLs", "UVM_Survival_sQTLs", "ACC_GWAS_sQTLs", "BLCA_GWAS_sQTLs", "BRCA_GWAS_sQTLs", "CESC_GWAS_sQTLs", "CHOL_GWAS_sQTLs", "COAD_GWAS_sQTLs", "DLBC_GWAS_sQTLs", "ESCA_GWAS_sQTLs", "GBM_GWAS_sQTLs", "HNSC_GWAS_sQTLs", "KICH_GWAS_sQTLs", "KIRC_GWAS_sQTLs", "KIRP_GWAS_sQTLs", "LAML_GWAS_sQTLs", "LGG_GWAS_sQTLs", "LIHC_GWAS_sQTLs", "LUAD_GWAS_sQTLs", "LUSC_GWAS_sQTLs", "MESO_GWAS_sQTLs", "OV_GWAS_sQTLs", "PAAD_GWAS_sQTLs", "PCPG_GWAS_sQTLs", "PRAD_GWAS_sQTLs", "READ_GWAS_sQTLs", "SARC_GWAS_sQTLs", "SKCM_GWAS_sQTLs", "STAD_GWAS_sQTLs", "TGCT_GWAS_sQTLs", "THCA_GWAS_sQTLs", "THYM_GWAS_sQTLs", "UCEC_GWAS_sQTLs", "UCS_GWAS_sQTLs", "UVM_GWAS_sQTLs"},
	},
	{
		Name:     "db/cellmarker",
		URL:      []string{"http://biocc.hrbmu.edu.cn/CellMarker/download/{{version}}_cell_markers.txt"},
		Versions: []string{"all", "Human", "Mouse", "Single"},
	},
	{
		Name:     "db/cgi",
		URL:      []string{"http://bioinfo.rjh.com.cn/download/bioinstaller/cgi/{{version}}_latest.zip"},
		Versions: []string{"cancer_bioactivities", "catalog_of_cancer_genes", "catalog_of_validated_oncogenic_mutations", "cgi_biomarkers"},
	},
	{
		Name:     "db/circbase",
		URL:      []string{"http://circrna.org/download/{{version}}"},
		Versions: []string{"hsa_hg19_circRNA.txt", "hsa_hg19_circRNA.xlsx", "hsa_hg19_circRNA.bed", "hsa_graph.png", "ce6_circID_to_name.txt", "mm9_circID_to_name.txt", "hg19_circID_to_name.txt", "find_circ.tar.gz", "mouse_mm9_circRNAs_putative_spliced_sequence.fa.gz", "human_hg19_circRNAs_putative_spliced_sequence.fa.gz"},
	},
	{
		Name:     "db/circnet",
		URL:      []string{"http://140.120.203.41/{{version}}.zip"},
		Versions: []string{"CircNetData"},
	},
	{
		Name:     "db/circrnadb",
		URL:      []string{"http://202.195.183.4:8000/circrnadb/doc/{{version}}.zip"},
		Versions: []string{"circRNADb", "genes.gtf", "Protein_features", "Protein_expression_evidence", "circRNA_parentalgene-diseases"},
	},
	{
		Name:     "db/civic",
		URL:      []string{"https://civic.genome.wustl.edu/downloads/nightly/nightly-{{version}}.tsv"},
		Versions: []string{"GeneSummaries", "VariantSummaries", "VariantGroupSummaries", "ClinicalEvidenceSummaries"},
	},
	{
		Name:     "db/consensuspathdb",
		URL:      []string{"http://cpdb.molgen.mpg.de/download/ConsensusPathDB_{{version}}.gz"},
		Versions: []string{"human_PPI", "human_PPI.psi25"},
	},
	{
		Name:     "db/cscd",
		URL:      []string{"http://gb.whu.edu.cn/CSCD/public/static/download/{{version}}.tar.gz"},
		Versions: []string{"hg19-cancer-circrna", "hg19-cancer-circrna-mre", "hg19-cancer-circrna-rbp", "hg19-cancer-circrna-orf", "hg19-normal-circrna", "hg19-normal-circrna-mre", "hg19-normal-circrna-rbp", "hg19-normal-circrna-orf", "hg19-common-circrna", "hg19-common-circrna-mre", "hg19-common-circrna-rbp", "hg19-common-circrna-orf", "hg38-cancer-circrna", "hg38-cancer-circrna-mre", "hg38-cancer-circrna-rbp", "hg38-cancer-circrna-orf", "hg38-normal-circrna", "hg38-normal-circrna-mre", "hg38-normal-circrna-rbp", "hg38-normal-circrna-orf", "hg38-common-circrna", "hg38-common-circrna-mre", "hg38-common-circrna-rbp", "hg38-common-circrna-orf"},
	},
	{
		Name:     "db/dbsno",
		URL:      []string{"http://140.138.144.145/~dbSNO/download/dbSNO{{version}}_all_data.txt.gz"},
		Versions: []string{"v2"},
	},
	{
		Name:     "db/dcdb",
		URL:      []string{"http://www.cls.zju.edu.cn/dcdb/downloadfile/{{version}}.zip"},
		Versions: []string{"DCDB2_plaintxt", "DCDB2.sql", "targets", "Drug_combinations", "components_identifier"},
	},
	{
		Name:     "db/denovo_db",
		URL:      []string{"http://denovo-db.gs.washington.edu/denovo-db.non-ssc-samples.variants.tsv.gz", "http://denovo-db.gs.washington.edu/denovo-db.non-ssc-samples.variants.vcf.gz", "http://denovo-db.gs.washington.edu/denovo-db.ssc-samples.variants.tsv.gz", "http://denovo-db.gs.washington.edu/denovo-db.ssc-samples.variants.vcf.gz", "http://denovo-db.gs.washington.edu/denovo-db.v.{{version}}.pdf"},
		Versions: []string{"1.6.1"},
	},
	{
		Name:     "db/dgidb",
		URL:      []string{"http://www.dgidb.org/data/{{version}}"},
		Versions: []string{"interactions.tsv", "genes.tsv", "drugs.tsv", "categories.tsv", "data/dGene_04-16-2013_2257.tsv", "data.sql"},
	},
	{
		Name:     "db/diseaseenhancer",
		URL:      []string{"http://biocc.hrbmu.edu.cn/DiseaseEnhancer/RFunctions/{{version}}.txt"},
		Versions: []string{"diseaseEnh5.1", "enhInfo", "enh2target"},
	},
	{
		Name:     "db/eggnog",
		URL:      []string{"http://eggnogdb.embl.de/download/eggnog_{{release}}/data/{{version}}/{{version}}.members.tsv.gz", "http://eggnogdb.embl.de/download/eggnog_{{release}}/data/{{version}}/{{version}}.annotations.tsv.gz", "http://eggnogdb.embl.de/download/eggnog_{{release}}/data/{{version}}/{{version}}.trees.tsv.gz", "http://eggnogdb.embl.de/download/eggnog_{{release}}/data/{{version}}/{{version}}.raw_algs.tsv.gz", "http://eggnogdb.embl.de/download/eggnog_{{release}}/data/{{version}}/{{version}}.trimmed_algs.tsv.gz", "http://eggnogdb.embl.de/download/eggnog_{{release}}/data/{{version}}/{{version}}.hmm.tsv.gz"},
		Versions: []string{"NOG", "aciNOG", "acidNOG", "acoNOG", "actNOG", "agaNOG", "agarNOG", "apiNOG", "aproNOG", "aquNOG", "arNOG", "arcNOG", "artNOG", "arthNOG", "ascNOG", "aveNOG", "bacNOG", "bactNOG", "bacteNOG", "basNOG", "bctoNOG", "biNOG", "bproNOG", "braNOG", "carNOG", "chaNOG", "chlNOG", "chlaNOG", "chloNOG", "chlorNOG", "chloroNOG", "chorNOG", "chrNOG", "cloNOG", "cocNOG", "creNOG", "cryNOG", "cyaNOG", "cytNOG", "debNOG", "defNOG", "dehNOG", "deiNOG", "delNOG", "dipNOG", "dotNOG", "dproNOG", "droNOG", "eproNOG", "eryNOG", "euNOG", "eurNOG", "euroNOG", "eurotNOG", "fiNOG", "firmNOG", "flaNOG", "fuNOG", "fusoNOG", "gproNOG", "haeNOG", "halNOG", "homNOG", "hymNOG", "hypNOG", "inNOG", "kinNOG", "lepNOG", "lilNOG", "maNOG", "magNOG", "meNOG", "metNOG", "methNOG", "methaNOG", "necNOG", "negNOG", "nemNOG", "onyNOG", "opiNOG", "perNOG", "plaNOG", "pleNOG", "poaNOG", "prNOG", "proNOG", "rhaNOG", "roNOG", "sacNOG", "saccNOG", "sorNOG", "sordNOG", "sphNOG", "spiNOG", "spriNOG", "strNOG", "synNOG", "tenNOG", "thaNOG", "theNOG", "therNOG", "thermNOG", "treNOG", "veNOG", "verNOG", "verrNOG", "virNOG"},
	},
	{
		Name:     "db/ewasdb",
		URL:      []string{"http://www.bioapp.org/ewasdb/Public/file/{{version}}.rar"},
		Versions: []string{"ewas_singlemarker", "GO_Category", "KEGG_Pathway"},
	},
	{
		Name:     "db/exorbase",
		URL:      []string{"http://www.exorbase.org/exoRBase/download/download?file={{version}}.txt"},
		Versions: []string{"Normal_circRNA_RPM", "CHD_circRNA_RPM", "CRC_circRNA_RPM", "HCC_circRNA_RPM", "PAAD_circRNA_RPM", "WhB_circRNA_RPM", "Samples_combined_circRNA_RPM", "Normal_lncRNA_TPM", "CHD_lncRNA_TPM", "CRC_lncRNA_TPM", "HCC_lncRNA_TPM", "PAAD_lncRNA_TPM", "WhB_lncRNA_TPM", "Samples_combined_lncRNA_TPM", "Normal_mRNA_TPM", "CHD_mRNA_TPM", "CRC_mRNA_TPM", "HCC_mRNA_TPM", "PAAD_mRNA_TPM", "WhB_mRNA_TPM", "Samples_combined_mRNA_TPM", "circRNA_ann", "lncRNA_ann", "mRNA_ann", "Experimental_revised"},
	},
	{
		Name:     "db/expression_atlas",
		URL:      []string{"ftp://ftp.ebi.ac.uk/pub/databases/microarray/data/atlas/experiments/atlas-{{version}}-data.tar.gz"},
		Versions: []string{"latest"},
	},
	{
		Name:     "db/exsnp",
		URL:      []string{"http://www.exsnp.org/data/{{version}}.txt"},
		Versions: []string{"raw_eQTL", "GSexSNP_allc_allp_ld8", "GSexSNP_allc_allp_ld5", "GSexSNP_allc_allp_ld3", "WTCCC_DZ_allc_allp_GSeQTL_0.05cM_list_BD", "WTCCC_DZ_allc_allp_GSeQTL_0.05cM_list_CAD", "WTCCC_DZ_allc_allp_GSeQTL_0.05cM_list_CD", "WTCCC_DZ_allc_allp_GSeQTL_0.05cM_list_HT", "WTCCC_DZ_allc_allp_GSeQTL_0.05cM_list_RA", "WTCCC_DZ_allc_allp_GSeQTL_0.05cM_list_T1D", "WTCCC_DZ_allc_allp_GSeQTL_0.05cM_list_T2D"},
	},
	{
		Name:     "db/fantom_cage_peaks",
		URL:      []string{"http://fantom.gsc.riken.jp/5/datafiles/latest/extra/{{version}}/"},
		Versions: []string{"CAGE_peaks"},
	},
	{
		Name:     "db/fantom_co_expression_clusters",
		URL:      []string{"http://fantom.gsc.riken.jp/5/datafiles/phase1.3/extra/{{version}}/"},
		Versions: []string{"Co-expression_clusters"},
	},
	{
		Name:     "db/fantom_enhancers",
		URL:      []string{"http://fantom.gsc.riken.jp/5/datafiles/latest/extra/{{version}}/"},
		Versions: []string{"Enhancers"},
	},
	{
		Name:     "db/fantom_motifs",
		URL:      []string{"http://fantom.gsc.riken.jp/5/datafiles/phase1.3/extra/Motifs/{{version}}/"},
		Versions: []string{"/", "TFBS", "jaspar_Significance_of_the_correlation_with_CAGE_expression", "novel_GREAT_analysis_results", "novel_Significance_of_the_correlation_with_CAGE_expression"},
	},
	{
		Name:     "db/fantom_ontology",
		URL:      []string{"http://fantom.gsc.riken.jp/5/datafiles/latest/extra/{{version}}/"},
		Versions: []string{"Ontology"},
	},
	{
		Name:     "db/fantom_tss_classifier",
		URL:      []string{"http://fantom.gsc.riken.jp/5/datafiles/phase1.3/extra/{{version}}/"},
		Versions: []string{"TSS_classifier"},
	},
	{
		Name:     "db/funcoup",
		URL:      []string{"http://funcoup.sbc.su.se/downloads/download.action?type=network&instanceID=7617155&fileName={{version}}.gz"},
		Versions: []string{"FC4.0_A.thaliana_full", "FC4.0_B.subtilis_full", "FC4.0_B.taurus_full", "FC4.0_C.elegans_full", "FC4.0_C.familiaris_full", "FC4.0_C.intestinalis_full", "FC4.0_D.melanogaster_full", "FC4.0_D.rerio_full", "FC4.0_E.coli_full", "FC4.0_G.gallus_full", "FC4.0_H.sapiens_full", "FC4.0_M.musculus_full", "FC4.0_O.sativa_full", "FC4.0_P.falciparum_full", "FC4.0_R.norvegicus_full", "FC4.0_S.cerevisiae_full", "FC4.0_S.pombe_full", "FC4.0_A.thaliana_compact", "FC4.0_B.subtilis_compact", "FC4.0_B.taurus_compact", "FC4.0_C.elegans_compact", "FC4.0_C.familiaris_compact", "FC4.0_C.intestinalis_compact", "FC4.0_D.melanogaster_compact", "FC4.0_D.rerio_compact", "FC4.0_E.coli_compact", "FC4.0_G.gallus_compact", "FC4.0_H.sapiens_compact", "FC4.0_M.musculus_compact", "FC4.0_O.sativa_compact", "FC4.0_P.falciparum_compact", "FC4.0_R.norvegicus_compact", "FC4.0_S.cerevisiae_compact", "FC4.0_S.pombe_compact"},
	},
	{
		Name:     "db/fusiongdb",
		URL:      []string{"https://ccsm.uth.edu/FusionGDB/tables/{{version}}.txt"},
		Versions: []string{"TCGA_ChiTaRS_combined_fusion_information_on_hg19", "TCGA_ChiTaRS_combined_fusion_ORF_analyzed_gencode_h19v19_fgID", "TCGA_ChiTaRS_combined_fusion_ORF_analyzed_gencode_h19v19", "TCGA_ChiTaRS_combined_fusion_Prot_feature_retention_search_UniProt_for_Hgene", "TCGA_ChiTaRS_combined_fusion_Prot_feature_retention_search_UniProt_for_Tgene", "fusion_ppi", "fusion_uniprot_related_drugs", "fgene_disease_associations", "uniprot_gsymbol", "TCGA_ChiTaRS_combined_fusion_ORF_analyzed_gencode_h19v19_In-frame_100k_check_aminoacid_seq", "TCGA_ChiTaRS_combined_fusion_ORF_analyzed_gencode_h19v19_In-frame_100k_check_cds_seq", "TCGA_ChiTaRS_combined_fusion_ORF_analyzed_gencode_h19v19_In-frame_100k_check_transcript_seq", "In-frame_fgene_num_samples"},
	},
	{
		Name:     "db/hgnc",
		URL:      []string{"ftp://ftp.ebi.ac.uk/pub/databases/genenames/new/tsv/locus_groups/{{version}}.txt", "ftp://ftp.ebi.ac.uk/pub/databases/genenames/new/tsv/{{version}}.txt"},
		Versions: []string{"protein-coding_gene", "non-coding_RNA", "phenotype", "pseudogene", "other", "withdrawn", "protein-coding_gene_alt_loci", "non-coding_RNA_alt_loci", "pseudogene_alt_loci", "other_alt_loci", "genefamilies"},
	},
	{
		Name:     "db/hmdb",
		URL:      []string{"http://www.hmdb.ca/system/downloads/current/{{version}}.zip", "http://specdb.wishartlab.com/downloads/exports/peak_lists_txt/{{version}}.zip", "http://www.hmdb.ca/system/downloads/current/sequences/{{version}}.zip", "http://specdb.wishartlab.com/downloads/exports/spectra_xml/{{version}}.zip", "http://specdb.wishartlab.com/downloads/exports/fid_files/{{version}}.zip", "http://specdb.wishartlab.com/downloads/exports/image_files/{{version}}.zip"},
		Versions: []string{""},
	},
	{
		Name:     "db/hpdi",
		URL:      []string{"http://bioinfo.wilmer.jhu.edu/PDI/{{version}}"},
		Versions: []string{"protein_chip_full_seq.csv", "protein_annotation.txt", "pro2motif.txt", "DNA_motifs.txt", "motif2protein.txt", "all_pwm.zip", "all_gpr_files.zip", "supplemental.pdf"},
	},
	{
		Name:     "db/hpo",
		URL:      []string{"http://compbio.charite.de/jenkins/job/hpo.annotations.monthly/{{version}}/artifact/annotation/*zip*/annotation.zip"},
		Versions: []string{"lastStableBuild"},
	},
	{
		Name:     "db/inbiomap",
		URL:      []string{"https://www.intomics.com/inbio/map/api/get_data?file=InBio_Map_core_{{version}}.tar.gz"},
		Versions: []string{"2016_09_12", "2016_05_31", "2016_02_05", "2015_11_02"},
	},
	{
		Name:     "db/instruct",
		URL:      []string{"http://instruct.yulab.org/download/{{version}}.sin"},
		Versions: []string{"sapiens", "thaliana", "elegans", "melanogaster", "musculus", "cerevisiae", "pombe"},
	},
	{
		Name:     "db/interpro",
		URL:      []string{"ftp://ftp.ebi.ac.uk/pub/databases/interpro/{{version}}"},
		Versions: []string{"entry.list", "interpro.xml.gz", "match_complete.xml.gz", "uniparc_match.tar.gz", "protein2ipr.dat.gz", "ParentChildTreeFile.txt", "interpro2go", "release_notes.txt"},
	},
	{
		Name:     "db/islandviewer",
		URL:      []string{"http://www.pathogenomics.sfu.ca/islandviewer/download/datasets/all_gis_{{version}}.txt.tar.gz"},
		Versions: []string{"islandviewer_iv4", "islandpick_iv4", "islandpath_dimob_iv4", "sigi_hmm_iv4", "islander_iv4"},
	},
	{
		Name:     "db/lnc2cancer",
		URL:      []string{"http://www.bio-bigdata.net/lnc2cancer/download/{{version}}.txt"},
		Versions: []string{"all_lnc_cancer_associations", "Circulating_lnc_cancer_associations", "Prognostic_lnc_cancer_associations", "Drug-resistant_lnc_cancer_associations", "LncRNAs%20regulated%20by%20TF%20in%20cancers", "LncRNAs%20regulated%20by%20miRNA%20in%20cancers", "LncRNAs%20regulated%20by%20variant%20in%20cancers", "LncRNAs%20regulated%20by%20methylation%20in%20cancers"},
	},
	{
		Name:     "db/lncediting",
		URL:      []string{"http://bioinfo.life.hust.edu.cn/LNCediting/static/human/download/{{version}}"},
		Versions: []string{"table1_lncrna_info", "table2_edit_sequence", "table3_edit_site_info", "table4_mirna_info", "table5_mirna_profile", "table7_function_gain", "table8_function_loss"},
	},
	{
		Name:     "db/lncrnadisease",
		URL:      []string{"http://www.rnanut.net/lncrnadisease/download/{{version}}.xlsx"},
		Versions: []string{"experimental%20circRNA-disease%20information", "experimental%20lncRNA-disease%20information", "predicted%20lncRNA-disease%20information", "all%20ncRNA-disease%20information"},
	},
	{
		Name:     "db/mircancer",
		URL:      []string{"http://mircancer.ecu.edu/downloads/{{version}}.txt"},
		Versions: []string{"miRCancerOctober2017", "miRCancerMarch2017", "miRCancerDecember2016", "miRCancerSeptember2016", "miRCancerJune2016", "miRCancerMarch2016", "miRCancerDecember2015", "miRCancerSeptember2015", "miRCancerJune2015", "miRCancerMarch2015", "miRCancerDecember2014", "miRCancerSeptember2014", "miRCancerJune2014", "miRCancerMarch2014", "miRCancerDecember2013", "miRCancerSeptember2013", "miRCancerJune2013", "miRCancerMarch2013", "miRCancerNovember2012"},
	},
	{
		Name:     "db/mirdb",
		URL:      []string{"http://mirdb.org/download/miRDB_v{{version}}_prediction_result.txt.gz"},
		Versions: []string{"5.0", "4.0", "3.0", "2.0", "1.0"},
	},
	{
		Name:     "db/mirnest",
		URL:      []string{"http://rhesus.amu.edu.pl/mirnest/copy/downloads/{{version}}.gz"},
		Versions: []string{"mirnest_EST_predictions", "mirnest_targets", "mirnest_deep_predictions", "mirnest_degradomes", "mirnest_mirtrons", "mirnest_mirna_gene_structure"},
	},
	{
		Name:     "db/mirtarbase",
		URL:      []string{"http://mirtarbase.mbc.nctu.edu.tw/cache/download/{{version}}/miRTarBase_MTI.xlsx"},
		Versions: []string{"7.0"},
	},
	{
		Name:     "db/mndr",
		URL:      []string{"http://www.rna-society.org/mndr/php/download.php?file={{version}}.zip"},
		Versions: []string{"All ncRNA-disease", "miRNA-disease", "lncRNA-disease"},
	},
	{
		Name:     "db/msdd",
		URL:      []string{"http://www.bio-bigdata.com/msdd/download/{{version}}.zip"},
		Versions: []string{"msdd"},
	},
	{
		Name:     "db/omim_open",
		URL:      []string{"https://omim.org/static/omim/data/{{version}}.txt", "https://raw.githubusercontent.com/macarthur-lab/omim/master/data/{{version}}.txt"},
		Versions: []string{"mim2gene", "full_omim_table", "gene_symbol_thesaurus", "use_omim_table"},
	},
	{
		Name:     "db/omim_private",
		URL:      []string{"https://data.omim.org/downloads/{{license}}/{{version}}.txt", "https://omim.org/static/omim/data/{{version}}.txt"},
		Versions: []string{"mim2gene", "genemap2", "genemap", "mimTitles", "morbidmap"},
	},
	{
		Name:     "db/oncokb",
		URL:      []string{"http://oncokb.org/api/v1/utils/{{version}}.txt"},
		Versions: []string{"allAnnotatedVariants", "allActionableVariants"},
	},
	{
		Name:     "db/oncomirdb",
		URL:      []string{"http://lifeome.net/database/oncomirdb/oncomirdb.v-{{version}}_download.txt"},
		Versions: []string{"1.1-20131217"},
	},
	{
		Name:     "db/oncotator",
		URL:      []string{"ftp://gsapubftp-anonymous@ftp.broadinstitute.org/bundle/oncotator/oncotator_{{version}}.tar.gz"},
		Versions: []string{"v1_ds_April052016"},
	},
	{
		Name:     "db/pancanqtl",
		URL:      []string{"http://bioinfo.life.hust.edu.cn/PancanQTL/static/download/{{version}}.xls"},
		Versions: []string{"ACC_tumor.cis_eQTL", "BLCA_tumor.cis_eQTL", "BRCA_tumor.cis_eQTL", "CESC_tumor.cis_eQTL", "CHOL_tumor.cis_eQTL", "COAD_tumor.cis_eQTL", "DLBC_tumor.cis_eQTL", "ESCA_tumor.cis_eQTL", "GBM_tumor.cis_eQTL", "HNSC_tumor.cis_eQTL", "KICH_tumor.cis_eQTL", "KIRC_tumor.cis_eQTL", "KIRP_tumor.cis_eQTL", "LAML_tumor.cis_eQTL", "LGG_tumor.cis_eQTL", "LIHC_tumor.cis_eQTL", "LUAD_tumor.cis_eQTL", "LUSC_tumor.cis_eQTL", "MESO_tumor.cis_eQTL", "OV_tumor.cis_eQTL", "PAAD_tumor.cis_eQTL", "PCPG_tumor.cis_eQTL", "PRAD_tumor.cis_eQTL", "READ_tumor.cis_eQTL", "SARC_tumor.cis_eQTL", "SKCM_tumor.cis_eQTL", "STAD_tumor.cis_eQTL", "TGCT_tumor.cis_eQTL", "THCA_tumor.cis_eQTL", "THYM_tumor.cis_eQTL", "UCEC_tumor.cis_eQTL", "UCS_tumor.cis_eQTL", "UVM_tumor.cis_eQTL", "ACC_tumor.trans_eQTL", "BLCA_tumor.trans_eQTL", "BRCA_tumor.trans_eQTL", "CESC_tumor.trans_eQTL", "CHOL_tumor.trans_eQTL", "COAD_tumor.trans_eQTL", "DLBC_tumor.trans_eQTL", "ESCA_tumor.trans_eQTL", "GBM_tumor.trans_eQTL", "HNSC_tumor.trans_eQTL", "KICH_tumor.trans_eQTL", "KIRC_tumor.trans_eQTL", "KIRP_tumor.trans_eQTL", "LAML_tumor.trans_eQTL", "LGG_tumor.trans_eQTL", "LIHC_tumor.trans_eQTL", "LUAD_tumor.trans_eQTL", "LUSC_tumor.trans_eQTL", "MESO_tumor.trans_eQTL", "OV_tumor.trans_eQTL", "PAAD_tumor.trans_eQTL", "PCPG_tumor.trans_eQTL", "PRAD_tumor.trans_eQTL", "READ_tumor.trans_eQTL", "SARC_tumor.trans_eQTL", "SKCM_tumor.trans_eQTL", "STAD_tumor.trans_eQTL", "TGCT_tumor.trans_eQTL", "THCA_tumor.trans_eQTL", "THYM_tumor.trans_eQTL", "UCEC_tumor.trans_eQTL", "UCS_tumor.trans_eQTL", "UVM_tumor.trans_eQTL"},
	},
	{
		Name:     "db/phosphonetworks",
		URL:      []string{""},
		Versions: []string{"rawKSI.csv", "refKSI.csv", "comKSI.csv", "motifSite.csv", "motifMatrix.csv", "motifLogo.tar", "highResolutionNetwork.csv", "supplemental.pdf"},
	},
	{
		Name:     "db/pmkb",
		URL:      []string{"https://pmkb.weill.cornell.edu/therapies/download.xlsx"},
		Versions: []string{"latest"},
	},
	{
		Name:     "db/proteinatlas",
		URL:      []string{"https://www.proteinatlas.org/download/{{version}}.tsv.zip", "https://www.proteinatlas.org/download/{{version}}.gz", "https://www.proteinatlas.org/images_static/{{version}}"},
		Versions: []string{"normal_tissue", "pathology", "subcellular_location", "rna_tissue", "transcript_rna_tissue", "transcript_rna_celline", "proteinatlas", "proteinatlas.xml", "proteinatlas.trig", "cell.svg"},
	},
	{
		Name:     "db/rbp_var",
		URL:      []string{"http://159.226.67.237/sun/RBP-Var/res/{{version}}.RBP-Var.score.p.txt.gz"},
		Versions: []string{"RNAedit", "dbSNP"},
	},
	{
		Name:     "db/rbpdb",
		URL:      []string{"http://rbpdb.ccbr.utoronto.ca/downloads/{{version}}"},
		Versions: []string{"RBPDB_v1.3.1_2012-11-21.sql", "RBPDB_v1.3.1_2012-11-21_TDT.zip", "RBPDB_v1.3.1_2012-11-21_CSV.zip"},
	},
	{
		Name:     "db/rddpred",
		URL:      []string{"http://epigenomics.snu.ac.kr/RDDpred/prior_data/{{version}}.MES.txt.gz"},
		Versions: []string{"Human", "Mouse", "Drosophila", "Arabidopsis"},
	},
	{
		Name:     "db/redoxdb",
		URL:      []string{"https://biocomputer.bio.cuhk.edu.hk/RedoxDB/download/{{version}}"},
		Versions: []string{"redoxdb.A.txt", "redoxdb.B.txt", "redoxdb.A.fa", "redoxdb.B.fa"},
	},
	{
		Name:     "db/remap",
		URL:      []string{"http://tagc.univ-mrs.fr/remap/download/remap1/ReMap2015_TF_{{version}}Peaks.tar.gz", "http://tagc.univ-mrs.fr/remap/download/ReMap1_lifted_hg38/{{version}}.bed.gz", "http://tagc.univ-mrs.fr/remap/download/remap1/All/{{version}}.bed.gz"},
		Versions: []string{"archive_all", "archive_nr", "filPeaks_public", "filPeaks_all", "nrPeaks_public", "nrPeaks_all", "hg38_allPeaks", "hg38_nrPeaks"},
	},
	{
		Name:     "db/remap2",
		URL:      []string{"http://tagc.univ-mrs.fr/remap/download/MACS/ReMap2_TF_{{version}}Peaks.tar.gz", "http://tagc.univ-mrs.fr/remap/download/MACS/ReMap2_{{version}}Peaks.bed.gz", "http://tagc.univ-mrs.fr/remap/download/MACS/ReMap2_{{version}}.bed.gz", "@<@if (str_detect('{{version}}', 'hg19')) {str_replace('{{version}}', '_hg19', '');'http://tagc.univ-mrs.fr/remap/download/MACS_lifted_hg19/ReMap2_TF_{{version}}Peaks_hg19.tar.gz'}@<@", "@<@if (str_detect('{{version}}', 'hg19')) {str_replace('{{version}}', '_hg19', '');'http://tagc.univ-mrs.fr/remap/download/MACS_lifted_hg19/ReMap2_{{version}}Peaks_hg19.bed.gz'}@<@", "@<@if (str_detect('{{version}}', 'hg19')) {str_replace('{{version}}', '_hg19', '');'http://tagc.univ-mrs.fr/remap/download/MACS_lifted_hg19/ReMap2_{{version}}_hg19.bed.gz'}@<@"},
		Versions: []string{"archive_all", "archive_nr", "archive_all_hg19", "archive_nr_hg19", "public_all", "public_crm", "public_nr", "encode_all", "encode_crm", "encode_nr", "all", "crm", "nr"},
	},
	{
		Name:     "db/rsnp3",
		URL:      []string{"ftp://rv.psych.ac.cn/pub/rsnp3/rSNP_info_{{version}}.tar.gz"},
		Versions: []string{"all"},
	},
	{
		Name:     "db/rvarbase",
		URL:      []string{"ftp://rv.psych.ac.cn/pub/rv/{{version}}.txt.gz"},
		Versions: []string{"dbsnp_info1", "dbvar_info"},
	},
	{
		Name:     "db/sedb",
		URL:      []string{"http://www.licpathway.net/sedb/download/{{version}}"},
		Versions: []string{"latest"},
	},
	{
		Name:     "db/seecancer",
		URL:      []string{"http://biocc.hrbmu.edu.cn/SEECancer/download/{{version}}.txt"},
		Versions: []string{"early_events", "late_events", "relapse_events", "metastasis_events", "drug_resistant_events", "drug_induced_events", "temporal_sequence"},
	},
	{
		Name:     "db/seeqtl",
		URL:      []string{"http://www.bios.unc.edu/research/genomic_software/seeQTL/data/{{version}}.tar.gz", "http://www.bios.unc.edu/research/genomic_software/seeQTL/data/{{version}}.txt"},
		Versions: []string{"pca_genotype", "genotype_expression_data", "eQTL_Qvalue_cutoff_hapmap3_cis_hg19", "eQTL_Qvalue_cutoff_hapmap3_trans_hg19", "eQTL_Qvalue_cutoff_Myers_cis_hg19", "eQTL_Qvalue_cutoff_Myers_trans_hg19", "data.montgomery", "data.argonne", "data.yale"},
	},
	{
		Name:     "db/sm2mir",
		URL:      []string{"http://210.46.85.180:8080/sm2mir/files/{{versin}}.xls"},
		Versions: []string{"SM2miR3", "SM2miR2n", "SM2miR"},
	},
	{
		Name:     "db/srnanalyzer",
		URL:      []string{"http://srnanalyzer.systemsbiology.net/downloads/{{version}}.tar.gz"},
		Versions: []string{"sRNA_DBs", "MainDBs", "NCBI_NonHuman"},
	},
	{
		Name:     "db/tumorfusions",
		URL:      []string{"http://www.tumorfusions.org/downloads/{{version}}.txt.gz"},
		Versions: []string{"pancanfus"},
	},
	{
		Name: "db/varcards",
		URL:  []string{"http://159.226.67.237/sun/varcards/resource/download/variant/chr{{chrom}}.variant.annotation.xls.gz"},
	},
	{
		Name: "db/clingov",
		URL:  []string{"https://ClinicalTrials.gov/AllAPIJSON.zip"},
	},
}
