package api

import (
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("xxx")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(client.authorization)
}

func TestCheckin(t *testing.T) {
	c, err := NewClient("xxx")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.Checkin("xxx", "bnb")
	if err != nil {
		fmt.Println(err)
		return
	}
}
