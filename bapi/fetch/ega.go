package fetch

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"reflect"
	"strings"

	"encoding/json"
	"net/http"
	"net/url"
	neturl "net/url"

	"github.com/openbiox/butils/log"

	cnet "github.com/openbiox/butils/net"
)

type requestInfo struct {
	Data map[string]interface{}
}

type egaPostFields struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Password     string `json:"password"`
	Scope        string `json:"scope"`
	Username     string `json:"username"`
}

type egaTokenRet struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

var egaClientVer = "3.0.39"
var egaSessionID = randInt(10)
var authURL = "https://ega.ebi.ac.uk:8443/ega-openid-connect-server/token"
var entry = "https://ega.ebi.ac.uk:8051/elixir/data"

// Egafetch get EGA files
func Egafetch(ega, fileID, outDir string, opt *cnet.BnetParams) (err error) {
	client := cnet.NewHTTPClient(opt.Timeout, opt.Proxy)
	//	token, _ := getEgaToken(opt)
	token := `eyJraWQiOiJyc2ExIiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiJodWFuZ2p5QHNqdHUuZWR1LmNuIiwiYXpwIjoiZjIwY2QyZDMtNjgyYS00NTY4LWE1M2UtNDI2MmVmNTRjOGY0IiwiaXNzIjoiaHR0cHM6XC9cL2VnYS5lYmkuYWMudWs6ODQ0M1wvZWdhLW9wZW5pZC1jb25uZWN0LXNlcnZlclwvIiwiZXhwIjoxNTY3NjE1NDczLCJpYXQiOjE1Njc2MTE4NzcsImp0aSI6IjkyZjE2MDA0LTlmNjctNDVjZS05NjMxLWNlOTQyY2RjYjljOCJ9.BvIcsmdHQ49KJgBieFk5WySBkHx2GXmZs8JHVoQOuuyThilXOrdqkXQ190bny0r5UIyLDc-C7MPBeckwFs7tZRra6T3M32AtVFhUtcfFA70QTOtnViCNDpHOg5wbs-BUqP7byFIu7w6Et8wT-rl0bj0FtuWqA62lqtzu9_uH4JI`
	req, err := http.NewRequest("GET", entry+"/metadata/datasets", nil)
	req, err = http.NewRequest("GET", entry+fmt.Sprintf("/metadata/datasets/%s/files", ega), nil)
	req, err = http.NewRequest("GET", entry+fmt.Sprintf("/metadata/files/%s", fileID), nil)
	setEgaStdHeader(req)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		log.Warn(err)
		return err
	}
	io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close()
	return nil
}

func setEgaStdHeader(req *http.Request) {
	req.Header.Add("Client-Version", egaClientVer)
	req.Header.Add("Session-Id", egaSessionID)
}

// Egafetch get EGA files
func getEgaToken(opt *cnet.BnetParams) (token string, err error) {
	client := cnet.NewHTTPClient(opt.Timeout, opt.Proxy)
	egaJSON := egaPostFields{
		ClientID:     "f20cd2d3-682a-4568-a53e-4262ef54c8f4",
		GrantType:    "password",
		ClientSecret: "AMenuDLjVdVo4BSwi0QD54LL6NeVDEZRzEQUJ7hJOM3g4imDZBHHX0hNfKHPeQIGkskhtCmqAJtt_jm7EKq-rWw",
		Scope:        "openid",
	}
	egaJSON.Username = ""
	egaJSON.Password = ""

	DataURLVal := url.Values{}
	setDataURLVal(egaJSON, &DataURLVal)

	req, err := http.NewRequest("POST", authURL, strings.NewReader(DataURLVal.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	setEgaStdHeader(req)
	if err != nil {
		log.Warn(err)
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Warn(err)
		return "", err
	}
	ret, _ := ioutil.ReadAll(resp.Body)
	var egaRet = egaTokenRet{}
	json.Unmarshal(ret, &egaRet)
	if egaRet.Error == "" {
		log.Infof("AccessToken: %s; Expired after 1 hour.", egaRet.AccessToken)
	} else {
		log.Warn(egaRet.Error)
	}
	io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close()
	return egaRet.AccessToken, nil
}

func setDataURLVal(dat interface{}, DataURLVal *neturl.Values) {

	typ := reflect.TypeOf(dat)
	val := reflect.ValueOf(dat)

	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	num := val.NumField()
	for i := 0; i < num; i++ {
		DataURLVal.Add(typ.Field(i).Tag.Get("json"), val.Field(i).String())
	}
}

func randInt(length int) string {
	var letterRunes = []rune("0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
