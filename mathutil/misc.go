package mathutil


func ReverseInt(number int) (reversed int) {
    for reversed = 0; number != 0; number /= 10 {
        reversed = reversed * 10 + number % 10
    }
    return
}


func Dummy ()  {
}
