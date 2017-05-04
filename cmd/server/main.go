//go:generate goagen bootstrap -d github.com/pmdcosta/dark-souls/design

package main

import (
	"flag"
	"github.com/go-kit/kit/log"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/kit"
	"github.com/goadesign/goa/middleware"
	"github.com/mitchellh/go-homedir"
	"github.com/pmdcosta/dark-souls/app"
	"github.com/pmdcosta/dark-souls/controllers"
	"github.com/pmdcosta/dark-souls/saves"
	"os"
	"path/filepath"
)

func main() {
	// Create service
	service := goa.New("souls")

	// create logger.
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	service.WithLogger(goakit.New(logger))

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, false))
	service.Use(middleware.Recover())

	// find home path.
	home, _ := homedir.Dir()

	// parse input flags.
	var (
		savePath = flag.String("save", filepath.Join(home, "AppData/Roaming/DarkSoulsIII/0110000100d5c663"), "Path to save directory.")
		bkPath   = flag.String("bk", filepath.Join(home, "AppData/Roaming/DarkSoulsIII/backup"), "Path to backup directory.")
		timer    = flag.Int("timer", 60, "Interval between saves in seconds.")
	)
	flag.Parse()

	// Start saves client.
	savesClient := saves.NewClient(logger, *savePath, *bkPath, *timer)
	if err := savesClient.Open(); err != nil {
		panic(err)
	}
	defer savesClient.Close()

	// Mount "saves" controller
	c := controllers.NewSavesController(service, savesClient.Service(), logger)
	app.MountSavesController(service, c)

	// Mount "public" controller
	cp := controllers.NewPublicController(service)
	app.MountPublicController(service, cp)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
