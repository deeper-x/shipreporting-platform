package db

// Assignment defines portinformer assigned to user
type Assignment struct {
	Name         string
	Portinformer int
	Customer     int // This will be foreign key, by Model(&Group).AddForeignKey
}

// Customer defines client data
type Customer struct {
	ID             int
	Email          string
	Password       string
	Portinformer_1 int
	Portinformer_2 int
	Portinformer_3 int
}
