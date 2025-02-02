// Package cmd is a package that command the service.
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/yogamandayu/go-boilerplate/internal/interfaces/grpc"

	"github.com/yogamandayu/go-boilerplate/internal/app"
	"github.com/yogamandayu/go-boilerplate/internal/config"
	"github.com/yogamandayu/go-boilerplate/internal/interfaces/rest"
	"github.com/yogamandayu/go-boilerplate/pkg/db"
	"github.com/yogamandayu/go-boilerplate/pkg/slog"
	"github.com/yogamandayu/go-boilerplate/util"

	"github.com/jackc/tern/v2/migrate"
	"github.com/urfave/cli/v2"
)

// Command is a run service command.
type Command struct {
	conf *config.Config
}

// NewCommand is a constructor.
func NewCommand(conf *config.Config) *Command {
	return &Command{
		conf: conf,
	}
}

// Commands is get list of commands.
func (cmd *Command) Commands() cli.Commands {
	return []*cli.Command{
		{
			Name:    "http:rest",
			Aliases: []string{"r"},
			Usage:   "Run Server API",
			Action: func(cCtx *cli.Context) error {
				slogger := slog.NewSlog()
				a := app.NewApp().WithOptions(
					app.WithConfig(cmd.conf),
					app.WithSlog(slogger),
				)

				r := rest.NewServer(a)
				opts := []rest.Option{
					rest.SetByConfig(cmd.conf),
				}
				if err := r.With(opts...).Run(); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "http:grpc",
			Aliases: []string{"g"},
			Usage:   "Run gRPC API",
			Action: func(cCtx *cli.Context) error {
				slogger := slog.NewSlog()
				a := app.NewApp().WithOptions(
					app.WithConfig(cmd.conf),
					app.WithSlog(slogger),
				)

				r := grpc.NewServer(a)
				opts := []grpc.Option{
					grpc.SetByConfig(cmd.conf),
				}
				if err := r.With(opts...).Run(); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "db:migrate",
			Aliases: []string{"dbm"},
			Usage:   "Run database migration with tern",
			Action: func(cCtx *cli.Context) error {
				dbConn, err := db.NewConnection(cmd.conf.DB.Config)
				if err != nil {
					log.Fatal(err)
				}
				defer dbConn.Close()

				migrationsDir := os.DirFS(fmt.Sprintf("%s/internal/infrastructure/database/migrations", util.RootDir()))

				pgConn, err := dbConn.Acquire(cCtx.Context)
				if err != nil {
					return err
				}
				defer pgConn.Release()

				migrator, err := migrate.NewMigrator(cCtx.Context, pgConn.Conn(), "schema_version")
				if err != nil {
					log.Fatalf("Unable to create migrator: %v\n", err)
				}

				// Load migrations from the specified directory
				err = migrator.LoadMigrations(migrationsDir)
				if err != nil {
					log.Fatalf("Unable to load migrations: %v\n", err)
				}

				// Apply the migrations (Up)
				err = migrator.Migrate(cCtx.Context)
				if err != nil {
					log.Fatalf("Migration failed: %v\n", err)
				}

				log.Println("Migrations applied successfully!")

				return nil
			},
		},
		{
			Name:    "db:generate",
			Aliases: []string{"dbg"},
			Usage:   "Run database migration with tern",
			Action: func(cCtx *cli.Context) error {
				err := exec.Command("cd", util.RootDir()).Run()
				if err != nil {
					log.Fatal(err)
				}
				err = exec.Command("sqlc", "generate").Run()
				if err != nil {
					log.Fatal(err)
				}

				log.Println("Database generate successfully!")

				return nil
			},
		},
		{
			Name:    "git:pre-commit",
			Aliases: []string{"hooks"},
			Usage:   "Install pre-commit",
			Action: func(cCtx *cli.Context) error {
				err := exec.Command("cp", fmt.Sprintf("%s/.githooks/pre-commit", util.RootDir()), fmt.Sprintf("%s/.git/hooks/pre-commit", util.RootDir())).Run()
				if err != nil {
					log.Fatal(err)
				}
				err = exec.Command("chmod", "+x", fmt.Sprintf("%s/.git/hooks/pre-commit", util.RootDir())).Run()
				if err != nil {
					log.Fatal(err)
				}

				log.Println("Pre-commit installed!")
				return nil
			},
		},
	}
}
