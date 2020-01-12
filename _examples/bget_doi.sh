#!/bin/bash
bget api ncbi -q '((The PARK10 gene USP24 is a negative regulator of autophagy and ULK1 protein stability[Title]) OR Coordinate regulation of autophagy and the ubiquitin proteasome system by MTOR[Title])' -o titleSearch.XML
dois=`bget api ncbi --xml2json pubmed titleSearch.XML |grep Doi| tr -d ' ,(Doi:)"'`
rm titleSearch.XML
echo ${dois}
bget doi ${dois}

bget --clean

