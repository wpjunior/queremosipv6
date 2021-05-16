package main

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"sort"
	"strings"

	"github.com/antchfx/htmlquery"
)

type site struct {
	Domain string `json:"domain"`
	IPV6   bool   `json:"ipv6"`
}

func main() {
	err := compile("BR")
	if err != nil {
		log.Fatal(err)
	}

	err = compile("Global")
	if err != nil {
		log.Fatal(err)
	}
}

func compile(origin string) error {
	top50Brazil, err := getAlexaWebsites(origin)
	if err != nil {
		return err
	}

	result, err := computeIPV6(top50Brazil)
	if err != nil {
		return err
	}

	brasil, err := os.Create(origin + "-status.json")
	if err != nil {
		return err
	}

	err = json.NewEncoder(brasil).Encode(result)
	if err != nil {
		return err
	}

	return nil
}

func computeIPV6(sites []string) ([]site, error) {
	result := make([]site, len(sites))

	for i, site := range sites {
		result[i].Domain = site
	}

	for i, site := range sites {
		conn, err := net.Dial("tcp6", site+":80")

		if err != nil {
			continue
		}

		result[i].IPV6 = true
		conn.Close()
	}

	return result, nil
}

func getAlexaWebsites(origin string) ([]string, error) {
	filePath := origin + ".html"
	doc, err := htmlquery.LoadDoc(filePath)

	if err != nil {
		return nil, err
	}

	nodes, err := htmlquery.QueryAll(doc, "//div[@class=\"tr site-listing\"]")
	if err != nil {
		return nil, err
	}

	result := []string{}
	for _, node := range nodes {
		subNode, err := htmlquery.Query(node, "//a")
		if err != nil {
			return nil, err
		}

		result = append(result, strings.ToLower(subNode.FirstChild.Data))
	}

	sort.Strings(result)

	return result, nil
}
