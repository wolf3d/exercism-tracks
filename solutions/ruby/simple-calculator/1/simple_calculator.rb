class SimpleCalculator
  class UnsupportedOperation < StandardError
  end

  ALLOWED_OPERATIONS = ['+', '/', '*'].freeze

  def self.calculate(first_operand, second_operand, operation)
    
    if !ALLOWED_OPERATIONS.include?(operation)
      raise UnsupportedOperation
    elsif first_operand.class.to_s != 'Integer' || second_operand.class.to_s != 'Integer'
      raise ArgumentError
    end

  if operation == '+'
    "#{first_operand} + #{second_operand} = #{first_operand + second_operand}"
  elsif operation == '/'
    "#{first_operand} / #{second_operand} = #{first_operand / second_operand}"
  elsif operation == '*'
    "#{first_operand} * #{second_operand} = #{first_operand * second_operand}"
  end
  rescue ZeroDivisionError
    "Division by zero is not allowed."
  end
end

