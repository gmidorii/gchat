package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestWrite(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Errorf("failed pwd:%v", err)
	}
	path := filepath.Join(pwd, "test")
	f, err := os.Create(path)
	if err != nil {
		t.Errorf("failed create test file:%v", err)
	}
	defer func() {
		err := os.Remove(path)
		if err != nil {
			t.Error(err)
		}
	}()

	h := &History{
		file: f,
	}

	err = h.Write("test", "sender", "message")
	if err != nil {
		t.Errorf("failed test method: %v", err)
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("failed read test file: %v", err)
	}

	if string(b) != "[test] sender: message\n" {
		t.Errorf("not expected write log\n e:%s\n a:%s",
			"[test] sender: message\n",
			string(b))
	}
}
