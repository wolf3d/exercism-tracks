class AssemblyLine
  def initialize(speed)
    @speed = speed
  end

  AssemblyLine::CARS_PER_HOUR = 221

  def production_rate_per_hour
    if @speed >= 1 && @speed <= 4
      @speed.to_f * CARS_PER_HOUR
    elsif @speed > 4 && @speed <= 8
      @speed.to_f * CARS_PER_HOUR * 0.9
    elsif @speed == 9
      @speed.to_f * CARS_PER_HOUR * 0.8
    elsif @speed == 10
      @speed.to_f * CARS_PER_HOUR * 0.77
    else
      0.0
    end
  end

  def working_items_per_minute
    (production_rate_per_hour / 60).to_i
  end
end
