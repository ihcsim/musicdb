package main

import (
	"fmt"

	"github.com/ihcsim/bluelens/cmd/blued/app"
	"github.com/ihcsim/bluelens/internal/core"
)

func mediaTypeRecommendations(r *core.Recommendations) *app.BluelensRecommendations {
	musicIDs := []string{}
	musicLinks := app.BluelensMusicLinkCollection{}
	for _, m := range r.List {
		musicIDs = append(musicIDs, m.ID)
		link := &app.BluelensMusicLink{Href: fmt.Sprintf("/music/%s", m.ID)}
		musicLinks = append(musicLinks, link)
	}

	links := &app.BluelensRecommendationsLinks{
		List: musicLinks,
		User: &app.BluelensUserLink{Href: fmt.Sprintf("/users/%s", r.UserID)},
	}
	return &app.BluelensRecommendations{
		MusicID: musicIDs,
		Links:   links,
	}
}

func mediaTypeMusic(m *core.Music) *app.BluelensMusic {
	return &app.BluelensMusic{
		ID:   m.ID,
		Href: fmt.Sprintf("/music/%s", m.ID),
		Tags: m.Tags,
	}
}

func mediaTypeMusicFull(m *core.Music) *app.BluelensMusicFull {
	return &app.BluelensMusicFull{
		ID:   m.ID,
		Href: fmt.Sprintf("/music/%s", m.ID),
		Tags: m.Tags,
	}
}

func mediaTypeMusicLink(m *core.Music) *app.BluelensMusicLink {
	return &app.BluelensMusicLink{
		Href: fmt.Sprintf("/music/%s", m.ID),
	}
}

func mediaTypeUser(u *core.User) *app.BluelensUser {
	followeesLinks := app.BluelensUserLinkCollection{}
	for _, followee := range u.Followees {
		link := &app.BluelensUserLink{
			Href: fmt.Sprintf("/users/%s", followee.ID),
		}
		followeesLinks = append(followeesLinks, link)
	}

	historyLinks := app.BluelensMusicLinkCollection{}
	for _, music := range u.History {
		link := &app.BluelensMusicLink{
			Href: fmt.Sprintf("/music/%s", music.ID),
		}
		historyLinks = append(historyLinks, link)
	}

	links := &app.BluelensUserLinks{
		Followees: followeesLinks,
		History:   historyLinks,
	}

	return &app.BluelensUser{
		ID:    u.ID,
		Href:  fmt.Sprintf("/users/%s", u.ID),
		Links: links,
	}
}

func mediaTypeUserFull(u *core.User) *app.BluelensUserFull {
	followees := app.BluelensUserCollection{}
	for _, followee := range u.Followees {
		user := mediaTypeUser(followee)
		followees = append(followees, user)
	}

	history := app.BluelensMusicCollection{}
	for _, music := range u.History {
		music := mediaTypeMusic(music)
		history = append(history, music)
	}

	return &app.BluelensUserFull{
		ID:        u.ID,
		Followees: followees,
		History:   history,
		Href:      fmt.Sprintf("/users/%s", u.ID),
	}
}

func mediaTypeUserLink(u *core.User) *app.BluelensUserLink {
	return &app.BluelensUserLink{
		Href: fmt.Sprintf("/users/%s", u.ID),
	}
}
