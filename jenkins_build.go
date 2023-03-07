package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Empty struct {
}

type BuildPhases struct {
	Builds []*BuildPhase
}

type BuildPhase struct {
	Build  string
	Target string
	Brunch string
}

func main() {
	http.HandleFunc("/build", BuildTrigger)
	http.HandleFunc("/get_build", GetBuildTrigger)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var triggers []string

func BuildTrigger(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	text := r.FormValue("text")
	fmt.Printf("BuildTrigger [%v]", text)
	triggers = append(triggers, text)
	var empty Empty
	result, err := json.Marshal(empty)
	if err != nil {
		fmt.Printf("returnUser:get hero info error [%s]", err)
	}
	fmt.Println(string(result))
	w.Write(result)
}
func GetBuildTrigger(w http.ResponseWriter, r *http.Request) {
	var builds BuildPhases
	if len(triggers) > 0 {
		for i := 0; i < len(triggers); i++ {
			var buildPhase BuildPhase
			var stars = strings.Fields(triggers[i])
			for j := 0; j < len(stars); j++ {
				switch j {
				case 0:
					buildPhase.Build = stars[j]
				case 1:
					buildPhase.Target = stars[j]
				case 2:
					buildPhase.Brunch = stars[j]
				}
			}
			builds.Builds = append(builds.Builds, &buildPhase)
		}
		triggers = make([]string, 0)
	}
	result, err := json.Marshal(builds)
	if err != nil {
		fmt.Printf("returnUser:get hero info error [%s]", err)
	}
	fmt.Println(string(result))
	w.Write(result)
}

//http://192.168.11.3:8080/unity-climber-client-ios/buildWithParameters?token=yingyugang&brunch=develop
