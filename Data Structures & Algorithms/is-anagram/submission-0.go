func isAnagram(s string, t string) bool {
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
    for _, v := range t {
        sm[v]--
        // same char show up in t more than it show up in s
        if sm[v] < 0 {
            return false
        }
    }
    return true
}