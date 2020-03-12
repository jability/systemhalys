package systemhalys

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

const newLine = '\n'

type Config struct {
	data map[string]string
}

func (c *Config) Get(key string) (string, error) {
	if value, ok := c.data[key]; ok {
		return value, nil
	}
	return "", fmt.Errorf("error key not found in config")
}

func newConfig() *Config {
	return &Config{
		data: make(map[string]string, 32),
	}
}

func newLineScanner(r io.Reader) *bufio.Scanner {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	return s
}

func loadFromReader(r io.Reader) *Config {
	s := newLineScanner(r)
	c := newConfig()

	regex := regexp.MustCompile(`^\s*(?P<Key>[\w\d-_]+)\s+(?P<Value>[^\s]+)\s*$`)
	for s.Scan() {
		match := regex.FindStringSubmatch(s.Text())
		if match != nil {
			c.data[match[1]] = match[2]
		}
	}
	return c
}

func Load(r io.Reader) *Config {
	return loadFromReader(r)
}

// Creates a Config from a filesystem path, manages its lifecycle
func LoadFromFile(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Load(f), nil
}
