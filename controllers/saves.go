package controllers

import (
	"github.com/go-kit/kit/log"
	"github.com/goadesign/goa"
	"github.com/pmdcosta/dark-souls"
	"github.com/pmdcosta/dark-souls/app"
)

// SavesController implements the saves resource.
type SavesController struct {
	*goa.Controller
	logger      log.Logger
	saveService dark_souls.SaveService
}

// NewSavesController creates a saves controller.
func NewSavesController(service *goa.Service, ss dark_souls.SaveService, logger log.Logger) *SavesController {
	return &SavesController{Controller: service.NewController("SavesController"), logger: logger, saveService: ss}
}

// Load runs the load action.
func (c *SavesController) Load(ctx *app.LoadSavesContext) error {
	// SavesController_Load: start_implement

	c.saveService.Load()

	// SavesController_Load: end_implement
	return nil
}

// Save runs the save action.
func (c *SavesController) Save(ctx *app.SaveSavesContext) error {
	// SavesController_Save: start_implement

	c.saveService.Save()

	// SavesController_Save: end_implement
	return nil
}

// Start runs the start action.
func (c *SavesController) Start(ctx *app.StartSavesContext) error {
	// SavesController_Start: start_implement

	c.saveService.Start()

	// SavesController_Start: end_implement
	return nil
}

// Stop runs the stop action.
func (c *SavesController) Stop(ctx *app.StopSavesContext) error {
	// SavesController_Stop: start_implement

	c.saveService.Stop()

	// SavesController_Stop: end_implement
	return nil
}
