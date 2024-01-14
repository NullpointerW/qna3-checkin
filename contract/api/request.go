package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Checkin(txh, via string) error {
	p := payload{
		"hash": txh,
		"via":  via,
	}
	m, err := p.marshal()
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", Checkin, m)
	c.setHeader(req, false)
	resp, err := c.h.Do(req)
	rb, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	var readMap map[string]any
	fmt.Println(string(rb))
	err = json.Unmarshal(rb, &readMap)
	if err != nil {
		return err
	}
	code := (readMap["statusCode"]).(float64)
	if code != http.StatusOK {
		err = errors.New((readMap["message"]).(string))
		return err
	}
	return nil
}

func (c *Client) ClaimStatus() (historyId, amount, nonce int, signature string, err error) {
	req, err := http.NewRequest("POST", ClaimStatus, nil)
	c.setHeader(req, true)
	resp, err := c.h.Do(req)
	if err != nil {
		return 0, 0, 0, "", err
	}
	rb, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return 0, 0, 0, "", err
	}
	var readMap map[string]any
	fmt.Println(string(rb))
	err = json.Unmarshal(rb, &readMap)
	if err != nil {
		return 0, 0, 0, "", err
	}
	data := (readMap["data"]).(map[string]any)
	//sign := (data["signature"]).(map[string]any)
	signature = ((data["signature"]).(map[string]any)["signature"]).(string)
	nonce = int(((data["signature"]).(map[string]any)["nonce"]).(float64))
	historyId = int((data["history_id"]).(float64))
	amount = int((data["amount"]).(float64))
	return
}

func (c *Client) ClaimAll(historyId int, signHash string) error {
	p := payload{
		"hash": signHash,
	}
	m, err := p.marshal()
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf(Claim, historyId), m)
	c.setHeader(req, false)
	resp, err := c.h.Do(req)
	if err != nil {
		return err
	}
	rb, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	var readMap map[string]any
	fmt.Println(string(rb))
	err = json.Unmarshal(rb, &readMap)
	if err != nil {
		return err
	}
	code := (readMap["statusCode"]).(float64)
	if code != http.StatusOK {
		err = errors.New((readMap["message"]).(string))
		return err
	}
	return nil
}
