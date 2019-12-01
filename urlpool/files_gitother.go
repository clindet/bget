package urlpool

var otherGitRepos = []bgetFilesURLType{
	{
		Name:         "gitlab/leen",
		URL:          []string{"https://gitlab.com/bioai/leen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/acdc",
		URL:          []string{"https://bitbucket.org/dudleylab/acdc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/phd",
		URL:          []string{"https://bitbucket.org/miroslavradojevic/phd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/bioescorte-suggestion",
		URL:          []string{"https://bitbucket.org/ipk_bit_team/bioescorte-suggestion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/pipit",
		URL:          []string{"https://bitbucket.org/biovizleuven/pipit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/purge-haplotigs",
		URL:          []string{"https://bitbucket.org/mroachawri/purge_haplotigs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/aa-stat",
		URL:          []string{"https://bitbucket.org/j_bale/aa_stat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/infiniumpurify",
		URL:          []string{"https://bitbucket.org/zhengxiaoqi/infiniumpurify"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/bacmeta",
		URL:          []string{"https://bitbucket.org/aleksisipola/bacmeta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/dd-detection",
		URL:          []string{"https://bitbucket.org/mkroon/dd_detection"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/ldetect",
		URL:          []string{"https://bitbucket.org/nygcresearch/ldetect"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/ldetect-data",
		URL:          []string{"https://bitbucket.org/nygcresearch/ldetect-data"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/flexscore",
		URL:          []string{"https://bitbucket.org/mjamroz/flexscore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/vdjpuzzle2",
		URL:          []string{"https://bitbucket.org/kirbyvisp/vdjpuzzle2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/pancanceramaretto",
		URL:          []string{"https://bitbucket.org/gevaertlab/pancanceramaretto"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/aikyatan",
		URL:          []string{"https://bitbucket.org/cellsandmachines/aikyatan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/biobranch",
		URL:          []string{"https://bitbucket.org/sulab/biobranch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/biosubg",
		URL:          []string{"https://bitbucket.org/akittas/biosubg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/m2lite",
		URL:          []string{"https://bitbucket.org/paiyetan/m2lite"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/bam-matcher",
		URL:          []string{"https://bitbucket.org/sacgf/bam-matcher"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/badtrip",
		URL:          []string{"https://bitbucket.org/nicofmay/badtrip"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/jqcml",
		URL:          []string{"https://bitbucket.org/proteinspector/jqcml"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/metaphinder",
		URL:          []string{"https://bitbucket.org/genomicepidemiology/metaphinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/pheno-deep-counter",
		URL:          []string{"https://bitbucket.org/tuttoweb/pheno-deep-counter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/metabat",
		URL:          []string{"https://bitbucket.org/berkeleylab/metabat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/jra-src",
		URL:          []string{"https://bitbucket.org/jointreadalignment/jra-src"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/wanding/duprecover/overview",
		URL:          []string{"https://bitbucket.org/wanding/duprecover/overview"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/pbrit",
		URL:          []string{"https://bitbucket.org/medgenua/pbrit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/pyseqlab",
		URL:          []string{"https://bitbucket.org/a_2/pyseqlab"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/l.u.st",
		URL:          []string{"https://bitbucket.org/afro-juju/l.u.st"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/pli",
		URL:          []string{"https://bitbucket.org/astexuk/pli"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/mageck-nest",
		URL:          []string{"https://bitbucket.org/liulab/mageck_nest"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/cophesim",
		URL:          []string{"https://bitbucket.org/izhbannikov/cophesim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/scg",
		URL:          []string{"https://bitbucket.org/aroth85/scg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/venomix",
		URL:          []string{"https://bitbucket.org/jasonmacrander/venomix"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/read",
		URL:          []string{"https://bitbucket.org/tguenther/read"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/iso-ktsp",
		URL:          []string{"https://bitbucket.org/regulatorygenomicsupf/iso-ktsp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/intervene",
		URL:          []string{"https://bitbucket.org/cbgr/intervene"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/cl-dash",
		URL:          []string{"https://bitbucket.org/booz-allen-sci-comp-team/cl-dash"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/norgal",
		URL:          []string{"https://bitbucket.org/kosaidtu/norgal"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/hogan",
		URL:          []string{"https://bitbucket.org/xavmeyer/hogan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/brainformat",
		URL:          []string{"https://bitbucket.org/oruebel/brainformat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/tcellxtalk",
		URL:          []string{"https://bitbucket.org/lp-csic-uab/tcellxtalk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/phyparts",
		URL:          []string{"https://bitbucket.org/blackrim/phyparts"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/lipidhunter",
		URL:          []string{"https://bitbucket.org/sysmedos/lipidhunter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/domainoid",
		URL:          []string{"https://bitbucket.org/sonnhammergroup/domainoid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/simlord",
		URL:          []string{"https://bitbucket.org/genomeinformatics/simlord"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/cabsflex",
		URL:          []string{"https://bitbucket.org/lcbio/cabsflex"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/alpha",
		URL:          []string{"https://bitbucket.org/thekswenson/alpha"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/gabi",
		URL:          []string{"https://bitbucket.org/mhowison/gabi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/metaprob",
		URL:          []string{"https://bitbucket.org/samu661/metaprob"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/svd-phy",
		URL:          []string{"https://bitbucket.org/andrea/svd-phy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/sparkseq",
		URL:          []string{"https://bitbucket.org/mwiewiorka/sparkseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/chimiric",
		URL:          []string{"https://bitbucket.org/leslielab/chimiric"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/granger-causality",
		URL:          []string{"https://bitbucket.org/dtyu/granger-causality"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/phylogenomic-dataset-construction",
		URL:          []string{"https://bitbucket.org/yangya/phylogenomic_dataset_construction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/fastsimbac",
		URL:          []string{"https://bitbucket.org/nicofmay/fastsimbac"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/digtyper",
		URL:          []string{"https://bitbucket.org/jana_ebler/digtyper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/oncodriveclustl",
		URL:          []string{"https://bitbucket.org/bbglab/oncodriveclustl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/mutationmotif",
		URL:          []string{"https://bitbucket.org/pycogent3/mutationmotif"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/fbb",
		URL:          []string{"https://bitbucket.org/irenerodriguez/fbb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/seql-nrps",
		URL:          []string{"https://bitbucket.org/dansondergaard/seql-nrps"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/biofilmanalyzer",
		URL:          []string{"https://bitbucket.org/rogex/biofilmanalyzer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/percolator-upgrade",
		URL:          []string{"https://bitbucket.org/jthalloran/percolator_upgrade"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/mngsg",
		URL:          []string{"https://bitbucket.org/sm_islam/mngsg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/sandpuma",
		URL:          []string{"https://bitbucket.org/chevrm/sandpuma"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/nightshift",
		URL:          []string{"https://bitbucket.org/akdehof/nightshift"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/xpiwit",
		URL:          []string{"https://bitbucket.org/jstegmaier/xpiwit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/chemdistillerpython",
		URL:          []string{"https://bitbucket.org/ianalytica/chemdistillerpython"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/hgvs",
		URL:          []string{"https://bitbucket.org/hgvs/hgvs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/logminer",
		URL:          []string{"https://bitbucket.org/kleinstein/logminer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/locality",
		URL:          []string{"https://bitbucket.org/thekswenson/locality"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/benchmarks",
		URL:          []string{"https://bitbucket.org/biodbqual/benchmarks"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/cutruntools",
		URL:          []string{"https://bitbucket.org/qzhudfci/cutruntools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/orange-reliability",
		URL:          []string{"https://bitbucket.org/biolab/orange-reliability"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/godon",
		URL:          []string{"https://bitbucket.org/davydov/godon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/libsleipnir.bitbucket.org",
		URL:          []string{"https://bitbucket.org/libsleipnir/libsleipnir.bitbucket.org"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/triovis",
		URL:          []string{"https://bitbucket.org/biovizleuven/triovis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/hicapp",
		URL:          []string{"https://bitbucket.org/mthjwu/hicapp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/kit",
		URL:          []string{"https://bitbucket.org/jarmond/kit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/elsa",
		URL:          []string{"https://bitbucket.org/charade/elsa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/deepscope",
		URL:          []string{"https://bitbucket.org/aschaumberg/deepscope"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/infusion",
		URL:          []string{"https://bitbucket.org/kokonech/infusion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/enzymer",
		URL:          []string{"https://bitbucket.org/casraz/enzymer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/mageck-vispr",
		URL:          []string{"https://bitbucket.org/liulab/mageck-vispr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/avishkar",
		URL:          []string{"https://bitbucket.org/cellsandmachines/avishkar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/ddr",
		URL:          []string{"https://bitbucket.org/rso24/ddr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/g-tris",
		URL:          []string{"https://bitbucket.org/bereste/g-tris"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/depthfinder",
		URL:          []string{"https://bitbucket.org/jerlar73/depthfinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/nsclc-paper",
		URL:          []string{"https://bitbucket.org/youngjh/nsclc_paper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/vsclust",
		URL:          []string{"https://bitbucket.org/veitveit/vsclust"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/oncorep",
		URL:          []string{"https://bitbucket.org/sulab/oncorep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/rucs",
		URL:          []string{"https://bitbucket.org/genomicepidemiology/rucs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/glbase",
		URL:          []string{"https://bitbucket.org/oaxiom/glbase"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/hdgmcm",
		URL:          []string{"https://bitbucket.org/cdal/hdgmcm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/classyfire-api",
		URL:          []string{"https://bitbucket.org/wishartlab/classyfire_api"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/srg-extractor",
		URL:          []string{"https://bitbucket.org/jerlar73/srg-extractor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/shoelaces",
		URL:          []string{"https://bitbucket.org/valenlab/shoelaces"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/miraw",
		URL:          []string{"https://bitbucket.org/account/user/bipous/projects/miraw"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/metamodules",
		URL:          []string{"https://bitbucket.org/alimay/metamodules"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/subcons-web-server",
		URL:          []string{"https://bitbucket.org/salvatore_marco/subcons-web-server"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/mlm-gs-supplement",
		URL:          []string{"https://bitbucket.org/jwliang/mlm_gs_supplement"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/scavager",
		URL:          []string{"https://bitbucket.org/markmipt/scavager"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/cellsandmachines",
		URL:          []string{"https://bitbucket.org/cellsandmachines"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/mistrvar",
		URL:          []string{"https://bitbucket.org/compbio/mistrvar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/protie",
		URL:          []string{"https://bitbucket.org/compbio/protie"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/rppa-preprocess",
		URL:          []string{"https://bitbucket.org/rppa_preprocess/rppa_preprocess"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/smart-viewpoint-computation-lib",
		URL:          []string{"https://bitbucket.org/rranon/smart-viewpoint-computation-lib"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/neri",
		URL:          []string{"https://bitbucket.org/sergionery/neri"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/kernelsvmspark",
		URL:          []string{"https://bitbucket.org/cellsandmachines/kernelsvmspark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/stereodist",
		URL:          []string{"https://bitbucket.org/oesperlab/stereodist"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/arpeggio",
		URL:          []string{"https://bitbucket.org/harryjubb/arpeggio"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/em-dawid",
		URL:          []string{"https://bitbucket.org/a_2/em_dawid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/bioner",
		URL:          []string{"https://bitbucket.org/ianalytica/bioner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/warpiv",
		URL:          []string{"https://bitbucket.org/berkeleylab/warpiv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/rgaugury",
		URL:          []string{"https://bitbucket.org/yaanlpc/rgaugury"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/drseq",
		URL:          []string{"https://bitbucket.org/tarela/drseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/libfmftsaxs",
		URL:          []string{"https://bitbucket.org/abc-group/libfmftsaxs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/forqs",
		URL:          []string{"https://bitbucket.org/dkessner/forqs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/mrmplusgui",
		URL:          []string{"https://bitbucket.org/paiyetan/mrmplusgui"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/vcf2networks",
		URL:          []string{"https://bitbucket.org/dalloliogm/vcf2networks"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/sequence-shape",
		URL:          []string{"https://bitbucket.org/wenxiu/sequence-shape"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/genediseaserepositioning",
		URL:          []string{"https://bitbucket.org/ncl-intbio/genediseaserepositioning"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/pia",
		URL:          []string{"https://bitbucket.org/osiris_phylogenetics/pia"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/phyloskeleton",
		URL:          []string{"https://bitbucket.org/lionelguy/phyloskeleton"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/cabsdock",
		URL:          []string{"https://bitbucket.org/lcbio/cabsdock"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/last-split-pe",
		URL:          []string{"https://bitbucket.org/splitpairedend/last-split-pe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/cytostruct",
		URL:          []string{"https://bitbucket.org/sergeyn/cytostruct"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/autometa",
		URL:          []string{"https://bitbucket.org/jason_c_kwan/autometa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/resistome-release",
		URL:          []string{"https://bitbucket.org/jdwinkler/resistome_release"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/laser-release",
		URL:          []string{"https://bitbucket.org/jdwinkler/laser_release"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/shades",
		URL:          []string{"https://bitbucket.org/satsumaimo/shades"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/mzjava",
		URL:          []string{"https://bitbucket.org/sib-pig/mzjava"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/jaspar",
		URL:          []string{"https://bitbucket.org/cbgr/jaspar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/vispa2",
		URL:          []string{"https://bitbucket.org/andreacalabria/vispa2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/omics-pipe",
		URL:          []string{"https://bitbucket.org/sulab/omics_pipe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/qc-analysis",
		URL:          []string{"https://bitbucket.org/proteinspector/qc_analysis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/fit3d",
		URL:          []string{"https://bitbucket.org/fkaiser/fit3d"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/bioa",
		URL:          []string{"https://bitbucket.org/nmancuso/bioa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/repgenhmm",
		URL:          []string{"https://bitbucket.org/yuvalel/repgenhmm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/diprog",
		URL:          []string{"https://bitbucket.org/farahani/diprog"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/jgi-sparc",
		URL:          []string{"https://bitbucket.org/berkeleylab/jgi-sparc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/aodp-v2.0-release",
		URL:          []string{"https://bitbucket.org/wenchen_aafc/aodp_v2.0_release"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/genomemining",
		URL:          []string{"https://bitbucket.org/batlabucd/genomemining"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/mulcch",
		URL:          []string{"https://bitbucket.org/roygroup/mulcch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/micomplete",
		URL:          []string{"https://bitbucket.org/evolegiolab/micomplete"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/sparcc",
		URL:          []string{"https://bitbucket.org/yonatanf/sparcc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/galaxytoolfactory",
		URL:          []string{"https://bitbucket.org/fubar/galaxytoolfactory"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/readdi",
		URL:          []string{"https://bitbucket.org/readdi/readdi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/fimtyper",
		URL:          []string{"https://bitbucket.org/genomicepidemiology/fimtyper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/fimtyper-db",
		URL:          []string{"https://bitbucket.org/genomicepidemiology/fimtyper_db"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/bart",
		URL:          []string{"https://bitbucket.org/luisa_amaral/bart"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/sorting-by-network-completion",
		URL:          []string{"https://bitbucket.org/mattbiggs/sorting_by_network_completion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/aggrescan3d",
		URL:          []string{"https://bitbucket.org/lcbio/aggrescan3d"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/spurio",
		URL:          []string{"https://bitbucket.org/bateman-group/spurio"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/remixt",
		URL:          []string{"https://bitbucket.org/dranew/remixt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/sist-codes",
		URL:          []string{"https://bitbucket.org/benhamlab/sist_codes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/agalma",
		URL:          []string{"https://bitbucket.org/caseywdunn/agalma"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/splot",
		URL:          []string{"https://bitbucket.org/lkalesinskas/splot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/densitycut-dev",
		URL:          []string{"https://bitbucket.org/jerry00/densitycut_dev"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/sl1p",
		URL:          []string{"https://bitbucket.org/fwhelan/sl1p"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/biotransformerjar",
		URL:          []string{"https://bitbucket.org/djoumbou/biotransformerjar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/banner-chemdner",
		URL:          []string{"https://bitbucket.org/tsendeemts/banner-chemdner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/taxmapper",
		URL:          []string{"https://bitbucket.org/dbeisser/taxmapper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/rawmsa",
		URL:          []string{"https://bitbucket.org/clami66/rawmsa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/fast",
		URL:          []string{"https://bitbucket.org/baderlab/fast"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/dmml",
		URL:          []string{"https://bitbucket.org/anthakki/dmml"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/comorbidity",
		URL:          []string{"https://bitbucket.org/ibi_group/comorbidity"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/proteogenomics",
		URL:          []string{"https://bitbucket.org/andreyto/proteogenomics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/tpes",
		URL:          []string{"https://bitbucket.org/l0ka/tpes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/hh-suite",
		URL:          []string{"https://bitbucket.org/soedinglab/hh-suite"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/magicmicrorna",
		URL:          []string{"https://bitbucket.org/mutgx/magicmicrorna"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/vptissue",
		URL:          []string{"https://bitbucket.org/vptissue/vptissue"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/catkern",
		URL:          []string{"https://bitbucket.org/elies_ramon/catkern"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/moleculedatabaseframework",
		URL:          []string{"https://bitbucket.org/kienerj/moleculedatabaseframework"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/svist4get",
		URL:          []string{"https://bitbucket.org/artegorov/svist4get"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/chromawalker",
		URL:          []string{"https://bitbucket.org/zhenwahtan/chromawalker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/cytodx-study-code",
		URL:          []string{"https://bitbucket.org/zichenghu_ucsf/cytodx_study_code"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/cystine-stabilized-proteins",
		URL:          []string{"https://bitbucket.org/sm_islam/cystine-stabilized-proteins"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/ht-altfrac",
		URL:          []string{"https://bitbucket.org/oesperlab/ht-altfrac"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/cfs-toolbox",
		URL:          []string{"https://bitbucket.org/mikkonuutinen/cfs_toolbox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/mgmapper",
		URL:          []string{"https://bitbucket.org/genomicepidemiology/mgmapper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/genespider",
		URL:          []string{"https://bitbucket.org/sonnhammergrni/genespider"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/phosfox",
		URL:          []string{"https://bitbucket.org/phintsan/phosfox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/defcom",
		URL:          []string{"https://bitbucket.org/bryancquach/defcom"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/phylogenize",
		URL:          []string{"https://bitbucket.org/pbradz/phylogenize"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/monovar",
		URL:          []string{"https://bitbucket.org/hamimzafar/monovar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/immunediversity",
		URL:          []string{"https://bitbucket.org/immunediversity/immunediversity"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/ciml-lib",
		URL:          []string{"https://bitbucket.org/ciml-ufjf/ciml-lib"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/profileseq",
		URL:          []string{"https://bitbucket.org/regulatorygenomicsupf/profileseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/clonocalc-plot",
		URL:          []string{"https://bitbucket.org/clonosuite/clonocalc-plot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/cafe",
		URL:          []string{"https://bitbucket.org/cob87icw6z/cafe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/rsbiodiv",
		URL:          []string{"https://bitbucket.org/rsbiodiv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/disconica",
		URL:          []string{"https://bitbucket.org/masauburn/disconica"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/protwis",
		URL:          []string{"https://bitbucket.org/gpcr/protwis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/cypreact",
		URL:          []string{"https://bitbucket.org/leon_ti/cypreact"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/gsoa",
		URL:          []string{"https://bitbucket.org/srp33/gsoa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/caviarbf",
		URL:          []string{"https://bitbucket.org/wenan/caviarbf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/glycoprofileassigner",
		URL:          []string{"https://bitbucket.org/fergaljd/glycoprofileassigner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/pymonaco",
		URL:          []string{"https://bitbucket.org/hgugmradiofisica/pymonaco"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/rna-motifs",
		URL:          []string{"https://bitbucket.org/rogrro/rna_motifs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/metamarker",
		URL:          []string{"https://bitbucket.org/mkoohim/metamarker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/decaf",
		URL:          []string{"https://bitbucket.org/marta-sd/decaf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/qtlsearch",
		URL:          []string{"https://bitbucket.org/alex-warwickvesztrocy/qtlsearch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/md-analysis-with-matlab",
		URL:          []string{"https://bitbucket.org/niekdeklerk/md-analysis-with-matlab"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/mechrna",
		URL:          []string{"https://bitbucket.org/compbio/mechrna"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/bbcontacts",
		URL:          []string{"https://bitbucket.org/soedinglab/bbcontacts"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/ccmpred",
		URL:          []string{"https://bitbucket.org/soedinglab/ccmpred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/3depiloop",
		URL:          []string{"https://bitbucket.org/4dnucleome/3depiloop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/mugbas",
		URL:          []string{"https://bitbucket.org/capemaster/mugbas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/sensa",
		URL:          []string{"https://bitbucket.org/floettma/sensa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/visualmcmc",
		URL:          []string{"https://bitbucket.org/rhali/visualmcmc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/hapzipper",
		URL:          []string{"https://bitbucket.org/pchanda/hapzipper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/whatshap",
		URL:          []string{"https://bitbucket.org/whatshap/whatshap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/samu661/fsh/overview",
		URL:          []string{"https://bitbucket.org/samu661/fsh/overview"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/suppa",
		URL:          []string{"https://bitbucket.org/regulatorygenomicsupf/suppa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/imondb",
		URL:          []string{"https://bitbucket.org/proteinspector/imondb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/dyss",
		URL:          []string{"https://bitbucket.org/ban-m/dyss"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/virus-vg",
		URL:          []string{"https://bitbucket.org/jbaaijens/virus-vg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/glycome-analytics-platform",
		URL:          []string{"https://bitbucket.org/scientificomputing/glycome-analytics-platform"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/xglycscan",
		URL:          []string{"https://bitbucket.org/paiyetan/xglycscan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/dendsort",
		URL:          []string{"https://bitbucket.org/biovizleuven/dendsort"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bitbucket/trust",
		URL:          []string{"https://bitbucket.org/liulab/trust"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
}
