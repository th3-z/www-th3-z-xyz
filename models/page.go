package models

type Page struct {
	SelectedTab int
	Title       string
	Id          string
	Session     *Session
}
