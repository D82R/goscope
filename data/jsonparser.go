package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type Burpfile struct {
	Target struct {
		Scope struct {
			AdvancedMode bool `json:"advanced_mode"`
			Exclude      []struct {
				Enabled  bool   `json:"enabled"`
				File     string `json:"file"`
				Host     string `json:"host"`
				Port     string `json:"port"`
				Protocol string `json:"protocol"`
			} `json:"exclude, omitempty"`

			Include []struct {
				Enabled  bool   `json:"enabled"`
				File     string `json:"file"`
				Host     string `json:"host"`
				Port     string `json:"port"`
				Protocol string `json:"protocol"`
			} `json:"include, omitempty"`
		} `json:"scope"`
	} `json:"target"`
}

// ParseJson - Parse supplied Burp configuration file
func ParseJson(f string) ([]string, []string) {

	var inscope []string
	var outscope []string
	var scope Burpfile
	var r = strings.NewReplacer("^", "", `\`, "", "$", "")    // replace invalid domain characters
	var re = regexp.MustCompile(`\.\*\.[a-zA-Z]+\.[a-zA-Z]+`) // match format .*.example.com

	jsonFile, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &scope)

	// pull out "host:" value from json and remove invalid characters for inscope
	for i := 0; i < len(scope.Target.Scope.Include); i++ {
		host := r.Replace(scope.Target.Scope.Include[i].Host)
		if host == re.FindString(host) {
			host = strings.Replace(host, ".", "", 1)
		}
		inscope = append(inscope, host)
	}

	// same but for outscope
	for i := 0; i < len(scope.Target.Scope.Exclude); i++ {
		host := r.Replace(scope.Target.Scope.Exclude[i].Host)
		if host == re.FindString(host) {
			host = strings.Replace(host, ".", "", 1)
		}
		outscope = append(outscope, host)
	}
	return inscope, outscope
}

// CleanDomains - Take raw domains from JsonParse() and clean them to standard format
func CleanDomains(inscope, outscope []string) ([]string, []string) {

	inscopeKeys := make(map[string]bool) // store non-duplicate inscope domains
	inscopeCleaned := []string{}         // hold clean inscope domains

	outscopeKeys := make(map[string]bool) // store non-duplicate inscope domains
	outscopeCleaned := []string{}         // hold clean inscope domains

	// remove duplicate inscope domains
	for _, v := range inscope {
		if _, value := inscopeKeys[v]; !value {
			inscopeKeys[v] = true
			inscopeCleaned = append(inscopeCleaned, v)
		}
	}

	// print clean inscope domains to stdout
	fmt.Println("[In Scope]")
	for _, v := range inscopeCleaned {
		fmt.Println(v)
	}

	// remove duplicate outscope domains
	for _, i := range outscope {
		if _, value := outscopeKeys[i]; !value {
			outscopeKeys[i] = true
			outscopeCleaned = append(outscopeCleaned, i)
		}
	}

	// print clean inscope domains to stdout
	fmt.Println("\n[Out Scope]")
	for _, v := range outscopeCleaned {
		fmt.Println(v)
	}

	return inscopeCleaned, outscopeCleaned
}
