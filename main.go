// main.go - Go Client with Interactive Inputs & Actions
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"habitauth"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 1. Initialize client with dashboard credentials
	client := habitauth.NewClient("TARGET_APP_NAME", "TARGET_APP_ID", "TARGET_APP_SECRET", "TARGET_PUBLIC_KEY", "1.0.0", "https://habitauth.com/api/v1")

	// 2. Handshake with server
	ok, err := client.Init("")
	if !ok || err != nil {
		fmt.Printf("[!] Init failed: %v\n", err)
		return
	}

	// 3. User Input Prompts (txtUsername, txtPassword, txtLicense)
	fmt.Print("Enter Username (txtUsername): ")
	txtUsername, _ := reader.ReadString('\n')
	txtUsername = strings.TrimSpace(txtUsername)

	fmt.Print("Enter Password (txtPassword): ")
	txtPassword, _ := reader.ReadString('\n')
	txtPassword = strings.TrimSpace(txtPassword)

	fmt.Print("Enter License Key (txtLicense): ")
	txtLicense, _ := reader.ReadString('\n')
	txtLicense = strings.TrimSpace(txtLicense)

	fmt.Println("\n[1] btnLogin  [2] btnRegister  [3] btnLicenseOnly")
	fmt.Print("Select Action: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice == "1" {
		if loggedIn, _ := client.Login(txtUsername, txtPassword); loggedIn {
			fmt.Printf("[+] Welcome %s! Expires: %s\n", client.User.Username, client.User.ExpiresAt)
			client.StartHeartbeat(30)
		} else {
			fmt.Printf("[-] Login failed: %v\n", client.LastResponse.Message)
		}
	} else if choice == "2" {
		if reg, _ := client.Register(txtUsername, txtPassword, txtLicense); reg {
			fmt.Println("[+] Registration successful! You can now log in.")
		} else {
			fmt.Printf("[-] Registration failed: %v\n", client.LastResponse.Message)
		}
	} else if choice == "3" {
		if lic, _ := client.License(txtLicense); lic {
			fmt.Printf("[+] License activated! Expires: %s\n", client.User.ExpiresAt)
			client.StartHeartbeat(30)
		} else {
			fmt.Printf("[-] License failed: %v\n", client.LastResponse.Message)
		}
	}
}