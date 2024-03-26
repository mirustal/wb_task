package contains

func ContainsSym(str string) (bool) {
    chars := []rune(str)
    length := len(chars)
    for i := 0; i < length; i++ {
        for j := i + 1; j < length; j++ {
            if chars[j] == chars[i] {
                return false
            }
        }
    }
    return true
}