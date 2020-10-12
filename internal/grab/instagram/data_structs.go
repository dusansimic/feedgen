package instagram

type DataJSON struct {
	EntryData struct {
		ProfilePage []struct {
			GraphQL struct {
				User struct {
					ID         string `json:"id"`
					Username   string `json:"username"`
					FullName   string `json:"full_name"`
					Biography  string `json:"biography"`
					ProfileURL string `json:"profile_pic_url_hd"`
					Media      struct {
						Edges []struct {
							Node     mediaNode `json:"node"`
							PageInfo pageInfo  `json:"page_info"`
						} `json:"edges"`
					} `json:"edge_owner_to_timeline_media"`
				} `json:"user"`
			} `json:"graphql"`
		} `json:"ProfilePage"`
	} `json:"entry_data"`
}

type pageInfo struct {
	EndCursor string `json:"end_cursor"`
	NextPage  bool   `json:"has_next_page"`
}

type caption struct {
	Edges []struct {
		Node struct {
			Text string `json:"text"`
		} `json:"node"`
	} `json:"edges"`
}

type dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type mediaNode struct {
	ImageURL     string     `json:"display_url"`
	ThumbnailURL string     `json:"thumbnail_src"`
	Shortcode    string     `json:"shortcode"`
	IsVideo      bool       `json:"is_video"`
	Date         int        `json:"taken_at_timestamp"`
	Dimensions   dimensions `json:"dimensions"`
	Caption      caption    `json:"edge_media_to_caption"`
	Children     struct {
		Edges []struct {
			Node childNode `json:"node"`
		} `json:"edges"`
	} `json:"edge_sidecar_to_children"`
}

type childNode struct {
	Shortcode  string     `json:"shortcode"`
	Dimensions dimensions `json:"dimensions"`
}
