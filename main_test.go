package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv(MYSQL_DATABASE, "testing")
	go main()
	code := m.Run()
	os.Exit(code)
}
