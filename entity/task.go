package entity

const (
	TaskStatusEmpty = iota
	TaskStatusArchive
	TaskStatusDone
)

type Task struct {
	Status   int64
	ZipPath  string
	URLSLice []string
}
