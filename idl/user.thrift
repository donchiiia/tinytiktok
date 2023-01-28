namespace go user

struct User {
    1: required i64 id // 用户id
    2: required string name // 用户名称
    3: optional i64 follow_count // 关注总数
    4: optional i64 follower_count // 粉丝总数
    5: required bool is_follow // true-已关注，false-未关注
}
struct DouyinUserRegisterRequest {
    1: required string username (vt.max_size = "32")// 注册用户名，最长32个字符
    2: required string password (vt.max_size = "32")// 密码，最长32个字符
}
struct DouyinUserRegisterResponse {
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
    3: required i64 user_id // 用户id
    4: required string token // 用户鉴权token
}
struct DouyinUserLoginRequest {
    1: required string username (vt.max_size = "32")// 注册用户名，最长32个字符
    2: required string password (vt.max_size = "32")// 密码，最长32个字符
}
struct DouyinUserLoginResponse {
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
    3: required i64 user_id // 用户id
    4: required string token // 用户鉴权token
}
struct DouyinUserRequest {
    1: required i64 user_id // 用户id
    2: required string token // 用户鉴权token
}
struct DouyinUserResponse {
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
    3: required User user // 用户信息
}

service UserService {
    DouyinUserRegisterResponse Register (1: DouyinUserRegisterRequest req)
    DouyinUserLoginResponse Login (1: DouyinUserLoginRequest req)
    DouyinUserResponse GetUserByID (1: DouyinUserRequest req)
}
