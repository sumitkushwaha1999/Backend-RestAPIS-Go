package requests

type UserData struct {
}

type Person struct {
	Name        string `json:"name" binding:"required"`
	Age         int    `json:"age" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Status      string `json:"status"`
}
