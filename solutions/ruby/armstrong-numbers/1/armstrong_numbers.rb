=begin
Write your code for the 'Armstrong Numbers' exercise in this file. Make the tests in
`armstrong_numbers_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/armstrong-numbers` directory.
=end

module ArmstrongNumbers
  module_function
  def include?(n)
    n.digits.sum {|digit| digit ** n.digits.size } == n    
  end
end