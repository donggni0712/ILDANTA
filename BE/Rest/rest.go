package Rest

import (
	"ILDANTA/Domain"
	"ILDANTA/Repository"
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
	result := Repository.ShowFirstRoute(sx, sy, ex, ey, Apikey)
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(result)
}

func Choose_TakeOn(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sx := vars["sx"]
	sy := vars["sy"]
	ex := vars["ex"]
	ey := vars["ey"]
	var tempresult []*Domain.Result
	tempresult = Repository.ShowFirstRoute(sx, sy, ex, ey, Apikey)

	//처음에 뭐 탈 지
	where := vars["whereOn"]
	what := vars["whatOn"]
	result := Repository.ClickRoute(where, what, tempresult)

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(result)
}

func Choose_TakeOffandTakeOn(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sx := vars["sx"]
	sy := vars["sy"]
	ex := vars["ex"]
	ey := vars["ey"]
	var tempresult []*Domain.Result
	tempresult = Repository.ShowFirstRoute(sx, sy, ex, ey, Apikey)

	//처음에 뭐 탈 지
	where := vars["whereOn"]
	what := vars["whatOn"]
	var result Domain.FirstPath
	result = Repository.ClickRoute(where, what, tempresult)
	//어디서 내려서 뭐 탈 지
	whereOff := vars["whereOff"]
	whereOn := vars["whereOn2"]
	whatOn := vars["whatOn2"]

	Repository.ClickSubPath(whereOff, whereOn, whatOn, result.AfterPathThemes)

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(vars)
}

func Start(aPort int) {
	router := mux.NewRouter()
	port = fmt.Sprintf(":%d", aPort)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/Search/{sx}&{sy}&{ex}&{ey}", Search).Methods("GET")
	router.HandleFunc("/Search/{sx}&{sy}&{ex}&{ey}/ChooseTakeOn/{whereOn}&{whatOn}", Choose_TakeOn).Methods("GET")
	router.HandleFunc("/Search/{sx}&{sy}&{ex}&{ey}/ChooseTakeOn/{whereOn}&{whatOn}/ChooseTakeOffOn/{whereOff}&{whereOn2}&{whatOn2}", Choose_TakeOffandTakeOn).Methods("GET")

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
