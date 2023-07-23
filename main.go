package main

import (
	"context"
	"net/http"
	"personalWeb/connection"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id				int
	ProjectName 	string
	StartDate 		time.Time
	EndDate 		time.Time
	Description 	string
	DistanceTime	string
	Technologies 	[]string
	NodeJs			bool
	ReactJs			bool
	Javascript		bool
	Html5			bool
	Image			string
	Author			string
}


// Data - data yang ditampung, yang kemudian data yang diisi harus sesuai dengan tipe data yang telah dibangun pada struct 
var dataProjects = [] Project{}
// 	{
// 		ProjectName:    "Design Web Apps 2023",
		
// 		Description: 	"Description",
// 		DistanceTime: 	"1 month",
// 		Javascript:     true,
// 		ReactJs:    	true,
// 		NodeJs:			true,
// 		Html5: 			true,
// 	},
// 	{
// 		ProjectName:    "Mobile Apps 2023",
		
// 		Description: 	"Description 2",
// 		DistanceTime: 	"1 month",
// 		Javascript:     true,
// 		ReactJs:    	true,
// 		NodeJs:			true,
// 		Html5: 			true,
// 	},


func main() {
	connection.DatabaseConnect()

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
	e.GET("/update-project/:id", updateProject )

	e.POST("/add-project", addProject)
	e.POST("/delete-project/:id", deleteProject)
	e.POST("/update-project/:id", updatedProject)

 
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

	databaseProjects, errProjects :=  connection.Conn.Query(context.Background(), "SELECT project_name, start_date, end_date, description, technologies, images FROM tb_project")

	if errProjects != nil {
		return c.JSON(http.StatusInternalServerError, errProjects.Error())
	}

	var resultProjects []Project
	for databaseProjects.Next() {
		var each = Project{}

		err := databaseProjects.Scan(&each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Technologies, &each.Image)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		
		each.DistanceTime = calculateDuration(each.StartDate, each.EndDate)
		if checkValue(each.Technologies, "reactjs") {
			each.NodeJs = true
		}
		if checkValue(each.Technologies, "nodejs") {
			each.ReactJs = true
		}
		if checkValue(each.Technologies, "javascript") {
			each.Javascript = true
		}
		if checkValue(each.Technologies, "html5") {
			each.Html5 = true
		}

		resultProjects = append(resultProjects, each)
	}

	data := map[string]interface{}{
		"Projects": resultProjects,
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

	idToInt, _ := strconv.Atoi(id)

	detailProject := Project {}

	for index, data := range dataProjects {
		// index += 1
		if index == idToInt { // 1 == 0
			detailProject= Project{
				ProjectName:    data.ProjectName,
				StartDate:		data.StartDate,
				EndDate: 		data.EndDate,
				Description: 	data.Description,
				DistanceTime: 	data.DistanceTime,
				Javascript:     data.Javascript,
				ReactJs:    	data.ReactJs,
				NodeJs:			data.NodeJs,
				Html5: 			data.Html5,
			}
		}
	}

	data := map[string]interface{}{ // interface -> tipe data apapun
		"Id":   id,
		"Project": detailProject,
	}

	return tmpl.Execute(c.Response(),data)
}

func calculateDuration(startDate time.Time, endDate time.Time ) string {
	// startTime, _ := time.Parse("2006-01-02", startDate)
	// endTime, _ := time.Parse("2006-01-02", endDate)

	durationTime := int(endDate.Sub(startDate).Hours())
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
	// distanceTime := calculateDuration(startDate, endDate)

	var nodeJs bool
	if c.FormValue("input-nodejs") == "on" {
		nodeJs = true
	}
	var reactJs bool
	if c.FormValue("input-reactjs") == "on"  {
		reactJs = true
	}
	var javascript bool
	if c.FormValue("input-javascript") == "on" {
		javascript = true

	}
	var html5 bool
	if c.FormValue("input-html5") == "on" {
		html5 = true
	}

	startTime, _ := time.Parse("2006-01-02", startDate)
	endTime, _ := time.Parse("2006-01-02", endDate)

	 newProject := Project{
		ProjectName:    projectName,
		StartDate:		startTime,
		EndDate: 		endTime,
		Description: 	description,
		// DistanceTime: 	distanceTime,
		NodeJs: 		nodeJs,		
		ReactJs: 		reactJs,		
		Javascript: 	javascript,		
		Html5:	 		html5,		
		
	} 



	// append berfungsi untuk menambahkan data newProject ke dalam slice dataProject
	// mirip dengan fungsi push() pada javascript
	// param1 = dimana datanya ditampung
	// param2 = data apa yang akan ditampung
	

	dataProjects = append(dataProjects, newProject) // reassign / timpa

	// fmt.Println("title: ", projectName)
	// fmt.Println("start date: ", startDate)
	// fmt.Println("end date: ", endDate)
	// fmt.Println("description: ", description)
	// fmt.Println("distance time: ", distanceTime)
	// fmt.Println("skill: ", nodeJs)
	// fmt.Println("skill: ", reactJs)
	// fmt.Println("skill: ", javascript)
	// fmt.Println("skill: ", html5)


	return c.Redirect(http.StatusMovedPermanently, "/project") 
}

func deleteProject(c echo.Context) error {
	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)
	// append

	// slice -> 3 struct (+ 1 struct)

	// slice = append(slice, structlagi)

	// fmt.Println("persiapan delete index : ", id)
	dataProjects = append(dataProjects[:idToInt], dataProjects[idToInt+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/project")
}

func updateProject(c echo.Context)error{
	id := c.Param("id")

	tmpl, err := template.ParseFiles("views/update-project.html")

	if err !=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	idToInt, _ := strconv.Atoi(id)

	detailProject := Project {}

	for index, data := range dataProjects {
		// index += 1
		if index == idToInt { // 1 == 0
			detailProject= Project{
				ProjectName:    data.ProjectName,
				StartDate:		data.StartDate,
				EndDate: 		data.EndDate,
				Description: 	data.Description,
				DistanceTime: 	data.DistanceTime,
				Javascript:     data.Javascript,
				ReactJs:    	data.ReactJs,
				NodeJs:			data.NodeJs,
				Html5: 			data.Html5,
			}
		}
	}

	data := map[string]interface{}{ // interface -> tipe data apapun
		"Id":   id,
		"Project": detailProject,
	}

	return tmpl.Execute(c.Response(),data)
}

func updatedProject(c echo.Context)error{
	// id := c.Param("id")
	// idToInt, _ := strconv.Atoi(id)
	id, _:= strconv.Atoi(c.Param("id"))

	projectName := c.FormValue("input-project-name")
	startDate := c.FormValue("input-start-date")
	endDate := c.FormValue("input-end-date")
	description := c.FormValue("input-description")
	// distanceTime := calculateDuration(startDate, endDate)

	var nodeJs bool
	if c.FormValue("input-nodejs") == "on" {
		nodeJs = true
	}
	var reactJs bool
	if c.FormValue("input-reactjs") == "on"  {
		reactJs = true
	}
	var javascript bool
	if c.FormValue("input-javascript") == "on" {
		javascript = true

	}
	var html5 bool
	if c.FormValue("input-html5") == "on" {
		html5 = true
	}

	startTime, _ := time.Parse("2006-01-02", startDate)
	endTime, _ := time.Parse("2006-01-02", endDate)

	 updatedProject := Project{
		ProjectName:    projectName,
		StartDate:		startTime,
		EndDate: 		endTime,
		Description: 	description,
		// DistanceTime: 	distanceTime,
		NodeJs: 		nodeJs,		
		ReactJs: 		reactJs,		
		Javascript: 	javascript,		
		Html5:	 		html5,		
		
	} 

	dataProjects[id] = updatedProject

	return c.Redirect(http.StatusMovedPermanently, "/project")
}

func checkValue(slice []string, value string) bool {
	for _, data := range slice {
		if data == value {
			return true
		}
	}
	return false
}