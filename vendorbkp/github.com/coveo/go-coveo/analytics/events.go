package analytics

// ActionEvent Basic incomplete analytics event type, contains most of the information
// sent from a search ui to the analytics
type ActionEvent struct {
	Language            string                 `json:"language"`
	Device              string                 `json:"device"`
	OriginLevel1        string                 `json:"originLevel1"`
	OriginLevel2        string                 `json:"originLevel2"`
	UserAgent           string                 `json:"userAgent,omitempty"`
	CustomData          map[string]interface{} `json:"customData,omitempty"`
	Anonymous           bool                   `json:"anonymous,omitempty"`
	Username            string                 `json:"username,omitempty"`
	UserDisplayName     string                 `json:"userDisplayName,omitempty"`
	Mobile              bool                   `json:"mobile,omitempty"`
	SplitTestRunName    string                 `json:"splitTestRunName,omitempty"`
	SplitTestRunVersion string                 `json:"splitTestRunVersion,omitempty"`
	OriginLevel3        string                 `json:"originLevel3,omitempty"`
}

// SearchEvent Is a structure reprensenting a search event sent to the analytics
// It incorporate an ActionEvent and adds more fields.
type SearchEvent struct {
	*ActionEvent
	SearchQueryUID  string       `json:"searchQueryUid"`
	QueryText       string       `json:"queryText"`
	ActionCause     string       `json:"actionCause"`
	AdvancedQuery   string       `json:"advancedQuery,omitempty"`
	NumberOfResults int          `json:"numberOfResults,omitempty"`
	Contextual      bool         `json:"contextual"`
	ResponseTime    int          `json:"responseTime,omitempty"`
	QueryPipeline   string       `json:"queryPipeline,omitempty"`
	UserGroups      []string     `json:"userGroups,omitempty"`
	Results         []ResultHash `json:"results,omitempty"`
}

// ResultHash Is a type used by the analytics to describe a result that was
// returned by a query that is usually sent with a search event.
type ResultHash struct {
	DocumentURI     string `json:"documentUri"`
	DocumentURIHash string `json:"documentUriHash"`
}

// ClickEvent Is a structure reprensenting a click event sent to the analytics
// It incorporate an ActionEvent and adds more fields.
type ClickEvent struct {
	*ActionEvent
	DocumentURI      string `json:"documentUri"`
	DocumentURIHash  string `json:"documentUriHash"`
	SearchQueryUID   string `json:"searchQueryUid"`
	CollectionName   string `json:"collectionName"`
	SourceName       string `json:"sourceName"`
	DocumentPosition int    `json:"documentPosition"`
	ActionCause      string `json:"actionCause"`
	DocumentTitle    string `json:"documentTitle,omitempty"`
	DocumentURL      string `json:"documentUrl,omitempty"`
	QueryPipeline    string `json:"queryPipeline,omitempty"`
	RankingModifier  string `json:"rankingModifier,omitempty"`
}

// CustomEvent Is a structure reprensenting a custom event sent to the analytics
// It incorporate an ActionEvent and adds more fields.
type CustomEvent struct {
	*ActionEvent
	EventType          string `json:"eventType"`
	EventValue         string `json:"eventValue"`
	LastSearchQueryUID string `json:"lastSearchQueryUid,omitempty"`
}

// ViewEvent Is a structure reprensenting a view event sent to the analytics
// It incorporate an ActionEvent and adds more fields.
type ViewEvent struct {
	*ActionEvent
	PageURI      string `json:"location"`
	PageReferrer string `json:"referrer"`
	PageTitle    string `json:"title"`
}
