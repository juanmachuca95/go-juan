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
)

// Database instance
var db *sql.DB 

func getPadron(c *fiber.Ctx) error {
	db = D.Connect( db )

	// Get Employee list from database
	rows, err := db.Query( SelectPadron() )
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := Padrones{}

	for rows.Next() {
		padron := Padron{}
		if err := rows.Scan(&padron.Dni, &padron.Nombre, &padron.Apellido, &padron.Voto); err != nil {
			return err // Exit if we get an error
		}

		// Append Employee to Employees
		result.Padrones = append(result.Padrones, padron)
	}

	// Return Employees in JSON format
	return c.JSON(result)
}


func storePadron(c *fiber.Ctx) error {
	db = D.Connect( db )
	u := new(Padron) //New Padron struct

	// Parse body into struct
	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Insert vote into database
	res, err := db.Query( InsertPadron() , u.Dni, u.Nombre, u.Apellido, u.Voto )
	if err != nil {
		return err
	}
	
	log.Println(res) // Print result

	return c.JSON(u)
}


func login(c *fiber.Ctx) error {

	login := new(Login)
	user := new(User)
	/* var id int
	var email string
	var password string */

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
	
	//print("Contraseña correcta")
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