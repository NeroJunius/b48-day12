package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Projects struct {
	ID                 int
	ProjectName        string
	Author             string
	StartDate          string
	EndDate            string
	Duration           string
	DescriptionProject string
	NodeJS             bool
	ReactJS            bool
	NextJS             bool
	TypeScript         bool
	Img                string
}

var dataProjects = []Projects{
	{
		ProjectName:        "The Wandering Knight",
		Author:             "N002",
		StartDate:          "2023-06-06",
		EndDate:            "2023-06-07",
		Duration:           "1 Month 1 Day",
		DescriptionProject: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin quis risus ut mi euismod sodales. Mauris id quam ut massa sodales faucibus consectetur sit amet dolor. ",
		NodeJS:             true,
		ReactJS:            true,
		NextJS:             true,
		TypeScript:         true,
	},

	{
		ProjectName:        "The Best True Damage",
		Author:             "Nero002",
		StartDate:          "2023-06-06",
		EndDate:            "2023-06-14",
		Duration:           "1 Month 8 Days",
		DescriptionProject: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin quis risus ut mi euismod sodales. Mauris id quam ut massa sodales faucibus consectetur sit amet dolor. ",
		NodeJS:             true,
		ReactJS:            true,
		NextJS:             true,
		TypeScript:         true,
	},

	{
		ProjectName:        "Karlan's Leader Skill",
		Author:             "Nero002",
		StartDate:          "2023-06-06",
		EndDate:            "2023-06-14",
		Duration:           "1 Month 8 Days",
		DescriptionProject: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin quis risus ut mi euismod sodales. Mauris id quam ut massa sodales faucibus consectetur sit amet dolor. ",
		NodeJS:             true,
		ReactJS:            true,
		NextJS:             true,
		TypeScript:         true,
	},
}

func main() {
	e := echo.New()
	e.Static("/assets", "assets")

	e.GET("/", Home)
	e.GET("/contactMe", contactMe)
	e.GET("/testimonial", testimonials)
	e.GET("/createProject", createProject)
	e.GET("/projectDetail/:id", projectDetail)
	e.GET("/editProject/:id", editProject)

	
	e.POST("/add-project", addProject)
	e.POST("/delete-project/:id", deleteProject)

	fmt.Println("server started on port 5900")
	e.Logger.Fatal(e.Start("localhost:5900"))
}

// List Fungsi GET Project nya /

func Home(c echo.Context) error {
	tmpl, err := template.ParseFiles("tabs/index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	projects := map[string]interface{}{
		"Project": dataProjects,
	}

	return tmpl.Execute(c.Response(), projects)
}

func contactMe(c echo.Context) error {
	var tmpl, err = template.ParseFiles("tabs/contact.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func createProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("tabs/project.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func testimonials(c echo.Context) error {
	var tmpl, err = template.ParseFiles("tabs/testimonial.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var tmpl, err = template.ParseFiles("tabs/project-detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	var listProjects = Projects{}
	for index, data := range dataProjects {
		if id == index {
			listProjects = Projects{
				ProjectName:        data.ProjectName,
				Author:             data.Author,
				StartDate:          data.StartDate,
				EndDate:            data.EndDate,
				Duration:           data.Duration,
				DescriptionProject: data.DescriptionProject,
				NodeJS:             data.NodeJS,
				ReactJS:            data.ReactJS,
				NextJS:             data.NextJS,
				TypeScript:         data.TypeScript,
			}
		}
	}

	data := map[string]interface{}{
		"Project": listProjects,
	}

	return tmpl.Execute(c.Response(), data)
}

func editProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("tabs/edit-form.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

// List Fungsi Post Project //

// time //
func countDuration(startDate, endDate string) string {
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

// buat project nya //
func addProject(c echo.Context) error {
	projectName := c.FormValue("projectName")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	duration := countDuration(endDate, startDate)
	descriptionProject := c.FormValue("projectDescription")

	var nodeJS bool
	if c.FormValue("nodeJS") == "yes" {
		nodeJS = true
	}
	var nextJS bool
	if c.FormValue("nextJS") == "yes" {
		nextJS = true
	}
	var reactJS bool
	if c.FormValue("reactJS") == "yes" {
		reactJS = true
	}
	var typeScript bool
	if c.FormValue("typeScript") == "yes" {
		typeScript = true
	}

	img := c.FormValue("imageProject")

	var createProject = Projects{
		ProjectName:        projectName,
		Author:             "Unknown",
		StartDate:          startDate,
		EndDate:            endDate,
		Duration:           duration,
		DescriptionProject: descriptionProject,
		NodeJS:             nodeJS,
		NextJS:             nextJS,
		ReactJS:            reactJS,
		TypeScript:         typeScript,
		Img:                img,
	}

	dataProjects = append(dataProjects, createProject)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

// delete project //
func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Index : ", id)

	dataProjects = append(dataProjects[:id], dataProjects[id+1:]...)

	fmt.Println("Berhasil menghapus project!")

	return c.Redirect(http.StatusMovedPermanently, "/")
}
