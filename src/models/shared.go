package models

import "errors"

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []T    `json:"data,freeCharacters,userGames,topRanks,userRank,user,userStats,omitempty"`
	Next    int    `json:"next,omitempty"`
}

func (resp Response[any]) CheckResponse() error {
	if resp.Message == "Forbidden" {
		return errors.New("noauth")
	}

	if resp.Code == 429 {
		return errors.New("ratelimit")
	}

	if resp.Code != 200 && resp.Message != "Success" {
		return errors.New("unhandled")
	}

	return nil
}
