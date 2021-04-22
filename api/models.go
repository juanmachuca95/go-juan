package api

// Padron struct
type Padron struct {
	Dni         int     `json:"dni"`
	Nombre      string  `json:"nombre"`
  	Apellido    string  `json:"apellido"`
  	Voto        bool    `json:"voto"`
}

// Padron struct
type Padrones struct {
	Padrones []Padron `json:"padron"`
}

// Usuario struct
type User struct {
	Id 			int `json:"id"`
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}

type Users struct{
	Users []User `json:"users"`
}

type Login struct{
	Email 		string `json:"email"`
	Password 	string `json:"password"`		
}

