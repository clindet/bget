install:
	go install .

doi:
	cd doc && Rscript -e "rmarkdown::render('doi.Rmd')"

check_todo:
	bioctl stat `grep todo doc/doi.list.journal.txt|cut -f 5 |cut -f 1 -d /` --freq | sort -k 1rn

