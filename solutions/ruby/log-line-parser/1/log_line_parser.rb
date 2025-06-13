class LogLineParser
  def initialize(line)
    @line = line.strip
  end

  def message
    @line.match(':[ \t]*([a-zA-Z ]+)')[1]    
  end

  def log_level
    @line.match('\[([a-zA-Z]+)\]')[1].downcase
  end

  def reformat
    "#{message} (#{log_level})"
  end
end
