package main

/**
* https://leetcode.com/problems/design-hashset/description/
* Example:
* MyHashSet hashSet = new MyHashSet();
* hashSet.add(1);
* hashSet.add(2);
* hashSet.contains(1); // returns true
* hashSet.contains(3); // returns false (not found)
* hashSet.add(2);
* hashSet.contains(2); // returns true
* hashSet.remove(2);
* hashSet.contains(2); // returns false (not found)
*
* 28 / 28 test cases passed.
* beats 100%
* Runtime:104ms
* https://leetcode.com/submissions/detail/169405715/
**/

type MyHashSet struct {
	Keys map[int]bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	keys := make(map[int]bool)
	return MyHashSet{
		Keys: keys,
	}
}

func (this *MyHashSet) Add(key int) {
	_, ok := this.Keys[key]
	if !ok {
		this.Keys[key] = true
	}
}

func (this *MyHashSet) Remove(key int) {
	_, ok := this.Keys[key]
	if ok {
		delete(this.Keys, key)
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	_, ok := this.Keys[key]
	if ok {
		return true
	} else {
		return false
	}
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
