package main

import "fmt"

const charset string = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type Partition struct {
	Keys []string
}

// partition takes an input set of keys and breaks it into numPartition different partitions of relatively equal size.
// If more partitions exist than
func partition(keys []string, numPartitions uint8) []Partition {
	keysPerThread := len(keys) / int(numPartitions)
	// If there is a remainder, add one to the keys per thread. The last partition will have fewer items.
	if len(keys)%int(numPartitions) != 0 {
		keysPerThread = keysPerThread + 1
	}

	keysCopy := make([]string, len(keys))
	copy(keysCopy, keys)
	partitions := make([]Partition, numPartitions)
	for i := 0; i < int(numPartitions); i++ {
		if keysPerThread > len(keysCopy) {
			partitions[i] = Partition{
				Keys: keysCopy,
			}
		} else {
			partitions[i] = Partition{
				Keys: keysCopy[:keysPerThread],
			}
			keysCopy = keysCopy[keysPerThread:]
		}
	}

	return partitions
}

func generateKeys(size int) []string {
	keys := make([]string, size)
	for i := 0; i < size; i++ {
		keys[i] = string(charset[i])
	}
	return keys
}

func main() {
	numPartitions := uint8(20)
	keys := generateKeys(42)
	partitions := partition(keys, numPartitions)
	fmt.Printf("NumPartitions: %d\n", numPartitions)
	fmt.Println("Partitions:")
	for _, p := range partitions {
		fmt.Printf("%v\n", p.Keys)
	}

}
