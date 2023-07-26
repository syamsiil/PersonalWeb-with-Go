package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"personalWeb/connection"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Struct is a collection of properties or methods wrapped as a new data type
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
 
type User struct {
	Id      	 	int
	Name     		string
	Email    		string
	HashedPassword 	string
}


type SessionData struct {
	IsLogin bool
	Name    string
}

var userData = SessionData{}

func main() {
    e := echo.New()
	connection.DatabaseConnect()
	
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))


	e.Static("/public","public")

    // e = echo package
	// GET = the method
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

	// Authentication
	e.GET("/form-login", formLogin)
	e.POST("/login", login)
	e.GET("/logout", logout)

	
	e.GET("/form-register", formRegister)
	e.POST("/register", register)

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

	session, _ := session.Get("session", c)

	if session.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = session.Values["isLogin"].(bool)
		userData.Name = session.Values["name"].(string)
	}

	dataSession := map[string]interface{}{
		"dataSession": userData,
		"FlashStatus":  session.Values["status"],
		"FlashMessage": session.Values["message"],
		"FlashName":    session.Values["name"],
	}

	delete(session.Values, "message")
	session.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), dataSession)
}

func contact (c echo.Context)error{
	tmpl, err := template.ParseFiles("views/contact.html")

	if err !=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	session, _ := session.Get("session", c)

	if session.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = session.Values["isLogin"].(bool)
		userData.Name = session.Values["name"].(string)
	}

	dataSession := map[string]interface{}{
		"dataSession": userData,
	}

	return tmpl.Execute(c.Response(), dataSession)
}

func formAddProject (c echo.Context)error{
	tmpl, err := template.ParseFiles("views/add-project.html")

	if err !=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	session, _ := session.Get("session", c)

	if session.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = session.Values["isLogin"].(bool)
		userData.Name = session.Values["name"].(string)
	}

	dataSession := map[string]interface{}{
		"dataSession": userData,
	}

	return tmpl.Execute(c.Response(), dataSession)
}

func project (c echo.Context)error{
	tmpl, err := template.ParseFiles("views/project.html")

	if err !=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	databaseProjects, errProjects :=  connection.Conn.Query(context.Background(), "SELECT id, project_name, start_date, end_date, description, technologies, images FROM tb_project") //

	if errProjects != nil {
		return c.JSON(http.StatusInternalServerError, errProjects.Error())
	}

	var resultProjects []Project
	for databaseProjects.Next() {
		var each = Project{}

		err := databaseProjects.Scan(&each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Technologies, &each.Image)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		each.DistanceTime = calculateDuration(each.StartDate, each.EndDate)

		if checkValue(each.Technologies, "nodejs") { //must macthing with value in html
			each.NodeJs = true
		}
		if checkValue(each.Technologies, "reactjs") {
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
		session, _ := session.Get("session", c)

		projects := map[string]interface{}{
			"Projects":     resultProjects,
			"dataSession":  userData,
			"FlashStatus":  session.Values["status"],
			"FlashMessage": session.Values["message"],
			"FlashName":    session.Values["name"],
		}
	
		delete(session.Values, "message")
		session.Save(c.Request(), c.Response())
	
		return tmpl.Execute(c.Response(), projects)
	

	//Kode map[string]int maknanya adalah, tipe data map dengan key bertipe string dan value bertipe interface/any.
	
}

func testimonials (c echo.Context)error{
	tmpl, err := template.ParseFiles("views/testimonials.html")

	if err !=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	session, _ := session.Get("session", c)

	if session.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = session.Values["isLogin"].(bool)
		userData.Name = session.Values["name"].(string)
	}

	dataSession := map[string]interface{}{
		"dataSession": userData,
	}

	return tmpl.Execute(c.Response(), dataSession)
}

func detailProject (c echo.Context)error{
	id := c.Param("id")

	tmpl, err := template.ParseFiles("views/detail-project.html")

	if err !=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	idToInt, _ := strconv.Atoi(id)

	detailProject := Project {}

	// query get 1 data
	connection.Conn.QueryRow(context.Background(), "SELECT id, project_name, start_date, end_date, description, technologies, images FROM tb_project WHERE id=$1", idToInt).Scan(&detailProject.Id, &detailProject.ProjectName, &detailProject.StartDate, &detailProject.EndDate, &detailProject.Description, &detailProject.Technologies, &detailProject.Image)

	// fmt.Println("ini data detail project: ", errQuery)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	
	detailProject.DistanceTime = calculateDuration(detailProject.StartDate, detailProject.EndDate)
	
	if checkValue(detailProject.Technologies, "nodejs") { //must macthing with value in html
		detailProject.NodeJs = true
	}
	if checkValue(detailProject.Technologies, "reactjs") {
		detailProject.ReactJs = true
	}
	if checkValue(detailProject.Technologies, "javascript") {
		detailProject.Javascript = true
	}
	if checkValue(detailProject.Technologies, "html5") {
		detailProject.Html5 = true
	}

	session, _ := session.Get("session", c)

	if session.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = session.Values["isLogin"].(bool)
		userData.Name = session.Values["name"].(string)
	}

	data := map[string]interface{}{ 
		"Id":   id,
		"Project": detailProject,
		"StartDateString": detailProject.StartDate.Format("2006-01-02"),
		"EndDateString":   detailProject.EndDate.Format("2006-01-02"),
	}

	return tmpl.Execute(c.Response(),data)
}

func calculateDuration(startDate time.Time, endDate time.Time ) string {
	
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
	TechNodeJs := c.FormValue("input-nodejs")
	TechReactJs := c.FormValue("input-reactjs")
	TechJavascript := c.FormValue("input-javascript")
	TechHtml5 := c.FormValue("input-html5")
	

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_project (project_name, start_date, end_date, description, technologies[1], technologies[2], technologies[3], technologies[4], images) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", projectName, startDate, endDate, description, TechNodeJs, TechReactJs, TechJavascript, TechHtml5, "default.jpg")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message ": err.Error()})
	}
		
	fmt.Println("title: ", projectName)
	fmt.Println("start date: ", startDate)
	fmt.Println("end date: ", endDate)
	fmt.Println("description: ", description)
	// fmt.Println("distance time: ", distanceTime)
	fmt.Println("skill: ", TechNodeJs)
	fmt.Println("skill: ", TechReactJs)
	fmt.Println("skill: ", TechJavascript)
	fmt.Println("skill: ", TechHtml5)
	return c.Redirect(http.StatusMovedPermanently, "/project") 
}

func deleteProject(c echo.Context) error {
	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)
	// dataProjects = append(dataProjects[:idToInt], dataProjects[idToInt+1:]...)

	connection.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", idToInt)

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

	// query get 1 data
	connection.Conn.QueryRow(context.Background(), "SELECT id, project_name, start_date, end_date, description, technologies, images FROM tb_project WHERE id=$1", idToInt).Scan(&detailProject.Id, &detailProject.ProjectName, &detailProject.StartDate, &detailProject.EndDate, &detailProject.Description, &detailProject.Technologies, &detailProject.Image)

	// fmt.Println("ini data detail project: ", errQuery)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	
	detailProject.DistanceTime = calculateDuration(detailProject.StartDate, detailProject.EndDate)
	
	if checkValue(detailProject.Technologies, "nodejs") { //must macthing with value in html
		detailProject.NodeJs = true
	}
	if checkValue(detailProject.Technologies, "reactjs") {
		detailProject.ReactJs = true
	}
	if checkValue(detailProject.Technologies, "javascript") {
		detailProject.Javascript = true
	}
	if checkValue(detailProject.Technologies, "html5") {
		detailProject.Html5 = true
	}


	// for index, data := range dataProjects {
		
	// 	if index == idToInt { 
	// 		detailProject= Project{
	// 			ProjectName:    data.ProjectName,
	// 			StartDate:		data.StartDate,
	// 			EndDate: 		data.EndDate,
	// 			Description: 	data.Description,
	// 			DistanceTime: 	data.DistanceTime,
	// 			Javascript:     data.Javascript,
	// 			ReactJs:    	data.ReactJs,
	// 			NodeJs:			data.NodeJs,
	// 			Html5: 			data.Html5,
	// 		}
	// 	}
	// }

	data := map[string]interface{}{ 
		"Id":   id,
		"Project": detailProject,
		"StartDateString": detailProject.StartDate.Format("2006-01-02"),
		"EndDateString":   detailProject.EndDate.Format("2006-01-02"),
	}

	return tmpl.Execute(c.Response(),data)
}

func updatedProject(c echo.Context)error{
	
	id, _:= strconv.Atoi(c.Param("id"))
	projectName := c.FormValue("input-project-name")
	startDate := c.FormValue("input-start-date")
	endDate := c.FormValue("input-end-date")
	description := c.FormValue("input-description")
	TechNodeJs := c.FormValue("input-nodejs")
	TechReactJs := c.FormValue("input-reactjs")
	TechJavascript := c.FormValue("input-javascript")
	TechHtml5 := c.FormValue("input-html5")
	

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_project SET project_name=$1, start_date=$2, end_date=$3, description=$4, technologies[1]=$5, technologies[2]=$6, technologies[3]=$7, technologies[4]=$8, images=$9 WHERE id=$10", projectName, startDate, endDate, description, TechNodeJs, TechReactJs, TechJavascript, TechHtml5, "default.jpg", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message ": err.Error()})
	}
	
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
	// 	javascript = true

	// }
	// var html5 bool
	// if c.FormValue("input-html5") == "on" {
	// 	html5 = true
	// }

	// startTime, _ := time.Parse("2006-01-02", startDate)
	// endTime, _ := time.Parse("2006-01-02", endDate)

	//  updatedProject := Project{
	// 	ProjectName:    projectName,
	// 	StartDate:		startTime,
	// 	EndDate: 		endTime,
	// 	Description: 	description,
	// 	NodeJs: 		nodeJs,		
	// 	ReactJs: 		reactJs,		
	// 	Javascript: 	javascript,		
	// 	Html5:	 		html5,		
	// } 

	// dataProjects[id] = updatedProject

	return c.Redirect(http.StatusMovedPermanently, "/project")
}

func checkValue(slice []string, value string) bool {
	for _, data := range slice { //
		if data == value {
			return true
		}
	}
	return false
}

func formLogin (c echo.Context)error{
	// bikin pengecekan
	// ngambil dari session datanya, misalnya isLogin -> false
	// sess, _ := session.Get("session", c)

	// if sess.Values["isLogin"] != true {
	// 	return c.Redirect(http.StatusMovedPermanently, "/")
	// }
	session, _ := session.Get("session", c)

	messageFlash := map[string]interface{}{
		"FlashStatus":  session.Values["status"],
		"FlashMessage": session.Values["message"],
	}

	var tmpl, err = template.ParseFiles("views/login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	delete(session.Values, "status")
	delete(session.Values, "message")
	session.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), messageFlash)
}

func formRegister (c echo.Context)error{
	tmpl, err := template.ParseFiles("views/register.html")

	if err !=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(),nil)
}

// authentication and session

func redirectWithMessage(c echo.Context, message string, status bool, redirectPath string) error {
	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, redirectPath)
}


func register(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_user (name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)
	if err != nil {
		redirectWithMessage(c, "RegistrationFailed, please try again!", false, "/form-register")
	}

	return redirectWithMessage(c, "Registration Success", true, "/form-login")
}

func login(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	email := c.FormValue("email")
	password := c.FormValue("password")

	var user = User{}

	errEmail := connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_user WHERE email=$1", email).Scan(&user.Id, &user.Name, &user.Email, &user.HashedPassword)
	errPass := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))

	if errEmail != nil {
		return redirectWithMessage(c, "Email wrong!", true, "/form-login")
	}

	if errPass != nil {
		return redirectWithMessage(c, "Password wrong!", true, "/form-login")
	}

	session, _ := session.Get("session", c)
	session.Options.MaxAge = 3600
	session.Values["message"] = "login Success"
	session.Values["status"] = true // show alert
	session.Values["name"] = user.Name
	session.Values["id"] = user.Id
	session.Values["isLogin"] = true // access login
	session.Save(c.Request(), c.Response())

	return redirectWithMessage(c, "Login Succes", true, "/")
}

func logout(c echo.Context) error {
	session, _ := session.Get("session", c)
	session.Options.MaxAge = -1
	session.Values["isLogin"] = false
	session.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}
