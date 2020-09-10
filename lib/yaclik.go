package yaclik

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func FetchFieldsJson(appid string, subdomain string, userid string, password string) ([]byte, error) {
	client := &http.Client{}
	type Body struct {
		App string `json:"app"`
	}
	data := Body{appid}
	body, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Error JSON Encode")
	}
	req, err := http.NewRequest("GET", "https://"+subdomain+".cybozu.com/k/v1/app/form/fields.json", bytes.NewReader(body))
	if err != nil {
		log.Fatal("Error request")
	}
	authData := userid + ":" + password
	req.Header.Add("X-Cybozu-Authorization", b64.StdEncoding.EncodeToString([]byte(authData)))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error client")
	}
	defer resp.Body.Close()
	ba, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error read")
	}
	return ba, err
}

func ParseFieldsJSON(ba *[]byte, iop io.Writer) {
	csvHeader := []string{"field_code", "label", "type"} // CSVヘッダー
	w := csv.NewWriter(iop)
	if err := w.Write(csvHeader); err != nil { // CSVヘッダー書き込み
		log.Fatal(err)
	}
	var f interface{}
	if err := json.Unmarshal(*ba, &f); err != nil { // ※baには外部からJSONデータが入ってくる想定
		log.Fatal(err)
	}
	m := f.(map[string]interface{})
	prop := m["properties"].(map[string]interface{})
	for k, v := range prop {
		record := []string{} // CSV行追加用スライス
		m := v.(map[string]interface{})
		record = append(record, k, m["label"].(string), m["type"].(string)) // データ追加
		if err := w.Write(record); err != nil {                             // CSV書き込み
			log.Fatal(err)
		}
		if m["type"].(string) == "SUBTABLE" {
			for k, v := range m {
				if k == "fields" {
					sv := v.(map[string]interface{})
					for k, v := range sv {
						value := v.(map[string]interface{})
						record := []string{}                                                                     // CSV行追加用スライス
						record = append(record, "~SUBTABLE~"+k, value["label"].(string), value["type"].(string)) // データ追加
						if err := w.Write(record); err != nil {                                                  // CSV書き込み
							log.Fatal(err)
						}
					}
				}
			}
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	return
}
