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

Examples:
  urls="https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe,http://download.oray.com/pgy/windows/PgyVPN_4.1.0.21693.exe,https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe" && echo $urls | sed 's/,/\n/g'> /tmp/urls.list

  bget -u ${urls} -n 2 -o /tmp/download
  bget -u ${urls} -n 3 -o /tmp/download -f -g wget
  bget -u ${urls} -n 3 -o /tmp/download -g wget --ignore
  bget -l /tmp/urls.list -o /tmp/download -f -n 3
  bget --get-spack
  bget --get-miniconda 3 -o /tmp/testenv
  bget --get-miniconda 3 --engine wget
  bget --get-miniconda 3 --engine axel
  bget --get-reffa GRCh38_defuse_97 -n 10
  bget --get-reffa hg38_ucsc,GRCh37_genecode_31

Flags:
      --autopath               Logical indicating that whether to create subdir in download dir: e.g. reffa/{{site}}/{{version}} (default true)
  -g, --engine string          Point the download engine: go-http, wget, curl, axel, git, and rsync. (default "go-http")
  -e, --extra-cmd string       Extra flags and values pass to internal CMDs
      --get-miniconda string   Install miniconda2 or miniconda3 in tools directory. Optional (2 or 3).
      --get-reffa string       Download reference in download directory. Format is genomeVersion_site_releaseVersion.
                               Optional (GRCh38_genecode_31, GRCh37_genecode_31, hg38_ucsc, hg19_ucsc, GRCh38_ensemble_97, GRCh38_defuse_97).
                               Multiple use comma to seperate (e.g. hg38_ucsc,hg19_ucsc).
      --get-spack              Logical indicating that whether to install spack in tools directory.
  -u, --get-urls string        URLs to be download.
  -l, --get-urls-list string   A file contains URLs for download.
  -h, --help                   help for bget
      --ignore                 Contine to download and skip the check of existed files.
      --log-dir string         Log dir. (default "/home/ljf/repositories/github/JhuangLab/bget/_log")
  -m, --mirror string          Set the mirror of resources.
  -o, --outdir string          Set the download dir for get-urls. (default "/home/ljf/repositories/github/JhuangLab/bget/_download")
  -f, --overwrite              Logical indicating that whether to overwrite existing files.
  -q, --quiet                  No output.
      --save-log               Save download log to local file]. (default true)
  -s, --separator string       Separator for --get-reffa and -u. (default ",")
  -k, --task-id string         Task ID. (default "bxmzndet179wx9k")
  -n, --thread int             Concurrency download thread. (default 1)
      --thread-axel int        Set the thread of axel. (default 5)
      --version                version for bget
```

## Maintainer

- [@Jianfeng](https://github.com/Miachol)

## License

Apache 2.0
