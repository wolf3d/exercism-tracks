=begin
Write your code for the 'Acronym' exercise in this file. Make the tests in
`acronym_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/acronym` directory.
=end

class Acronym
  def self.abbreviate(text)
    abbreviation = ""
    text.scan(/\w+/) {|word| abbreviation.concat(word[0].upcase)}
    abbreviation
  end
end