package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
	"errors"
	"regexp"
	"test/callNotion"
)

type AlgCollection struct {
	Algs []Alg
}

type Alg struct {
	Letter string 
	Process string
}

func main() {

    raw, err := ioutil.ReadFile("./myalgs.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    // var ac AlgCollection

    // json.Unmarshal(raw, &ac)

	// for _, alg := range ac.Algs {
    //     fmt.Println(alg.Letter)
    //     fmt.Println(alg.Process)
    // }

	var result map[string]interface{}
	json.Unmarshal(raw, &result)

	algs := AlgCollection{}
	for key, value := range result {
		if value != nil {
			// fmt.Println(key, value.(string))
			algs.Algs = append(algs.Algs, Alg{key, value.(string)})
		}
	}

	algRawProperties := callNotion.AlgRawPropertyCollection{}
	for _, alg := range algs.Algs {
		fmt.Println("ALG", alg)
		first, err := getFirstSticker(&alg)
		if err != nil {panic(err)}
		fmt.Println("FIRST STICKER: ", first)
		second, err := getSecondSticker(&alg)
		if err != nil {panic(err)}
		fmt.Println("SECOND STICKER: ", second)
		setup, err := getSetup(&alg)
		if err != nil {panic(err)}
		fmt.Println("SETUP: ", setup)
		commFirstHalf, err := getCommutatorFirstHalf(&alg)
		if err != nil {panic(err)}
		fmt.Println("COMMUTATOR FIRST HALF: ", commFirstHalf)
		commSecondHalf, err := getCommutatorSecondHalf(&alg)
		if err != nil {panic(err)}
		fmt.Println("COMMUTATOR SECOND HALF: ", commSecondHalf)
		fmt.Println("\n")

		algRawProperties.Algs = append(algRawProperties.Algs, callNotion.AlgRawProperty{alg.Letter, alg.Process, first, second, setup, commFirstHalf, commSecondHalf})
	}

	callNotion.CallNotion(&algRawProperties)
}

func getFirstSticker(alg *Alg) (string, error) {
	if (len([]rune(alg.Letter)) != 2) {
		fmt.Println("\n")
		fmt.Println(len(alg.Letter))
		return "", errors.New("ステッカーの数が不正です。")
	}

	return string([]rune(alg.Letter)[:1]), nil
}

func getSecondSticker(alg *Alg) (string, error) {
	if (len([]rune(alg.Letter)) != 2) {
		fmt.Println("\n")
		fmt.Println(len(alg.Letter))
		return "", errors.New("ステッカーの数が不正です。")
	}

	return string([]rune(alg.Letter)[1:]), nil
}

func getCommutatorFirstHalf(alg *Alg) (string, error) {
	re := regexp.MustCompile(`\[([^\[\]]*),([^\[\]]*)\]`)
	res := re.FindStringSubmatch(alg.Process)
	if len(res) == 0 {
		return "", nil
	}
	
	return res[1], nil
}

func getCommutatorSecondHalf(alg *Alg) (string, error) {
	re := regexp.MustCompile(`\[([^\[\]]*),([^\[\]]*)\]`)
	res := re.FindStringSubmatch(alg.Process)
	if len(res) == 0 {
		return "", nil
	}
	
	return res[2], nil
}

func getSetup(alg *Alg) (string, error) {
	re := regexp.MustCompile(`\[([^\[\]]*):.*`)
	res := re.FindStringSubmatch(alg.Process)
	if len(res) == 0 {
		return "", nil
	}

	return res[1], nil

}