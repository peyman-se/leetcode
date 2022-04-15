/*
You are given an integer array nums that is sorted in non-decreasing order.

Determine if it is possible to split nums into one or more subsequences such that both of the following conditions are true:

Each subsequence is a consecutive increasing sequence (i.e. each integer is exactly one more than the previous integer).
All subsequences have a length of 3 or more.
Return true if you can split nums according to the above conditions, or false otherwise.

A subsequence of an array is a new array that is formed from the original array by deleting some (can be none) of the elements without disturbing the relative positions of the remaining elements. (i.e., [1,3,5] is a subsequence of [1,2,3,4,5] while [1,3,2] is not).

 

Example 1:

Input: nums = [1,2,3,3,4,5]
Output: true
Explanation: nums can be split into the following subsequences:
[1,2,3,3,4,5] --> 1, 2, 3
[1,2,3,3,4,5] --> 3, 4, 5
Example 2:

Input: nums = [1,2,3,3,4,4,5,5]
Output: true
Explanation: nums can be split into the following subsequences:
[1,2,3,3,4,4,5,5] --> 1, 2, 3, 4, 5
[1,2,3,3,4,4,5,5] --> 3, 4, 5
Example 3:

Input: nums = [1,2,3,4,4,5]
Output: false
Explanation: It is impossible to split nums into consecutive increasing subsequences of length 3 or more.
 

Constraints:

1 <= nums.length <= 104
-1000 <= nums[i] <= 1000
nums is sorted in non-decreasing order.
*/

package main

func isPossible(nums []int) bool {
    s := []*SubArray{}
    ls := []*SubArray{}
    
    var lastElement int
    for _, i := range nums {
        if len(s) == 0 && len(ls) == 0 {
            e := &SubArray{
                LastElement: i,
                Size: 1,
            }
            // s = append(s, e)
            ls = append(ls, e)
            lastElement = i
            continue
        }
        
        if i > (lastElement + 1) {
            //create a new sub-array
            e := &SubArray{
                LastElement: i,
                Size: 1,
            }
            s = append(s, ls...)
            ls = []*SubArray{e}
            lastElement = i
            continue
        }
        
        //add it to the sub-array that has the smallest size
        sort.Slice(ls, func(i, j int) bool {
            return ls[i].Size < ls[j].Size
        })
        
        
        found := false
        for index, e := range ls {
            if (e.LastElement + 1) != i {
                continue
            }
            
            e.LastElement = i
            e.Size = e.Size + 1
            ls[index] = e
            found = true
            break
        }
        
        if !found {
            //create a new sub-array
            e := &SubArray{
                LastElement: i,
                Size: 1,
            }
            
            ls = append(ls, e)
            // ls = []*SubArray{e}
        }

        lastElement = i
    }
    
    s = append(s, ls...)
    for _, e := range s {
        if e.Size < 3 {
            return false
        }
    }
    
    return true
}


type SubArray struct {
    LastElement int
    Size int
}