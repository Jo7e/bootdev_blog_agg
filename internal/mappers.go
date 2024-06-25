package internal

import (
	"bootdev_blog_agg/internal/database"
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
}

func DatabaseFeedToFeed(f database.Feed) Feed {
	return Feed{
		ID:        f.ID,
		Name:      f.Name,
		Url:       f.Url,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		UserID:    f.UserID,
	}
}

func DataBaseFeedsToFeeds(feeds []database.Feed) []Feed {
	feedsResponse := make([]Feed, 0, len(feeds))

	for _, feed := range feeds {
		feedResponse := Feed{
			ID:        feed.ID,
			Name:      feed.Name,
			Url:       feed.Url,
			CreatedAt: feed.CreatedAt,
			UpdatedAt: feed.UpdatedAt,
			UserID:    feed.UserID,
		}
		feedsResponse = append(feedsResponse, feedResponse)
	}

	return feedsResponse
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func DatabaseFeedFollowToFeedFollow(feedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feedFollow.ID,
		FeedID:    feedFollow.FeedID,
		UserID:    feedFollow.UserID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
	}
}

func DatabaseFeedFollowsToFeedFollows(feedFollows []database.FeedFollow) []FeedFollow {
	feedFollowsResponse := make([]FeedFollow, 0, len(feedFollows))

	for _, feedFollow := range feedFollows {
		feedFollowResponse := FeedFollow{
			ID:        feedFollow.ID,
			FeedID:    feedFollow.FeedID,
			UserID:    feedFollow.UserID,
			CreatedAt: feedFollow.CreatedAt,
			UpdatedAt: feedFollow.UpdatedAt,
		}
		feedFollowsResponse = append(feedFollowsResponse, feedFollowResponse)
	}

	return feedFollowsResponse
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Apikey    string    `json:"apikey"`
}

func DatabaseUserToUser(u database.User) User {
	return User{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Apikey:    u.Apikey,
	}
}
