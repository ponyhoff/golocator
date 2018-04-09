package main

import (
	"github.com/ponyhoff/golocator/rest"
	"net/http"
	"fmt"
)

func parseURLParams() rest.RESTHandler {
	return func(request *rest.Request, response *rest.Response) {
		uri := request.HTTPRequest().URL.Path[len("/ip/"):]
		// nice hack to hardcode :)
		request.AddParams("ip", uri)
	}
}


func getLocation(s LocatorService) rest.RESTHandler {
	return func(request *rest.Request, response *rest.Response) {
		ipAddr, found := request.Params["ip"]
		if !found {
			response.Terminate(
				http.StatusNotFound,
				rest.NewError(
					fmt.Sprintf("requested ip address `%s` was not found.", ipAddr),
					nil,
				),
			)
		}

		addr,_ := ipAddr.(string)

		loc, err := s.GetLocation(addr)

		if err != nil {
			response.Terminate(
				http.StatusBadRequest,
				rest.NewError(err.Error(), nil),
			)
		}

		response.StatusCode = http.StatusOK
		response.ResponseBody = loc
	}

}