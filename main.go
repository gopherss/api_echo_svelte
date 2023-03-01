package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string    `json:"name"  gorm:"<-"`
	Age       int       `json:"age"  gorm:"<-"`
	CreatedAt time.Time `json:"created"`
}

var DB *gorm.DB
var users = make([]User, 0, 4)

// var seq int = 1

const (
	host     = "localhost"
	user     = "postgres"
	password = "hacker"
	dbname   = "api_svelte_echo"
	port     = 5432
)

var err error

func connectionDB() {
	chainConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", host, user, password, dbname, port)

	DB, err = gorm.Open(postgres.Open(chainConnection), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database " + dbname)
	} else {
		fmt.Println("Connection success to database " + dbname)
	}

	DB.AutoMigrate(&User{})

	//DB.Create(&User{Name: "Avicii", Age: 13})

}

func main() {

	connectionDB()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
	}))

	e.GET("/", getUser)
	e.POST("/save", saveUser)
	e.DELETE("/delete/:id", deleteUser)
	e.PUT("/update/:id", updateUser)

	e.Logger.Fatal(e.Start(":" + port))

}

func getUser(c echo.Context) error {

	DB.Find(&users)

	return c.JSON(http.StatusOK, &users)
}

func saveUser(c echo.Context) error {

	u := new(User)

	if err := c.Bind(u); err != nil {
		return err
	}

	/* u.Id = seq
	u.CreatedAt = time.Now().UTC()
	users = append(users, *u)
	seq++ */

	DB.Create(u)

	return c.JSON(http.StatusCreated, &users)
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	/* for index, u := range users {
		if u.Id == id {
			users = removeElement(users, index)
		}
	} */

	DB.Delete(&User{}, id)

	return c.NoContent(http.StatusNoContent)
}

func updateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	u := new(User)

	if err := c.Bind(u); err != nil {
		return err
	}

	/* var myId int

	for index, u := range users {
		if u.Id == id {
			myId = index
		}
	}

	users[myId].Name = u.Name
	users[myId].Age = u.Age
	*/

	DB.Model(&User{}).Where("ID = ?", id).Updates(User{Name: u.Name, Age: u.Age})
	return c.NoContent(http.StatusNoContent)
}

/* func removeElement(array []User, position int) []User {
	return append(array[:position], array[position+1:]...)
} */
