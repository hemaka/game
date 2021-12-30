package user

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"game/common/db"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var langFile = "./assets/cn/auth.json"

type Auth struct {
}

type AuthTxt struct {
	Welcome string
	Help    string
}

func (this *AuthTxt) Load() {
	data, err := ioutil.ReadFile(langFile)

	if err == nil {
		err = json.Unmarshal(data, this)
	}

	if err != nil {
		log.Printf("Error parsing %s: %v", langFile, err)
	}

	executable, err := os.Executable()
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(langFile); err != nil {
		executablePath, err := filepath.Abs(filepath.Dir(executable))
		if err != nil {
			panic(err)
		}

		log.Printf("Going to folder %v...", executablePath)

		os.Chdir(executablePath)
	}
}

func (this *Auth) handle(act string) []Ret {
	// var rds = server.RedisRun()
	var authTxt = AuthTxt{}
	authTxt.Load()
	if act == "help" {
		var rets = []Ret{{Act: "msg", Data: "", Msg: authTxt.Help}}
		return rets
	}
	actArr := strings.Fields(act)

	switch actArr[0] {
	case "create":
		if len(actArr) != 3 {
			var rets = []Ret{{Act: "msg", Data: "", Msg: "參數錯誤"}}
			return rets
		}
		return this.createPlayer(actArr[1], actArr[2])
	case "login":
		if len(actArr) != 3 {
			var rets = []Ret{{Act: "msg", Data: "", Msg: "參數錯誤"}}
			return rets
		}
		return this.loginPlayer(actArr[1], actArr[2])
	default:
		var rets = []Ret{{Act: "msg", Data: "", Msg: authTxt.Welcome}}
		return rets
	}
}

func (this *Auth) createPlayer(username string, password string) []Ret {
	var db = db.MysqlRun()
	var player Player
	result := db.First(&player, "username = ?", username)

	if result.RowsAffected > 0 {
		var rets = []Ret{{Act: "msg", Data: "", Msg: "用戶名已被使用"}}
		return rets
	}
	var pwd = makePassword(username, password)
	p := Player{Username: username, Password: pwd, Token: uuid.New().String()}
	db.Create(&p)
	msg, _ := json.Marshal(map[string]string{"action": "create", "token": p.Token, "msg": "注册成功"})
	var rets = []Ret{{Act: "msg", Data: "", Msg: string(msg)}, {Act: "world", Data: "", Msg: p.Username + " 加入了小世界"}}
	return rets
}

func (this *Auth) loginPlayer(username string, password string) []Ret {
	var db = db.MysqlRun()
	var player Player
	result := db.First(&player, "username = ?", username)
	if result.RowsAffected > 0 {
		var pwd = makePassword(username, password)
		if player.Password == pwd {
			player.Token = uuid.New().String()
			db.Save(&player)

			msg, _ := json.Marshal(map[string]string{"action": "login", "token": player.Token, "msg": "登陸成功"})
			var rets = []Ret{{Act: "msg", Data: "", Msg: string(msg)}, {Act: "world", Data: "", Msg: player.Username + " 回到了小世界"}}
			return rets
		}
	}
	var rets = []Ret{{Act: "msg", Data: "", Msg: "用戶名或密碼錯誤"}}
	return rets
}

func (this *Auth) tokenLogin(token string) []Ret {
	var db = db.MysqlRun()
	var player Player
	result := db.First(&player, "token = ?", token)
	if result.RowsAffected > 0 {
		player.Token = uuid.New().String()
		db.Save(&player)
		msg, _ := json.Marshal(map[string]string{"action": "login", "token": player.Token, "msg": "登陸成功"})
		var rets = []Ret{{Act: "msg", Data: "", Msg: string(msg)}, {Act: "world", Data: "", Msg: player.Username + " 回到了小世界"}}
		return rets
	}
	msg, _ := json.Marshal(map[string]string{"action": "loginfail", "token": "", "msg": "登陸失败"})
	var rets = []Ret{{Act: "msg", Data: "", Msg: string(msg)}}
	return rets
}

func makePassword(username string, password string) string {
	h := md5.New()
	io.WriteString(h, password)
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
	io.WriteString(h, "nkgame")
	io.WriteString(h, username)
	io.WriteString(h, pwmd5)
	last := fmt.Sprintf("%x", h.Sum(nil))
	return last
}

func mapAction(act string) string {
	switch act {
	case "create":
	case "make":
	case "注册":
	case "reg":
	case "register":
	case "註冊":
		return "create"
	case "login":
	case "登陸":
	case "connect":
	case "登录":
	case "登陆":
		return "login"
	default:
		return "no"
	}
	return "no"
}
