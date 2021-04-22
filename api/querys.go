package api

func SelectPadron() string {
	return "SELECT dni, nombre, apellido, voto FROM padron"
}

func InsertPadron() string {
	return "INSERT INTO padron (dni, nombre, apellido, voto ) VALUES (?, ?, ?, ?)"
}

func UpdatePadron() string {
	return "UPDATE padron SET voto=1 WHERE id= ?"
}

func GetVotante() string {
	return "SELECT id FROM padron WHERE dni= ?"
}

// Jefes de mesa
func LoginUser() string {
	return "SELECT id, password, email FROM users WHERE email = ?"
}