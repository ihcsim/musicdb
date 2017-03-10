package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var apiHost = "localhost:8080"

var _ = API("bluelens", func() {
	Title("The bluelens API")
	Description("This API provides a set of endpoints to manage users' followees, music history and recommendations.")
	Host(apiHost)
	Scheme("http")
	BasePath("/bluelens")
})

var _ = Resource("recommendations", func() {
	BasePath("/recommendations")
	DefaultMedia(Recommendations)

	Action("list", func() {
		Routing(GET("/:userID/:maxCount"))
		Description("List all the music recommendations for a user.")
		Params(func() {
			Param("userID", String, "ID of the user these recommendations are meant for.")
			Param("maxCount", Integer, "Maximum number of recommendations to be returned to the user.")
		})
		Response(OK)
	})
})

var _ = Resource("user", func() {
	BasePath("/user")

	Action("listens", func() {
		Routing(POST("/:userID/listen/:musicID"))
		Description("A user listens to a music.")
		Params(func() {
			Param("userID", String, "ID of the user.")
			Param("musicID", Integer, "ID of the music.")
		})
		Response(OK)
	})

	Action("follows", func() {
		Routing(POST("/:followerID/follows/:followeeID"))
		Description("A user follows another user.")
		Params(func() {
			Param("followeeID", String, "ID of the followee.")
			Param("followerID", Integer, "ID of the follower.")
		})
		Response(OK)
	})
})

var Recommendations = MediaType("application/vnd.bluelens.recommendations+json", func() {
	Description("A list of recommendations for the specified user")
	ContentType("application/json")
	Attributes(func() {
		Attribute("list", CollectionOf("application/vnd.bluelens.music+json"))
		Attribute("user", User)

		Links(func() {
			Link("list")
			Link("user")
		})

		Required("list", "user")
	})

	View("default", func() {
		Attribute("list")
		Attribute("links")
	})

	View("all", func() {
		Attribute("list")
		Attribute("user")
	})
})

var User = MediaType("application/vnd.bluelens.user+json", func() {
	Description("A user resource")
	ContentType("application/json")
	Attributes(func() {
		Attribute("id", String)
		Attribute("followees", CollectionOf("application/vnd.bluelens.user+json"))
		Attribute("history", CollectionOf("application/vnd.bluelens.music+json"))
		Attribute("href", String)

		Links(func() {
			Link("followees")
			Link("history")
		})

		Required("id", "href")
	})

	View("default", func() {
		Attribute("id")
		Attribute("links")
		Attribute("href")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})

	View("all", func() {
		Attribute("id")
		Attribute("followees")
		Attribute("history")
		Attribute("href")
	})
})

var Music = MediaType("application/vnd.bluelens.music+json", func() {
	Description("A music resource")
	ContentType("application/json")
	Attributes(func() {
		Attribute("id", String)
		Attribute("tags", ArrayOf(String))
		Attribute("href", String)

		Required("id", "href")
	})

	View("default", func() {
		Attribute("id")
		Attribute("tags")
		Attribute("href")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/swagger.json", "server/swagger/swagger.json")
})