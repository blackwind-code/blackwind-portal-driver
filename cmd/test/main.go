package main

import (
	"log"
	"net/http"
	"os"

	"github.com/blackwind-code/blackwind-portal-driver/internal/zerotier"
)

func main() {
	ZT_ADDR := os.Args[1]

	log.Printf("Zerotier Address: %v\n", ZT_ADDR)

	m := http.NewServeMux()
	zerotier.Init(m)

	log.Printf("[Test 0] Creating device: %v\n", ZT_ADDR)
	log.Printf("[Output 0] %v\n\n\n", zerotier.DeviceCreate(ZT_ADDR))

	log.Printf("[Test 1] Updating device: %v\n", ZT_ADDR)
	log.Printf("[Output 1] %v\n\n\n", zerotier.DeviceUpdate(ZT_ADDR, 120))

	log.Printf("[Test 2] Deleting device: %v\n", ZT_ADDR)
	log.Printf("[Output 2] %v\n\n\n", zerotier.DeviceDelete(ZT_ADDR))
}
