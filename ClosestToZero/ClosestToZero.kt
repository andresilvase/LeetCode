import kotlin.math.abs

fun findClosestNumber(nums: IntArray): Int {
    var closest = abs(nums[0])
    var output = nums[0]

    nums.forEach { currentNumber ->
        val absValue = abs(currentNumber)

        if (absValue < closest) {

            closest = absValue
            output = currentNumber
        } else if (absValue == closest && currentNumber > output) {
            output = currentNumber
        }
    }

    return output
}
