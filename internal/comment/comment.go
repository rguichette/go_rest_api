package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented ")
)

//Comment - a representation of the comment stucture for our service
type Comment struct {
	Id     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface defines ALL of the methods that our SERVICE needs in order to operate
type Store interface {
	GetCommentFromStore(context.Context, string) (Comment, error)
}

//Service - is the struct on which all our logic will be built on top of.
type Service struct {
	commentStore Store
}

//NewService - returns a pointer to a new service
func NewService(store Store) *Service {

	return &Service{commentStore: store}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retreiving a comment")

	cmt, err := s.commentStore.GetCommentFromStore(ctx, id) //

	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil

}

func (s *Service) updateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}
func (s *Service) deleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) createComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
