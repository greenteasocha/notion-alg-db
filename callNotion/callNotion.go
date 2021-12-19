package callNotion

import (
    "fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"bytes"

	"github.com/joho/godotenv"
)

type AlgRawPropertyCollection struct {
	Algs []AlgRawProperty
}

type AlgRawProperty struct {
	Letter string 
	Process string
	FirstSticker string
	SecondSticker string
	Setup string
	CommutatorFirstHalf string
	CommutatorSecondHalf string
}

type AlgNotionPropertyCollection struct {
	Algs []AlgRawProperty
}

type AlgNotionBody struct {
	Parent Parent `json:"parent"`
	Properties AlgNotionProperties `json:"properties""` 
}

type Parent struct {
	DatabeseId string `json:"database_id"`
}

type AlgNotionProperties struct {
	// Letter string 
	Process TitleBaseProperty
	FirstSticker MultiSelectBaseProperty
	SecondSticker MultiSelectBaseProperty
	Setup RichTextBaseProperty
	CommutatorFirstHalf RichTextBaseProperty
	CommutatorSecondHalf RichTextBaseProperty
}

type FirstStickerProperty struct {
	FirstSticker MultiSelectBaseProperty
}

type MultiSelectBaseProperty struct {
	MultiSelect []MultiSelect `json:"multi_select"`
}

type MultiSelect struct {
	Name string `json:"name"`
}

type ColumnProperty struct {
    Column RichTextBaseProperty
}

type RichTextBaseProperty struct {
	RichText []RichText `json:"rich_text"`
}

type RichText struct {
	Type string `json:"type"`
	Text Text `json:"text"`
}

type NameProperty struct {
    Name TitleBaseProperty
}

type TitleBaseProperty struct {
	Title []Title `json:"title"`
}

type Title struct {
	Text Text `json:"text"`
}

type Text struct {
	Content string `json:"content"`
}

func CallNotion(algs *AlgRawPropertyCollection) {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	} 
	parent_database_id := os.Getenv("PARENT_DATABASE_ID")
	bearer_token := "Bearer " + os.Getenv("NOTION_INTEGTATION_TOKEN")
	fmt.Println(parent_database_id)

	for _, alg := range algs.Algs {
		parent := Parent{parent_database_id}
		properties := AlgNotionProperties{}
		fmt.Println(alg)
		properties.Process = TitleBaseProperty{Title: []Title{{Text: Text{alg.Process}}}}
		properties.FirstSticker = MultiSelectBaseProperty{MultiSelect: []MultiSelect{{Name: alg.FirstSticker}}}
		properties.SecondSticker = MultiSelectBaseProperty{MultiSelect: []MultiSelect{{Name: alg.SecondSticker}}}
		properties.Setup = RichTextBaseProperty{RichText: []RichText{{Type: "text", Text: Text{alg.Setup}}}}
		properties.CommutatorFirstHalf = RichTextBaseProperty{RichText: []RichText{{Type: "text", Text: Text{alg.CommutatorFirstHalf}}}}
		properties.CommutatorSecondHalf = RichTextBaseProperty{RichText: []RichText{{Type: "text", Text: Text{alg.CommutatorSecondHalf}}}}

		body := AlgNotionBody{Parent: parent, Properties: properties}
		jsonString, _ := json.Marshal(body)
		fmt.Println(string(jsonString1))


		url := "https://api.notion.com/v1/pages/"
		fmt.Println("URL:>", url)

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonString))
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
		resBody, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(resBody))
	}

}