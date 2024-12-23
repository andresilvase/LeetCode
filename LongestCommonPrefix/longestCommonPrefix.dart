import 'dart:math';

import 'test_use_case.dart';

class Solution {
  String longestCommonPrefix(List<String> strs) {
    if (strs.first == "" || strs.length > 1 && strs[1] == "") return "";
    if (strs.length == 1) return strs.first;

    String first = strs.first;
    String second = strs[1];
    String prefix = "";

    for (int i = 0; i < min(first.length, second.length); i++) {
      if (first[i] == second[i]) {
        prefix += first[i];
      } else {
        break;
      }
    }

    if (prefix == "") return prefix;

    for (int i = 2; i < strs.length; i++) {
      prefix = prefix.substring(0, min(strs[i].length, prefix.length));

      while (strs[i].substring(0, prefix.length) != prefix) {
        prefix = prefix.substring(0, prefix.length - 1);
      }
    }

    return prefix;
  }
}

void main() {
  Solution s = Solution();
  for (final useCase in longestCommonPrefixUseCase) {
    print("Output: ${s.longestCommonPrefix(useCase)}");
  }
}
