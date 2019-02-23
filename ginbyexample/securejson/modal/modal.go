package modal

import "time"

/*
Login ...
*/
type Login struct {
	User     string `json:"user" form:"user" xml:"user" binding:"required,min=10,max=12"`
	Password string `json:"password" form:"password" xml:"password" binding:"required,PhoneNumerValidator"`
}

/*
User ...
*/
type User struct {
	Name string    `json:"name"`
	DOB  time.Time `json:"dob"`
}

/*
Success ...
*/
type Success struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

/*
ErrorMessage ...
*/
type ErrorMessage struct {
	Message string `json:"error"`
	Status  int    `json:"status"`
}
