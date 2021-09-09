package db

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type db struct {
	name string
	rh   *os.File
	wh   *os.File
}

// func createDbIfnotExist(path string) (*os.File, error) {
// 		os.Lstat(path)
// }

func New(name string) db {
	wh, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	rh, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return db{name: name, rh: rh, wh: wh}
}

func (d db) Set(k string, val interface{}) bool {
	_, err := d.wh.WriteString(fmt.Sprintf("%s,%v\n", k, val))
	if err != nil {
		println(err.Error())
		return false
	}
	return true
}

func (d db) Get(k string) (string, bool) {
	d.wh.Sync()
	d.rh.Seek(0, 0)
	r := bufio.NewReader(d.rh)
	for {
		line, isPrefix, err := r.ReadLine()
		if line == nil || err != nil {
			break
		}
		words := strings.Split(string(line), ",")
		if strings.TrimSpace(words[0]) == k {
			value := string(words[1])
			if isPrefix {
				println("line is big, not handled in this version, update :)")
				// value
				// TODO: handle in future
			}
			return value, true
		}
	}
	return ``, false
}

func (d db) Close() {
	d.rh.Sync()
	d.rh.Close()
}
