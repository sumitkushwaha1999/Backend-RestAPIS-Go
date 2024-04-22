package responses

type Person struct {
	Id          int    `json:"id" `
	Name        string `json:"name"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phoneNumber"`
	Status      string `json:"status"`
}
