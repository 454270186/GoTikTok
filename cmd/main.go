package main

func main() {
	router := NewRouter()

	router.Run("172.20.10.2:8181")
}