package main

import "fmt"

func strStro(haystack string, needle string) int {
    s := -1
    match := true

    if len(needle) == 0 {
        return 0        
    }

    for i,_ := range haystack{
        //fmt.Println(haystack[i])
        if(haystack[i]==needle[0]){
            match = true
            s = i
        }

        if(match == true && s != -1) {
            
            for j := 0; j < len(needle); j++{
                if(haystack[i+j] != needle[j]){
                    match = false
                    s = -1
                    break
                }
            }
            return s
        }
    }
    return s
}

func maine() {
    haystack := "aaaaa"
    needle := "aab"
    result := strStr(haystack, needle)
    fmt.Println(result)
}