package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)


type ObjectStore struct {
	storage map[string]Object
}

func (store *ObjectStore) Add(obj Object) {
	store.storage[obj.Key] = obj
}

func (store *ObjectStore) Get(key string) Object {
	return store.storage[key]
}

type Object struct {
	Key   string
	Value int
	Message string
}

type Worker struct {
	id int
	store *ObjectStore
}

func (store *ObjectStore) generateObjects(size int) []string {
	keys := make([]string, size)
	for i := 0; i < size; i++ {
		i := i
		key := strconv.Itoa(i)
		keys[i] = key
		store.Add(Object{
			Key:   key,
			Value: rand.Int(),
		})
	}
	return keys
}

func (worker Worker) Do(keys <-chan string, objects chan<- Object, wg *sync.WaitGroup) {
	defer wg.Done()
	for key := range keys {
		obj := worker.store.Get(key)
		time.Sleep(50 * time.Millisecond) //sleep to simulate I/O bound process
		obj.Message = fmt.Sprintf("worker %d got object: %2s, %d", worker.id, obj.Key, obj.Value)
		objects <- obj
	}
}

func main() {
	var wg sync.WaitGroup
	store := ObjectStore{
		storage: map[string]Object{},
	}
	keys := store.generateObjects(50)

	// Create buffered channels by specifying the channel size
	keysChannel := make(chan string, len(keys))
	objectsChannel := make(chan Object, len(keys))

	// Prepare all of our workers and provide the input and output channels
	for i := 0; i < 4; i++ {
		i := i
		worker := Worker{
			id:    i,
			store: &store,
		}
		wg.Add(1)
		go worker.Do(keysChannel, objectsChannel, &wg)
	}

	// Start writing to our input channel, once the channel is closed our worker for loops will be able to exit
	// and the WaitGroup for that worker will be closed.
	for i := 0; i < len(keys); i++ {
		i := i
		keysChannel <- keys[i]
	}
	close(keysChannel)

	// Wait for all of the WaitGroups to finish. Once completed we know that we're done writing to the objectsChannel
	// and can close it. The objectsChannel must be closed BEFORE range looping over it so we don't hit a deadlock.
	wg.Wait()
	close(objectsChannel)

	// Read all objects from the channel and process
	for object := range objectsChannel {
		fmt.Println(object.Message)
	}
}
