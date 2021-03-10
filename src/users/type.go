package users

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Response struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	User    User   `json:"user"`
}
