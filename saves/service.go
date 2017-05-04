package saves

import (
	"fmt"
	"github.com/pmdcosta/dark-souls"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"sync"
	"time"
)

// Ensure service implements interface.
var _ dark_souls.SaveService = &Service{}

// Service represents a service for managing saves.
type Service struct {
	client     *Client
	state      bool
	stateMutex sync.Mutex
}

// Save saves the game.
func (s *Service) Save() error {
	saveDir := filepath.Join(s.client.BackupPath, strconv.FormatInt(s.client.Now().Unix(), 10))
	if err := os.Mkdir(saveDir, os.ModePerm); err != nil {
		s.client.Logger.Log("err", ErrDirFailed, "path", saveDir)
		return err
	}
	files, _ := filepath.Glob(filepath.Join(s.client.SavePath, "*"))
	for _, path := range files {
		_, file := filepath.Split(path)
		s.copyFile(path, filepath.Join(saveDir, file))
	}
	return nil
}

// Save saves the game.
func (s *Service) Load() error {
	// get last save.
	dir, _ := filepath.Glob(filepath.Join(s.client.BackupPath, "*"))
	sort.Strings(dir)
	last := dir[len(dir)-1]

	// backup current.
	s.Save()

	// load last save.
	files, _ := filepath.Glob(filepath.Join(last, "*"))
	for _, path := range files {
		_, file := filepath.Split(path)
		s.copyFile(path, filepath.Join(s.client.SavePath, file))
	}

	return nil
}

// Save saves the game.
func (s *Service) Start() error {
	s.stateMutex.Lock()
	s.state = true
	s.stateMutex.Unlock()
	return nil
}

// Save saves the game.
func (s *Service) Stop() error {
	s.stateMutex.Lock()
	s.state = false
	s.stateMutex.Unlock()
	return nil
}

func (s *Service) copyFile(src, dst string) error {
	if _, err := os.Stat(src); err != nil {
		return err
	}
	if _, err := os.Stat(dst); err != nil && !os.IsNotExist(err) {
		return err
	}
	return s.copyFileContents(src, dst)
}

func (s *Service) copyFileContents(src, dst string) error {
	fmt.Printf("Copy from: %s to %s\n", src, dst)

	in, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer in.Close()

	// delete if exists.
	os.Remove(dst)

	// create new file.
	out, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		fmt.Println(err)
		return err
	}

	if err := out.Sync(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *Service) timer(interval int) {
	ticker := time.NewTicker(time.Second * time.Duration(interval))
	for range ticker.C {
		s.stateMutex.Lock()
		if s.state {
			s.Save()
		}
		s.stateMutex.Unlock()
	}
}
