=begin
Write your code for the 'Triangle' exercise in this file. Make the tests in
`triangle_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/triangle` directory.
=end

class Triangle
  def initialize(sides)
    @sides = sides
  end

  def equilateral?
    return false if not_a_triangle?
    @sides[0] == @sides[1] && @sides[1] == @sides[2]
  end

  def scalene?
    return false if not_a_triangle?
    @sides[0] != @sides[1] && @sides[1] != @sides[2] && @sides[0] != @sides[2]
  end

  def isosceles?
    return false if not_a_triangle?
    #2*@sides[0] >= @sides[1] + @sides[2] || 2*@sides[1] >= @sides[0] + @sides[2] || 2*@sides[2] >= @sides[1] + @sides[0]
    @sides[0] == @sides[1] || @sides[1] == @sides[2] || @sides[2] == @sides[0]
  end

  def not_a_triangle?
    @sides[0] == 0 || @sides[1] == 0 || @sides[2] == 0 || @sides[0] + @sides[1] < @sides[2] || @sides[0] + @sides[2] < @sides[1] || @sides[1] + @sides[2] < @sides[0]
  end
end