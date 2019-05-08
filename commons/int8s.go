
package commons

import (
	"sort"
	"math/rand"
)

const EmptyInt8 int8 =0

type Int8Collection struct{
	value	[]int8
}

func NewInt8Collection(value []int8) *Int8Collection {
	return &Int8Collection{value:value}
}

func(c *Int8Collection) Concate(given []int8)  *Int8Collection {
	value := make([]int8, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Int8Collection) Drop(n int)  *Int8Collection {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Int8Collection) Filter(fn func(int, int8)bool)  *Int8Collection {
	value := make([]int8, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Int8Collection) First() int8 {
	if len(c.value) <= 0 {
		return EmptyInt8
	} 
	return c.value[0]
}

func(c *Int8Collection) Last() int8 {
	if len(c.value) <= 0 {
		return EmptyInt8
	} 
	return c.value[len(c.value)-1]
}

func(c *Int8Collection) Map(fn func(int, int8)) *Int8Collection {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Int8Collection) Reduce(fn func(int8, int8, int) int8,initial int8) int8   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Int8Collection) Reverse()  *Int8Collection {
	value := make([]int8, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Int8Collection) Unique()  *Int8Collection{
	value := make([]int8, 0, len(c.value))
	seen:=make(map[int8]struct{})
	for _, each := range c.value {
		if _,exist:=seen[each];exist{
			continue
		}		
		seen[each]=struct{}{}
		value=append(value,each)			
	}
	c.value = value
	return c
}

func(c *Int8Collection) Append(given int8) *Int8Collection {
	c.value=append(c.value,given)
	return c
}

func(c *Int8Collection) Len() int {
	return len(c.value)
}

func(c *Int8Collection) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Int8Collection) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Int8Collection)  Sort()  *Int8Collection {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] <= (c.value[j])
	})
	return c 
}

func(c *Int8Collection) All(fn func(int, int8)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Int8Collection) Any(fn func(int, int8)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Int8Collection) Paginate(size int)  [][]int8 {
	var pages  [][]int8
	prev := -1
	for i := range c.value {
		if (i-prev) <= size-1 && i != (len(c.value)-1) {
			continue
		}
		pages=append(pages,c.value[prev+1:i+1])
		prev=i
	}
	return pages
}

func(c *Int8Collection) Pop() int8{
	if len(c.value) <= 0 {
		return EmptyInt8 
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Int8Collection) Prepend(given int8) *Int8Collection {
	c.value = append([]int8{given},c.value...)
	return c
}

func(c *Int8Collection) Max() int8{
	if len(c.value) <= 0 {
		return EmptyInt8 
	}
	var max int8
	for idx,each := range c.value {
		if idx==0{
			max=each
			continue
		}
		if max <= each {
			max = each
		}
	}
	return max
}


func(c *Int8Collection) Min() int8{
	if len(c.value) <= 0 {
		return EmptyInt8 
	}
	var min int8
	for idx,each := range c.value {
		if idx==0{
			min=each
			continue
		}
		if each  <= min {
			min = each
		}
	}
	return min
}

func(c *Int8Collection) Random() int8{
	if len(c.value) <= 0 {
		return EmptyInt8 
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Int8Collection) Shuffle() *Int8Collection {
	if len(c.value) <= 0 {
		return nil
	}
	indexes := make([]int, len(c.value))
	for i := range c.value {
		indexes[i] = i
	}
	
	rand.Shuffle(len(c.value), func(i, j int) {
		c.value[i], c.value[j] = 	c.value[j], c.value[i] 
	})
	
	return c
}

func(c *Int8Collection) Collect() []int8{
	return c.value
}
