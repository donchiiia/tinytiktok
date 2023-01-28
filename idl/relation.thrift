namespace go relation

include "user.thrift"

struct DouyinRelationActionRequest {
    1: required i64 user_id // 用户id
    2: required string token // 用户鉴权token
    3: required i64 to_user_id // 对方用户id
    4: required i32 action_type // 1-关注，2-取消关注
}
struct DouyinRelationActionResponse {
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
}
struct DouyinRelationFollowListRequest {
    1: required i64 user_id // 用户id
    2: required string token // 用户鉴权token
}
struct DouyinRelationFollowListResponse {
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
    3: required list<user.User> user_list // 用户信息列表
}
struct DouyinRelationFollowerListRequest {
    1: required i64 user_id // 用户id
    2: required string token // 用户鉴权token
}
struct DouyinRelationFollowerListResponse {
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
    3: required list<user.User> user_list // 用户列表
}
struct DouyinRelationFriendListRequest {
    1: required i64 user_id // 用户id
    2: required string token // 用户鉴权token
}
struct DouyinRelationFriendListResponse {
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
    3: required list<user.User> user_list // 用户列表
}

service RelationService {
    DouyinRelationActionResponse RelationAction (1: DouyinRelationActionRequest req)
    DouyinRelationFollowListResponse RelationFollowList (1: DouyinRelationFollowListRequest req)
    DouyinRelationFollowerListResponse RelationFollowerList (1: DouyinRelationFollowerListRequest req)
    DouyinRelationFriendListResponse RelationFriendList (1: DouyinRelationFriendListRequest req)
}
