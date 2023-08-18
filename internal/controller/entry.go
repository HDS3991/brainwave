package controller

type Controller struct {
	User IUser
}

var Entry = Controller{
	User: NewUser(),
}
