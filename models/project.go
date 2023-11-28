package models

type Project struct {
	Meta        Internals
	Name        string `json:"name" firestore:"name"`
	Description string `json:"description" firestore:"description"`
}

func (user *User) NewProject(name, description string) *Project {
	p := &Project{
		Meta:        (Internals{}).NewInternals("projects"),
		Name:        name,
		Description: description,
	}
	p.Meta.Moderation.Admins = []string{user.Meta.ID}
	return p
}

type ProjectOptions struct {
}
