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
		return "empty"
	case TaskStatusArchive:
		return "archive"
	case TaskStatusDone:
		return "done"
	default:
		return "unknown"
	}
}
func (s TaskStatus) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}
