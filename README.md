<img src="https://img.shields.io/badge/lifecycle-experimental-orange.svg" alt="Life cycle: experimental"> [![GoDoc](https://godoc.org/github.com/JhuangLab/bget?status.svg)](https://godoc.org/github.com/JhuangLab/bget)

# bget

Lightweight downloader for bioinformatics data, databases and files. Golang `http` library, `wget`, `curl`, `axel`, `git`, and `rsync` were supported as the download engine.

[![asciicast](https://asciinema.org/a/262357.svg)](https://asciinema.org/a/262357)

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
  bget [url1 url2... | -k key1 key2... | --doi doi1 doi2...] [flags]

Examples:
  urls="https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe,http://download.oray.com/pgy/windows/PgyVPN_4.1.0.21693.exe,https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe" && echo $urls | tr "," "\n"> /tmp/urls.list

  bget ${urls}
  bget https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe
  bget -u ${urls} -t 2 -o /tmp/download
  bget -u ${urls} -t 3 -o /tmp/download -f -g wget
  bget -u ${urls} -t 3 -o /tmp/download -g wget --ignore
  bget -l /tmp/urls.list -o /tmp/download -f -t 3
  bget -k bwa
  bget --doi 10.5281/zenodo.3363060 10.5281/zenodo.3357455 10.5281/zenodo.3351812 -t 3
  bget --spack
  bget --miniconda 3 -o /tmp/testenv
  bget --miniconda 3 --engine wget
  bget --miniconda 3 --engine axel
  bget -k "reffa@GRCh38 %defuse #97" -t 10 -f
  bget --reffa "GRCh38 %defuse #97" -t 10
  bget --reffa "hg38 %ucsc, GRCh37 %genecode #31"

Flags:
      --autopath           Logical indicating that whether to create subdir in download dir (for --reffa): e.g. reffa/{{site}}/{{version}} (default true)
      --doi string         Doi to be download.
  -g, --engine string      Point the download engine: go-http, wget, curl, axel, git, and rsync. (default "go-http")
  -e, --extra-cmd string   Extra flags and values pass to internal CMDs
  -h, --help               help for bget
      --ignore             Contine to download and skip the check of existed files.
  -k, --keys string        String key to be download. item@version %site #releaseVersion, e.g. bwa, GRCh38 %defuse #97
  -a, --keys-all           Show all available string key can be download.
      --log-dir string     Log dir. (default "/home/ljf/repositories/github/JhuangLab/bget/_log")
      --miniconda string   Install miniconda2 or miniconda3 in tools directory. Optional (2 or 3).
  -m, --mirror string      Set the mirror of resources.
  -o, --outdir string      Set the download dir for get-urls. (default "/home/ljf/repositories/github/JhuangLab/bget/_download")
  -f, --overwrite          Logical indicating that whether to overwrite existing files.
  -q, --quiet              No output.
      --reffa string       Download reference in download directory. Format is genomeVersion %site #releaseVersion.
                           Optional (GRCh38 %genecode #31, GRCh37 %genecode #31, hg38 %ucsc, hg19 %ucsc, GRCh38 %ensemble #97, GRCh38 %defuse #97).
                           Multiple use comma to seperate (e.g. GRCh38 %genecode #31,GRCh37 %genecode #31, %fusioncatcher #95).
      --save-log           Save download log to local file]. (default true)
  -s, --separator string   Separator for --reffa,-k, and -u flag. (default ",")
      --spack              Logical indicating that whether to install spack in tools directory.
      --task-id string     Task ID (random). (default "gp82no1hz9kygmk")
  -t, --thread int         Concurrency download thread. (default 1)
      --thread-axel int    Set the thread of axel. (default 5)
  -u, --urls string        URLs to be download.
  -l, --urls-list string   A file contains URLs for download.
      --version            version for bget
```

You can also to use DOI download article and its supplementary Data. The supported website and journals will be continue increased.

```bash
bget --doi 10.5281/zenodo.3363060 10.5281/zenodo.3357455 10.5281/zenodo.3351812 -t 3
bget --doi 10.1016/j.chembiol.2017.08.011 10.1016/j.ccell.2016.11.002 10.1016/j.cell.2017.07.016 -t 2
bget --doi 10.1016/j.molcel.2017.10.018 10.1126/sciadv.aax3387 10.1016/j.neuron.2017.09.008 -t 2
```

## Maintainer

- [@Jianfeng](https://github.com/Miachol)

## License

Apache 2.0
