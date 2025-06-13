=begin
Write your code for the 'ETL' exercise in this file. Make the tests in
`etl_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/etl` directory.
=end

class ETL
  def self.transform(old)
    h=Hash.new
    old.each {|key, val| val.each {|v| h[v.downcase]=key } }
    h
  end
end