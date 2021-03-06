// Code generated by goagen v1.1.0, command line:
// $ goagen
// --design=github.com/ihcsim/bluelens/design
// --out=$(GOPATH)/src/github.com/ihcsim/bluelens/cmd/blued
// --version=v1.1.0
//
// API "bluelens": Application Contexts
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"strconv"
)

// CreateMusicContext provides the music create action context.
type CreateMusicContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *Music
}

// NewCreateMusicContext parses the incoming request URL and body, performs validations and creates the
// context used by the music controller create action.
func NewCreateMusicContext(ctx context.Context, service *goa.Service) (*CreateMusicContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateMusicContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// CreatedLink sends a HTTP response with status code 201.
func (ctx *CreateMusicContext) CreatedLink(r *BluelensMusicLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateMusicContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *CreateMusicContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateMusicContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ListMusicContext provides the music list action context.
type ListMusicContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Limit  int
	Offset int
}

// NewListMusicContext parses the incoming request URL and body, performs validations and creates the
// context used by the music controller list action.
func NewListMusicContext(ctx context.Context, service *goa.Service) (*ListMusicContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListMusicContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramLimit := req.Params["limit"]
	if len(paramLimit) > 0 {
		rawLimit := paramLimit[0]
		if limit, err2 := strconv.Atoi(rawLimit); err2 == nil {
			rctx.Limit = limit
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("limit", rawLimit, "integer"))
		}
	}
	paramOffset := req.Params["offset"]
	if len(paramOffset) > 0 {
		rawOffset := paramOffset[0]
		if offset, err2 := strconv.Atoi(rawOffset); err2 == nil {
			rctx.Offset = offset
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("offset", rawOffset, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListMusicContext) OK(r BluelensMusicCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.bluelens.music+json; type=collection")
	if r == nil {
		r = BluelensMusicCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKFull sends a HTTP response with status code 200.
func (ctx *ListMusicContext) OKFull(r BluelensMusicFullCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.bluelens.music+json; type=collection")
	if r == nil {
		r = BluelensMusicFullCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ListMusicContext) OKLink(r BluelensMusicLinkCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.bluelens.music+json; type=collection")
	if r == nil {
		r = BluelensMusicLinkCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ListMusicContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListMusicContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ShowMusicContext provides the music show action context.
type ShowMusicContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewShowMusicContext parses the incoming request URL and body, performs validations and creates the
// context used by the music controller show action.
func NewShowMusicContext(ctx context.Context, service *goa.Service) (*ShowMusicContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowMusicContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OKFull sends a HTTP response with status code 200.
func (ctx *ShowMusicContext) OKFull(r *BluelensMusicFull) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ShowMusicContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowMusicContext) NotFound(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowMusicContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// RecommendRecommendationsContext provides the recommendations recommend action context.
type RecommendRecommendationsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Limit  int
	UserID string
}

// NewRecommendRecommendationsContext parses the incoming request URL and body, performs validations and creates the
// context used by the recommendations controller recommend action.
func NewRecommendRecommendationsContext(ctx context.Context, service *goa.Service) (*RecommendRecommendationsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := RecommendRecommendationsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramLimit := req.Params["limit"]
	if len(paramLimit) > 0 {
		rawLimit := paramLimit[0]
		if limit, err2 := strconv.Atoi(rawLimit); err2 == nil {
			rctx.Limit = limit
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("limit", rawLimit, "integer"))
		}
		if rctx.Limit < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`limit`, rctx.Limit, 0, true))
		}
	}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		rctx.UserID = rawUserID
	}
	return &rctx, err
}

// OKAll sends a HTTP response with status code 200.
func (ctx *RecommendRecommendationsContext) OKAll(r *BluelensRecommendationsAll) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OK sends a HTTP response with status code 200.
func (ctx *RecommendRecommendationsContext) OK(r *BluelensRecommendations) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *RecommendRecommendationsContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *RecommendRecommendationsContext) NotFound(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *RecommendRecommendationsContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// CreateUserContext provides the user create action context.
type CreateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *User
}

// NewCreateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller create action.
func NewCreateUserContext(ctx context.Context, service *goa.Service) (*CreateUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// CreatedLink sends a HTTP response with status code 201.
func (ctx *CreateUserContext) CreatedLink(r *BluelensUserLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateUserContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *CreateUserContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateUserContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// FollowUserContext provides the user follow action context.
type FollowUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	FolloweeID string
	ID         string
	Payload    *FollowUserPayload
}

// NewFollowUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller follow action.
func NewFollowUserContext(ctx context.Context, service *goa.Service) (*FollowUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := FollowUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramFolloweeID := req.Params["followeeID"]
	if len(paramFolloweeID) > 0 {
		rawFolloweeID := paramFolloweeID[0]
		rctx.FolloweeID = rawFolloweeID
	}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// followUserPayload is the user follow action payload.
type followUserPayload struct {
	// ID of the followee.
	FolloweeID *string `form:"followeeID,omitempty" json:"followeeID,omitempty" xml:"followeeID,omitempty"`
}

// Publicize creates FollowUserPayload from followUserPayload
func (payload *followUserPayload) Publicize() *FollowUserPayload {
	var pub FollowUserPayload
	if payload.FolloweeID != nil {
		pub.FolloweeID = payload.FolloweeID
	}
	return &pub
}

// FollowUserPayload is the user follow action payload.
type FollowUserPayload struct {
	// ID of the followee.
	FolloweeID *string `form:"followeeID,omitempty" json:"followeeID,omitempty" xml:"followeeID,omitempty"`
}

// OK sends a HTTP response with status code 200.
func (ctx *FollowUserContext) OK(r *BluelensUser) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKFull sends a HTTP response with status code 200.
func (ctx *FollowUserContext) OKFull(r *BluelensUserFull) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *FollowUserContext) OKLink(r *BluelensUserLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *FollowUserContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *FollowUserContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *FollowUserContext) NotFound(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *FollowUserContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ListUserContext provides the user list action context.
type ListUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Limit  int
	Offset int
}

// NewListUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller list action.
func NewListUserContext(ctx context.Context, service *goa.Service) (*ListUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramLimit := req.Params["limit"]
	if len(paramLimit) > 0 {
		rawLimit := paramLimit[0]
		if limit, err2 := strconv.Atoi(rawLimit); err2 == nil {
			rctx.Limit = limit
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("limit", rawLimit, "integer"))
		}
	}
	paramOffset := req.Params["offset"]
	if len(paramOffset) > 0 {
		rawOffset := paramOffset[0]
		if offset, err2 := strconv.Atoi(rawOffset); err2 == nil {
			rctx.Offset = offset
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("offset", rawOffset, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListUserContext) OK(r BluelensUserCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.bluelens.user+json; type=collection")
	if r == nil {
		r = BluelensUserCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKFull sends a HTTP response with status code 200.
func (ctx *ListUserContext) OKFull(r BluelensUserFullCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.bluelens.user+json; type=collection")
	if r == nil {
		r = BluelensUserFullCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ListUserContext) OKLink(r BluelensUserLinkCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.bluelens.user+json; type=collection")
	if r == nil {
		r = BluelensUserLinkCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ListUserContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListUserContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ListenUserContext provides the user listen action context.
type ListenUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      string
	MusicID string
	Payload *ListenUserPayload
}

// NewListenUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller listen action.
func NewListenUserContext(ctx context.Context, service *goa.Service) (*ListenUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListenUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	paramMusicID := req.Params["musicID"]
	if len(paramMusicID) > 0 {
		rawMusicID := paramMusicID[0]
		rctx.MusicID = rawMusicID
	}
	return &rctx, err
}

// listenUserPayload is the user listen action payload.
type listenUserPayload struct {
	// ID of the music.
	MusicID *string `form:"musicID,omitempty" json:"musicID,omitempty" xml:"musicID,omitempty"`
}

// Publicize creates ListenUserPayload from listenUserPayload
func (payload *listenUserPayload) Publicize() *ListenUserPayload {
	var pub ListenUserPayload
	if payload.MusicID != nil {
		pub.MusicID = payload.MusicID
	}
	return &pub
}

// ListenUserPayload is the user listen action payload.
type ListenUserPayload struct {
	// ID of the music.
	MusicID *string `form:"musicID,omitempty" json:"musicID,omitempty" xml:"musicID,omitempty"`
}

// OK sends a HTTP response with status code 200.
func (ctx *ListenUserContext) OK(r *BluelensUser) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKFull sends a HTTP response with status code 200.
func (ctx *ListenUserContext) OKFull(r *BluelensUserFull) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ListenUserContext) OKLink(r *BluelensUserLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ListenUserContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ListenUserContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListenUserContext) NotFound(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListenUserContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ShowUserContext provides the user show action context.
type ShowUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewShowUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller show action.
func NewShowUserContext(ctx context.Context, service *goa.Service) (*ShowUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OKFull sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OKFull(r *BluelensUserFull) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ShowUserContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowUserContext) NotFound(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowUserContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}
