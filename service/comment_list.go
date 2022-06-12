package service

import (
	"dousheng/proto/proto"
	"dousheng/repository"
)

func CommentList(vid int64) ([]*proto.Comment, error) {

	commentList, err := repository.NewCommentDaoInstance().QueryByVId(vid)
	if err != nil {
		return nil, err
	}
	var protoCommentList []*proto.Comment
	for _, comment := range *commentList {
		user, err := repository.NewUserDaoInstance().QueryUserById(comment.UId)
		if err != nil {
			return nil, err
		}
		demoUser := &proto.User{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      false,
		}
		month := comment.CreateTime.Format("01")
		date := comment.CreateTime.Format("02")

		protoCommentList = append(protoCommentList, &proto.Comment{
			Id:         comment.Id,
			User:       demoUser,
			Content:    comment.Content,
			CreateDate: month + "-" + date,
		})
	}
	return protoCommentList, nil
}
