package services

type AppService struct{}

var Version = "dev"

func NewAppService() *AppService {
	return &AppService{}
}

func (s *AppService) Version() string {
	return Version
}