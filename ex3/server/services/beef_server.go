package services

import (
	"context"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type beefSerivceServer struct {
}

func NewBeefSerivceServer() BeefSerivceServer {
	return beefSerivceServer{}
}

func (beefSerivceServer) mustEmbedUnimplementedBeefSerivceServer() {}

func (beefSerivceServer) GetBeef(ctx context.Context, req *BeefRequest) (*BeefResponse, error) {
	data, _ := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	body, _ := io.ReadAll(data.Body)

	var mapBeef = make(map[string]int32)

	for _, b := range regexp.MustCompile(`\s+|\.|,`).Split(string(body), -1) {
		if b != "" {
			if _, ok := mapBeef[strings.ToLower(b)]; !ok {
				mapBeef[strings.ToLower(b)] = 1
			} else {
				mapBeef[strings.ToLower(b)]++
			}
		}
	}

	res := BeefResponse{
		Beef: mapBeef,
	}

	return &res, nil
}
