package instago

type IGReelMediaItem struct {
	Audience              string `json:"audience"`
	EdgeStoryMediaViewers struct {
		Count    int64 `json:"count"`
		PageInfo struct {
			HasNextPage bool `json:"has_next_page"`
			//end_cursor
		} `json:"page_info"`
		//edges
	} `json:"edge_story_media_viewers"`
	Typename   string `json:"__typename"`
	Id         string `json:"id"`
	Dimensions struct {
		Height int64 `json:"height"`
		Width  int64 `json:"width"`
	} `json:"dimensions"`
	DisplayResources []struct {
		Src          string `json:"src"`
		ConfigWidth  int64  `json:"config_width"`
		ConfigHeight int64  `json:"config_height"`
	} `json:"display_resources"`
	DisplayUrl   string `json:"display_url"`
	MediaPreview string `json:"media_preview"`
	//gating_info
	//fact_check_overall_rating
	//fact_check_information
	//media_overlay_info
	//sensitivity_friction_info
	TakenAtTimestamp    int64 `json:"taken_at_timestamp"`
	ExpiringAtTimestamp int64 `json:"expiring_at_timestamp"`
	//story_cta_url
	//story_view_count
	IsVideo       bool            `json:"is_video"`
	Owner         IGReelMediaUser `json:"owner"`
	TrackingToken string          `json:"tracking_token"`

	HasAudio bool `json:"has_audio"`
	//overlay_image_resources
	VideoDuration  float64 `json:"video_duration"`
	VideoResources []struct {
		Src          string `json:"src"`
		ConfigWidth  int64  `json:"config_width"`
		ConfigHeight int64  `json:"config_height"`
		MimeType     string `json:"mime_type"`
		Profile      string `json:"profile"`
	} `json:"video_resources"`

	TappableObjects []struct {
		Typename string  `json:"__typename"`
		X        float64 `json:"x"`
		Y        float64 `json:"y"`
		Width    float64 `json:"width"`
		Height   float64 `json:"height"`
		Rotation float64 `json:"rotation"`
		//custom_title
		//attribution
		TappableType string `json:"tappable_type"`
		Username     string `json:"username"`
		FullName     string `json:"full_name"`
		IsPrivate    bool   `json:"is_private"`
	} `json:"tappable_objects"`
	//story_app_attribution
	//edge_media_to_sponsor_user
	//muting_info
}

func (i IGReelMediaItem) GetUsername() string {
	return i.Owner.Username
}

func (i IGReelMediaItem) GetUserId() string {
	return i.Owner.Id
}

func (i IGReelMediaItem) GetMediaUrl() string {
	switch i.Typename {
	case "GraphStoryImage":
		return i.GetStoryImageUrl()
	case "GraphStoryVideo":
		return i.GetStoryVideoUrl()
	default:
		return ""
	}
	return ""
}

func (i IGReelMediaItem) GetStoryImageUrl() string {
	src := ""
	width := int64(0)

	for _, dr := range i.DisplayResources {
		if dr.ConfigWidth > width {
			width = dr.ConfigWidth
			src = dr.Src
		}
	}

	return src
}

func (i IGReelMediaItem) GetStoryVideoUrl() string {
	src := ""
	width := int64(0)

	for _, vr := range i.VideoResources {
		if vr.ConfigWidth > width {
			width = vr.ConfigWidth
			src = vr.Src
		}
	}

	return src
}
