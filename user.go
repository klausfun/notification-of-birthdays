package NotificationOfBirthdays

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type Author struct {
	Id    int    `json:"-" db:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
