
package main

import (
	"net/http"
	"fmt"
	//"encoding/json"
	//"io/ioutil"
	"mime/multipart"
	"bytes"
	"io"
	"io/ioutil"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("C:/Users/zhou1/Documents/html/")))
	http.Handle("/test",&transister{})
	//err:=http.ListenAndServe(":9090",nil)
	err:=http.ListenAndServeTLS(":9090", "cert.pem","key.pem",nil)
	if err!=nil{
		fmt.Println(err)
	}
}

func postFile(fileHeader *multipart.FileHeader)(io.Reader, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("file", fileHeader.Filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return nil,err
	}


	//打开文件句柄操作
	fh, err := fileHeader.Open()
	if err != nil {
		fmt.Println("error opening file")
		return nil,err
	}
	defer fh.Close()

	//iocopy
	copiedLen, err := io.Copy(fileWriter, fh)
	if err != nil {
		return nil,err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	fmt.Printf("contentType %s, copyLen %d\n",contentType,copiedLen)

	//fieldWriter, err := bodyWriter.CreateFormField("file", fileHeader.Filename)
	client := http.Client{}
	req, _ := http.NewRequest("POST", "http://10.0.69.65:8082/test?apiCode=YH002",
		bodyBuf)
	req.Header.Set("Authorization","eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzZWNyZX" +
		"RJZCI6Inl1bnNoYW5nIn0.C5o2k0foWC2ySPoN26brYdGla-NzvEkcTrDWCq2lX_4")
	req.Header.Set("Content-Type", contentType)

	resp, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	//defer resp.Body.Close()
	//resp_body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return err
	//}
	//fmt.Println(resp.Status)
	//fmt.Println(string(resp_body))
	return resp.Body,nil
}


type transister struct {}

func (t *transister) ServeHTTP(w http.ResponseWriter,r *http.Request){
	//fmt.Println(r.Body)

	//client := http.Client{}
	//client.Post("http://10.0.69.65:8082/test?apiCode=YH002",)



	//body, err := ioutil.ReadAll(r.Body)


	_,fileHeader,err := r.FormFile("file")
	if err!=nil{
		fmt.Println(err)
		return
	}

	//fmt.Println(fileHeader.Size)

	body,err :=postFile(fileHeader)
	if err!=nil{
		fmt.Println(err)
		return
	}


	//err = r.ParseMultipartForm(32 << 20)
	//for k,v := range r.Form {
	//	fmt.Printf("in the form: k:%s,v:%s\n",k,v)
	//}
	//fmt.Println(err)




	//data := &outSide{
	//	Msg: "",
	//	ReturnCode: "200",
	//	Success: "Y",
	//	Data: innerSide{
	//		Msg:"成功",
	//		Data:"你好",
	//		Code: "000",
	//	},
	//}
	//
	//jsonData, _ := json.Marshal(data)

	writeBody, err := ioutil.ReadAll(body)
	////resContent := []byte{}
	////res.Body.Read(resContent)
	////fmt.Println(res.StatusCode)
	fmt.Println(string(writeBody))
	//fmt.Println(err)
	//
	w.Write(writeBody)
}

type outSide struct {
	Msg string     `json:"msg"`
	Success   string     `json:"success"`
	ReturnCode string `json:"returnCode"`
	Data innerSide `json:"data"`
}

type innerSide struct {
	Msg string     `json:"msg"`
	Code   string     `json:"code"`
	Data string `json:"data"`
}

