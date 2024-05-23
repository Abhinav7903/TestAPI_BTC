package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Function to get transaction IDs
func getTxids(btcaddress string) ([]string, error) {
	txurl := fmt.Sprintf("https://vayu.hornet.technology/api/address/%s", btcaddress)

	resp, err := http.Get(txurl)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var resdata map[string]interface{}
	if err := json.Unmarshal(body, &resdata); err != nil {
		return nil, fmt.Errorf("error parsing JSON response: %v", err)
	}

	txHistory, ok := resdata["txHistory"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("txHistory not found in response")
	}

	txids, ok := txHistory["txids"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("txids not found in txHistory")
	}

	var txidStrings []string
	for _, txid := range txids {
		txidStr, ok := txid.(string)
		if ok {
			txidStrings = append(txidStrings, txidStr)
		}
	}

	return txidStrings, nil
}

// Function to get transaction details and extract timestamp
func getTransactionTimestamp(txid string) (string, error) {
	url := fmt.Sprintf("https://vayu.hornet.technology/api/tx/%s", txid)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return "", fmt.Errorf("error parsing JSON response: %v", err)
	}

	timestamp, ok := data["time"].(float64)
	if !ok {
		return "", fmt.Errorf("time field not found or invalid")
	}

	humanReadableTime := time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
	return humanReadableTime, nil
}

// Handler to get the last active timestamp for a Bitcoin address
func getLastActiveTimestamp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	btcaddress := vars["btcaddress"]

	txids, err := getTxids(btcaddress)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting txids: %v", err), http.StatusInternalServerError)
		return
	}

	if len(txids) == 0 {
		http.Error(w, fmt.Sprintf("No transactions found for address: %s", btcaddress), http.StatusNotFound)
		return
	}

	latestTxid := txids[0]
	timestamp, err := getTransactionTimestamp(latestTxid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting transaction timestamp: %v", err), http.StatusInternalServerError)
		return
	}

	responseData := map[string]interface{}{
		"message":               "success",
		"last_active_timestamp": timestamp,
	}

	responseJSON, err := json.MarshalIndent(responseData, "", "  ")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello! This is the home page.")
	w.Write([]byte("\n"))
	w.Write([]byte("To get the last active timestamp for a transaction ID, make a GET request to /api/address/{btcaddress}"))
	w.Write([]byte("\n"))
	w.Write([]byte("Example: /api/address/1a2b3c4d5e6f7g8h9i0j"))

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	r.HandleFunc("/api/address/{btcaddress}", getLastActiveTimestamp).Methods("GET")

	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
