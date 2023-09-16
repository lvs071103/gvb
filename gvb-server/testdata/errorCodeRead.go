package main

import (
	"encoding/json"
	"fmt"
	"gvb-server/global"
	"gvb-server/models/resp"
	"os"
)


const file = "/Users/zhoufurong/gocode/src/vblog/gvb-server/models/resp/err_code.json"

// type ErrMap map[string]string
// type ErrMap map[int]string
type ErrMap map[resp.ErrorCode]string

func main()  {
	bytedata, err := os.ReadFile(file)
	if err != nil {
		global.Log.Errorln(err)
		return
	}
	var errMap = ErrMap{}
	err = json.Unmarshal(bytedata, &errMap)
	if err != nil{
		global.Log.Errorln(err)
	}
	fmt.Println(errMap)
	fmt.Println(errMap[resp.SettingsError])
}