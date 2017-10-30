package lassie

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var (
	addr  = flag.String("addr", DefaultAddr, "address")
	token = flag.String("token", "", "token")
)

func TestMain(m *testing.M) {
	flag.Parse()
	if *token == "" {
		if _, err := NewWithAddr(*addr, ""); err != nil {
			fmt.Println("Error creating client:", err)
			fmt.Println("You might need to specify a token when running the tests against", *addr)
			fmt.Println("Run `go test -args -token <your-api-token>`.")
			os.Exit(1)
		}
	}

	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	_, err := NewWithAddr(*addr, *token)
	if err != nil {
		t.Fatal(err)
	}
}
