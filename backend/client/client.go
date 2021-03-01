package client

import (
	b64 "encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type MemeClient struct {

}

type Meme struct {
	Name string
	Slug string
	Url string
}

func NewMemeClient() *MemeClient {
	return &MemeClient{}
}

func (mc *MemeClient) GetMemes() []Meme {
	return []Meme{
		{
			Name: "Disaster Girl",
			Slug: "Disaster-Girl",
			Url: "http://apimeme?name=Disaster-Girl",
		},
		{
			Name: "10 Guy",
			Slug: "10-Guy",
			Url: "http://apimeme?name=10-Guy",
		},
		{
			Name: "Afraid To Ask Andy",
			Slug: "Afraid-To-Ask-Andy",
			Url: "http://apimeme?name=Afraid-To-Ask-Andy",
		},
		{
			Name: "Ancient Aliens",
			Slug: "Afraid-To-Ask-Andy",
			Url: "http://apimeme?name=Ancient-Aliens",
		},
		{
			Name: "Archer",
			Slug: "Archer",
			Url: "http://apimeme?name=Ancient-Aliens",
		},{
			Name: "Bad Luck Brian",
			Slug: "Bad-Luck-Brian",
			Url: "http://apimeme?name=Ancient-Aliens",
		},
	}
}

func (mc *MemeClient) CreateMeme(name string,top string,bottom string) string {

	r,err := http.Get(buildUrl(name,top,bottom))
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	photoBase64 := b64.StdEncoding.EncodeToString(responseData)

	return  photoBase64
}

func buildUrl(name string,top string,bottom string) string {

	q := "meme=" + name + "&top=" + top + "&bottom=" + bottom

	u := url.URL{
		Scheme:   "https",
		Host:     "apimeme.com",
		Path:     "meme",
		RawQuery: q,
	}

	log.Printf("url built for new meme => %s",u.String())
	
	return u.String()
}
