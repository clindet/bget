<img src="https://img.shields.io/badge/lifecycle-experimental-orange.svg" alt="Life cycle: experimental"> [![GoDoc](https://godoc.org/github.com/openbiox/bget?status.svg)](https://godoc.org/github.com/openbiox/bget)

# bget

Lightweight downloader for bioinformatics data, databases and files. Golang `http` library, `wget`, `curl`, `axel`, `git`, and `rsync` were supported as the download engine.

Possible URLs pool:

- Reference genomes
- Source code of bioinformatics tools
- Bioinformatics databases and files
- Papers material
- ......

## Installation

```bash
go get -u github.com/openbiox/bget
```

## Usage

```bash
Lightweight downloader for bioinformatics data, databases and files (under development). It will provides a simple and parallelized method to access various bioinformatics resoures. More see here https://github.com/openbiox/bget.

Usage:
  bget [flags]
  bget [command]

Available Commands:
  doi         Can be used to access files via DOI.
  help        Help about any command
  key         Can be used to access URLs via a key string.
  seq         Can be used to access sequence data via unique id (dbGAP and EGA) or manifest files (TCGA).
  url         Can be used to access URLs via Golang http, wget, curl, axel and git, and rsync.

Flags:
      --clean                    Remove _download and _log in current dir.
  -g, --engine string            Point the download engine: go-http, wget, curl, axel, git, and rsync. (default "go-http")
  -e, --extra-cmd string         Extra flags and values pass to internal CMDs
  -h, --help                     help for bget
      --ignore                   Contine to download and skip the check of existed files.
  -l, --list-file string         A file contains dois for download.
      --log-dir string           Log dir. (default "/home/ljf/repositories/github/openbiox/bget/_log")
  -m, --mirror string            Set the mirror of resources.
  -o, --outdir string            Set the download dir. (default "/home/ljf/repositories/github/openbiox/bget")
  -f, --overwrite                Logical indicating that whether to overwrite existing files.
      --proxy string             HTTP proxy to download.
  -q, --quiet                    No output.
  -n, --remote-name              Use remote defined filename.
  -r, --retries int              Retry specifies the number of attempts to retrieve the data. (default 5)
      --retries-sleep-time int   Sleep time after one retry. (default 5)
      --save-log                 Save download log to local file]. (default true)
  -s, --seperator string         Optional 'url1{seperator}url2' for multiple keys, urls, or seqs. (default ",")
      --task-id string           Task ID (random). (default "wpyk8931z450351")
  -t, --thread int               Concurrency download thread. (default 1)
      --thread-axel int          Set the thread of axel. (default 5)
      --timeout int              Set the timeout of per request. (default 35)
  -u, --uncompress               Uncompress download files for .zip, .tar.gz, and .gz suffix files.
      --version                  version for bget

Use "bget [command] --help" for more information about a command.
```

You can use DOI to download article and its supplementary Data. The supported website and journals will be continue increased.

**Warn**: If you do not follow the policies of the relevant website (i.e. continuous download or limited copyright), you will lose the authorization to use this tool.

```bash
Can be used to access files via DOI. More see here https://github.com/openbiox/bget.

Usage:
  bget doi [doi1 doi2 doi3...] [flags]

Examples:
  bget doi 10.5281/zenodo.3363060 10.5281/zenodo.3357455 10.5281/zenodo.3351812 -t 3
  bget doi 10.1016/j.devcel.2017.03.001 10.1016/j.stem.2019.07.009 10.1016/j.celrep.2018.03.072 -t 2

  bapi ncbi -q '((The PARK10 gene USP24 is a negative regulator of autophagy and ULK1 protein stability[Title]) OR Coordinate regulation of autophagy and the ubiquitin proteasome system by MTOR[Title])' -o titleSearch.XML
  dois=`bapi ncbi --xml2json pubmed titleSearch.XML |grep Doi| tr -d ' ,(Doi:)"'`
  echo ${dois}
  bget doi ${dois}
  bget doi 10.1080/15548627.2018.1505155 --proxy http://username:password@hostname:port

Flags:
  -g, --engine string      Point the download engine: go-http, wget, curl, axel, git, and rsync. (default "go-http")
      --full-text          Access full text. (default true)
  -h, --help               help for doi
  -l, --list-file string   A file contains dois for download.
  -m, --mirror string      Set the mirror of resources.
      --pmc                Try PMC database.
      --suppl              Access supplementary files.
      --thread-axel int    Set the thread of axel. (default 5)
```

`bget seq` can be used to access [Gene Expression Omnibus (GEO)](https://www.ncbi.nlm.nih.gov/geo), [Sequence Read Archive (SRA)](https://www.ncbi.nlm.nih.gov/sra/), and [GDC Data Portal](https://portal.gdc.cancer.gov/) are supported.

```bash
Can be used to access sequence data via unique id or manifest files. More see here https://github.com/openbiox/bget.

Usage:
  bget seq [id1 id2 id3... | manifest1 manifest2 manifest3...] [flags]

Examples:
  bget seq ERR3324530 SRR544879 # download files from SRA databaes
  bget seq GSE23543 # download files from GEO databaes (auto download SRA acc list and run info)
  bget dbgap.krt # download files from dbGap database using krt files

  # download TCGA files using file id
  bget seq b7670817-9d6b-494e-9e22-8494e2fd430d

  # download TCGA files using manifest files
  # split for parallel
  split -a 3 --additional-suffix=.txt -l 100 gdc_manifest.2019-08-23-TCGA.txt -d
  for i in x*.txt
  do
    head -n 1 x000.txt > ${i}.tmp
    cat ${i} >> ${i}.tmp
    mv ${i}.tmp ${i}
  done
  sed -i '1d' x000.txt
  bget seq *.txt -t 5

  # support auto (if you do not have *.krt, TCGA manifest, please not include it for test)
  bget seq SRR544879 GSE23543 b7670817-9d6b-494e-9e22-8494e2fd430d dbgap.krt *.txt -t 5

Flags:
  -g, --engine string      Point the download engine: go-http, wget, curl, and axel. (default "go-http")
      --gdc-token string   Token to access TCGA portal files.
  -h, --help               help for seq
  -l, --list-file string   A file contains seq id (e.g. SRR) or manifest files for download.
  -u, --uncompress         Uncompress download files for .zip, .tar.gz, and .gz suffix files (now support GEO database).
```

`bget url` can be used to access files via input URLs. Golang http, wget, curl, axel and git, and rsync are support for download process. 
```bash
Can be used to access URLs via Golang http, wget, curl, axel and git, and rsync. More see here https://github.com/openbiox/bget.

Usage:
  bget url [url1 url2 url3...] [flags]

Examples:
  urls="https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe,http://download.oray.com/pgy/windows/PgyVPN_4.1.0.21693.exe,https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe" && echo $urls | tr "," "\n"> /tmp/urls.list

  bget url ${urls}
  bget url https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe
  bget url ${urls} -t 2 -o /tmp/download
  bget url ${urls} -t 3 -o /tmp/download -f -g wget
  bget url ${urls} -t 3 -o /tmp/download -g wget --ignore
  bget url -l /tmp/urls.list -o /tmp/download -f -t 3

Flags:
  -g, --engine string      Point the download engine: go-http, wget, curl, axel, git, and rsync. (default "go-http")
  -h, --help               help for url
  -l, --list-file string   A file contains URLs for download.
  -m, --mirror string      Set the mirror of resources.
      --thread-axel int    Set the thread of axel. (default 5)
  -u, --uncompress         Uncompress download files for .zip, .tar.gz, and .gz suffix files.
```

## Maintainer

- [@Jianfeng](https://github.com/Miachol)

## License

Apache 2.0
