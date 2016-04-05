package ipalloc

import (
	"device"
	"fmt"
	. "github.com/gorilla/mux"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	vars := Vars(r)
	ip := vars["ip"]

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	device, err := Registry.Find(ip)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, err.Error())
		return
	}

	data, err := device.Json()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, data)
}

func AddNewDevice(w http.ResponseWriter, r *http.Request) {
	device, err := device.FromJson(r.Body)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if nil != err {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
		return
	}

	err = Registry.Register(device)

	if nil != err {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
		return
	}

	data, err := device.Json()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, data)
}
