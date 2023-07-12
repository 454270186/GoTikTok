package main

// test user
func main() {
	router := NewRouter()

	// router.Run("172.20.10.2:8181") // 热点
	// router.Run("192.168.2.44:8181") // 家里
	router.Run("10.14.13.212:8181") // BG
}