package byte_ascii_op_benchmark

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenAesKeyWithMap(initKey string) (aesKey, key string) {
	keyMap := map[string]string{
		"0": "A",
		"1": "B",
		"2": "C",
		"3": "D",
		"4": "E",
		"5": "F",
		"6": "G",
		"7": "H",
		"8": "I",
		"9": "J",
	}

	now := time.Now().Format("20060102150405")
	for i := 0; i < len(now); i++ {
		if rand.Intn(2) == 0 {
			key += strings.ToUpper(keyMap[now[i:i+1]])
		} else {
			key += strings.ToLower(keyMap[now[i:i+1]])
		}
	}
	aesKey = fmt.Sprintf("%32x", md5.Sum([]byte(initKey+now)))
	return
}

func GenAesKeyWithSlice(initKey string) (aesKey, key string) {
	keys := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J'}
	now := []byte(time.Now().Format("20060102150405"))
	ln := len(now)
	keybytes := make([]byte, ln)
	for i := 0; i < ln; i++ {
		if rand.Intn(2) == 0 {
			keybytes[i] = keys[now[i]-'0']
		} else {
			keybytes[i] = keys[now[i]-'0'] + 32
		}
	}
	lk := len(initKey)
	key = string(keybytes)
	tmp := make([]byte, lk+ln)
	copy(tmp, []byte(initKey))
	copy(tmp[lk:], now)
	aesKey = fmt.Sprintf("%32x", md5.Sum(tmp))
	return
}
