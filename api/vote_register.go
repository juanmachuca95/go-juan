package api

import(
	"go-juan/internal/database"
	"go-juan/internal/logs"
)


type CreateVoteCMD struct {
	nomyape       	string `json:"nomyape"`
	dni        		string `json:"dni"`
	registrado		string `json:"registrado"`
	created_at 		string `json:"created_at"`
	updated_at 		string `json:"updated_at"`
}


func (us *UserService) SaveDni(cmd CreateVoteCMD) (*CreateVoteCMD, error) {
	/* fecha */
	t := time.Now()

	fecha := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	/* Crear usuario */
	_, err := us.Exec(CreateDniQuery(), cmd.Numero, fecha)

	if err != nil {

		logs.Error("cannot insert user" + err.Error())

		return nil, err
	}

	return &CreateDniCMD{
		Numero: cmd.Numero,
	}, nil
}