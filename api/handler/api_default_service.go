package handler

import (
	"context"
	"net/http"

	"github.com/ttakuya50/go-architecture-sample/api/service"
)

// DefaultApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
	userService service.UserService
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService(userService service.UserService) DefaultApiServicer {
	return &DefaultApiService{
		userService: userService,
	}
}

func (s *DefaultApiService) UserPost(ctx context.Context, object InlineObject) (ImplResponse, error) {
	if err := s.userService.Register(ctx, object.Name); err != nil {
		return Response(http.StatusNotImplemented, nil), err
	}
	return Response(http.StatusOK, nil), nil
}

func (s *DefaultApiService) UserDelete(ctx context.Context, object1 InlineObject1) (ImplResponse, error) {
	if err := s.userService.Delete(ctx, object1.UserId); err != nil {
		return Response(http.StatusNotImplemented, nil), err
	}
	return Response(http.StatusOK, nil), nil
}

// UserListPost -
func (s *DefaultApiService) UserListPost(ctx context.Context, object2 InlineObject2) (ImplResponse, error) {
	if err := s.userService.AddList(ctx, object2.UserId, object2.Title); err != nil {
		return Response(http.StatusNotImplemented, nil), err
	}

	return Response(http.StatusOK, nil), nil
}
