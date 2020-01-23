package webserver

import (
	"github.com/gin-contrib/sessions"
	"log"
	"net/http"
	"strconv"

	shiprepdb "github.com/deeper-x/shipreporting-platform/db"
	"github.com/deeper-x/shipreporting-platform/password"
	"github.com/gin-gonic/gin"
)

// ShowCustomers todo context
func ShowCustomers(c *gin.Context) {
	db := shiprepdb.Connector()

	defer db.Close()

	c.HTML(http.StatusOK, "show_customers", gin.H{
		"message": "hi",
	})
}

// FormAddCustomer todo context
func FormAddCustomer(c *gin.Context) {
	c.HTML(http.StatusOK, "form_add_customer", gin.H{
		"message": "hi",
	})
}

// AddCustomer todo context
func AddCustomer(c *gin.Context) {
	session := sessions.Default(c)
	db := shiprepdb.Connector()

	defer db.Close()

	hashed, err := password.HashPassword(c.PostForm("password"))

	if err != nil {
		log.Println(err)
	}

	strPortinformer := (session.Get("managedPortinformer")).(string)

	intPortinformer, err := strconv.Atoi(strPortinformer)

	if err != nil {
		log.Println(err)
	}

	customer := shiprepdb.Customer{Email: c.PostForm("email"), Password: hashed, Portinformer_1: intPortinformer}

	err = db.Insert(&customer)

	if err != nil {
		log.Println(err)
	}

	c.JSON(200, gin.H{
		"message": "ok",
	})
}

// DelCustomer todo description
func DelCustomer(c *gin.Context) {
	db := shiprepdb.Connector()

	defer db.Close()

	id := c.Param("id")

	param, _ := strconv.Atoi(id)

	customer := shiprepdb.Customer{ID: param}

	db.Delete(&customer) // Delete it permanently

	c.JSON(200, gin.H{
		"message": "ok",
	})
}

// AllCustomers todo description
func AllCustomers(c *gin.Context) {
	db := shiprepdb.Connector()
	session := sessions.Default(c)

	defer db.Close()

	idPortinformer := (session.Get("managedPortinformer")).(string)

	intPortinformer, err := strconv.Atoi(idPortinformer)

	if err != nil {
		log.Println(err)
	}

	customers := []shiprepdb.Customer{}

	err = db.Model(&customers).Where("portinformer_1 = ?", intPortinformer).Select()

	if err != nil {
		log.Println(err)
	}

	log.Println(customers)

	c.HTML(200, "all_customers", gin.H{
		"all_customers": customers,
	})
}
