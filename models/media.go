package models

type MediaFolder struct {
	Meta Internals
	User UserRef
	Name string `json:"name" firestore:"name"`
}

func (user *User) NewMediaFolder(name string) *MediaFolder {
	return &MediaFolder{
		Meta: NewInternals("mediafolder"),
		User: user.Ref(),
		Name: name,
	}
}

type MediaFile struct {
	Meta        Internals
	User        UserRef
	Folder      string `json:"folder" firestore:"folder"`
	Kind        string `json:"kind" firestore:"kind"`
	URI         string `json:"uri" firestore:"uri"`
	Description string `json:"description" firestore:"description"`
}

func (folder *MediaFolder) NewMediaFile(user *User, kind, uri, description string) *MediaFile {
	return &MediaFile{
		Meta:        NewInternals("mediafile"),
		User:        user.Ref(),
		Folder:      folder.Name,
		Kind:        kind,
		URI:         uri,
		Description: description,
	}
}
