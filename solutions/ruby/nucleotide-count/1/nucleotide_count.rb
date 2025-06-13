=begin
Write your code for the 'Nucleotide Count' exercise in this file. Make the tests in
`nucleotide_count_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/nucleotide-count` directory.
=end

class Nucleotide

  def initialize(chain)
    @chain = chain
  end
  
  def Nucleotide.from_dna(chain)
    raise ArgumentError if chain.chars.any?(/[^ATCG]/)
    new(chain)
  end

  def count(char)
    @chain.count(char)
  end

  def histogram
    m = {"A"=>0, "T"=>0, "C"=>0, "G"=>0}
    @chain.chars.tally.each {|key, value| m[key] = value }    
    m
  end
 
end