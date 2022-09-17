package main

import "fmt"

type Passenger struct{
	Name		 string
	TicketNumber int
	Boarded 	 bool
}
type Bus struct{
	FrontSeat Passenger
}

func main() {
	
	casey:=Passenger{"casey",222,false}
	fmt.Println(casey)

	var(
		bill 	= Passenger{Name: "bill",TicketNumber: 2222}
		ella 	= Passenger{Name: "ella",TicketNumber: 22  }
	)
	fmt.Println(ella,bill)

	var heidi Passenger
	heidi.Name		   = "heidi"
	heidi.TicketNumber = 222222
	heidi.Boarded	   = true
	fmt.Println(heidi)

	casey.Boarded = true
	bill.Boarded  = true

	if bill.Boarded{
		fmt.Println("Bill has boarded the bus")
	}
	if casey.Boarded{
		fmt.Println("Casey has boarded the bus")
	}

	heidi.Boarded = true
	bus:=Bus{heidi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name,"is in the front seat")	
}
