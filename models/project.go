package models

type Project struct {
	Meta        Internals
	Name        string `json:"name" firestore:"name"`
	Description string `json:"description" firestore:"description"`
}

func (user *User) NewProject(name, description string) *Project {
	p := &Project{
		Meta:        user.Meta.NewInternals("project"),
		Name:        name,
		Description: description,
	}
	p.Meta.Moderation.Admins = []string{user.ID}
	return p
}

type ProjectOptions struct {
}
