package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type DateDiff struct {
	InYears   int `json:"InYears"`
	InMonths  int `json:"inMonths"`
	InWeeks   int `json:"inWeeks"`
	InDays    int `json:"inDays"`
	InHours   int `json:"inHours"`
	InMinutes int `json:"inMinutes"`
	InSeconds int `json:"inSeconds"`
}

func DateDiffHandler(w http.ResponseWriter, r *http.Request) {
	date1str := r.FormValue("d1")
	date2str := r.FormValue("d2")

	layout := "2006-01-02"

	date1, err := time.Parse(layout, date1str)

	if err != nil {
		fmt.Println(err)
	}

	date2, err := time.Parse(layout, date2str)

	if err != nil {
		fmt.Println(err)
	}

	year, month, week, day, hour, min, sec := diff(date2, date1)

	res2D := &DateDiff{
		InYears:   year,
		InMonths:  month,
		InWeeks:   week,
		InDays:    day,
		InHours:   hour,
		InMinutes: min,
		InSeconds: sec,
	}
	res2B, _ := json.Marshal(res2D)

	fmt.Fprintf(w, string(res2B))

}

func diff(a, b time.Time) (year, month, week, day, hour, min, sec int) {

	year = a.Year() - b.Year()
	month = (int(a.Month()) + 12*a.Year()) - (int(b.Month()) + 12*b.Year())
	week = int(a.Sub(b).Hours() / (24 * 7))
	day = int(a.Sub(b).Hours() / 24)
	hour = int(a.Sub(b).Hours())
	min = int(a.Sub(b).Minutes())
	sec = int(a.Sub(b).Seconds())

	return
}
