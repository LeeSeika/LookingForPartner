package dao

import "lookingforpartner/model"

type UserInterface interface {
	FirstOrCreateUser(user *model.User) error
	SetUser(user *model.User) error
	GetUser(wxUid string) (*model.User, error)
}
