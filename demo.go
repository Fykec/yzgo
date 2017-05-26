package main

import (
	"bytes"
	"fmt"
	"encoding/json"
)

func printResult(ret string) {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(ret), "", "\t")
	if err != nil {
		println(ret)
	} else {
		fmt.Println(out.String())
	}

}

func main() {
	//client := YZClient{isOAuth: true, accessToken: ""}
	client := YZClient{appID: "<AppID>", appSecret: "<AppSecret>"}
	ret := client.Invoke("kdt.shop.basic.get", "1.0.0", "GET", map[string]string{}, map[string]string{})
	printResult(ret)
}
