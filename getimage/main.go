package main

import(
	//"crypto/tls"
	"net/http"
	"fmt"
	"os"
)

func main() {

	tr := &http.Transport{
		//TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://captchas.oss-cn-hangzhou.aliyuncs.com/simulator-captcha/login/949946.png")

	if err != nil {
		panic(err)
	}


	body := make([]byte,4096,4096)
	len,_ := resp.Body.Read(body)
	fmt.Println(len)
	f, _ := os.Create("validcode.png")

	f.Write(body)
	f.Close()
}
