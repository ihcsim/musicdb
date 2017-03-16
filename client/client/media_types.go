// Code generated by goagen v1.1.0, command line:
// $ goagen
// --design=github.com/ihcsim/bluelens/design
// --out=$(GOPATH)/src/github.com/ihcsim/bluelens/client
// --tool=blue
// --version=v1.1.0
//
// API "bluelens": Application Media Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// A music resource (default view)
//
// Identifier: application/vnd.bluelens.music+json; view=default
type BluelensMusic struct {
	Href string   `form:"href" json:"href" xml:"href"`
	ID   string   `form:"id" json:"id" xml:"id"`
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
}

// Validate validates the BluelensMusic media type instance.
func (mt *BluelensMusic) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	return
}

// A music resource (link view)
//
// Identifier: application/vnd.bluelens.music+json; view=link
type BluelensMusicLink struct {
	Href string `form:"href" json:"href" xml:"href"`
	ID   string `form:"id" json:"id" xml:"id"`
}

// Validate validates the BluelensMusicLink media type instance.
func (mt *BluelensMusicLink) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	return
}

// DecodeBluelensMusic decodes the BluelensMusic instance encoded in resp body.
func (c *Client) DecodeBluelensMusic(resp *http.Response) (*BluelensMusic, error) {
	var decoded BluelensMusic
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeBluelensMusicLink decodes the BluelensMusicLink instance encoded in resp body.
func (c *Client) DecodeBluelensMusicLink(resp *http.Response) (*BluelensMusicLink, error) {
	var decoded BluelensMusicLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// BluelensMusicCollection is the media type for an array of BluelensMusic (default view)
//
// Identifier: application/vnd.bluelens.music+json; type=collection; view=default
type BluelensMusicCollection []*BluelensMusic

// Validate validates the BluelensMusicCollection media type instance.
func (mt BluelensMusicCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// BluelensMusicCollection is the media type for an array of BluelensMusic (link view)
//
// Identifier: application/vnd.bluelens.music+json; type=collection; view=link
type BluelensMusicLinkCollection []*BluelensMusicLink

// Validate validates the BluelensMusicLinkCollection media type instance.
func (mt BluelensMusicLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeBluelensMusicCollection decodes the BluelensMusicCollection instance encoded in resp body.
func (c *Client) DecodeBluelensMusicCollection(resp *http.Response) (BluelensMusicCollection, error) {
	var decoded BluelensMusicCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeBluelensMusicLinkCollection decodes the BluelensMusicLinkCollection instance encoded in resp body.
func (c *Client) DecodeBluelensMusicLinkCollection(resp *http.Response) (BluelensMusicLinkCollection, error) {
	var decoded BluelensMusicLinkCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A list of recommendations for the specified user (all view)
//
// Identifier: application/vnd.bluelens.recommendations+json; view=all
type BluelensRecommendationsAll struct {
	// Links to related resources
	Links *BluelensRecommendationsLinks `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	List  BluelensMusicCollection       `form:"list" json:"list" xml:"list"`
	User  *BluelensUser                 `form:"user" json:"user" xml:"user"`
}

// Validate validates the BluelensRecommendationsAll media type instance.
func (mt *BluelensRecommendationsAll) Validate() (err error) {
	if mt.List == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "list"))
	}
	if mt.User == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "user"))
	}
	if mt.Links != nil {
		if err2 := mt.Links.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if err2 := mt.List.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if mt.User != nil {
		if err2 := mt.User.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// A list of recommendations for the specified user (default view)
//
// Identifier: application/vnd.bluelens.recommendations+json; view=default
type BluelensRecommendations struct {
	// Links to related resources
	Links   *BluelensRecommendationsLinks `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	MusicID []string                      `form:"musicID,omitempty" json:"musicID,omitempty" xml:"musicID,omitempty"`
}

// Validate validates the BluelensRecommendations media type instance.
func (mt *BluelensRecommendations) Validate() (err error) {
	if mt.Links != nil {
		if err2 := mt.Links.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// BluelensRecommendationsLinks contains links to related resources of BluelensRecommendations.
type BluelensRecommendationsLinks struct {
	List BluelensMusicLinkCollection `form:"list,omitempty" json:"list,omitempty" xml:"list,omitempty"`
	User *BluelensUserLink           `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
}

// Validate validates the BluelensRecommendationsLinks type instance.
func (ut *BluelensRecommendationsLinks) Validate() (err error) {
	if err2 := ut.List.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if ut.User != nil {
		if err2 := ut.User.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeBluelensRecommendationsAll decodes the BluelensRecommendationsAll instance encoded in resp body.
func (c *Client) DecodeBluelensRecommendationsAll(resp *http.Response) (*BluelensRecommendationsAll, error) {
	var decoded BluelensRecommendationsAll
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeBluelensRecommendations decodes the BluelensRecommendations instance encoded in resp body.
func (c *Client) DecodeBluelensRecommendations(resp *http.Response) (*BluelensRecommendations, error) {
	var decoded BluelensRecommendations
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A user resource (all view)
//
// Identifier: application/vnd.bluelens.user+json; view=all
type BluelensUserAll struct {
	Followees BluelensUserCollection  `form:"followees,omitempty" json:"followees,omitempty" xml:"followees,omitempty"`
	History   BluelensMusicCollection `form:"history,omitempty" json:"history,omitempty" xml:"history,omitempty"`
	Href      string                  `form:"href" json:"href" xml:"href"`
	ID        string                  `form:"id" json:"id" xml:"id"`
}

// Validate validates the BluelensUserAll media type instance.
func (mt *BluelensUserAll) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if err2 := mt.Followees.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if err2 := mt.History.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// A user resource (default view)
//
// Identifier: application/vnd.bluelens.user+json; view=default
type BluelensUser struct {
	Href string `form:"href" json:"href" xml:"href"`
	ID   string `form:"id" json:"id" xml:"id"`
	// Links to related resources
	Links *BluelensUserLinks `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// Validate validates the BluelensUser media type instance.
func (mt *BluelensUser) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Links != nil {
		if err2 := mt.Links.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// A user resource (link view)
//
// Identifier: application/vnd.bluelens.user+json; view=link
type BluelensUserLink struct {
	Href string `form:"href" json:"href" xml:"href"`
	ID   string `form:"id" json:"id" xml:"id"`
}

// Validate validates the BluelensUserLink media type instance.
func (mt *BluelensUserLink) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	return
}

// BluelensUserLinks contains links to related resources of BluelensUser.
type BluelensUserLinks struct {
	Followees BluelensUserLinkCollection  `form:"followees,omitempty" json:"followees,omitempty" xml:"followees,omitempty"`
	History   BluelensMusicLinkCollection `form:"history,omitempty" json:"history,omitempty" xml:"history,omitempty"`
}

// Validate validates the BluelensUserLinks type instance.
func (ut *BluelensUserLinks) Validate() (err error) {
	if err2 := ut.Followees.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if err2 := ut.History.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeBluelensUserAll decodes the BluelensUserAll instance encoded in resp body.
func (c *Client) DecodeBluelensUserAll(resp *http.Response) (*BluelensUserAll, error) {
	var decoded BluelensUserAll
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeBluelensUser decodes the BluelensUser instance encoded in resp body.
func (c *Client) DecodeBluelensUser(resp *http.Response) (*BluelensUser, error) {
	var decoded BluelensUser
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeBluelensUserLink decodes the BluelensUserLink instance encoded in resp body.
func (c *Client) DecodeBluelensUserLink(resp *http.Response) (*BluelensUserLink, error) {
	var decoded BluelensUserLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// BluelensUserCollection is the media type for an array of BluelensUser (all view)
//
// Identifier: application/vnd.bluelens.user+json; type=collection; view=all
type BluelensUserAllCollection []*BluelensUserAll

// Validate validates the BluelensUserAllCollection media type instance.
func (mt BluelensUserAllCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// BluelensUserCollection is the media type for an array of BluelensUser (default view)
//
// Identifier: application/vnd.bluelens.user+json; type=collection; view=default
type BluelensUserCollection []*BluelensUser

// Validate validates the BluelensUserCollection media type instance.
func (mt BluelensUserCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// BluelensUserCollection is the media type for an array of BluelensUser (link view)
//
// Identifier: application/vnd.bluelens.user+json; type=collection; view=link
type BluelensUserLinkCollection []*BluelensUserLink

// Validate validates the BluelensUserLinkCollection media type instance.
func (mt BluelensUserLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// BluelensUserLinksArray contains links to related resources of BluelensUserCollection.
type BluelensUserLinksArray []*BluelensUserLinks

// Validate validates the BluelensUserLinksArray type instance.
func (ut BluelensUserLinksArray) Validate() (err error) {
	for _, e := range ut {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeBluelensUserAllCollection decodes the BluelensUserAllCollection instance encoded in resp body.
func (c *Client) DecodeBluelensUserAllCollection(resp *http.Response) (BluelensUserAllCollection, error) {
	var decoded BluelensUserAllCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeBluelensUserCollection decodes the BluelensUserCollection instance encoded in resp body.
func (c *Client) DecodeBluelensUserCollection(resp *http.Response) (BluelensUserCollection, error) {
	var decoded BluelensUserCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeBluelensUserLinkCollection decodes the BluelensUserLinkCollection instance encoded in resp body.
func (c *Client) DecodeBluelensUserLinkCollection(resp *http.Response) (BluelensUserLinkCollection, error) {
	var decoded BluelensUserLinkCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
