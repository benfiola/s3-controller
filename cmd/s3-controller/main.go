package main

import (
	"context"
	"fmt"
	"os"

	"github.com/benfiola/s3-controller/internal/controller"
	"github.com/benfiola/s3-controller/internal/info"
	"github.com/benfiola/s3-controller/internal/logging"
	"github.com/urfave/cli/v3"
)

func main() {
	cli.VersionPrinter = func(cmd *cli.Command) {
		fmt.Fprint(cmd.Root().Writer, cmd.Root().Version)
	}

	command := &cli.Command{
		Action: func(ctx context.Context, c *cli.Command) error {
			healthAddress := c.String("health-address")
			kubeconfig := c.String("kubeconfig")
			leaderElection := c.Bool("leader-election")
			logFormat := c.String("log-format")
			logLevel := c.String("log-level")
			metricsAddress := c.String("metrics-address")

			logger, err := logging.New(&logging.Opts{Format: logFormat, Level: logLevel})
			if err != nil {
				return err
			}

			sctx := logging.WithLogger(ctx, logger)

			controller, err := controller.New(&controller.Opts{
				HealthAddress:  healthAddress,
				Kubeconfig:     kubeconfig,
				LeaderElection: leaderElection,
				MetricsAddress: metricsAddress,
			})
			if err != nil {
				return err
			}

			return controller.Run(sctx)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "health-address",
				Sources: cli.EnvVars("HEALTH_ADDRESS"),
			},
			&cli.StringFlag{
				Name:    "kubeconfig",
				Sources: cli.EnvVars("KUBECONFIG"),
			},
			&cli.BoolFlag{
				Name:    "leader-election",
				Sources: cli.EnvVars("LEADER_ELECTION"),
			},
			&cli.StringFlag{
				Name:    "log-format",
				Sources: cli.EnvVars("LOG_FORMAT"),
				Value:   "text",
			},
			&cli.StringFlag{
				Name:    "log-level",
				Sources: cli.EnvVars("LOG_LEVEL"),
				Value:   "info",
			},
			&cli.StringFlag{
				Name:    "metrics-address",
				Sources: cli.EnvVars("METRICS_ADDRESS"),
			},
		},
		Version: info.Version,
	}

	err := command.Run(context.Background(), os.Args)
	code := 0
	if err != nil {
		fmt.Printf("command failed, error: %v", err)
		code = 1
	}
	os.Exit(code)
}
