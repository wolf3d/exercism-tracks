=begin
Write your code for the 'Isogram' exercise in this file. Make the tests in
`isogram_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/isogram` directory.
=end

class Isogram
  def Isogram.isogram?(text)
    !text.upcase.scan(/[A-Z]/).group_by(&:itself).transform_values(&:count).any? {|item| item[1]>1}
  end
end