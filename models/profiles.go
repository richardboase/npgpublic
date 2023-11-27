package models

type Profiles struct {
	Social       Socials             `json:"socials" firestore:"socials"`
	Personal     PersonalProfile     `json:"personal" firestore:"personal"`
	Practitioner PractitionerProfile `json:"practitioner" firestore:"practitioner"`
	Business     BusinessProfile     `json:"business" firestore:"business"`
}

type PersonalProfile struct {
	Meta      Internals
	Firstname string `json:"firstname" firestore:"firstname"`
	Lastname  string `json:"lastname" firestore:"lastname"`
	Phone     string `json:"phone" firestore:"phone"`
}

type PractitionerProfile struct {
	Meta        Internals
	Description string   `json:"description" firestore:"description"`
	Phone       string   `json:"phone" firestore:"phone"`
	Website     string   `json:"website" firestore:"website"`
	Keywords    []string `json:"keywords" firestore:"keywords"`
}

type BusinessProfile struct {
	Meta        Internals
	Name        string `json:"name" firestore:"name"`
	Description string `json:"description" firestore:"description"`
	Address     string `json:"address" firestore:"address"`
	Phone       string `json:"phone" firestore:"phone"`
	Website     string `json:"website" firestore:"website"`
	VAT         string `json:"vat" firestore:"vat"`
}

type Socials struct {
	Facebook    string `json:"facebook" firestore:"facebook"`
	Telegram    string `json:"telegram" firestore:"telegram"`
	WhatsApp    string `json:"whatsapp" firestore:"whatsapp"`
	Instagram   string `json:"instagram" firestore:"instagram"`
	Snapchat    string `json:"snapchat" firestore:"snapchat"`
	MetaThreads string `json:"metathreads" firestore:"metathreads"`
	X           string `json:"x" firestore:"x"`
	What3Words  string `json:"w3w" firestore:"w3w"`
}
