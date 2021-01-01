package main


func main() {
	mux := Routes()
	server := NewServer(mux)
	server.Run()
}