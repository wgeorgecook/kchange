package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

type ktype string

const (
	deployment ktype = "deployment"
	cronjob    ktype = "cronjob"
)

type kobject struct {
	kind ktype
	name string
}

func addToRequests(k ktype, s []string, r []kobject) []kobject {
	var ks []kobject
	for _, t := range s {
		ko := kobject{kind: k, name: t}
		ks = append(ks, ko)
	}

	r = append(r, ks...)
	return r
}

func printChangeCause(r []kobject) error {
	for _, o := range r {
		out, err := exec.Command("kubectl", "get", string(o.kind), o.name, "-o", "jsonpath={.metadata.annotations.kubernetes\\.io/change-cause}").Output()
		if err != nil {
			return fmt.Errorf("cannot print %s %s change cause: %v", o.kind, o.name, err)
		}
		fmt.Println(string(out))
	}
	return nil
}

func main() {
	var d *cli.StringSlice
	var j *cli.StringSlice
	var requests []kobject
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
					requests = addToRequests(deployment, s, requests)
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
					requests = addToRequests(cronjob, s, requests)
					return nil
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			if err := printChangeCause(requests); err != nil {
				log.Println(err)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
