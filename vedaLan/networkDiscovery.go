package vedalan

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/schollz/peerdiscovery"
)

type DiscoveryJSON struct {
	IPAddress string `json:"ipAddress"`
	UID       string `json:"uid"`
}

func generateRandomUsername() string {
	words := []string{"momentary", "culture", "random", "example", "words", "go"}

	// Generate two random indices
	rand.Seed(time.Now().UnixNano())
	index1 := rand.Intn(len(words))
	index2 := rand.Intn(len(words))

	// Retrieve the random words
	word1 := words[index1]
	word2 := words[index2]

	// Join the words to form the username
	username := word1 	+ word2

	return username
}


func GetPeersJson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Started Network Discovery")
	discoveries, _ := peerdiscovery.Discover(peerdiscovery.Settings{Limit: 2})

	peers := make([]DiscoveryJSON, 0)
	for _, d := range discoveries {
		// Generate a random username
		uid := generateRandomUsername()

		// Create the discovery JSON
		discovery := DiscoveryJSON{
			IPAddress: d.Address,
			UID:       uid,
		}

		peers = append(peers, discovery)
	}

	// Convert the discoveries to JSON
	jsonData, err := json.Marshal(peers)
	if err != nil {
		log.Println("Error marshaling JSON:", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write the JSON data to the response
	_, err = w.Write(jsonData)
	if err != nil {
		log.Println("Error writing JSON response:", err.Error())
	}
}

