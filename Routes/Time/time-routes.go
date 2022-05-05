package Time

import (
	"TimeAPI/DB"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"time"
)

func RegisterTimeRoutes(router *mux.Router) {
	router.HandleFunc("/time", getTime)
}

func getTime(w http.ResponseWriter, r *http.Request) {
	timezoneQueryParam := r.URL.Query().Get("tz")
	hasPassedTimezone := true
	if timezoneQueryParam == "" {
		timezoneQueryParam = "UTC"
		hasPassedTimezone = false
	}

	timezones := strings.Split(timezoneQueryParam, ",")
	times := make(map[string]string)

	times = convertToTimezones(timezones, times)

	if !hasPassedTimezone {
		newtimes := make(map[string]string)
		newtimes["current_time"] = times["UTC"]
		times = newtimes
	}

	w.Header().Add("Content-Type", "application/json")
	if times != nil {
		_ = json.NewEncoder(w).Encode(times)
	} else {
		w.WriteHeader(404)
	}
}

func convertToTimezones(timezones []string, times map[string]string) map[string]string {
	for _, t := range timezones {
		if !DB.IsTimezoneValid(t) {
			times = nil
			break
		}
		loc, _ := time.LoadLocation(t)
		times[t] = fmt.Sprint(time.Now().In(loc))
	}
	return times
}
