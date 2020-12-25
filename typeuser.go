package instago

import (
	"strconv"
)

type User interface {
	GetUserId() string
	GetUsername() string
	IsPublic() bool
}

type IGFollowUser struct {
	Pk                         int64  `json:"pk"` // Pk is user id
	Username                   string `json:"username"`
	FullName                   string `json:"full_name"`
	IsPrivate                  bool   `json:"is_private"`
	ProfilePicUrl              string `json:"profile_pic_url"`
	ProfilePicId               string `json:"profile_pic_id"`
	IsVerified                 bool   `json:"is_verified"`
	HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
	//account_badges
	LatestReelMedia int64 `json:"latest_reel_media"`
	//story_reel_media_ids
	IsFavorite bool `json:"is_favorite"`
}

func (u IGFollowUser) GetUserId() string {
	return strconv.FormatInt(u.Pk, 10)
}

func (u IGFollowUser) GetUsername() string {
	return u.Username
}

func (u IGFollowUser) IsPublic() bool {
	return !u.IsPrivate
}

type IGUser struct {
	Pk               int64  `json:"pk"`
	Username         string `json:"username"`
	FullName         string `json:"full_name"`
	IsPrivate        bool   `json:"is_private"`
	IsVerified       bool   `json:"is_verified"`
	ProfilePicUrl    string `json:"profile_pic_url"`
	ProfilePicId     string `json:"profile_pic_id"`
	FriendshipStatus struct {
		Following       bool `json:"following"`
		FollowedBy      bool `json:"followed_by"`
		Blocking        bool `json:"blocking"`
		IsPrivate       bool `json:"is_private"`
		IncomingRequest bool `json:"incoming_request"`
		OutgoingRequest bool `json:"outgoing_request"`
		IsBestie        bool `json:"is_bestie"`
		IsRestricted    bool `json:"is_restricted"`
	} `json:"friendship_status"`
	HasAnonymousProfilePicture bool `json:"has_anonymous_profile_picture"`
	IsUnpublished              bool `json:"is_unpublished"`
	IsFavorite                 bool `json:"is_favorite"`

	// used in topsearch
	FollowerCount        int64   `json:"follower_count"`
	Byline               string  `json:"byline"`
	MutualFollowersCount float64 `json:"mutual_followers_count"`
	LatestReelMedia      int64   `json:"latest_reel_media"`
}

func (u IGUser) GetUserId() string {
	return strconv.FormatInt(u.Pk, 10)
}

func (u IGUser) GetUsername() string {
	return u.Username
}

func (u IGUser) IsPublic() bool {
	return !u.IsPrivate
}

type IGTaggedUser struct {
	Username string
	Id       string
}
