class Solution {
  List<int> spiralOrder(List<List<int>> matrix) {
    List<int> output = [];
    int matrixLength = matrix.length * matrix[0].length;

    for (int i = 0; output.length < matrixLength; i++) {
      for (int j = i; j < matrix[i].length - i; j++) {
        output.add(matrix[i][j]);
      }
      if (output.length >= matrixLength) break;

      int k = i + 1;
      int lineLength = matrix[k].length;

      while (k < matrix.length - i) {
        output.add(matrix[k][lineLength - i - 1]);
        k++;
      }

      if (output.length >= matrixLength) break;

      for (int m = lineLength - i - 2; m >= i; m--) {
        output.add(matrix[k - 1][m]);
      }
      if (output.length >= matrixLength) break;

      for (int n = k - 2; n > i; n--) {
        output.add(matrix[n][i]);
      }
      if (output.length >= matrixLength) break;
    }

    return output;
  }
}

void main() {
  Solution s = Solution();

  List<List<int>> matrix = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
  ];

  List<List<int>> matrix2 = [
    [01, 02, 03, 04],
    [05, 06, 07, 08],
    [09, 10, 11, 12],
    [13, 14, 15, 16],
    [17, 18, 19, 20],
    [21, 22, 23, 24],
  ];

  print(s.spiralOrder(matrix));
  print(s.spiralOrder(matrix2));
}
