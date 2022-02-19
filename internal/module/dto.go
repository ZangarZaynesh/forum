package module

type CreateUserDTO struct {
	Login           string `json: "login"`
	Password        string `json: "password"`
	ConfirmPassword string `json: "confirmpassword"`
	Email           string `json: "email"`
}

type ReadUserDTO struct {
	Id    int    `json: "id"`
	Login string `json: "login"`
}

type UpdateUserDTO struct {
	Login    string `json: "login"`
	Password string `json: "password"`
	Email    string `json: "email"`
}

type DeleteUserDTO struct {
	Id int `json: "id"`
}
