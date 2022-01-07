package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/ttakuya50/go-architecture-sample/api/domain/model"
	"github.com/ttakuya50/go-architecture-sample/api/infra/mysql"
)

//go:generate mockgen -source=$GOFILE -destination=task_mock.go -package=$GOPACKAGE -self_package=github.com/ttakuya50/go-architecture-sample/api/domain/$GOPACKAGE

type TaskRepo interface {
	BulkDelete(ctx context.Context, db ContextExecutor, tasks []*model.Task) error
}

type taskRepo struct {
}

func NewTaskRepo() TaskRepo {
	return &taskRepo{}
}

func (t taskRepo) BulkDelete(ctx context.Context, db ContextExecutor, tasks []*model.Task) error {
	ids := make([]interface{}, 0, len(tasks))
	for _, t := range tasks {
		ids = append(ids, t.ID)
	}

	rowsAff, err := mysql.Tasks(qm.WhereIn(mysql.TaskColumns.ID+" IN ?", ids...)).DeleteAll(ctx, db)
	if err != nil {
		return err
	}

	if rowsAff != int64(len(tasks)) {
		return errors.New(fmt.Sprintf("削除された件数が一致しません.rowsAff=%d len=%d\n", rowsAff, len(tasks)))
	}

	return nil
}
