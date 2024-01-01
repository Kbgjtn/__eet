package cases

import (
    "sort"
)


func FindContentChildren(g []int, s []int) (r int) {
    sort.Ints(g)
    sort.Ints(s)

    for i := 0; i < len(s) && r < len(g); i++ {
        if s[i] >= g[r] {
            r++
        }
    }

    return
}
