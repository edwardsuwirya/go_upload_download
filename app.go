package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Init Server
	serverHost := "0.0.0.0"
	serverPort := "8888"
	r := gin.Default()

	// Repository Temporary
	var student Student

	// Upload controller using POST form-data
	r.POST("/student", func(c *gin.Context) {
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
		fileLocation := filepath.Join("/Users/edwardsuwirya/Downloads/upload/", fileHeader.Filename)
		out, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer func(out *os.File) {
			err := out.Close()
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}(out)
		_, err = io.Copy(out, file)
		student = Student{
			Id:           studentId,
			PhotoProfile: fileLocation,
		}
		c.JSON(http.StatusOK, "File uploaded")
	})

	// Download controller using GET
	r.GET("/student-image/:id", func(c *gin.Context) {
		id := c.Param("id")
		log.Println("Search the student by this id ", id)
		// Set mime types
		c.Header("Content-Type", "image/png")
		c.FileAttachment(student.PhotoProfile, fmt.Sprintf("profile-%s.png", id))
	})

	// Run server
	host := fmt.Sprintf("%s:%s", serverHost, serverPort)
	err := r.Run(host)
	if err != nil {
		panic(err)
	}
}
