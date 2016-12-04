**** Dev Server is up and running!! ***
**** Navigate to 52.66.134.228:8080 to check out the Dashboard template.***
**** Post JSON Data to 52.66.134.228:8080/api to find the nearest Emergency Base.***

---------------------------------------------------------------------------------------------------------------------------

This folder contains the static website template and the source code, executable for the backend Go Server.

Run the Binary "108" to start the server. 

Incase the binary doesn't work, the server can be started by running "go run main.go NearestBase.go" from the terminal. (Golang needs to be installed).

The Server run on port 8080.

The Dashboard runs on the / endpoint. (localhost:8080/)
Location of user from the spot of emergency is POST'ed to the /api endpoint (localhost:8080/api)

The Server uses the Google Maps DistanceMatrix API to estimate the closest Emergency base based on traffic data.(Closest by time.)
---------------------------------------------------------------------------------------------------------------------------

The Co-ordinates of the Bases have been encoded into a Polyline.
The encoded strings are present in the "Base Station Polyline Encoded.txt" file.

---------------------------------------------------------------------------------------------------------------------------

Example of JSON sent to Server from App :
{
  "lat": "12.840639",
  "long": "80.170417",
  "tag": "emergency",
  "etype": "fire"
}

Corresponding JSON returned from Server :
{
	"District": "KANCHEEPURAM",
	"Locality": "MARAIMALAINAGAR MUNICIPALITY OFFICE",
	"Lat": "12.802211",
	"Long": "80.025733"
}

---------------------------------------------------------------------------------------------------------------------------


