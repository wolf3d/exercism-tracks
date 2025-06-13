=begin
Write your code for the 'Raindrops' exercise in this file. Make the tests in
`raindrops_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/raindrops` directory.
=end

class Raindrops
  F = {3=>"Pling",5=>"Plang",7=>"Plong"}
  def self.convert(num)
    val = ""
    F.each do |factor,value|
      if num % factor == 0
        val += value
      end
    end
  val.empty? ? num.to_s : val
  end
end