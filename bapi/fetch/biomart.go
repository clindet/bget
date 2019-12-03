package fetch

type BiomartEndpoints struct {
	Martservice bool
}

const host = "https://www.ensembl.org/biomart"

// ListEnsembl returns JSON string of BioMart databases to which biomaRt
// can connect to. By default all public BioMart databases are
// displayed.
func ListEnsembl(quite bool){

}