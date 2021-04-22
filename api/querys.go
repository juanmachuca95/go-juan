package api

func SelectPadron() string {
	return "SELECT dni, nombre, apellido, voto FROM padron"
}

func InsertPadron() string {
	return "INSERT INTO padron (dni, nombre, apellido, voto ) VALUES (?, ?, ?, ?)"
}

// Jefes de mesa
func LoginUser() string {
	return "SELECT id, password, email FROM users WHERE email = ?"
}