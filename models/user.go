package models

type User struct {
	Name    string
	Address string
	Job     string
	Reason  string
}

func (user *User) NewUser(name string, address string, job string, reason string) *User {
	return &User{
		Name:    name,
		Address: address,
		Job:     job,
		Reason:  reason,
	}
}
