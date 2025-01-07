package app

import (
	"log/slog"
	"os"

	"github.com/acronix0/XML-Parser-Golang/internal/config"
	"github.com/acronix0/XML-Parser-Golang/internal/database"
	"github.com/acronix0/XML-Parser-Golang/internal/repository"
	"github.com/acronix0/XML-Parser-Golang/internal/service"
	"github.com/acronix0/XML-Parser-Golang/internal/service/parser"
)

type serviceProvider struct {
	config *config.Config
	dataBase *database.Database
	logger *slog.Logger
	parserService service.Parser
	parserRepository repository.RepositoryManager
}

func newServiceProvider(cfg *config.Config, logger *slog.Logger) *serviceProvider {
	return &serviceProvider{ config: cfg, logger: logger}
}

func (s *serviceProvider) Config() *config.Config{
	return s.config
}

func (s *serviceProvider) Database() *database.Database{
	if s.dataBase == nil {
		cfg := s.Config()
		db, err := database.NewDatabase(
			cfg.Database.Port,
			cfg.Database.Host,
			cfg.Database.UserName,
			cfg.Database.Password,
			cfg.Database.Name,
		)
		if err != nil {
			panic(err)
		}
		s.dataBase = db
	}
	return s.dataBase
}

func (s *serviceProvider) Logger() *slog.Logger{
	if s.logger == nil {
		s.logger = slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelInfo},
			),
		)
	}
	return s.logger
}

func (s *serviceProvider) RepositoryManager() repository.RepositoryManager{
	if s.parserRepository == nil{
		s.parserRepository = repository.NewRepositoryManager(s.Database().GetDB())
	}
	return s.parserRepository
}

func (s *serviceProvider) ParserService() service.Parser {
	if s.parserService == nil {
    s.parserService = parser.NewParserService(
			s.logger, 
			s.RepositoryManager().Category(), 
			s.RepositoryManager().Product(),
			s.config.XMLConfig.FirstFilePath, 
			s.config.XMLConfig.SecondFilePath,
		)
  }

  return s.parserService

}