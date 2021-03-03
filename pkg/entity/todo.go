package entity

import "errors"

type TodoList struct {
	Id          int
	Title       string
	Description string
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int
	Title       string
	Description string
	Done        bool
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string
	Description *string
}

func (u UpdateListInput) Validate() error {
	if u.Title == nil && u.Description == nil {
		return errors.New("update structure cannot have nil values")
	}
	return nil
}
