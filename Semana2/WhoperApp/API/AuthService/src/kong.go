package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

/*
 * AddPlugin add a Kong plugin to api
 */
func AddPlugin(p Plugins) error {
	_, err := http.PostForm(pluginsURL+p.Plugin+"/plugins",
		p.values)
	if err != nil {
		log.Printf("func\tAddPlugincould:\tnot set plugin\t%s:\t%v\n", p.Plugin, err)
		return err
	} else {
		log.Printf("func\tAddPluginPlugin:\t%s added to Kong", p.Plugin)
		return nil
	}
}

/*
 * CreateConsumer a Consumer on Kong
 */
func CreateConsumer(c Consumer) (Consumer, error) {
	resp, err := http.PostForm(consumersURL,
		url.Values{"username": {c.Username}, "custom_id": {c.Custom_id}})
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("func\tCreateConsumer:\tcould not create Kong customer:\t%v\n", err)
	}
	res := Consumer{}
	json.Unmarshal(body, &res)
	resp.Body.Close()
	log.Printf("func\tCreateConsumer:\tUser %s created as a Kong consumer:\t%s", c.Username, res.Custom_id)
	return res, nil
}

/*
 * Create an API Key for a Consumer on Kong
 */
func CreateAPIKey(c Consumer) (string, error) {
	resp, err := http.PostForm(consumersURL+c.Username+"/key-auth",
		url.Values{})
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("func\tCreateAPIKey:\tcould not create Kong customer:\t%v\n", err)
	}
	res := APIKey{}
	json.Unmarshal(body, &res)
	resp.Body.Close()
	log.Printf("func\tCreateAPIKey:\tCreated Kong API Key for User %s\twith Consumer Id:\t%s", c.Username, res.ConsumerId)
	return res.Key, nil
}

/*
 * EOF!
 */
