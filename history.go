package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type Historier interface {
	Write(room, mamber, message string) error
	Close() error
}

type History struct {
	dir  string
	file *os.File
}

func NewHistory(dir string, name string) (Historier, error) {
	// mode append or create if not exist
	f, err := os.OpenFile(filepath.Join(dir, name),
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		0666,
	)
	if err != nil {
		return nil, err
	}
	return &History{
		dir:  dir,
		file: f,
	}, nil
}

func (l *History) Write(room, member, message string) error {
	out := fmt.Sprintf("[%s] %s: %s\n", room, member, message)
	_, err := l.file.WriteString(out)
	return err
}

func (l *History) Close() error {
	err := l.file.Close()
	if err != nil {
		return errors.Wrap(err, "falied log file closed")
	}
	log.Printf("%s log file is closed", l.file.Name())
	return nil
}
