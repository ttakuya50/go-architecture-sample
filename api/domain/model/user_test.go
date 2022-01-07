package model

import (
	"strconv"
	"testing"
)

func TestUser_IsExceededListMax(t *testing.T) {
	userID := int64(11)

	type fields struct {
		ID    int64
		Name  string
		Lists []*List
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "正常:所持可能な最大リスト数を超えている場合",
			fields: fields{
				ID:    userID,
				Name:  "ほげほげ",
				Lists: createLists(userID, maxList),
			},
			want: true,
		},
		{
			name: "正常:所持可能な最大リスト数を超えていない場合",
			fields: fields{
				ID:    userID,
				Name:  "ほげほげ",
				Lists: createLists(userID, maxList-1),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:    tt.fields.ID,
				Name:  tt.fields.Name,
				Lists: tt.fields.Lists,
			}
			if got := u.IsExceededListMax(); got != tt.want {
				t.Errorf("IsExceededListMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createLists(userID, len int64) []*List {
	list := make([]*List, 0, len)
	for i := int64(1); i <= len; i++ {
		list = append(list, &List{
			ID:     i,
			UserID: userID,
			Title:  "リストタイトル" + strconv.FormatInt(i, 10),
			Tasks:  nil,
		})
	}
	return list
}
