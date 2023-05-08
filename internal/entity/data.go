package entity

type InstagramData struct {
	GraphQL struct {
		User struct {
			EdgeOwnerToTimelineMedia struct {
				Edges []Edge `json:"edges"`
			} `json:"edge_owner_to_timeline_media"`
		} `json:"user"`
	} `json:"graphql"`
}

type Edge struct {
	Node struct {
		ID                   string               `json:"id"`
		EdgeMediaToCaption   EdgeMediaToCaption   `json:"edge_media_to_caption"`
		ThumbnailSrc         string               `json:"thumbnail_src"`
		TakenAtTimestamp     int64                `json:"taken_at_timestamp"`
		EdgeMediaPreviewLike EdgeMediaPreviewLike `json:"edge_media_preview_like"`
	} `json:"node"`
}

type EdgeMediaToCaption struct {
	Edges []EdgeCaption `json:"edges"`
}

type EdgeCaption struct {
	Node struct {
		Text string `json:"text"`
	} `json:"node"`
}

type EdgeMediaPreviewLike struct {
	Count int `json:"count"`
}
