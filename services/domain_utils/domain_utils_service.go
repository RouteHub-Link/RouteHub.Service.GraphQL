package services_domain_utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/RouteHub-Link/routehub-service-graphql/config"
)

type DomainUtilsService struct {
	host string
}

func NewDomainUtilsService() *DomainUtilsService {
	applicationConfig := config.ConfigurationService{}.Get()

	return &DomainUtilsService{host: applicationConfig.Services.DomainUtilsHost}
}

func (ds *DomainUtilsService) PostValidateURL(payload *URLValidationPayload) (id string, err error) {
	if payload.Link == "" {
		err = errors.New("link is required")
		return
	}

	url := strings.Join([]string{ds.host, ValidateURL}, "/")
	payloadAsjson, err := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadAsjson))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New("request failed")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	id = trimByteString(body, id)
	return
}

func (ds *DomainUtilsService) GetValidateURL(id string) (state TaskState, result *TaskResultPayload, err error) {
	if id == "" {
		err = errors.New("id is required")
		return
	}

	url := strings.Join([]string{ds.host, ValidateURL, id}, "/")
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New("request failed")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	taskInfo := new(TaskInfo)

	err = json.Unmarshal(body, &taskInfo)
	if err != nil {
		return
	}

	state = taskInfo.State
	emptyTime := time.Time{}
	if taskInfo.CompletedAt == emptyTime {
		return
	}

	byteToJson := json.NewDecoder(bytes.NewReader(taskInfo.Result))
	err = byteToJson.Decode(&result)

	return
}

func (ds *DomainUtilsService) PostValidateDNS(payload *DNSValidationPayload) (id string, err error) {
	if payload.Link == "" {
		err = errors.New("link is required")
		return
	}

	url := strings.Join([]string{ds.host, ValidateDNS}, "/")
	payloadAsjson, err := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadAsjson))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New("request failed")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	id = string(body)
	return
}

func (ds *DomainUtilsService) GetValidateDNS(id string) (state TaskState, result *TaskResultPayload, err error) {
	if id == "" {
		err = errors.New("id is required")
		return
	}

	url := strings.Join([]string{ds.host, ValidateDNS, id}, "/")
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New("request failed")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	taskInfo := new(TaskInfo)

	err = json.Unmarshal(body, &taskInfo)
	if err != nil {
		return
	}

	state = taskInfo.State
	emptyTime := time.Time{}
	if taskInfo.CompletedAt == emptyTime {
		return
	}

	byteToJson := json.NewDecoder(bytes.NewReader(taskInfo.Result))
	err = byteToJson.Decode(&result)

	return
}

func (ds *DomainUtilsService) PostValidateSite(payload *SiteValidationPayload) (id string, err error) {
	if payload.Link == "" {
		err = errors.New("link is required")
		return
	}

	url := strings.Join([]string{ds.host, ValidateSite}, "/")
	payloadAsjson, err := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadAsjson))
	req.Header.Set("Content-Type", "application/json")
	log.Printf("request : +%v", req)
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New("request failed")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	id = trimByteString(body, id)
	return
}

func trimByteString(body []byte, id string) string {
	id = string(body)
	id = strings.Replace(id, "\"", "", -1)
	id = strings.Replace(id, "\n", "", -1)
	return id
}

func (ds *DomainUtilsService) GetValidateSite(id string) (state TaskState, result *TaskResultPayload, taskInfo *TaskInfo, err error) {
	if id == "" {
		err = errors.New("id is required")
		return
	}

	url := strings.Join([]string{ds.host, ValidateSite, id}, "/")
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New("request failed")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	taskInfo = new(TaskInfo)

	err = json.Unmarshal(body, &taskInfo)
	if err != nil {
		return
	}

	state = taskInfo.State
	emptyTime := time.Time{}
	if taskInfo.CompletedAt == emptyTime {
		return
	}

	byteToJson := json.NewDecoder(bytes.NewReader(taskInfo.Result))
	err = byteToJson.Decode(&result)

	return
}
