package comment

import (
	"context"
	"fmt"
)

//Comment - a representation of the comment stucture for our service
type Comment struct {
	Id     string
	Slug   string
	Body   string
	Author string
}

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
		return Comment{}, err
	}
	return cmt, nil

}
