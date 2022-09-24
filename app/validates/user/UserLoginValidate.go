package user

type UserLogin struct {
	User     string `validate:"required"`
	Password string `validate:"required"`
}

func LoginValidate(data map[string]interface{}) {
}
