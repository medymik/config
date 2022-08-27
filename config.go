package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Configure(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		splitted := strings.Split(string(line), "=")
		if len(splitted) != 2 {
			return fmt.Errorf("error occured while parsing the env file")
		}
		os.Setenv(splitted[0], splitted[1])
	}
}
