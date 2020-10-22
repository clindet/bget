#!/bin/bash
urls="https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe,http://download.oray.com/pgy/windows/PgyVPN_4.1.0.21693.exe,https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe" && echo $urls | tr "," "\n"> /tmp/urls.list

bget url ${urls}
bget url https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe --save-log
bget url ${urls} -t 3 -o /tmp/download -f -g wget --save-log --verbose 2
bget url ${urls} -t 2 -o /tmp/download --save-log --verbose 2
bget url ${urls} -t 3 -o /tmp/download -g wget --resume
bget url -l /tmp/urls.list -o /tmp/download -f -t 3

bget url Miachol/github_demo --github
bget url PapenfussLab/gridss clindet/bget --with-github-assets -t 5 --github
bget url PapenfussLab/gridss clindet/bget --only-github-assets -t 5 --github
bget url PapenfussLab/gridss clindet/bget --with-github-assets --with-assets-versions v2.7.2,v0.1.3 -t 5 --github

bget --clean

