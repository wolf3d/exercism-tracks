=begin
Write your code for the 'Gigasecond' exercise in this file. Make the tests in
`gigasecond_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/gigasecond` directory.
=end

class Gigasecond
  Gigasecond::GIGASECOND = 1000000000
  def Gigasecond.from(start)
    start + Gigasecond::GIGASECOND
  end
end