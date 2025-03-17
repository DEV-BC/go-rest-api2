package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by ID")
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment - a representation of the comment structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store interface defines all the methods our service needs to operate
type Store interface {
	GetComment(ctx context.Context, id string) (Comment, error)
}

// Service - where all our logic will be on top of. CRUD logic
type Service struct {
	Store Store
}

// NewService returns a pointer to a new service
func NewService(store Store) *Service {
	return &Service{Store: store}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving a comment")
	comment, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return comment, nil
}

func (s *Service) UpdateComment(ctx context.Context, comment Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, comment Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
