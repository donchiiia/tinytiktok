// // Code generated by hertz generator.
package api

//
//import (
//	"context"
//
//	"github.com/cloudwego/hertz/pkg/app"
//	"github.com/cloudwego/hertz/pkg/protocol/consts"
//	api "tinytiktok/cmd/api/biz/model/api"
//)
//
//// Register .
//// @router /douyin/user/register/ [POST]
//func Register(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req api.DouyinUserRegisterRequest
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//
//	resp := new(api.DouyinUserRegisterResponse)
//
//	c.JSON(consts.StatusOK, resp)
//}
//
//// Login .
//// @router /douyin/user/login/ [POST]
//func Login(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req api.DouyinUserLoginRequest
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//
//	resp := new(api.DouyinUserLoginResponse)
//
//	c.JSON(consts.StatusOK, resp)
//}
//
//// GetUserByID .
//// @router /douyin/user/ [GET]
//func GetUserByID(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req api.DouyinUserRequest
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//
//	resp := new(api.DouyinUserResponse)
//
//	c.JSON(consts.StatusOK, resp)
//}
//
//// GetUserFeed .
//// @router /douyin/feed/ [GET]
//func GetUserFeed(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req api.DouyinFeedRequest
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//
//	resp := new(api.DouyinFeedResponse)
//
//	c.JSON(consts.StatusOK, resp)
//}
//
//// PublishAction .
//// @router /douyin/publish/action/ [POST]
//func PublishAction(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req api.DouyinPublishActionRequest
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//
//	resp := new(api.DouyinPublishActionResponse)
//
//	c.JSON(consts.StatusOK, resp)
//}
//
//// PublishList .
//// @router /douyin/publish/list/ [GET]
//func PublishList(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req api.DouyinPublishListRequest
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//
//	resp := new(api.DouyinPublishListResponse)
//
//	c.JSON(consts.StatusOK, resp)
//}
//
//// FavoriteAction .
//// @router /douyin/favorite/action/ [POST]
//func FavoriteAction(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req api.DouyinFavoriteActionRequest
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//
//	resp := new(api.DouyinFavoriteActionResponse)
//
//	c.JSON(consts.StatusOK, resp)
//}
//
//// FavoriteList .
//// @router /douyin/favorite/list/ [GET]
//func FavoriteList(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req api.DouyinFavoriteListRequest
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//
//	resp := new(api.DouyinFavoriteListResponse)
//
//	c.JSON(consts.StatusOK, resp)
//}
//
//// RelationAction .
//// @router /douyin/relation/action/ [POST]
//func RelationAction(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req api.DouyinRelationActionRequest
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//
//	resp := new(api.DouyinRelationActionResponse)
//
//	c.JSON(consts.StatusOK, resp)
//}
//
//// RelationFollowList .
//// @router /douyin/relation/follow/list/ [GET]
//func RelationFollowList(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req api.DouyinRelationFollowListRequest
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//
//	resp := new(api.DouyinRelationFollowListResponse)
//
//	c.JSON(consts.StatusOK, resp)
//}
//
//// RelationFollowerList .
//// @router /douyin/relation/follower/list/ [GET]
//func RelationFollowerList(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req api.DouyinRelationFollowerListRequest
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//
//	resp := new(api.DouyinRelationFollowerListResponse)
//
//	c.JSON(consts.StatusOK, resp)
//}
//
//// RelationFriendList .
//// @router /douyin/relation/friend/list/ [GET]
//func RelationFriendList(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req api.DouyinRelationFriendListRequest
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//
//	resp := new(api.DouyinRelationFriendListResponse)
//
//	c.JSON(consts.StatusOK, resp)
//}
//
//// CommentAction .
//// @router /douyin/comment/action/ [POST]
//func CommentAction(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req api.DouyinCommentActionRequest
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//
//	resp := new(api.DouyinCommentActionResponse)
//
//	c.JSON(consts.StatusOK, resp)
//}
//
//// CommentList .
//// @router /douyin/comment/list/ [GET]
//func CommentList(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req api.DouyinCommentListRequest
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//
//	resp := new(api.DouyinCommentListResponse)
//
//	c.JSON(consts.StatusOK, resp)
//}
//
//// MessageAction .
//// @router /douyin/message/action/ [POST]
//func MessageAction(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req api.DouyinMessageActionRequest
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//
//	resp := new(api.DouyinMessageActionResponse)
//
//	c.JSON(consts.StatusOK, resp)
//}
//
//// MessageChat .
//// @router /douyin/message/chat/ [GET]
//func MessageChat(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req api.DouyinMessageChatRequest
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//
//	resp := new(api.DouyinMessageChatResponse)
//
//	c.JSON(consts.StatusOK, resp)
//}
