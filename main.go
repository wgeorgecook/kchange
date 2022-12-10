package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func main() {
	var d *cli.StringSlice
	var deployments []string
	app := &cli.App{
		Name:  "Kubernetes Change Cause",
		Usage: "Prints the kubernetes.io/change-cause for the provided deployment",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:        "deployment",
				Aliases:     []string{"d"},
				Usage:       "deployment to request change cause for",
				Destination: d,
				Required:    true,
				Action: func(ctx *cli.Context, s []string) error {
					if s == nil {
						return errors.New("deployment cannot be blank")
					}
					deployments = append(deployments, s...)
					return nil
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			for _, deployment := range deployments {
				out, err := exec.Command("kubectl", "get", "deployment", deployment, "-o", "jsonpath={.metadata.annotations.kubernetes\\.io/change-cause}").Output()
				if err != nil {
					return err
				}
				fmt.Println(string(out))
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
