package role

type Role string

const (
	User       Role = "user"
	Backoffice Role = "backoffice"
	System     Role = "system"
)

func (r Role) IsBackoffice() bool {
	return r == Backoffice
}

func (r Role) IsUser() bool {
	return r == User
}
