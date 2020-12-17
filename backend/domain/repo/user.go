package repo

type User interface {
	Users() ([]model.User, error)
	User(id int) (*model.User, error)
}
