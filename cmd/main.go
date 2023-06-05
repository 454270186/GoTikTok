package main

func main() {
	router := NewRouter()

	router.Run("192.168.2.8:8181")
}