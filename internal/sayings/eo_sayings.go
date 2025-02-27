package sayings

import (
	"bufio"
	"errors"
	"math/rand"
	"os"
	"sync"
	"time"
)

var (
	mu      sync.RWMutex
	sayings []string
	rng     = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// LoadSayings reads sayings from a file and stores them in memory
func LoadSayings(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var newSayings []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newSayings = append(newSayings, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if len(newSayings) == 0 {
		return errors.New("no sayings found in file")
	}

	// Shuffle sayings
	rng.Shuffle(len(newSayings), func(i, j int) { newSayings[i], newSayings[j] = newSayings[j], newSayings[i] })

	// Update sayings safely
	mu.Lock()
	sayings = newSayings
	mu.Unlock()

	return nil
}

// GetRandomSaying returns a random saying from the list
func GetRandomSaying() (string, error) {
	mu.RLock()
	defer mu.RUnlock()

	if len(sayings) == 0 {
		return "", errors.New("no sayings available")
	}

	return sayings[rng.Intn(len(sayings))], nil
}

// LoadSayingsFromSlice allows loading sayings from a slice (used for testing)
func LoadSayingsFromSlice(s []string) {
	mu.Lock()
	defer mu.Unlock()
	sayings = s
}

// GetAllSayings returns all sayings (used for testing)
func GetAllSayings() []string {
	mu.RLock()
	defer mu.RUnlock()
	return sayings
}
