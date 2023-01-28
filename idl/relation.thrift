namespace go relation

include "user.thrift"

struct DouyinRelationActionRequest {
    1: i64 user_id // 用户id
    2: string token // 用户鉴权token
    3: i64 to_user_id // 对方用户id
    4: i32 action_type // 1-关注，2-取消关注
}
struct DouyinRelationActionResponse {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
}
struct DouyinRelationFollowListRequest {
    1: i64 user_id // 用户id
    2: string token // 用户鉴权token
}
struct DouyinRelationFollowListResponse {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<user.User> user_list // 用户信息列表
}
struct DouyinRelationFollowerListRequest {
    1: i64 user_id // 用户id
    2: string token // 用户鉴权token
}
struct DouyinRelationFollowerListResponse {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<user.User> user_list // 用户列表
}

service RelationService {
    DouyinRelationActionResponse RelationAction (1: DouyinRelationActionRequest req)
    DouyinRelationFollowListResponse RelationFollowList (1: DouyinRelationFollowListRequest req)
    DouyinRelationFollowerListResponse RelationFollowerList (1: DouyinRelationFollowerListRequest req)
}
