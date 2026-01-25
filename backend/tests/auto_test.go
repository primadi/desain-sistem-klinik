package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting Auto Test for Sistem Klinik Backend...")

	baseURL := "http://localhost:8080"

	// Wait for server to start if running in parallel, but here we assume it's running
	// Simple health check or list tenants check

	fmt.Println("[TEST] Fetching Tenants List...")
	resp, err := http.Get(baseURL + "/tenants")
	if err != nil {
		fmt.Printf("[FAIL] Could not connect to %s: %v\n", baseURL, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("[FAIL] Expected status 200, got %d\n", resp.StatusCode)
		os.Exit(1)
	}

	fmt.Println("[PASS] Fetching Tenants List")

	// Add more assertions here

	fmt.Println("All tests passed!")
}
