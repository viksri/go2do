package tasks

type TaskStatus int

const (
	Unknown TaskStatus = iota
	NotStarted
	InProgress
	Blocked
	Complete
)

func (ts TaskStatus) String() string {
	switch ts {
	case NotStarted:
		return "Not Started"
	case InProgress:
		return "In Progress"
	case Blocked:
		return "Blocked"
	case Complete:
		return "Complete"
	default:
		return "Unknown"
	}
}

func FromString(str string) TaskStatus {
	switch str {
	case "Not Started":
		return NotStarted
	case "In Progress":
		return InProgress
	case "Blocked":
		return Blocked
	case "Complete":
		return Complete
	default:
		return Unknown
	}
}
