=begin
Write your code for the 'Protein Translation' exercise in this file. Make the tests in
`protein_translation_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/protein-translation` directory.
=end

class InvalidCodonError < StandardError
end

class Translation

  CODON_PROTEIN_MAP = {
    AUG: 	"Methionine",
UUU: 	"Phenylalanine",
UUC: 	"Phenylalanine",
UUA: 	"Leucine",
UUG: 	"Leucine",   
UCU: 	"Serine",
UCC: "Serine",
UCA: "Serine",
UCG: "Serine",    
UAU: 	"Tyrosine",
UAC: 	"Tyrosine",    
UGU:"Cysteine",
UGC:"Cysteine",    
UGG:"Tryptophan",
UAA:"STOP",
UAG:"STOP",
UGA:"STOP"    
  }


  
  def self.of_codon(codon)
    CODON_PROTEIN_MAP[codon.to_sym]
  end

  def self.of_rna(strand)
    codons = strand.chars.each_slice(3).map {|item|
      item.join }

    codons.each_with_object([]) {|i, a|
      codon = self.of_codon(i)
    raise InvalidCodonError if codon == nil
      
    return a if codon == "STOP" 
    a.push(codon)
    } 
  end
end