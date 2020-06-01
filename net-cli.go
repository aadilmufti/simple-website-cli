package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

var myFlags = []cli.Flag{
	&cli.StringFlag{
		Name:  "host",
		Value: "google.com",
	},
}

func info() {
	app.Name = "Simple Website Lookup CLI"
	app.Usage = "Query IPs, CNAMEs, MX, TXT records and Name Servers!"
}

func commands() {
	app.Commands = []*cli.Command{
		{
			Name:  "ns",
			Usage: "Looks up the Name Servers for a Particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up IP adresses for a Particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Looks up the CNAME of a Particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				fmt.Println(cname)
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Looks up MX records for a Particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i])
				}
				return nil
			},
		},
		{
			Name:  "txt",
			Usage: "Looks up TXT records for a Particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				txt, err := net.LookupTXT(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				for i := 0; i < len(txt); i++ {
					fmt.Println(txt[i])
				}
				return nil
			},
		},
	}
}

func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
