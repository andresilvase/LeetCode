class Solution {
  List<String> summaryRanges(List<int> nums) {
    if (nums.length == 1) return [nums.first.toString()];

    List<String> output = [];
    int rangeStartIndex = 0;
    int sequenceCount = 0;

    for (int i = 1; i < nums.length; i++) {
      if (nums[i] > nums[i - 1] + 1) {
        if (sequenceCount == 0) {
          output.add("${nums[i - 1]}");
        } else {
          output.add("${nums[rangeStartIndex]}->${nums[i - 1]}");
        }

        rangeStartIndex = i;
        sequenceCount = -1;
      }
      sequenceCount++;

      if (i == nums.length - 1) {
        if (sequenceCount == 0) {
          output.add("${nums[i]}");
        } else {
          output.add("${nums[rangeStartIndex]}->${nums[i]}");
        }
      }
    }

    return output;
  }
}

void main() {
  Solution s = Solution();
  List<int> nums = [0, 2, 3, 4, 6, 8, 9];

  print("Output: ${s.summaryRanges(nums)}");
}
