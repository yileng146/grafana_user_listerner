package grafana_org_invite

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func PuPuTech_Invite(ad_num string, grafana_url string, graf_admintoken string, org_id string) {
	username := ad_num
	url := grafana_url + "/api/orgs/" + org_id + "/users"
	method := "POST"
	str := "{\"loginOrEmail\":" + "\"" + username + "\"" + "," + "\"role\":\"viewer\"}"
	payload := strings.NewReader(str)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", graf_admintoken)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
