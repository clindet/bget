package types

type BapiClisT struct {
	TaskID       string
	Quiet        string
	SaveLog      string
	LogDir       string
	HelpFlags    bool
	Version      string
	Proxy        string
	Retries      int
	Query        string
	Format       string
	Outfn        string
	Email        string
	Thread       int
	From         int
	Size         int
	RemoteName   bool
	Timeout      int
	RetSleepTime int
	CallCor      bool
	PrettyJSON   bool
	Indent       int
	SortKeys     bool
	Extra        string
}

type NcbiClisT struct {
	NcbiDB        string
	NcbiRetmax    int
	NcbiXMLToJSON string
	NcbiXMLPaths  []string
	NcbiKeywords  string
}
