package codenames

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type WordCard struct {
	Front string
	Back  string
}

func (wc *WordCard) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&wc.Front, &wc.Back}
	expLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if l, e := len(tmp), expLen; l != e {
		return fmt.Errorf("wrong number of words for card. Expected %d, got %d", expLen, l)
	}
	return nil
}

func (wc *WordCard) String() string {
	if rand.Float32() < 0.5 {
		return wc.Front
	} else {
		return wc.Back
	}
}

func GetWordCardsFromFile(file string) *[]WordCard {
	// Make sure the file exists
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Println("Unable to load wordlist file.")
		os.Exit(1)
	}

	// Read the file contents
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Unable to read word list file!")
		os.Exit(1)
	}

	cards := &[]WordCard{}
	err = json.Unmarshal(data, &cards)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Unable to decode word list file. Ensure you have valid JSON and it matches the format defined in the README.")
	}

	return cards
}

func ShuffleCards(cards *[]WordCard) *[]WordCard {
	rand.Shuffle(len(*cards), func(i, j int) {
		(*cards)[i], (*cards)[j] = (*cards)[j], (*cards)[i]
	})

	return cards
}

func PickCards(cards *[]WordCard, n int) *[]WordCard {
	pickedCards := append([]WordCard{}, (*cards)[0:n]...)
	return &pickedCards
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
