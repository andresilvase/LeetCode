fun isHappy(n: Int): Boolean {
    if (n == 1) return true
    if (n < 0) return false

    var num: Int = n

    val set: MutableSet<Int> = mutableSetOf(num)

    val digitSquaresSumup = { number: Int ->
        number.toString().map { (it - '0') * (it - '0') }.sum()
    }

    while (true) {
        var newNumber: Int = digitSquaresSumup(num)

        if (newNumber == 1) return true
        if (set.contains(newNumber)) return false

        num = newNumber
        set.add(num)
    }
}

fun main() {
    val num: Int = 2

    println(isHappy(num))
}
