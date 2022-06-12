package service

import (
	"dousheng/repository"
	"time"
)

type CommentActionFlow struct {
	UId     int64
	VId     int64
	Content string
}

func CommentAction(userId int64, videoId int64, content string) error {
	return NewCommentActionFlow(userId, videoId, content).Do()
}

func NewCommentActionFlow(userId int64, videoId int64, content string) *CommentActionFlow {
	return &CommentActionFlow{
		UId:     userId,
		VId:     videoId,
		Content: content,
	}
}

func (c *CommentActionFlow) Do() error {
	if err := c.checkParam(); err != nil {
		return err
	}
	err := c.action()
	if err != nil {
		return err
	}
	return nil
}

func (c *CommentActionFlow) checkParam() error {
	return nil
}

func (c *CommentActionFlow) action() error {
	err := repository.NewCommentDaoInstance().CreateComment(&repository.Comment{
		VId:        c.VId,
		UId:        c.UId,
		Content:    c.Content,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		IsDeleted:  false,
	})
	if err != nil {
		return err
	}
	return nil
}
