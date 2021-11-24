package main
 
import (
    "bytes"
    "io/ioutil"
    "net/http"
)

func main() {
    // 간단한 http.Post 예제
    reqBody := bytes.NewBufferString("Post plain text")
	//var URL string = "https://api.odsay.com/v1/api/searchBusLane?lang=0&busNo=10&apiKey="
	//var apikey string = "Mi%2B95EDTMwWb2pbwhatNbhwx4tE4XkBsZ1GiAS2HoGI"
    resp, err := http.Post("https://api.odsay.com/v1/api/searchBusLane?lang=0&busNo=10&apiKey=Mi%2B95EDTMwWb2pbwhatNbhwx4tE4XkBsZ1GiAS2HoGI", "text/plain", reqBody)
    if err != nil {
        panic(err)
    }
 
    defer resp.Body.Close()
 
    // Response 체크.
    respBody, err := ioutil.ReadAll(resp.Body)
    if err == nil {
        str := string(respBody)
        println(str)
    }
}