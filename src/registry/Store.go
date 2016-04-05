package registry

import (
	"bufio"
	. "bytes"
	. "common"
	"device"
	"io/ioutil"
	"os"
)

var DataStream = make(chan string, 1)

func ReadStore(store string) error {

	hdData, err := os.Open(store)
	if err != nil {
		return NewError(ERR_OPEN_DATAFILE, store, err.Error())
	}

	defer func() {
		if err := hdData.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(hdData)

	for scanner.Scan() {
		line := scanner.Text()
		DataStream <- line
	}
	defer close(DataStream)
	return nil
}

func (this *Registry) Save() error {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()
	return ioutil.WriteFile(this.store, this.TextBytes().Bytes(), 0644)
}

func (this *Registry) Load(store string) {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()

	this.store = store
	if FileExist(this.store) {
		go ReadStore(this.store)

		for record := range DataStream {
			device := device.Deserialize(record)
			this.addDevice(device)
		}
	}
}
