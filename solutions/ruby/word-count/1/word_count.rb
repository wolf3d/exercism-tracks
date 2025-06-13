=begin
Write your code for the 'Word Count' exercise in this file. Make the tests in
`word_count_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/word-count` directory.
=end
class Phrase
  def initialize(phrase)
    @phrase = phrase
  end

  def word_count
    result = Hash.new
    @phrase.scan(/\b[a-zA-Z0-9]+[']?[a-zA-Z0-9]*\b/).each { |word|
      result.has_key?(word.downcase) ? result[word.downcase] = result[word.downcase] + 1 : result[word.downcase] = 1
    }
    result
  end
end