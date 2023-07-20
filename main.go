package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	ProjectName 	string
	StartDate 		string
	EndDate 		string
	Description 	string
	DistanceTime	string
	PostDate 		string
	NodeJs			bool
	ReactJs			bool
	Javascript		bool
	Html5			bool
	Image			string
	Author			string
}


// Data - data yang ditampung, yang kemudian data yang diisi harus sesuai dengan tipe data yang telah dibangun pada struct 
var dataProjects = [] Project{
	{
		ProjectName:    "Design Web Apps 2023",
		StartDate:		"2023-06-01",
		EndDate: 		"2023-06-06",
		Description: 	"Description",
		DistanceTime: 	"1 month",
		PostDate: 		"20/07/2023",
		Javascript:     true,
		ReactJs:    	true,
		NodeJs:			true,
		Html5: 			true,
	},
	{
		ProjectName:    "Mobile Apps 2023",
		StartDate:		"2023-06-01",
		EndDate: 		"2023-06-06",
		Description: 	"Description 2",
		DistanceTime: 	"1 month",
		PostDate: 		"20/07/2023",
		Javascript:     true,
		ReactJs:    	true,
		NodeJs:			true,
		Html5: 			true,
	},
} 

func main() {
    e := echo.New()

	e.Static("/public","public")

    // e = echo package
	// GET =  run the method
	// "/" = endpoint/routing ("localhost:5000 , ex. "/home")
	// helloWorld = function that will run if the route are opened
    e.GET("/hello", helloWorld)
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/add-project", formAddProject)
	e.GET("/project", project)
	e.GET("/testimonials", testimonials)
	e.GET("/detail-project/:id", detailProject)

	e.POST("/add-project", addProject)

 
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

func project (c echo.Context)error{
	tmpl, err := template.ParseFiles("views/project.html")

	if err !=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	data := map[string]interface{}{
		"Projects": dataProjects,
	}

	return tmpl.Execute(c.Response(),data)
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

func calculateDuration(startDate, endDate string) string {
	startTime, _ := time.Parse("2006-01-02", startDate)
	endTime, _ := time.Parse("2006-01-02", endDate)

	durationTime := int(endTime.Sub(startTime).Hours())
	durationDays := durationTime / 24
	durationWeeks := durationDays / 7
	durationMonths := durationWeeks / 4
	durationYears := durationMonths / 12

	var duration string

	if durationYears > 1 {
		duration = strconv.Itoa(durationYears) + " years"
	} else if durationYears > 0 {
		duration = strconv.Itoa(durationYears) + " year"
	} else {
		if durationMonths > 1 {
			duration = strconv.Itoa(durationMonths) + " months"
		} else if durationMonths > 0 {
			duration = strconv.Itoa(durationMonths) + " month"
		} else {
			if durationWeeks > 1 {
				duration = strconv.Itoa(durationWeeks) + " weeks"
			} else if durationWeeks > 0 {
				duration = strconv.Itoa(durationWeeks) + " week"
			} else {
				if durationDays > 1 {
					duration = strconv.Itoa(durationDays) + " days"
				} else {
					duration = strconv.Itoa(durationDays) + " day"
				}
			}
		}
	}

	return duration
}

func addProject(c echo.Context)error{
	projectName := c.FormValue("input-project-name")
	startDate := c.FormValue("input-start-date")
	endDate := c.FormValue("input-end-date")
	description := c.FormValue("input-description")
	distanceTime := calculateDuration(startDate, endDate)

	// var nodeJs bool
	// if c.FormValue("input-nodejs") == "on" {
	// 	nodeJs = true
	// }
	// var reactJs bool
	// if c.FormValue("input-reactjs") == "on"  {
	// 	reactJs = true
	// }
	// var javascript bool
	// if c.FormValue("input-javascript") == "on" {
	// javascript = true

	// }
	// var html5 bool
	// if c.FormValue("input-html5") == "on" {
	// 	html5 = true
	// }else{
	// 	c.JSON(http.StatusBadRequest, map[string]string{"message": "Test"})
	// }

	 newProject := Project{
		ProjectName:    projectName,
		StartDate:		startDate,
		EndDate: 		endDate,
		Description: 	description,
		DistanceTime: 	distanceTime,
		// NodeJs: 		nodeJs,		
		// ReactJs: 		reactJs,		
		// Javascript: 	javascript,		
		// Html5:	 		html5,		
		
	} 

	// append berfungsi untuk menambahkan data newProject ke dalam slice dataProject
	// mirip dengan fungsi push() pada javascript
	// param1 = dimana datanya ditampung
	// param2 = data apa yang akan ditampung
	

	dataProjects = append(dataProjects, newProject) // reassign / timpa

	fmt.Println("title: ", projectName)
	fmt.Println("start date: ", startDate)
	fmt.Println("end date: ", endDate)
	fmt.Println("description: ", description)
	fmt.Println("distance time: ", distanceTime)
	// fmt.Println("skill: ", nodeJs)
	// fmt.Println("skill: ", reactJs)
	// fmt.Println("skill: ", javascript)
	// fmt.Println("skill: ", html5)


	return c.Redirect(http.StatusMovedPermanently, "/project") 
}