package tumblr

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// APIURLBLOG Blog Method API URL.
	APIURLBLOG = "https://api.tumblr.com/v2/blog/%s"
)

// Tumblr struct
type Tumblr struct {
	Token          string
	TokenSecret    string
	BaseHost       string
	consumerKey    string
	consumerSecret string
	args           map[string]string
}

// NewTumblr new tumblr.
func NewTumblr(consumerKey string, consumerSecret string) *Tumblr {
	args := make(map[string]string)
	args["api_kei"] = consumerKey
	return &Tumblr{
		consumerKey:    consumerKey,
		consumerSecret: consumerSecret,
		args:           args,
	}
}

// HTTPPost post.
func (t Tumblr) HTTPPost(path string, data map[string]string) *http.Response {
	log.Println("API Path:", path)

	query := url.Values{}
	for key, val := range data {
		query.Set(key, val)
	}

	for key, val := range t.args {
		query.Set(key, val)
	}

	resp, err := http.PostForm(path, t.Sign("POST", path, query))

	//req, err := http.NewRequest("POST", path,
	//	strings.NewReader(t.Sign("POST", path, query)))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	////req.Header.Set("Authorization", "Bearea")
	////req.Header.Add("Authorization", t.args["api_kei"])

	//client := &http.Client{}

	////resp, err := http.PostForm(path, query)
	//log.Printf("%+v", req)
	//resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

// Post method
func (t Tumblr) Post(args map[string]string) *http.Response {
	args["type"] = "text"
	args["state"] = "queue"
	args["tags"] = "api,test"
	return t.HTTPPost(fmt.Sprintf(APIURLBLOG, t.BaseHost)+"/post", args)
}

// Sign sign all data.
func (t Tumblr) Sign(method, path string, args url.Values) url.Values {
	args.Set("oauth_version", "1.0")
	args.Set("oauth_signature_method", "HMAC-SHA1")
	args.Set("oauth_nonce", fmt.Sprint(time.Now().Unix()))
	args.Set("oauth_timestamp", fmt.Sprint(time.Now().Unix()))
	args.Set("oauth_consumer_key", t.consumerKey)
	args.Set("oauth_token", t.Token)

	h := hmac.New(sha1.New,
		[]byte(fmt.Sprintf("%s&%s", t.consumerSecret, t.TokenSecret)))

	sourceString := fmt.Sprintf("%s&%s&%s",
		method,
		url.QueryEscape(path),
		url.QueryEscape(strings.Replace(args.Encode(), "+", "%20", -1)))
	h.Write([]byte(sourceString))

	args.Set("oauth_signature", Base64Encode(h.Sum(nil)))
	return args
}

// Base64Encode encodes a value using base64.
func Base64Encode(value []byte) string {
	return base64.StdEncoding.WithPadding(base64.StdPadding).EncodeToString(value)
}
