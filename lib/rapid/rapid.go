package rapid

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type JokeResponse struct {
	Body   []*Joke `json:"body"`
	Status bool    `json:"success"`
}

type Joke struct {
	ID        string `json:"_id"`
	Punchline string `json:"punchline"`
	Setup     string `json:"setup"`
	Type      string `json:"type"`
}

func GetRandomJoke() *Joke {
	url := "https://dad-jokes.p.rapidapi.com/random/joke"

	req, _ := http.NewRequest("GET", url, nil)
	apiKey := os.Getenv("API_KEY_DAD_JOKES")

	req.Header.Add("x-rapidapi-key", apiKey)
	req.Header.Add("x-rapidapi-host", "dad-jokes.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	resp := JokeResponse{}
	json.Unmarshal(body, &resp)

	return resp.Body[0]
}
