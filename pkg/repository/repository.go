package repository

type Authorization interface {
}

type Subscription interface {
}

type Repository struct {
	Authorization
	Subscription
}

func NewRepository() *Repository {
	return &Repository{}
}
