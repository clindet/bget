<img src="https://img.shields.io/badge/lifecycle-experimental-orange.svg" alt="Life cycle: experimental"> [![GoDoc](https://godoc.org/github.com/JhuangLab/bget?status.svg)](https://godoc.org/github.com/JhuangLab/bget)

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
go get -u github.com/JhuangLab/bget
```

## Usage

```bash
Lightweight downloader for bioinformatics data, databases and files (under development). It will provides a simple and parallelized method to access various bioinformatics resoures. More see here https://github.com/JhuangLab/bget.

Usage:
  bget [flags]
  bget [command]

Available Commands:
  doi         Can be used to access files via DOI.
  help        Help about any command
  seq         Can be used to access sequence data via unique id (dbGAP and EGA) or manifest files (TCGA).
  url         Can be used to access URLs via Golang http, wget, curl, axel and git, and rsync.
  url-key     Can be used to access URLs via a key string.

Flags:
      --clean              Remove _download and _log in current dir.
  -e, --extra-cmd string   Extra flags and values pass to internal CMDs
  -h, --help               help for bget
      --ignore             Contine to download and skip the check of existed files.
      --log-dir string     Log dir. (default "/Users/apple/Documents/bget/_log")
  -o, --outdir string      Set the download dir for get-urls. (default "/Users/apple/Documents/bget/_download")
  -f, --overwrite          Logical indicating that whether to overwrite existing files.
  -q, --quiet              No output.
      --save-log           Save download log to local file]. (default true)
  -s, --separator string   Separator for --reffa,-k, and -u flag. (default ",")
      --task-id string     Task ID (random). (default "9iscp3s1s8sg6f1")
  -t, --thread int         Concurrency download thread. (default 1)
      --version            version for bget

Use "bget [command] --help" for more information about a command.
```

You can also to use DOI download article and its supplementary Data. The supported website and journals will be continue increased.

```bash
bget doi 10.5281/zenodo.3363060 10.5281/zenodo.3357455 10.5281/zenodo.3351812 -t 3
bget doi 10.1016/j.chembiol.2017.08.011 10.1016/j.ccell.2016.11.002 10.1016/j.cell.2017.07.016 -t 2
bget doi 10.1016/j.molcel.2017.10.018 10.1126/sciadv.aax3387 10.1016/j.neuron.2017.09.008 -t 2
```

`bget seq` can be used to access [Gene Expression Omnibus (GEO)](https://www.ncbi.nlm.nih.gov/geo), [Sequence Read Archive (SRA)](https://www.ncbi.nlm.nih.gov/sra/), and [GDC Data Portal](https://portal.gdc.cancer.gov/) are supported.

```bash
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
```

`bget url` can be used to access files via input URLs. Golang http, wget, curl, axel and git, and rsync are support for download process. 

```bash
urls="https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe,http://download.oray.com/pgy/windows/PgyVPN_4.1.0.21693.exe,https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe" && echo $urls | tr "," "\n"> /tmp/urls.list

bget url ${urls}
bget url https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe
bget url ${urls} -t 2 -o /tmp/download
bget url ${urls} -t 3 -o /tmp/download -f -g wget
bget url ${urls} -t 3 -o /tmp/download -g wget --ignore
bget url -l /tmp/urls.list -o /tmp/download -f -t 3
```

`bget url-key` can be used to access files via using a key string. key-value pairs for URLs.

```bash
# Get all supported key
bget url-key -a
# Get all versions of supported key (table format)
bget url-key gdc-client bwa -v table
# Get all versions of supported key (json format)
bget url-key gdc-client bwa -v json
# Get all versions of supported key (text format)
bget url-key gdc-client bwa -v txt

# Download files (2 thread)
bget url-key gdc-client bwa -t 2
```

## Maintainer

- [@Jianfeng](https://github.com/Miachol)

## License

Apache 2.0
