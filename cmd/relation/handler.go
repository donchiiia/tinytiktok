package main

import (
	"context"
	"tinytiktok/cmd/relation/pack"
	"tinytiktok/cmd/relation/service"
	relation "tinytiktok/kitex_gen/relation"
	"tinytiktok/pkg/errno"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	currID, err := Jwt.GetUserIDFromToken(req.Token)
	if err != nil {
		resp = pack.BuildRelationActionResp(err)
		return resp, nil
	}

	req.UserId = currID

	if req.ActionType != 1 && req.ActionType != 2 {
		resp = pack.BuildRelationActionResp(errno.ActionTypeErr)
		return resp, nil
	}
	err = service.NewRelationActionService(ctx).FollowAction(req)
	if err != nil {
		resp = pack.BuildRelationActionResp(err)
		return resp, nil
	}
	resp = pack.BuildRelationActionResp(errno.Success)
	return resp, nil
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	currID, err := Jwt.GetUserIDFromToken(req.Token)
	if err != nil {
		resp = pack.BuildFollowListResp(err)
		return resp, nil
	}

	req.UserId = currID

	followList, err := service.NewFollowListService(ctx).FollowList(req)
	if err != nil {
		resp = pack.BuildFollowListResp(err)
		return resp, nil
	}

	resp = pack.BuildFollowListResp(errno.Success)
	resp.UserList = followList
	return resp, nil
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	currID, err := Jwt.GetUserIDFromToken(req.Token)
	if err != nil {
		resp = pack.BuildFollowerListResp(err)
		return resp, nil
	}

	req.UserId = currID

	followerList, err := service.NewFollowerListService(ctx).FollowerList(req)
	if err != nil {
		resp = pack.BuildFollowerListResp(err)
		return resp, nil
	}

	resp = pack.BuildFollowerListResp(errno.Success)
	resp.UserList = followerList
	return resp, nil
}

func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *relation.DouyinRelationFriendListRequest) (resp *relation.DouyinRelationFriendListResponse, err error) {
	currID, err := Jwt.GetUserIDFromToken(req.Token)
	if err != nil {
		resp = pack.BuildFriendListResp(err)
		return resp, nil
	}

	req.UserId = currID

	friendList, err := service.NewFriendListService(ctx).FriendList(req)
	if err != nil {
		resp = pack.BuildFriendListResp(err)
		return resp, nil
	}

	resp = pack.BuildFriendListResp(errno.Success)
	resp.UserList = friendList
	return resp, nil
}
