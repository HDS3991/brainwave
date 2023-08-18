package service

import "brainwave/internal/service/user"

type Services struct {
	User IUser
}

var Entry = Services{
	User: user.NewUser(),
}
