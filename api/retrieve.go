package api

// RetrieveOption is the options for retrieve API.
type RetrieveOption struct {
	State       State          `json:"state,omitempty"`
	Favorite    FavoriteFilter `json:"favorite,omitempty"`
	Tag         string         `json:"tag,omitempty"`
	ContentType ContentType    `json:"contentType,omitempty"`
	Sort        Sort           `json:"sort,omitempty"`
	DetailType  DetailType     `json:"detailType,omitempty"`
	Search      string         `json:"search,omitempty"`
	Domain      string         `json:"domain,omitempty"`
	Since       int            `json:"since,omitempty"`
	Count       int            `json:"count,omitempty"`
	Offset      int            `json:"offset,omitempty"`
}

type State string

const (
	StateUnread  State = "unread"
	StateArchive       = "archive"
	StateAll           = "all"
)

type ContentType string

const (
	ContentTypeArticle ContentType = "article"
	ContentTypeVideo               = "video"
	ContentTypeImage               = "image"
)

type Sort string

const (
	SortNewest Sort = "newest"
	SortOldest      = "oldest"
	SortTitle       = "title"
	SortSite        = "site"
)

type DetailType string

const (
	DetailTypeSimple   DetailType = "simple"
	DetailTypeComplete            = "complete"
)

type FavoriteFilter string

const (
	FavoriteFilterUnspecified FavoriteFilter = ""
	FavoriteFilterUnfavorited                = "0"
	FavoriteFilterFavorited                  = "1"
)

type retrieveAPIOptionWithAuth struct {
	*RetrieveOption
	authInfo
}

type RetrieveResult struct {
	List     map[string]Item
	Status   int
	Complete int
	Since    int
}

type ItemStatus int

const (
	ItemStatusUnread   ItemStatus = 0
	ItemStatusArchived            = 1
	ItemStatusDeleted             = 2
)

type ItemMediaAttachment int

const (
	ItemMediaAttachmentNoMedia  ItemMediaAttachment = 0
	ItemMediaAttachmentHasMedia                     = 1
	ItemMediaAttachmentIsMedia                      = 2
)

type Item struct {
	ItemID        int        `json:"item_id,string"`
	ResolvedId    int        `json:"resolved_id,string"`
	GivenURL      string     `json:"given_url"`
	ResolvedURL   string     `json:"resolved_url"`
	GivenTitle    string     `json:"given_title"`
	ResolvedTitle string     `json:"resolved_title"`
	Favorite      int        `json:",string"`
	Status        ItemStatus `json:",string"`
	Excerpt       string
	IsArticle     int                 `json:"is_article,string"`
	HasImage      ItemMediaAttachment `json:"has_image,string"`
	HasVideo      ItemMediaAttachment `json:"has_video,string"`
	WordCount     int                 `json:"word_count,string"`

	// Fields for detailed response
	Tags    map[string]map[string]interface{}
	Authors map[string]map[string]interface{}
	Images  map[string]map[string]interface{}
	Videos  map[string]map[string]interface{}

	// Fields that are not documented but exist
	SortId int `json:"sort_id"`
}

// URL is an alias for ResolvedURL
func (item Item) URL() string {
	return item.ResolvedURL
}

// Title is an alias for ResolvedTitle
func (item Item) Title() string {
	return item.ResolvedTitle
}

// Retrieve returns the in Pocket
func (c *Client) Retrieve(options *RetrieveOption) (*RetrieveResult, error) {
	data := retrieveAPIOptionWithAuth{
		authInfo:       c.authInfo,
		RetrieveOption: options,
	}

	res := &RetrieveResult{}
	err := PostJSON("/v3/get", data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
