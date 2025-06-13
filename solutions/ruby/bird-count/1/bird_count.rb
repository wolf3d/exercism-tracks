class BirdCount
  def self.last_week
    #@@birds_last_week
    [0, 2, 5, 3, 7, 8, 4]
  end

  def initialize(birds_per_day)
    #@@birds_last_week = [0, 2, 5, 3, 7, 8, 4]
    @birds_per_day = birds_per_day
  end

  def yesterday
    @birds_per_day[-2]
  end

  def total
    sum = 0
    @birds_per_day.each do |bird_count_on_day|
      sum += bird_count_on_day
    end
    sum
  end

  def busy_days
    @birds_per_day.count {|bird_count_on_day| bird_count_on_day >= 5}      
  end

  def day_without_birds?
    @birds_per_day.any? {|bird_count_on_day| bird_count_on_day == 0}      
  end
end
