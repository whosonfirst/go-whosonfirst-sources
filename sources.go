package sources

import (
	"encoding/json"
	_ "errors"
	"github.com/whosonfirst/go-whosonfirst-sources/sources"
	"log"
)

type WOFSource struct {
	Id          int    `json:"id"`
	Fullname    string `json:"fullname"`
	Name        string `json:"name"`
	Prefix      string `json:"prefix"`
	Key         string `json:"key"`
	URL         string `json:"url"`
	License     string `json:"license"`
	Description string `json:"description"`
}

type WOFSourceSpecification map[string]WOFSource

var specification *WOFSourceSpecification

func init() {

	var err error

	specification, err = Spec()

	if err != nil {
		log.Fatal("Failed to parse specification", err)
	}
}

func Spec() (*WOFSourceSpecification, error) {

	var spec WOFSourceSpecification
	err := json.Unmarshal([]byte(sources.Specification), &spec)

	if err != nil {
		return nil, err
	}

	return &spec, nil
}

func IsValidSource(source string) bool {

	log.Println("Please write me")
	return false
}