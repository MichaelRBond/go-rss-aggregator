package types

// FeverAPIPostRequest fevers post body for access
type FeverAPIPostRequest struct {
	APIKey string `json:"api_key"`
}

// FeverAPIQueryParams allowed query params for fever api. Most of these will be
// empty. Their presents changes behavior
type FeverAPIQueryParams struct {
	API           string `json:"api"` // Returns a 501 if set to xml
	Groups        string `json:"groups"`
	Feeds         string `json:"feeds"`
	Favicons      string `json:"favicons"`
	Items         string `json:"items"`
	Links         string `json:"links"`
	UnreadItemIDs string `json:"uread_item_ids"`
	SavedItemIDs  string `json:"saved_item_ids"`
}

// FeverAPIOptions denotes options passed in with FeverAPIQueryParams
type FeverAPIOptions struct {
	API           string
	Groups        bool
	Feeds         bool
	Favicons      bool
	Items         bool
	Links         bool
	UnreadItemIDs bool
	SavedItemIDs  bool
}

// FeverFeedGroups is the Fever API return for feed Groups.
// -- FeedIDs : A string of FeedIds that belong too the group, comma delmited
type FeverFeedGroups struct {
	GroupID int32  `json:"group_id"`
	FeedIDs string `json:"feed_ids"`
}

// FeverGroup a FeverAPI group object
type FeverGroup struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
}
