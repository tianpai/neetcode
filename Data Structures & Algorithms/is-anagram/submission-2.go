func isAnagram(s string, t string) bool {
    // check len first, anagram must be equal length.
    if len(s) != len(t){
        return false
    }
    sm := make(map[rune]int)

    //make a map of s: k -> char, int frequency
    for _, c := range s {
        sm[c]++
    }
    // t must cancel out the map perfectly. 
    // if key does not in sm, false
    // if char appears more than it does in s, false
    // > [!NOTE]
    // this will not work if we did not check length first
    //  If t is shorter than s (e.g., s="aa", t="a"), 
    //  the loop finishes without error (no negative counts), 
    //  but s has leftover characters. 
    //  Without the length check, return true incorrectly.
    for _, v := range t {
        // even if v is not in s map, in go sm[v] is zero value
        // so for the char in t that is not in s -> sm[v]-- is -1 
        sm[v]--
        // same char show up in t more than it show up in s
        if sm[v] < 0 {
            return false
        }
    }
    return true
}