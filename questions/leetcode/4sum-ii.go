/*
Given four integer arrays nums1, nums2, nums3, and nums4 all of length n, return the number of tuples (i, j, k, l) such that:

0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
 

Example 1:

Input: nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
Output: 2
Explanation:
The two tuples are:
1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0
Example 2:

Input: nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
Output: 1
 

Constraints:

n == nums1.length
n == nums2.length
n == nums3.length
n == nums4.length
1 <= n <= 200
-228 <= nums1[i], nums2[i], nums3[i], nums4[i] <= 228
*/
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    m1 := map[int]int{}
    for _, i := range nums1 {
        for _,j := range nums2 {
            _, ok := m1[i+j]
            if !ok {
                m1[i+j] = 0
            }
            m1[i+j] = m1[i+j] + 1
        }
    }
    
    m2 := map[int]int{}
    for _, k := range nums3 {
        for _,l := range nums4 {
            _, ok := m2[k+l]
            if !ok {
                m2[k+l] = 0
            }
            m2[k+l] = m2[k+l] + 1
        }
    }
    
    output := 0
    
    for index, _ := range m1 {
        if _, ok := m2[-index];ok {
            output = output + m1[index]*m2[-index]
        }
    }
    
    return output
}