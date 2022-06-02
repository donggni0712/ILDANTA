package Rest

import (
	"ILDANTA/Domain"
	"ILDANTA/Service"
	"ILDANTA/Utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var port string

var Apikey = "Mi%2B95EDTMwWb2pbwhatNbhwx4tE4XkBsZ1GiAS2HoGI"

func Search(rw http.ResponseWriter, r *http.Request) {
	var requestBody Domain.Search
	Utils.HandleErr(json.NewDecoder(r.Body).Decode(&requestBody))

	//처음 경로 출력
	var response []*Domain.Result
	response = Service.GetFirstRoute(requestBody.Sx, requestBody.Sy, requestBody.Ex, requestBody.Ey, Apikey)
	res := Service.GetFirstPage(response)
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "*")
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(res)
}

func SearchSub(rw http.ResponseWriter, r *http.Request) {
	var requestBody Domain.SearchSubPath
	Utils.HandleErr(json.NewDecoder(r.Body).Decode(&requestBody))

	secondResponse := Service.GetSubPage(requestBody, Apikey)

	rw.Header().Add("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "*")
	json.NewEncoder(rw).Encode(secondResponse)
}
func RawData(rw http.ResponseWriter, r *http.Request) {

	var requestBody Domain.Search
	Utils.HandleErr(json.NewDecoder(r.Body).Decode(&requestBody))

	//처음 경로 출력
	var response []*Domain.Result
	response = Service.GetFirstRoute(requestBody.Sx, requestBody.Sy, requestBody.Ex, requestBody.Ey, Apikey)

	rw.Header().Add("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "*")
	json.NewEncoder(rw).Encode(response)
}
func Start(aPort int) {
	router := mux.NewRouter()
	port = fmt.Sprintf(":%d", aPort)

	router.HandleFunc("/RawData", RawData).Methods("POST")

	router.HandleFunc("/Search", Search).Methods("POST")
	router.HandleFunc("/Search/Choose", SearchSub).Methods("POST")

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
