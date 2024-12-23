import 'dart:math';

class Solution {
  int maxProfit(List<int> prices) {
    int buyValue = prices.first;
    int gain = 0;

    for (int i = 1; i < prices.length; i++) {
      buyValue = min(buyValue, prices[i]);
      gain = max(gain, prices[i] - buyValue);
    }

    return gain;
  }
}

void main() {
  Solution s = Solution();
  List<int> prices = [7, 1, 5, 3, 6, 4];

  print("Output: ${s.maxProfit(prices)}");
}
