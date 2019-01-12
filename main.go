package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type tmplData struct {
	Ss []strs
	S  strs
}

type session struct {
	Start  time.Time
	End    time.Time
	D      time.Duration
	Active bool
}

type strs struct {
	Start  string
	End    string
	D      string
	Hours  string
	Active bool
}

var ss = make([]strs, 0)

var s = &session{
	Start:  time.Now(),
	End:    time.Now(),
	D:      time.Second,
	Active: true,
}

func main() {

	readSessions()

	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/stop", http.HandlerFunc(stop))
	http.Handle("/start", http.HandlerFunc(start))
	http.Handle("/download", http.HandlerFunc(serveFile))

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if s.Active {
		s.D = time.Now().Sub(s.Start)
	}
	d := tmplData{ss, sessionString(s)}
	mainTmpl.Execute(w, d)
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("db.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	w.Header().Set("Content-Disposition", "attachment; filename=tracker.csv")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
	io.Copy(w, f)
	//http.Redirect(w, r, "/", http.StatusSeeOther)
}

func stop(w http.ResponseWriter, r *http.Request) {
	s.Active = false
	s.End = time.Now()
	s.D = s.End.Sub(s.Start)
	str := sessionString(s)
	ss = append(ss, str)
	writeResult(str)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func start(w http.ResponseWriter, r *http.Request) {
	s.Active = true
	s.Start = time.Now()
	s.End = time.Now()
	s.D = time.Second
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func writeResult(s strs) {
	f, err := os.OpenFile("db.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := csv.NewWriter(f)

	data := []string{
		s.Start,
		s.End,
		s.D,
		s.Hours,
	}
	w.Write(data)
	w.Flush()
}

func readSessions() {
	f, err := os.OpenFile("db.csv", os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, rec := range records {
		temp := strs{
			Start: rec[0],
			End:   rec[1],
			D:     rec[2],
			Hours: rec[3],
		}
		ss = append(ss, temp)
	}
}

const tlayout = "2006-01-02 15:04:05"

func tf(t time.Time) string {
	return t.Format(tlayout)
}

func tp(s string) time.Time {
	t, err := time.Parse(tlayout, s)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func sessionString(s *session) strs {
	h := float64((s.End.Unix() - s.Start.Unix())) / 60.0 / 60.0
	res := strs{
		Start:  tf(s.Start),
		End:    tf(s.End),
		D:      s.D.String(),
		Hours:  fmt.Sprintf("%0.2f", h),
		Active: s.Active,
	}
	return res
}
