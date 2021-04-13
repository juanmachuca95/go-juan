package main   

import "database/sql"
import "fmt"
import	_ "github.com/go-sql-driver/mysql"

type RegistroVotos struct {
  Dni           string `json:"dni"`
  Registrado    bool   `json:"registrado"`
  Created_at    string `json:"created_at"`
}

func main() {

  /*Conección a la base de datos*/
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/gopruebas")

	//Manejo de error
	if err != nil {		
		panic(err.Error())
	}

	defer db.Close()
	
	fmt.Println("Succesfully conected to MySQL database")


  /* Insertar datos desde go */
  insert, err := db.Query("INSERT INTO registrovotos VALUES (DEFAULT, '40048379', 1, '2021-03-04', null)")

  if err != nil {
    panic(err.Error())
  }

  defer insert.Close()

  fmt.Println("Se insertado un dato en la base de datos correctamente")

  /* Leer datos desde go*/
  results, err := db.Query("SELECT dni, registrado, created_at FROM registrovotos")

  if err != nil {
    panic(err.Error())
  }

  for results.Next() {
    
    var registrosvotos RegistroVotos 

    err = results.Scan(&registrosvotos.Dni, &registrosvotos.Registrado, &registrosvotos.Created_at)
    if err != nil {
      panic(err.Error())
    }

    fmt.Println(registrosvotos.Dni)
  }
}