package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/ttakuya50/go-architecture-sample/api/domain/model"

	"github.com/golang/mock/gomock"

	"github.com/ttakuya50/go-architecture-sample/api/domain/repository"
)

func Test_userService_Register(t *testing.T) {
	testErr := fmt.Errorf("エラー")
	userID := int64(231121)

	type fields struct {
		db         func(ctrl *gomock.Controller) repository.DB
		userRepo   func(ctrl *gomock.Controller) repository.UserRepo
		randomRepo func(ctrl *gomock.Controller) repository.RandomRepo
	}

	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "正常",
			fields: fields{
				db: func(ctrl *gomock.Controller) repository.DB {
					mock := repository.NewMockDB(ctrl)
					return mock
				},
				userRepo: func(ctrl *gomock.Controller) repository.UserRepo {
					mock := repository.NewMockUserRepo(ctrl)
					mock.EXPECT().Create(context.TODO(), gomock.Any(), userID, "ほげほげ").Return(nil)
					return mock
				},
				randomRepo: func(ctrl *gomock.Controller) repository.RandomRepo {
					mock := repository.NewMockRandomRepo(ctrl)
					mock.EXPECT().Int63().Return(userID)
					return mock
				},
			},
			args: args{
				ctx:  context.TODO(),
				name: "ほげほげ",
			},
			wantErr: false,
		},
		{
			name: "異常",
			fields: fields{
				db: func(ctrl *gomock.Controller) repository.DB {
					mock := repository.NewMockDB(ctrl)
					return mock
				},
				userRepo: func(ctrl *gomock.Controller) repository.UserRepo {
					mock := repository.NewMockUserRepo(ctrl)
					mock.EXPECT().Create(context.TODO(), gomock.Any(), userID, "ほげほげ").Return(testErr)
					return mock
				},
				randomRepo: func(ctrl *gomock.Controller) repository.RandomRepo {
					mock := repository.NewMockRandomRepo(ctrl)
					mock.EXPECT().Int63().Return(userID)
					return mock
				},
			},
			args: args{
				ctx:  context.TODO(),
				name: "ほげほげ",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			s := &userService{
				db:         tt.fields.db(ctrl),
				userRepo:   tt.fields.userRepo(ctrl),
				randomRepo: tt.fields.randomRepo(ctrl),
			}
			if err := s.Register(tt.args.ctx, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userService_Delete(t *testing.T) {
	userID := int64(12345)
	testErr := fmt.Errorf("エラー")

	type fields struct {
		db         func(ctrl *gomock.Controller) repository.DB
		userRepo   func(ctrl *gomock.Controller) repository.UserRepo
		randomRepo func(ctrl *gomock.Controller) repository.RandomRepo
		taskRepo   func(ctrl *gomock.Controller) repository.TaskRepo
		listRepo   func(ctrl *gomock.Controller) repository.ListRepo
	}
	type args struct {
		ctx    context.Context
		userID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "正常:ユーザーがリストやタスクを持っていない場合",
			fields: fields{
				db: func(ctrl *gomock.Controller) repository.DB {
					ctx := context.TODO()
					mock := repository.NewMockDB(ctrl)
					txMock := repository.NewMockTx(ctrl)
					mock.EXPECT().BeginTx(ctx, nil).Return(txMock, nil)
					txMock.EXPECT().Commit().Return(nil)
					return mock
				},
				userRepo: func(ctrl *gomock.Controller) repository.UserRepo {
					ctx := context.TODO()
					user := &model.User{
						ID:    userID,
						Name:  "ほげほげ",
						Lists: nil,
					}

					mock := repository.NewMockUserRepo(ctrl)
					mock.EXPECT().Find(ctx, gomock.Any(), userID).Return(user, nil)
					mock.EXPECT().Delete(ctx, gomock.Any(), user).Return(nil)
					return mock
				},
				randomRepo: func(ctrl *gomock.Controller) repository.RandomRepo {
					mock := repository.NewMockRandomRepo(ctrl)
					return mock
				},
				taskRepo: func(ctrl *gomock.Controller) repository.TaskRepo {
					ctx := context.TODO()
					mock := repository.NewMockTaskRepo(ctrl)
					mock.EXPECT().BulkDelete(ctx, gomock.Any(), nil).Return(nil)
					return mock
				},
				listRepo: func(ctrl *gomock.Controller) repository.ListRepo {
					ctx := context.TODO()
					mock := repository.NewMockListRepo(ctrl)
					mock.EXPECT().BulkDelete(ctx, gomock.Any(), nil).Return(nil)
					return mock
				},
			},
			args: args{
				ctx:    context.TODO(),
				userID: userID,
			},
			wantErr: false,
		},
		{
			name: "異常:エラーが発生してロールバックが行われた場合",
			fields: fields{
				db: func(ctrl *gomock.Controller) repository.DB {
					ctx := context.TODO()
					mock := repository.NewMockDB(ctrl)
					txMock := repository.NewMockTx(ctrl)
					mock.EXPECT().BeginTx(ctx, nil).Return(txMock, nil)
					txMock.EXPECT().Rollback().Return(nil)
					return mock
				},
				userRepo: func(ctrl *gomock.Controller) repository.UserRepo {
					ctx := context.TODO()
					user := &model.User{
						ID:    userID,
						Name:  "ほげほげ",
						Lists: []*model.List{},
					}

					mock := repository.NewMockUserRepo(ctrl)
					mock.EXPECT().Find(ctx, gomock.Any(), userID).Return(user, nil)
					return mock
				},
				randomRepo: func(ctrl *gomock.Controller) repository.RandomRepo {
					mock := repository.NewMockRandomRepo(ctrl)
					return mock
				},
				taskRepo: func(ctrl *gomock.Controller) repository.TaskRepo {
					ctx := context.TODO()
					mock := repository.NewMockTaskRepo(ctrl)
					mock.EXPECT().BulkDelete(ctx, gomock.Any(), nil).Return(testErr)
					return mock
				},
				listRepo: func(ctrl *gomock.Controller) repository.ListRepo {
					mock := repository.NewMockListRepo(ctrl)
					return mock
				},
			},
			args: args{
				ctx:    context.TODO(),
				userID: userID,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			s := &userService{
				db:         tt.fields.db(ctrl),
				userRepo:   tt.fields.userRepo(ctrl),
				randomRepo: tt.fields.randomRepo(ctrl),
				taskRepo:   tt.fields.taskRepo(ctrl),
				listRepo:   tt.fields.listRepo(ctrl),
			}
			if err := s.Delete(tt.args.ctx, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
