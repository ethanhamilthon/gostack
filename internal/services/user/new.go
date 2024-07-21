package user

type UserService struct {
}

type DB interface{}

func New(db DB) *UserService {
	return &UserService{}
}
