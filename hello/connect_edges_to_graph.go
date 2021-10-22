package main

import "fmt"

// RemoveDumplicate 过滤掉重复的元素
func RemoveDumplicate(slc []int) []int {
    result := []int{}
    tempMap := map[int]byte{}
    for _, e := range slc{
        l := len(tempMap)
        tempMap[e] = 0
        if len(tempMap) != l{ 
            result = append(result, e)
        }
    }
    return result
}

// DeleteSliceItem 删除指定的元素
func DeleteSliceItem(a []int, item int) []int{
	for i := 0; i < len(a); i++ {
		if a[i] == item {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}
	return a
}

func GetFather(father *map[int]int, u int) int {
    idx := make([]int, 0)
    v := u
    for v != (*father)[v] {
        idx = append(idx, v)
        v = (*father)[v]
    }
    for _, i := range idx {
        (*father)[i] = v
    }
    return v
}

func EdgesConnectToGraph(edges [][]int) (map[int] int, error) {
    father := make(map[int]int)
    
    labelSet := make([]int, 0)
    for _, pair := range edges {
        src := pair[0]
        dst := pair[1]
        labelSet = append(labelSet, src)
        labelSet = append(labelSet, dst)
    }
    labelSet = RemoveDumplicate(labelSet)
    labelSet = DeleteSliceItem(labelSet, -1)
    for _, label := range labelSet {
        father[label] = label 
    }
    for i := 0; i < len(edges); i++ {
        u, v := edges[i][0], edges[i][1]
        fatherU := GetFather(&father, u)
        fatherV := GetFather(&father, v)
        father[fatherU] = fatherV
    }
    for _, label := range labelSet {
        GetFather(&father, label)
    }

    toDelK := make([]int, 0)
    for k, _ := range father {
        if father[k] == k {
            toDelK = append(toDelK, k)
        }
    }
    for _, k := range toDelK {
        delete(father, k)
    }

    return father, nil
}


func main() {
    edges := [][]int {  {0, 5},
                        {0, 7},
                        {2, 15},
                        {3, 18},
                        {4, 12},
                        {4, 14},
                        {5,  8},
                        {9, 17},
                        {12, 14},
                        {12, 19},
                        {13, 15},
                        {18, 19}}
    father, _ := EdgesConnectToGraph(edges)
    fmt.Println(father)
}
