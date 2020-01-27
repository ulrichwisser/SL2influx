# SL2influx
Save current trafik data from Stockholm public transport (SL) to influxdb.

This program queries the SL real time traffic data API and saves the next arrival of each line to influxdb.

The program takes a few options
- -conf <filename> must be given and contains the configuration in YAML syntax
- -verbose prints out information on the program run
- -dryrun skips saving data to influxdb
  
## Install
Prerequistes are of course an installation of the Go language https://golang.org/
and the following modules
- YAML https://gopkg.in/yaml.v2
- InfluxDB client https://github.com/influxdata/influxdb1-client 

## Configuration options
```
serverroot: http://api.sl.se/api2/realtimedeparturesV4.json
apikey: example-key
siteid: 1
sitedirection: 1
traveltype: 1

influxserver: http://127.0.0.1:8086
influxdb: SL
influxuser: SLuser
influxpasswd: SLpasswd
```

- serverroot is the URL of the real time traffic API (use the value provided)
- apikey You have to get your own API key at https://trafiklab.se
- siteid Trafiklab provides an API to find the siteid of the stop you want to monitor
- sitedirection is used to distinguish between side of the road for bus stops.
- traveltype is the type of transport to be saved (0=Metros,1=Buses,2=Trains,3=Trams,4=Ships)


- influxserver gives the URL to the influxdb server including port
- influxdb name of the database to save data to (must exist)
- influxuser username to access influxdb
- influxpasswd password to access influxdb

## Influx
The database must be created before you start to save data.
```influx -execute 'create database "SL"'```

All data is saved to a measurement called "SLlatest".

The measurement contains the following tags
- linenumber
- destination
- siteid
- sitedirection

and the following values
- TimeTableDateTime string in format YYYYMMDDTHH:MM:SS
- TimeTableTimestamp unix timestamp computed from TimeTableDateTime
- ExpectedDateTime strimg in format YYYYMMDDTHH:MM:SS
- ExpectedTimestamp unix timestamp computed from ExpectedDateTime
- DisplayTime string in varying format as is displayed on signs at stops

All times are given in Stockholm local time. Display time can be in format HH:MM or xx min or "Nu" (=Swedish for now).

## Future development
Please let me know if you would like to cover any other use cases.

