module Port
  # TODO: define the 'IDENTIFIER' constant
  Port::IDENTIFIER = :PALE

  def self.get_identifier(city)
    city[0,4].upcase.to_sym
  end

  def self.get_terminal(ship_identifier)
    sip = ship_identifier.to_s[0,3].upcase
    case sip
      when "OIL", "GAS"
        :A
      else
        :B
    end
  end
end
