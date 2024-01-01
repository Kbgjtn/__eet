package cases

/*

"()[]{}" -> true

"([)]" -> false
stack: ( [ ) ]

every we see open bracket we push it to stack
then we see close bracket we pop from stack and compare
**/

func IsValid(s string) bool {
    stack := make([]rune, 0)

    if len(s) == 0 {
        return true
    }

    for _, v := range s {
        if v == '(' || v == '{' || v == '[' {
            stack = append(stack, v)
        } else if v == ')' || v == '}' || v == ']' {
            if len(stack) == 0 {
                // no other way to close
                return false
            }
            // pop it
            // get the last added
            last := stack[len(stack)-1]
            if (v == ')' && last == '(') || (v == '}' && last == '{') || (v == ']' && last == '[') {
                stack = stack[:len(stack)-1]
            } else {
                return false
            }
        }
    }

    return len(stack) == 0
}
