package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	yaclik "./lib"
)

type Env struct {
	Domain   string
	AppID    string
	Login    string
	Password string
	Format   string
}

func main() {
	var env Env
	flag.StringVar(&env.Domain, "d", "", "Domain name")
	flag.StringVar(&env.AppID, "a", "0", "App ID")
	flag.StringVar(&env.Login, "n", "", "User login name")
	flag.StringVar(&env.Password, "p", "", "User login password")
	flag.StringVar(&env.Format, "o", "csv", "Output format")

	flag.Parse()

	if flag.Parsed() {
		if len(os.Args) == 0 || env.AppID == "0" || env.Domain == "" || env.Login == "" {
			os.Exit(1)
		}

		ba, _ := yaclik.FetchFieldsJson(env.AppID, env.Domain, env.Login, env.Password)
		if env.Format == "json" {
			var out bytes.Buffer
			if err := json.Indent(&out, ba, "", "  "); err != nil {
				log.Fatal("Error json indent", err)
			}
			out.WriteTo(os.Stdout)
		} else {
			yaclik.ParseFieldsJSON(&ba, os.Stdout)
		}
	} else {
		fmt.Println("flag parse error")
		os.Exit(1)
	}
}
