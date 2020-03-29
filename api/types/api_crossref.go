package types

type CrossRefEndpoints struct {
	Doi          CrossRefDoiPost
	ArticleTitle CrossRefTitlePost
	MetaData     CrossRefMetaPost
	SearchType   string `json:"search_type"`
	QueryType    string `json:"queryType"`
}

type CrossRefMetaPost struct {
	Auth        string `json:"auth"`
	ISSN        string `json:"issn"`
	Title       string `json:"title"`
	Atitle      string `json:"atitle"`
	Volume      string `json:"volume"`
	Issue       string `json:"issue"`
	Page        string `json:"page"`
	Year        string `json:"year"`
	ISBN        string `json:"isbn"`
	Compum      string `json:"compnum"`
	Stitle      string `json:"stitle"`
	ViewRecords string `json:"view_records"`
	SearchType  string `json:"search_type"`
	QueryType   string `json:"queryType"`
}

type CrossRefTitlePost struct {
	ArticleTitleSearch string `json:"article_title_search"`
	MultiHit           string `json:"multi_hit"`
	Atitle2            string `json:"atitle2"`
	Auth2              string `json:"auth2"`
	QueryType          string `json:"queryType"`
}

type CrossRefDoiPost struct {
	Doi       string `json:"doi"`
	QueryType string `json:"queryType"`
	ResType   string `json:"resType"`
	DoiSearch string `json:"doi_search"`
}
