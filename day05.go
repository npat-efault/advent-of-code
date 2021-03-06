package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

func BadSeq(pc, c rune) bool {
	if (pc == 'a' && c == 'b') ||
		(pc == 'c' && c == 'd') ||
		(pc == 'p' && c == 'q') ||
		(pc == 'x' && c == 'y') {
		return true
	} else {
		return false
	}
}

func IsNice(s string) bool {
	vowels, doubles, badseq := 0, 0, 0
	var pc rune = -1
	for _, c := range s {
		if strings.IndexRune("aeiou", c) != -1 {
			vowels++
		}
		if c == pc {
			doubles++
		}
		if BadSeq(pc, c) {
			badseq++
		}
		pc = c
	}
	if vowels >= 3 && doubles > 0 && badseq == 0 {
		return true
	} else {
		return false
	}
}

func IsNice1(s string) bool {
	var ppc, pc rune = -1, -1
	seqs := make(map[string]int)
	seqFound, repFound := false, false
	for i, c := range s {
		if !seqFound && pc >= 0 {
			sq := string([]rune{pc, c})
			idx, ok := seqs[sq]
			if ok {
				if i-idx > 1 {
					seqFound = true
				}
			} else {
				seqs[sq] = i
			}
		}
		if ppc == c {
			repFound = true
		}
		if seqFound && repFound {
			return true
		}
		ppc = pc
		pc = c
	}
	return false
}

// IsNice2 is equivalent to IsNice1, but it's cleaner and a bit more
// general.
func IsNice2(s string) bool {
	const seqLen = 2
	const repSkip = 1
	seqs := make(map[string]int)
	seqFound, repFound := false, false
	rv := []rune(s)
	for i := 0; i < len(rv); i++ {
		if !seqFound && i <= len(rv)-seqLen {
			sq := string(rv[i : i+seqLen])
			idx, ok := seqs[sq]
			if ok {
				if i-idx >= seqLen {
					seqFound = true
				}
			} else {
				seqs[sq] = i
			}
		}
		if !repFound && i > repSkip {
			if rv[i] == rv[i-repSkip-1] {
				repFound = true
			}
		}
		if seqFound && repFound {
			return true
		}
	}
	return false
}

func Day05(input io.Reader) {
	sc := bufio.NewScanner(input)
	nice, nice1 := 0, 0
	for sc.Scan() {
		str := sc.Text()
		if IsNice(str) {
			nice++
		}
		if IsNice2(str) {
			nice1++
		}
	}
	if sc.Err() != nil {
		log.Fatal(sc.Err().Error())
	}
	fmt.Println(nice, nice1)
}
