package main

func main() {
	ow := NewOpenWeather("524901", "26f04947ccd1591b1fdea1d787dee4c7")

	server := NewAPIServer(":8080")
	server.AddExternalAPI("open-weather", ow)
	server.Run()
}