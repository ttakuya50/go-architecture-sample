package model

type List struct {
	ID     int64
	UserID int64
	Title  string
	Tasks  []*Task
}

func NewList(ID int64, userID int64, title string, tasks []*Task) *List {
	return &List{
		ID:     ID,
		UserID: userID,
		Title:  title,
		Tasks:  tasks,
	}
}

// maxList 所持可能な最大リスト数
const maxList = 10

// IsExceededTaskMax 所持可能な最大タスク数を超えているか
func (l *List) IsExceededTaskMax() bool {
	if len(l.Tasks) < maxTask {
		return true
	}

	return false
}
