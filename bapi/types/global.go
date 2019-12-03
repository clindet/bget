package types

type BapiClisT struct {
	Quiet        bool
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
}
