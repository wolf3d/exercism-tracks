=begin
Write your code for the 'Luhn' exercise in this file. Make the tests in
`luhn_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/luhn` directory.
=end

class Luhn
  def self.valid?(number)
    return false if number.length < 2 || number.scan(/[^0-9 ]+/).length > 0 || number.scan(/[0-9]/).length < 2
    number.scan(/[0-9]/).reverse.map.each_with_index {|val, index| index.odd? && val.to_i < 9 ? self.multiply(val.to_i) : val.to_i }.sum % 10 == 0
  end
  
  def self.multiply(num)
      v = num * 2
      v > 9 ? v - 9 : v
  end
end