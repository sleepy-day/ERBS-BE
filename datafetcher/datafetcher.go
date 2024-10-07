package datafetcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/time/rate"
)

type ERBSResponse[T any] struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      T      `json:"data,omitempty"`
	UserGames T      `json:"userGames,omitempty"`
	TopRanks  T      `json:"topRanks,omitempty"`
	UserRank  T      `json:"userRank,omitempty"`
	User      T      `json:"user,omitmepty"`
	UserStats T      `json:"userStats,omitempty"`
	Result    T      `json:"result,omitempty"`
}

type L10 struct {
	L10Path string `json:"l10Path"`
}

type JSONFile struct {
	Data      json.RawMessage `json:"data,omitempty"`
	UserGames json.RawMessage `json:"userGames,omitempty"`
	TopRanks  json.RawMessage `json:"topRanks,omitempty"`
}

type FieldInfo struct {
	LargestInt int64
	Flags      int
}

const (
	HashURL     = "https://open-api.bser.io/v2/data/hash"
	DataURL     = "https://open-api.bser.io/v2/data/"
	L10URL      = "https://open-api.bser.io/v1/l10n/"
	GameURL     = "https://open-api.bser.io/v1/games/40164089"
	TopRanksURL = "https://open-api.bser.io/v1/rank/top/27/3"

	StringFlag = 0x00000001
	IntFlag    = 0x00000010
	FloatFlag  = 0x00000100
	ArrayFlag  = 0x00001000
	MapVarFlag = 0x00010000
	MapArrFlag = 0x00100000
)

var (
	codeFile  *os.File
	apiKey    = ""
	limiter   = rate.NewLimiter(rate.Every(1200*time.Millisecond), 1)
	languages = []string{"English", "Korean", "Japanese", "ChineseSimplified", "ChineseTraditional"}
	miscURLs  = map[string]string{
		"Game":         "https://open-api.bser.io/v1/user/games/2694939",
		"TopRanks":     "https://open-api.bser.io/v1/rank/top/27/3",
		"Nickname":     "https://open-api.bser.io/v1/user/nickname?query=sterence",
		"Rank":         "https://open-api.bser.io/v1/rank/2694939/27/3",
		"UserStats":    "https://open-api.bser.io/v1/user/stats/2694939/27",
		"WeaponRoutes": "https://open-api.bser.io/v1/weaponRoutes/recommend",
	}

	client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:          10,
			IdleConnTimeout:       15 * time.Second,
			ResponseHeaderTimeout: 15 * time.Second,
			DisableKeepAlives:     false,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
)

func httpGet(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	return body, nil
}

func createHttpRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-key", apiKey)

	return req, nil
}

func makeHttpRequest[T any](url string) (*ERBSResponse[T], error) {
	req, err := createHttpRequest(url)
	if err != nil {
		return nil, err
	}

	var resp *http.Response

	for i := 0; i < 3; i++ {
		err = limiter.Wait(context.Background())
		if err != nil {
			return nil, err
		}

		resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusTooManyRequests {
			break
		}
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(rune(resp.StatusCode)))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var erbsResp ERBSResponse[T]
	err = json.Unmarshal(body, &erbsResp)
	if err != nil {
		return nil, err
	}

	return &erbsResp, nil
}

func getDataTypes() (map[string]int64, error) {
	req, err := makeHttpRequest[json.RawMessage](HashURL)
	if err != nil {
		return nil, err
	}

	jsonData, err := req.Data.MarshalJSON()
	if err != nil {
		return nil, err
	}

	var hashes map[string]int64
	err = json.Unmarshal(jsonData, &hashes)
	if err != nil {
		return nil, err
	}

	return hashes, nil
}

func getHashData() {
	hashes, err := getDataTypes()
	if err != nil {
		fmt.Printf("Unable to get data types: %s\n", err.Error())
		return
	}

	for name := range hashes {
		f, err := os.OpenFile(filepath.Join("data", "hash", name+".json"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			fmt.Printf("Error opening file %s.json err: %s\n", name, err.Error())
			return
		}
		defer f.Close()

		req, err := makeHttpRequest[json.RawMessage](DataURL + name)
		if err != nil {
			fmt.Printf("Error making request for type [%s] err: %s\n", name, err.Error())
			return
		}

		_, err = f.Write(req.Data)
		if err != nil {
			fmt.Printf("Error writing to file for type\n")
			return
		}

		f.Close()
	}
}

func getL10n() {
	for _, v := range languages {
		resp, err := makeHttpRequest[L10](L10URL + v)
		if err != nil {
			fmt.Printf("Error getting l10n data: %s\n", err.Error())
			return
		}

		data, err := httpGet(resp.Data.L10Path)
		if err != nil {
			fmt.Printf("Error getting l10n file: %s\n", err.Error())
			return
		}

		f, err := os.OpenFile(filepath.Join("data", "l10n", v+".txt"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			fmt.Printf("Error opening file %s.txt: %s\n", v, err.Error())
			return
		}
		defer f.Close()

		_, err = f.Write(data)
		if err != nil {
			fmt.Printf("Error writing file %s.txt: %s\n", v, err.Error())
			return
		}

		f.Close()
	}
}

func getMiscInfo() {
	for k, v := range miscURLs {
		resp, err := makeHttpRequest[json.RawMessage](v)
		if err != nil {
			fmt.Printf("Error getting URL [%s] : %s", v, err.Error())
			return
		}

		var raw json.RawMessage
		switch {
		case resp.Data != nil:
			raw = resp.Data
		case resp.UserGames != nil:
			raw = resp.UserGames
		case resp.TopRanks != nil:
			raw = resp.TopRanks
		case resp.User != nil:
			raw = resp.User
		case resp.UserRank != nil:
			raw = resp.UserRank
		case resp.UserStats != nil:
			raw = resp.UserStats
		case resp.Result != nil:
			raw = resp.Result
		}

		f, err := os.OpenFile(filepath.Join("data", "misc", k+".json"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			fmt.Printf("Error opening %s.json: %s", k, err.Error())
			return
		}
		defer f.Close()

		_, err = f.Write(raw)
		if err != nil {
			fmt.Printf("Error writing to %s.json: %s", k, err.Error())
			return
		}
	}
}

func FetchData(key string) {
	apiKey = key

	os.Mkdir(filepath.Join("data"), os.ModePerm)
	os.Mkdir(filepath.Join("data", "hash"), os.ModePerm)
	os.Mkdir(filepath.Join("data", "l10n"), os.ModePerm)
	os.Mkdir(filepath.Join("data", "misc"), os.ModePerm)

	var err error
	codeFile, err = os.OpenFile("datatypes.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Printf("Error opening datatypes.go file: %s", err.Error())
		return
	}

	getHashData()
	getL10n()
	getMiscInfo()
}
