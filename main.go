package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	epoch        int64 = 1577836800000 // Custom epoch (e.g., Jan 1, 2020 in milliseconds)
	nodeBits     uint  = 10            // Number of bits for node ID
	sequenceBits uint  = 12            // Number of bits for sequence
	maxNodeID          = -1 ^ (-1 << nodeBits)
	maxSequence        = -1 ^ (-1 << sequenceBits)
	timeShift          = nodeBits + sequenceBits
	nodeShift          = sequenceBits
)

type Snowflake struct {
	mu        sync.Mutex
	timestamp int64
	nodeID    int64
	sequence  int64
}

// NewSnowflake initializes a new Snowflake generator
func NewSnowflake(nodeID int64) (*Snowflake, error) {
	if nodeID < 0 || nodeID > maxNodeID {
		return nil, fmt.Errorf("node ID must be between 0 and %d", maxNodeID)
	}

	return &Snowflake{
		timestamp: 0,
		nodeID:    nodeID,
		sequence:  0,
	}, nil
}

// Generate creates a new unique ID
func (sf *Snowflake) Generate() int64 {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	now := time.Now().UnixMilli()

	if now == sf.timestamp {
		sf.sequence = (sf.sequence + 1) & maxSequence
		if sf.sequence == 0 {
			// Wait for the next millisecond
			for now <= sf.timestamp {
				fmt.Print("a")
				now = time.Now().UnixMilli()
			}
		}
	} else {
		sf.sequence = 0
	}

	sf.timestamp = now

	id := (((now - epoch) << timeShift) |
		(sf.nodeID << nodeShift) |
		sf.sequence)
	return id
}

func main() {
	nodeID := int64(1)

	sf, err := NewSnowflake(nodeID)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// p := fmt.Println
	// Generate a few IDs
	//
	// fmt.Println(time.Now().UnixMilli())
	// fmt.Println(math.Pow(2, float64(nodeBits)) - 1)
	// fmt.Println(-1 ^ (-1 << nodeBits))

	// n := time.Now().UnixMilli()
	// count := 0
	// for {
	// 	now := time.Now().UnixMilli()

	// 	if now != n {
	// 		fmt.Print(count, "diff \n")
	// 		n = now
	// 		count = 0
	// 	} else {
	// 		count++
	// 		fmt.Print("same \n")
	// 	}
	// }

	for i := 0; i < 1; i++ {
		uuid := sf.Generate()
		fmt.Println(uuid)
		fmt.Printf("%b", uuid|0)
		// fmt.Printf(" %T: ", uuid)
		// fmt.Printf("%d \n", unsafe.Sizeof(uuid))
	}

	var y int64

	y = 10

	fmt.Println(y)
}
