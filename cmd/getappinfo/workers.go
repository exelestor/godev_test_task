package main

import "github.com/exelestor/godev_test_task/pkg/appinfo"

const WORKERS = 24

type Workers struct {
	AppsToProcess chan string
	ProcessedApps chan appinfo.Result
	done          chan bool
}

func (w *Workers) worker() {
	for a := range w.AppsToProcess {
		res, err := appinfo.Get(a, lang)
		if err != nil {
			w.ProcessedApps <- appinfo.Result{ID: a, Error: err.Error()}
			continue
		}
		w.ProcessedApps <- res
	}
	w.done <- true
}

func (w *Workers) Init() {
	w.AppsToProcess = make(chan string, 100)
	w.ProcessedApps = make(chan appinfo.Result, 100)
	w.done = make(chan bool)
}

func (w *Workers) Run() {
	for n := 0; n < WORKERS; n++ {
		go w.worker()
	}
}

func (w *Workers) WaitUntilDone() {
	for n := 0; n < WORKERS; n++ {
		<-w.done
	}
}
