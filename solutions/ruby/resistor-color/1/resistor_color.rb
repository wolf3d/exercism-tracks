=begin
Write your code for the 'Resistor Color' exercise in this file. Make the tests in
`resistor_color_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/resistor-color` directory.
=end

class ResistorColor

  ResistorColor::COLORS = %w[black brown red orange yellow green blue violet grey white]

  def ResistorColor.color_code(code)
    ResistorColor::COLORS.index(code)
  end  
end