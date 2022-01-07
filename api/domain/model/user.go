package model

import "fmt"

type User struct {
	ID    int64
	Name  string
	Lists []*List
}

func NewUser(ID int64, name string, lists []*List) *User {
	return &User{
		ID:    ID,
		Name:  name,
		Lists: lists,
	}
}

// Tasks リストを超えてユーザーが所持している全てのタスクを取得
func (u *User) Tasks() []*Task {
	var res []*Task
	for _, l := range u.Lists {
		for _, t := range l.Tasks {
			res = append(res, t)
		}
	}
	return res
}

// IsExceededListMax 所持可能な最大リスト数を超えているか
func (u *User) IsExceededListMax() bool {
	fmt.Printf("len(u.Lists)=%v\n", len(u.Lists))

	if maxList <= len(u.Lists) {
		return true
	}

	return false
}
