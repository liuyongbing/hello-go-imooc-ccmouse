package infra

import (
	"io/ioutil"
	"net/http"
)

type Retriever struct {
}

func (Retriever) Get(url string) string {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic("ioutil.ReadAll err" + err.Error())
	}

	return string(bytes)
}
