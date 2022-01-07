package repository

import (
	"context"
	"errors"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/ttakuya50/go-architecture-sample/api/domain/model"
	"github.com/ttakuya50/go-architecture-sample/api/infra/mysql"
)

//go:generate mockgen -source=$GOFILE -destination=user_mock.go -package=$GOPACKAGE -self_package=github.com/ttakuya50/go-architecture-sample/api/domain/$GOPACKAGE

type UserRepo interface {
	Find(ctx context.Context, db ContextExecutor, userID int64) (*model.User, error)
	Create(ctx context.Context, db ContextExecutor, userID int64, name string) error
	Delete(ctx context.Context, db ContextExecutor, user *model.User) error
}

type userRepo struct {
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (r *userRepo) Find(ctx context.Context, db ContextExecutor, userID int64) (*model.User, error) {
	user, err := mysql.Users(
		qm.Where(mysql.UserColumns.ID+"=?", userID),
		qm.Load("Lists"),
		qm.Load("Tasks"),
	).One(ctx, db)
	if err != nil {
		return nil, err
	}

	tasks := make([]*model.Task, 0, len(user.R.Tasks))
	for _, l := range user.R.Tasks {
		var memo string
		if l.Memo.Valid {
			memo = l.Memo.String
		}

		tasks = append(tasks, model.NewTask(
			l.ID,
			l.ListID,
			l.UserID,
			l.Title,
			memo,
			l.IsDone,
		))
	}

	lists := make([]*model.List, 0, len(user.R.Lists))
	for _, l := range user.R.Lists {
		lists = append(lists, model.NewList(
			l.ID,
			l.UserID,
			l.Title,
			tasks,
		))
	}

	return model.NewUser(user.ID, user.Name, lists), nil
}

func (r *userRepo) Create(ctx context.Context, db ContextExecutor, userID int64, name string) error {
	u := &mysql.User{
		ID:        userID,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return u.Insert(ctx, db, boil.Infer())
}

func (r *userRepo) Delete(ctx context.Context, db ContextExecutor, user *model.User) error {
	u := &mysql.User{
		ID: user.ID,
	}

	res, err := u.Delete(ctx, db)
	if err != nil {
		return err
	}
	if res != 1 {
		return errors.New("ユーザーの削除件数が一致しません")
	}
	return nil
}
