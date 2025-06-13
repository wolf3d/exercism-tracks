=begin
Write your code for the 'Series' exercise in this file. Make the tests in
`series_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/series` directory.
=end

class Series
  def initialize(seq)
    @seq = seq
  end
def slices(len)
  raise ArgumentError if len > @seq.length
  @seq.chars.each_cons(len).map {|slice| slice.join}
end
end