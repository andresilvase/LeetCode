func mergeAlternately(_ word1: String, _ word2: String) -> String {
    var output: String = ""

    for i in 0..<max(word1.count, word2.count) {
        output += "\(word1[word1.index(word1.startIndex, offsetBy: i)])"
        output += "\(word2[word2.index(word2.startIndex, offsetBy: i)])"

        let lastCharW1: Bool = word1.count == i + 1
        let lastCharW2: Bool = word2.count == i + 1

        if word2.count != word1.count {
            if lastCharW1 {
                output += word2.dropFirst(i + 1)
                break
            }

            if lastCharW2 {
                output += word1.dropFirst(i + 1)
                break
            }
        }

    }
    return output
}
