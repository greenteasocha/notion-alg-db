package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"github.com/joho/godotenv"
)

/**
* goでのPOST REQUESTとnotionへ送るべきbodyの確認用
**/
func main() {
    url := "https://api.notion.com/v1/pages/"
    fmt.Println("URL:>", url)

	// database_idは正式なものを
    var jsonStr = []byte(`{
		"parent": {
			"database_id": "***dummy-databese-id***"
		},
		"properties": {
			"Tags": {
			   "multi_select": [
					{
						"name": "LFD"
					},
					{
						"name": "LBD"
					}
				]
			},
			"Name": {
				"title": [
			  		{
						"text": {
				  			"content": "R D' : [U , R' D R]"
						}
			  		}
				]		
		  	},
			"Column": {
				"rich_text": [
					{
						"type": "text",
						"text": {
							"content": "abctest",
							"link": null
						}
					}
				]	
		  	}
		}
	}`)


	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	} 
	bearer_token := "Bearer " + os.Getenv("NOTION_INTEGTATION_TOKEN")
	
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer_token)
	req.Header.Add("Notion-Version", "2021-05-13")
	client := &http.Client{}
    resp, err := client.Do(req)

    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
	

    fmt.Println("response Status:", resp.Status)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}