package model

type UserInterface interface {
	FirstOrCreateUser(user *User) error
	SetUser(user *User) error
	GetUser(wxUid string) (User, error)
}

type UserCacheInterface interface {
}
