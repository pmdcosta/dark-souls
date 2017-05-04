package saves

import (
	"github.com/go-kit/kit/log"
	"github.com/pmdcosta/dark-souls"
	"os"
	"time"
)

const Timer int = 10

// Client represents a client used to manage save files.
type Client struct {
	Logger log.Logger

	// location of the target directory.
	SavePath   string
	BackupPath string

	// timer interval.
	Timer int

	// returns the current time.
	Now func() time.Time

	// services.
	saveService Service
}

// NewClient creates a new client.
func NewClient(l log.Logger, path, backup string, t int) *Client {
	c := &Client{
		Now:        time.Now,
		Logger:     l,
		SavePath:   path,
		BackupPath: backup,
		Timer:      Timer,
	}
	if t > 0 {
		c.Timer = t
	}
	c.saveService.client = c
	return c
}

// Open starts the client.
func (c *Client) Open() error {
	// checks if the directories exist.
	if _, err := os.Stat(c.SavePath); err != nil {
		c.Logger.Log("err", ErrDirNotFound, "path", c.SavePath)
		return err
	} else if _, err := os.Stat(c.BackupPath); err != nil {
		if err := os.Mkdir(c.BackupPath, os.ModePerm); err != nil {
			c.Logger.Log("err", ErrDirFailed, "path", c.BackupPath)
			return err
		}
	}
	c.saveService.state = false
	go c.saveService.timer(c.Timer)
	return nil
}

// Close stops the client.
func (c *Client) Close() error {
	return nil
}

// SaveService returns the service associated with the client.
func (c *Client) Service() dark_souls.SaveService { return &c.saveService }
