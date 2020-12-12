package arrays

func StringIn(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }

    return false
}

func IntIn(a int, list []int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func ByteIn(a byte, list []byte) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func IntMax(list []int) int {
    var max int
    for i, b := range list {
        if i == 0 {
            max = b
        } else if b > max {
            max = b
        }
    }
    return max
}

func IntMin(list []int) int {
    var min int
    for i, b := range list {
        if i == 0 {
            min = b
        } else if b < min{
            min = b
        }
    }
    return min
}
