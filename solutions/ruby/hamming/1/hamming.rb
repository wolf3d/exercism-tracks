=begin
Write your code for the 'Hamming' exercise in this file. Make the tests in
`hamming_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/hamming` directory.
=end
 class Hamming
   def self.compute(s1, s2)
     distance = 0
     if s1.length != s2.length
       raise ArgumentError
     elsif s1.length == 0 && s2.length == 0
     0
     end
   s1.chars.map.with_index {|val, index|
   if s2[index] != val
     distance += 1
   end
   }
     distance
   end
 end