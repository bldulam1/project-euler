package problem22

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// FetchNames fetches the names from project euler
func FetchNames(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	names := strings.Split(strings.Replace(
		string(body), "\"", "", -1),
		",",
	)

	return names
}

func characterScore(char rune) uint8 {
	return uint8(char - 'A' + 1)
}

// NameScore calculates the score of a name
func NameScore(name string) int {
	score := 0
	for _, char := range name {
		score += int(characterScore(char))
	}
	return score
}
