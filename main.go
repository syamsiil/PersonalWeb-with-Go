package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

	e.Static("/public","public")

    // e = echo package
	// GET =  run the method
	// "/" = endpoint/routing ("localhost:5000 , ex. "/home")
	// helloWorld = function that will run if the route are opened
    e.GET("/", helloWorld)
	e.GET("/home", home)
	e.GET("/contact", contact)
	e.GET("/add-project", formAddProject)
	e.GET("/testimonials", testimonials)
	e.GET("/detail-project/:id", detailProject)

	e.POST("add-project", addProject)

 
    e.Logger.Fatal(e.Start("localhost:5000"))
}

func helloWorld(c echo.Context)error {
	return c.String(http.StatusOK, "Hello World")
}

func home (c echo.Context)error{
	tmpl, err := template.ParseFiles("views/index.html")

	if err !=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(),nil)
}

func contact (c echo.Context)error{
	tmpl, err := template.ParseFiles("views/contact.html")

	if err !=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(),nil)
}

func formAddProject (c echo.Context)error{
	tmpl, err := template.ParseFiles("views/add-project.html")

	if err !=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(),nil)
}

func testimonials (c echo.Context)error{
	tmpl, err := template.ParseFiles("views/testimonials.html")

	if err !=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(),nil)
}

func detailProject (c echo.Context)error{
	id := c.Param("id")

	tmpl, err := template.ParseFiles("views/detail-project.html")

	if err !=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	detailProject := map[string]interface{}{ // interface -> tipe data apapun
		"Id":      id,
		"Title":   "Dumbways ID memang keren",
		"Content": "Dumbways ID adalah bootcamp terbaik sedunia seakhirat!",
	}

	return tmpl.Execute(c.Response(),detailProject)
}

func addProject(c echo.Context)error{
	projectName := c.FormValue("input-project-name")
	starDate := c.FormValue("input-start-date")
	endDate := c.FormValue("input-end-date")
	description := c.FormValue("input-description")

	fmt.Println("title: ", projectName)
	fmt.Println("start date: ", starDate)
	fmt.Println("end date: ", endDate)
	fmt.Println("description: ", description)


	return c.Redirect(http.StatusMovedPermanently, "/add-project") 
}