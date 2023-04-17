package controller

import (
	"enigmacamp.com/uploadDownload/model"
	"enigmacamp.com/uploadDownload/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type StudentController struct {
	router                     *gin.Engine
	studentRegistrationUseCase usecase.StudentRegistrationUseCase
}

func (s *StudentController) createStudentHandler(c *gin.Context) {
	studentId := c.PostForm("id")
	file, fileHeader, err := c.Request.FormFile("photo")
	if err != nil {
		log.Println("Bad Request")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fileName := strings.Split(fileHeader.Filename, ".")
	if len(fileName) != 2 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fileExt := fileName[1]
	if strings.ToLower(fileExt) != "png" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	log.Println("studentId ", studentId)
	log.Println("file", fileHeader.Filename)
	log.Println("fileExt", fileExt)

	// File IO save file
	student := model.Student{
		Id: studentId,
	}
	errSaving := s.studentRegistrationUseCase.Register(&student, &file, fileExt)
	if errSaving != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, "File uploaded")
}

func (s *StudentController) getStudentProfileImageHandler(c *gin.Context) {
	id := c.Param("id")
	log.Println("Search the student by this id ", id)
	// Set mime types
	c.Header("Content-Type", "image/png")
	student, err := s.studentRegistrationUseCase.GetPhotoProfile(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.FileAttachment(student.PhotoProfile, fmt.Sprintf("profile-%s.png", id))
}
func NewStudentController(r *gin.Engine, studentRegistrationUseCase usecase.StudentRegistrationUseCase) *StudentController {
	controller := StudentController{router: r, studentRegistrationUseCase: studentRegistrationUseCase}
	r.POST("/student", controller.createStudentHandler)
	r.GET("/student-image/:id", controller.getStudentProfileImageHandler)
	return &controller
}
