package config_test

import (
	"io/fs"
	"log"
	"os"
	"path"
	"testing"

	"github.com/medymik/config"
)

func TestCanExportTheVariables(t *testing.T) {
	tmpFile, err := os.Create(".env_test")
	if err != nil {
		log.Fatalln(err)
	}

	envVar := []byte("APP_NAME=hello\nPORT=9898")
	if err = os.WriteFile(path.Join(path.Dir("."), tmpFile.Name()), envVar, fs.ModeAppend); err != nil {
		log.Fatalln(err)
	}
	config.Configure(tmpFile.Name())
	if os.Getenv("APP_NAME") != "hello" {
		t.Errorf("expect %v but got %v", "hello", os.Getenv("APP_NAME"))
	}

	if os.Getenv("PORT") != "9898" {
		t.Errorf("expect %v but got %v", "9898", os.Getenv("PORT"))
	}

	if err := os.Remove(tmpFile.Name()); err != nil {
		log.Fatalln(err)
	}
}
