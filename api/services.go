package api 

import (
 	"github.com/dgrijalva/jwt-go"
	D "go-juan/internal/database"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
	"log"
	"time"
	"os" 
	"fmt"
)

// Database instance
var db *sql.DB 

func updatePadron(c *fiber.Ctx) error {
	
	dni := new(Padron)
	user := new(User)
	err := c.BodyParser(&dni)
	fmt.Println(dni.Dni)
	if  err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Cannot parse json",
		})
	}
	db = D.Connect(db)
	res := db.QueryRow(GetVotante(), dni.Dni)
	if res != nil {
		err := res.Scan(&user.Id)
		if(err != nil){
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"No existe usuario con este dni.",
			})
		}
	}

	resp, err := db.Query(UpdatePadron(), user.Id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"No se ha podido actualizar este registro.",
		})
	}
	log.Println(resp)
    fmt.Println(user.Id)

	defer db.Close();
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"mensaje": "Se ha actualizado el padron correctamente.",
		"ok": true,
 	})
}


func getPadron(c *fiber.Ctx) error {

	db = D.Connect( db )
	rows, err := db.Query( SelectPadron() )
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	defer rows.Close()
	result := Padrones{}

	for rows.Next() {
		padron := Padron{}
		if err := rows.Scan(&padron.Dni, &padron.Nombre, &padron.Apellido, &padron.Voto); err != nil {
			return err
		}
		result.Padrones = append(result.Padrones, padron)
	}
	return c.JSON(result)
}


func storeDniPadron(c *fiber.Ctx) error {
	db = D.Connect( db )

	p := new(Padron)
	err := c.BodyParser(&p.Dni)
	p.Voto = true
	if  err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Cannot parse json",
		})
	}

	res, err := db.Query( InsertPadron(), &p.Dni, "", "", &p.Voto )
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Dni duplicado o con valores incorrectos.",
		})
	}

	if( res != nil ){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"No se ha podido registrar este documento.",
		})
	 }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Se ha actualizado un registro en el padron.",
	});

}


func storePadron(c *fiber.Ctx) error {

	db = D.Connect( db )
	u := new(Padron)
	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	res, err := db.Query( InsertPadron() , u.Dni, u.Nombre, u.Apellido, u.Voto )
	if err != nil {
		return err
	}

	log.Println(res) 
	return c.JSON(u)
}


// LOGIN APLICATION
func login(c *fiber.Ctx) error {

	login := new(Login)
	user := new(User)
	err := c.BodyParser(&login)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Cannot parse json",
		})
	}

	db = D.Connect( db )
	res := db.QueryRow(LoginUser(),login.Email) // Buscamos el usuario por el email
	if res != nil {
		err := res.Scan(&user.Id, &user.Password, &user.Email)
		if(err != nil){
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":"Bad credentials",
			})
		}
	}
	
	hash := user.Password
	hashComoByte := []byte(hash)
	contraseña := login.Password
	contraseñaComoByte := []byte(contraseña)
	error := bcrypt.CompareHashAndPassword(hashComoByte, contraseñaComoByte) // Validación de la contrasenia por el hash
	if error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":"Bad credentials",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7) // Una semana
	jwtSecret := os.Getenv("TOKEN_KEY")
	s, err := token.SignedString([]byte( jwtSecret ))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": s,
		"user" : struct {
			Id 		int `json:"id"`
			Email 	string `json:"email"`
		}{
			Id: user.Id,
			Email: user.Email,
		},
 	})
}


