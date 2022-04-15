//https://leetcode.com/problems/longest-substring-without-repeating-characters/

/*
3. Longest Substring Without Repeating Characters
Medium

21670

967

Add to List

Share
Given a string s, find the length of the longest substring without repeating characters.

 

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
 

Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/

func lengthOfLongestSubstring(s string) int {
    result := make(map[byte]int)
    max := 0
    index := 0
    sc := s
    for {
        if index >= len(sc) {
            substringLen := len(result)
        
            if substringLen > max {
                max = substringLen
            }
            break
        }
        c := sc[index]
        _, exists := result[c]
        if !exists {
            result[c] = index
            index++
            continue
        }
        
        substringLen := len(result)
        
        if substringLen > max {
            max = substringLen
        }
        //get back to previous occurance and restart
        sc = sc[1:]
        index = 0
        
        result = map[byte]int{}
    }

    
    return max
}