=begin
Write your code for the 'Scrabble Score' exercise in this file. Make the tests in
`scrabble_score_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/scrabble-score` directory.
=end
class Scrabble 
POINTS = { A: 1, E: 1, I: 1, O: 1, U: 1, L: 1, N: 1, R: 1, S: 1, T:1,
D: 2, G:2,
B: 3, C: 3, M: 3, P:3,
F: 4, H: 4, V: 4, W: 4, Y:4,
K:5,
J: 8, X:8,
Q: 10, Z:10
}

def initialize(word)
  @word = word
end

def Scrabble.score(word)
  if word == nil
    0
  else
    word.scan(/[a-zA-Z]/).map {|item| item.upcase}.map {|item| POINTS[item.to_sym]}.sum
  end  
end

def score
  Scrabble.score(@word)
end
  
end