package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func CheckPrefix(s *[md5.Size]byte) bool {
	if s[0] == 0 && s[1] == 0 && s[2]&0xf0 == 0 {
		return true
	} else {
		return false
	}
}

func CheckPrefix1(s *[md5.Size]byte) bool {
	if s[0] == 0 && s[1] == 0 && s[2] == 0 {
		return true
	} else {
		return false
	}
}

func Day04(input io.Reader) {
	key, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal(err.Error())
	}
	key = bytes.Trim(key, "\n\r\t ")
	var coin, coin1 int
	coinFound := false
	for i := 0; ; i++ {
		data := append(key, fmt.Sprintf("%d", i)...)
		sum := md5.Sum(data)
		if !coinFound {
			if CheckPrefix(&sum) {
				coin = i
				coinFound = true
			} else {
				// If not Prefix then also not Prefix1
				continue
			}
		}
		if CheckPrefix1(&sum) {
			coin1 = i
			break
		}
	}
	fmt.Println(coin, coin1)
}
