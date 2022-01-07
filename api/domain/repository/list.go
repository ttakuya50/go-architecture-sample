package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/ttakuya50/go-architecture-sample/api/infra/mysql"

	"github.com/ttakuya50/go-architecture-sample/api/domain/model"
)

//go:generate mockgen -source=$GOFILE -destination=list_mock.go -package=$GOPACKAGE -self_package=github.com/ttakuya50/go-architecture-sample/api/domain/$GOPACKAGE

type ListRepo interface {
	BulkDelete(ctx context.Context, db ContextExecutor, lists []*model.List) error
	Create(ctx context.Context, db ContextExecutor, ID, userID int64, title string) error
}

type listRepo struct {
}

func NewListRepo() ListRepo {
	return &listRepo{}
}

func (l *listRepo) BulkDelete(ctx context.Context, db ContextExecutor, lists []*model.List) error {
	ids := make([]interface{}, 0, len(lists))
	for _, l := range lists {
		ids = append(ids, l.ID)
	}

	rowsAff, err := mysql.Lists(qm.WhereIn(mysql.ListColumns.ID+" IN ?", ids...)).DeleteAll(ctx, db)
	if err != nil {
		return err
	}

	if rowsAff != int64(len(lists)) {
		return errors.New(fmt.Sprintf("削除された件数が一致しません.rowsAff=%d len=%d\n", rowsAff, len(lists)))
	}

	return nil
}

func (l *listRepo) Create(ctx context.Context, db ContextExecutor, ID, userID int64, title string) error {
	list := &mysql.List{
		ID:        ID,
		UserID:    userID,
		Title:     title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return list.Insert(ctx, db, boil.Infer())
}
