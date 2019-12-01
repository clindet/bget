package urlpool

import (
	"time"
)

// BitbucketObj is the struct of https://api.bitbucket.org/2.0/repositories/user/repo/refs
type BitbucketObj struct {
	Pagelen int                `json:"pagelen"`
	Values  []BitbucketTagsVal `json:"values"`
	Page    int                `json:"page"`
}

// BitbucketTagsVal is the values of BitbucketTags
type BitbucketTagsVal struct {
	Name  string `json:"name"`
	Links struct {
		Commits struct {
			Href string `json:"href"`
		} `json:"commits"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
	} `json:"links"`
	Tagger  interface{} `json:"tagger"`
	Date    interface{} `json:"date"`
	Message interface{} `json:"message"`
	Type    string      `json:"type"`
	Target  struct {
		Hash       string `json:"hash"`
		Repository struct {
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
			Type     string `json:"type"`
			Name     string `json:"name"`
			FullName string `json:"full_name"`
			UUID     string `json:"uuid"`
		} `json:"repository"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Comments struct {
				Href string `json:"href"`
			} `json:"comments"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Diff struct {
				Href string `json:"href"`
			} `json:"diff"`
			Approve struct {
				Href string `json:"href"`
			} `json:"approve"`
			Statuses struct {
				Href string `json:"href"`
			} `json:"statuses"`
		} `json:"links"`
		Author struct {
			Raw  string `json:"raw"`
			Type string `json:"type"`
			User struct {
				DisplayName string `json:"display_name"`
				UUID        string `json:"uuid"`
				Links       struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
					Avatar struct {
						Href string `json:"href"`
					} `json:"avatar"`
				} `json:"links"`
				Nickname  string `json:"nickname"`
				Type      string `json:"type"`
				AccountID string `json:"account_id"`
			} `json:"user"`
		} `json:"author"`
		Parents []struct {
			Hash  string `json:"hash"`
			Type  string `json:"type"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
			} `json:"links"`
		} `json:"parents"`
		Date    time.Time `json:"date"`
		Message string    `json:"message"`
		Type    string    `json:"type"`
	} `json:"target"`
}

type GitHubAssets []struct {
	URL      string      `json:"url"`
	ID       int         `json:"id"`
	NodeID   string      `json:"node_id"`
	Name     string      `json:"name"`
	Label    interface{} `json:"label"`
	Uploader struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"uploader"`
	ContentType        string    `json:"content_type"`
	State              string    `json:"state"`
	Size               int       `json:"size"`
	DownloadCount      int       `json:"download_count"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	BrowserDownloadURL string    `json:"browser_download_url"`
}
