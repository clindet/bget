Main interface
--------------

    bget -h

    #> Lightweight downloader for bioinformatics data, databases and files (under development). It will provides a simple and parallelized method to access various bioinformatics resoures. More see here https://github.com/openbiox/bget.
    #> 
    #> Usage:
    #>   bget [flags]
    #>   bget [command]
    #> 
    #> Available Commands:
    #>   api         Query bioinformatics website APIs.
    #>   doi         Can be used to access files via DOI.
    #>   fmt         A set of format (fmt) command.
    #>   help        Help about any command
    #>   key         Can be used to access URLs via a key string.
    #>   seq         Can be used to access sequence data via unique id (dbGAP and EGA) or manifest files (TCGA).
    #>   url         Can be used to access URLs via Golang http, wget, curl, axel and git, and rsync.
    #> 
    #> Flags:
    #>       --clean     Remove _download and _log in current dir.
    #>   -h, --help      help for bget
    #>       --version   version for bget
    #> 
    #> Use "bget [command] --help" for more information about a command.

bget api
--------

    bget api -h
    #> Query bioinformatics website APIs.
    #> 
    #> Usage:
    #>   bget api [flags]
    #>   bget api [command]
    #> 
    #> Available Commands:
    #>   dta         Query dataset2tools website APIs: datasets (d), tools (t), and canned analysis (a).
    #>   gdc         Query GDC portal website APIs.
    #>   ncbi        Query ncbi website APIs.
    #> 
    #> Flags:
    #>   -e, --email string             Email specifies the email address to be sent to the server (NCBI website is required). (default "your_email@domain.com")
    #>       --format string            Rettype specifies the format of the returned data (CSV, TSV, JSON for gdc; XML/TEXT for ncbi).
    #>       --from int                 Parameters of API control the start item of retrived data. (default -1)
    #>   -h, --help                     help for api
    #>       --indent int               Control the indent of output json files. (default 4)
    #>       --json-pretty              Pretty json files.
    #>   -q, --query string             Query specifies the search query for record retrieval (required).
    #>       --quiet                    No log output.
    #>   -r, --retries int              Retry specifies the number of attempts to retrieve the data. (default 5)
    #>       --retries-sleep-time int   Sleep time after one retry. (default 5)
    #>       --size int                 Parameters of API control the lenth of retrived data. Default is auto determined. (default -1)
    #>       --sort-keys                Control wheather to sort JSON key.
    #>       --timeout int              Set the timeout of per request. (default 35)
    #>       --version                  version for api
    #> 
    #> Global Flags:
    #>       --clean   Remove _download and _log in current dir.
    #> 
    #> Use "bget api [command] --help" for more information about a command.

    bget api ncbi -h
    #> Query ncbi website APIs.
    #> 
    #> Usage:
    #>   bget api ncbi [flags]
    #> 
    #> Examples:
    #>   # query pubmed with 'B-ALL'
    #>   bget api ncbi -d pubmed -q B-ALL --format XML -e your_email@domain.com
    #> 
    #>   # query pubmed and convert it to json format that also extract all URLs and calculate the words connections
    #>   bget api ncbi -q "Galectins control MTOR and AMPK in response to lysosomal damage to induce autophagy OR MTOR-independent autophagy induced by interrupted endoplasmic reticulum-mitochondrial Ca2+ communication: a dead end in cancer cells. OR The PARK10 gene USP24 is a negative regulator of autophagy and ULK1 protein stability OR Coordinate regulation of autophagy and the ubiquitin proteasome system by MTOR." | bget api ncbi --xml2json pubmed -k "MAPK, MTOR, autophagy" --call-cor - | sed 's;}{;,;g' | bget fmt --json-to-slice - > final.json
    #> 
    #>   # query larger items
    #>   k="algorithm, tool, model, pipleline, method, database, workflow, dataset, bioinformatics, sequencing, http, github.com, gitlab.com, bitbucket.org, RNA-Seq, DNA, profile, landscape"
    #>   bget api ncbi -q "dataset and RNA-seq and bioinformatics[journal]" -e "your_email@domain.com" -m 20 | awk '/<[?]xml version="1.0" [?]>/{close(f); f="abstract.http.XML.tmp" ++c;next} {print>f;}' && bget api ncbi --xml2json pubmed abstract.http.XML.tmp* -k "${k}" --call-cor -t 11 | sed 's;}{;,;g' > final.json
    #> 
    #> Flags:
    #>       --call-cor          Wheather to calculate the corelated keywords, and return the sentence contains >=2 keywords.
    #>   -d, --db string         Db specifies the database to search (default "pubmed")
    #>   -h, --help              help for ncbi
    #>   -k, --keywords string   Keywords to extracted from abstract. (default "algorithm, tool, model, pipleline, method, database, workflow, dataset, bioinformatics, sequencing, http, github.com, gitlab.com, bitbucket.org")
    #>   -o, --outfn string      Out specifies destination of the returned data (default to stdout).
    #>   -m, --per-size int      Retmax specifies the number of records to be retrieved per request. (default 100)
    #>   -t, --thread int        Thread to process. (default 2)
    #>       --xml2json string   Convert XML files to json [e.g. pubmed].
    #> 
    #> Global Flags:
    #>       --clean                    Remove _download and _log in current dir.
    #>   -e, --email string             Email specifies the email address to be sent to the server (NCBI website is required). (default "your_email@domain.com")
    #>       --format string            Rettype specifies the format of the returned data (CSV, TSV, JSON for gdc; XML/TEXT for ncbi).
    #>       --from int                 Parameters of API control the start item of retrived data. (default -1)
    #>       --indent int               Control the indent of output json files. (default 4)
    #>       --json-pretty              Pretty json files.
    #>   -q, --query string             Query specifies the search query for record retrieval (required).
    #>       --quiet                    No log output.
    #>   -r, --retries int              Retry specifies the number of attempts to retrieve the data. (default 5)
    #>       --retries-sleep-time int   Sleep time after one retry. (default 5)
    #>       --size int                 Parameters of API control the lenth of retrived data. Default is auto determined. (default -1)
    #>       --sort-keys                Control wheather to sort JSON key.
    #>       --timeout int              Set the timeout of per request. (default 35)

    bget api gdc -h
    #> Query GDC portal website APIs.
    #> 
    #> Usage:
    #>   bget api gdc [flags]
    #> 
    #> Examples:
    #>   # retrive projects meta info from GDC portal
    #>   bget api gdc -p
    #>   bget api gdc -p --json-pretty
    #>   bget api gdc -p -q TARGET-NBL --json-pretty
    #>   bget api gdc -p --format tsv > tcga_projects.tsv
    #>   bget api gdc -p --format csv > tcga_projects.csv
    #>   bget api gdc -p --from 1 --szie 2
    #>   # check GDC portal status (https://portal.gdc.cancer.gov/)
    #>   bget api gdc -s
    #>   # retrive cases info from GDC portal
    #>   bget api gdc -c
    #>   # retrive files info from GDC portal
    #>   bget api gdc -f
    #>   # retrive annotations info from GDC portal
    #>   bget api gdc -a
    #>   # query manifest for gdc-client
    #>   bget api gdc -m -q "5b2974ad-f932-499b-90a3-93577a9f0573,556e5e3f-0ab9-4b6c-aa62-c42f6a6cf20c" -o my_manifest.txt
    #>   bget api gdc -m -q "5b2974ad-f932-499b-90a3-93577a9f0573,556e5e3f-0ab9-4b6c-aa62-c42f6a6cf20c" > my_manifest.txt
    #>   bget api gdc -m -q "5b2974ad-f932-499b-90a3-93577a9f0573,556e5e3f-0ab9-4b6c-aa62-c42f6a6cf20c" -n
    #>   # query data that only support the samll filesize
    #>   bget api gdc -d -q "5b2974ad-f932-499b-90a3-93577a9f0573" -n
    #> 
    #> Flags:
    #>   -a, --annotations     Retrive annotations info from GDC portal.
    #>   -c, --cases           Retrive cases info from GDC portal.
    #>   -d, --data            Retrive /data from GDC portal.
    #>       --fields string   Fields parameters.
    #>   -f, --files           Retrive files info from GDC portal.
    #>       --filter string   Retrive data with GDC filter.
    #>   -h, --help            help for gdc
    #>   -l, --legacy          Use legacy API of GDC portal.
    #>   -m, --manifest        Retrive /manifest data from GDC portal.
    #>   -o, --outfn string    Out specifies destination of the returned data (default to stdout).
    #>   -p, --projects        Retrive projects meta info from GDC portal.
    #>   -n, --remote-name     Use remote defined filename.
    #>       --slicing         Retrive BAM slicing from GDC portal.
    #>       --sort string     Sort parameters.
    #>   -s, --status          Check GDC portal status (https://portal.gdc.cancer.gov/).
    #>       --token string    Token to access GDC.
    #> 
    #> Global Flags:
    #>       --clean                    Remove _download and _log in current dir.
    #>   -e, --email string             Email specifies the email address to be sent to the server (NCBI website is required). (default "your_email@domain.com")
    #>       --format string            Rettype specifies the format of the returned data (CSV, TSV, JSON for gdc; XML/TEXT for ncbi).
    #>       --from int                 Parameters of API control the start item of retrived data. (default -1)
    #>       --indent int               Control the indent of output json files. (default 4)
    #>       --json-pretty              Pretty json files.
    #>   -q, --query string             Query specifies the search query for record retrieval (required).
    #>       --quiet                    No log output.
    #>   -r, --retries int              Retry specifies the number of attempts to retrieve the data. (default 5)
    #>       --retries-sleep-time int   Sleep time after one retry. (default 5)
    #>       --size int                 Parameters of API control the lenth of retrived data. Default is auto determined. (default -1)
    #>       --sort-keys                Control wheather to sort JSON key.
    #>       --timeout int              Set the timeout of per request. (default 35)

    bget api dta -h
    #> Query dataset2tools website APIs: datasets (d), tools (t), and canned analysis (a).
    #> 
    #> Usage:
    #>   bget api dta [flags]
    #> 
    #> Examples:
    #>   # query canned analysis accession  , e.g. DCA00000060.
    #>   bget api dta -a DCA00000060
    #>   # query dataset accession number, e.g. GSE31106 
    #>   bget api dta -s GSE31106 | bget fmt --json-pretty -
    #>   # query via object type
    #>   bget api dta --type dataset | bget fmt --json-pretty --indent 2 -
    #>   # props of dataset accession, e.g. upregulated.
    #>   bget api dta -g upregulated | json2csv -o out.csv
    #> 
    #> Flags:
    #>   -a, --analysis-acc string   Canned analysis accession  , e.g. DCA00000060.
    #>   -s, --dataset-acc string    Dataset accession number, e.g. GSE31106.
    #>   -d, --disease string        Disease name, e.g. prostate cancer
    #>   -g, --geneset string        With dataset accession, e.g. upregulated.
    #>   -h, --help                  help for dta
    #>   -o, --outfn string          Out specifies destination of the returned data (default to stdout).
    #>   -t, --tool string           Tool name, e.g. bwa.
    #>       --type string           Object type [tool, dataset, canned_analysis].
    #> 
    #> Global Flags:
    #>       --clean                    Remove _download and _log in current dir.
    #>   -e, --email string             Email specifies the email address to be sent to the server (NCBI website is required). (default "your_email@domain.com")
    #>       --format string            Rettype specifies the format of the returned data (CSV, TSV, JSON for gdc; XML/TEXT for ncbi).
    #>       --from int                 Parameters of API control the start item of retrived data. (default -1)
    #>       --indent int               Control the indent of output json files. (default 4)
    #>       --json-pretty              Pretty json files.
    #>   -q, --query string             Query specifies the search query for record retrieval (required).
    #>       --quiet                    No log output.
    #>   -r, --retries int              Retry specifies the number of attempts to retrieve the data. (default 5)
    #>       --retries-sleep-time int   Sleep time after one retry. (default 5)
    #>       --size int                 Parameters of API control the lenth of retrived data. Default is auto determined. (default -1)
    #>       --sort-keys                Control wheather to sort JSON key.
    #>       --timeout int              Set the timeout of per request. (default 35)

bget doi
--------

    bget doi -h
    #> Can be used to access files via DOI. More see here https://github.com/openbiox/bget.
    #> 
    #> Usage:
    #>   bget doi [doi1 doi2 doi3...] [flags]
    #> 
    #> Examples:
    #>   bget doi 10.5281/zenodo.3363060 10.5281/zenodo.3357455 10.5281/zenodo.3351812 -t 3
    #>   bget doi 10.1016/j.devcel.2017.03.001 10.1016/j.stem.2019.07.009 10.1016/j.celrep.2018.03.072 -t 2
    #> 
    #>   bapi ncbi -q '((The PARK10 gene USP24 is a negative regulator of autophagy and ULK1 protein stability[Title]) OR Coordinate regulation of autophagy and the ubiquitin proteasome system by MTOR[Title])' -o titleSearch.XML
    #>   dois=`bapi ncbi --xml2json pubmed titleSearch.XML |grep Doi| tr -d ' ,(Doi:)"'`
    #>   echo ${dois}
    #>   bget doi ${dois}
    #>   bget doi 10.1080/15548627.2018.1505155 --proxy http://username:password@hostname:port
    #> 
    #> Flags:
    #>   -g, --engine string            Point the download engine: go-http, wget, curl, axel, git, and rsync. (default "go-http")
    #>   -e, --extra-cmd string         Extra flags and values pass to internal CMDs
    #>       --full-text                Access full text. (default true)
    #>   -h, --help                     help for doi
    #>       --ignore                   Contine to download and skip the check of existed files.
    #>   -l, --list-file string         A file contains dois for download.
    #>       --log-dir string           Log dir. (default "_log")
    #>   -m, --mirror string            Set the mirror of resources.
    #>   -o, --outdir string            Set the download dir.
    #>   -f, --overwrite                Logical indicating that whether to overwrite existing files.
    #>       --pmc                      Try PMC database.
    #>       --proxy string             HTTP proxy to download.
    #>   -q, --quiet                    No output.
    #>   -n, --remote-name              Use remote defined filename.
    #>   -r, --retries int              Retry specifies the number of attempts to retrieve the data. (default 5)
    #>       --retries-sleep-time int   Sleep time after one retry. (default 5)
    #>       --save-log                 Save download log to local file]. (default true)
    #>   -s, --seperator string         Optional 'url1{seperator}url2' for multiple keys, urls, or seqs. (default ",")
    #>       --suppl                    Access supplementary files.
    #>       --task-id string           Task ID (random). (default "kdgbobvax225938")
    #>   -t, --thread int               Concurrency download thread. (default 1)
    #>       --thread-axel int          Set the thread of axel. (default 5)
    #>       --timeout int              Set the timeout of per request. (default 35)
    #>       --universe                 Try universe spider. (default true)
    #> 
    #> Global Flags:
    #>       --clean   Remove _download and _log in current dir.

### bget seq

    bget seq -h
    #> Can be used to access sequence data via unique id or manifest files. More see here https://github.com/openbiox/bget.
    #> 
    #> Usage:
    #>   bget seq [id1 id2 id3... | manifest1 manifest2 manifest3...] [flags]
    #> 
    #> Examples:
    #>   bget seq ERR3324530 SRR544879 # download files from SRA databaes
    #>   bget seq GSE23543 GSM1098572 -t 2 # download files from GEO databaes (auto download SRA acc list and run info)
    #>   bget seq dbgap.krt # download files from dbGap database using krt files
    #>   bget seq EGAD00001000951 # download dataset from EGA databaes
    #>   bget seq EGAF00000585895 # download file from EGA databaes
    #>  
    #>   # download TCGA files using file id
    #>   bget seq b7670817-9d6b-494e-9e22-8494e2fd430d
    #> 
    #>   # download TCGA files using manifest files
    #>   # split for parallel
    #>   split -a 3 --additional-suffix=.txt -l 100 gdc_manifest.2019-08-23-TCGA.txt -d
    #>   for i in x*.txt
    #>   do
    #>     head -n 1 x000.txt > ${i}.tmp && cat ${i} >> ${i}.tmp &&mv ${i}.tmp ${i}
    #>   done
    #>   sed -i '1d' x000.txt
    #>   bget seq *.txt -t 5
    #> 
    #>   # support auto (if you do not have *.krt, TCGA manifest, please not include it for test)
    #>   bget seq SRR544879 GSE23543 EGAD00001000951 b7670817-9d6b-494e-9e22-8494e2fd430d dbgap.krt *.txt -t 5
    #> 
    #> Flags:
    #>   -g, --engine string            Point the download engine: go-http, wget, curl, axel, git, and rsync. (default "go-http")
    #>   -e, --extra-cmd string         Extra flags and values pass to internal CMDs
    #>   -h, --help                     help for seq
    #>       --ignore                   Contine to download and skip the check of existed files.
    #>   -l, --list-file string         A file contains accession ids for download.
    #>       --log-dir string           Log dir. (default "/Users/apple/Documents/github/bget/doc/_log")
    #>   -m, --mirror string            Set the mirror of resources.
    #>   -o, --outdir string            Set the download dir. (default "/Users/apple/Documents/github/bget/doc")
    #>   -f, --overwrite                Logical indicating that whether to overwrite existing files.
    #>       --proxy string             HTTP proxy to download.
    #>       --query-gpl                Wheather fetch GPL files from GEO database.
    #>   -q, --quiet                    No output.
    #>   -n, --remote-name              Use remote defined filename.
    #>   -r, --retries int              Retry specifies the number of attempts to retrieve the data. (default 5)
    #>       --retries-sleep-time int   Sleep time after one retry. (default 5)
    #>       --save-log                 Save download log to local file]. (default true)
    #>   -s, --seperator string         Optional 'url1{seperator}url2' for multiple keys, urls, or seqs. (default ",")
    #>       --task-id string           Task ID (random). (default "mb9x24scjkjc964")
    #>   -t, --thread int               Concurrency download thread. (default 1)
    #>       --thread-axel int          Set the thread of axel. (default 5)
    #>       --timeout int              Set the timeout of per request. (default 35)
    #>       --token-file-ega string    Credential file to access EGA archive files, {"username": "{your_user_name}", 
    #>                                    "password": "{your_password}","client_secret":"AMenuDLjVdVo4BSwi0QD54LL6NeVDEZRzEQUJ7h
    #>                                      JOM3g4imDZBHHX0hNfKHPeQIGkskhtCmqAJtt_jm7EKq-rWw"}.
    #>       --token-gdc string         Token to access TCGA portal files.
    #> 
    #> Global Flags:
    #>       --clean   Remove _download and _log in current dir.

### bget url

    bget url -h
    #> Can be used to access URLs via Golang http, wget, curl, axel and git, and rsync. More see here https://github.com/openbiox/bget.
    #> 
    #> Usage:
    #>   bget url [url1 url2 url3...] [flags]
    #> 
    #> Examples:
    #>   urls="https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe,http://download.oray.com/pgy/windows/PgyVPN_4.1.0.21693.exe,https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe" && echo $urls | tr "," "\n"> /tmp/urls.list
    #> 
    #>   bget url ${urls}
    #>   bget url https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe
    #>   bget url ${urls} -t 2 -o /tmp/download
    #>   bget url ${urls} -t 3 -o /tmp/download -f -g wget
    #>   bget url ${urls} -t 3 -o /tmp/download -g wget --ignore
    #>   bget url -l /tmp/urls.list -o /tmp/download -f -t 3
    #> 
    #>   bget url Miachol/github_demo --github
    #>   bget url PapenfussLab/gridss openbiox/bget --with-github-assets -t 5 --github
    #>   bget url PapenfussLab/gridss openbiox/bget --only-github-assets -t 5 --github
    #>   bget url PapenfussLab/gridss openbiox/bget --with-github-assets --with-assets-versions v2.7.2,v0.1.3 -t 5 --github
    #> 
    #> Flags:
    #>   -g, --engine string                   Point the download engine: go-http, wget, curl, axel, git, and rsync. (default "go-http")
    #>   -e, --extra-cmd string                Extra flags and values pass to internal CMDs
    #>       --github                          GitHub mode.
    #>       --github-assets-versions string   Required to get specific tagname of github assets (e.g. v2.7.1,v1.0.0).
    #>   -h, --help                            help for url
    #>       --ignore                          Contine to download and skip the check of existed files.
    #>   -l, --list-file string                A file contains urls for download.
    #>       --log-dir string                  Log dir. (default "/Users/apple/Documents/github/bget/doc/_log")
    #>   -m, --mirror string                   Set the mirror of resources.
    #>       --only-github-assets              Logical indicating that whether to only download github repo assets files.
    #>   -o, --outdir string                   Set the download dir. (default "/Users/apple/Documents/github/bget/doc")
    #>   -f, --overwrite                       Logical indicating that whether to overwrite existing files.
    #>       --proxy string                    HTTP proxy to download.
    #>   -q, --quiet                           No output.
    #>   -n, --remote-name                     Use remote defined filename.
    #>   -r, --retries int                     Retry specifies the number of attempts to retrieve the data. (default 5)
    #>       --retries-sleep-time int          Sleep time after one retry. (default 5)
    #>       --save-log                        Save download log to local file]. (default true)
    #>   -s, --seperator string                Optional 'url1{seperator}url2' for multiple keys, urls, or seqs. (default ",")
    #>       --task-id string                  Task ID (random). (default "4ey91oi13565u48")
    #>   -t, --thread int                      Concurrency download thread. (default 1)
    #>       --thread-axel int                 Set the thread of axel. (default 5)
    #>       --timeout int                     Set the timeout of per request. (default 35)
    #>   -u, --uncompress                      Uncompress download files for .zip, .tar.gz, and .gz suffix files.
    #>       --with-github-assets              Logical indicating that whether to download associated assets files of github repo.
    #> 
    #> Global Flags:
    #>       --clean   Remove _download and _log in current dir.

### bget key

    bget key -h
    #> Can be used to access URLs via a key string. e.g. 'item' or 'item@version #releaseVersion', : bwa, reffa-defuse@GRCh38 #97. More see here https://github.com/openbiox/bget.
    #> 
    #> Usage:
    #>   bget key [key1 key2 key3...] [flags]
    #> 
    #> Examples:
    #>   bget key bwa
    #>   bget key -a // get all available keys
    #>   bget key samtools -v // view all samtools available versions in table
    #>   bget key samtools -v --formt json // view all samtools available versions in JSON
    #>   bget key "reffa/defuse@GRCh38 #97" -t 10 -f
    #>   bget key reffa/defuse@GRCh38 release=97 -t 10 -f
    #>   bget key db/annovar@clinvar_20170501 db/annovar@clinvar_20180603 builder=hg38
    #> 
    #>   bget key db/annovar -v --out-text
    #>   bget key db/annovar version='clinvar_20131105, clinvar_20140211, clinvar_20140303, clinvar_20140702, clinvar_20140902, clinvar_20140929, clinvar_20150330, clinvar_20150629, clinvar_20151201, clinvar_20160302, clinvar_20161128, clinvar_20170130, clinvar_20170501, clinvar_20170905, clinvar_20180603, avsnp150, avsnp147, avsnp144, avsnp142, avsnp138, cadd, caddgt10, caddgt20, cadd13, cadd13gt10, cadd13gt20, cg69, cg46, cosmic70, cosmic68wgs, cosmic68, cosmic67wgs, cosmic67, cosmic65, cosmic64, dbnsfp35a, dbnsfp33a, dbnsfp31a_interpro, dbnsfp30a, dbscsnv11, eigen, esp6500siv2_ea, esp6500siv2_aa, esp6500siv2_all, exac03nontcga, exac03nonpsych, exac03, fathmm, gerp++gt2, gme, gnomad_exome, gnomad_genome, gwava, hrcr1, icgc21, intervar_20170202, kaviar_20150923, ljb26_all, mcap, mitimpact2, mitimpact24, nci60, popfreq_max_20150413, popfreq_all_20150413, revel, regsnpintron' builder=hg19 -t 10 -f
    #> 
    #> Flags:
    #>       --autopath                 Logical indicating that whether to create subdir in download dir: e.g. reffa/{{key}}/
    #>   -g, --engine string            Point the download engine: go-http, wget, curl, axel, git, and rsync. (default "go-http")
    #>   -e, --extra-cmd string         Extra flags and values pass to internal CMDs
    #>       --format string            Output format (text, json, table)
    #>   -h, --help                     help for key
    #>       --ignore                   Contine to download and skip the check of existed files.
    #>   -a, --keys-all                 Show all available string key can be download.
    #>   -l, --list-file string         A file contains keys for download.
    #>       --log-dir string           Log dir. (default "/Users/apple/Documents/github/bget/doc/_log")
    #>   -m, --mirror string            Set the mirror of resources.
    #>   -o, --outdir string            Set the download dir. (default "/Users/apple/Documents/github/bget/doc")
    #>   -f, --overwrite                Logical indicating that whether to overwrite existing files.
    #>       --proxy string             HTTP proxy to download.
    #>   -q, --quiet                    No output.
    #>   -n, --remote-name              Use remote defined filename.
    #>   -r, --retries int              Retry specifies the number of attempts to retrieve the data. (default 5)
    #>       --retries-sleep-time int   Sleep time after one retry. (default 5)
    #>       --save-log                 Save download log to local file]. (default true)
    #>   -s, --seperator string         Optional 'url1{seperator}url2' for multiple keys, urls, or seqs. (default ",")
    #>   -v, --show-versions            Show all available versions of key.
    #>       --task-id string           Task ID (random). (default "if9vsdch2ug9z8h")
    #>   -t, --thread int               Concurrency download thread. (default 1)
    #>       --thread-axel int          Set the thread of axel. (default 5)
    #>       --timeout int              Set the timeout of per request. (default 35)
    #>   -u, --uncompress               Uncompress download files for .zip, .tar.gz, and .gz suffix files.
    #>       --with-assets              Logical indicating that whether to download associated assets files.
    #> 
    #> Global Flags:
    #>       --clean   Remove _download and _log in current dir.

    ## show all supported items
    bget key -a
    #> +++
    #> | 3dchromatin-replicateqc                                                 | abyss                                                           |
    #> | advntr                                                                  | agfusion                                                        |
    #> | aligner/blast                                                           | aligner/blat                                                    |
    #> | anchor                                                                  | anno/vcfanno                                                    |
    #> | annovar                                                                 | annovarr                                                        |
    #> | app/babun                                                               | app/cmder                                                       |
    #> | app/iontorrent-suite                                                    | app/orange3                                                     |
    #> | arnapipe                                                                | asap                                                            |
    #> | assemble/edena                                                          | assemble/velvet                                                 |
    #> | atlas2                                                                  | autochrom3d                                                     |
    #> | backspin                                                                | ballgown                                                        |
    #> | bamtools                                                                | bamutil                                                         |
    #> | bapi                                                                    | bazam                                                           |
    #> | bcbio-nextgen                                                           | bcftools                                                        |
    #> | beagle                                                                  | bearscc                                                         |
    #> | bedops                                                                  | bedtools2                                                       |
    #> | bget                                                                    | bigstitcher                                                     |
    #> | bin3c                                                                   | biobloom                                                        |
    #> | bioinstaller                                                            | biopython                                                       |
    #> | bitbucket/3depiloop                                                     | bitbucket/aa-stat                                               |
    #> | bitbucket/acdc                                                          | bitbucket/agalma                                                |
    #> | bitbucket/aggrescan3d                                                   | bitbucket/aikyatan                                              |
    #> | bitbucket/alpha                                                         | bitbucket/aodp-v2.0-release                                     |
    #> | bitbucket/arpeggio                                                      | bitbucket/autometa                                              |
    #> | bitbucket/avishkar                                                      | bitbucket/bacmeta                                               |
    #> | bitbucket/badtrip                                                       | bitbucket/bam-matcher                                           |
    #> | bitbucket/banner-chemdner                                               | bitbucket/bart                                                  |
    #> | bitbucket/bbcontacts                                                    | bitbucket/benchmarks                                            |
    #> | bitbucket/bioa                                                          | bitbucket/biobranch                                             |
    #> | bitbucket/bioescorte-suggestion                                         | bitbucket/biofilmanalyzer                                       |
    #> | bitbucket/bioner                                                        | bitbucket/biosubg                                               |
    #> | bitbucket/biotransformerjar                                             | bitbucket/brainformat                                           |
    #> | bitbucket/cabsdock                                                      | bitbucket/cabsflex                                              |
    #> | bitbucket/cafe                                                          | bitbucket/catkern                                               |
    #> | bitbucket/caviarbf                                                      | bitbucket/ccmpred                                               |
    #> | bitbucket/cellsandmachines                                              | bitbucket/cfs-toolbox                                           |
    #> | bitbucket/chemdistillerpython                                           | bitbucket/chimiric                                              |
    #> | bitbucket/chromawalker                                                  | bitbucket/ciml-lib                                              |
    #> | bitbucket/cl-dash                                                       | bitbucket/classyfire-api                                        |
    #> | bitbucket/clonocalc-plot                                                | bitbucket/comorbidity                                           |
    #> | bitbucket/cophesim                                                      | bitbucket/cutruntools                                           |
    #> | bitbucket/cypreact                                                      | bitbucket/cystine-stabilized-proteins                           |
    #> | bitbucket/cytodx-study-code                                             | bitbucket/cytostruct                                            |
    #> | bitbucket/dd-detection                                                  | bitbucket/ddr                                                   |
    #> | bitbucket/decaf                                                         | bitbucket/deepscope                                             |
    #> | bitbucket/defcom                                                        | bitbucket/dendsort                                              |
    #> | bitbucket/densitycut-dev                                                | bitbucket/depthfinder                                           |
    #> | bitbucket/digtyper                                                      | bitbucket/diprog                                                |
    #> | bitbucket/disconica                                                     | bitbucket/dmml                                                  |
    #> | bitbucket/domainoid                                                     | bitbucket/drseq                                                 |
    #> | bitbucket/dyss                                                          | bitbucket/elsa                                                  |
    #> | bitbucket/em-dawid                                                      | bitbucket/enzymer                                               |
    #> | bitbucket/fast                                                          | bitbucket/fastsimbac                                            |
    #> | bitbucket/fbb                                                           | bitbucket/fimtyper                                              |
    #> | bitbucket/fimtyper-db                                                   | bitbucket/fit3d                                                 |
    #> | bitbucket/flexscore                                                     | bitbucket/forqs                                                 |
    #> | bitbucket/g-tris                                                        | bitbucket/gabi                                                  |
    #> | bitbucket/galaxytoolfactory                                             | bitbucket/genediseaserepositioning                              |
    #> | bitbucket/genespider                                                    | bitbucket/genomemining                                          |
    #> | bitbucket/glbase                                                        | bitbucket/glycome-analytics-platform                            |
    #> | bitbucket/glycoprofileassigner                                          | bitbucket/godon                                                 |
    #> | bitbucket/granger-causality                                             | bitbucket/gsoa                                                  |
    #> | bitbucket/hapzipper                                                     | bitbucket/hdgmcm                                                |
    #> | bitbucket/hgvs                                                          | bitbucket/hh-suite                                              |
    #> | bitbucket/hicapp                                                        | bitbucket/hogan                                                 |
    #> | bitbucket/ht-altfrac                                                    | bitbucket/immunediversity                                       |
    #> | bitbucket/imondb                                                        | bitbucket/infiniumpurify                                        |
    #> | bitbucket/infusion                                                      | bitbucket/intervene                                             |
    #> | bitbucket/iso-ktsp                                                      | bitbucket/jaspar                                                |
    #> | bitbucket/jgi-sparc                                                     | bitbucket/jqcml                                                 |
    #> | bitbucket/jra-src                                                       | bitbucket/kernelsvmspark                                        |
    #> | bitbucket/kit                                                           | bitbucket/l.u.st                                                |
    #> | bitbucket/laser-release                                                 | bitbucket/last-split-pe                                         |
    #> | bitbucket/ldetect                                                       | bitbucket/ldetect-data                                          |
    #> | bitbucket/libfmftsaxs                                                   | bitbucket/libsleipnir.bitbucket.org                             |
    #> | bitbucket/lipidhunter                                                   | bitbucket/locality                                              |
    #> | bitbucket/logminer                                                      | bitbucket/m2lite                                                |
    #> | bitbucket/mageck-nest                                                   | bitbucket/mageck-vispr                                          |
    #> | bitbucket/magicmicrorna                                                 | bitbucket/md-analysis-with-matlab                               |
    #> | bitbucket/mechrna                                                       | bitbucket/metabat                                               |
    #> | bitbucket/metamarker                                                    | bitbucket/metamodules                                           |
    #> | bitbucket/metaphinder                                                   | bitbucket/metaprob                                              |
    #> | bitbucket/mgmapper                                                      | bitbucket/micomplete                                            |
    #> | bitbucket/miraw                                                         | bitbucket/mistrvar                                              |
    #> | bitbucket/mlm-gs-supplement                                             | bitbucket/mngsg                                                 |
    #> | bitbucket/moleculedatabaseframework                                     | bitbucket/monovar                                               |
    #> | bitbucket/mrmplusgui                                                    | bitbucket/mugbas                                                |
    #> | bitbucket/mulcch                                                        | bitbucket/mutationmotif                                         |
    #> | bitbucket/mzjava                                                        | bitbucket/neri                                                  |
    #> | bitbucket/nightshift                                                    | bitbucket/norgal                                                |
    #> | bitbucket/nsclc-paper                                                   | bitbucket/omics-pipe                                            |
    #> | bitbucket/oncodriveclustl                                               | bitbucket/oncorep                                               |
    #> | bitbucket/orange-reliability                                            | bitbucket/pancanceramaretto                                     |
    #> | bitbucket/pbrit                                                         | bitbucket/percolator-upgrade                                    |
    #> | bitbucket/phd                                                           | bitbucket/pheno-deep-counter                                    |
    #> | bitbucket/phosfox                                                       | bitbucket/phylogenize                                           |
    #> | bitbucket/phylogenomic-dataset-construction                             | bitbucket/phyloskeleton                                         |
    #> | bitbucket/phyparts                                                      | bitbucket/pia                                                   |
    #> | bitbucket/pipit                                                         | bitbucket/pli                                                   |
    #> | bitbucket/profileseq                                                    | bitbucket/proteogenomics                                        |
    #> | bitbucket/protie                                                        | bitbucket/protwis                                               |
    #> | bitbucket/purge-haplotigs                                               | bitbucket/pymonaco                                              |
    #> | bitbucket/pyseqlab                                                      | bitbucket/qc-analysis                                           |
    #> | bitbucket/qtlsearch                                                     | bitbucket/rawmsa                                                |
    #> | bitbucket/read                                                          | bitbucket/readdi                                                |
    #> | bitbucket/remixt                                                        | bitbucket/repgenhmm                                             |
    #> | bitbucket/resistome-release                                             | bitbucket/rgaugury                                              |
    #> | bitbucket/rna-motifs                                                    | bitbucket/rppa-preprocess                                       |
    #> | bitbucket/rsbiodiv                                                      | bitbucket/rucs                                                  |
    #> | bitbucket/samu661/fsh/overview                                          | bitbucket/sandpuma                                              |
    #> | bitbucket/scavager                                                      | bitbucket/scg                                                   |
    #> | bitbucket/sensa                                                         | bitbucket/seql-nrps                                             |
    #> | bitbucket/sequence-shape                                                | bitbucket/shades                                                |
    #> | bitbucket/shoelaces                                                     | bitbucket/simlord                                               |
    #> | bitbucket/sist-codes                                                    | bitbucket/sl1p                                                  |
    #> | bitbucket/smart-viewpoint-computation-lib                               | bitbucket/sorting-by-network-completion                         |
    #> | bitbucket/sparcc                                                        | bitbucket/sparkseq                                              |
    #> | bitbucket/splot                                                         | bitbucket/spurio                                                |
    #> | bitbucket/srg-extractor                                                 | bitbucket/stereodist                                            |
    #> | bitbucket/subcons-web-server                                            | bitbucket/suppa                                                 |
    #> | bitbucket/svd-phy                                                       | bitbucket/svist4get                                             |
    #> | bitbucket/taxmapper                                                     | bitbucket/tcellxtalk                                            |
    #> | bitbucket/tpes                                                          | bitbucket/triovis                                               |
    #> | bitbucket/trust                                                         | bitbucket/vcf2networks                                          |
    #> | bitbucket/vdjpuzzle2                                                    | bitbucket/venomix                                               |
    #> | bitbucket/virus-vg                                                      | bitbucket/vispa2                                                |
    #> | bitbucket/visualmcmc                                                    | bitbucket/vptissue                                              |
    #> | bitbucket/vsclust                                                       | bitbucket/wanding/duprecover/overview                           |
    #> | bitbucket/warpiv                                                        | bitbucket/whatshap                                              |
    #> | bitbucket/xglycscan                                                     | bitbucket/xpiwit                                                |
    #> | bowtie                                                                  | bowtie2                                                         |
    #> | bpipe                                                                   | breakdancer                                                     |
    #> | breakmer                                                                | breakpointsurveyor                                              |
    #> | brie                                                                    | bwa                                                             |
    #> | bystro                                                                  | bzip2                                                           |
    #> | caveman                                                                 | cdeep3m                                                         |
    #> | cellfishing.jl                                                          | cellprofiler                                                    |
    #> | cellsius                                                                | cesa                                                            |
    #> | chia-pet2                                                               | chicmaxima                                                      |
    #> | chimeraviz                                                              | chromhmm                                                        |
    #> | chromtime                                                               | chromvar                                                        |
    #> | chronqc                                                                 | circbrain                                                       |
    #> | cistopic                                                                | clonealign                                                      |
    #> | clustergrammer                                                          | cn-learn                                                        |
    #> | cnvkit                                                                  | cnvnator                                                        |
    #> | conbase                                                                 | confined                                                        |
    #> | conos                                                                   | cromwell                                                        |
    #> | curl                                                                    | dart                                                            |
    #> | dash                                                                    | db/annovar                                                      |
    #> | db/annovar-1000g                                                        | db/annovar-ensgene                                              |
    #> | db/annovar-knowngene                                                    | db/annovar-noidx                                                |
    #> | db/annovar-refgene                                                      | db/appris                                                       |
    #> | db/atcircdb                                                             | db/awesome                                                      |
    #> | db/biosystems                                                           | db/cancer-hotspots                                              |
    #> | db/cancersplicingqtl                                                    | db/cellmarker                                                   |
    #> | db/cgi                                                                  | db/circbase                                                     |
    #> | db/circnet                                                              | db/circrnadb                                                    |
    #> | db/civic                                                                | db/consensuspathdb                                              |
    #> | db/cscd                                                                 | db/dbsno                                                        |
    #> | db/dcdb                                                                 | db/denovo-db                                                    |
    #> | db/dgidb                                                                | db/diseaseenhancer                                              |
    #> | db/eggnog                                                               | db/ewasdb                                                       |
    #> | db/exorbase                                                             | db/expression-atlas                                             |
    #> | db/exsnp                                                                | db/fantom-cage-peaks                                            |
    #> | db/fantom-co-expression-clusters                                        | db/fantom-enhancers                                             |
    #> | db/fantom-motifs                                                        | db/fantom-ontology                                              |
    #> | db/fantom-tss-classifier                                                | db/funcoup                                                      |
    #> | db/fusiongdb                                                            | db/hgnc                                                         |
    #> | db/hmdb                                                                 | db/hpdi                                                         |
    #> | db/hpo                                                                  | db/inbiomap                                                     |
    #> | db/instruct                                                             | db/interpro                                                     |
    #> | db/islandviewer                                                         | db/journal-doaj                                                 |
    #> | db/lnc2cancer                                                           | db/lncediting                                                   |
    #> | db/lncrnadisease                                                        | db/mircancer                                                    |
    #> | db/mirdb                                                                | db/mirnest                                                      |
    #> | db/mirtarbase                                                           | db/mndr                                                         |
    #> | db/msdd                                                                 | db/omim-open                                                    |
    #> | db/omim-private                                                         | db/oncokb                                                       |
    #> | db/oncomirdb                                                            | db/oncotator                                                    |
    #> | db/pancanqtl                                                            | db/phosphonetworks                                              |
    #> | db/pmkb                                                                 | db/proteinatlas                                                 |
    #> | db/rbp-var                                                              | db/rbpdb                                                        |
    #> | db/rddpred                                                              | db/redoxdb                                                      |
    #> | db/remap                                                                | db/remap2                                                       |
    #> | db/rsnp3                                                                | db/rvarbase                                                     |
    #> | db/sedb                                                                 | db/seecancer                                                    |
    #> | db/seeqtl                                                               | db/sm2mir                                                       |
    #> | db/srnanalyzer                                                          | db/tumorfusions                                                 |
    #> | db/ucsc-cytoband                                                        | db/ucsc-dnase-clustered                                         |
    #> | db/ucsc-ensgene                                                         | db/ucsc-knowngene                                               |
    #> | db/ucsc-refgene                                                         | db/ucsc-tfbs-clustered                                          |
    #> | db/varcards                                                             | dca                                                             |
    #> | deepcell-tf                                                             | deepnovo-dia                                                    |
    #> | deepvariant                                                             | delly                                                           |
    #> | detin                                                                   | divers                                                          |
    #> | doc/awesome-single-cell                                                 | doc/awosome-bioinformatics                                      |
    #> | doc/phatdocs                                                            | doc/pypdb-docs                                                  |
    #> | doc/splatter-paper                                                      | doc/squigglekitdocs                                             |
    #> | doc/trackviewer                                                         | doc/visordoc                                                    |
    #> | dreg                                                                    | dropclust                                                       |
    #> | dstruct                                                                 | easeq                                                           |
    #> | easysvg                                                                 | echarts                                                         |
    #> | effusion                                                                | f-sclvm                                                         |
    #> | facets                                                                  | fastp                                                           |
    #> | fastq-tools                                                             | fastqc                                                          |
    #> | fastx-toolkit                                                           | fatotwobit                                                      |
    #> | feast                                                                   | fmriprep                                                        |
    #> | forge                                                                   | freebayes                                                       |
    #> | freec                                                                   | fusioncatcher                                                   |
    #> | g2s                                                                     | gatk-bundle                                                     |
    #> | gatk4                                                                   | gdc-client                                                      |
    #> | gemini                                                                  | genomedisco                                                     |
    #> | genomeuplot                                                             | genvisr                                                         |
    #> | geogrid                                                                 | ggdag                                                           |
    #> | ggseqlogo                                                               | ggthemr                                                         |
    #> | giggle                                                                  | github/13check-rna                                              |
    #> | github/16gt                                                             | github/2kplus2                                                  |
    #> | github/2matrix                                                          | github/2stepqa                                                  |
    #> | github/3d-printed-radiographic-test-tools                               | github/3dchromatin-replicateqc                                  |
    #> | github/3dec                                                             | github/3dface                                                   |
    #> | github/3dmax                                                            | github/3dpatch                                                  |
    #> | github/435271                                                           | github/4d-nucleome-analysis-toolbox                             |
    #> | github/4dassign                                                         | github/4pipe4                                                   |
    #> | github/aalto-ics-kepaco                                                 | github/aascatterplot                                            |
    #> | github/abb                                                              | github/abbyyan3/bhglm                                           |
    #> | github/abc                                                              | github/abdominal-mr-phantom                                     |
    #> | github/abis                                                             | github/able                                                     |
    #> | github/abmda                                                            | github/abra                                                     |
    #> | github/abra2                                                            | github/abscan                                                   |
    #> | github/absnf                                                            | github/abus-code                                                |
    #> | github/ac-diamond                                                       | github/ac-pca                                                   |
    #> | github/accumulate                                                       | github/aces                                                     |
    #> | github/acgh-viewer                                                      | github/aci                                                      |
    #> | github/acnc-dame                                                        | github/acnviewer                                                |
    #> | github/acp-dl                                                           | github/actinn                                                   |
    #> | github/ad-zcc                                                           | github/adapt-mix                                                |
    #> | github/adaptive-geometric-search-for-protein-design                     | github/adaptivehm                                               |
    #> | github/adcp                                                             | github/addit                                                    |
    #> | github/additive-fnnrw                                                   | github/addo                                                     |
    #> | github/aditya-88/asap                                                   | github/adjutant                                                 |
    #> | github/admixem                                                          | github/admixture-graph                                          |
    #> | github/ads-hcspark                                                      | github/advanced-multiloops                                      |
    #> | github/adversarial-relation-classification                              | github/aether                                                   |
    #> | github/afcluster                                                        | github/affylumcna                                               |
    #> | github/affypipe                                                         | github/afresh                                                   |
    #> | github/agennt                                                           | github/aggregategenefunctionprediction                          |
    #> | github/agin                                                             | github/agotron-detector                                         |
    #> | github/agplus                                                           | github/agrp                                                     |
    #> | github/ags-and-acn-tools                                                | github/airnet-pytorch                                           |
    #> | github/aivar                                                            | github/ajia                                                     |
    #> | github/aksmooth                                                         | github/akt-selective                                            |
    #> | github/al3c                                                             | github/alarm                                                    |
    #> | github/aldenleung/omtools                                               | github/ale                                                      |
    #> | github/algorithm-performance-analysis                                   | github/algorithmcomparison                                      |
    #> | github/alidetection                                                     | github/alifreefold                                              |
    #> | github/align-linguistic-alignment                                       | github/align3d                                                  |
    #> | github/alignerboost                                                     | github/aligngraph                                               |
    #> | github/alleleanalyzer                                                   | github/allelic-inclusion                                        |
    #> | github/allo                                                             | github/allonkleinlab/spring                                     |
    #> | github/almostsignificant                                                | github/alo-algorithm-for-kidney-exchanges                       |
    #> | github/alpha                                                            | github/alpha-centauri                                           |
    #> | github/als-deeplearning                                                 | github/althap                                                   |
    #> | github/althapalignr                                                     | github/altre                                                    |
    #> | github/alview                                                           | github/amap                                                     |
    #> | github/amas                                                             | github/amda                                                     |
    #> | github/amplimap                                                         | github/amplisolve                                               |
    #> | github/ampumi                                                           | github/amylogramanalysis                                        |
    #> | github/ananas                                                           | github/ananke                                                   |
    #> | github/anaquin                                                          | github/anatomy-modality-decomposition                           |
    #> | github/ancestral-blocks-reconstruction                                  | github/ancestry-viz                                             |
    #> | github/ancis-pytorch                                                    | github/and                                                      |
    #> | github/andi                                                             | github/andy-s-algorithm                                         |
    #> | github/aneuvis                                                          | github/angsd-wrapper                                            |
    #> | github/angular-ripleys-k                                                | github/ann-glycolysis-flux-prediction                           |
    #> | github/ann-solo                                                         | github/anndata                                                  |
    #> | github/annocript                                                        | github/annopeak                                                 |
    #> | github/annotatr                                                         | github/anonimme                                                 |
    #> | github/antibody-2019                                                    | github/antibodyinterfaceprediction                              |
    #> | github/anticancer-peptides-review                                       | github/antigenpredictor                                         |
    #> | github/antivpp                                                          | github/aozan                                                    |
    #> | github/ap11-samifier                                                    | github/apero                                                    |
    #> | github/aphid                                                            | github/apinet                                                   |
    #> | github/apostl                                                           | github/apples                                                   |
    #> | github/appscangeo                                                       | github/aptablocks                                               |
    #> | github/aqua                                                             | github/ar-pred-source                                           |
    #> | github/arachne                                                          | github/arcas                                                    |
    #> | github/arcashla                                                         | github/architect                                                |
    #> | github/arcs                                                             | github/ardiss                                                   |
    #> | github/argdit                                                           | github/argon                                                    |
    #> | github/args-oap-v2.0                                                    | github/argyle                                                   |
    #> | github/arioc                                                            | github/ark                                                      |
    #> | github/arl-eegmodels                                                    | github/armsd                                                    |
    #> | github/arnapipe                                                         | github/arraylasso                                               |
    #> | github/arraymaker                                                       | github/artgan                                                   |
    #> | github/artifusion                                                       | github/artmap                                                   |
    #> | github/arts                                                             | github/aryana-aligner                                           |
    #> | github/asafe                                                            | github/asap                                                     |
    #> | github/asar                                                             | github/asciigenome                                              |
    #> | github/asd-genes-prediction                                             | github/asdpex                                                   |
    #> | github/aselux                                                           | github/asgart                                                   |
    #> | github/ashr                                                             | github/asja                                                     |
    #> | github/aspc                                                             | github/assemblosis                                              |
    #> | github/assemblytics                                                     | github/assexon                                                  |
    #> | github/assign                                                           | github/assocplots                                               |
    #> | github/astrap                                                           | github/atac-pipe                                                |
    #> | github/atc                                                              | github/aten                                                     |
    #> | github/atlas-rat                                                        | github/atlases                                                  |
    #> | github/atma                                                             | github/atminter                                                 |
    #> | github/atropos                                                          | github/atsnp                                                    |
    #> | github/att-chemdner                                                     | github/att-chemprot                                             |
    #> | github/augmentor                                                        | github/augmentor.jl                                             |
    #> | github/author-detection                                                 | github/autoimmune-research                                      |
    #> | github/av-segmentation                                                  | github/avesim                                                   |
    #> | github/awfisher                                                         | github/awol-mrf                                                 |
    #> | github/axe                                                              | github/axondeepseg                                              |
    #> | github/axonpacking                                                      | github/axonseg                                                  |
    #> | github/azahar                                                           | github/b-lore                                                   |
    #> | github/b-mis-normalization                                              | github/b-nem                                                    |
    #> | github/backclip                                                         | github/bacpacs                                                  |
    #> | github/bactdating                                                       | github/bacterial-colonization-model                             |
    #> | github/bacteriamslf                                                     | github/bactsnp                                                  |
    #> | github/badlands-model                                                   | github/badock                                                   |
    #> | github/baerhunter                                                       | github/bagse                                                    |
    #> | github/baitfisher-package                                               | github/baldr                                                    |
    #> | github/bam-abs                                                          | github/bam2ssj                                                  |
    #> | github/bamchop                                                          | github/bamclipper                                               |
    #> | github/bamfa                                                            | github/bamgineer                                                |
    #> | github/bamhash                                                          | github/bamixchecker                                             |
    #> | github/bamse                                                            | github/bamtools                                                 |
    #> | github/barnaba                                                          | github/bart-bma                                                 |
    #> | github/bartender-1.1                                                    | github/basics                                                   |
    #> | github/batch-ge                                                         | github/batcheffectremoval                                       |
    #> | github/batchqc                                                          | github/batmeth2                                                 |
    #> | github/bayescat                                                         | github/bayesembler                                              |
    #> | github/bayesiandatafusion.jl                                            | github/bayesianpgmm                                             |
    #> | github/bayexer                                                          | github/baynorm                                                  |
    #> | github/baynorm-papercode                                                | github/bbarker/falcon                                           |
    #> | github/bbgp                                                             | github/bbknn                                                    |
    #> | github/bc5cidtask                                                       | github/bc6pm-hrnn                                               |
    #> | github/bcalm                                                            | github/bcftools                                                 |
    #> | github/bcgtree                                                          | github/bcigepred                                                |
    #> | github/bcool                                                            | github/bcrystal                                                 |
    #> | github/bdbg                                                             | github/bdc                                                      |
    #> | github/bdchemo                                                          | github/bdmma                                                    |
    #> | github/bdmma-macos                                                      | github/bdss                                                     |
    #> | github/beacon-network-inference                                         | github/beam                                                     |
    #> | github/beam-propagation-method                                          | github/bel-enrichment                                           |
    #> | github/bel2abm                                                          | github/bemkl-rbps                                               |
    #> | github/benchmark-models                                                 | github/benchmarking-tsdiscretizations                           |
    #> | github/benchmarkncvtools                                                | github/bereta                                                   |
    #> | github/bermuda                                                          | github/besst                                                    |
    #> | github/betaboost                                                        | github/betaserpentine                                           |
    #> | github/betaturn18                                                       | github/beyondbinaryparcellationdata                             |
    #> | github/bfc                                                              | github/bfmem                                                    |
    #> | github/bgsa                                                             | github/bgsc                                                     |
    #> | github/bgt                                                              | github/bhklab                                                   |
    #> | github/bib                                                              | github/bicolor                                                  |
    #> | github/biddsat                                                          | github/bide-2d                                                  |
    #> | github/bidifuse                                                         | github/bigbwa                                                   |
    #> | github/bigdatagenomics/mango                                            | github/bigld                                                    |
    #> | github/bigred                                                           | github/bilouvain                                                |
    #> | github/bimberlab/discvrseq                                              | github/bin-passing-analyzer                                     |
    #> | github/bin3c                                                            | github/bindash                                                  |
    #> | github/bindpredict                                                      | github/binm                                                     |
    #> | github/binning                                                          | github/binning-refiner                                          |
    #> | github/bio-gradient-descent                                             | github/bio-quinn2013                                            |
    #> | github/bio-scores                                                       | github/bio-tradis                                               |
    #> | github/bio3d                                                            | github/bioattribution                                           |
    #> | github/biobert                                                          | github/biobert-pretrained                                       |
    #> | github/bioblend                                                         | github/bioc                                                     |
    #> | github/biocaddie                                                        | github/biocaddie2016mayodata                                    |
    #> | github/biocemid                                                         | github/bioclipse                                                |
    #> | github/biocompute-objects                                               | github/biocontainers                                            |
    #> | github/biocppi-extraction                                               | github/biocreativevi-bioid-assignment                           |
    #> | github/biodataome                                                       | github/biodiscml                                                |
    #> | github/bioflosoftware                                                   | github/biograph                                                 |
    #> | github/bioinfo                                                          | github/bioinformatics-sourcecode                                |
    #> | github/bioinstaller                                                     | github/biojava-tutorial                                         |
    #> | github/biojs                                                            | github/biojs-io-biom                                            |
    #> | github/biokeen                                                          | github/biolab/red                                               |
    #> | github/biolitmap                                                        | github/biomaj2galaxy                                            |
    #> | github/biomake                                                          | github/biomartr                                                 |
    #> | github/biomedical-corpora                                               | github/biomedical-qa                                            |
    #> | github/biomethyl                                                        | github/biomsef                                                  |
    #> | github/bioner-cross-sharing                                             | github/bionet-mining                                            |
    #> | github/bionetgen                                                        | github/bionev                                                   |
    #> | github/bionitio                                                         | github/bionmf-gpu                                               |
    #> | github/bioont-search-benchmark                                          | github/biopartsbuilder                                          |
    #> | github/biopartsdb                                                       | github/biopax.viz                                               |
    #> | github/bioposdep                                                        | github/biopyramid                                               |
    #> | github/bioqueue                                                         | github/bioruby-svgenes                                          |
    #> | github/bioruby-ucsc-api                                                 | github/bioshake                                                 |
    #> | github/biostructmap                                                     | github/biostructurem                                            |
    #> | github/biosual                                                          | github/biotite                                                  |
    #> | github/biotoolscompose                                                  | github/bipspi                                                   |
    #> | github/biren                                                            | github/bispark                                                  |
    #> | github/bisque                                                           | github/bits                                                     |
    #> | github/bitseq                                                           | github/bitseqvb-benchmarking                                    |
    #> | github/bives-statsgenerator                                             | github/biwalklda                                                |
    #> | github/bixgboost                                                        | github/bjass                                                    |
    #> | github/blant                                                            | github/blasst                                                   |
    #> | github/blastgraph                                                       | github/blastgui                                                 |
    #> | github/blastjs                                                          | github/blca                                                     |
    #> | github/blend4php                                                        | github/blisar                                                   |
    #> | github/blmrm                                                            | github/blobology                                                |
    #> | github/blobsplorer                                                      | github/bluesnp                                                  |
    #> | github/bmdexpress-2                                                     | github/bmf-qsar                                                 |
    #> | github/bminerimportancebacillus                                         | github/bmsim                                                    |
    #> | github/bnbr                                                             | github/bner                                                     |
    #> | github/bnnr                                                             | github/bnrr                                                     |
    #> | github/boa                                                              | github/bofdat                                                   |
    #> | github/boiler                                                           | github/boltzmannmachines.jl                                     |
    #> | github/bonita                                                           | github/boolean-t2dm                                             |
    #> | github/boolesim                                                         | github/boost-hic                                                |
    #> | github/boostgapfill                                                     | github/bootejtk                                                 |
    #> | github/boss                                                             | github/bowtie-scaling                                           |
    #> | github/bp-quant                                                         | github/bpbi                                                     |
    #> | github/bpipe                                                            | github/bpp                                                      |
    #> | github/bprmeth                                                          | github/bpsc                                                     |
    #> | github/brainimager                                                      | github/branchinggps                                             |
    #> | github/brapes                                                           | github/brca-analyzer                                            |
    #> | github/breakid                                                          | github/breakpointer                                             |
    #> | github/breakpointsurveyor                                               | github/breastcancerclassifier                                   |
    #> | github/bret-analyzer                                                    | github/brides                                                   |
    #> | github/bridge                                                           | github/bridges                                                  |
    #> | github/brm                                                              | github/broad-peaks                                              |
    #> | github/broccoli                                                         | github/bronchomix                                               |
    #> | github/browniealigner                                                   | github/browniecorrector                                         |
    #> | github/browsevcf                                                        | github/bruno                                                    |
    #> | github/brwhnha                                                          | github/bs-snper                                                 |
    #> | github/bsevaluationremotesceneir                                        | github/bstools                                                  |
    #> | github/bsvf                                                             | github/bth                                                      |
    #> | github/bts-dsn                                                          | github/btw                                                      |
    #> | github/btyper                                                           | github/bulkvis                                                  |
    #> | github/bvnlabscattome                                                   | github/bvp-pred-unb                                             |
    #> | github/bwas                                                             | github/bwmr                                                     |
    #> | github/c-deva                                                           | github/c-hmm                                                    |
    #> | github/c-intersecture                                                   | github/c3d                                                      |
    #> | github/caars                                                            | github/cabergh/ebdims                                           |
    #> | github/cadrres                                                          | github/cafe                                                     |
    #> | github/cafe-plugin                                                      | github/cafemocha                                                |
    #> | github/caffe-3d-faster-rcnn                                             | github/cafu                                                     |
    #> | github/calib                                                            | github/calour                                                   |
    #> | github/calq                                                             | github/cam                                                      |
    #> | github/camisim                                                          | github/camstyle                                                 |
    #> | github/canalizingpower                                                  | github/canary                                                   |
    #> | github/cancer-subtyping                                                 | github/cancerdiscover                                           |
    #> | github/candi                                                            | github/canet                                                    |
    #> | github/cansnper                                                         | github/canvas                                                   |
    #> | github/canvasdb                                                         | github/capc-map                                                 |
    #> | github/capsim                                                           | github/capsnet-ptm                                              |
    #> | github/capssa                                                           | github/cardiacpbpk                                              |
    #> | github/care-rcortex                                                     | github/carpools                                                 |
    #> | github/cartoon-network                                                  | github/casas                                                    |
    #> | github/case-bbb-prediction-data                                         | github/casian                                                   |
    #> | github/caslocusanno                                                     | github/casmap                                                   |
    #> | github/casper                                                           | github/caspo                                                    |
    #> | github/caspo-ts                                                         | github/castin                                                   |
    #> | github/catana                                                           | github/catfish                                                  |
    #> | github/cath-tools                                                       | github/causaltrail                                              |
    #> | github/cbig                                                             | github/cc-mds                                                   |
    #> | github/ccbgpipe                                                         | github/ccc                                                      |
    #> | github/cclasso                                                          | github/cd-trace                                                 |
    #> | github/cddapp                                                           | github/cddforfmri                                               |
    #> | github/cdpc                                                             | github/cell-maps                                                |
    #> | github/cellminercompanion                                               | github/cellmissy                                                |
    #> | github/cellnomenclaturestudy                                            | github/cellprofiler-analyst                                     |
    #> | github/cellsim                                                          | github/celltrans                                                |
    #> | github/cepics                                                           | github/cerebro                                                  |
    #> | github/cerebroapp                                                       | github/cerena                                                   |
    #> | github/cesar                                                            | github/cesar2.0                                                 |
    #> | github/cfdnapattern                                                     | github/cfnet                                                    |
    #> | github/cgan                                                             | github/cgdm                                                     |
    #> | github/cgheliparm                                                       | github/cglasso                                                  |
    #> | github/cgmisc                                                           | github/cgrtools                                                 |
    #> | github/cgvr                                                             | github/chainrank                                                |
    #> | github/chance                                                           | github/changlab                                                 |
    #> | github/chaperism                                                        | github/charger                                                  |
    #> | github/checkmyblob                                                      | github/checkmyindex                                             |
    #> | github/chem-preview                                                     | github/chemogenomicalg4dtipred                                  |
    #> | github/chemps2                                                          | github/chemts                                                   |
    #> | github/chewbbaca                                                        | github/chexmix                                                  |
    #> | github/chfs                                                             | github/chia-pet2                                                |
    #> | github/chiapop                                                          | github/chicdiff                                                 |
    #> | github/chiimp                                                           | github/chilay                                                   |
    #> | github/chilin                                                           | github/chimeraviz                                               |
    #> | github/chimericognizer                                                  | github/chimerscope                                              |
    #> | github/chippcr                                                          | github/chipsad                                                  |
    #> | github/chipseqspikeinfree                                               | github/chipulate                                                |
    #> | github/chipwig-v2                                                       | github/chmannot                                                 |
    #> | github/chopbaicontact:birte.kehr@decode.is                              | github/chopstitch                                               |
    #> | github/chordomics                                                       | github/christopherblum                                          |
    #> | github/chroma-clade                                                     | github/chromatra                                                |
    #> | github/chromdet                                                         | github/chromdragonn                                             |
    #> | github/chromozoom                                                       | github/chronqc                                                  |
    #> | github/cidr                                                             | github/cima                                                     |
    #> | github/circacompare                                                     | github/circadb                                                  |
    #> | github/circcode                                                         | github/circdeep                                                 |
    #> | github/circmeta                                                         | github/circompara                                               |
    #> | github/circsplice                                                       | github/circtools                                                |
    #> | github/circuitservice                                                   | github/cirgo                                                    |
    #> | github/cissticp                                                         | github/cistopic                                                 |
    #> | github/cjbitseq                                                         | github/clairvoyante                                             |
    #> | github/clamms                                                           | github/classifier-selection-code                                |
    #> | github/cld                                                              | github/clickx                                                   |
    #> | github/clifinder                                                        | github/cline                                                    |
    #> | github/clinical-citation-sentiment                                      | github/clinical-sentences                                       |
    #> | github/cliquems                                                         | github/clj-biosequence                                          |
    #> | github/clockstarg                                                       | github/clonefinderapi                                           |
    #> | github/clonevol                                                         | github/cloops                                                   |
    #> | github/close                                                            | github/cloudauthz                                               |
    #> | github/cloudbrush                                                       | github/cloudforest                                              |
    #> | github/cloudneo                                                         | github/cloudphylo                                               |
    #> | github/cluesv1                                                          | github/clumsid                                                  |
    #> | github/clusdca                                                          | github/clust                                                    |
    #> | github/cluster-wgs-data                                                 | github/clusterdv                                                |
    #> | github/clustermap                                                       | github/clustermq                                                |
    #> | github/clustermq-performance                                            | github/clustertad                                               |
    #> | github/cmap                                                             | github/cmb-labeler                                              |
    #> | github/cmfsm                                                            | github/cmonkey2                                                 |
    #> | github/cmsa                                                             | github/cnanalysis450k                                           |
    #> | github/cnape                                                            | github/cnara                                                    |
    #> | github/cnefinder                                                        | github/cner                                                     |
    #> | github/cnet                                                             | github/cnn                                                      |
    #> | github/cnn-brain-strains                                                | github/cnn-smoothie                                             |
    #> | github/cnr                                                              | github/cnr-analyses                                             |
    #> | github/cnspector                                                        | github/cnt-ilp                                                  |
    #> | github/cnv-prioritization                                               | github/cnvalidator                                              |
    #> | github/cnvkit                                                           | github/co-fuse                                                  |
    #> | github/coac                                                             | github/coala                                                    |
    #> | github/cobamp                                                           | github/cobasi                                                   |
    #> | github/cobia                                                            | github/cobralab-atlases                                         |
    #> | github/cobratoolbox                                                     | github/cocacola                                                 |
    #> | github/cocos                                                            | github/cocoscore                                                |
    #> | github/coda                                                             | github/codeml-modl                                              |
    #> | github/codonhmm                                                         | github/coelho2015-netsdetermination                             |
    #> | github/coge                                                             | github/cogena                                                   |
    #> | github/cognate                                                          | github/cogstack                                                 |
    #> | github/cola                                                             | github/collector                                                |
    #> | github/colliderapp                                                      | github/colocalizr                                               |
    #> | github/colocr                                                           | github/colormap                                                 |
    #> | github/combat-tb                                                        | github/combination-index                                        |
    #> | github/combine-cnn-enhancer-and-promoters                               | github/combined-pvalues                                         |
    #> | github/combiningdependentpvalues                                        | github/come                                                     |
    #> | github/comer                                                            | github/cometa                                                   |
    #> | github/cometsc                                                          | github/comida                                                   |
    #> | github/comm                                                             | github/comorbidity4j                                            |
    #> | github/comoto-pasteur-fr                                                | github/comparisonfastafiles                                     |
    #> | github/compath                                                          | github/complete-striped-smith-waterman-library                  |
    #> | github/complexviewer                                                    | github/compms2miner                                             |
    #> | github/compomics-utilities                                              | github/compound-eye-simulator                                   |
    #> | github/compressgv                                                       | github/computational                                            |
    #> | github/computel                                                         | github/comsa                                                    |
    #> | github/concise                                                          | github/condition-dependent-correlation-subgroups-ccs            |
    #> | github/condo                                                            | github/conffuse                                                 |
    #> | github/confidence-information-ontology                                  | github/confindr                                                 |
    #> | github/confined                                                         | github/confit                                                   |
    #> | github/confold2                                                         | github/conformationalchange                                     |
    #> | github/conics                                                           | github/connjur-sandbox                                          |
    #> | github/connor-lab/vapor                                                 | github/conpair                                                  |
    #> | github/consensus                                                        | github/consensx.itk.ppke.hu                                     |
    #> | github/consexpression                                                   | github/conspecifix                                              |
    #> | github/contaminationx                                                   | github/contextual-regression-for-circrna                        |
    #> | github/convlstmforgr                                                    | github/cooler                                                   |
    #> | github/copal                                                            | github/copmem                                                   |
    #> | github/coptea                                                           | github/cordova                                                  |
    #> | github/coretracker                                                      | github/cornas                                                   |
    #> | github/corpora                                                          | github/correlation-between-rna-seq-and-rrbs                     |
    #> | github/cosmomvpa                                                        | github/cosmos                                                   |
    #> | github/cosr                                                             | github/counting-consistent-sub-dag                              |
    #> | github/countsimqc                                                       | github/couplenet                                                |
    #> | github/coverview                                                        | github/cox-nnet                                                 |
    #> | github/coxboost                                                         | github/cpdock                                                   |
    #> | github/cpi-em                                                           | github/cpinsim                                                  |
    #> | github/cpm-cytoscape                                                    | github/cpr                                                      |
    #> | github/craftgbd                                                         | github/cram-js                                                  |
    #> | github/crdata                                                           | github/creamino                                                 |
    #> | github/cregulome                                                        | github/crenet                                                   |
    #> | github/crisflash                                                        | github/crispr                                                   |
    #> | github/crispr-dav                                                       | github/crisprdisco                                              |
    #> | github/crispritz                                                        | github/crisprmatch                                              |
    #> | github/crisprpred                                                       | github/crm                                                      |
    #> | github/croco                                                            | github/crossbar-net                                             |
    #> | github/crosslink-viewer                                                 | github/crossmapper                                              |
    #> | github/crossplan                                                        | github/crowd-cid-relex                                          |
    #> | github/crowdnodnet                                                      | github/crtl                                                     |
    #> | github/crumble                                                          | github/cryfa                                                    |
    #> | github/cryostage                                                        | github/csar                                                     |
    #> | github/csbfinder                                                        | github/cscoretool                                               |
    #> | github/csdm                                                             | github/csi-utr                                                  |
    #> | github/csmd                                                             | github/cssscl                                                   |
    #> | github/csvd                                                             | github/ctap                                                     |
    #> | github/ctc-screening-cnn                                                | github/ctcf-mp                                                  |
    #> | github/ctdquerier                                                       | github/cuclark                                                  |
    #> | github/cudagsea                                                         | github/cudampf                                                  |
    #> | github/cuddi                                                            | github/cuedemotion                                              |
    #> | github/cufusion                                                         | github/cumulative-fourier-power-spectrum                        |
    #> | github/cupsoda                                                          | github/curatio                                                  |
    #> | github/curatr                                                           | github/cusnn                                                    |
    #> | github/cwdtw                                                            | github/cwords                                                   |
    #> | github/cwsdtwnano                                                       | github/cymer                                                    |
    #> | github/cymira                                                           | github/cysbar                                                   |
    #> | github/cysmotifsearcher                                                 | github/cytobackbone                                             |
    #> | github/cytocompare                                                      | github/cytofmerge                                               |
    #> | github/cytokit                                                          | github/cytometry-clustering-comparison                          |
    #> | github/cyversewarwick                                                   | github/d-gex                                                    |
    #> | github/d2r-codes                                                        | github/d2sbin                                                   |
    #> | github/d3m                                                              | github/d3ner                                                    |
    #> | github/dace                                                             | github/dacum                                                    |
    #> | github/dada2                                                            | github/dafga                                                    |
    #> | github/dag-viewer-biojs                                                 | github/dagitty                                                  |
    #> | github/dairydb                                                          | github/daisy                                                    |
    #> | github/dalliance                                                        | github/dangertrack                                              |
    #> | github/danq                                                             | github/dante                                                    |
    #> | github/dart                                                             | github/darts                                                    |
    #> | github/dasc                                                             | github/data-example                                             |
    #> | github/data-in-brief-influence                                          | github/data-of-genome-rearrangement                             |
    #> | github/data-prism                                                       | github/data2dynamics                                            |
    #> | github/databaseui                                                       | github/datamonkey-js                                            |
    #> | github/datarlatools                                                     | github/dataset-retrieval-pipeline                               |
    #> | github/datasets                                                         | github/datasink                                                 |
    #> | github/dave-the-scientist                                               | github/dbc                                                      |
    #> | github/dbgap2x                                                          | github/dbh                                                      |
    #> | github/dbtora                                                           | github/dcars                                                    |
    #> | github/dcaxl                                                            | github/dcj-c                                                    |
    #> | github/dcm                                                              | github/dcmqi                                                    |
    #> | github/dcnlp                                                            | github/dde-bd                                                   |
    #> | github/ddot                                                             | github/ddpcrclust                                               |
    #> | github/ddpcrvis                                                         | github/ddsae                                                    |
    #> | github/ddseeker                                                         | github/dearwxpub                                                |
    #> | github/debga                                                            | github/debgr                                                    |
    #> | github/deblender                                                        | github/deblur                                                   |
    #> | github/debruijn-motif                                                   | github/debwt                                                    |
    #> | github/debwtcontact                                                     | github/decent                                                   |
    #> | github/decometdia                                                       | github/decon1d                                                  |
    #> | github/decontam                                                         | github/deconvolution-of-essential-gene-signitures               |
    #> | github/deconvseq                                                        | github/dectdec                                                  |
    #> | github/deep-collaborative-filtering                                     | github/deep-learning-examples                                   |
    #> | github/deep-plant                                                       | github/deep-resp-forest                                         |
    #> | github/deepacet                                                         | github/deepaffinity                                             |
    #> | github/deepathology                                                     | github/deepbinner                                               |
    #> | github/deepboost                                                        | github/deepbound                                                |
    #> | github/deepcas9                                                         | github/deepchem                                                 |
    #> | github/deepcircode                                                      | github/deepcnf-auc                                              |
    #> | github/deepcoil                                                         | github/deepcon                                                  |
    #> | github/deepconv-dti                                                     | github/deepcov                                                  |
    #> | github/deepcropping                                                     | github/deepcrystal                                              |
    #> | github/deepcseqsite                                                     | github/deepcytof                                                |
    #> | github/deepdiffraction                                                  | github/deepdive-genegene-app                                    |
    #> | github/deepdna                                                          | github/deepdom                                                  |
    #> | github/deepdr                                                           | github/deepdrug3d                                               |
    #> | github/deepdta                                                          | github/deepem3d                                                 |
    #> | github/deepexpression                                                   | github/deepfe-ppi                                               |
    #> | github/deepfold                                                         | github/deepfr                                                   |
    #> | github/deepgo                                                           | github/deephint                                                 |
    #> | github/deephistone                                                      | github/deephlapan                                               |
    #> | github/deepimpute                                                       | github/deepisofun                                               |
    #> | github/deepligand                                                       | github/deeplightfieldssr                                        |
    #> | github/deepm6aseq                                                       | github/deepmetapsicov                                           |
    #> | github/deepmirtar-sda                                                   | github/deepmito                                                 |
    #> | github/deepmspeptide                                                    | github/deepnet-rbp                                              |
    #> | github/deepnf                                                           | github/deeppasta                                                |
    #> | github/deeppath                                                         | github/deepphos                                                 |
    #> | github/deeppicker-python                                                | github/deepppisp                                                |
    #> | github/deepram                                                          | github/deepred                                                  |
    #> | github/deeprtplus                                                       | github/deepseenet                                               |
    #> | github/deepseqpan                                                       | github/deepsequence                                             |
    #> | github/deepsignal                                                       | github/deepsimulator                                            |
    #> | github/deepslide                                                        | github/deepsnr                                                  |
    #> | github/deepsnvminer                                                     | github/deept3                                                   |
    #> | github/deepubi                                                          | github/deepubiquitylation                                       |
    #> | github/deerect-polya                                                    | github/defor                                                    |
    #> | github/defume                                                           | github/degps                                                    |
    #> | github/degsm                                                            | github/deicode                                                  |
    #> | github/deisom                                                           | github/delta                                                    |
    #> | github/delta-statistic                                                  | github/deltabs                                                  |
    #> | github/demix-q                                                          | github/demixtallmaterials                                       |
    #> | github/dendropy                                                         | github/dendrosplit                                              |
    #> | github/denoptim                                                         | github/densitypath                                              |
    #> | github/deopen                                                           | github/depecher                                                 |
    #> | github/deplanckelab/asap                                                | github/deplogo                                                  |
    #> | github/deploid                                                          | github/deploid-r                                                |
    #> | github/depop                                                            | github/derep-np                                                 |
    #> | github/derfinder                                                        | github/description-extractor                                    |
    #> | github/descs-standalone                                                 | github/designmetabolicrnalabeling                               |
    #> | github/designmpra                                                       | github/desir                                                    |
    #> | github/destin                                                           | github/detect-v2                                                |
    #> | github/detectron                                                        | github/detime                                                   |
    #> | github/dets                                                             | github/deuterater                                               |
    #> | github/deuteros                                                         | github/deviate                                                  |
    #> | github/dexsi                                                            | github/dfim                                                     |
    #> | github/dfm                                                              | github/dgcox                                                    |
    #> | github/dgendoo                                                          | github/dgenies                                                  |
    #> | github/dgmdl                                                            | github/dhaka                                                    |
    #> | github/diagram                                                          | github/dialignr                                                 |
    #> | github/dib-lab/khmer                                                    | github/dieterich-lab                                            |
    #> | github/differential-mutation-analysis                                   | github/differential-privacy-genomic-inference-attack            |
    #> | github/diffexpy                                                         | github/diffgraph                                                |
    #> | github/diffnetfdr                                                       | github/diffpop                                                  |
    #> | github/diffsplicing                                                     | github/diffuse                                                  |
    #> | github/diffusivedynamics                                                | github/digitalcellsorter                                        |
    #> | github/dinamo                                                           | github/dinar                                                    |
    #> | github/dingo                                                            | github/dinosaur                                                 |
    #> | github/dipartite                                                        | github/disambiguate                                             |
    #> | github/discoal                                                          | github/discomark                                                |
    #> | github/discordant                                                       | github/discorhythm                                              |
    #> | github/discoversl                                                       | github/disdrugpred                                              |
    #> | github/disease-similarity-fusion                                        | github/diseasediscovery                                         |
    #> | github/diseaseextract                                                   | github/dislncrf                                                 |
    #> | github/dispot                                                           | github/dispredict-v1.0                                          |
    #> | github/disseqt.jl                                                       | github/distanced                                                |
    #> | github/distributedlearningpredictor                                     | github/distruct                                                 |
    #> | github/disvis                                                           | github/dive                                                     |
    #> | github/diversity                                                        | github/divpop                                                   |
    #> | github/divvier                                                          | github/django-chicp                                             |
    #> | github/dlbi                                                             | github/dligand2                                                 |
    #> | github/dlprb                                                            | github/dmbc                                                     |
    #> | github/dmcm                                                             | github/dmdeepm6a1.0                                             |
    #> | github/dmdtoolkit                                                       | github/dmfinder                                                 |
    #> | github/dmi                                                              | github/dmrfinder                                                |
    #> | github/dmsc                                                             | github/dna-nn                                                   |
    #> | github/dna-rchitect                                                     | github/dnabinding                                               |
    #> | github/dnacc                                                            | github/dnadataaugmentation                                      |
    #> | github/dnadu                                                            | github/dnam-based-age-predictor                                 |
    #> | github/dnapi                                                            | github/dnascan                                                  |
    #> | github/dnascent                                                         | github/dnbfa                                                    |
    #> | github/dncon2                                                           | github/dnea                                                     |
    #> | github/dnn-hmm                                                          | github/do-ms                                                    |
    #> | github/doc2hpo                                                          | github/docker-galaxy-hicexplorer                                |
    #> | github/docker-rrbssim                                                   | github/dockq                                                    |
    #> | github/docks                                                            | github/doepipeline                                              |
    #> | github/dogfinder                                                        | github/dognet                                                   |
    #> | github/domex                                                            | github/doraemon                                                 |
    #> | github/doritool                                                         | github/dot2dot                                                  |
    #> | github/dove                                                             | github/downloads                                                |
    #> | github/dp-gp-cluster                                                    | github/dp-representation-transfer                               |
    #> | github/dpetstep                                                         | github/dphansti/coral                                           |
    #> | github/dphansti/mango                                                   | github/dpm-lgcp                                                 |
    #> | github/dppi                                                             | github/dpre                                                     |
    #> | github/dptdt                                                            | github/dpuc2                                                    |
    #> | github/dqbioinformatics                                                 | github/dr-thermo                                                |
    #> | github/dr2di                                                            | github/draco-stem                                               |
    #> | github/drdiff                                                           | github/dream-gene-essentiality-challenge                        |
    #> | github/dreamtools                                                       | github/dreg                                                     |
    #> | github/drimpute                                                         | github/driverml                                                 |
    #> | github/driversub                                                        | github/drl4cellmovement                                         |
    #> | github/drme                                                             | github/dropclust                                                |
    #> | github/drseq2                                                           | github/drug-drug-interaction                                    |
    #> | github/drug-target-prediction                                           | github/drugcombination                                          |
    #> | github/drugcombinationprediction                                        | github/drugdiseasenet                                           |
    #> | github/drugtargetor                                                     | github/drugthatgene                                             |
    #> | github/drugz                                                            | github/drvae                                                    |
    #> | github/drwh                                                             | github/dsm-framework                                            |
    #> | github/dspred                                                           | github/dt-all                                                   |
    #> | github/dto                                                              | github/dtwscore                                                 |
    #> | github/duett                                                            | github/dugongbioinformatics                                     |
    #> | github/duphold                                                          | github/dwcox                                                    |
    #> | github/dycluster                                                        | github/dyconet                                                  |
    #> | github/dynadup                                                          | github/dynalogo                                                 |
    #> | github/dynamicdbg                                                       | github/dysprog                                                  |
    #> | github/e-driver                                                         | github/e-prospect                                               |
    #> | github/e2fm                                                             | github/eagle                                                    |
    #> | github/earomatic                                                        | github/easycodeml                                               |
    #> | github/ebcall                                                           | github/ebgsea                                                   |
    #> | github/ebi-metabolights                                                 | github/ebolavirussdpsbioinformatics2019                         |
    #> | github/ebwt2snp                                                         | github/echo                                                     |
    #> | github/ecogems                                                          | github/ecpred                                                   |
    #> | github/edetect                                                          | github/edge-in-tcga                                             |
    #> | github/edgescaping                                                      | github/edlib                                                    |
    #> | github/eecog-comp                                                       | github/eeg-annotate                                             |
    #> | github/eeg-spindles                                                     | github/effect-of-the-operational-parameters                     |
    #> | github/effusion                                                         | github/efindsite                                                |
    #> | github/efp-seq-browser                                                  | github/egad                                                     |
    #> | github/egcpi                                                            | github/ehr                                                      |
    #> | github/eigen-scripts                                                    | github/eklipse                                                  |
    #> | github/el-smurf                                                         | github/electron-micrograph-denoiser                             |
    #> | github/electronfactors                                                  | github/elemcor                                                  |
    #> | github/elm                                                              | github/elmeri                                                   |
    #> | github/emase                                                            | github/embeddings-reproduction                                  |
    #> | github/emblmygff3                                                       | github/emclarity                                                |
    #> | github/emdn                                                             | github/emdunifrac                                               |
    #> | github/emep                                                             | github/emerald                                                  |
    #> | github/emirge                                                           | github/emmckmer                                                 |
    #> | github/emolfrag                                                         | github/emringer                                                 |
    #> | github/emsar                                                            | github/emwer                                                    |
    #> | github/encoded                                                          | github/end2end-all-conv                                         |
    #> | github/enigma                                                           | github/enimpute                                                 |
    #> | github/ennet                                                            | github/enscat                                                   |
    #> | github/ensembl-rest                                                     | github/ensembler                                                |
    #> | github/enve                                                             | github/enzynet                                                  |
    #> | github/epa-ng                                                           | github/epanet-dev                                               |
    #> | github/epee                                                             | github/epem-gmm                                                 |
    #> | github/epga2                                                            | github/ephagen                                                  |
    #> | github/epiabc                                                           | github/epic2                                                    |
    #> | github/epicode                                                          | github/epicom                                                   |
    #> | github/epicseg                                                          | github/epidish                                                  |
    #> | github/epigbs                                                           | github/epigraph                                                 |
    #> | github/epilog                                                           | github/epinem                                                   |
    #> | github/epiprofile2.0-family                                             | github/epiq                                                     |
    #> | github/episafari                                                        | github/episurg                                                  |
    #> | github/epitracer                                                        | github/epivan                                                   |
    #> | github/eplmi                                                            | github/epof                                                     |
    #> | github/epooling                                                         | github/epsdc                                                    |
    #> | github/epslasso                                                         | github/eqp-cluster                                              |
    #> | github/eqtloverlapper                                                   | github/eqtlseq                                                  |
    #> | github/equation-generator                                               | github/equivalent-junctions                                     |
    #> | github/equivariantnetworks                                              | github/erd2                                                     |
    #> | github/error-correction                                                 | github/esigen                                                   |
    #> | github/esmo                                                             | github/esp-dnn                                                  |
    #> | github/espresso-research                                                | github/esprit2                                                  |
    #> | github/estitics                                                         | github/etcomp                                                   |
    #> | github/ethseq                                                           | github/etoxpred                                                 |
    #> | github/eval                                                             | github/evcouplings                                              |
    #> | github/evidenceaggregateddriverranking                                  | github/evidenceontology                                         |
    #> | github/eviltools                                                        | github/evntextrc                                                |
    #> | github/evoclust3d                                                       | github/evolclust                                                |
    #> | github/evolstruct-phogly                                                | github/ewas-twin-simulation                                     |
    #> | github/ewca                                                             | github/exmachina                                                |
    #> | github/exondel                                                          | github/exonsuite                                                |
    #> | github/expansionhunter                                                  | github/explobatch                                               |
    #> | github/express-d                                                        | github/exranges                                                 |
    #> | github/extramp                                                          | github/extrarg                                                  |
    #> | github/extreme                                                          | github/exttada                                                  |
    #> | github/ezmap                                                            | github/ezqsar                                                   |
    #> | github/eztree                                                           | github/f-anogan                                                 |
    #> | github/f3das                                                            | github/fab-phmm                                                 |
    #> | github/facets                                                           | github/factorial-hmm                                            |
    #> | github/factorizations                                                   | github/fait                                                     |
    #> | github/falco                                                            | github/famplex                                                  |
    #> | github/faot                                                             | github/fardeep                                                  |
    #> | github/fasp                                                             | github/fast                                                     |
    #> | github/fast-gep                                                         | github/fast-lmm                                                 |
    #> | github/fastbaps                                                         | github/fastbma                                                  |
    #> | github/fastderain                                                       | github/fastheatmap                                              |
    #> | github/fastlzerospikeinference                                          | github/fastore                                                  |
    #> | github/fastp                                                            | github/fastq-brew                                               |
    #> | github/fastqpuri                                                        | github/fastrfs                                                  |
    #> | github/fastspar                                                         | github/fatslim                                                  |
    #> | github/faucet                                                           | github/favites                                                  |
    #> | github/fbdenovo                                                         | github/fbm                                                      |
    #> | github/fcnv                                                             | github/fcs-point-correlator                                     |
    #> | github/fda-ars                                                          | github/feast                                                    |
    #> | github/featureselect                                                    | github/feelnc                                                   |
    #> | github/feizhe/pgs                                                       | github/fem                                                      |
    #> | github/fermi                                                            | github/fermikit                                                 |
    #> | github/fertility-gru                                                    | github/fetalfraction-snpimpute                                  |
    #> | github/few-shot-segmentation                                            | github/fgcs17                                                   |
    #> | github/fiddlercrab.info                                                 | github/figs                                                     |
    #> | github/filtus                                                           | github/fimtrack                                                 |
    #> | github/findgco                                                          | github/findgse                                                  |
    #> | github/findis-db                                                        | github/findr                                                    |
    #> | github/fineradstructure                                                 | github/fiqt                                                     |
    #> | github/fisheye-orb-slam                                                 | github/fishplot                                                 |
    #> | github/fit-sne                                                          | github/fitild                                                   |
    #> | github/fizzy                                                            | github/flaimapper                                               |
    #> | github/flas                                                             | github/flashpca                                                 |
    #> | github/flashweave.jl                                                    | github/fldgen                                                   |
    #> | github/flea-pipeline                                                    | github/flea-web-app                                             |
    #> | github/flexbar                                                          | github/flexgeo                                                  |
    #> | github/flexicoclusteringpackage                                         | github/flexidot                                                 |
    #> | github/flexor                                                           | github/flib-coevo                                               |
    #> | github/flint                                                            | github/flnadd                                                   |
    #> | github/floc                                                             | github/flowermorphology                                         |
    #> | github/flowgrid                                                         | github/flowlearn                                                |
    #> | github/flowpeaks                                                        | github/flowsom                                                  |
    #> | github/flps                                                             | github/fluff                                                    |
    #> | github/fluidigm2purc                                                    | github/fluka-ct                                                 |
    #> | github/flycop                                                           | github/fmlc                                                     |
    #> | github/fmoc                                                             | github/fmrat                                                    |
    #> | github/fmtree                                                           | github/fnbtools                                                 |
    #> | github/foldatlas                                                        | github/foresee                                                  |
    #> | github/foreseqs                                                         | github/forwardgenomics                                          |
    #> | github/founder-sequences                                                | github/fourrussiansrnafolding                                   |
    #> | github/fp2vec                                                           | github/fpkit                                                    |
    #> | github/fqc                                                              | github/fqtools                                                  |
    #> | github/fragalcode                                                       | github/fragbuilder                                              |
    #> | github/frama                                                            | github/framework                                                |
    #> | github/franklin                                                         | github/frcnn-cad                                                |
    #> | github/friends-and-family                                               | github/frogs                                                    |
    #> | github/fromage                                                          | github/fsda                                                     |
    #> | github/fselector                                                        | github/fsepsa                                                   |
    #> | github/fsmkl                                                            | github/fsqn                                                     |
    #> | github/fssemr                                                           | github/fsva                                                     |
    #> | github/fta                                                              | github/ftc                                                      |
    #> | github/fuento                                                           | github/fuma                                                     |
    #> | github/fumoso                                                           | github/functree-ng                                              |
    #> | github/fundus-fractal-analysis                                          | github/fungap                                                   |
    #> | github/funmap2                                                          | github/funmappone                                               |
    #> | github/funnel-gsea-r-package                                            | github/funpred-3.0                                              |
    #> | github/funtoonorm                                                       | github/funvar                                                   |
    #> | github/fusenet                                                          | github/fuseq                                                    |
    #> | github/g-msr                                                            | github/g2p                                                      |
    #> | github/g3lollipop.js                                                    | github/g4catchall                                               |
    #> | github/ga-bagging-svm                                                   | github/gac                                                      |
    #> | github/galaxy                                                           | github/galaxy-rna-workbench                                     |
    #> | github/galaxydocker                                                     | github/gap                                                      |
    #> | github/gapest                                                           | github/gappa                                                    |
    #> | github/gappadder                                                        | github/gapreduce                                                |
    #> | github/garfield-ngs                                                     | github/garleek                                                  |
    #> | github/garlic                                                           | github/gasol                                                    |
    #> | github/gatc                                                             | github/gated-xnor                                               |
    #> | github/gatekeeper                                                       | github/gatk                                                     |
    #> | github/gaudi                                                            | github/gavin-yinld                                              |
    #> | github/gaynor-sun-gbj-breast-cancer                                     | github/gaze3dfix                                                |
    #> | github/gbgfa                                                            | github/gbs-pacecar                                              |
    #> | github/gbs-snp-crop                                                     | github/gbwt                                                     |
    #> | github/gc                                                               | github/gcevobase                                                |
    #> | github/gcmdr                                                            | github/gcsa2                                                    |
    #> | github/gdcrnatools                                                      | github/ge-lda-survival                                          |
    #> | github/geck                                                             | github/ged-lab/khmer                                            |
    #> | github/gedfn                                                            | github/geiger-v2                                                |
    #> | github/gembassy                                                         | github/gemini                                                   |
    #> | github/gemmer                                                           | github/gempro                                                   |
    #> | github/gemprot                                                          | github/gemse                                                    |
    #> | github/gemsplice-code                                                   | github/gemtools                                                 |
    #> | github/gencof                                                           | github/gene-block-evolution                                     |
    #> | github/gene-is                                                          | github/gene-surrounder                                          |
    #> | github/genecnv                                                          | github/geneevolve                                               |
    #> | github/genefu                                                           | github/genefuse                                                 |
    #> | github/genegraphics                                                     | github/genenotebook                                             |
    #> | github/genepheno                                                        | github/genepopedit                                              |
    #> | github/genesthycan                                                      | github/genesurv                                                 |
    #> | github/geneticscreen.jl                                                 | github/geneticthesaurus                                         |
    #> | github/gengis                                                           | github/genhap                                                   |
    #> | github/genmap-comparator                                                | github/gennet                                                   |
    #> | github/geno-diver                                                       | github/genocore                                                 |
    #> | github/genome-bin-tools                                                 | github/genome-maps                                              |
    #> | github/genome-sketching                                                 | github/genome-topology-network                                  |
    #> | github/genomealignmenttools                                             | github/genomeartist                                             |
    #> | github/genomed3plot                                                     | github/genomedisco                                              |
    #> | github/genomeflow                                                       | github/genomemining                                             |
    #> | github/genomescope                                                      | github/genometester4                                            |
    #> | github/genomeuplot                                                      | github/genomewarp                                               |
    #> | github/genomicgraphcoords                                               | github/genomon-itdetector                                       |
    #> | github/genotype-corrector                                               | github/genplay                                                  |
    #> | github/genssi                                                           | github/georacle                                                 |
    #> | github/george                                                           | github/georgettetanner                                          |
    #> | github/gep2pep                                                          | github/get-phylomarkers                                         |
    #> | github/getisdmr                                                         | github/gevitanalysisrelease                                     |
    #> | github/gfakluge                                                         | github/gfapy                                                    |
    #> | github/gfaviz                                                           | github/gfcom                                                    |
    #> | github/gff3toolkit                                                      | github/gfplain                                                  |
    #> | github/gfusion                                                          | github/ggm-shrinkage                                            |
    #> | github/ggrasp                                                           | github/ggsashimi                                                |
    #> | github/ggtk                                                             | github/ggv                                                      |
    #> | github/ggv-api                                                          | github/ghost-tree                                               |
    #> | github/ght                                                              | github/gi-cluster                                               |
    #> | github/gift                                                             | github/giggle                                                   |
    #> | github/gimpute                                                          | github/gips                                                     |
    #> | github/giremi                                                           | github/github-bioinformatics                                    |
    #> | github/github.com:                                                      | github/gkmexplain                                               |
    #> | github/gladiatox                                                        | github/glanet                                                   |
    #> | github/glass                                                            | github/glep                                                     |
    #> | github/glint                                                            | github/glmvc                                                    |
    #> | github/glp                                                              | github/glsnn                                                    |
    #> | github/glutton                                                          | github/glycanformatconverter                                    |
    #> | github/glycoseq                                                         | github/glycresoft                                               |
    #> | github/glyxtoolms                                                       | github/gm-cml                                                   |
    #> | github/gma                                                              | github/gmalign                                                  |
    #> | github/gmm-marker-tracking                                              | github/gmrad                                                    |
    #> | github/gms                                                              | github/gmxapi                                                   |
    #> | github/gne                                                              | github/gneiss                                                   |
    #> | github/gnetlmm                                                          | github/gnn                                                      |
    #> | github/gnuplot-iostream                                                 | github/goatools                                                 |
    #> | github/gobe                                                             | github/gobyweb2-plugins                                         |
    #> | github/goeckslab                                                        | github/gofeat                                                   |
    #> | github/gokit                                                            | github/goldilocks                                               |
    #> | github/goleft                                                           | github/gooogle                                                  |
    #> | github/gopher                                                           | github/gorampage                                                |
    #> | github/gospel                                                           | github/gotrapper                                                |
    #> | github/gpa-package                                                      | github/gpca                                                     |
    #> | github/gpds                                                             | github/gpga                                                     |
    #> | github/gpica                                                            | github/gpmat                                                    |
    #> | github/gpmicrobiome                                                     | github/gppf                                                     |
    #> | github/gppfst                                                           | github/gpqtlmapping                                             |
    #> | github/gprotpred                                                        | github/gpseudoclust                                             |
    #> | github/gpseudorank                                                      | github/gptime                                                   |
    #> | github/gpu-daemon                                                       | github/gqt                                                      |
    #> | github/grabb                                                            | github/graftm-gpkgs                                             |
    #> | github/grail                                                            | github/grailbio-bio                                             |
    #> | github/gram-cnn                                                         | github/grandprix                                                |
    #> | github/graph-annotation                                                 | github/graph2pro                                                |
    #> | github/graph2tab                                                        | github/graphaligner                                             |
    #> | github/graphlet-naming                                                  | github/graphmap                                                 |
    #> | github/graphspace                                                       | github/grassmanncluster                                         |
    #> | github/greenwoodlab                                                     | github/grein                                                    |
    #> | github/grep                                                             | github/grepsearch                                               |
    #> | github/gridlmm                                                          | github/grimm                                                    |
    #> | github/grimon                                                           | github/grnmf                                                    |
    #> | github/grnte                                                            | github/grnvbem                                                  |
    #> | github/gromacs-mic                                                      | github/grompy                                                   |
    #> | github/groot                                                            | github/grouper                                                  |
    #> | github/groupk                                                           | github/growthestimate                                           |
    #> | github/grpslopemc                                                       | github/grridgecodata                                            |
    #> | github/gsa4mirna                                                        | github/gsalightning                                             |
    #> | github/gsbc                                                             | github/gsca                                                     |
    #> | github/gsea-incontext                                                   | github/gsia                                                     |
    #> | github/gsimp                                                            | github/gsl                                                      |
    #> | github/gsmodutils                                                       | github/gsp                                                      |
    #> | github/gsso                                                             | github/gsu                                                      |
    #> | github/gtc                                                              | github/gtdbtk                                                   |
    #> | github/gtm-generativetopographicmapping                                 | github/gtotree                                                  |
    #> | github/gtrac                                                            | github/gtshark                                                  |
    #> | github/gtz                                                              | github/gubbins                                                  |
    #> | github/guoweifeng                                                       | github/guv-ap                                                   |
    #> | github/gwa-tutorial                                                     | github/gwalpha                                                  |
    #> | github/gwas-scripts                                                     | github/gwasrapidd                                               |
    #> | github/gwc                                                              | github/gxbrowser                                                |
    #> | github/h-blast                                                          | github/h-popg                                                   |
    #> | github/hal                                                              | github/halc                                                     |
    #> | github/hammock                                                          | github/haoran2014                                               |
    #> | github/haploclique                                                      | github/haploconduct                                             |
    #> | github/haploforge                                                       | github/haplomerger2                                             |
    #> | github/hapmc                                                            | github/hapmuc                                                   |
    #> | github/happe                                                            | github/happytools                                               |
    #> | github/harc                                                             | github/harmony                                                  |
    #> | github/harvest                                                          | github/hasappy                                                  |
    #> | github/hasn                                                             | github/hatspil                                                  |
    #> | github/haveyouswappedyoursamples                                        | github/hayai-annotation-plants                                  |
    #> | github/haystack-bio                                                     | github/hbm                                                      |
    #> | github/hcsa                                                             | github/hdx-viewer                                               |
    #> | github/heap                                                             | github/hedicim                                                  |
    #> | github/hefpipe-rpos                                                     | github/helmsman                                                 |
    #> | github/helo                                                             | github/hemodonacion                                             |
    #> | github/herg-chembl-jcim                                                 | github/hergepred                                                |
    #> | github/heritability                                                     | github/hetero-rp                                                |
    #> | github/heteroplasmyworkflow                                             | github/hfufs                                                    |
    #> | github/hg-color                                                         | github/hgtector                                                 |
    #> | github/hgtsim                                                           | github/hgvs                                                     |
    #> | github/hh-suite                                                         | github/hhcompare                                                |
    #> | github/hhconpred                                                        | github/hic-pro                                                  |
    #> | github/hic-spector                                                      | github/hicap                                                    |
    #> | github/hicaptools                                                       | github/hiceekr                                                  |
    #> | github/hicgan                                                           | github/hiclipr                                                  |
    #> | github/hicnvtranslocation                                               | github/hicomet                                                  |
    #> | github/hicplotter                                                       | github/hictranshi-c                                             |
    #> | github/hierarchical-hotnet                                              | github/hierarchical-rnns-model-for-ddi-extraction               |
    #> | github/hierarchical3dgenome                                             | github/hifresp                                                  |
    #> | github/high-quality-ellipse-detection                                   | github/hihmm                                                    |
    #> | github/hima                                                             | github/himc                                                     |
    #> | github/hintra                                                           | github/hippodeep                                                |
    #> | github/hipred                                                           | github/hirgc                                                    |
    #> | github/historian                                                        | github/hitime                                                   |
    #> | github/hitselect                                                        | github/hitwalker2                                               |
    #> | github/hiv-gag-immunogens                                               | github/hivmmer                                                  |
    #> | github/hla-bind                                                         | github/hla-la                                                   |
    #> | github/hlama                                                            | github/hlaprofiler                                              |
    #> | github/hlatyphon                                                        | github/hlmethy                                                  |
    #> | github/hmc-grt                                                          | github/hmeta-d                                                  |
    #> | github/hmm-based-method                                                 | github/hmmerctter                                               |
    #> | github/hmmpe                                                            | github/hmtgo                                                    |
    #> | github/hnematzadeh                                                      | github/hnmdrp                                                   |
    #> | github/hnmf                                                             | github/hocnnlb                                                  |
    #> | github/hodgkin-huxley-spde                                              | github/hogmmnc                                                  |
    #> | github/hollytibble                                                      | github/homblocks                                                |
    #> | github/home                                                             | github/hopland                                                  |
    #> | github/hops                                                             | github/horvathlab-ngs                                           |
    #> | github/hot-scan                                                         | github/hotspot                                                  |
    #> | github/howdesbt                                                         | github/hpc-workflow-manager                                     |
    #> | github/hpccs                                                            | github/hpg-aligner                                              |
    #> | github/hpg-pore                                                         | github/hpo2go                                                   |
    #> | github/hppi-tensorflow                                                  | github/hpviewer                                                 |
    #> | github/hr1912/treeexp                                                   | github/hsa                                                      |
    #> | github/hsesumo                                                          | github/htdp                                                     |
    #> | github/hts-barcode-checker                                              | github/hts-nim                                                  |
    #> | github/hts-nim-tools                                                    | github/htsint                                                   |
    #> | github/htssip                                                           | github/htsvis                                                   |
    #> | github/hulk                                                             | github/hupan                                                    |
    #> | github/hurdlenormal                                                     | github/hvisaclassifier                                          |
    #> | github/hvsm                                                             | github/hyasp                                                    |
    #> | github/hyb                                                              | github/hybphylomaker                                            |
    #> | github/hybpiper                                                         | github/hybriddetective                                          |
    #> | github/hyde                                                             | github/hydra                                                    |
    #> | github/hyper                                                            | github/hyper-docs                                               |
    #> | github/hypercubeme                                                      | github/hypergate                                                |
    #> | github/hypergraphdynamiccorrelation                                     | github/hypocotyl-unet                                           |
    #> | github/hyppo                                                            | github/hysxe/acme                                               |
    #> | github/ia-lab                                                           | github/iamcomparison                                            |
    #> | github/iatc-nrakel                                                      | github/ibiomes                                                  |
    #> | github/ibpp                                                             | github/ibrel                                                    |
    #> | github/icb-docker                                                       | github/icepop                                                   |
    #> | github/ichseg                                                           | github/icma                                                     |
    #> | github/icn3d                                                            | github/icnv                                                     |
    #> | github/icopydav                                                         | github/icr142-benchmarker                                       |
    #> | github/ideepe                                                           | github/ideeps                                                   |
    #> | github/identification-of-brain-based-disorders                          | github/identificationwitharme                                   |
    #> | github/idepi                                                            | github/idhi-mirw                                                |
    #> | github/idingo                                                           | github/idlp                                                     |
    #> | github/idssr                                                            | github/if                                                       |
    #> | github/ifeature                                                         | github/iform                                                    |
    #> | github/igb-mi-bundle                                                    | github/igess                                                    |
    #> | github/igloo                                                            | github/iguide                                                   |
    #> | github/iham                                                             | github/iht.jl                                                   |
    #> | github/ikap                                                             | github/illumina-utils                                           |
    #> | github/image-engine                                                     | github/image-quality-evaluation                                 |
    #> | github/imagebasedtranscriptomics                                        | github/imagebox                                                 |
    #> | github/imagecn                                                          | github/imagepy                                                  |
    #> | github/imaging-of-drugs-distribution-and-quantifications                | github/imap                                                     |
    #> | github/imapsplice                                                       | github/imcmda                                                   |
    #> | github/imglib                                                           | github/imgui-electron-packages                                  |
    #> | github/immboost                                                         | github/immclass2019                                             |
    #> | github/immucc                                                           | github/immune-deconvolution-benchmark                           |
    #> | github/immunedb                                                         | github/immunedeconv                                             |
    #> | github/impact                                                           | github/imperfect-gold-standard                                  |
    #> | github/impre                                                            | github/improver2013                                             |
    #> | github/imputation                                                       | github/ims-informed-library                                     |
    #> | github/ims-project                                                      | github/ims-quality                                              |
    #> | github/imsa-a                                                           | github/imspectr                                                 |
    #> | github/imtrbm                                                           | github/indecut                                                  |
    #> | github/index                                                            | github/index-app-public                                         |
    #> | github/indra-pathway-map                                                | github/infinity                                                 |
    #> | github/infiniumpurify                                                   | github/infmod3dgen                                              |
    #> | github/inform                                                           | github/inmembrane                                               |
    #> | github/inmf                                                             | github/innate2adaptive/decombinator                             |
    #> | github/inpactor                                                         | github/insilico-subtyping                                       |
    #> | github/insilicoseq                                                      | github/inspiired                                                |
    #> | github/intarna                                                          | github/integrate                                                |
    #> | github/integrate-neo                                                    | github/integrate-vis                                            |
    #> | github/integratedphasing                                                | github/integrative                                              |
    #> | github/integrativeanalysis                                              | github/integreat                                                |
    #> | github/interactiverosetta                                               | github/interface-classifier                                     |
    #> | github/interminer                                                       | github/intlim                                                   |
    #> | github/invertinggan                                                     | github/ionmf                                                    |
    #> | github/ionspattern                                                      | github/ipa                                                      |
    #> | github/ipbt                                                             | github/ipct                                                     |
    #> | github/iphloc-es                                                        | github/ipo                                                      |
    #> | github/ipsa                                                             | github/irap                                                     |
    #> | github/irecspot                                                         | github/irscope                                                  |
    #> | github/irspot-sf                                                        | github/is-cellr                                                 |
    #> | github/isa-tools                                                        | github/isc                                                      |
    #> | github/iscore                                                           | github/isdb                                                     |
    #> | github/iseeu                                                            | github/isescan                                                  |
    #> | github/isglobal-brge/lambda                                             | github/isgpt                                                    |
    #> | github/ishidalab-titech/3dcnn-mqa                                       | github/ishspsy-project                                          |
    #> | github/iskana/pbwt-sec                                                  | github/islandpath                                               |
    #> | github/ismags                                                           | github/ismb15                                                   |
    #> | github/ismb2017-lstm                                                    | github/iso2flux                                                 |
    #> | github/isocket                                                          | github/isocor                                                   |
    #> | github/isoem2                                                           | github/isofishr                                                 |
    #> | github/isomir2function                                                  | github/isop                                                     |
    #> | github/isoscm                                                           | github/isosegmenter                                             |
    #> | github/isotope-correction-toolbox                                       | github/isotopiclabelling                                        |
    #> | github/isotree                                                          | github/isown                                                    |
    #> | github/ispp                                                             | github/issues                                                   |
    #> | github/iswathx                                                          | github/iterassemble                                             |
    #> | github/iterativeerrorcorrection                                         | github/ivc                                                      |
    #> | github/ivis                                                             | github/ivus-ultrasonic                                          |
    #> | github/iwgs                                                             | github/jackhou2/c3                                              |
    #> | github/jaffa                                                            | github/jama16-retina-replication                                |
    #> | github/james-e-barrett                                                  | github/jami                                                     |
    #> | github/jamss                                                            | github/jannovar                                                 |
    #> | github/janus                                                            | github/japsa                                                    |
    #> | github/jarvis                                                           | github/jaxbd2k-shortcourse                                      |
    #> | github/jbcb2018                                                         | github/jbudis/lambda                                            |
    #> | github/jcvi                                                             | github/jdinac                                                   |
    #> | github/jepegmix2                                                        | github/jesse                                                    |
    #> | github/jester                                                           | github/jinchaofeng-code                                         |
    #> | github/jingwyang/treeexp                                                | github/jiyuanhu                                                 |
    #> | github/jms                                                              | github/jmztab-m                                                 |
    #> | github/joa                                                              | github/johnbarton/ace                                           |
    #> | github/joint-marker-free-alignment                                      | github/jointrn                                                  |
    #> | github/joker                                                            | github/jomayer/smurf                                            |
    #> | github/jrgui                                                            | github/jsms                                                     |
    #> | github/juchmme                                                          | github/juliapalacios/phylodyn                                   |
    #> | github/justorthologs                                                    | github/juzhe1120-matlab-software                                |
    #> | github/jvarkit                                                          | github/kablammo                                                 |
    #> | github/kalign                                                           | github/kanizsa-prime                                            |
    #> | github/kaptive                                                          | github/kaptive-web                                              |
    #> | github/kart                                                             | github/kasimcloud                                               |
    #> | github/kasimjs                                                          | github/kat                                                      |
    #> | github/kbet                                                             | github/kbmtl                                                    |
    #> | github/kbranches                                                        | github/kbws                                                     |
    #> | github/kcmbt-mt                                                         | github/kdetrees                                                 |
    #> | github/kdsnp                                                            | github/keanu                                                    |
    #> | github/kekulescope                                                      | github/kema                                                     |
    #> | github/kent                                                             | github/keratocyte                                               |
    #> | github/kernelized-rank-learning                                         | github/kescases                                                 |
    #> | github/kestrel                                                          | github/keyregulatorygenes                                       |
    #> | github/kf-arabidopsis-grna                                              | github/kfits                                                    |
    #> | github/kgs                                                              | github/kinact                                                   |
    #> | github/kinconform                                                       | github/kinmut2                                                  |
    #> | github/kinship-privacy                                                  | github/kipper                                                   |
    #> | github/kirp-topological-features                                        | github/kltepigenome                                             |
    #> | github/km2gcn                                                           | github/kmcex                                                    |
    #> | github/kmer-db                                                          | github/kmer-ssr                                                 |
    #> | github/kmerind                                                          | github/kmerpyramid                                              |
    #> | github/kneeanalysis                                                     | github/knn                                                      |
    #> | github/knot-pull                                                        | github/knoto-id                                                 |
    #> | github/knotty                                                           | github/knowledge-elicitation-for-precision-medicine             |
    #> | github/knowledgediscovery                                               | github/kodicollins/pasta                                        |
    #> | github/kollector                                                        | github/kover                                                    |
    #> | github/kpal                                                             | github/kpax3.jl                                                 |
    #> | github/kpic2                                                            | github/kpm                                                      |
    #> | github/krab-znf                                                         | github/krait                                                    |
    #> | github/kraken                                                           | github/krakenuniq                                               |
    #> | github/kreationsupplementary                                            | github/krm15/acme                                               |
    #> | github/kronos                                                           | github/ksea                                                     |
    #> | github/ksp-puel                                                         | github/ksrepo                                                   |
    #> | github/kubenow-plugin                                                   | github/kubernetes/charts                                        |
    #> | github/kurtosis-conservation                                            | github/kusterlab                                                |
    #> | github/kwip                                                             | github/kyzhu/assembltrie                                        |
    #> | github/l0adridge                                                        | github/l1em                                                     |
    #> | github/la-isla-de-tomato                                                | github/lacroixlaurent                                           |
    #> | github/lacytools                                                        | github/laggedorderedlassonetwork                                |
    #> | github/lambda                                                           | github/lamoureux-lab/3dcnn-mqa                                  |
    #> | github/lamsa                                                            | github/lancet                                                   |
    #> | github/largegopred                                                      | github/larvalign                                                |
    #> | github/lasagne4bio                                                      | github/lattice-metage                                           |
    #> | github/laura-orellana/ebdims                                            | github/lb-impute                                                |
    #> | github/lbeep                                                            | github/lc-ms-pachyderm                                          |
    #> | github/lc50                                                             | github/lca-cnn                                                  |
    #> | github/lcslcis                                                          | github/lcv                                                      |
    #> | github/ld                                                               | github/ldjump                                                   |
    #> | github/ldserv                                                           | github/leaf                                                     |
    #> | github/leaf-gp                                                          | github/leap                                                     |
    #> | github/learningdl                                                       | github/ledpred                                                  |
    #> | github/lem                                                              | github/lemon                                                    |
    #> | github/lemon-tree                                                       | github/lenup                                                    |
    #> | github/lep                                                              | github/lfc                                                      |
    #> | github/lgcm                                                             | github/lhqxinghun-bioinformatics                                |
    #> | github/lhqxinghun/bioinformatics                                        | github/libchebi                                                 |
    #> | github/libcsam                                                          | github/libgaba                                                  |
    #> | github/libgtftk                                                         | github/libra-x                                                  |
    #> | github/librfn                                                           | github/libstructural                                            |
    #> | github/lifs-tools                                                       | github/lightassembler                                           |
    #> | github/lightdock                                                        | github/lightdock-bm5                                            |
    #> | github/likelihoodfreephylogenetics                                      | github/lilikoi                                                  |
    #> | github/limbr                                                            | github/limitfluxtocore                                          |
    #> | github/limix                                                            | github/lims                                                     |
    #> | github/lincs-rnaseq-cpp                                                 | github/linearfold                                               |
    #> | github/linker                                                           | github/linkhd                                                   |
    #> | github/lionessr                                                         | github/lipase-reclassification                                  |
    #> | github/lipidfinder                                                      | github/lipidminion                                              |
    #> | github/liquid                                                           | github/lir                                                      |
    #> | github/lis-context-viewer                                               | github/litemol                                                  |
    #> | github/lll-fva                                                          | github/lm-lstm-crf                                              |
    #> | github/lme4qtl                                                          | github/lmmel-mir-miner                                          |
    #> | github/lmmo                                                             | github/lncadeep                                                 |
    #> | github/lncdiff                                                          | github/lncmirsrn                                                |
    #> | github/lncrna-disease-link                                              | github/lncrna-homologs                                          |
    #> | github/lncrna-id                                                        | github/lncscore                                                 |
    #> | github/lobstahs                                                         | github/localtadsim                                              |
    #> | github/locings                                                          | github/locnuclei                                                |
    #> | github/locusexplorer                                                    | github/locustreeinference                                       |
    #> | github/logical-modelling-pipeline                                       | github/logloss-beraf                                            |
    #> | github/lollipops                                                        | github/lolopicker                                               |
    #> | github/longitudinal-microbiome-analysis-public                          | github/longo                                                    |
    #> | github/look4trs                                                         | github/lordfast                                                 |
    #> | github/loregic                                                          | github/lorisnanni                                               |
    #> | github/lpg                                                              | github/lrcstats                                                 |
    #> | github/lrdkt                                                            | github/lrf-dtis                                                 |
    #> | github/lrsim                                                            | github/lrssl                                                    |
    #> | github/lsdbs-mpi                                                        | github/lsgkm                                                    |
    #> | github/lsm-worker                                                       | github/lsmm                                                     |
    #> | github/lstd                                                             | github/lstmvoter                                                |
    #> | github/lstrap-lite                                                      | github/lsue                                                     |
    #> | github/ltmgsca                                                          | github/ltr-retriever                                            |
    #> | github/lumpy-sv                                                         | github/lvq-knn                                                  |
    #> | github/lw-fqzip2                                                        | github/lwlrtest                                                 |
    #> | github/lyve-set                                                         | github/lzerospikeinference                                      |
    #> | github/lzw-kernel                                                       | github/m2align                                                  |
    #> | github/m2c                                                              | github/m2c-rel-nb                                               |
    #> | github/m3drop                                                           | github/ma-compbio/multires                                      |
    #> | github/macarthur-lab/clinvar                                            | github/macau                                                    |
    #> | github/mackinac                                                         | github/macs                                                     |
    #> | github/macsyfinder                                                      | github/mad-hype                                                 |
    #> | github/magellan                                                         | github/magetbrain                                               |
    #> | github/magi                                                             | github/magi-s                                                   |
    #> | github/magic                                                            | github/magpy                                                    |
    #> | github/magus                                                            | github/maize-tes                                                |
    #> | github/maize-tissue-specific-grn                                        | github/maizedig                                                 |
    #> | github/maizesnpdb                                                       | github/maizete-variation                                        |
    #> | github/malax                                                            | github/maligner                                                 |
    #> | github/malini                                                           | github/manga                                                    |
    #> | github/manta                                                            | github/manuscript-avsec-bioinformatics-2017                     |
    #> | github/map3k8-thyroid-spheres-v-3.0                                     | github/mapcomp                                                  |
    #> | github/mapexr                                                           | github/mapoptics                                                |
    #> | github/mappi-dat                                                        | github/maprepeat                                                |
    #> | github/maps                                                             | github/mapseq                                                   |
    #> | github/maracluster                                                      | github/marathon                                                 |
    #> | github/mark2curedata                                                    | github/marker2sequence                                          |
    #> | github/markyt                                                           | github/marsi                                                    |
    #> | github/marvel                                                           | github/marvin                                                   |
    #> | github/masashitsubaki                                                   | github/mascot                                                   |
    #> | github/mash                                                             | github/mashmap                                                  |
    #> | github/mason                                                            | github/mass                                                     |
    #> | github/mass-simulator                                                   | github/masscomp                                                 |
    #> | github/massif                                                           | github/massytools                                               |
    #> | github/mast                                                             | github/mastermsm                                                |
    #> | github/matam                                                            | github/matataki                                                 |
    #> | github/matched-forest                                                   | github/mathbiocu                                                |
    #> | github/mathelab/ramp-db                                                 | github/matrixconverter                                          |
    #> | github/matrixepistasis                                                  | github/matryoshka                                               |
    #> | github/matscholar                                                       | github/maxent-ppi                                               |
    #> | github/mayonlppapcdss                                                   | github/mbeccuti/pgs                                             |
    #> | github/mbirw                                                            | github/mc-cbn                                                   |
    #> | github/mcclintock                                                       | github/mccortex                                                 |
    #> | github/mcee-2.0                                                         | github/mcimpute-scrnaseq                                        |
    #> | github/mcindoor20000                                                    | github/mcmc-ce-codes                                            |
    #> | github/mcorr                                                            | github/mcs                                                      |
    #> | github/mcsc-decontamination                                             | github/mcsmrt                                                   |
    #> | github/mct                                                              | github/mctandem                                                 |
    #> | github/mctcodes                                                         | github/mctl                                                     |
    #> | github/mcts-rna                                                         | github/md-task                                                  |
    #> | github/mda-cnn                                                          | github/mdine                                                    |
    #> | github/mdkarcher/phylodyn                                               | github/mdnnmd                                                   |
    #> | github/mdpbiome                                                         | github/mdplot                                                   |
    #> | github/mds2                                                             | github/mdts                                                     |
    #> | github/mea                                                              | github/mea-tools                                                |
    #> | github/means                                                            | github/mec                                                      |
    #> | github/mecat                                                            | github/mecors                                                   |
    #> | github/med4way                                                          | github/medcrawler                                               |
    #> | github/medestrand                                                       | github/medication-qa-medinfo2019                                |
    #> | github/medici                                                           | github/medline-code-v2                                          |
    #> | github/medline-collaboration-datasets                                   | github/medtest                                                  |
    #> | github/medusa                                                           | github/meffil                                                   |
    #> | github/mefit                                                            | github/mega                                                     |
    #> | github/megagta                                                          | github/megan-ce                                                 |
    #> | github/meltos                                                           | github/memo                                                     |
    #> | github/mendelian                                                        | github/mendelprob                                               |
    #> | github/mentalist                                                        | github/mepurity                                                 |
    #> | github/mer                                                              | github/meristogram                                              |
    #> | github/merlot                                                           | github/meshsim                                                  |
    #> | github/meta-gdbp                                                        | github/meta-proteome-analyzer                                   |
    #> | github/metaassemblyeval                                                 | github/metabmf                                                  |
    #> | github/metaboanalystr                                                   | github/metabocraft                                              |
    #> | github/metabodiff                                                       | github/metabolicframework                                       |
    #> | github/metabolomics-filtering                                           | github/metabolomicstools                                        |
    #> | github/metabridge-shiny                                                 | github/metacache                                                |
    #> | github/metacherchant                                                    | github/metachip                                                 |
    #> | github/metacluster                                                      | github/metacomp                                                 |
    #> | github/metacrast                                                        | github/metacrispr                                               |
    #> | github/metacycle                                                        | github/metadiff                                                 |
    #> | github/metaerg                                                          | github/metafast                                                 |
    #> | github/metagen                                                          | github/metagenome-pfam-score                                    |
    #> | github/metagenomic-benchmark                                            | github/metago                                                   |
    #> | github/metagrn                                                          | github/metakallisto                                             |
    #> | github/metalab                                                          | github/metamos                                                  |
    #> | github/metamsd                                                          | github/metaomics                                                |
    #> | github/metapalette                                                      | github/metapathways2                                            |
    #> | github/metapheno                                                        | github/metaphinder                                              |
    #> | github/metaplotr                                                        | github/metapoap                                                 |
    #> | github/metapred2cs                                                      | github/metaqubic                                                |
    #> | github/metarep                                                          | github/metaseek                                                 |
    #> | github/metashot                                                         | github/metasmc                                                  |
    #> | github/metaspark                                                        | github/metasra-pipeline                                         |
    #> | github/metatopics                                                       | github/metawrap                                                 |
    #> | github/metaxcan                                                         | github/metcirc                                                  |
    #> | github/metdiff                                                          | github/metgem                                                   |
    #> | github/methcomp                                                         | github/methrafo                                                 |
    #> | github/methylcode                                                       | github/methylflow                                               |
    #> | github/methyliftover                                                    | github/methylimp                                                |
    #> | github/methylmaps                                                       | github/metica                                                   |
    #> | github/metitree                                                         | github/metlab                                                   |
    #> | github/metpeak                                                          | github/metquest                                                 |
    #> | github/metqy                                                            | github/metriculator                                             |
    #> | github/metumpx-bin                                                      | github/mexcowalk                                                |
    #> | github/mfeprimer                                                        | github/mfmd                                                     |
    #> | github/mfsba                                                            | github/mgicnn                                                   |
    #> | github/mgpfusion                                                        | github/mgrfe-garfe                                              |
    #> | github/mhc-typer                                                        | github/mhfp                                                     |
    #> | github/mhn                                                              | github/mhtrajectoryr                                            |
    #> | github/miams                                                            | github/mic-meme                                                 |
    #> | github/micado                                                           | github/micmic                                                   |
    #> | github/micop                                                            | github/micro                                                    |
    #> | github/microaggregation-based-anonymization-tool                        | github/microbe-id                                               |
    #> | github/microbedb                                                        | github/microbial-conditions-ontology                            |
    #> | github/microbiome-fixed-reference                                       | github/microbiome-helper                                        |
    #> | github/microbiome-v2.1                                                  | github/microbiomedda                                            |
    #> | github/microbiomegwas                                                   | github/microbiota-resistome                                     |
    #> | github/micropro                                                         | github/microrna-nhpred                                          |
    #> | github/microscope                                                       | github/mid-correct                                              |
    #> | github/miergolf                                                         | github/miic                                                     |
    #> | github/mills-lab/vapor                                                  | github/mim                                                      |
    #> | github/mimeanto                                                         | github/mimi                                                     |
    #> | github/mina                                                             | github/mind                                                     |
    #> | github/mines                                                            | github/miniasm                                                  |
    #> | github/minicom                                                          | github/minimap                                                  |
    #> | github/minimap2                                                         | github/minimds                                                  |
    #> | github/minimizers                                                       | github/minion-qc                                                |
    #> | github/mintmap                                                          | github/mio                                                      |
    #> | github/mionsite                                                         | github/mipepid                                                  |
    #> | github/miphy                                                            | github/mipup                                                    |
    #> | github/mir-prefer                                                       | github/mira                                                     |
    #> | github/miraimplications                                                 | github/mirfinder                                                |
    #> | github/mirfix                                                           | github/mirge                                                    |
    #> | github/mirgff3                                                          | github/mirlocator                                               |
    #> | github/mirmark                                                          | github/mirmeta                                                  |
    #> | github/mirmmr                                                           | github/mirseqviewer                                             |
    #> | github/mirtardann                                                       | github/mirtime                                                  |
    #> | github/mirtop                                                           | github/mirtrace                                                 |
    #> | github/mirureader                                                       | github/mirvial                                                  |
    #> | github/miscoto                                                          | github/misfinder                                                |
    #> | github/miso-example                                                     | github/mispu                                                    |
    #> | github/mistic                                                           | github/mitefinder                                               |
    #> | github/mitobim                                                          | github/mitodix                                                  |
    #> | github/mitoimp                                                          | github/mitophast                                                |
    #> | github/mitoseek                                                         | github/mitoz                                                    |
    #> | github/mitre                                                            | github/mix                                                      |
    #> | github/mixclone                                                         | github/mixed-effect-gams                                        |
    #> | github/mixedperfectphylogeny                                            | github/mixmir                                                   |
    #> | github/mixmpln                                                          | github/mixsubhap                                                |
    #> | github/mlrmbo                                                           | github/mlsf                                                     |
    #> | github/mlstar                                                           | github/mltm                                                     |
    #> | github/mmcnn                                                            | github/mme-apis                                                 |
    #> | github/mmpdb                                                            | github/mmr                                                      |
    #> | github/mmseq                                                            | github/mmseqs                                                   |
    #> | github/mmtf-javascript                                                  | github/mmvec                                                    |
    #> | github/mnem                                                             | github/moabb                                                    |
    #> | github/mob-suite                                                        | github/mobiusklein                                              |
    #> | github/mocaprecovery                                                    | github/moccasin                                                 |
    #> | github/moccs                                                            | github/mod-bio                                                  |
    #> | github/mode-task                                                        | github/model-discovery-tool                                     |
    #> | github/model-fitting-cbs                                                | github/modelexplorer                                            |
    #> | github/modeltest                                                        | github/modentify                                                |
    #> | github/modict                                                           | github/modistools                                               |
    #> | github/modmaps3d                                                        | github/modsara2                                                 |
    #> | github/moff                                                             | github/molecular-clock                                          |
    #> | github/molgenis-imputation                                              | github/molgenis/molgenis                                        |
    #> | github/moli                                                             | github/moloc                                                    |
    #> | github/molti                                                            | github/molti-dream                                              |
    #> | github/molyso                                                           | github/momi-g                                                   |
    #> | github/monster                                                          | github/montra                                                   |
    #> | github/moomin                                                           | github/morfpred-plus                                            |
    #> | github/morph-bulk                                                       | github/morpholex-en                                             |
    #> | github/mosbat                                                           | github/mosdepth                                                 |
    #> | github/motif-scraper                                                    | github/motifbreakr                                              |
    #> | github/motion-detection                                                 | github/mountainclimber                                          |
    #> | github/mousetrap-os                                                     | github/mp                                                       |
    #> | github/mp-lamp                                                          | github/mpicbg-scicomp/snakemake-workflows                       |
    #> | github/mpies                                                            | github/mplnclust                                                |
    #> | github/mpradesigntools                                                  | github/mprascore                                                |
    #> | github/mptm                                                             | github/mptp                                                     |
    #> | github/mqap                                                             | github/mqc                                                      |
    #> | github/mr-demuxy                                                        | github/mrbait                                                   |
    #> | github/mrcnn-scene-recognition                                          | github/mrgbp                                                    |
    #> | github/mrgsea                                                           | github/mrlr                                                     |
    #> | github/mrrad                                                            | github/ms-data-core-api                                         |
    #> | github/ms-empire                                                        | github/ms-lda                                                   |
    #> | github/ms2gs                                                            | github/ms2ldaviz                                                |
    #> | github/msc-tree-resampling                                              | github/msf                                                      |
    #> | github/msisensor                                                        | github/msmexplorer                                              |
    #> | github/mspac                                                            | github/msparsm                                                  |
    #> | github/mspire                                                           | github/msproteomicstools                                        |
    #> | github/msqrob                                                           | github/msr                                                      |
    #> | github/msreduce                                                         | github/mst-minor-allele-caller                                  |
    #> | github/mswaldhmp                                                        | github/mtb-report                                               |
    #> | github/mtb-sig-miner                                                    | github/mtbgt                                                    |
    #> | github/mtbnn                                                            | github/mtbseq-source                                            |
    #> | github/mtgipick                                                         | github/mtgosc                                                   |
    #> | github/mtipatchsearch                                                   | github/mtmr-net                                                 |
    #> | github/mucor                                                            | github/mulrfrepo                                                |
    #> | github/multi-csar                                                       | github/multiclass-segmentation                                  |
    #> | github/multigems                                                        | github/multiguidescan                                           |
    #> | github/multimed                                                         | github/multimodal-brain-synthesis                               |
    #> | github/multimodal-medical-image-synthesis                               | github/multimodal-network-diffusion                             |
    #> | github/multimodalprognosis                                              | github/multiphate                                               |
    #> | github/multiplierz                                                      | github/multitable-review                                        |
    #> | github/multivarnetwork                                                  | github/multiwayclassification                                   |
    #> | github/multiwayregression                                               | github/musica                                                   |
    #> | github/musitedeep                                                       | github/mutationtcn                                              |
    #> | github/mutscan                                                          | github/muxstep                                                  |
    #> | github/mvbc                                                             | github/mvlr                                                     |
    #> | github/mvp-aaa-codelabs                                                 | github/mwas-biomarkers                                          |
    #> | github/mwfft2                                                           | github/mxfold                                                   |
    #> | github/mychembl                                                         | github/mycofier                                                 |
    #> | github/mycorrhiza                                                       | github/myelinj                                                  |
    #> | github/myelinq                                                          | github/mynclist                                                 |
    #> | github/mypmfs                                                           | github/mytai                                                    |
    #> | github/mz.unity                                                         | github/n1pas                                                    |
    #> | github/nacho                                                            | github/naf                                                      |
    #> | github/nanatex                                                          | github/nandb                                                    |
    #> | github/nanocall                                                         | github/nanodj                                                   |
    #> | github/nanomark                                                         | github/nanook                                                   |
    #> | github/nanopack                                                         | github/nanopype                                                 |
    #> | github/nanor                                                            | github/napeasy                                                  |
    #> | github/narx-qoe-release                                                 | github/nb-distribution                                          |
    #> | github/nbfpeg                                                           | github/nblda                                                    |
    #> | github/nbss                                                             | github/nbzimm                                                   |
    #> | github/ncappc                                                           | github/ncbi-hackathons                                          |
    #> | github/ncddetect2                                                       | github/ncgocr                                                   |
    #> | github/nclaes/coral                                                     | github/nclcomparator                                            |
    #> | github/ncphlda                                                          | github/ncyc                                                     |
    #> | github/ndd                                                              | github/ndexr                                                    |
    #> | github/neat                                                             | github/neat-genreads                                            |
    #> | github/neatfreq                                                         | github/neep                                                     |
    #> | github/negationdistribution                                             | github/negbio                                                   |
    #> | github/nemo                                                             | github/neo-pep-tool                                             |
    #> | github/neodti                                                           | github/neoepitope                                               |
    #> | github/neopredpipe                                                      | github/nephele                                                  |
    #> | github/neptune                                                          | github/nessie                                                   |
    #> | github/nest                                                             | github/nested-sampling                                          |
    #> | github/nestly                                                           | github/netgem                                                   |
    #> | github/nethypgeom                                                       | github/netics                                                   |
    #> | github/netminer                                                         | github/netml                                                    |
    #> | github/netpathminer                                                     | github/netprophet-2.0                                           |
    #> | github/netprot                                                          | github/netsmooth                                                |
    #> | github/network-classifier                                               | github/network-snps                                             |
    #> | github/networkannotation                                                | github/networkinference                                         |
    #> | github/networkregularisedcox                                            | github/networkrobustnesstoolbox                                 |
    #> | github/neubi                                                            | github/neural-circuits-vr                                       |
    #> | github/neuroinformatics                                                 | github/neuromatic                                               |
    #> | github/neuromeasure                                                     | github/neuromorphovis                                           |
    #> | github/neuroner                                                         | github/newton-method-mds                                        |
    #> | github/nextclip                                                         | github/nextflu                                                  |
    #> | github/nextpolish                                                       | github/nextstrain                                               |
    #> | github/nexus                                                            | github/ngl                                                      |
    #> | github/nglview                                                          | github/ngly                                                     |
    #> | github/ngmaster                                                         | github/ngmerge                                                  |
    #> | github/ngmlr                                                            | github/ngo2019-authortopic                                      |
    #> | github/ngrhmda                                                          | github/ngs-bits                                                 |
    #> | github/ngs-pipe                                                         | github/ngs-qcbox                                                |
    #> | github/ngsane                                                           | github/ngscheckmate                                             |
    #> | github/ngscloud                                                         | github/ngsf-hmm                                                 |
    #> | github/ngsld                                                            | github/ngspanpipe                                               |
    #> | github/ngsphy                                                           | github/ngsplusplus                                              |
    #> | github/ngsrelate                                                        | github/ngstools                                                 |
    #> | github/ngtas-pipeline                                                   | github/nicebot                                                  |
    #> | github/nichesimulation                                                  | github/nifpred                                                  |
    #> | github/niftylink                                                        | github/nimbus                                                   |
    #> | github/nimfa                                                            | github/nipter                                                   |
    #> | github/nitpicker                                                        | github/nitrilase                                                |
    #> | github/nitrogenuptake2016                                               | github/nitumid                                                  |
    #> | github/njmerge                                                          | github/nldmseq                                                  |
    #> | github/nlpar                                                            | github/nmfem                                                    |
    #> | github/nmflibrary                                                       | github/nmrlipids                                                |
    #> | github/nnlda                                                            | github/node.dating                                              |
    #> | github/noisecancellingrepeatfinder                                      | github/nojah                                                    |
    #> | github/non-parametric-algorithm-to-isolate-chunks-in-response-sequences | github/nonlineargradientsui                                     |
    #> | github/nonparametricsummarypsm                                          | github/nonpareil                                                |
    #> | github/novo-stitch                                                      | github/novocaller                                               |
    #> | github/novoplasty                                                       | github/npa                                                      |
    #> | github/npd-micro-saccade-detection                                      | github/nplb                                                     |
    #> | github/npnn                                                             | github/npreader                                                 |
    #> | github/npyc-toolbox                                                     | github/npyc-toolbox-tutorials                                   |
    #> | github/nquire                                                           | github/nrc                                                      |
    #> | github/nrl2drp                                                          | github/nrlmfb                                                   |
    #> | github/nrps-linker                                                      | github/nsclc-subtype-classification                             |
    #> | github/nseq                                                             | github/nsg                                                      |
    #> | github/nss-of-lwir-and-vissible-images                                  | github/nssrfbinary                                              |
    #> | github/nssrfpackage                                                     | github/ntcard                                                   |
    #> | github/ntedit                                                           | github/ntp-ersn                                                 |
    #> | github/nucdiff                                                          | github/nuclei-segmentation                                      |
    #> | github/nucleosee                                                        | github/nuclitrack                                               |
    #> | github/nurse                                                            | github/nvenn                                                    |
    #> | github/nvr                                                              | github/nvt                                                      |
    #> | github/nx4                                                              | github/nxtrim                                                   |
    #> | github/nyiuab/bhglm                                                     | github/obcs                                                     |
    #> | github/oclust                                                           | github/oddt                                                     |
    #> | github/ode-rmhmc                                                        | github/odin                                                     |
    #> | github/odonata-traits                                                   | github/oefinder                                                 |
    #> | github/off-target-prediction                                            | github/off-target-prediction                                    |
    #> | github/officechromatography                                             | github/offsetbasedgraph                                         |
    #> | github/offtargetpredict                                                 | github/og-consistency-pipeline                                  |
    #> | github/ogcleaner                                                        | github/ogfsc                                                    |
    #> | github/ogt-prediction                                                   | github/ohana                                                    |
    #> | github/ohvarfinder                                                      | github/olga                                                     |
    #> | github/ols-client                                                       | github/ols-dialog                                               |
    #> | github/omaat                                                            | github/omblast                                                  |
    #> | github/omgraph                                                          | github/omhismb2019                                              |
    #> | github/omic-features-successful-targets                                 | github/omicpred                                                 |
    #> | github/omicsintegrator                                                  | github/oncosimul                                                |
    #> | github/oncotator                                                        | github/online-chem                                              |
    #> | github/onto2graph                                                       | github/onto2vec                                                 |
    #> | github/ontologies                                                       | github/ontologyeval                                             |
    #> | github/ontologyevaluate                                                 | github/ontomaton                                                |
    #> | github/ontoquery                                                        | github/opa2vec                                                  |
    #> | github/opal                                                             | github/opal-plus                                                |
    #> | github/openbehavior                                                     | github/opencapsule                                              |
    #> | github/openfmrianalysis                                                 | github/openhealth                                               |
    #> | github/openms                                                           | github/opennft-demo                                             |
    #> | github/opentree                                                         | github/optical                                                  |
    #> | github/optimal-seed-solver                                              | github/optimir                                                  |
    #> | github/optimize-it                                                      | github/optimum-rlda                                             |
    #> | github/optnetaligncpp                                                   | github/optpipe                                                  |
    #> | github/orbkit                                                           | github/orbslamm                                                 |
    #> | github/orca                                                             | github/orcacode                                                 |
    #> | github/orchid                                                           | github/orderedpainting                                          |
    #> | github/ordino                                                           | github/orfm                                                     |
    #> | github/organelle-pba                                                    | github/organsegrstn                                             |
    #> | github/orna                                                             | github/orthofinder                                              |
    #> | github/orthopara-constraintchecker                                      | github/orthoscope                                               |
    #> | github/osdb                                                             | github/osfp                                                     |
    #> | github/osprey-refactor                                                  | github/ost-pymodules                                            |
    #> | github/outcome-prediction-nse                                           | github/outlier-in-blast-hits                                    |
    #> | github/ovarcall                                                         | github/overlap                                                  |
    #> | github/oxsa                                                             | github/oyster-river-protocol                                    |
    #> | github/p2rank                                                           | github/p5-bpwrapper                                             |
    #> | github/paamia2015                                                       | github/pacemaker-clustering-methods                             |
    #> | github/pachyderm/pachyderm                                              | github/pacificbiosciences/falcon                                |
    #> | github/package-gfe                                                      | github/pacl                                                     |
    #> | github/pacom                                                            | github/pact                                                     |
    #> | github/pad                                                              | github/padimi                                                   |
    #> | github/paga                                                             | github/pagal                                                    |
    #> | github/paladin                                                          | github/pals                                                     |
    #> | github/pam                                                              | github/pames                                                    |
    #> | github/pamir                                                            | github/pandelos                                                 |
    #> | github/panfp                                                            | github/pangenome                                                |
    #> | github/pangenomepipeline                                                | github/panget                                                   |
    #> | github/panisa                                                           | github/panoptes                                                 |
    #> | github/pantools                                                         | github/panviz                                                   |
    #> | github/papago                                                           | github/papst                                                    |
    #> | github/para-dpmm                                                        | github/para-suite                                               |
    #> | github/para-suite-aligner                                               | github/paragsea                                                 |
    #> | github/parapred                                                         | github/parasail                                                 |
    #> | github/parasor                                                          | github/parastagonospora-nodorum-sn15                            |
    #> | github/parentage                                                        | github/parfit                                                   |
    #> | github/pargenes                                                         | github/parkour                                                  |
    #> | github/parsnip                                                          | github/partie                                                   |
    #> | github/parties                                                          | github/partitionfinder                                          |
    #> | github/partitionuce                                                     | github/pascal                                                   |
    #> | github/pasdqc                                                           | github/pasnet                                                   |
    #> | github/pass                                                             | github/pastas                                                   |
    #> | github/pastaspark                                                       | github/patchdca                                                 |
    #> | github/path2surv                                                        | github/pathdeseq                                                |
    #> | github/pathemb                                                          | github/pathogenportal                                           |
    #> | github/pathos                                                           | github/pathrings                                                |
    #> | github/pathscore                                                        | github/pathway-mapper                                           |
    #> | github/pathwaymatcher                                                   | github/pathwaymatrix                                            |
    #> | github/patientexplorer                                                  | github/pats                                                     |
    #> | github/patterncnv                                                       | github/pattrec                                                  |
    #> | github/pavian                                                           | github/pavooc                                                   |
    #> | github/pbrowse                                                          | github/pbtest-r-package                                         |
    #> | github/pbxplore                                                         | github/pcasa                                                    |
    #> | github/pcdslab                                                          | github/pcetk                                                    |
    #> | github/pcgr                                                             | github/pclean-release                                           |
    #> | github/pcmf                                                             | github/pcnem                                                    |
    #> | github/pconsc4                                                          | github/pcr-error-correction                                     |
    #> | github/pcrduplicates                                                    | github/pctpred                                                  |
    #> | github/pdb-tools                                                        | github/pdb2entropy                                              |
    #> | github/pdb2trent                                                        | github/pdeep                                                    |
    #> | github/pdhs-elm                                                         | github/pdhs-svm                                                 |
    #> | github/pdna                                                             | github/pdv                                                      |
    #> | github/pdx-analysis-workflows                                           | github/pea                                                      |
    #> | github/pea-m5c                                                          | github/peak-grouping-alignment                                  |
    #> | github/peakc                                                            | github/peakerror                                                |
    #> | github/peakplotter                                                      | github/peakxus                                                  |
    #> | github/peakzilla                                                        | github/peath                                                    |
    #> | github/ped                                                              | github/pedigreenet                                              |
    #> | github/pedla                                                            | github/pegasus                                                  |
    #> | github/pehaplo                                                          | github/peitho                                                   |
    #> | github/penndiff                                                         | github/pep                                                      |
    #> | github/pepfoot                                                          | github/pepkalc                                                  |
    #> | github/pepr                                                             | github/pepvis                                                   |
    #> | github/pepx                                                             | github/perf                                                     |
    #> | github/perfect                                                          | github/perga                                                    |
    #> | github/perpetually-updated-trees                                        | github/personalized-regression                                  |
    #> | github/pets                                                             | github/pez                                                      |
    #> | github/pfam-map-loader                                                  | github/pfam-maps                                                |
    #> | github/pfinderuce-swsc-en                                               | github/pga                                                      |
    #> | github/pgen-genomicvariations-workflow                                  | github/pgrnafinder                                              |
    #> | github/phage                                                            | github/phandango                                                |
    #> | github/phanotate                                                        | github/pharmaconer                                              |
    #> | github/phat                                                             | github/phelim                                                   |
    #> | github/pheno4j                                                          | github/phenogocon                                               |
    #> | github/phenomenet-vp                                                    | github/phenomescape                                             |
    #> | github/phenominer                                                       | github/phenopolis                                               |
    #> | github/phenorank                                                        | github/phenos                                                   |
    #> | github/phenotype-prediction-pipeline                                    | github/phenotypeseeker                                          |
    #> | github/phenotypesimulator                                               | github/phesant                                                  |
    #> | github/phewas                                                           | github/phnmnl                                                   |
    #> | github/phocos                                                           | github/phoglystruct                                             |
    #> | github/phosphoproteome-prediction                                       | github/phraider                                                 |
    #> | github/phy-mer                                                          | github/phyc                                                     |
    #> | github/phyclip                                                          | github/phyd3                                                    |
    #> | github/phydelity                                                        | github/phyfoo                                                   |
    #> | github/phyinformr                                                       | github/phylesystem                                              |
    #> | github/phylesystem-api                                                  | github/phylip-enhance                                           |
    #> | github/phylo-io                                                         | github/phylo-node                                               |
    #> | github/phyloctcflooping                                                 | github/phylogeotool                                             |
    #> | github/phyloligo                                                        | github/phylomad                                                 |
    #> | github/phylomagnet                                                      | github/phylonetworks.jl                                         |
    #> | github/phylopi                                                          | github/phyloprofile                                             |
    #> | github/phylorad                                                         | github/phyloscanner                                             |
    #> | github/phylostratr                                                      | github/phylosub                                                 |
    #> | github/phylowgs                                                         | github/phyml                                                    |
    #> | github/physca                                                           | github/physiboss                                                |
    #> | github/phytoassembly                                                    | github/phyx                                                     |
    #> | github/pi2                                                              | github/pia                                                      |
    #> | github/piblup                                                           | github/picalculax                                               |
    #> | github/picky                                                            | github/pidgin                                                   |
    #> | github/piil                                                             | github/pileup.js                                                |
    #> | github/pimms                                                            | github/pimp                                                     |
    #> | github/pinda                                                            | github/pinetree                                                 |
    #> | github/pingpongpro                                                      | github/pipeline-for-isoseq                                      |
    #> | github/pipeliner                                                        | github/pipred                                                   |
    #> | github/pirnapredictor                                                   | github/pirnn                                                    |
    #> | github/pirsupplementary                                                 | github/pisces                                                   |
    #> | github/pixel                                                            | github/pixie                                                    |
    #> | github/plaac                                                            | github/plaidoh                                                  |
    #> | github/planteome                                                        | github/plasmidseeker                                            |
    #> | github/plasmidtron                                                      | github/pldist                                                   |
    #> | github/plier                                                            | github/pll-modules                                              |
    #> | github/ploidyinfer                                                      | github/ploidyngs                                                |
    #> | github/pluto                                                            | github/pmanalyzer                                               |
    #> | github/pmanalyzerweb                                                    | github/pmd                                                      |
    #> | github/pmfa                                                             | github/pmsignature                                              |
    #> | github/pnc                                                              | github/pnerf                                                    |
    #> | github/pnnl                                                             | github/poap                                                     |
    #> | github/pocketpipe                                                       | github/polster                                                  |
    #> | github/polya-prediction-lrm-dnn                                         | github/polyainv                                                 |
    #> | github/polyclustr                                                       | github/polydfe                                                  |
    #> | github/polyploid-genotyping                                             | github/polyqtl                                                  |
    #> | github/pong                                                             | github/popcluster                                               |
    #> | github/popdemog                                                         | github/popfba                                                   |
    #> | github/pophelper                                                        | github/popins                                                   |
    #> | github/poplddecay                                                       | github/popstr                                                   |
    #> | github/poptrt                                                           | github/posdatools                                               |
    #> | github/pose                                                             | github/pose-regularization                                      |
    #> | github/posigene                                                         | github/positionalmi                                             |
    #> | github/positive-selection                                               | github/postplsr                                                 |
    #> | github/pourrna                                                          | github/powerfit                                                 |
    #> | github/powsimr                                                          | github/ppa-assembler                                            |
    #> | github/ppi                                                              | github/ppis-wdsvm                                               |
    #> | github/ppnid                                                            | github/ppr-meta                                                 |
    #> | github/ppsp                                                             | github/pqhe                                                     |
    #> | github/praline                                                          | github/prankwebapp                                              |
    #> | github/precise                                                          | github/predcrp                                                  |
    #> | github/predgly                                                          | github/predictability-of-cancer-evolution                       |
    #> | github/prediction-prostate-surveillance                                 | github/predixcan                                                |
    #> | github/predpsi-svr                                                      | github/predrad                                                  |
    #> | github/prefersim                                                        | github/prefmd                                                   |
    #> | github/prequal                                                          | github/presm                                                    |
    #> | github/presogenesis                                                     | github/price                                                    |
    #> | github/pride-toolsuite                                                  | github/primer3                                                  |
    #> | github/primer3-masker                                                   | github/princces                                                 |
    #> | github/prince                                                           | github/princess-opensource                                      |
    #> | github/priorknowledgeepistasisrank                                      | github/prism                                                    |
    #> | github/probabilitymapviewer                                             | github/probannopy                                               |
    #> | github/probhap                                                          | github/procell                                                  |
    #> | github/prodigy                                                          | github/profet                                                   |
    #> | github/progressivecactus                                                | github/promodule                                                |
    #> | github/promotepredictor                                                 | github/promoterpredict                                          |
    #> | github/proper-extension                                                 | github/propip                                                   |
    #> | github/proq-scripts                                                     | github/prosampler                                               |
    #> | github/prosave                                                          | github/prosstt                                                  |
    #> | github/protaeljs                                                        | github/protasr                                                  |
    #> | github/protein-folding-decoy-set                                        | github/protein-ml                                               |
    #> | github/protein-recon-mcriemman                                          | github/protein-sequence-analysis                                |
    #> | github/protein-solubility                                               | github/protein2genetree                                         |
    #> | github/proteinpocketdetection                                           | github/proteoformer                                             |
    #> | github/proteomicsbrowser                                                | github/proteomodlr                                              |
    #> | github/proteosign                                                       | github/protspam                                                 |
    #> | github/prottrace                                                        | github/proxdist                                                 |
    #> | github/proxl-web-app                                                    | github/psbvb                                                    |
    #> | github/pseaac-psepssm-wd                                                | github/psearch                                                  |
    #> | github/psepssm-dcca-lfda                                                | github/psi                                                      |
    #> | github/psi-blastexb                                                     | github/psi-sigma                                                |
    #> | github/psims                                                            | github/psite                                                    |
    #> | github/psiz                                                             | github/psmvc                                                    |
    #> | github/psucce                                                           | github/psynteract                                               |
    #> | github/ptempest                                                         | github/ptm-x                                                    |
    #> | github/ptmapper                                                         | github/ptolemy                                                  |
    #> | github/ptuneos                                                          | github/ptw                                                      |
    #> | github/ptxqc                                                            | github/pubmedportable                                           |
    #> | github/pufferfish                                                       | github/puffin                                                   |
    #> | github/pulmon                                                           | github/pulp                                                     |
    #> | github/pulser                                                           | github/pupil-size                                               |
    #> | github/pureclip                                                         | github/pvac-seq                                                 |
    #> | github/pvalue-weighting-matlab                                          | github/pviz                                                     |
    #> | github/pwc-net                                                          | github/pwrewas                                                  |
    #> | github/pwtees                                                           | github/pyabc                                                    |
    #> | github/pyar                                                             | github/pyaudioanalysis                                          |
    #> | github/pyboolnet                                                        | github/pycellerator                                             |
    #> | github/pycgm                                                            | github/pycgtool                                                 |
    #> | github/pychemia                                                         | github/pychimera                                                |
    #> | github/pycotools                                                        | github/pydca                                                    |
    #> | github/pydream                                                          | github/pyefp                                                    |
    #> | github/pyfeat                                                           | github/pygenprop                                                |
    #> | github/pygtftk                                                          | github/pyham                                                    |
    #> | github/pyhla                                                            | github/pyhsiclasso                                              |
    #> | github/pykeen                                                           | github/pyloh                                                    |
    #> | github/pymbar                                                           | github/pymcpsc                                                  |
    #> | github/pymethylprocess                                                  | github/pymod                                                    |
    #> | github/pymzml                                                           | github/pynbs                                                    |
    #> | github/pyncs                                                            | github/pypanda                                                  |
    #> | github/pypdb                                                            | github/pypedia                                                  |
    #> | github/pypedia-server                                                   | github/pyphi                                                    |
    #> | github/pypore                                                           | github/pyro-nn                                                  |
    #> | github/pyrod                                                            | github/pyscestoolbox                                            |
    #> | github/pyseer                                                           | github/pysfd                                                    |
    #> | github/pysm                                                             | github/pysradb                                                  |
    #> | github/pysster                                                          | github/pytag                                                    |
    #> | github/pytanfinder                                                      | github/pytc                                                     |
    #> | github/pyvolve                                                          | github/pywater                                                  |
    #> | github/pywindow                                                         | github/q2-feature-classifier                                    |
    #> | github/q2-fragment-insertion                                            | github/q2-longitudinal                                          |
    #> | github/q2mm                                                             | github/qanalysis-project                                        |
    #> | github/qbioimages                                                       | github/qc-dmet                                                  |
    #> | github/qc3                                                              | github/qiaseq-dna                                               |
    #> | github/qmlgalaxyportal                                                  | github/qpa                                                      |
    #> | github/qparse                                                           | github/qpath                                                    |
    #> | github/qreadselector                                                    | github/qs-net                                                   |
    #> | github/qsimscan                                                         | github/qsmooth                                                  |
    #> | github/qsubsec                                                          | github/qtc                                                      |
    #> | github/quadmutex                                                        | github/quagmir                                                  |
    #> | github/quandico                                                         | github/quantifly                                                |
    #> | github/quantilebootstrap                                                | github/quantinemo                                               |
    #> | github/quantnorm                                                        | github/quartetscores                                            |
    #> | github/quasar                                                           | github/qubic2                                                   |
    #> | github/quentin                                                          | github/queries                                                  |
    #> | github/querymed                                                         | github/questionnaire                                            |
    #> | github/quimp                                                            | github/quin                                                     |
    #> | github/quota-alignment                                                  | github/qvz                                                      |
    #> | github/r-samstrt                                                        | github/r.sambada                                                |
    #> | github/r2dgc                                                            | github/r4kappa                                                  |
    #> | github/raamlab                                                          | github/rabifier                                                 |
    #> | github/racat                                                            | github/racipe-1.0                                               |
    #> | github/ractip                                                           | github/radaa                                                    |
    #> | github/radscripts                                                       | github/ragoo                                                    |
    #> | github/ragout                                                           | github/ragp                                                     |
    #> | github/rail-dbgap                                                       | github/raincloudplots                                           |
    #> | github/rambl                                                            | github/ramf                                                     |
    #> | github/ramp-backend                                                     | github/rampart                                                  |
    #> | github/ramwas                                                           | github/randal                                                   |
    #> | github/random-forest-protein-ligand-decoy-detection                     | github/ranger                                                   |
    #> | github/rapid                                                            | github/rapidaci                                                 |
    #> | github/rapidmic                                                         | github/rapmap                                                   |
    #> | github/rappas                                                           | github/raptr-sv                                                 |
    #> | github/rarefaction                                                      | github/raser                                                    |
    #> | github/rasp                                                             | github/raspberry                                                |
    #> | github/ratschlab/pbwt-sec                                               | github/raunaq-m/multires                                        |
    #> | github/raven                                                            | github/ravensoftware                                            |
    #> | github/raxml-light-1.0.5                                                | github/raxml-ng                                                 |
    #> | github/rbapy                                                            | github/rbbt-workflows                                           |
    #> | github/rbiomirgs                                                        | github/rbv                                                      |
    #> | github/rcorrector                                                       | github/rcount                                                   |
    #> | github/rcpa.tools                                                       | github/rcppocc                                                  |
    #> | github/rcsm                                                             | github/rcupcake                                                 |
    #> | github/rcupcake-case-studies                                            | github/rd-analyzer                                              |
    #> | github/rdb2c                                                            | github/rdolphin                                                 |
    #> | github/rea                                                              | github/reactidr                                                 |
    #> | github/reactiondecoder                                                  | github/reactionflow                                             |
    #> | github/reactionminer                                                    | github/reactome-pengine                                         |
    #> | github/read-split-run                                                   | github/readtools                                                |
    #> | github/reago                                                            | github/recap                                                    |
    #> | github/recov                                                            | github/recovery                                                 |
    #> | github/recursive-chemprot                                               | github/recycler                                                 |
    #> | github/red-lesion-detection                                             | github/red-ml                                                   |
    #> | github/redcap-ddp                                                       | github/redetector/red                                           |
    #> | github/redundans                                                        | github/refbool                                                  |
    #> | github/refeditor                                                        | github/regate                                                   |
    #> | github/regbase                                                          | github/regbenchmark                                             |
    #> | github/regcombims                                                       | github/regent                                                   |
    #> | github/region-templates                                                 | github/reglinker                                                |
    #> | github/regmex                                                           | github/regnet                                                   |
    #> | github/regresshaplo                                                     | github/regshape                                                 |
    #> | github/regularized-manova                                               | github/regulatorynetworkgmpmodel                                |
    #> | github/rehunt                                                           | github/relz-fs                                                  |
    #> | github/remap                                                            | github/remilo                                                   |
    #> | github/remma                                                            | github/removepli                                                |
    #> | github/rendertoolbox3                                                   | github/rentplus                                                 |
    #> | github/reoa                                                             | github/reparation                                               |
    #> | github/repdenovo                                                        | github/repeatcraftp                                             |
    #> | github/repeatseq                                                        | github/repeatsoaker                                             |
    #> | github/replong                                                          | github/repository                                               |
    #> | github/reptile                                                          | github/repurpose                                                |
    #> | github/repviz                                                           | github/reqtl                                                    |
    #> | github/rerconverge                                                      | github/rescore                                                  |
    #> | github/rescore-metabolites                                              | github/rescue                                                   |
    #> | github/resimnet                                                         | github/response-logic                                           |
    #> | github/response-logic-projects                                          | github/respre                                                   |
    #> | github/resseq                                                           | github/retroseq                                                 |
    #> | github/retrotransposons-spread                                          | github/reverse-engineering-bc-grns                              |
    #> | github/rfathm6a                                                         | github/rfscorevs                                                |
    #> | github/rfscorevs-binary                                                 | github/rgaat-v2                                                 |
    #> | github/rgeneticthesaurus                                                | github/rgv                                                      |
    #> | github/rhat                                                             | github/rhe-reg                                                  |
    #> | github/riblast                                                          | github/ribodiff                                                 |
    #> | github/riborex                                                          | github/ribotricer                                               |
    #> | github/ribotricer-results                                               | github/riboviz                                                  |
    #> | github/ribowaltz                                                        | github/rifraf.jl                                                |
    #> | github/riginv                                                           | github/rinbix                                                   |
    #> | github/ringo                                                            | github/risa                                                     |
    #> | github/ritan                                                            | github/ritornello                                               |
    #> | github/rl-gensvm                                                        | github/rl-skat                                                  |
    #> | github/rlatools                                                         | github/rlowpc                                                   |
    #> | github/rlscore                                                          | github/rm-seq                                                   |
    #> | github/rmats-dvr                                                        | github/rmetl                                                    |
    #> | github/rmfam                                                            | github/rmfilter                                                 |
    #> | github/rmne                                                             | github/rmscca                                                   |
    #> | github/rmsi                                                             | github/rmtl                                                     |
    #> | github/rmut                                                             | github/rna-interactions-benchmark                               |
    #> | github/rna-secondary-structure-database                                 | github/rna3dcnn                                                 |
    #> | github/rnadetect                                                        | github/rnaeditingindexer                                        |
    #> | github/rnaentropy                                                       | github/rnaifold                                                 |
    #> | github/rnaindel                                                         | github/rnaplonc                                                 |
    #> | github/rnapuzzler                                                       | github/rnaredprint                                              |
    #> | github/rnaseq-protocol                                                  | github/rnaseqeval                                               |
    #> | github/rnaseqview                                                       | github/rnie                                                     |
    #> | github/rnmrfit                                                          | github/rnn-for-membrane-protein-types-prediction                |
    #> | github/robokop                                                          | github/robustclusteringpatientsubtyping                         |
    #> | github/robustsparsecorrelation                                          | github/rocketship                                               |
    #> | github/rodin                                                            | github/roicnn                                                   |
    #> | github/roma                                                             | github/romop                                                    |
    #> | github/romulus                                                          | github/rop                                                      |
    #> | github/ropebwt2                                                         | github/rosids                                                   |
    #> | github/roulattice                                                       | github/rpowerlib                                                |
    #> | github/rqt                                                              | github/rrbssim                                                  |
    #> | github/rse                                                              | github/rsfp                                                     |
    #> | github/rsmlm                                                            | github/rsmlm-tutorials                                          |
    #> | github/rsnapsim                                                         | github/rss                                                      |
    #> | github/rsss                                                             | github/rtfbs-db                                                 |
    #> | github/rtm-gwas                                                         | github/runbng                                                   |
    #> | github/rupee                                                            | github/rvolcano                                                 |
    #> | github/rvs                                                              | github/rvtests                                                  |
    #> | github/rwen                                                             | github/rwr-mh                                                   |
    #> | github/s-predixcan                                                      | github/s2b                                                      |
    #> | github/s3m                                                              | github/s4vdpca                                                  |
    #> | github/sa-ssr                                                           | github/saarclust                                                |
    #> | github/saber                                                            | github/safeclustering                                           |
    #> | github/sage2                                                            | github/sail                                                     |
    #> | github/sailfish-cir                                                     | github/saint2                                                   |
    #> | github/salivaprint                                                      | github/salmon                                                   |
    #> | github/salsa                                                            | github/samblaster                                               |
    #> | github/samdude                                                          | github/samfire                                                  |
    #> | github/sammi                                                            | github/samovar                                                  |
    #> | github/sampl5-dc-surface-empirical                                      | github/sampled-ancestors                                        |
    #> | github/samtools                                                         | github/sandres                                                  |
    #> | github/sanefalcon                                                       | github/sap                                                      |
    #> | github/sap01-tgs-lite-supplem                                           | github/sarks                                                    |
    #> | github/sarw-lnkf                                                        | github/sasm-vgwas                                               |
    #> | github/saspy                                                            | github/sativa                                                   |
    #> | github/saturn                                                           | github/sbb                                                      |
    #> | github/sbc                                                              | github/sbd-evaluation                                           |
    #> | github/sbetoolbox                                                       | github/sbgntikz                                                 |
    #> | github/sbl                                                              | github/sbml-diff                                                |
    #> | github/sbml-mod-ws                                                      | github/sbol-validator                                           |
    #> | github/sbvb0                                                            | github/sc2p                                                     |
    #> | github/scae-ir-spectral-imaging                                         | github/scanfold                                                 |
    #> | github/scanindel                                                        | github/scanneo                                                  |
    #> | github/scanpav                                                          | github/scanpy                                                   |
    #> | github/scanvis                                                          | github/scbgfengchao                                             |
    #> | github/scbs-map                                                         | github/sccaller                                                 |
    #> | github/sccg                                                             | github/scct                                                     |
    #> | github/scdesign                                                         | github/scdnet                                                   |
    #> | github/scellsupplementary                                               | github/scent                                                    |
    #> | github/scepath                                                          | github/scfba                                                    |
    #> | github/scgeatoolbox                                                     | github/scgen                                                    |
    #> | github/schinter                                                         | github/scifil                                                   |
    #> | github/scikit-ribo                                                      | github/scipion                                                  |
    #> | github/scistree                                                         | github/scmarker                                                 |
    #> | github/scmatch                                                          | github/scmut                                                    |
    #> | github/scoary                                                           | github/scode                                                    |
    #> | github/scop                                                             | github/scot                                                     |
    #> | github/scoup                                                            | github/scout                                                    |
    #> | github/scphaser                                                         | github/scrat                                                    |
    #> | github/screenbeam                                                       | github/screenmasker                                             |
    #> | github/script-hmm-replication-timing                                    | github/scrnabatchqc                                             |
    #> | github/scrnaseq-benchmark                                               | github/scrnaseq-cell-cluster-labeling                           |
    #> | github/scrnaseq-clustering-comparison                                   | github/scrublet                                                 |
    #> | github/scssim                                                           | github/sct                                                      |
    #> | github/sctcrseq                                                         | github/sctree-test                                              |
    #> | github/scuba                                                            | github/scvdmc                                                   |
    #> | github/scvi                                                             | github/sda                                                      |
    #> | github/sdeap                                                            | github/sdg                                                      |
    #> | github/seaflow-cluster                                                  | github/sealer-release                                           |
    #> | github/sechidis                                                         | github/secretsanta                                              |
    #> | github/section-sort                                                     | github/securema                                                 |
    #> | github/secureqc                                                         | github/sedef                                                    |
    #> | github/seed                                                             | github/seeksv                                                   |
    #> | github/seevis                                                           | github/segmentation-software                                    |
    #> | github/segmitos-mitosis-detection                                       | github/selam                                                    |
    #> | github/self-blm                                                         | github/self-organization                                        |
    #> | github/selfish                                                          | github/selscan                                                  |
    #> | github/sem-cpp                                                          | github/sema                                                     |
    #> | github/semanticprogramming                                              | github/semanticsco                                              |
    #> | github/semehr                                                           | github/semgen                                                   |
    #> | github/senet                                                            | github/sepp                                                     |
    #> | github/seq-ppi                                                          | github/seqan                                                    |
    #> | github/seqbuster                                                        | github/seqenv                                                   |
    #> | github/seqkit                                                           | github/seqlib                                                   |
    #> | github/seqlm                                                            | github/seqotron                                                 |
    #> | github/seqpatch                                                         | github/seqsero2                                                 |
    #> | github/seqsleepnet                                                      | github/sequana                                                  |
    #> | github/sequence-prediction-using-cnn-and-lstms                          | github/sequence-search-tool-for-antimicrobial-resistance-sstar- |
    #> | github/sequence2vec                                                     | github/sereega                                                  |
    #> | github/seroba                                                           | github/serpent                                                  |
    #> | github/set-cover-tools                                                  | github/sexcmd                                                   |
    #> | github/sexy                                                             | github/sf-abc                                                   |
    #> | github/sgspls                                                           | github/sgtk                                                     |
    #> | github/shaker                                                           | github/shamsaraj                                                |
    #> | github/shaoweinuaa                                                      | github/shape-component-analysis-matlab                          |
    #> | github/sharpvisu                                                        | github/shawn-xu/ppip                                            |
    #> | github/sheikhizadeh/ace                                                 | github/shesisplus                                               |
    #> | github/shiny-revecor                                                    | github/shiny-seq                                                |
    #> | github/shinycircos                                                      | github/shinycnv                                                 |
    #> | github/shinyheatmap                                                     | github/shinymethyl                                              |
    #> | github/shiver                                                           | github/shouji                                                   |
    #> | github/shubhamchandak94/spring                                          | github/sia                                                      |
    #> | github/sian                                                             | github/siapopr                                                  |
    #> | github/sibjoin.git+                                                     | github/sid                                                      |
    #> | github/side-effects                                                     | github/sieve                                                    |
    #> | github/sieve-score                                                      | github/sieve-ub                                                 |
    #> | github/sifrproject                                                      | github/sifter-t                                                 |
    #> | github/sigma                                                            | github/sigmat                                                   |
    #> | github/sigmod                                                           | github/signature-sj                                             |
    #> | github/sigprofilermatrixgenerator                                       | github/silentmutations                                          |
    #> | github/silkslider                                                       | github/sim-tmle-tutorial                                        |
    #> | github/simbac                                                           | github/simclda                                                  |
    #> | github/simd                                                             | github/simextargid                                              |
    #> | github/simgwas                                                          | github/simka                                                    |
    #> | github/simkern                                                          | github/simlrgithub                                              |
    #> | github/simpel                                                           | github/simphony                                                 |
    #> | github/simphy                                                           | github/simpleitk-notebooks                                      |
    #> | github/simra                                                            | github/simsem                                                   |
    #> | github/simug                                                            | github/simulatingcochlearimplants                               |
    #> | github/simurg                                                           | github/sinbad                                                   |
    #> | github/sincerities                                                      | github/sine-scan                                                |
    #> | github/single-cell-review                                               | github/singlecellassay                                          |
    #> | github/sinnlrr                                                          | github/sioor                                                    |
    #> | github/sisrs                                                            | github/sitar                                                    |
    #> | github/sixparam                                                         | github/sjaracne                                                 |
    #> | github/skandlab/smurf                                                   | github/skesa                                                    |
    #> | github/skpcr                                                            | github/skygrowth                                                |
    #> | github/slamem                                                           | github/slantbrainseg                                            |
    #> | github/sle                                                              | github/sleepeegnet                                              |
    #> | github/slicemap                                                         | github/slide                                                    |
    #> | github/slimenrich                                                       | github/slimer                                                   |
    #> | github/slimmer                                                          | github/slin                                                     |
    #> | github/slr                                                              | github/sm-engine                                                |
    #> | github/smallgenometools                                                 | github/smartr                                                   |
    #> | github/smash                                                            | github/smcgenedeconv                                            |
    #> | github/smilesvecproteinrepresentation                                   | github/smirarab/pasta                                           |
    #> | github/smith2012-insulin-signalling                                     | github/smlocalizer                                              |
    #> | github/smrt                                                             | github/smrt-pipeline                                            |
    #> | github/sms                                                              | github/smudge                                                   |
    #> | github/smurf                                                            | github/snakepipes                                               |
    #> | github/snapperdb                                                        | github/snapperdb-references                                     |
    #> | github/snappy                                                           | github/snaptron                                                 |
    #> | github/snaptron-experiments                                             | github/snftool                                                  |
    #> | github/sniffles                                                         | github/sniper                                                   |
    #> | github/snipviz                                                          | github/snovault                                                 |
    #> | github/snowball                                                         | github/snp-select                                               |
    #> | github/snp-sites                                                        | github/snp2sim                                                  |
    #> | github/snpclust                                                         | github/snpenrichr                                               |
    #> | github/snpfilt                                                          | github/snpgenie                                                 |
    #> | github/snpr                                                             | github/snpsvm                                                   |
    #> | github/snvphyl-galaxy                                                   | github/snyder-birth-death-code                                  |
    #> | github/sodapop                                                          | github/sofia                                                    |
    #> | github/solarius                                                         | github/solidbin                                                 |
    #> | github/solidmon                                                         | github/solrplantapi                                             |
    #> | github/solver-for-the-jeffreys-type-equations-system                    | github/som                                                      |
    #> | github/sonar                                                            | github/sonics                                                   |
    #> | github/sopang                                                           | github/sora                                                     |
    #> | github/sorting-cancer-karyotypes                                        | github/sourmash                                                 |
    #> | github/sowhat                                                           | github/spaced-seeds-for-metagenomics                            |
    #> | github/spacescanner                                                     | github/spades                                                   |
    #> | github/spadevizr                                                        | github/spagi                                                    |
    #> | github/spar-k                                                           | github/sparc                                                    |
    #> | github/spark-cpvs                                                       | github/spark-vs                                                 |
    #> | github/sparkbwa                                                         | github/sparse-lowrank-regression                                |
    #> | github/sparseiso                                                        | github/sparsemcmm                                               |
    #> | github/spcapp                                                           | github/spcovr                                                   |
    #> | github/spdrank                                                          | github/speaq                                                    |
    #> | github/spectacle                                                        | github/specter                                                  |
    #> | github/spectraph                                                        | github/spectrassembler                                          |
    #> | github/spectre                                                          | github/spectreok                                                |
    #> | github/spectrogene                                                      | github/spectrum                                                 |
    #> | github/spedeimportance                                                  | github/spherical                                                |
    #> | github/sphyr                                                            | github/spiddor                                                  |
    #> | github/spike-compression-autoencoder                                    | github/spin-scenario                                            |
    #> | github/spind                                                            | github/spladder                                                 |
    #> | github/splicefamalign                                                   | github/splicejumper                                             |
    #> | github/splicelauncher                                                   | github/spliceogen                                               |
    #> | github/splicesites-ann-cgr                                              | github/splicev                                                  |
    #> | github/splicify                                                         | github/splicing-analysis-pipeline                               |
    #> | github/spm                                                              | github/spm-code                                                 |
    #> | github/sports1.0                                                        | github/spot-a-neonatalrabbit                                    |
    #> | github/spotyping-v2.0                                                   | github/sppider                                                  |
    #> | github/spread                                                           | github/springsalad                                              |
    #> | github/sprint                                                           | github/sprint-race                                              |
    #> | github/sprites                                                          | github/sprites2                                                 |
    #> | github/spurf                                                            | github/sputnik                                                  |
    #> | github/spv-signaling-pathway-visualizer-v1.0                            | github/squadbookchapter                                         |
    #> | github/squat                                                            | github/squeakr                                                  |
    #> | github/squeezemeta                                                      | github/squiggle                                                 |
    #> | github/squigglekit                                                      | github/squire                                                   |
    #> | github/srbreak                                                          | github/srf                                                      |
    #> | github/srhtests                                                         | github/ssages-public                                            |
    #> | github/ssasu                                                            | github/ssbio                                                    |
    #> | github/sscclust                                                         | github/ssdfa                                                    |
    #> | github/ssgsea2.0                                                        | github/sskat                                                    |
    #> | github/ssp                                                              | github/ssp-lsdr                                                 |
    #> | github/ssrsc                                                            | github/sss-test                                                 |
    #> | github/sst                                                              | github/ssusearch                                                |
    #> | github/st-pipeline                                                      | github/st-viewer                                                |
    #> | github/stablede                                                         | github/star-genome-browser                                      |
    #> | github/starchip                                                         | github/starcode                                                 |
    #> | github/startapp                                                         | github/stasnet                                                  |
    #> | github/statistical-agglomeration                                        | github/steadystateconc                                          |
    #> | github/stego                                                            | github/stells2                                                  |
    #> | github/stemcellsva                                                      | github/stereoseg                                                |
    #> | github/stillbirth                                                       | github/stk                                                      |
    #> | github/stochastic-diffusion                                             | github/stochhmm                                                 |
    #> | github/stopgap                                                          | github/stp-gui                                                  |
    #> | github/stra-net                                                         | github/straingems                                               |
    #> | github/strainhub                                                        | github/strainseeker                                             |
    #> | github/strandscript                                                     | github/strassembly                                              |
    #> | github/strawberry                                                       | github/streamingtrim                                            |
    #> | github/strelka                                                          | github/streptomyces-coelicolor-gem                              |
    #> | github/stretch                                                          | github/stride                                                   |
    #> | github/stripes                                                          | github/structure-threader                                       |
    #> | github/structurefold2                                                   | github/struo                                                    |
    #> | github/studyprotocolsandbox                                             | github/stytra                                                   |
    #> | github/subcellular-localization                                         | github/subcloneseeker                                           |
    #> | github/subclpred                                                        | github/subdyquency                                              |
    #> | github/submatrixselectionsvd.jl                                         | github/submito-xgboost                                          |
    #> | github/submodulartrack                                                  | github/subnetwork-inference                                     |
    #> | github/subrecon                                                         | github/subseq                                                   |
    #> | github/substra                                                          | github/suget                                                    |
    #> | github/suicidalaltruismdissertation                                     | github/sumac                                                    |
    #> | github/sumer                                                            | github/summaryauc                                               |
    #> | github/sumsec                                                           | github/sunbeam                                                  |
    #> | github/suns                                                             | github/suns-cmd                                                 |
    #> | github/suns-search                                                      | github/supcp                                                    |
    #> | github/super-i2b2                                                       | github/super-resolution-adversarial-defense                     |
    #> | github/superangle                                                       | github/superbubble                                              |
    #> | github/superencoder                                                     | github/superproteintree                                         |
    #> | github/superquantnode                                                   | github/suretypesc                                               |
    #> | github/surfelmeshing                                                    | github/surpriseme                                               |
    #> | github/surrogateminimaldepth                                            | github/survindel                                                |
    #> | github/survival                                                         | github/survivalcausaltree                                       |
    #> | github/sushi                                                            | github/sv-bay                                                   |
    #> | github/sv-plauditdemonstration                                          | github/sv2                                                      |
    #> | github/svaseq                                                           | github/svdquest                                                 |
    #> | github/sve                                                              | github/svict                                                    |
    #> | github/svim                                                             | github/svmine                                                   |
    #> | github/svpv                                                             | github/svtools                                                  |
    #> | github/sw-tandem                                                        | github/swapcounter                                              |
    #> | github/swarmsight                                                       | github/swift-t-variant-calling                                  |
    #> | github/swiftlink                                                        | github/swiftortho                                               |
    #> | github/swinger                                                          | github/switchablenorms                                          |
    #> | github/swne                                                             | github/sword                                                    |
    #> | github/swpepnovo                                                        | github/symmetry                                                 |
    #> | github/synergy-screen                                                   | github/synima                                                   |
    #> | github/synquant                                                         | github/synrio                                                   |
    #> | github/syntenybrowser                                                   | github/synthesizer                                              |
    #> | github/sysbiolux/falcon                                                 | github/szczepankiewicz-dib-2019                                 |
    #> | github/t-cell-classification                                            | github/t-lex3                                                   |
    #> | github/t-sne-heatmaps                                                   | github/t2ga                                                     |
    #> | github/taba                                                             | github/tabsat                                                   |
    #> | github/tad-fusion-score                                                 | github/tada                                                     |
    #> | github/tadbit                                                           | github/tadoss                                                   |
    #> | github/tadtool                                                          | github/tagfinder                                                |
    #> | github/tagger                                                           | github/taiji                                                    |
    #> | github/tailor                                                           | github/tangle                                                   |
    #> | github/tango-final                                                      | github/tangram                                                  |
    #> | github/tankche1                                                         | github/tapas                                                    |
    #> | github/tapes                                                            | github/tappm-cli                                                |
    #> | github/tar-vir                                                          | github/tardis                                                   |
    #> | github/targetclone                                                      | github/targetmine-gradle                                        |
    #> | github/targetseqview                                                    | github/targetspecificgatksequencingpipeline                     |
    #> | github/taska                                                            | github/tawny-karyotype                                          |
    #> | github/tax-credit-data                                                  | github/tbe                                                      |
    #> | github/tc-recon                                                         | github/tcc-gui                                                  |
    #> | github/tcga-assembler-2                                                 | github/tcga-rnaseq-clinical                                     |
    #> | github/tcgenerators                                                     | github/tciapathfindersignificance                               |
    #> | github/tcrclusteringpaper                                               | github/tcsbu                                                    |
    #> | github/tcsm                                                             | github/tea                                                      |
    #> | github/teacheng                                                         | github/teachopencadd                                            |
    #> | github/team                                                             | github/tecandidates                                             |
    #> | github/tefma                                                            | github/tegs                                                     |
    #> | github/telescope                                                        | github/temp                                                     |
    #> | github/temt                                                             | github/tensqr                                                   |
    #> | github/tepic                                                            | github/terapca                                                  |
    #> | github/terraphast                                                       | github/tersect                                                  |
    #> | github/testsetbias                                                      | github/tetools                                                  |
    #> | github/tetyper                                                          | github/texomer                                                  |
    #> | github/text-image-deblur                                                | github/tf-chan-lab/omtools                                      |
    #> | github/tf2exp                                                           | github/tfanalysis                                               |
    #> | github/tfbayes                                                          | github/tfce-mediation                                           |
    #> | github/tfea.chip-downloads                                              | github/tfit                                                     |
    #> | github/tfregulomer                                                      | github/tft                                                      |
    #> | github/tgnet                                                            | github/tgvaughan/icytree                                        |
    #> | github/thapmix                                                          | github/theprecisionlasso                                        |
    #> | github/thermonet                                                        | github/theta                                                    |
    #> | github/thetamater                                                       | github/tibanna                                                  |
    #> | github/tide-for-ptm-search                                              | github/tidehunter                                               |
    #> | github/tigar                                                            | github/tigeri                                                   |
    #> | github/timeseriesnem                                                    | github/timing2                                                  |
    #> | github/tisan                                                            | github/tisk1.0                                                  |
    #> | github/tissue-enrichment-test                                           | github/titer                                                    |
    #> | github/tivan                                                            | github/tksamc                                                   |
    #> | github/tlcr-rl                                                          | github/tmcrys                                                   |
    #> | github/tmtscca                                                          | github/tn-seqexplorer                                           |
    #> | github/tn-test                                                          | github/tner                                                     |
    #> | github/tnm                                                              | github/tnrs                                                     |
    #> | github/tnseq                                                            | github/to-mi                                                    |
    #> | github/toast                                                            | github/tobmi                                                    |
    #> | github/toggle                                                           | github/togogenome                                               |
    #> | github/togostanza                                                       | github/toner                                                    |
    #> | github/tongchf                                                          | github/toolbox-romano-et-al                                     |
    #> | github/topas                                                            | github/tophat-recondition                                       |
    #> | github/topological-measures-wide-analysis                               | github/tormes                                                   |
    #> | github/towards-reliable-bioner                                          | github/toxreporter                                              |
    #> | github/tpglda                                                           | github/tpmcalculator                                            |
    #> | github/tpot                                                             | github/tpot-fss                                                 |
    #> | github/tqec                                                             | github/trac                                                     |
    #> | github/tractaviewer                                                     | github/trader                                                   |
    #> | github/trainbenchmark                                                   | github/traitar                                                  |
    #> | github/transcriptclean                                                  | github/transcriptome-assembly-pipeline                          |
    #> | github/transfer-learning-bner-bioinformatics-2018                       | github/transformphenotype                                       |
    #> | github/transit                                                          | github/transmart-xnat-connector                                 |
    #> | github/transphylo                                                       | github/transposons                                              |
    #> | github/trap                                                             | github/traveler                                                 |
    #> | github/traxtile-public                                                  | github/trca-ssvep                                               |
    #> | github/tree-hmm                                                         | github/treecluster                                              |
    #> | github/treeexp-tutorial                                                 | github/treegrafter                                              |
    #> | github/treehouse                                                        | github/treemerge                                                |
    #> | github/treepl                                                           | github/treescaper                                               |
    #> | github/treesgs                                                          | github/treeshapestats                                           |
    #> | github/treeshrink                                                       | github/treewas                                                  |
    #> | github/trig                                                             | github/triocnv                                                  |
    #> | github/triomdr                                                          | github/tripal-analysis-expression                               |
    #> | github/tripal-elasticsearch                                             | github/tripoly                                                  |
    #> | github/trna-read-mapping                                                | github/tron                                                     |
    #> | github/trove                                                            | github/trs                                                      |
    #> | github/trtcrd/slim                                                      | github/truc                                                     |
    #> | github/trup                                                             | github/tsam                                                     |
    #> | github/tsapa                                                            | github/tscan                                                    |
    #> | github/tsd                                                              | github/tsis                                                     |
    #> | github/tsnad                                                            | github/tsnet                                                    |
    #> | github/tsvdb                                                            | github/ttree-source                                             |
    #> | github/ttt                                                              | github/tumor-clones                                             |
    #> | github/tumorcnv                                                         | github/tunnex                                                   |
    #> | github/tusv                                                             | github/tvb-empirical-data-pipeline                              |
    #> | github/tw2                                                              | github/twisst                                                   |
    #> | github/two-layer-qrcode                                                 | github/twopaco                                                  |
    #> | github/typeloader                                                       | github/u-pass                                                   |
    #> | github/ubd                                                              | github/ublastx-stageone                                         |
    #> | github/uclinfectionimmunity/decombinator                                | github/ucscnanopore-data                                        |
    #> | github/uea-srna-workbench                                               | github/ukbrest                                                  |
    #> | github/ularcirc                                                         | github/umediation                                               |
    #> | github/umib                                                             | github/umis                                                     |
    #> | github/umons-taichi                                                     | github/uncertainty-components                                   |
    #> | github/uncover                                                          | github/uncurl-python                                            |
    #> | github/under                                                            | github/undwive                                                  |
    #> | github/uniconsig                                                        | github/unicycler                                                |
    #> | github/uniqtag                                                          | github/uniqtag-paper                                            |
    #> | github/unity                                                            | github/universal-tts                                            |
    #> | github/uorf-tools                                                       | github/upho                                                     |
    #> | github/uproc                                                            | github/upsetr                                                   |
    #> | github/upstrap                                                          | github/uqsa                                                     |
    #> | github/usefulplants-indicator                                           | github/uslcount                                                 |
    #> | github/uspikehunter                                                     | github/uv-posit                                                 |
    #> | github/uvp                                                              | github/vacceed                                                  |
    #> | github/valor                                                            | github/vankrevelenlocal                                         |
    #> | github/vapid                                                            | github/vapper                                                   |
    #> | github/vapr                                                             | github/varan-gie                                                |
    #> | github/varanto                                                          | github/varcap                                                   |
    #> | github/varcmp                                                           | github/vargenius                                                |
    #> | github/vargeno                                                          | github/vari                                                     |
    #> | github/vari-merge                                                       | github/variancecomponenttest.jl                                 |
    #> | github/variant-validation                                               | github/variantbam                                               |
    #> | github/variantsdwh                                                      | github/varianttoolchest                                         |
    #> | github/varmatch                                                         | github/varscot                                                  |
    #> | github/vashr                                                            | github/vasp                                                     |
    #> | github/vbt-trioanalysis                                                 | github/vccri/c3                                                 |
    #> | github/vcf-kit                                                          | github/vcf-plotein                                              |
    #> | github/vcfanno                                                          | github/vcftotree-3.0.0                                          |
    #> | github/vcmicrobiome                                                     | github/vcnet                                                    |
    #> | github/vcselection                                                      | github/vdjdb-db                                                 |
    #> | github/vdjer                                                            | github/vdjtools                                                 |
    #> | github/vdjviz                                                           | github/vecscreen-plus-taxonomy                                  |
    #> | github/vep-plugins                                                      | github/verse                                                    |
    #> | github/versioned-data                                                   | github/vertical-tracts                                          |
    #> | github/vgteam/vg                                                        | github/vhost-classifier                                         |
    #> | github/vic                                                              | github/video-database                                           |
    #> | github/viennarna                                                        | github/viewbs                                                   |
    #> | github/vifi                                                             | github/vigen                                                    |
    #> | github/vimco                                                            | github/vip                                                      |
    #> | github/viper                                                            | github/vipr-mat-peptide                                         |
    #> | github/virapipe                                                         | github/virbin                                                   |
    #> | github/virgena                                                          | github/viromic-diversity                                        |
    #> | github/virtual-reality-chemical-space                                   | github/virulign                                                 |
    #> | github/viruses-classifier                                               | github/visfeature                                               |
    #> | github/visid                                                            | github/visioprot-ms                                             |
    #> | github/visjs2jupyter                                                    | github/vislab/eeg-blinks                                        |
    #> | github/visnormsc                                                        | github/visor                                                    |
    #> | github/vizardous                                                        | github/vomm                                                     |
    #> | github/voronto                                                          | github/voxelmorph                                               |
    #> | github/vp                                                               | github/vpac                                                     |
    #> | github/vpot                                                             | github/vseams                                                   |
    #> | github/vsearch                                                          | github/vt                                                       |
    #> | github/w4cseq                                                           | github/walking-rdf-and-owl                                      |
    #> | github/walt                                                             | github/warfarin-cohort                                          |
    #> | github/warp                                                             | github/warpgroup                                                |
    #> | github/waternn                                                          | github/wavecnv                                                  |
    #> | github/wcmc-research-informatics                                        | github/wdnfinder                                                |
    #> | github/wdstar                                                           | github/webgivi                                                  |
    #> | github/webmolcs                                                         | github/wenda                                                    |
    #> | github/wenpingd                                                         | github/wfes                                                     |
    #> | github/wft4galaxy                                                       | github/wgd                                                      |
    #> | github/wgfast                                                           | github/wgs-analysis-demo                                        |
    #> | github/whatshap                                                         | github/whisper                                                  |
    #> | github/wholecellsimdb                                                   | github/wipp                                                     |
    #> | github/wisecondorx                                                      | github/wish                                                     |
    #> | github/workflow-of-chipseq                                              | github/workflowr                                                |
    #> | github/workflows                                                        | github/worm-single-cell-data-and-codes                          |
    #> | github/wscunmix                                                         | github/wtchg-rg                                                 |
    #> | github/wuetal2017                                                       | github/wyrm                                                     |
    #> | github/x-cnn                                                            | github/x2h                                                      |
    #> | github/xander-assembler                                                 | github/xenofilter                                               |
    #> | github/xenophile                                                        | github/xgsa                                                     |
    #> | github/xifdr                                                            | github/xispec                                                   |
    #> | github/xitosbml                                                         | github/xlsearch                                                 |
    #> | github/xmabio                                                           | github/xmanv2                                                   |
    #> | github/xmrf                                                             | github/xmwas                                                    |
    #> | github/xolotl                                                           | github/xomeblender                                              |
    #> | github/xulabs/projects                                                  | github/yalff                                                    |
    #> | github/yalla                                                            | github/yamda                                                    |
    #> | github/yeast-segmentation                                               | github/ymap                                                     |
    #> | github/yoann-dufresne/slim                                              | github/yoc                                                      |
    #> | github/yoheirosen/vg                                                    | github/ytzhong/projects                                         |
    #> | github/z-spacing                                                        | github/z-spacing-spark                                          |
    #> | github/z-transform-method                                               | github/zerone                                                   |
    #> | github/zerosum                                                          | github/zibr                                                     |
    #> | github/zing                                                             | github/zombi                                                    |
    #> | github/zorro                                                            | gitlab/leen                                                     |
    #> | gmap                                                                    | golang                                                          |
    #> | gridss                                                                  | gvmap                                                           |
    #> | higlass                                                                 | hisat2                                                          |
    #> | htseq                                                                   | htslib                                                          |
    #> | i-boost                                                                 | igraph                                                          |
    #> | iguide                                                                  | igv                                                             |
    #> | image/imagej                                                            | in-silico-labeling                                              |
    #> | interproscan                                                            | irfinder                                                        |
    #> | iseq                                                                    | isop                                                            |
    #> | ivtnmr                                                                  | jaffa                                                           |
    #> | jvarkit                                                                 | ldsc                                                            |
    #> | leafcutter                                                              | leap                                                            |
    #> | libgtextutils                                                           | liftover                                                        |
    #> | liteq                                                                   | lofreq                                                          |
    #> | lsmm                                                                    | lzo                                                             |
    #> | lzop                                                                    | macs                                                            |
    #> | maftools                                                                | magic                                                           |
    #> | manta                                                                   | maorm                                                           |
    #> | mapsplice2                                                              | marginphase                                                     |
    #> | marvel                                                                  | mask-rcnn                                                       |
    #> | mcorr                                                                   | mdseq                                                           |
    #> | mecat                                                                   | melissa                                                         |
    #> | micropro                                                                | mimosca                                                         |
    #> | miniconda2                                                              | miniconda3                                                      |
    #> | mistic                                                                  | mixcr                                                           |
    #> | mmsplice                                                                | mosdepth                                                        |
    #> | mrnn                                                                    | mscentipede                                                     |
    #> | multiqc                                                                 | mutect                                                          |
    #> | mutscan                                                                 | ncboost                                                         |
    #> | nextflow                                                                | ngb                                                             |
    #> | ngs-qc-toolkit                                                          | ngstk                                                           |
    #> | nodejs                                                                  | novoalign                                                       |
    #> | oases                                                                   | olego                                                           |
    #> | oncotator                                                               | opencpu                                                         |
    #> | other/paradigm                                                          | outrigger                                                       |
    #> | paga                                                                    | paletter                                                        |
    #> | pathopredictor                                                          | pcgr                                                            |
    #> | pcre                                                                    | phenograph                                                      |
    #> | phenopredict                                                            | picard                                                          |
    #> | picky                                                                   | pigz                                                            |
    #> | pindel                                                                  | plier                                                           |
    #> | plotly.js                                                               | plyranges                                                       |
    #> | prada                                                                   | prinseq                                                         |
    #> | prosit                                                                  | protein/pv                                                      |
    #> | pxz                                                                     | qap                                                             |
    #> | quest                                                                   | qvalue                                                          |
    #> | r-color-palettes                                                        | r2d3                                                            |
    #> | raceid                                                                  | radia                                                           |
    #> | rapid                                                                   | rca                                                             |
    #> | recount                                                                 | reditools                                                       |
    #> | reffa/defuse                                                            | reffa/encode-hg19                                               |
    #> | reffa/encode-hg19-ataqc                                                 | reffa/encode-hg19-bowtie2-index                                 |
    #> | reffa/encode-hg19-bwa-index                                             | reffa/encode-hg38                                               |
    #> | reffa/encode-hg38-ataqc                                                 | reffa/encode-hg38-bowtie2-index                                 |
    #> | reffa/encode-hg38-bwa-index                                             | reffa/encode-mm10                                               |
    #> | reffa/encode-mm10-ataqc                                                 | reffa/encode-mm10-bowtie2-index                                 |
    #> | reffa/encode-mm10-bwa-index                                             | reffa/encode-mm9                                                |
    #> | reffa/encode-mm9-ataqc                                                  | reffa/encode-mm9-bowtie2-index                                  |
    #> | reffa/encode-mm9-bwa-index                                              | reffa/ensemble                                                  |
    #> | reffa/fusioncatcher                                                     | reffa/genecode                                                  |
    #> | reffa/rmats                                                             | reffa/ucsc                                                      |
    #> | relaxed                                                                 | resm                                                            |
    #> | rhat                                                                    | rmats                                                           |
    #> | root                                                                    | rop                                                             |
    #> | rum                                                                     | samblaster                                                      |
    #> | samstat                                                                 | samtools                                                        |
    #> | saver                                                                   | scgen                                                           |
    #> | sclvm                                                                   | scnorm                                                          |
    #> | scrabble                                                                | scvi                                                            |
    #> | sda                                                                     | se-mei                                                          |
    #> | selene                                                                  | sellerslab-gemini                                               |
    #> | seqtk                                                                   | seurat                                                          |
    #> | simlr-py                                                                | singlesplice                                                    |
    #> | skmer                                                                   | sleuth                                                          |
    #> | snakemake/dna-seq-gatk-variant-calling                                  | snakemake/sequana                                               |
    #> | snpeff                                                                  | solexaqa                                                        |
    #> | somatic-sniper                                                          | spack                                                           |
    #> | sparsehash                                                              | speedseq                                                        |
    #> | sqlite                                                                  | sratools                                                        |
    #> | srnanalyzer                                                             | ssaha2                                                          |
    #> | star                                                                    | strawberry                                                      |
    #> | strelka                                                                 | stringtie                                                       |
    #> | subread                                                                 | sv2                                                             |
    #> | svaba                                                                   | svscore                                                         |
    #> | svtools                                                                 | taxmaps                                                         |
    #> | tbb                                                                     | test/github-demo                                                |
    #> | threadpool                                                              | three.js                                                        |
    #> | tmap                                                                    | topconfects                                                     |
    #> | tophat                                                                  | tracer                                                          |
    #> | trendsceek                                                              | trimgalore                                                      |
    #> | trinityrnaseq                                                           | tvc                                                             |
    #> | ucsc                                                                    | unet-segmentation                                               |
    #> | unifrac                                                                 | useq                                                            |
    #> | vadir                                                                   | varscan                                                         |
    #> | vcf2maf                                                                 | vcflib                                                          |
    #> | vcftools                                                                | vep                                                             |
    #> | vg                                                                      | vifi                                                            |
    #> | viper                                                                   | viral-ngs                                                       |
    #> | wdl/antonkulaga                                                         | wdl/bgm-cwg-whole-exome-sequencing                              |
    #> | wdl/bgm-cwg-whole-genome-sequencing                                     | wdl/biowdl-aligning                                             |
    #> | wdl/biowdl-bam-to-gvcf                                                  | wdl/biowdl-chip-seq                                             |
    #> | wdl/biowdl-gams                                                         | wdl/biowdl-germline-dna                                         |
    #> | wdl/biowdl-jointgenotyping                                              | wdl/biowdl-pipeline-template                                    |
    #> | wdl/biowdl-qc                                                           | wdl/biowdl-rna-seq                                              |
    #> | wdl/biowdl-somatic-variantcalling                                       | wdl/biowdl-tasks                                                |
    #> | wdl/biowdl-virus-assembly                                               | wdl/brass                                                       |
    #> | wdl/deepvar                                                             | wdl/deepvariant-glnexus                                         |
    #> | wdl/encode-chipseq2                                                     | wdl/firecloud-tools                                             |
    #> | wdl/gatk-3-4-intel-germline-snps-indels                                 | wdl/gatk-3-4-rnaseq-germline-snps-indels                        |
    #> | wdl/gatk-3-data-processing                                              | wdl/gatk-3-intel-germline-snps-indels                           |
    #> | wdl/gatk-4-data-processing                                              | wdl/gatk-4-germline-snps-indels                                 |
    #> | wdl/gatk-4-intel-germline-snps-indels                                   | wdl/gatk-4-intel-somatic-with-preprocessing                     |
    #> | wdl/gatk-4-somatic-cnvs                                                 | wdl/gatk-4-somatic-snvs-indels                                  |
    #> | wdl/gatk-4-somatic-with-preprocessing                                   | wdl/gatk-five-dollar-genome-analysis                            |
    #> | wdl/gatk-intel-faster-alternative-samples                               | wdl/gatk-seq-format-conversion                                  |
    #> | wdl/gatk-seq-format-validation                                          | wdl/gatk-wgs-germline-snps-indels                               |
    #> | wdl/genomics-iter-developers-wgs-wes-germline                           | wdl/hongiiv-gatk-workflows                                      |
    #> | wdl/jimmybgammyknee-workflows                                           | wdl/johnmous-chip-seq                                           |
    #> | wdl/kurt-hetrick-variants                                               | wdl/mayomics-vc                                                 |
    #> | wdl/metagenophile                                                       | wdl/mgs-sop                                                     |
    #> | wdl/oskarvid-germline-pipeline                                          | wdl/ottojolanki-kallisto                                        |
    #> | wdl/pacbio-tools                                                        | wdl/raisok-rnaseq                                               |
    #> | wdl/saige                                                               | wdl/sanderest-cromwell                                          |
    #> | wdl/seonwhee-genome-gatk                                                | wdl/tsca-firecloud                                              |
    #> | wdl/uofabio                                                             | wdl/wgsa-parsr                                                  |
    #> | wham                                                                    | xz                                                              |
    #> | zifa                                                                    | zinbwave                                                        |
    #> | zlib                                                                    |
    #> +++

Maintainer
----------

-   \[@Jianfeng\](<a href="https://github.com/Miachol" class="uri">https://github.com/Miachol</a>)

License
-------

Academic Free License version 3.0
