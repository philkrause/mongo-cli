package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bobziuchkovski/digest"
)

func main() {
	TokenUN, _ := os.LookupEnv("MONGO_UN")
	TokenKEY, _ := os.LookupEnv("MONGO_KEY")
	GroupID, _ := os.LookupEnv("MONGO_GROUP_ID")
	URL := `https://cloud.mongodb.com/api/atlas/v1.0/groups/` + GroupID + `/clusters/HarmonyProd/snapshots`
	// currentDate := time.Now()
	// TODAY := currentDate.Format("2006-01-02 15:04:05")

	t := digest.NewTransport(TokenUN, TokenKEY)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := t.RoundTrip(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)

	var data Data
	if err := json.Unmarshal(b, &data); err != nil {
		fmt.Println(err)
	}

	for i := 0; i <= len(data.Results)-1; i++ {
		if data.Results[i].Complete != true {
			fmt.Println("There was a failed backup on Cluster: ", data.Results[i].ID, "@", data.Results[i].Created.Date)
		} else {
			fmt.Println(data.Results[i].ID, ": Backup was completed on ", data.Results[i].Created.Date)
		}
	}

}
