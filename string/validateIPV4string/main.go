package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Correcting the input string to use a forward slash instead of backslash
	fmt.Println(validateIPv4("123.123.123.123/32"))
}

// Function to validate IPv4 address with optional CIDR notation
func validateIPv4(str string) string {
	// Split by '/' to separate the IP address and the CIDR notation if present
	parts := strings.Split(str, "/")
	ip := parts[0]
	fmt.Println(ip)
	cidr := ""
	if len(parts) > 1 {
		cidr = parts[1]
	}

	// Split IP address by '.'
	ipv4 := strings.Split(ip, ".")
	if len(ipv4) != 4 {
		return "Invalid IP address format: wrong number of octets"
	}

	// Check each part of the IP address
	for _, part := range ipv4 {
		num, err := strconv.Atoi(part)
		if err != nil || num < 0 || num > 255 {
			return fmt.Sprintf("Invalid IP address part: %s", part)
		}
	}

	// If CIDR is provided, validate it
	if cidr != "" {
		cidrNum, err := strconv.Atoi(cidr)
		if err != nil || cidrNum < 0 || cidrNum > 32 {
			return fmt.Sprintf("Invalid CIDR notation: %s", cidr)
		}
	}

	return "Valid IP address format"
}
