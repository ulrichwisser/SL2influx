package main

type ResponseOfDepartures struct {
	StatusCode    int
	Message       string
	ExecutionTime int
	ResponseData  ResponseData
}

type ResponseData struct {
	LatestUpdate        string
	DataAge             int
	Metros              []Response
	Buses               []Response
	Trains              []Response
	Trams               []Response
	Ships               []Response
	StopPointDeviations []Response
}

type Response struct {
	GroupOfLine          string
	TransportMode        string
	LineNumber           string
	Destination          string
	JourneyDirection     int
	StopAreaName         string
	StopAreaNumber       int
	StopPointNumber      int
	StopPointDesignation string
	TimeTabledDateTime   string
	ExpectedDateTime     string
	DisplayTime          string
	JourneyNumber        int
	Deviations           string
}
