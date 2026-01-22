# Basic transformation
nums = [1, 2, 3, 4]
result = nums.map { |n| n * 2 }
# => [2, 4, 6, 8]

# String transformation
langs = ["ruby", "go", "java"]
langs.map { |l| l.upcase }
# => ["RUBY", "GO", "JAVA"]

# Using do...end
prices = [100, 200, 300]
prices.map do |p|
  p - (p * 0.1)
end
# => [90.0, 180.0, 270.0]

# Map with hash
users = { alice: 21, bob: 25 }
users.map { |name, age| "#{name}: #{age}" }
# => ["alice: 21", "bob: 25"]

# map! modifies original array
nums = [1, 2, 3]
nums.map! { |n| n * 10 }
# => [10, 20, 30]

# map vs each
[1, 2, 3].each { |n| n * 2 }
# => [1, 2, 3] (no change)

# map vs select
[1, 2, 3, 4].select { |n| n.even? }
# => [2, 4]
