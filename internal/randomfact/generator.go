package randomfact

import (
	"encoding/json"
	"math/rand"
	"time"
)

// GetRandomFact return a random fact from the list provided
func GetRandomFact(factList []byte) *string {
	var facts = new(FactAPIResponse)
	err := json.Unmarshal(factList, &facts)
	if err != nil {
		panic("Could not decode")
	}

	newFactList := facts.Facts

	rand.Seed(time.Now().UnixNano())

	chosen := newFactList[rand.Intn(len(newFactList)-1)]

	return chosen.Text
}
