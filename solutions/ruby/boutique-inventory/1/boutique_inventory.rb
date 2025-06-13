class BoutiqueInventory
  def initialize(items)
    @items = items
  end

  def item_names
    @items.map {|item|
      item[:name]
      }.sort
  end

  def cheap
    @items.map.select {|item|
    item[:price] < 30.00 
    }
  end

  def out_of_stock
    @items.map.select {|item|      
      item[:quantity_by_size].empty?
    }
  end

  def stock_for_item(name)
    @items.map.select {|item|
      item[:name] == name
    }.map { |itm|
    itm[:quantity_by_size]
    }[0]
  end

  def total_stock
    sum = 0
    @items.map {|item|
    item[:quantity_by_size].each {|_, val| sum += val}
    }
    sum
  end

  private
  attr_reader :items
end
