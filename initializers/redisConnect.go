package initializers

import (
	"fmt"
	"github.com/go-redis/redis"
)

func ConnectToRedis() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := client.Set("token", "FBB2ub23yo@5B%6~098*643NdsaGgs", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("token").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("::: Redis :::", val)

	err = client.Incr("counter").Err()
	if err != nil {
		panic(err)
	}

}
