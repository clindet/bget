<img src="https://img.shields.io/badge/lifecycle-experimental-orange.svg" alt="Life cycle: experimental"> <a href="https://godoc.org/github.com/openbiox/bget"><img src="https://godoc.org/github.com/openbiox/bget?status.svg" alt="GoDoc"></a>

# bget

bget is an portable tool with several sub-commands to query bioinformatics APIs, data, databases and files. The Golang `http` library, `wget`, `curl`, `axel`, `git`, and `rsync` were supported as the query engine.

Supported types:

- Reference genomes
- Source code of bioinformatics tools
- Bioinformatics databases and files
- Papers material
- ......

Downstream tool:

- [bioctl](https://github.com/openbiox/bioctl): convert, format, and other functions
- [bioextr](https://github.com/openbiox/bioextr): text-mining functions

## Prerequisities

For website spider (optional):

- Headless Chrome is required for some of website with JavaScript driven render pages. For windows users, you may need to create an alias of Chrome to make [chromedp](https://github.com/chromedp/chromedp) work.

For raw sequencing data query (optional):

- [sra-tools](https://github.com/ncbi/sra-tools) for SRA and dbGAP database: MAC and Windows user `bget key base/sratools@2.9.6-1`, Linux user `bget key base/sratools`;
- [pyega3](https://github.com/EGA-archive/ega-download-client) for EGA database: `pip3 install pyega3`;
- [gdc-client](https://gdc.cancer.gov/access-data/gdc-data-transfer-tool) for GDC portal: `bget key base/gdc-client@1.4.0 -u`.

## Installation

```bash
# windows
wget https://github.com/openbiox/bget/releases/download/v0.2.4/bget.exe

# osx
wget https://github.com/openbiox/bget/releases/download/v0.2.4/bget_osx
mv bget_osx bget
chmod a+x bget

# linux
wget https://github.com/openbiox/bget/releases/download/v0.2.4/bget_linux64
mv bget_linux64 bget
chmod a+x bget

# get latest version
go get -u github.com/openbiox/bget
```

## Usage

Command line outputs see [here](https://github.com/openbiox/bget/blob/master/doc/cli.md)

### Query webiste API

`bget api` can be used to query serveral website APIs, such as PubMed, Datasetdataset2tools, and GDC portal website.

In addition, you can use the downstream tool [bioctl](https://github.com/openbiox/bioctl) to conduct the simple text-mining of PubMed abstract at the sentence level.

```bash
# NCBI eutils
# query pubmed
bget api ncbi -d pubmed -q B-ALL --format XML -e your_email@domain.com

# query pubmed and convert it to json format that also extract all URLs and calculate the words connections
bget api ncbi -q "Galectins control MTOR and AMPK in response to lysosomal damage to induce autophagy OR MTOR-independent autophagy induced by interrupted endoplasmic reticulum-mitochondrial Ca2+ communication: a dead end in cancer cells. OR The PARK10 gene USP24 is a negative regulator of autophagy and ULK1 protein stability OR Coordinate regulation of autophagy and the ubiquitin proteasome system by MTOR." | bioctl cvrt --xml2json pubmed -

# datasetdataset2tools API
# query canned analysis accession	, e.g. DCA00000060.
bget api dta -a DCA00000060
# query dataset accession number, e.g. GSE31106 
bget api dta -s GSE31106 | bioctl fmt --json-pretty -
# query via object type
bget api dta --type dataset | bioctl fmt --json-pretty --indent 2 -
# props of dataset accession, e.g. upregulated.
bget api dta -g upregulated | json2csv -o out.csv

# GDC portal API
# retrive projects meta info from GDC portal
bget api gdc -p
bget api gdc -p --json-pretty
bget api gdc -p -q TARGET-NBL --json-pretty
bget api gdc -p --format tsv > tcga_projects.tsv
bget api gdc -p --format csv > tcga_projects.csv
bget api gdc -p --from 1 --szie 2
# check GDC portal status (https://portal.gdc.cancer.gov/)
bget api gdc -s
# retrive cases info from GDC portal
bget api gdc -c
# retrive files info from GDC portal
bget api gdc -f
# retrive annotations info from GDC portal
bget api gdc -a
# query manifest for gdc-client
bget api gdc -m -q "5b2974ad-f932-499b-90a3-93577a9f0573,556e5e3f-0ab9-4b6c-aa62-c42f6a6cf20c" -o my_manifest.txt
bget api gdc -m -q "5b2974ad-f932-499b-90a3-93577a9f0573,556e5e3f-0ab9-4b6c-aa62-c42f6a6cf20c" > my_manifest.txt
bget api gdc -m -q "5b2974ad-f932-499b-90a3-93577a9f0573,556e5e3f-0ab9-4b6c-aa62-c42f6a6cf20c" -n
# query data that only support the samll filesize
bget api gdc -d -q "5b2974ad-f932-499b-90a3-93577a9f0573" -n

# clinicaltrials.gov API
# returns the date when the ClinicalTrials.gov dataset was posted.
bget api cligov --info-dat-vers
# returns the current version number of the ClinicalTrials.gov API
bget api cligov --info-api-vers
# returns detailed definitions.
bget api cligov --info-api-defs
# returns all available data elements for a single study record.
bget api cligov --info-study-struct
# returns all data elements.
bget api cligov --info-study-fields
# returns an annotated version of the Study Structure info URL.
bget api cligov --info-study-stat
# returns groups of weighted study fields, or "search areas"
bget api cligov --info-search-area
	
# query functions
bget api cligov -q heart+attack --full-studies --format json
bget api cligov -q heart+attack --fields NCTId,Condition,BriefTitle --study-fields
bget api cligov -q heart+attack --field Condition --field-values

# bio.tools API
# query item detail
bget api biots --tool signalp

# search item
bget api biots --name signalp
bget api biots --topic Proteomics
bget api biots --dtype 'Protein sequence'
bget api biots --dfmt FASTA
bget api biots --ofmt 'ClustalW format'
```

### Query DOI resources

`bget doi` can be used to query DOI resources from website and journals that the supported items are continuely increasing.

```bash
## query zendo website with 3 thread
bget doi 10.5281/zenodo.3363060 10.5281/zenodo.3357455 10.5281/zenodo.3351812 -t 3

## query fulltext of publications (proxy may needed)
bget doi 10.1016/j.devcel.2017.03.001 10.1016/j.stem.2019.07.009 10.1016/j.celrep.2018.03.072 -t 2

## query publications with supplementary files
bget doi 10.1038/s41586-019-1844-5 --suppl

```

We can query PDF of the manuscript via using Endnote or sci-hub. However, you can not easily get the supplementary files of scientific papers based on the two ways.

![doi demo](https://github.com/openbiox/bget/raw/master/doc/static/doi.gif)

Here, we are developing and sharing an open-source tool bget with `doi` subcommand to query supplementary files of scientific papers. The journals with high impact factors or those integrative publishers are a higher priority in our development plan, see [here](http://openbiox.github.io/bget/doi.html)

**Warn**: It is noted that we do not want to distribute any pirated resources or cause unnecessary network congestion. We hope this tool can provide an optional method to more easily query related files of scientific papers. Please use it in a non-invasive way (i.e. high concurrency, long continuous request). If you do not follow the policies of the relevant website (i.e. continuous download or limited copyright), you will lose the authorization to use this tool.

### Query files via alias key

`bget key` can be used to query a set of files via the alias key, such as bwa, samtools, reffa/defuse, and db/annovar.

```bash
# download bwa source (with task env info)
bget key bwa --verbose 2
# get all available keys
bget key -a
# in JSON format
bget key -a --format json
# view all bwa and samtools available tags in table
bget key bwa samtools -v
# view all bwa and samtools available tags in json
bget key bwa samtools -v --format json

# force download defuse reference (with task env info and save log to file)
bget key "reffa/defuse@GRCh38 #97" -t 10 -f
bget key reffa/defuse@GRCh38 release=97 -t 10 -f
# download annovar reference
bget key db/annovar@clinvar_20170501 db/annovar@clinvar_20180603 builder=hg38

bget key db/annovar -v --formt text
bget key db/annovar version='clinvar_20131105, clinvar_20140211, clinvar_20140303, clinvar_20140702, clinvar_20140902, clinvar_20140929, clinvar_20150330, clinvar_20150629, clinvar_20151201, clinvar_20160302, clinvar_20161128, clinvar_20170130, clinvar_20170501, clinvar_20170905, clinvar_20180603, avsnp150, avsnp147, avsnp144, avsnp142, avsnp138, cadd, caddgt10, caddgt20, cadd13, cadd13gt10, cadd13gt20, cg69, cg46, cosmic70, cosmic68wgs, cosmic68, cosmic67wgs, cosmic67, cosmic65, cosmic64, dbnsfp35a, dbnsfp33a, dbnsfp31a_interpro, dbnsfp30a, dbscsnv11, eigen, esp6500siv2_ea, esp6500siv2_aa, esp6500siv2_all, exac03nontcga, exac03nonpsych, exac03, fathmm, gerp++gt2, gme, gnomad_exome, gnomad_genome, gwava, hrcr1, icgc21, intervar_20170202, kaviar_20150923, ljb26_all, mcap, mitimpact2, mitimpact24, nci60, popfreq_max_20150413, popfreq_all_20150413, revel, regsnpintron' builder=hg19 -t 10 -f
```

### Query FASTQ/CEL files from GEO/SRA/EGA/dbGAP/GDC

`bget seq` can be used to query files from [Gene Expression Omnibus (GEO)](https://www.ncbi.nlm.nih.gov/geo), [Sequence Read Archive (SRA)](https://www.ncbi.nlm.nih.gov/sra/), and [GDC Data Portal](https://portal.gdc.cancer.gov/).

```bash
# download files from SRA databaes using prefetch
bget seq ERR3324530 SRR544879

# download files from GEO databaes, auto download SRA acc list and run info
bget seq GSE23543 GSM1098572 -t 2

# download files from dbGap database using krt files
bget seq dbgap.krt using prefetch

# download dataset from EGA databaes using pyega3
bget seq EGAD00001000951

# download file from EGA databaes using pyega3
bget seq EGAF00000585895

# download TCGA files using file id using gdc-client
bget seq b7670817-9d6b-494e-9e22-8494e2fd430d

# download TCGA files using manifest files using gdc-client
# split for parallel
split -a 3 --additional-suffix=.txt -l 100 gdc_manifest.2019-08-23-TCGA.txt -d
for i in x*.txt
do
  head -n 1 x000.txt > ${i}.tmp && cat ${i} >> ${i}.tmp &&mv ${i}.tmp ${i}
done
sed -i '1d' x000.txt
bget seq *.txt -t 5

# support auto (if you do not have *.krt, TCGA manifest, please not include it for test)
bget seq SRR544879 GSE23543 EGAD00001000951 b7670817-9d6b-494e-9e22-8494e2fd430d dbgap.krt *.txt -t 5
```

### Query URLs

`bget url` can be used to query files using URLs.

```bash
urls="https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe,http://download.oray.com/pgy/windows/PgyVPN_4.1.0.21693.exe,https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe" && echo $urls | tr "," "\n"> /tmp/urls.list

bget url ${urls}
bget url https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe --save-log
bget url ${urls} -t 3 -o /tmp/download -f -g wget --save-log --verbose 2
bget url ${urls} -t 2 -o /tmp/download --save-log --verbose 2
bget url ${urls} -t 3 -o /tmp/download -g wget --ignore
bget url -l /tmp/urls.list -o /tmp/download -f -t 3

# query github repo (support assets files)
bget url Miachol/github_demo --github
bget url PapenfussLab/gridss openbiox/bget --with-github-assets -t 5 --github
bget url PapenfussLab/gridss openbiox/bget --only-github-assets -t 5 --github
bget url PapenfussLab/gridss openbiox/bget --with-github-assets --with-assets-versions v2.7.2,v0.1.3 -t 5 --github
```

## Maintainer

- [@Jianfeng](https://github.com/Miachol)

## License

Academic Free License version 3.0

