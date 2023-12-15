package utils

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/itishrishikesh/gofeed/internal/database"
)

func StartScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("E# - Failed to get next feeds to fetch. E: %v", err)
			continue
		}
		if err != nil {
			log.Println("Error fetching feeds:", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(feed, db, wg)
		}
		wg.Wait()
	}
}

func scrapeFeed(feed database.Feed, db *database.Queries, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("E# - Failed to mark feed as fetched. E: %v", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("E# - Failed to fetch feed. E: %v", err)
		return
	}
	for _, item := range rssFeed.Channel.Items {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("E# - Failed to parse date. E: %v", err)
			continue
		}
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: description,
			PublishedAt: pubAt,
			Url:         item.Link,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("E# - Failed to create post. E: %v", err)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Items))
}
