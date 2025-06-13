=begin
Write your code for the 'Space Age' exercise in this file. Make the tests in
`space_age_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/space-age` directory.
=end

class SpaceAge
  SpaceAge::EARTH_YEAR_IN_SECONDS = 31557600
  SpaceAge::ORBITAL_PERIODS = 
{
        Mercury: 0.2408467,        
        Venus: 0.61519726,
        Earth: 1.0,
        Mars: 1.8808158,
        Jupiter: 11.862615,
        Saturn: 29.447498,
        Uranus: 84.016846,
		Neptune: 164.79132,
}

  def initialize(seconds)
    @on_earth = calculate_age(seconds, "Earth")
    @on_mercury = calculate_age(seconds, "Mercury")
    @on_venus = calculate_age(seconds, "Venus")
    @on_mars = calculate_age(seconds, "Mars")
    @on_jupiter = calculate_age(seconds, "Jupiter")
    @on_saturn = calculate_age(seconds, "Saturn")    
    @on_uranus = calculate_age(seconds, "Uranus")
    @on_neptune = calculate_age(seconds, "Neptune")
  end

  def calculate_age(seconds, planet)
    seconds / (SpaceAge::EARTH_YEAR_IN_SECONDS * SpaceAge::ORBITAL_PERIODS[planet.to_sym])
  end

  def on_earth
    @on_earth
  end

def on_mercury
    @on_mercury
end

def on_venus	
    @on_venus
end

def on_mars 	
	@on_mars
end

def	on_jupiter
    @on_jupiter
end

def	on_saturn
  @on_saturn
end
  
def on_uranus
    @on_uranus 
end

def on_neptune
    @on_neptune
end
  
end