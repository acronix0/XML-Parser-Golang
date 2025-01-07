package app

import (
	"context"
	"log/slog"
	"os"

	grpcapp "github.com/acronix0/XML-Parser-Golang/internal/app/grpc"
	"github.com/acronix0/XML-Parser-Golang/internal/config"
	"github.com/acronix0/XML-Parser-Golang/internal/database"
)

type App struct {
	serviceProvider *serviceProvider
	GRPCServer      *grpcapp.App
	log *slog.Logger
	config *config.Config
}

func NewApp(ctx context.Context) (*App, error) {
	a:= &App{config: config.Load()}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) Run() error{
	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error{
	inits := []func(context.Context) error {
		a.initLogger,
		a.initServiceProvider,
		a.InitMigrations,
		a.initGRPCServer,
	}
	for _, f:= range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	a.config = config.Load()
	return nil
}

func (a *App) InitMigrations(_ context.Context) error {
	return database.InitMigrations(
		a.config.MigrationsFolder,
		a.config.Database.Host,
		a.config.Database.UserName,
		a.config.Database.Name,
		a.config.Database.Password,
		a.config.Database.Port,
	)
}

func (a *App) initLogger(_ context.Context) error {
	var log *slog.Logger
	switch a.config.Env {
	case config.EnvLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	a.log = log
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.config, a.log)
	return nil
}


func (a *App) initGRPCServer(_ context.Context) error{
	a.GRPCServer = grpcapp.New(a.log,a.config.GRPC.Port, a.serviceProvider.ParserService())
	return nil
}

func (a *App) runGRPCServer() error{
	const op = "app.runGRPCServer"
	log := a.log.With(slog.String("op",op))
	log.Info("Start gRPC server")
	
	go func() {
		a.GRPCServer.Run()
	}()

	stop := make(chan os.Signal, 1)

	<-stop
	a.GRPCServer.Stop()
	log.Info("Server stopped")
	return nil
}