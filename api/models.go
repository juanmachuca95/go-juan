package api

// Employee struct
type Padron struct {
	Dni         int     `json:"dni"`
	Nombre      string  `json:"nombre"`
  	Apellido    string  `json:"apellido"`
  	Voto        bool    `json:"voto"`
}

// Employees struct
type Padrones struct {
	Padrones []Padron `json:"padron"`
}

