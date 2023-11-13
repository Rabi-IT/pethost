package schedule_status

type Status string

var (
	Open     Status = "open"
	Rejected Status = "rejected"
	Finished Status = "finished"
)
