package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/mshto/team-project/src/config"
	"github.com/mshto/team-project/src/services"
	"github.com/urfave/negroni"
)

func main() {
	configFile := flag.String("config", "./team_project_config.json", "Configuration file in JSON-format")
	flag.Parse()

	if len(*configFile) > 0 {
		config.FilePath = *configFile
	}

	err := config.Load()
	if err != nil {
		log.Fatalf("error while reading config: %s", err)
	}

	f, err := os.OpenFile(config.Config.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Fatalf("error while close log file: %s", err)
		}
	}()

	log.SetOutput(f)

	// setting up web server middlewares
	middlewareManager := negroni.New()
	middlewareManager.Use(negroni.NewRecovery())
	middlewareManager.UseHandler(services.NewRouter())

	log.Println("Starting HTTP listener...")
	err = http.ListenAndServe(config.Config.ListenURL, middlewareManager)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Stop running application: %s", err)
}
