=begin
Write your code for the 'Matrix' exercise in this file. Make the tests in
`matrix_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/matrix` directory.
=end
class Matrix
  def initialize(lines)
    @rows = []
    @columns = []
    lines.split("\n").each do |raw_row|
      r = raw_row.split(' ').map(&:to_i)      
      rows.push(r)
    end    
   
    @rows.each do |row|

        row.map.with_index do |number, nindex|
        if @columns[nindex] == nil
            @columns[nindex] = []
        end
        @columns[nindex].push(number)    
        end
    end
  end

  def columns
    @columns
  end

  def rows
    @rows
  end
end
