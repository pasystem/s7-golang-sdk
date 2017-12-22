package main

import (
	"gitlab.tm-consulting.ru/avia/services/s7-sdk"

	"encoding/xml"
	"log"
)

func main() {
	request := sdk.MakeSessionAirShoppingRQ()

	request.Body.AirShoppingRQ.Party.Sender.AgentUserSender = sdk.MakeAgentUserSender(
		"OVB", "OVB902/1OVB2TR",
		"RU", "00000011Z", "1",
		"OVB902", "1001/1001A", "115")

	request.Body.AirShoppingRQ.Travelers = sdk.MakeTravelers(map[string]int{
		sdk.PASSENGER_TYPE_CODE_ADULT: 1,
	})

	request.Body.AirShoppingRQ.CoreQuery.ShoppingResponseIDs = sdk.MakeShoppingResponseIDs("S7",
		"Jv00QiuPgYJRD3xLFFUGBpWHV/0TzgowP3ti4MNPGHaIhWtLH/StrP48gUXoVPQ93KzTWlZ4007/0/1")

	request.Body.AirShoppingRQ.Metadata = sdk.MakeResultType(sdk.RESULT_TYPE_SMARTCHOICE)

	output, err := xml.MarshalIndent(request, "", "  ")
	log.Printf("out: \n%s\nerr: %v\n", output, err)
}
