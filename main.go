package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type consultant struct {
	ID    string    `json:"id"`
	Name  string    `json:"name"`
	Role  string    `json:"role"`
	Since time.Time `json:"since"`
}

func getTime() time.Time {
	t := time.Now()
	return t
}

var consultants = []consultant{
	{ID: "1", Name: "Khizar Iqbal", Role: "Delivery Consultant", Since: getTime()},
	{ID: "2", Name: "Johnny Appleseed", Role: "Delivery Consultant", Since: getTime()},
}

func getAllConsultants(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, consultants)
}

func addConsultant(c *gin.Context) {
	var newConsultant consultant
	if err := c.BindJSON(&newConsultant); err != nil {
		return
	}
	t := time.Now()
	newConsultant.Since = t
	consultants = append(consultants, newConsultant)
	c.IndentedJSON(http.StatusCreated, consultants)
}

func removeConsultant(c *gin.Context) {
	id := c.Param("id")
	var i int
	for i = 1; i < len(consultants); i++ {
		if consultants[i].ID == id {
			break
		} else if (i == len(consultants)-1) && (consultants[i].ID != id) {
			i = -1
			break
		}
	}
	if i == -1 {
		fmt.Println("inside")
		c.IndentedJSON(http.StatusAccepted, "Already deleted")
		return
	}
	toDelete := consultants[i]
	copy(consultants[i:], consultants[i+1:])
	consultants = consultants[:len(consultants)-1]
	c.IndentedJSON(http.StatusAccepted, toDelete)

}

func home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Welcome to the api")
}

func main() {
	fmt.Println(consultants)
	app := gin.Default()
	app.GET("/", home)
	app.GET("/consultants", getAllConsultants)
	app.POST("/consultant", addConsultant)
	app.DELETE("/consultant/:id", removeConsultant)
	app.Run(":8080")
}
