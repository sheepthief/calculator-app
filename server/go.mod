module example.com/server

go 1.25.3

replace example.com/RPN => ../RPN

replace example.com/addition => ../addition

require example.com/RPN v0.0.0-00010101000000-000000000000 // indirect

require (
	example.com/addition v0.0.0-00010101000000-000000000000 // indirect
	example.com/division v0.0.0-00010101000000-000000000000 // indirect
	example.com/multiplication v0.0.0-00010101000000-000000000000 // indirect
	example.com/subtraction v0.0.0-00010101000000-000000000000 // indirect
    example.com/calculator v0.0.0-00010101000000-000000000000
)

replace example.com/subtraction => ../subtraction

replace example.com/division => ../division

replace example.com/multiplication => ../multiplication

replace example.com/calculator => ../calculator 