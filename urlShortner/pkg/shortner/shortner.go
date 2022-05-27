package shortner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"pkg/data"

	// "net/http" // REST framework
	"time"

	"github.com/gorilla/mux"
)

const (
	Alphabet   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	Prefix     = "http://localhost:8080/api/v1/"
	HashLength = 6
)

var cache = make(map[string]string)
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

//This is endpoint handler
func UrlShortner(w http.ResponseWriter, r *http.Request) {

	fmt.Print("Request param:", r.URL.Path)
	if r.URL.Path == "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Print("method type:", r.Method)
	switch r.Method {

	case "POST":
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "BAD Request.", http.StatusBadRequest)
		}

		shortUrl, err := Process(body)
		if err != nil {
			http.Error(w, "Couldn't process the provided Url", http.StatusBadRequest)
		}

		var response = data.Response{Url: shortUrl}
		var output []byte
		output, err = json.Marshal(response)
		if err != nil {
			http.Error(w, "Failed to convert.", http.StatusExpectationFailed)
		}

		w.Write([]byte(output))

	case "GET":
		param := r.URL.Path
		fmt.Print("param :", param)
		longUrl := cache[param]
		fmt.Print("longUrl:", longUrl)

	}
}

func GetFullUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recevied get request")
	// param := r.URL.Path
	vars := mux.Vars(r)
	param := vars["code"]
	fmt.Println("param:", param)
	longUrl := cache[param]
	fmt.Println("lognUrl", longUrl)
	http.Redirect(w, r, longUrl, http.StatusSeeOther)
}

// the processing of request happen here
func Process(body []byte) (url string, err error) {
	var input data.Request

	err = json.Unmarshal(body, &input)
	if err != nil && input.Url == "" {
		return
	}

	inputUrl := input.Url

	// check if alredy processed
	cacheValue := checkCache(inputUrl)
	if cacheValue != "" {
		url = cacheValue
		return
	}

	url = Shortner(inputUrl)
	// populate in cache
	cache[inputUrl] = url

	fmt.Printf("Input : %s Output: %s", inputUrl, url)

	return
}

// this is get the random string and append it to the short url home address(prefix).
func Shortner(url string) (sUrl string) {
	code := StringWithCharset(HashLength, Alphabet)
	sUrl = Prefix + code
	cache[code] = url
	return
}

// this is actully give the random string for the url
// TODO RFC?base32 encode (5 random char)?
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// check if the value already present in cache
// TODO REDIS
func checkCache(key string) (value string) {
	if cache[key] != "" {
		value = cache[key]

		return
	}
	return
}

// Ls7DyF : https://google.com/hello234324434324sasasa324324"
// https://google.com/hello234324434324sasasa324324 : https://tiny.com/Ls7DyF

// 10 mn req /day
// red-wri 10^1
// 10yrs live-run time
