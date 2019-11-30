<img src="https://img.shields.io/badge/lifecycle-experimental-orange.svg" alt="Life cycle: experimental">
<a href="https://godoc.org/github.com/openbiox/bget"><img src="https://godoc.org/github.com/openbiox/bget?status.svg" alt="GoDoc"></a>

bget
====

Lightweight downloader for bioinformatics data, databases and files.
Golang `http` library, `wget`, `curl`, `axel`, `git`, and `rsync` were
supported as the download engine.

Possible URLs pool:

-   Reference genomes
-   Source code of bioinformatics tools
-   Bioinformatics databases and files
-   Papers material
-   ……

Why we want to create this tool?

We can query PDF of the manuscript via using Endnote or sci-hub.
However, you can not easily get the supplementary files of scientific
papers based on the two ways.

![doi
demo](https://github.com/openbiox/bget/raw/master/doc/static/doi.gif)

Here, we are developing and sharing an open-source tool bget with `doi`
subcommand to query supplementary files of scientific papers. The
journals with high impact factors or those integrative publishers are a
higher priority in our development plan:
<a href="https://github.com/openbiox/bget/blob/master/doc/doi.md" class="uri">https://github.com/openbiox/bget/blob/master/doc/doi.md</a>

It is noted that we do not want to distribute any pirated resources or
cause unnecessary network congestion. We hope this tool can provide an
optional method to more easily query related files of scientific papers.
Please use it in a non-invasive way (i.e. high concurrency, long
continuous request).

Prerequisities
--------------

-   Headless Chrome is required for some of website with JavaScript
    driven render pages.

Install Chrome on Linux:

-   [How to Install Google Chrome Web Browser on
    Ubuntu](https://linuxize.com/post/how-to-install-google-chrome-web-browser-on-ubuntu-18-04/)
-   [How to Install Google Chrome Web Browser on
    CentOS](https://linuxize.com/post/how-to-install-google-chrome-web-browser-on-centos-7/)

For windows users, you may need to create an alias of Chrome to make
[chromedp](https://github.com/chromedp/chromedp) work.

Installation
------------

    # download bget on MAC OSX
    wget -c https://github.com/openbiox/bget/releases/download/v0.1.3/bget_osx
    mv bget_osx bget
    chmod a+x bget
    #
    # download bget on Linux
    wget -c https://github.com/openbiox/bget/releases/download/v0.1.3/bget_linux64
    mv bget_linux64 bget
    chmod a+x bget
    #
    # download bget on Windows
    wget -c https://github.com/openbiox/bget/releases/download/v0.1.3/bget.exe
    #
    # get latest version
    go get -u github.com/openbiox/bget

Usage
-----

Demo Video:
<a href="https://www.notion.so/sjtu/Demo-of-bget-doi-key-seq-url-78c2c334bf894668aa17fd128bd3255c" class="uri">https://www.notion.so/sjtu/Demo-of-bget-doi-key-seq-url-78c2c334bf894668aa17fd128bd3255c</a>

    bget -h

    #> Lightweight downloader for bioinformatics data, databases and files (under development). It will provides a simple and parallelized method to access various bioinformatics resoures. More see here https://github.com/openbiox/bget.
    #> 
    #> Usage:
    #>   bget [flags]
    #>   bget [command]
    #> 
    #> Available Commands:
    #>   doi         Can be used to access files via DOI.
    #>   help        Help about any command
    #>   key         Can be used to access URLs via a key string.
    #>   seq         Can be used to access sequence data via unique id (dbGAP and EGA) or manifest files (TCGA).
    #>   url         Can be used to access URLs via Golang http, wget, curl, axel and git, and rsync.
    #> 
    #> Flags:
    #>       --clean                    Remove _download and _log in current dir.
    #>   -g, --engine string            Point the download engine: go-http, wget, curl, axel, git, and rsync. (default "go-http")
    #>   -e, --extra-cmd string         Extra flags and values pass to internal CMDs
    #>   -h, --help                     help for bget
    #>       --ignore                   Contine to download and skip the check of existed files.
    #>   -l, --list-file string         A file contains dois for download.
    #>       --log-dir string           Log dir. (default "/Users/apple/Documents/github/bget/_log")
    #>   -m, --mirror string            Set the mirror of resources.
    #>   -o, --outdir string            Set the download dir. (default "/Users/apple/Documents/github/bget")
    #>   -f, --overwrite                Logical indicating that whether to overwrite existing files.
    #>       --proxy string             HTTP proxy to download.
    #>   -q, --quiet                    No output.
    #>   -n, --remote-name              Use remote defined filename.
    #>   -r, --retries int              Retry specifies the number of attempts to retrieve the data. (default 5)
    #>       --retries-sleep-time int   Sleep time after one retry. (default 5)
    #>       --save-log                 Save download log to local file]. (default true)
    #>   -s, --seperator string         Optional 'url1{seperator}url2' for multiple keys, urls, or seqs. (default ",")
    #>       --task-id string           Task ID (random). (default "l6yd3ys032iq807")
    #>   -t, --thread int               Concurrency download thread. (default 1)
    #>       --thread-axel int          Set the thread of axel. (default 5)
    #>       --timeout int              Set the timeout of per request. (default 35)
    #>   -u, --uncompress               Uncompress download files for .zip, .tar.gz, and .gz suffix files.
    #>       --version                  version for bget
    #> 
    #> Use "bget [command] --help" for more information about a command.

### bget doi

The `bget doi` supported website and journals are continuely increasing.

**Warn**: If you do not follow the policies of the relevant website
(i.e. continuous download or limited copyright), you will lose the
authorization to use this tool.

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
    #>       --full-text   Access full text. (default true)
    #>   -h, --help        help for doi
    #>       --pmc         Try PMC database.
    #>       --suppl       Access supplementary files.
    #>       --universe    Try universe spider. (default true)
    #> 
    #> Global Flags:
    #>   -g, --engine string            Point the download engine: go-http, wget, curl, axel, git, and rsync. (default "go-http")
    #>   -e, --extra-cmd string         Extra flags and values pass to internal CMDs
    #>       --ignore                   Contine to download and skip the check of existed files.
    #>   -l, --list-file string         A file contains dois for download.
    #>       --log-dir string           Log dir. (default "/Users/apple/Documents/github/bget/_log")
    #>   -m, --mirror string            Set the mirror of resources.
    #>   -o, --outdir string            Set the download dir. (default "/Users/apple/Documents/github/bget")
    #>   -f, --overwrite                Logical indicating that whether to overwrite existing files.
    #>       --proxy string             HTTP proxy to download.
    #>   -q, --quiet                    No output.
    #>   -n, --remote-name              Use remote defined filename.
    #>   -r, --retries int              Retry specifies the number of attempts to retrieve the data. (default 5)
    #>       --retries-sleep-time int   Sleep time after one retry. (default 5)
    #>       --save-log                 Save download log to local file]. (default true)
    #>   -s, --seperator string         Optional 'url1{seperator}url2' for multiple keys, urls, or seqs. (default ",")
    #>       --task-id string           Task ID (random). (default "lpzqo8faon2xqsu")
    #>   -t, --thread int               Concurrency download thread. (default 1)
    #>       --timeout int              Set the timeout of per request. (default 35)
    #>   -u, --uncompress               Uncompress download files for .zip, .tar.gz, and .gz suffix files.

### bget key

`bget key` can be used to download a set of files via a alias name key.

    bget key -h

    #> Can be used to access URLs via a key string. e.g. 'item' or 'item@version #releaseVersion', : bwa, reffa-defuse@GRCh38 #97. More see here https://github.com/openbiox/bget.
    #> 
    #> Usage:
    #>   bget key [key1 key2 key3...] [flags]
    #> 
    #> Examples:
    #>   bget key aligner/bwa
    #>   bget key -a // get all available keys
    #>   bget key seq/samtools -v // view all samtools available versions in CMD table
    #>   bget key seq/samtools -v --out-json // view all samtools available versions in JSON format
    #>   bget key "reffa/defuse@GRCh38 #97" -t 10 -f
    #>   bget key reffa/defuse@GRCh38 release=97 -t 10 -f
    #>   bget key db/annovar@clinvar_20170501 db/annovar@clinvar_20180603 builder=hg38
    #> 
    #>   bget key db/annovar -v --out-text
    #>   bget key db/annovar version='clinvar_20131105, clinvar_20140211, clinvar_20140303, clinvar_20140702, clinvar_20140902, clinvar_20140929, clinvar_20150330, clinvar_20150629, clinvar_20151201, clinvar_20160302, clinvar_20161128, clinvar_20170130, clinvar_20170501, clinvar_20170905, clinvar_20180603, avsnp150, avsnp147, avsnp144, avsnp142, avsnp138, cadd, caddgt10, caddgt20, cadd13, cadd13gt10, cadd13gt20, cg69, cg46, cosmic70, cosmic68wgs, cosmic68, cosmic67wgs, cosmic67, cosmic65, cosmic64, dbnsfp35a, dbnsfp33a, dbnsfp31a_interpro, dbnsfp30a, dbscsnv11, eigen, esp6500siv2_ea, esp6500siv2_aa, esp6500siv2_all, exac03nontcga, exac03nonpsych, exac03, fathmm, gerp++gt2, gme, gnomad_exome, gnomad_genome, gwava, hrcr1, icgc21, intervar_20170202, kaviar_20150923, ljb26_all, mcap, mitimpact2, mitimpact24, nci60, popfreq_max_20150413, popfreq_all_20150413, revel, regsnpintron' builder=hg19 -t 10 -f
    #> 
    #> Flags:
    #>       --autopath        Logical indicating that whether to create subdir in download dir: e.g. reffa/{{key}}/
    #>   -h, --help            help for key
    #>   -a, --keys-all        Show all available string key can be download.
    #>       --out-json        Output information in JSON string
    #>       --out-text        Output information in plain text
    #>   -v, --show-versions   Show all available versions of key.
    #> 
    #> Global Flags:
    #>   -g, --engine string            Point the download engine: go-http, wget, curl, axel, git, and rsync. (default "go-http")
    #>   -e, --extra-cmd string         Extra flags and values pass to internal CMDs
    #>       --ignore                   Contine to download and skip the check of existed files.
    #>   -l, --list-file string         A file contains dois for download.
    #>       --log-dir string           Log dir. (default "/Users/apple/Documents/github/bget/_log")
    #>   -m, --mirror string            Set the mirror of resources.
    #>   -o, --outdir string            Set the download dir. (default "/Users/apple/Documents/github/bget")
    #>   -f, --overwrite                Logical indicating that whether to overwrite existing files.
    #>       --proxy string             HTTP proxy to download.
    #>   -q, --quiet                    No output.
    #>   -n, --remote-name              Use remote defined filename.
    #>   -r, --retries int              Retry specifies the number of attempts to retrieve the data. (default 5)
    #>       --retries-sleep-time int   Sleep time after one retry. (default 5)
    #>       --save-log                 Save download log to local file]. (default true)
    #>   -s, --seperator string         Optional 'url1{seperator}url2' for multiple keys, urls, or seqs. (default ",")
    #>       --task-id string           Task ID (random). (default "05xfxc544iimmut")
    #>   -t, --thread int               Concurrency download thread. (default 1)
    #>       --timeout int              Set the timeout of per request. (default 35)
    #>   -u, --uncompress               Uncompress download files for .zip, .tar.gz, and .gz suffix files.

    ## show all supported items
    bget key -a
    #> ++++
    #> | aligner/bazam                             | aligner/blast                                 | aligner/blat                           |
    #> | aligner/bowtie                            | aligner/bowtie2                               | aligner/bwa                            |
    #> | aligner/dart                              | aligner/gmap                                  | aligner/hisat2                         |
    #> | aligner/novoalign                         | aligner/olego                                 | aligner/rhat                           |
    #> | aligner/rum                               | aligner/ssaha2                                | aligner/star                           |
    #> | aligner/taxmaps                           | aligner/tmap                                  | aligner/tophat                         |
    #> | anno/annovar                              | anno/annovarr                                 | anno/bystro                            |
    #> | anno/ensembl-vep                          | anno/gemini                                   | anno/giggle                            |
    #> | anno/oncotator                            | anno/pathopredictor                           | anno/pcgr                              |
    #> | anno/snpeff                               | anno/svscore                                  | anno/vcfanno                           |
    #> | app/babun                                 | app/cmder                                     | app/iontorrent-suite                   |
    #> | app/orange3                               | assemble/abyss                                | assemble/edena                         |
    #> | assemble/marvel                           | assemble/mecat                                | assemble/oases                         |
    #> | assemble/sda                              | assemble/stringtie                            | assemble/trinityrnaseq                 |
    #> | assemble/velvet                           | base/bapi                                     | base/bcbio-nextgen                     |
    #> | base/bget                                 | base/bioinstaller                             | base/bpipe                             |
    #> | base/bzip2                                | base/cromwell                                 | base/curl                              |
    #> | base/fatotwobit                           | base/gdc-client                               | base/golang                            |
    #> | base/htslib                               | base/igraph                                   | base/igv                               |
    #> | base/liftover                             | base/lzo                                      | base/lzop                              |
    #> | base/miniconda2                           | base/miniconda3                               | base/nextflow                          |
    #> | base/ngstk                                | base/nodejs                                   | base/pcre                              |
    #> | base/pigz                                 | base/pxz                                      | base/root                              |
    #> | base/spack                                | base/sqlite                                   | base/sratools                          |
    #> | base/tbb                                  | base/ucsc                                     | base/xz                                |
    #> | base/zlib                                 | db/annovar                                    | db/annovar-1000g                       |
    #> | db/annovar-ensgene                        | db/annovar-knowngene                          | db/annovar-noidx                       |
    #> | db/annovar-refgene                        | db/appris                                     | db/atcircdb                            |
    #> | db/awesome                                | db/biosystems                                 | db/cancer_hotspots                     |
    #> | db/cancersplicingqtl                      | db/cellmarker                                 | db/cgi                                 |
    #> | db/circbase                               | db/circnet                                    | db/circrnadb                           |
    #> | db/civic                                  | db/consensuspathdb                            | db/cscd                                |
    #> | db/dbsno                                  | db/dcdb                                       | db/denovo_db                           |
    #> | db/dgidb                                  | db/diseaseenhancer                            | db/eggnog                              |
    #> | db/ewasdb                                 | db/exorbase                                   | db/expression_atlas                    |
    #> | db/exsnp                                  | db/fantom_cage_peaks                          | db/fantom_co_expression_clusters       |
    #> | db/fantom_enhancers                       | db/fantom_motifs                              | db/fantom_ontology                     |
    #> | db/fantom_tss_classifier                  | db/funcoup                                    | db/fusiongdb                           |
    #> | db/hgnc                                   | db/hmdb                                       | db/hpdi                                |
    #> | db/hpo                                    | db/inbiomap                                   | db/instruct                            |
    #> | db/interpro                               | db/islandviewer                               | db/journal-doaj                        |
    #> | db/lnc2cancer                             | db/lncediting                                 | db/lncrnadisease                       |
    #> | db/mircancer                              | db/mirdb                                      | db/mirnest                             |
    #> | db/mirtarbase                             | db/mndr                                       | db/msdd                                |
    #> | db/omim_open                              | db/omim_private                               | db/oncokb                              |
    #> | db/oncomirdb                              | db/oncotator                                  | db/pancanqtl                           |
    #> | db/phosphonetworks                        | db/pmkb                                       | db/proteinatlas                        |
    #> | db/rbp_var                                | db/rbpdb                                      | db/rddpred                             |
    #> | db/redoxdb                                | db/remap                                      | db/remap2                              |
    #> | db/rsnp3                                  | db/rvarbase                                   | db/sedb                                |
    #> | db/seecancer                              | db/seeqtl                                     | db/sm2mir                              |
    #> | db/srnanalyzer                            | db/tumorfusions                               | db/ucsc-cytoband                       |
    #> | db/ucsc-dnase-clustered                   | db/ucsc-ensgene                               | db/ucsc-knowngene                      |
    #> | db/ucsc-refgene                           | db/ucsc-tfbs-clustered                        | db/varcards                            |
    #> | doc/awesome-single-cell                   | doc/awosome-bioinformatics                    | doc/splatter-paper                     |
    #> | doc/trackviewer                           | exp/arnapipe                                  | exp/asap                               |
    #> | exp/backspin                              | exp/ballgown                                  | exp/bearscc                            |
    #> | exp/brie                                  | exp/cellfishing.jl                            | exp/cellsius                           |
    #> | exp/circbrain                             | exp/clonealign                                | exp/conos                              |
    #> | exp/dca                                   | exp/deepcell-tf                               | exp/dropclust                          |
    #> | exp/f-sclvm                               | exp/htseq                                     | exp/irfinder                           |
    #> | exp/isop                                  | exp/leafcutter                                | exp/magic                              |
    #> | exp/mapsplice2                            | exp/mdseq                                     | exp/mimosca                            |
    #> | exp/mistic                                | exp/mixcr                                     | exp/mmsplice                           |
    #> | exp/mrnn                                  | exp/outrigger                                 | exp/phenograph                         |
    #> | exp/phenopredict                          | exp/plier                                     | exp/prada                              |
    #> | exp/raceid                                | exp/rca                                       | exp/resm                               |
    #> | exp/rmats                                 | exp/saver                                     | exp/scgen                              |
    #> | exp/sclvm                                 | exp/scnorm                                    | exp/scrabble                           |
    #> | exp/scvi                                  | exp/seurat                                    | exp/simlr-py                           |
    #> | exp/singlesplice                          | exp/sleuth                                    | exp/srnanalyzer                        |
    #> | exp/strawberry                            | exp/subread                                   | exp/topconfects                        |
    #> | exp/tracer                                | exp/trendsceek                                | exp/viper                              |
    #> | exp/zifa                                  | exp/zinbwave                                  | image/cdeep3m                          |
    #> | image/cellprofiler                        | image/fmriprep                                | image/imagej                           |
    #> | image/in-silico-labeling                  | image/leap                                    | image/mask-rcnn                        |
    #> | image/unet-segmentation                   | mut/advntr                                    | mut/agfusion                           |
    #> | mut/annovar                               | mut/atlas2                                    | mut/bcftools                           |
    #> | mut/beagle                                | mut/breakdancer                               | mut/breakmer                           |
    #> | mut/breakpointsurveyor                    | mut/caveman                                   | mut/chia-pet2                          |
    #> | mut/cn-learn                              | mut/cnvkit                                    | mut/cnvnator                           |
    #> | mut/conbase                               | mut/deepvariant                               | mut/delly                              |
    #> | mut/effusion                              | mut/facets                                    | mut/forge                              |
    #> | mut/freebayes                             | mut/freec                                     | mut/fusioncatcher                      |
    #> | mut/gatk-bundle                           | mut/gatk4                                     | mut/gridss                             |
    #> | mut/iseq                                  | mut/jaffa                                     | mut/lofreq                             |
    #> | mut/maftools                              | mut/manta                                     | mut/marginphase                        |
    #> | mut/mutect                                | mut/mutscan                                   | mut/ncboost                            |
    #> | mut/picky                                 | mut/pindel                                    | mut/radia                              |
    #> | mut/rapid                                 | mut/reditools                                 | mut/se-mei                             |
    #> | mut/somatic-sniper                        | mut/speedseq                                  | mut/strelka                            |
    #> | mut/sv2                                   | mut/svaba                                     | mut/svtools                            |
    #> | mut/tvc                                   | mut/vadir                                     | mut/varscan                            |
    #> | mut/vcf2maf                               | mut/vcflib                                    | mut/vcftools                           |
    #> | mut/vg                                    | mut/wham                                      | other/bigstitcher                      |
    #> | other/confined                            | other/g2s                                     | other/libgtextutils                    |
    #> | other/paga                                | other/paradigm                                | other/recount                          |
    #> | other/sellerslab-gemini                   | other/sparsehash                              | other/unifrac                          |
    #> | qc/3dchromatin-replicateqc                | qc/chronqc                                    | qc/detin                               |
    #> | qc/fastqc                                 | qc/fastx-toolkit                              | qc/multiqc                             |
    #> | qc/ngs-qc-toolkit                         | qc/prinseq                                    | qc/solexaqa                            |
    #> | qc/trimgalore                             | reffa/defuse                                  | reffa/encode-hg19                      |
    #> | reffa/encode-hg19-ataqc                   | reffa/encode-hg19-bowtie2-index               | reffa/encode-hg19-bwa-index            |
    #> | reffa/encode-hg38                         | reffa/encode-hg38-ataqc                       | reffa/encode-hg38-bowtie2-index        |
    #> | reffa/encode-hg38-bwa-index               | reffa/encode-mm10                             | reffa/encode-mm10-ataqc                |
    #> | reffa/encode-mm10-bowtie2-index           | reffa/encode-mm10-bwa-index                   | reffa/encode-mm9                       |
    #> | reffa/encode-mm9-ataqc                    | reffa/encode-mm9-bowtie2-index                | reffa/encode-mm9-bwa-index             |
    #> | reffa/ensemble                            | reffa/fusioncatcher                           | reffa/genecode                         |
    #> | reffa/rmats                               | reffa/ucsc                                    | rlib/chimeraviz                        |
    #> | rlib/dash                                 | rlib/easysvg                                  | rlib/genvisr                           |
    #> | rlib/geogrid                              | rlib/ggdag                                    | rlib/ggseqlogo                         |
    #> | rlib/ggthemr                              | rlib/gvmap                                    | rlib/liteq                             |
    #> | rlib/paletter                             | rlib/r-color-palettes                         | rlib/threadpool                        |
    #> | seq/anchor                                | seq/autochrom3d                               | seq/bamtools                           |
    #> | seq/bamutil                               | seq/bedops                                    | seq/bedtools2                          |
    #> | seq/bin3c                                 | seq/biobloom                                  | seq/biopython                          |
    #> | seq/cesa                                  | seq/chicmaxima                                | seq/chromhmm                           |
    #> | seq/chromtime                             | seq/chromvar                                  | seq/cistopic                           |
    #> | seq/deepnovo-dia                          | seq/divers                                    | seq/dreg                               |
    #> | seq/dstruct                               | seq/easeq                                     | seq/fastp                              |
    #> | seq/fastq-tools                           | seq/feast                                     | seq/genomedisco                        |
    #> | seq/iguide                                | seq/interproscan                              | seq/ivtnmr                             |
    #> | seq/jvarkit                               | seq/macs                                      | seq/maorm                              |
    #> | seq/mcorr                                 | seq/melissa                                   | seq/micropro                           |
    #> | seq/mosdepth                              | seq/mscentipede                               | seq/picard                             |
    #> | seq/plyranges                             | seq/prosit                                    | seq/qap                                |
    #> | seq/quest                                 | seq/rop                                       | seq/samblaster                         |
    #> | seq/samstat                               | seq/samtools                                  | seq/selene                             |
    #> | seq/seqtk                                 | seq/skmer                                     | seq/useq                               |
    #> | seq/vifi                                  | seq/viral-ngs                                 | snakemake/dna-seq-gatk-variant-calling |
    #> | snakemake/sequana                         | stats/armadillo                               | stats/i-boost                          |
    #> | stats/ldsc                                | stats/lsmm                                    | stats/qvalue                           |
    #> | test/github-demo                          | wdl/antonkulaga                               | wdl/bgm-cwg-whole-exome-sequencing     |
    #> | wdl/bgm-cwg-whole-genome-sequencing       | wdl/biowdl-aligning                           | wdl/biowdl-bam-to-gvcf                 |
    #> | wdl/biowdl-chip-seq                       | wdl/biowdl-gams                               | wdl/biowdl-germline-dna                |
    #> | wdl/biowdl-jointgenotyping                | wdl/biowdl-pipeline-template                  | wdl/biowdl-qc                          |
    #> | wdl/biowdl-rna-seq                        | wdl/biowdl-somatic-variantcalling             | wdl/biowdl-tasks                       |
    #> | wdl/biowdl-virus-assembly                 | wdl/brass                                     | wdl/deepvar                            |
    #> | wdl/deepvariant-glnexus                   | wdl/encode-chipseq2                           | wdl/firecloud-tools                    |
    #> | wdl/gatk-3-4-intel-germline-snps-indels   | wdl/gatk-3-4-rnaseq-germline-snps-indels      | wdl/gatk-3-data-processing             |
    #> | wdl/gatk-3-intel-germline-snps-indels     | wdl/gatk-4-data-processing                    | wdl/gatk-4-germline-snps-indels        |
    #> | wdl/gatk-4-intel-germline-snps-indels     | wdl/gatk-4-intel-somatic-with-preprocessing   | wdl/gatk-4-somatic-cnvs                |
    #> | wdl/gatk-4-somatic-snvs-indels            | wdl/gatk-4-somatic-with-preprocessing         | wdl/gatk-five-dollar-genome-analysis   |
    #> | wdl/gatk-intel-faster-alternative-samples | wdl/gatk-seq-format-conversion                | wdl/gatk-seq-format-validation         |
    #> | wdl/gatk-wgs-germline-snps-indels         | wdl/genomics-iter-developers-wgs-wes-germline | wdl/hongiiv-gatk-workflows             |
    #> | wdl/jimmybgammyknee-workflows             | wdl/johnmous-chip-seq                         | wdl/kurt-hetrick-variants              |
    #> | wdl/mayomics-vc                           | wdl/metagenophile                             | wdl/mgs-sop                            |
    #> | wdl/oskarvid-germline-pipeline            | wdl/ottojolanki-kallisto                      | wdl/pacbio-tools                       |
    #> | wdl/raisok-rnaseq                         | wdl/saige                                     | wdl/sanderest-cromwell                 |
    #> | wdl/seonwhee-genome-gatk                  | wdl/tsca-firecloud                            | wdl/uofabio                            |
    #> | wdl/wgsa-parsr                            | web/clustergrammer                            | web/echarts                            |
    #> | web/genomeuplot                           | web/higlass                                   | web/ngb                                |
    #> | web/opencpu                               | web/plotly.js                                 | web/protein/pv                         |
    #> | web/r2d3                                  | web/relaxed                                   | web/three.js                           |
    #> ++++

### bget seq

`bget seq` can be used to access [Gene Expression Omnibus
(GEO)](https://www.ncbi.nlm.nih.gov/geo), [Sequence Read Archive
(SRA)](https://www.ncbi.nlm.nih.gov/sra/), and [GDC Data
Portal](https://portal.gdc.cancer.gov/) are supported.

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
    #>       --gpl-geo                 Wheather fetch GPL files from GEO database.
    #>   -h, --help                    help for seq
    #>       --token-file-ega string   Credential file to access EGA archive files, {"username": "{your_user_name}", 
    #>                                   "password": "{your_password}","client_secret":"AMenuDLjVdVo4BSwi0QD54LL6NeVDEZRzEQUJ7h
    #>                                   JOM3g4imDZBHHX0hNfKHPeQIGkskhtCmqAJtt_jm7EKq-rWw"}.
    #>       --token-gdc string        Token to access TCGA portal files.
    #> 
    #> Global Flags:
    #>   -g, --engine string            Point the download engine: go-http, wget, curl, axel, git, and rsync. (default "go-http")
    #>   -e, --extra-cmd string         Extra flags and values pass to internal CMDs
    #>       --ignore                   Contine to download and skip the check of existed files.
    #>   -l, --list-file string         A file contains dois for download.
    #>       --log-dir string           Log dir. (default "/Users/apple/Documents/github/bget/_log")
    #>   -m, --mirror string            Set the mirror of resources.
    #>   -o, --outdir string            Set the download dir. (default "/Users/apple/Documents/github/bget")
    #>   -f, --overwrite                Logical indicating that whether to overwrite existing files.
    #>       --proxy string             HTTP proxy to download.
    #>   -q, --quiet                    No output.
    #>   -n, --remote-name              Use remote defined filename.
    #>   -r, --retries int              Retry specifies the number of attempts to retrieve the data. (default 5)
    #>       --retries-sleep-time int   Sleep time after one retry. (default 5)
    #>       --save-log                 Save download log to local file]. (default true)
    #>   -s, --seperator string         Optional 'url1{seperator}url2' for multiple keys, urls, or seqs. (default ",")
    #>       --task-id string           Task ID (random). (default "ernucgi73gt0xiv")
    #>   -t, --thread int               Concurrency download thread. (default 1)
    #>       --timeout int              Set the timeout of per request. (default 35)
    #>   -u, --uncompress               Uncompress download files for .zip, .tar.gz, and .gz suffix files.

### bget url

`bget url` can be used to access files via input URLs. Golang http,
wget, curl, axel and git, and rsync are support for download process.

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
    #> Flags:
    #>   -h, --help   help for url
    #> 
    #> Global Flags:
    #>   -g, --engine string            Point the download engine: go-http, wget, curl, axel, git, and rsync. (default "go-http")
    #>   -e, --extra-cmd string         Extra flags and values pass to internal CMDs
    #>       --ignore                   Contine to download and skip the check of existed files.
    #>   -l, --list-file string         A file contains dois for download.
    #>       --log-dir string           Log dir. (default "/Users/apple/Documents/github/bget/_log")
    #>   -m, --mirror string            Set the mirror of resources.
    #>   -o, --outdir string            Set the download dir. (default "/Users/apple/Documents/github/bget")
    #>   -f, --overwrite                Logical indicating that whether to overwrite existing files.
    #>       --proxy string             HTTP proxy to download.
    #>   -q, --quiet                    No output.
    #>   -n, --remote-name              Use remote defined filename.
    #>   -r, --retries int              Retry specifies the number of attempts to retrieve the data. (default 5)
    #>       --retries-sleep-time int   Sleep time after one retry. (default 5)
    #>       --save-log                 Save download log to local file]. (default true)
    #>   -s, --seperator string         Optional 'url1{seperator}url2' for multiple keys, urls, or seqs. (default ",")
    #>       --task-id string           Task ID (random). (default "wts26egkedkygea")
    #>   -t, --thread int               Concurrency download thread. (default 1)
    #>       --timeout int              Set the timeout of per request. (default 35)
    #>   -u, --uncompress               Uncompress download files for .zip, .tar.gz, and .gz suffix files.

Maintainer
----------

-   \[@Jianfeng\](<a href="https://github.com/Miachol" class="uri">https://github.com/Miachol</a>)

License
-------

Academic Free License version 3.0
