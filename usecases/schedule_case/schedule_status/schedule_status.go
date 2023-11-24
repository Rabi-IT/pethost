package schedule_status

type Status string

var (
	Open         Status = "open"
	Accepted     Status = "accepted"
	Paid         Status = "paid"
	PaidAccepted Status = "paid_accepted"
	Canceled     Status = "canceled"
	Refunded     Status = "refunded"
	Finished     Status = "finished"
)
