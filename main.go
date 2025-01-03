package main

import (
    "fmt"
    "github.com/urfave/cli/v2"
    "log"
    "os"
)

func main() {
    app := &cli.App{
        Name:  "app",
        Usage: "check the website health",
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:     "url",
                Aliases:  []string{"u"},
                Usage:    "url to check",
                Required: true,
            },
            &cli.StringFlag{
                Name:    "port",
                Aliases: []string{"p"},
                Usage:   "port to check",
            },
        },
        Action: func(c *cli.Context) error {
            port := c.String("port")
            if port == "" {
                port = "80"
            }
            status := check(c.String("url"), port)
            fmt.Println(status)
            return nil
        },
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}