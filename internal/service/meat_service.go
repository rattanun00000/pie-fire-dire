package service

import (
	"io"
	"net/http"
	"time"
)

type MeatIpsumService struct {
	client *http.Client
}

func NewMeatIpsumService() *MeatIpsumService {
	return &MeatIpsumService{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (s *MeatIpsumService) FetchMeatIpsum() (string, error) {
	resp, err := s.client.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
