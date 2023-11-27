package models

func (user *User) NewHealthJournal() *HealthJournal {
	return &HealthJournal{
		User: user.Ref(),
	}
}

type HealthJournal struct {
	User    UserRef
	Entries []TopicUpdate
}
