package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func dbHandlerFunc(db *DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := db.Show()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed show db: %v", err), 500)
			return
		}
		w.Write(res)
	}
}

func scheduleHandlerFunc(ctx context.Context, scheduler *Scheduler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		jst, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed get location: %v", err), 500)
			return
		}
		if err := scheduler.Reserve(ctx, time.Now().In(jst).Add(time.Duration(3) * time.Minute)); err != nil {
			http.Error(w, fmt.Sprintf("Failed create schdule job: %v", err), 500)
			return
		}
		fmt.Fprintln(w, "Scheduled")
	}
}
