package model

type Task struct {
	ID     int64
	ListID int64
	UserID int64
	Title  string
	Memo   string
	IsDone bool
}

func NewTask(ID int64, listID int64, userID int64, title string, memo string, isDone bool) *Task {
	return &Task{
		ID:     ID,
		ListID: listID,
		UserID: userID,
		Title:  title,
		Memo:   memo,
		IsDone: isDone,
	}
}

// maxTask リストが所持可能な最大タスク数
const maxTask = 10
