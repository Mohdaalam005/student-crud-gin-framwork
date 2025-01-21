package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var students = []student{
	{
		ID:   "1",
		Name: "mohd aalam",
		Age:  24,
	}, {
		ID:   "2",
		Name: "mohd ateek",
		Age:  43,
	}, {
		ID:   "3",
		Name: "mohd lateef",
		Age:  55,
	}, {
		ID:   "4",
		Name: "mohd aladeen",
		Age:  72,
	},
}

func main() {

	engine := gin.Default()
	engine.GET("/students", getAllStudent)
	engine.POST("/students", createStudent)
	engine.GET("/students/:id", getStudent)
	engine.PUT("/students/:id", updateStudent)
	engine.DELETE("/students/:id", deleteStudent)
	engine.Run("localhost:9000")
}

type student struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getAllStudent(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, students)
}
func createStudent(c *gin.Context) {
	var student student
	err := c.BindJSON(&student)
	if err != nil {
		return
	}
	students = append(students, student)
	c.IndentedJSON(http.StatusOK, student)
}

func getStudent(c *gin.Context) {
	id := c.Param("id")
	for _, student := range students {
		if student.ID == id {
			c.IndentedJSON(http.StatusOK, student)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})

}

func updateStudent(c *gin.Context) {
	var studentUpdate student
	err := c.BindJSON(&studentUpdate)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, student := range students {
		if student.ID == c.Param("id") {
			students[i].Name = studentUpdate.Name
			students[i].Age = studentUpdate.Age
			students[i].ID = studentUpdate.ID
		}
	}
	c.IndentedJSON(http.StatusOK, studentUpdate)
}

func deleteStudent(c *gin.Context) {
	for i, student := range students {
		if student.ID == c.Param("id") {
			students = append(students[:i], students[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}
