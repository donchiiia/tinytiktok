namespace go favorite

include "feed.thrift"

struct DouyinFavoriteActionRequest {
    1: required i64 user_id // 用户id (api文档无此字段，但数据库表的用户与视频点赞关系表需要此字段)
    2: required string token // 用户鉴权token
    3: required i64 video_id // 视频id
    4: required i32 action_type // 1-点赞，2-取消点赞
}
struct douyinFavoriteActionResponse {
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
}
struct DouyinFavoriteListRequest {
    1: required i64 user_id // 用户id
    2: required string token // 用户鉴权token
}
struct DouyinFavoriteListResponse {
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
    3: required list<feed.Video> video_list // 用户点赞视频列表
}

service FavoriteService {
    douyinFavoriteActionResponse FavoriteAction (1: DouyinFavoriteActionRequest req)
    douyinFavoriteActionResponse FavoriteList (1: DouyinFavoriteListRequest req)
}
