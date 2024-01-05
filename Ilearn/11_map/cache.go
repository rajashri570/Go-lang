package main

import (
	"errors"
	"fmt"
)

type Cache struct {
	table map[string]string
}

func (c *Cache) Put(key, value string) error {
	if len(c.table) >= 3 {
		//fmt.Println("full")
		return errors.New("cache is full")
	}
	c.table[key] = value
	return nil
}

func main() {
	var myCache *Cache = new(Cache)
	myCache.table = make(map[string]string)

	myCache.Put("raj", "rajashri")
	myCache.Put("jay", "jayashri")
	myCache.Put("bhagu", "bhagyashri")
	myCache.Put("kishu", "Krishna")

	fmt.Println("Cache contents:")
	for key, value := range myCache.table {
		fmt.Printf("%s: %s\n", key, value)
	}

}
