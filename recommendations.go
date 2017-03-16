package core

import (
	"fmt"
	"sort"
	"strings"
)

// Recommendations is a collection of music recommendation for a user.
type Recommendations struct {

	// A List of recommended music for the user.
	List MusicList

	// ID of the user who these recommendations are for.
	UserID string
}

// NewRecommendations returns a list of new recommendations to the specified user.
func NewRecommendations(userID string) *Recommendations {
	return &Recommendations{UserID: userID}
}

// RecommendSort picks a list of recommended music from the system sorted by the music ID.
func RecommendSort(userID string, maxCount int, store Store) (*Recommendations, error) {
	r, err := Recommend(userID, maxCount, store)
	if err != nil {
		return nil, err
	}

	sort.Slice(r.List, func(i, j int) bool {
		return strings.Compare(r.List[i].ID, r.List[j].ID) == -1
	})

	return r, nil
}

// Recommend picks a list of recommended music for the specified user.
func Recommend(userID string, maxCount int, store Store) (*Recommendations, error) {
	user, err := store.FindUser(userID)
	if err != nil {
		return nil, err
	}

	r := NewRecommendations(userID)
	if user.IsNew() {
		recommendations, err := store.ListMusic(maxCount)
		if err != nil {
			return nil, err
		}
		r.List = recommendations
		return r, nil
	}

	doneFolloweesTask, doneHistoryTask := make(chan struct{}), make(chan struct{})
	musicChan := make(chan *Music)
	exclude := buildExcludeList(user)

	go func() {
		defer close(doneFolloweesTask)
		for _, f := range user.Followees {
			for _, m := range f.History {
				musicChan <- m
			}
		}
	}()

	go func() {
		defer close(doneHistoryTask)
		for _, history := range user.History {
			for _, tag := range history.Tags {
				musicByTags, err := store.FindMusicByTags(tag)
				if err != nil {
				}
				for _, music := range musicByTags {
					musicChan <- music
				}
			}
		}
	}()

	// Exit loops only when both goroutines are done.
	// Refer https://dave.cheney.net/2013/04/30/curious-channels for info on nil channel's behaviour.
	for doneFolloweesTask != nil || doneHistoryTask != nil {
		select {
		case m := <-musicChan:
			if _, exists := exclude[m.ID]; !exists {
				exclude[m.ID] = struct{}{}
				r.List = append(r.List, m)
			}
		case <-doneFolloweesTask:
			doneFolloweesTask = nil
		case <-doneHistoryTask:
			doneHistoryTask = nil
		}
	}

	return r, nil
}

func buildExcludeList(u *User) map[string]struct{} {
	// initialize the 'exclude' map with the user's music history to avoid adding duplicates later
	exclude := make(map[string]struct{})
	for _, h := range u.History {
		exclude[h.ID] = struct{}{}
	}
	return exclude
}

// String returns the string representation of the recommendations.
func (r Recommendations) String() string {
	return fmt.Sprintf("%T {\nUser ID:%q\nList: [%s]\n}", r, r.UserID, r.List)
}
