package in

type User struct {
	Id       string  `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Profile  Profile `json:"profile"`
}

type Profile struct {
	Gender  string `json:"gender"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Email   string `json:"email"`
}
