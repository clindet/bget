<img src="https://img.shields.io/badge/lifecycle-experimental-orange.svg" alt="Life cycle: experimental"> [![GoDoc](https://godoc.org/github.com/openbiox/bget/bapi?status.svg)](https://godoc.org/github.com/openbiox/bget/bapi)

# bapi

[bapi](https://github.cmo/openbiox/bget/bapi) is an open source and cross-platform command line tool to query the APIs of bioinformatics related websites and conduct the text-mining of PubMed abstract.

Now, it can be used to query the APIs of [dataset2tools](http://amp.pharm.mssm.edu/datasets2tools/search), [GDC Portal](https://portal.gdc.cancer.gov/), and [NCBI](https://www.ncbi.nlm.nih.gov/).

In the future, we will include more useful APIs of bioinformatics, such as [BioMart](http://www.biomart.org/), [Fairsharing](https://fairsharing.org/), [EGA-Archive](https://ega-archive.org/), and your request.

In addition, you can use the `bapi` to conduct the simple text-mining of PubMed abstract at the sentence level.

- support calculate correlations between any keywords at the sentence level (PubMed)
- support extract URLs from Pubmed abstract
- support convert PubMed abstract from XML => JSON
- support to prettify JSON stream
- support convert key-value JSON to slice JSON

## Installation

```bash
# Binary
Linux: https://github.com/openbiox/bapi/releases/download/v0.1.0/bapi_linux64
Mac: https://github.com/openbiox/bapi/releases/download/v0.1.0/bapi_osx
Win: https://github.com/openbiox/bapi/releases/download/v0.1.0/bapi.exe

# For windows user, we recommend to run bapi at git bash mode (support awk and other GNU tools)
# See more: https://git-scm.com/

# Golang developer
go get -u github.com/openbiox/bapi
```

## Usage

You can direct to type `bapi` to see all subcommands and global flags of `bapi`. Now, this project is unstable, the name or flags may changed in the future.

```bash
Query bioinformatics website APIs. More see here https://github.com/openbiox/bget/bapi.

Usage:
  bapi [flags]
  bapi [command]

Available Commands:
  dta         Query dataset2tools website APIs: datasets (d), tools (t), and canned analysis (a).
  fmt         A set of file format (fmt) command of bapi.
  gdc         Query GDC portal website APIs.
  help        Help about any command
  ncbi        Query ncbi website APIs.

Flags:
  -e, --email string             Email specifies the email address to be sent to the server (NCBI website is required). (default "your_email@domain.com")
      --format string            Rettype specifies the format of the returned data (CSV, TSV, JSON for gdc; XML/TEXT for ncbi).
      --from int                 Parameters of API control the start item of retrived data. (default -1)
  -h, --help                     help for bapi
  -q, --query string             Query specifies the search query for record retrieval (required).
      --quiet                    No log output.
  -r, --retries int              Retry specifies the number of attempts to retrieve the data. (default 5)
      --retries-sleep-time int   Sleep time after one retry. (default 5)
      --size int                 Parameters of API control the lenth of retrived data. Default is auto determined. (default -1)
      --timeout int              Set the timeout of per request. (default 35)
      --version                  version for bapi

Use "bapi [command] --help" for more information about a command.
```

**GDC query:**

```bash
Query GDC portal APIs. More see here https://github.com/openbiox/bget/bapi.

Usage:
  bapi gdc [flags]

Examples:
  # Query projects in GDC portal
  bapi gdc -p
  # Query projects in GDC portal in pretty JSON
  bapi gdc -p --json-pretty
  # Query TARGET-NBL project in GDC portal in pretty JSON
  bapi gdc -p -q TARGET-NBL --json-pretty
  # Query projects in TSV format
  bapi gdc -p --format TSV > tcga_projects.tsv
  # Query projects in CSV format
  bapi gdc -p --format CSV > tcga_projects.csv
  # Query projects via control the start item and total item
  bapi gdc -p --from 1 --szie 2
  # See status of GDC portal
  bapi gdc -s
  # Retrive cases info from GDC portal
  bapi gdc -c
  # Retrive files info from GDC portal
  bapi gdc -f
  # Retrive annotations info from GDC portal. 
  bapi gdc -a

  // Download manifest for gdc-client
  bapi gdc -m -q "5b2974ad-f932-499b-90a3-93577a9f0573,556e5e3f-0ab9-4b6c-aa62-c42f6a6cf20c" -o my_manifest.txt
  bapi gdc -m -q "5b2974ad-f932-499b-90a3-93577a9f0573,556e5e3f-0ab9-4b6c-aa62-c42f6a6cf20c" > my_manifest.txt
  bapi gdc -m -q "5b2974ad-f932-499b-90a3-93577a9f0573,556e5e3f-0ab9-4b6c-aa62-c42f6a6cf20c" -n

  // Download data
  bapi gdc -d -q "5b2974ad-f932-499b-90a3-93577a9f0573" -n
```

**Query NCBI:**

```bash
Query ncbi website APIs. More see here https://github.com/openbiox/bget/bapi.

Usage:
  bapi ncbi [flags]

Examples:
  # Search Pubmed database with "B-ALL" query and returns XML format file
  bapi ncbi -d pubmed -q B-ALL --format XML -e your_email@domain.com
  # Split returns XML files
  bapi ncbi -q "RNA-seq and bioinformatics[journal]" -e "your_email@domain.com" -m 100 | awk '/<[?]xml version="1.0" [?]>/{close(f); f="abstract.http.XML.tmp" ++c;next} {print>f;}'

  # Calculate the words corelations at sentence level
  k="algorithm, tool, model, pipleline, method, database, workflow, dataset, bioinformatics, sequencing, http, github.com, gitlab.com, bitbucket.org, RNA-Seq, DNA, profile, landscape"
  bapi ncbi --xml2json pubmed abstract.http.XML.tmp* -k "${k}" --call-cor | sed 's;}{;,;g' > final.json

  # Support to input the pipe stream with '-' flag
  bapi ncbi -q "Galectins control MTOR and AMPK in response to lysosomal damage to induce autophagy OR MTOR-independent autophagy induced by interrupted endoplasmic reticulum-mitochondrial Ca2+ communication: a dead end in cancer cells. OR The PARK10 gene USP24 is a negative regulator of autophagy and ULK1 protein stability OR Coordinate regulation of autophagy and the ubiquitin proteasome system by MTOR." | bapi ncbi --xml2json pubmed -k "MAPK, MTOR, autophagy" --call-cor - | sed 's;}{;,;g' | bapi fmt --json-to-slice - > final.json

  # json2csv install from https://github.com/zemirco/json2csv#readme
  json2csv -i final.json -o final.csv  

```

**Query dataset2tools:**

```bash
Query dataset2tools APIs. More see here https://github.com/openbiox/bget/bapi.

Usage:
  bapi dta [flags]

Examples:
  # retrive canned_analysis
  bapi dta -a DCA00000060
  # retrive dataset in pretty JSON format
  bapi dta -s GSE31106 | bapi fmt --json-pretty -
  # pretty JSON format with indent control
  bapi dta --type dataset | bapi fmt --json-pretty --indent 2 -
  # retrive geneset
  bapi dta -g upregulated | json2csv -o out.csv
```

## Maintainer

- [@Jianfeng](https://github.com/Miachol)

## License

Academic Free License version 3.0
