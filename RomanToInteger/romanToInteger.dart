int romanToInt(String s) {
  Map<String, int> romanValues = {
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
  };

  List<String> validCombinations = ["IV", "IX", "XL", "XC", "CD", "CM"];
  int sum = 0;

  bool isValidCombination(String combination) => validCombinations.contains(combination);

  for (int i = 0; i < s.length; i++) {
    if (i > 0) {
      String possibleCombination = "${s[i - 1]}${s[i]}";

      if (isValidCombination(possibleCombination)) {
        sum -= romanValues[s[i - 1]]!;
        sum += (romanValues[s[i]]! - romanValues[s[i - 1]]!);
      } else {
        sum += romanValues[s[i]]!;
      }
    } else {
      sum += romanValues[s[i]]!;
    }
  }

  return sum;
}

bool twoIntegerListAreEqual(List<int> first, List<int> second) {
  if (first.length != second.length) return false;

  bool equal = true;

  for (int i = 0; i < first.length; i++) {
    equal = first[i] == second[i];

    if (!equal) break;
  }

  return equal;
}

void main() {
  List<String> romanNumerals = [
    'I',
    'II',
    'III',
    'IV',
    'V',
    'VI',
    'VII',
    'VIII',
    'IX',
    'X',
    'XI',
    'XII',
    'XIII',
    'XIV',
    'XV',
    'XVI',
    'XVII',
    'XVIII',
    'XIX',
    'XX',
    'XXX',
    'XL',
    'L',
    'LX',
    'LXX',
    'LXXX',
    'XC',
    'C',
    'CI',
    'CII',
    'CIII',
    'CIV',
    'CV',
    'CVI',
    'CVII',
    'CVIII',
    'CIX',
    'CX',
    'CL',
    'CLX',
    'CLXX',
    'CLXXX',
    'CC',
    'CCC',
    'CD',
    'D',
    'DC',
    'DCC',
    'DCCC',
    'CM',
    'M',
    'MC',
    'MCC',
    'MCCC',
    'MCD',
    'MD',
    'MDC',
    'MDCC',
    'MDCCC',
    'MCM',
    'MM',
    'MMM'
  ];

  List<int> integerValues = [
    1,
    2,
    3,
    4,
    5,
    6,
    7,
    8,
    9,
    10,
    11,
    12,
    13,
    14,
    15,
    16,
    17,
    18,
    19,
    20,
    30,
    40,
    50,
    60,
    70,
    80,
    90,
    100,
    101,
    102,
    103,
    104,
    105,
    106,
    107,
    108,
    109,
    110,
    150,
    160,
    170,
    180,
    200,
    300,
    400,
    500,
    600,
    700,
    800,
    900,
    1000,
    1100,
    1200,
    1300,
    1400,
    1500,
    1600,
    1700,
    1800,
    1900,
    2000,
    3000
  ];

  List<int> output = [];

  for (int i = 0; i < romanNumerals.length; i++) {
    output.add(romanToInt(romanNumerals[i]));
  }

  print("The output is correct: ${twoIntegerListAreEqual(integerValues, output)}");
}
