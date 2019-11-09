install:
	go install .

doi:
	cd doc && Rscript -e "rmarkdown::render('doi.Rmd')"

