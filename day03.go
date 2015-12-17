package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

type Addr struct{ Lon, Lat int }

func (a *Addr) Move(r rune) bool {
	switch r {
	case '<':
		a.Lon--
	case '>':
		a.Lon++
	case '^':
		a.Lat++
	case 'v':
		a.Lat--
	default:
		return false
	}
	return true
}

func Day03(input io.Reader) {
	bin := bufio.NewReader(input)
	saddr, raddr := Addr{}, Addr{}
	visited := make(map[Addr]int)
	visited[saddr]++
	visited[raddr]++
	for {
		c, _, err := bin.ReadRune()
		if err != nil {
			if err != io.EOF {
				log.Fatal(err.Error())
			}
			break
		}
		saddr.Move(c)
		visited[saddr]++

		c, _, err = bin.ReadRune()
		if err != nil {
			if err != io.EOF {
				log.Fatal(err.Error())
			}
			break
		}
		raddr.Move(c)
		visited[raddr]++
	}
	fmt.Println(len(visited))
}
