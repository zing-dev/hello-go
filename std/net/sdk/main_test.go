package main

import "testing"

func TestServer(t *testing.T) {
	NewSdkServer()
	group.Wait()
}
func TestClient(t *testing.T) {
	run()
}
