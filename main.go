package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

type kobject int

func (k kobject) String() string {
	switch k {
	case deployment:
		return "deployment"
	case cronjob:
		return "cronjob"
	default:
		return "undefined"
	}
}

const (
	deployment kobject = iota
	cronjob
)

func printChangeCause(o kobject, r []string) error {
	for _, request := range r {
		out, err := exec.Command("kubectl", "get", o.String(), request, "-o", "jsonpath={.metadata.annotations.kubernetes\\.io/change-cause}").Output()
		if err != nil {
			return err
		}
		fmt.Println(string(out))
	}
	return nil
}
func main() {
	var d *cli.StringSlice
	var j *cli.StringSlice
	var deployments []string
	var cronjobs []string
	app := &cli.App{
		Name:  "Kubernetes Change Cause",
		Usage: "Prints the kubernetes.io/change-cause for the provided deployment",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:        "deployment",
				Aliases:     []string{"d"},
				Usage:       "deployment to request change cause for",
				Destination: d,
				Required:    false,
				Action: func(ctx *cli.Context, s []string) error {
					deployments = append(deployments, s...)
					return nil
				},
			},
			&cli.StringSliceFlag{
				Name:        "cronjob",
				Aliases:     []string{"j"},
				Usage:       "cronjob to request change cause for",
				Destination: j,
				Required:    false,
				Action: func(ctx *cli.Context, s []string) error {
					cronjobs = append(cronjobs, s...)
					return nil
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			if deployments == nil && cronjobs == nil {
				return errors.New("must provide at least one kubernetes object to request change cause")
			}
			if err := printChangeCause(deployment, deployments); err != nil {
				log.Printf("cannot print deployment change cause: %v", err)
			}
			if err := printChangeCause(cronjob, cronjobs); err != nil {
				log.Printf("cannot print cronjob change cause: %v", err)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
