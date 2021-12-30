package user

import (
	"encoding/json"
	"log"
	"strings"
)

type Action struct {
	Act   string
	Token string
}

type Ret struct {
	Act  string
	Data string
	Msg  string
}

func (this *Action) ActionHandler(data []byte) []Ret {
	// var result map[string]interface{}
	var action = &Action{}
	log.Printf(string(data))
	var err = json.Unmarshal(data, action)
	if err != nil {
		log.Printf("Error parsing %s: %v", data, err)
		var rets = []Ret{{Act: "msg", Data: "", Msg: "error command"}}
		return rets
	}
	this.Act = strings.TrimSpace(this.Act)
	log.Printf(action.Token)
	if !checkAction(action.Act) {
		var rets = []Ret{{Act: "msg", Data: "", Msg: "不要太調皮喔"}}
		return rets
	}
	if action.Token == "" {
		var auth = &Auth{}
		return auth.handle(strings.TrimSpace(action.Act))
	}

	if action.Act == "tokenLogin" && action.Token != "" {
		var auth = &Auth{}
		return auth.tokenLogin(action.Token)
	}
	log.Printf(action.Act)
	log.Printf(action.Token)
	var rets = []Ret{{Act: "msg", Data: "", Msg: "你太顽皮了"}}
	return rets
}

func checkAction(act string) bool {
	if act == "" {
		return false
	}
	var actionlist = []string{"help", "I", "login", "create", "tokenLogin"}
	actarr := strings.Fields(act)
	for _, eachItem := range actionlist {
		if eachItem == actarr[0] {
			return true
		}
	}
	return false
}
