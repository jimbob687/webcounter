package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"os"
	"strconv"
)

var ctx = context.Background()

func main() {

	http.HandleFunc("/", webCounter)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func webCounter(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		redisUrl := os.Getenv("REDIS_URL")
		redisAddr := fmt.Sprintf("%s:6379", redisUrl)
		rdb := redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		val, err := rdb.Get(ctx, "key1").Result()
		if err != nil {
			fmt.Println("key1 not found")
			storeValue(rdb, "1")
                        val = "1"
		} else {
			intVar, err := strconv.Atoi(val)
			fmt.Println("Got value: ", val)
			if err != nil {
				fmt.Println("error: ", err)
			} else {
				intVar++
				storeValue(rdb, strconv.Itoa(intVar))
			}
		}

		fmt.Fprintf(w, "Holla! we have hit %s times", val)
	}
}

func storeValue(rdb *redis.Client, value string) {
	fmt.Println("About to store value")
	err := rdb.Set(ctx, "key1", value, 0).Err()
	if err != nil {
		fmt.Println("Unable to store value")
	}
}
