package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/ttakuya50/go-architecture-sample/api/domain/repository"
)

type UserService interface {
	Register(ctx context.Context, name string) error
	Delete(ctx context.Context, userID int64) error
	AddList(ctx context.Context, userID int64, title string) error
}

type userService struct {
	db         repository.DB
	userRepo   repository.UserRepo
	randomRepo repository.RandomRepo
	taskRepo   repository.TaskRepo
	listRepo   repository.ListRepo
}

func NewUserService(
	db repository.DB,
	userRepo repository.UserRepo,
	randomRepo repository.RandomRepo,
	taskRepo repository.TaskRepo,
	listRepo repository.ListRepo,
) *userService {
	return &userService{
		db:         db,
		userRepo:   userRepo,
		randomRepo: randomRepo,
		taskRepo:   taskRepo,
		listRepo:   listRepo,
	}
}

func (s *userService) Register(ctx context.Context, name string) error {
	userID := s.randomRepo.Int63()
	return s.userRepo.Create(ctx, s.db, userID, name)
}

func (s *userService) Delete(ctx context.Context, userID int64) error {
	user, err := s.userRepo.Find(ctx, s.db, userID)
	if err != nil {
		return err
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := s.taskRepo.BulkDelete(ctx, tx, user.Tasks()); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return txErr
		}
		return err
	}

	if err := s.listRepo.BulkDelete(ctx, tx, user.Lists); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return txErr
		}
		return err
	}

	if err := s.userRepo.Delete(ctx, tx, user); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return txErr
		}
		return err
	}

	return tx.Commit()
}

func (s *userService) AddList(ctx context.Context, userID int64, title string) error {
	user, err := s.userRepo.Find(ctx, s.db, userID)
	if err != nil {
		return err
	}
	if user.IsExceededListMax() {
		return errors.New(fmt.Sprintf("所持数を超えるためリストが追加できません"))
	}

	ID := s.randomRepo.Int63()
	if err := s.listRepo.Create(ctx, s.db, ID, userID, title); err != nil {
		return err
	}

	return nil
}
