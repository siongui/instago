package instago

import (
	"errors"
	"fmt"
)

func PrintTaggedUsers(tu TaggedUsers) {
	fmt.Println("tagged users:", tu.GetTaggedUsernames())
	for _, pair := range tu.GetIdUsernamePairs() {
		fmt.Println("id:", pair[0], ", username:", pair[1])
	}
}

type TaggedUsers interface {
	GetTaggedUsernames() []string
	GetIdUsernamePairs() [][2]string
}

func PrintPostItem(pi PostItem) (err error) {
	fmt.Println("self type:", pi.GetSelfType())
	fmt.Println("username:", pi.GetUsername())
	fmt.Println("id:", pi.GetUserId())
	fmt.Println("url:", pi.GetPostUrl())
	fmt.Println("code:", pi.GetPostCode())
	fmt.Println("timestamp:", pi.GetTimestamp())
	urls, err := pi.GetMediaUrls()
	if err != nil {
		return
	}
	for _, url := range urls {
		fmt.Println("url:", url)
	}
	return
}

// Common methods of IG posts for IGMedia and IGItem
type PostItem interface {
	GetSelfType() string
	GetUsername() string
	GetUserId() string
	GetPostUrl() string
	GetPostCode() string
	GetTimestamp() int64
	GetMediaUrls() ([]string, error)
}

// Main data structure returned by https://www.instagram.com/p/{{CODE}}/?__a=1
type IGMedia struct {
	Typename   string `json:"__typename"`
	Id         string `json:"id"`
	Shortcode  string `json:"shortcode"`
	Dimensions struct {
		Height int64 `json:"height"`
		Width  int64 `json:"width"`
	} `json:"dimensions"`

	// "gating_info"

	MediaPreview     string `json:"media_preview"`
	DisplayUrl       string `json:"display_url"`
	DisplayResources []struct {
		Src          string `json:"src"`
		ConfigWidth  int64  `json:"config_width"`
		ConfigHeight int64  `json:"config_height"`
	} `json:"display_resources"`

	DashInfo struct {
		IsDashEligible bool `json:"is_dash_eligible"`
		//"video_dash_manifest"
		NumberOfQualities int64 `json:"number_of_qualities"`
	} `json:"dash_info"`

	VideoUrl       string `json:"video_url"`
	VideoViewCount int64  `json:"video_view_count"`
	IsVideo        bool   `json:"is_video"`

	ShouldLogClientEvent bool   `json:"should_log_client_event"`
	TrackingToken        string `json:"tracking_token"`

	EdgeMediaToTaggedUser MediaUsertags `json:"edge_media_to_tagged_user"`

	EdgeMediaToCaption struct {
		Edges []struct {
			Node struct {
				Text string `json:"text"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"edge_media_to_caption"`
	CaptionIsEdited bool `json:"caption_is_edited"`

	EdgeMediaToComment struct {
		Count    int64 `json:"count"`
		PageInfo struct {
			HasNextPage bool   `json:"has_next_page"`
			EndCursor   string `json:"end_cursor"`
		} `json:"page_info"`
		Edges []struct {
			Node struct {
				Id        string `json:"id"`
				Text      string `json:"text"`
				CreatedAt int64  `json:"created_at"`
				Owner     struct {
					Id            string `json:"id"`
					ProfilePicUrl string `json:"profile_pic_url"`
					Username      string `json:"username"`
				} `json:"owner"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"edge_media_to_comment"`
	CommentsDisabled bool `json:"comments_disabled"`

	TakenAtTimestamp int64 `json:"taken_at_timestamp"`

	EdgeMediaPreviewLike struct {
		Count int64 `json:"count"`
		Edges []struct {
			Node struct {
				Id            string `json:"id"`
				ProfilePicUrl string `json:"profile_pic_url"`
				Username      string `json:"username"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"edge_media_preview_like"`

	//"edge_media_to_sponsor_user"

	Location struct {
		Id            string `json:"id"`
		HasPublicPage bool   `json:"has_public_page"`
		Name          string `json:"name"`
		Slug          string `json:"slug"`
	} `json:"location"`

	ViewerHasLiked             bool `json:"viewer_has_liked"`
	ViewerHasSaved             bool `json:"viewer_has_saved"`
	ViewerHasSavedToCollection bool `json:"viewer_has_saved_to_collection"`

	Owner struct {
		Id                string `json:"id"`
		ProfilePicUrl     string `json:"profile_pic_url"`
		Username          string `json:"username"`
		BlockedByViewer   bool   `json:"blocked_by_viewer"`
		FollowedByViewer  bool   `json:"followed_by_viewer"`
		FullName          string `json:"full_name"`
		HasBlockedViewer  bool   `json:"has_blocked_viewer"`
		IsPrivate         bool   `json:"is_private"`
		IsUnpublished     bool   `json:"is_unpublished"`
		IsVerified        bool   `json:"is_verified"`
		RequestedByViewer bool   `json:"requested_by_viewer"`
	} `json:"owner"`

	IsAd bool `json:"is_ad"`

	// "edge_web_media_to_related_media"

	EdgeSidecarToChildren struct {
		Edges []struct {
			Node IGMedia `json:"node"`
		} `json:"edges"`
	} `json:"edge_sidecar_to_children"`
}

// return URL of image with best resolution
func (em *IGMedia) getImageUrl() string {
	res := em.DisplayResources
	return res[len(res)-1].Src
}

func (em *IGMedia) getVideoUrl() string {
	return em.VideoUrl
}

// Get username of the post owner
func (em *IGMedia) GetUsername() string {
	return em.Owner.Username
}

// Get id of the post owner
func (em *IGMedia) GetUserId() string {
	return em.Owner.Id
}

// Get timestamp of the post
func (em *IGMedia) GetTimestamp() int64 {
	return em.TakenAtTimestamp
}

// Get URLs of media (photos/videos) in the post
func (em *IGMedia) GetMediaUrls() (urls []string, err error) {
	switch em.Typename {
	case "GraphImage":
		urls = append(urls, em.getImageUrl())
	case "GraphVideo":
		urls = append(urls, em.getVideoUrl())
	case "GraphSidecar":
		for _, edge := range em.EdgeSidecarToChildren.Edges {
			nodeUrls, err := edge.Node.GetMediaUrls()
			if err != nil {
				return urls, err
			}
			urls = append(urls, nodeUrls...)
		}
	default:
		err = errors.New("Not Regular Media Type: " + em.Typename)
		return
	}

	/*
		for i, _ := range urls {
			urls[i], _ = StripQueryString(urls[i])
		}
	*/

	return
}

// Get URL of the post
func (em *IGMedia) GetPostUrl() string {
	return "https://www.instagram.com/p/" + em.Shortcode + "/"
}

// Return code of post
func (em *IGMedia) GetPostCode() string {
	return em.Shortcode
}

// Return self type name
func (em *IGMedia) GetSelfType() string {
	return "IGMedia"
}

// tagged users in post
type MediaUsertags struct {
	Edges []struct {
		Node struct {
			User struct {
				FullName      string `json:"full_name"`
				Id            string `json:"id"`
				IsVerified    bool   `json:"is_verified"`
				ProfilePicUrl string `json:"profile_pic_url"`
				Username      string `json:"username"`
			} `json:"user"`
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"node"`
	} `json:"edges"`
}

func (m MediaUsertags) GetTaggedUsernames() (usernames []string) {
	for _, edge := range m.Edges {
		usernames = append(usernames, edge.Node.User.Username)
	}
	return
}

func (m MediaUsertags) GetIdUsernamePairs() (pairs [][2]string) {
	for _, edge := range m.Edges {
		pair := [2]string{edge.Node.User.Id, edge.Node.User.Username}
		pairs = append(pairs, pair)
	}
	return
}
