package delivery

import (
	"enigmacamp.com/uploadDownload/delivery/controller"
	"enigmacamp.com/uploadDownload/repository"
	"enigmacamp.com/uploadDownload/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	useCase usecase.StudentRegistrationUseCase
	engine  *gin.Engine
	host    string
}

func (s *Server) initController() {
	controller.NewStudentController(s.engine, s.useCase)
}

func (s *Server) Run() {
	s.initController()

	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}
func NewServer() *Server {
	baseFilePath := "/Users/edwardsuwirya/Downloads/upload/"
	serverHost := "0.0.0.0"
	serverPort := "8888"

	fileRepo := repository.NewFileRepository(baseFilePath)
	studentRepo := repository.NewStudentRepository()
	useCase := usecase.NewStudentRegistrationUseCase(fileRepo, studentRepo)

	r := gin.Default()

	return &Server{
		useCase: useCase,
		engine:  r,
		host:    fmt.Sprintf("%s:%s", serverHost, serverPort),
	}
}
