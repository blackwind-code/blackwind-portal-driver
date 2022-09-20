package zerotier

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	Client int = 130
	Server int = 120
)

type HttpRes struct {
	status string
	body   string
}

// curl -X POST \"http://localhost:9993/controller/network/{:s}/member/{:s}\" -H \"X-ZT1-AUTH: {:s}\" -d '{{\"authorized\": true
func DeviceCreate(ztAddr string) HttpRes {
	URL := ZEROTIER_API_URL + "/controller/network/" + ZEROTIER_NETWORK_ID + "/member/" + ztAddr

	var reqJson ZT_Member_Approve
	reqJson.Authorized = true
	reqPayload, err := json.Marshal(reqJson)
	if err != nil {
		Log.Printf("Error: %v\n", err)
		return HttpRes{status: "-1 ERROR", body: "Failed to connect to Zerotier API"}
	}

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(reqPayload))
	if err != nil {
		Log.Printf("Error: %v\n", err)
		return HttpRes{status: "-1 ERROR", body: "Failed to connect to Zerotier API"}
	}
	req.Header.Set("X-ZT1-AUTH", ZEROTIER_TOKEN)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		Log.Printf("Error: %v\n", err)
		return HttpRes{status: "-1 ERROR", body: "Failed to connect to Zerotier API"}
	}
	defer res.Body.Close()
	resBody, _ := ioutil.ReadAll(res.Body)

	return HttpRes{status: res.Status, body: string(resBody)}
}

// curl -X POST \"http://localhost:9993/controller/network/{:s}/member/{:s}\" -H \"X-ZT1-AUTH: {:s}\" -d '{:s}'
func DeviceUpdate(ztAddr string, ztType int) HttpRes {
	URL := ZEROTIER_API_URL + "/controller/network/" + ZEROTIER_NETWORK_ID + "/member/" + ztAddr

	var reqJson ZT_Member_Change_Tag
	reqJson.Tags = [][]int{{100, ztType}}
	reqPayload, err := json.Marshal(reqJson)
	if err != nil {
		Log.Printf("Error: %v\n", err)
		return HttpRes{status: "-1 ERROR", body: "Failed to connect to Zerotier API"}
	}

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(reqPayload))
	if err != nil {
		Log.Printf("Error: %v\n", err)
		return HttpRes{status: "-1 ERROR", body: "Failed to connect to Zerotier API"}
	}
	req.Header.Set("X-ZT1-AUTH", ZEROTIER_TOKEN)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		Log.Printf("Error: %v\n", err)
		return HttpRes{status: "-1 ERROR", body: "Failed to connect to Zerotier API"}
	}
	defer res.Body.Close()
	resBody, _ := ioutil.ReadAll(res.Body)

	return HttpRes{status: res.Status, body: string(resBody)}
}

// curl -X POST \"http://localhost:9993/controller/network/{:s}/member/{:s}\" -H \"X-ZT1-AUTH: {:s}\" -d '{{\"authorized\": false}}'
// curl -X DELETE \"http://localhost:9993/controller/network/{:s}/member/{:s}\" -H \"X-ZT1-AUTH: {:s}\"
func DeviceDelete(ztAddr string) HttpRes {
	URL := ZEROTIER_API_URL + "/controller/network/" + ZEROTIER_NETWORK_ID + "/member/" + ztAddr

	var reqJson ZT_Member_Approve
	reqJson.Authorized = false
	reqPayload, err := json.Marshal(reqJson)
	if err != nil {
		Log.Printf("Error: %v\n", err)
		return HttpRes{status: "-1 ERROR", body: "Failed to connect to Zerotier API"}
	}

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(reqPayload))
	if err != nil {
		Log.Printf("Error: %v\n", err)
		return HttpRes{status: "-1 ERROR", body: "Failed to connect to Zerotier API"}
	}
	req.Header.Set("X-ZT1-AUTH", ZEROTIER_TOKEN)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		Log.Printf("Error: %v\n", err)
		return HttpRes{status: "-1 ERROR", body: "Failed to connect to Zerotier API"}
	}
	res.Body.Close()

	req, err = http.NewRequest("DELETE", URL, nil)
	if err != nil {
		Log.Printf("Error: %v\n", err)
		return HttpRes{status: "-1 ERROR", body: "Failed to connect to Zerotier API"}
	}
	req.Header.Set("X-ZT1-AUTH", ZEROTIER_TOKEN)
	req.Header.Set("Content-Type", "application/json")

	client = &http.Client{}
	res, err = client.Do(req)
	if err != nil {
		Log.Printf("Error: %v\n", err)
		return HttpRes{status: "-1 ERROR", body: "Failed to connect to Zerotier API"}
	}
	defer res.Body.Close()
	resBody, _ := ioutil.ReadAll(res.Body)

	return HttpRes{status: res.Status, body: string(resBody)}
}
