package dtos

// SignInDTO represents the data required for user sign-in
type SignInDTO struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

// SignUpDTO represents the data required for user sign-up
type SignUpDTO struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8" json:"password"`
	FullName string `validate:"required" json:"fullName"`
	Gender   string `validate:"omitempty,oneof=Male Female" json:"gender"`
	Birthday string `validate:"omitempty,customDate" json:"birthday"`
	Phone    string `validate:"omitempty,numeric,len=10" json:"phone"`
}
