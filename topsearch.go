package instago

// This file implements topsearch of Instagram web.

import (
	"encoding/json"
)

const urlTopsearch = `https://www.instagram.com/web/search/topsearch/?context=blended&include_reel=true&query=`

// Decode JSON data returned from Instagram topsearch API
type IGTopsearchResp struct {
	Users []struct {
		Position int64 `json:"position"`
		// json: cannot unmarshal string into Go struct field IGUser.users.user.pk of type int64
		User struct {
			Pk                         string `json:"pk"`
			Username                   string `json:"username"`
			FullName                   string `json:"full_name"`
			IsPrivate                  bool   `json:"is_private"`
			ProfilePicUrl              string `json:"profile_pic_url"`
			ProfilePicId               string `json:"profile_pic_id"`
			IsVerified                 bool   `json:"is_verified"`
			HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
			MutualFollowersCount       int64  `json:"mutual_followers_count"`
			//account_badges
			SocialContext       string `json:"social_context"`
			SearchSocialContext string `json:"search_social_context"`
			UnseenCount         int64  `json:"unseen_count"`
			FriendshipStatus    struct {
				Following       bool `json:"following"`
				IsPrivate       bool `json:"is_private"`
				IncomingRequest bool `json:"incoming_request"`
				OutgoingRequest bool `json:"outgoing_request"`
				IsBestie        bool `json:"is_bestie"`
				IsRestricted    bool `json:"is_restricted"`
			} `json:"friendship_status"`
			LatestReelMedia int64 `json:"latest_reel_media"`
			Seen            int64 `json:"seen"`
		} `json:"user"`
	} `json:"users"`

	Places []struct {
		Position int64 `json:"position"`
		Place    struct {
			Location struct {
				Pk               string  `json:"pk"`
				Name             string  `json:"name"`
				Address          string  `json:"address"`
				City             string  `json:"city"`
				ShortName        string  `json:"short_name"`
				Lng              float64 `json:"lng"`
				Lat              float64 `json:"lat"`
				ExternalSource   string  `json:"external_source"`
				FacebookPlacesId int64   `json:"facebook_places_id"`
			} `json:"location"`
			Title    string `json:"title"`
			Subtitle string `json:"subtitle"`
			//media_bundles
			Slug string `json:"slug"`
		} `json:"place"`
	} `json:"places"`

	Hashtags []struct {
		Position int64 `json:"position"`
		Hashtag  struct {
			Name                 string `json:"name"`
			Id                   int64  `json:"id"`
			MediaCount           int64  `json:"media_count"`
			UseDefaultAvatar     bool   `json:"use_default_avatar"`
			SearchResultSubtitle string `json:"search_result_subtitle"`
		} `json:"hashtag"`
	} `json:"hashtags"`
	HasMore          bool   `json:"has_more"`
	RankToken        string `json:"rank_token"`
	ClearClientCache bool   `json:"clear_client_cache"`
	Status           string `json:"status"`
}

// Given a string, return the users that best matches the string. This is
// actually *topsearch* on Instagram web.
func (m *IGApiManager) Topsearch(str string) (tr IGTopsearchResp, err error) {
	url := urlTopsearch + str
	b, err := m.getHTTPResponse(url, "GET")
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &tr)
	return
}
