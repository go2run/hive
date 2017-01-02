package performance

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/uber-go/zap"
)

//PerforData to store performance data
type PerforData struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

var perforDatas []PerforData

// HandlePerfor handle CPU info from bee
//curl -H "Content-Type:application/json" -X POST --data '{"id": 946,"name":"cpu max","value":56}' http://localhost:61616/performance/cpu/946
func HandlePerfor(w http.ResponseWriter, r *http.Request) {
	logger := zap.New(
		zap.NewJSONEncoder(zap.NoTime()), // drop timestamps in tests
	)

	params := mux.Vars(r)

	var data PerforData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		logger.Fatal("Decode json data failed.", zap.Error(err))
	} else {
		data.ID = params["deviceid"]
		json.NewEncoder(w).Encode(data)
		perforDatas = append(perforDatas, data)
	}

	logger.Info("data is", zap.String("id", data.ID), zap.String("name", data.Name), zap.Float64("value", data.Value))

}

// GetPerfor return CPU info to web page
func GetPerfor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["deviceid"]
	for _, item := range perforDatas {
		if strings.EqualFold(item.ID, id) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&PerforData{})
}

//Init use to init performance collection framework
func Init() {
	logger := zap.New(
		zap.NewJSONEncoder(zap.NoTime()), // drop timestamps in tests
	)

	logger.Info("Begin to Init Performance Collection Framework.")
	router := mux.NewRouter()
	router.HandleFunc("/performance/{deviceid}", HandlePerfor).Methods("POST")
	router.HandleFunc("/performance/{deviceid}", GetPerfor).Methods("GET")

	err := http.ListenAndServe(":61616", router)
	if err != nil {
		logger.Fatal("Start to listen&serve port 91919 failed.", zap.Error(err))
	}

	logger.Info("Init Performance Collection Framework successfully.")
}
