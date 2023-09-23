package user

type RegisterUserInput struct {
	Name       string `json:"name" db:"name" binding:"required"`
	Occupation string `json:"occupation" db:"occupation" binding:"required"`
	Email      string `json:"email" db:"email" binding:"required"`
	Password   string `json:"password" db:"password" binding:"required"`
}

// type LoginInput struct {
// 	Email    string `json:"email" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
