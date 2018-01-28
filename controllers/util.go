package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
)

func GetBlock(blockId string) (map[string]interface{}, error) {
	url := "https://api.tierion.com/v1/records/" + blockId
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		beego.Debug(err)
	}
	req.Header.Set("X-Api-Key", "vZvkeRz/J+pd8mDuSPMZNVkl0TJ9qgfI3JphLV5S1uw=")
	req.Header.Set("X-Username", "leonie.reif@yahoo.com")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		beego.Error(err)
	}
	if val, ok := dat["json"]; ok {
		beego.Debug(val)
		json.Unmarshal([]byte(val.(string)), &dat)
		return dat, nil
	}
	return dat, nil
}

func PostBlock(c *RegisterController) (string, error) {
	url := "https://api.tierion.com/v1/records"
	jsonStr := []byte(fmt.Sprintf(`{    
		"datastoreId":"6543",
		"name":"%s",
		"medicalhistory": "%s",
		"birthplace": "%s",
		"locations": "%s",
		"health": "%s"
	}`, c.GetString("name"), c.GetString("medicalhistory"), c.GetString("birthplace"), c.GetString("locations"),
		c.GetString("generalhealth")))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		beego.Debug(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", "vZvkeRz/J+pd8mDuSPMZNVkl0TJ9qgfI3JphLV5S1uw=")
	req.Header.Set("X-Username", "leonie.reif@yahoo.com")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		beego.Error(err)
	}
	if val, ok := dat["id"]; ok {
		return val.(string), nil
	}
	return "", nil
}

func IdentifyFace(faceId string) (fi string, err error) {
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if recover() != nil {
			err = errors.New("array index out of bounds")
		}
	}()
	url := "https://westcentralus.api.cognitive.microsoft.com/face/v1.0/identify"
	jsonStr := []byte(`{    
		"personGroupId":"refugee",
		"faceIds":[
			"` + faceId + `"
		],
		"maxNumOfCandidatesReturned":1,
		"confidenceThreshold": 0.5
		
	}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Ocp-Apim-Subscription-Key", "5af1d1692f4c4148b4ddce8a3119bad4")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var arr []map[string]interface{}
	_ = json.Unmarshal([]byte(body), &arr)
	if val, ok := arr[0]["candidates"].([]interface{})[0].(map[string]interface{})["personId"]; ok {
		return val.(string), nil
	}
	return "", err
}

func GetPerson(personId string) (string, error) {
	url := "https://westcentralus.api.cognitive.microsoft.com/face/v1.0/persongroups/refugee/persons/" + personId
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Ocp-Apim-Subscription-Key", "5af1d1692f4c4148b4ddce8a3119bad4")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var dat map[string]interface{}
	beego.Debug(string(body))
	if err := json.Unmarshal(body, &dat); err != nil {
		beego.Error(err)
	}
	if val, ok := dat["userData"]; ok {
		return val.(string), nil
	}
	return "", nil
}

func DetectFace(image *[]byte) (string, error) {
	url := "https://westcentralus.api.cognitive.microsoft.com/face/v1.0/detect"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(*image))
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Ocp-Apim-Subscription-Key", "5af1d1692f4c4148b4ddce8a3119bad4")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	beego.Debug(string(body))
	var arr []map[string]interface{}
	_ = json.Unmarshal([]byte(body), &arr)
	if val, ok := arr[0]["faceId"]; ok {
		return val.(string), nil
	}
	return "", nil
}

func Train() {
	url := "https://westcentralus.api.cognitive.microsoft.com/face/v1.0/persongroups/refugee/train"
	req, err := http.NewRequest("POST", url, nil)
	req.Header.Set("Ocp-Apim-Subscription-Key", "5af1d1692f4c4148b4ddce8a3119bad4")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
}

func AddFace(personId string, image *[]byte) (string, error) {
	url := "https://westcentralus.api.cognitive.microsoft.com/face/v1.0/persongroups/refugee/persons/" + personId + "/persistedFaces"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(*image))
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Ocp-Apim-Subscription-Key", "5af1d1692f4c4148b4ddce8a3119bad4")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		beego.Error(err)
	}
	if val, ok := dat["persistedFaceId"]; ok {
		return val.(string), nil
	}
	return "", nil
}

func NewPerson(name string, userData string) string {
	url := "https://westcentralus.api.cognitive.microsoft.com/face/v1.0/persongroups/refugee/persons"
	jsonStr := []byte(`{"name":"` + name + `", "userData":"` + userData + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Ocp-Apim-Subscription-Key", "5af1d1692f4c4148b4ddce8a3119bad4")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		beego.Error(err)
	}
	return dat["personId"].(string)
}
