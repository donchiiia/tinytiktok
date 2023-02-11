package pack

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/comment"
)

// CommentInfo 将 db.Comment 包装成 comment.Comment
func CommentInfo(dbComment *db.Comment, dbUser *db.User, isFollow bool) *comment.Comment {
	c := &comment.Comment{
		Id:         dbComment.ID,
		User:       UserInfo(dbUser, isFollow),
		Content:    dbComment.Content,
		CreateDate: dbComment.CreatedAt.Format("01-02"),
	}
	return c
}

// CommentListInfo 将 db.Comment列表 包装成 comment.Comment列表
// 涉及db操作的包装函数要传入context
// 评论列表要展示用户信息，而用户信息存在与当前登录用户的关注信息，每次查询都调用db很消耗资源，直接传入内存中的followSet
func CommentListInfo(ctx context.Context, dbCommentList []*db.Comment, currID int64, dbFollowSet map[int64]struct{}) ([]*comment.Comment, error) {
	commentList := make([]*comment.Comment, 0)
	for _, v := range dbCommentList {
		user, err := db.GetUsersByID(ctx, v.UserID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		var isFollow = false

		if currID != -1 {
			_, isFollow = dbFollowSet[v.UserID]
		}

		commentList = append(commentList, CommentInfo(v, user[0], isFollow))
	}
	return commentList, nil
}
