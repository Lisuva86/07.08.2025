package entity

type TaskStatus int

const (
	TaskStatusEmpty TaskStatus = iota
	TaskStatusArchive
	TaskStatusDone
	TaskStatusNone
)

type Task struct {
	Status   TaskStatus `json:"status"`
	ZipPath  string     `json:"zip_path"`
	URLSLice []URLResult   `json:"urls"`
	TaskName string `json:"task_name"`
}

func (s TaskStatus) String() string {
	switch s {
	case TaskStatusEmpty:
		return "task is waiting for URL"
	case TaskStatusArchive:
		return "task performs archiving"
	case TaskStatusDone:
		return "task complete archiving"
	default:
		return "unknown"
	}
}
func (s TaskStatus) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}
