package Rest

import (
	"ILDANTA/ApiClient"
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

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "Main",
		},
		{
			URL:         url("/Search?{SX}&{SY}&{EX}&{EY}"),
			Method:      "POST",
			Description: "ShowFirstRoute",
		},
		{
			URL:         url("/Search?{SX}&{SY}&{EX}&{EY}/{choice}&{where}"),
			Method:      "POST",
			Description: "ShowFirstPath",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func Search(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sx := vars["sx"]
	sy := vars["sy"]
	ex := vars["ex"]
	ey := vars["ey"]
	//처음 경로 출력
	result := Service.ShowFirstRoute(sx, sy, ex, ey, Apikey)
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "*")
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(result)
}

func Choose_TakeOn(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var tempresult []*Domain.Result

	Utils.HandleErr(json.NewDecoder(r.Body).Decode(&tempresult))

	//처음에 뭐 탈 지
	where := vars["whereOn"]
	what := vars["whatOn"]
	result := Service.ClickRoute(where, what, tempresult)

	rw.Header().Add("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "*")
	json.NewEncoder(rw).Encode(result)
}

func Choose_TakeOffandTakeOn(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var tempresult Domain.FirstPath
	Utils.HandleErr(json.NewDecoder(r.Body).Decode(&tempresult))
	//어디서 내려서 뭐 탈 지
	whereOff := vars["whereOff"]
	whereOn := vars["whereOn2"]
	whatOn := vars["whatOn2"]

	result := Service.ClickSubPath(whereOff, whereOn, whatOn, tempresult.AfterPathThemes)

	rw.Header().Add("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "*")
	json.NewEncoder(rw).Encode(result)
}

func SeeRawData(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sx := vars["sx"]
	sy := vars["sy"]
	ex := vars["ex"]
	ey := vars["ey"]
	//처음 경로 출력
	result := ApiClient.CallAPI(sx, sy, ex, ey, Apikey)
	rw.Header().Add("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "*")
	json.NewEncoder(rw).Encode(result)
}

func Start(aPort int) {
	router := mux.NewRouter()
	port = fmt.Sprintf(":%d", aPort)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/development/RawData/{sx}&{sy}&{ex}&{ey}", SeeRawData).Methods("GET")
	router.HandleFunc("/Search/{sx}&{sy}&{ex}&{ey}", Search).Methods("GET")
	router.HandleFunc("/Search/ChooseTakeOn/{whereOn}&{whatOn}", Choose_TakeOn).Methods("GET")
	router.HandleFunc("/Search/ChooseTakeOffOn/{whereOff}&{whereOn2}&{whatOn2}", Choose_TakeOffandTakeOn).Methods("GET")

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
