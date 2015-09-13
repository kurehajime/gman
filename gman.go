// gman.go
package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
)

type searchResult struct {
	TotalCount int `json:"total_count"`
	Items      []struct {
		Name     string `json:"name"`
		FullName string `json:"full_name"`
	} `json:"items"`
}
type repo struct {
	Name          string `json:"name"`
	FullName      string `json:"full_name"`
	Description   string `json:"description"`
	DefaultBranch string `json:"default_branch"`
	HtmlUrl       string `json:"html_url"`
}

func Gman(text string) (string, error) {
	r, err := selRepo(text)
	if err != nil {
		return "", err
	}
	return getReadMe(r)
}
func OpenRepo(text string) (string,error) {
	r, err := selRepo(text)
	if err != nil {
		return "",err
	}
	url:=r.HtmlUrl
	//open web browser
	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", url).Start()
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	}	
	return r.HtmlUrl,nil
}
func ShowList(text string) (string, error) {
	var rtnStr string
	res, err := getRepoList(text)
	if err != nil {
		return "", err
	}
	for i := 0; i < len(res.Items); i++ {
		rtnStr += res.Items[i].FullName+"\n"
	}
	return rtnStr, nil
}
func selRepo(text string) (repo, error) {
	var full_name string
	var r repo
	if strings.Contains(text, "/") {
		full_name = text
	} else {
		res, err := getRepoList(text)
		if err != nil {
			return r, err
		}
		for i := 0; i < len(res.Items); i++ {
			if strings.ToUpper(res.Items[i].Name) == strings.ToUpper(text) {
				full_name = res.Items[i].FullName
				break
			}
		}
		if full_name == "" {
			return r, errors.New("Repository(" + text + ") not found ...")
		}
	}
	r, err := getRepo(full_name)
	if err != nil {
		return r, err
	}
	return r, nil
}

func getReadMe(r repo) (string, error) {
	var url = "https://raw.githubusercontent.com/" + r.FullName + "/" + r.DefaultBranch + "/README.md"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}
	readme, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(readme), nil
}
func getRepoList(text string) (searchResult, error) {
	var gURL string = "https://api.github.com/search/repositories?q="
	var res searchResult
	resp, err := http.Get(gURL + url.QueryEscape(text))
	if err != nil {
		return res, err
	}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}
func getRepo(full_name string) (repo, error) {
	var gURL string = "https://api.github.com/repos/"
	var res repo
	resp, err := http.Get(gURL + full_name)
	if err != nil {
		return res, err
	}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&res)
	if err != nil {
		return res, err
	}
	if res.FullName == "" {
		return res, errors.New("Repository(" + full_name + ") not found ...")
	}
	return res, nil
}
