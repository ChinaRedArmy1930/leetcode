package main

import (
	"fmt"
	"math/rand"
)

//TODO
type RandomizedCollection struct {
	Nums  map[int]int
	count int
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{Nums: map[int]int{}}
}

func (this *RandomizedCollection) Insert(val int) bool {
	this.count++
	if _, ok := this.Nums[val]; ok {
		this.Nums[val] += 1
		return false
	}
	this.Nums[val] += 1
	return true
}

func (this *RandomizedCollection) Remove(val int) bool {
	if _, ok := this.Nums[val]; ok {
		this.count--
		this.Nums[val] -= 1
		if this.Nums[val] <= 0 {
			delete(this.Nums, val)
		}

		return true
	}

	return false
}

func (this *RandomizedCollection) GetRandom() int {
	return this.Nums[rand.Intn(this.count+1)]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	obj := Constructor()
	obj.Insert(1)
	obj.Insert(1)
	obj.Remove(1)
	param_3 := obj.GetRandom()
	fmt.Println(param_3)
}
