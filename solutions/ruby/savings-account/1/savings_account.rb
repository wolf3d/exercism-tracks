module SavingsAccount
  def self.interest_rate(balance)
    if balance < 0
      -3.213
    elsif balance >= 0 && balance < 1000
    0.5
    elsif balance >= 1000 && balance < 5000
    1.621
    elsif balance >= 5000
    2.475
    end
  end

  def self.annual_balance_update(balance)
    if balance.negative?
      -(self.interest_rate(balance).abs / 100 * balance.abs) + balance
    else
    (self.interest_rate(balance) / 100 * balance) + balance
    end
  end

  def self.years_before_desired_balance(current_balance, desired_balance)
    years = 0
    balance = current_balance
    until balance >= desired_balance
      years += 1
      balance = self.annual_balance_update(balance)
    end
  years
  end
end
