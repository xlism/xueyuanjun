package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

type DataBucket struct {
	buf   *bytes.Buffer
	cond  *sync.Cond
	mutex *sync.RWMutex
}

func NewDataBucket() *DataBucket {
	buf := make([]byte, 0)
	db := &DataBucket{
		buf:   bytes.NewBuffer(buf),
		mutex: new(sync.RWMutex),
	}
	db.cond = sync.NewCond(db.mutex.RLocker())
	return db
}

func (db *DataBucket) read(i int) {
	var data []byte
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	for {
		d, err := db.buf.ReadByte()
		if err != nil {
			if err == io.EOF {
				if string(data) != "" {
					fmt.Printf("reader-%d: %s\n", i, data)
				}
			}
			db.cond.Wait()
			continue
		}
		data = append(data, d)
	}
}

func (db *DataBucket) put(d []byte) (int, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	n, err := db.buf.Write(d)
	db.cond.Signal()
	return n, err
}

func Run() {
	db := NewDataBucket()
	go db.read(1)
	go func(i int) {
		d := fmt.Sprintf("data-%d", i)
		db.put([]byte(d))
	}(1)
	time.Sleep(100 * time.Millisecond)
}
