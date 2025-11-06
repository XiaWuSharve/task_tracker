package tasktracker

type Status int

const (
	todo       Status = iota
	inProgress Status = iota
	done       Status = iota
)
