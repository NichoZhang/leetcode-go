package main

/**
* https://leetcode.com/problems/design-hashmap/description/
* Example:
* MyHashMap hashMap = new MyHashMap();
* hashMap.put(1, 1);
* hashMap.put(2, 2);
* hashMap.get(1);            // returns 1
* hashMap.get(3);            // returns -1 (not found)
* hashMap.put(2, 1);          // update the existing value
* hashMap.get(2);            // returns 1
* hashMap.remove(2);          // remove the mapping for 2
* hashMap.get(2);            // returns -1 (not found)
*
* 33 / 33 test cases passed.
* beats 100%
* Runtime:120ms
* https://leetcode.com/submissions/detail/169409226/
**/

type MyHashMap struct {
	HashMap map[int]int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	hashMap := make(map[int]int)
	return MyHashMap{
		HashMap: hashMap,
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.HashMap[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	_, ok := this.HashMap[key]
	if ok {
		return this.HashMap[key]
	} else {
		return -1
	}
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	_, ok := this.HashMap[key]
	if ok {
		delete(this.HashMap, key)
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
