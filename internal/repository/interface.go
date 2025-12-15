package repository

type UserRepository interface {
	FindAll()
	Create()
	GetByID()
	Update()
	Delete()
}
