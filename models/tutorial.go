package models

type Tutorial struct {
	Id       int
	Title    string
	Author   Author
	Comments []Comment
}
