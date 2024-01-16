package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"qna3-checkin/contract"
)

const (
	Login       = "https://api.qna3.ai/api/v2/auth/login?via=wallet"
	ClaimStatus = "https://api.qna3.ai/api/v2/my/claim-all"
	Checkin     = "https://api.qna3.ai/api/v2/my/check-in"
	Claim       = "https://api.qna3.ai/api/v2/my/claim/%d"
)

type payload map[string]any

func (p payload) marshal() (io.Reader, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(b))
	return bytes.NewBuffer(b), nil
}

type Client struct {
	h             *http.Client
	authorization string
	ua            string
}

func NewClient(pk string) (*Client, error) {
	c := new(Client)
	c.h = new(http.Client)
	c.ua = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36"
	const msg = "AI + DYOR = Ultimate Answer to Unlock Web3 Universe"
	signed, addr, err := contract.SignMessage(pk, msg)
	if err != nil {
		return nil, err
	}
	p := payload{
		"wallet_address": addr,
		"signature":      signed,
	}
	m, err := p.marshal()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", Login, m)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.ua)
	resp, err := c.h.Do(req)
	if err != nil {
		return nil, err
	}
	rb, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var readMap map[string]any
	//fmt.Println(string(rb))
	err = json.Unmarshal(rb, &readMap)
	if err != nil {
		return nil, err
	}
	data := (readMap["data"]).(map[string]any)
	c.authorization = "Bearer " + (data["accessToken"]).(string)
	return c, nil
}

func (c *Client) setHeader(r *http.Request, nonJson bool) *http.Request {
	if !nonJson {
		r.Header.Set("Content-Type", "application/json")
	}
	r.Header.Set("Authorization", c.authorization)
	r.Header.Set("User-Agent", c.ua)
	return r
}
