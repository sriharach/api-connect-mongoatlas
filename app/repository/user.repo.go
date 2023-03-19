package repository

type User struct {
	User_name  string `json:"user_name"`
	E_mail     string `json:"e_mail"`
	Password   string `json:"password"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}
