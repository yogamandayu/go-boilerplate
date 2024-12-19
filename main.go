package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"github.com/yogamandayu/go-boilerplate/internal/cmd"
	"github.com/yogamandayu/go-boilerplate/internal/config"
)

// @title go-boilerplate API
// @version 1.0
// @description go-boilerplate is an simple API for request and confirm OTP.
// @contact.name Yoga
// @contact.email yoga.grahadi@gmail.com
// @accept application/json
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file provided")
	}

	conf := config.NewConfig()
	conf.WithOptions(
		config.WithDBConfig(),
		config.WithRedisAPIConfig(),
		config.WithRedisWorkerNotificationConfig(),
		config.WithRESTConfig(),
		config.WithTelegramBotConfig(),
		config.WithRollbarConfig(),
	)

	cliApp := cli.NewApp()
	commands := cmd.NewCommand(conf).Commands()
	cliApp.Commands = commands
	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatalf("Unable to run CLI command, err: %v", err)
	}
}
