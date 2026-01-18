# ENUM 1: Integer enum using constants
class Status
  PENDING  = 0
  APPROVED = 1
  REJECTED = 2
end

# ENUM 2: String enum (API friendly)
class PaymentMethod
  CREDIT_CARD = "credit_card"
  UPI         = "upi"
  NET_BANKING = "net_banking"
end

# ENUM 3: Hash-based enum (frozen for safety)
ROLE = {
  admin: "ADMIN",
  user: "USER",
  guest: "GUEST"
}.freeze

# ENUM validation
def valid_status?(status)
  [Status::PENDING, Status::APPROVED, Status::REJECTED].include?(status)
end

# -------- Usage (single flow) --------
status = Status::APPROVED
payment = PaymentMethod::UPI
role = ROLE[:admin]

puts "Status: #{status}"
puts "Valid Status? #{valid_status?(status)}"
puts "Payment Method: #{payment}"
puts "Role: #{role}"
