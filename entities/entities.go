package entities

//User struct entities
type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

//Poduct struct entities
type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

//Order struct entities
type Order struct {
	ID      uint    `json:"id"`
	User    User    `json:"user"`
	Product Product `json:"product"`
}
