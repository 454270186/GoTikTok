package main

func main() {
	router := NewRouter()

	router.Run("127.0.0.1:8181")
}