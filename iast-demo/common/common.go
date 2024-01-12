package common

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"io"
	"strings"
)

func LowerMapKey(dataMap map[string]string) map[string]string {
	lowerKeyMap := map[string]string{}
	for k, v := range dataMap {
		lowerKeyMap[strings.ToLower(k)] = v
	}
	return lowerKeyMap
}

func Map2StringLine(dataMap map[string]string, removeKeys ...string) string {
	set := mapset.NewSet()
	for _, v := range removeKeys {
		set.Add(v)
	}
	var result []string
	for k, v := range dataMap {
		if set.Contains(strings.ToLower(k)) {
			continue
		}
		data := fmt.Sprintf("%s: %s", k, v)
		result = append(result, data)
	}
	return strings.Join(result, "\n")
}

func GenRedisKey(data ...string) string {
	if len(data) == 0 {
		return ""
	}
	return strings.Join(data, ":")
}

func MD5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func PrintReader(outReader *bufio.Reader) {
	for {
		line, err := outReader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Print(line)
	}
}
