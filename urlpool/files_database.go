package urlpool

var annovarLinks = []bgetFilesURLType{
	{
		Name: "db_annovar",
		URL: []string{"http://www.openbioinformatics.org/annovar/download/{{builder}}_{{version}}.txt.gz",
			"http://www.openbioinformatics.org/annovar/download/{{builder}}_{{version}}.txt.idx.gz"},
		Versions: []string{"clinvar_20131105", "clinvar_20140211", "clinvar_20140303", "clinvar_20140702", "clinvar_20140902", "clinvar_20140929", "clinvar_20150330", "clinvar_20150629", "clinvar_20151201", "clinvar_20160302", "clinvar_20161128", "clinvar_20170130", "clinvar_20170501", "clinvar_20170905", "clinvar_20180603", "avsnp150", "avsnp147", "avsnp144", "avsnp142", "avsnp138", "cadd", "caddgt10", "caddgt20", "cadd13", "cadd13gt10", "cadd13gt20", "cg69", "cg46", "cosmic70", "cosmic68wgs", "cosmic68", "cosmic67wgs", "cosmic67", "cosmic65", "cosmic64", "dbnsfp35a", "dbnsfp33a", "dbnsfp31a_interpro", "dbnsfp30a", "dbscsnv11", "eigen", "esp6500siv2_ea", "esp6500siv2_aa", "esp6500siv2_all", "exac03nontcga", "exac03nonpsych", "exac03", "fathmm", "gerp++elem", "gerp++gt2", "gme", "gnomad_exome", "gnomad_genome", "gwava", "hrcr1", "icgc21", "intervar_20170202",
			"kaviar_20150923", "ljb26_all", "mcap", "mitimpact2", "mitimpact24", "nci60", "popfreq_max_20150413", "popfreq_all_20150413", "revel", "regsnpintron"},
	},
	{
		Name:     "db_annovar_1000g",
		URL:      []string{"http://www.openbioinformatics.org/annovar/download/{{builder}}_{{version}}.zip"},
		Versions: []string{"1000g2015aug", "1000g2014oct", "1000g2014sep", "1000g2014aug", "1000g2012apr", "1000g2012feb", "1000g2011may", "1000g2010nov", "1000g2012apr", "1000g2010jul", "1000g2010", "1000g"},
	},
	{
		Name: "db_annovar_noidx",
		URL: []string{"http://www.openbioinformatics.org/annovar/download/{{version}}.txt.gz",
			"http://www.openbioinformatics.org/annovar/download/{{builder}}_{{version}}.txt.gz"},
		Versions: []string{"GDI_full_10282015", "RVIS_ExAC_4KW", "LoFtool_scores", "tmcsnpdb"},
	},
	{
		Name: "db_annovar_knowngene",
		URL: []string{"http://www.openbioinformatics.org/annovar/download/{{builder}}_knownGene.txt.gz",
			"http://www.openbioinformatics.org/annovar/download/{{builder}}_knownGeneMrna.fa.gz",
			"http://www.openbioinformatics.org/annovar/download/{{builder}}_kgXref.txt.gz"},
	},
	{
		Name: "db_annovar_ensgene",
		URL: []string{"http://www.openbioinformatics.org/annovar/download/{{builder}}_ensGene.txt.gz",
			"http://www.openbioinformatics.org/annovar/download/{{builder}}_ensGeneMrna.fa.gz"},
	},
	{
		Name: "db_annovar_refgene",
		URL: []string{"http://www.openbioinformatics.org/annovar/download/{{builder}}_refGene.txt.gz",
			"http://www.openbioinformatics.org/annovar/download/{{builder}}_refGeneMrna.fa.gz",
			"http://www.openbioinformatics.org/annovar/download/{{builder}}_refGeneVersion.txt.gz"},
	},
	{
		Name: "db_ucsc_cytoband",
		URL:  []string{"http://hgdownload.cse.ucsc.edu/goldenPath/{{builder}}/database/cytoBand.txt.gz"},
	},
	{
		Name: "db_ucsc_dnase_clustered",
		URL:  []string{"http://hgdownload.cse.ucsc.edu/goldenPath/{{builder}}/encodeDCC/wgEncodeRegDnaseClustered/wgEncodeRegDnaseClustered{{version}}.bed.gz"},
	},
	{
		Name: "db_ucsc_ensgene",
		URL:  []string{"http://hgdownload.cse.ucsc.edu/goldenPath/{{builder}}/database/ensGene.txt.gz"},
	},
	{
		Name: "db_ucsc_knowngene",
		URL: []string{"http://hgdownload.cse.ucsc.edu/goldenPath/{{builder}}/database/knownGene.txt.gz",
			"http://hgdownload.cse.ucsc.edu/goldenPath/{{builder}}/database/kgXref.txt.gz"},
	},
	{
		Name: "db_ucsc_refgene",
		URL:  []string{"http://hgdownload.cse.ucsc.edu/goldenPath/{{builder}}/database/refGene.txt.gz"},
	},
	{
		Name: "db_ucsc_tfbs_clustered",
		URL:  []string{"http://hgdownload.cse.ucsc.edu/goldenPath/{{builder}}/encodeDCC/wgEncodeRegTfbsClustered/wgEncodeRegTfbsClustered{{version}}.bed.gz"},
	},
}
