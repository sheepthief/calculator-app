module example.com/RPN

go 1.25.3

replace example.com/addition => ../addition

require (
	example.com/addition v0.0.0-00010101000000-000000000000
	example.com/division v0.0.0-00010101000000-000000000000
	example.com/multiplication v0.0.0-00010101000000-000000000000
	example.com/subtraction v0.0.0-00010101000000-000000000000
)

replace example.com/subtraction => ../subtraction

replace example.com/multiplication => ../multiplication

replace example.com/division => ../division
