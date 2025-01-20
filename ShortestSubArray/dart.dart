import 'dart:collection';
import 'dart:math';

List<int> prefixedSumOf(List<int> nums) {
  List<int> prefixSum = [0];

  for (int i = 1; i <= nums.length; i++) {
    prefixSum.add(prefixSum[i - 1] + nums[i - 1]);
  }
  return prefixSum;
}

int shortestSubArray(List<int> nums, int k) {
  DoubleLinkedQueue<int> deque = DoubleLinkedQueue<int>();
  int numSize = nums.length;

  int minLength = numSize + 1;

  List<int> prefixSum = prefixedSumOf(nums);
  print(prefixSum);

  for (int i = 0; i <= numSize; i++) {
    while (deque.isNotEmpty && prefixSum[i] - prefixSum[deque.first] >= k) {
      minLength = min(minLength, i - deque.first);
      deque.removeFirst();
    }

    while (deque.isNotEmpty && prefixSum[i] <= prefixSum[deque.last]) {
      deque.removeLast();
    }

    deque.add(i);
  }

  if (minLength == numSize + 1) return -1;

  return min(minLength, numSize + 1);
}

void main() {
  List<int> nums = [56, 21, 56, 35, 9];
  int k = 61;

  int result = shortestSubArray(nums, k);

  print("Result: $result");
}
