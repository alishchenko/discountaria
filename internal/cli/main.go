package cli

import (
	"github.com/alecthomas/kingpin"
	"github.com/alishchenko/discountaria/internal/config"
	"github.com/alishchenko/discountaria/internal/server"
	"github.com/pkg/errors"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

var (
	app = kingpin.New("generator-svc", "service responsible for generating a book's pdf with a custom signature on it, handling status of uploading process, and storing tokens")

	// Run commands
	runCommand = app.Command("run", "run command")
	apiCommand = runCommand.Command("service", "run service")
	// Migration commands
	migrateCommand     = app.Command("migrate", "migrate command")
	migrateUpCommand   = migrateCommand.Command("up", "migrate database up")
	migrateDownCommand = migrateCommand.Command("down", "migrate database down")
)

func Run(args []string) bool {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env.Level)

	defer func() {
		if rvr := recover(); rvr != nil {
			log.Error("app panicked")
		}
	}()

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.Error(errors.Wrap(err, "failed to parse arguments").Error())
		panic(errors.Wrap(err, "failed to parse arguments"))
	}
	switch cmd {
	case apiCommand.FullCommand():
		log.Info("started api...")
		server.Run(cfg, log)
	case migrateUpCommand.FullCommand():
		err = MigrateUp(cfg)
	case migrateDownCommand.FullCommand():
		err = MigrateDown(cfg)
		// handle any custom commands here in the same way
	default:
		log.Error("unknown command %s", cmd)
		return false
	}
	if err != nil {
		if err != nil {
			log.Error(errors.Wrap(err, "failed to exec cmd").Error())
			return false
		}
	}

	return true
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
