// +build ignore

package main

import (
	"fmt"
	"github.com/m3hm3t/go-concurrency-patterns-library/pkg/sharding"
)

func main() {
	shardedMap := sharding.NewShardedMap(5)

	shardedMap.Set("alpha", 1)
	shardedMap.Set("beta", 2)
	shardedMap.Set("gamma", 3)

	fmt.Println(shardedMap.Get("alpha"))
	fmt.Println(shardedMap.Get("beta"))
	fmt.Println(shardedMap.Get("gamma"))

	keys := shardedMap.Keys()
	for _, k := range keys {
		fmt.Println(k)
	}

	shardedMap.Delete("beta")

	keys = shardedMap.Keys()
	for _, k := range keys {
		fmt.Println(k)
	}
}