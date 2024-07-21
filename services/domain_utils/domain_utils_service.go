package services_domain_utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/RouteHub-Link/routehub-service-graphql/config"
	"github.com/go-resty/resty/v2"
)

type DomainUtilsService struct {
	host  string
	resty *resty.Client
}

func NewDomainUtilsService() *DomainUtilsService {
	applicationConfig := config.ConfigurationService{}.Get()
	restyClient := resty.New()

	return &DomainUtilsService{host: applicationConfig.Services.DomainUtilsHost, resty: restyClient}
}

func (ds *DomainUtilsService) PostValidateURL(payload *URLValidationPayload) (id string, err error) {
	if payload.Link == "" {
		return "", errors.New("link is required")
	}

	url := strings.Join([]string{ds.host, ValidateURL}, "/")
	payloadAsjson, _ := json.Marshal(payload)

	resp, err := ds.resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payloadAsjson).
		Post(url)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode() != http.StatusOK {
		return "", errors.New("request failed")
	}

	body := resp.Body()
	id = trimByteString(body, id)
	return id, nil
}

func (ds *DomainUtilsService) GetValidateURL(id string) (state TaskState, result *TaskResultPayload, err error) {
	if id == "" {
		err = errors.New("id is required")
		return
	}

	url := strings.Join([]string{ds.host, ValidateURL, id}, "/")
	resp, err := ds.resty.R().Get(url)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode() != http.StatusOK {
		err = errors.New("request failed")
		return
	}

	taskInfo := new(TaskInfo)

	err = json.Unmarshal(resp.Body(), &taskInfo)
	if err != nil {
		return
	}

	state = taskInfo.State
	if taskInfo.CompletedAt.IsZero() {
		return
	}

	err = json.NewDecoder(bytes.NewReader(taskInfo.Result)).Decode(&result)

	return
}

func (ds *DomainUtilsService) PostValidateDNS(payload *DNSVerificationPayload) (id string, err error) {
	if payload.Link == "" {
		return "", errors.New("link is required")
	}

	url := strings.Join([]string{ds.host, ValidateDNS}, "/")
	payloadAsjson, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	resp, err := ds.resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payloadAsjson).
		Post(url)
	if err != nil {
		return "", err
	}

	if resp.StatusCode() != http.StatusOK {
		return "", errors.New("request failed")
	}

	body := resp.Body()
	id = trimByteString(body, id)
	return id, nil
}

func (ds *DomainUtilsService) GetValidateDNS(id string) (state TaskState, result *TaskResultPayload, taskInfo *TaskInfo, err error) {
	if id == "" {
		err = errors.New("id is required")
		return
	}

	url := strings.Join([]string{ds.host, ValidateDNS, id}, "/")

	resp, err := ds.resty.R().
		SetHeader("Content-Type", "application/json").
		Get(url)
	if err != nil {
		return
	}

	if resp.StatusCode() != http.StatusOK {
		err = errors.New("request failed")
		return
	}

	taskInfo = new(TaskInfo)

	err = json.Unmarshal(resp.Body(), &taskInfo)
	if err != nil {
		return
	}

	state = taskInfo.State
	emptyTime := time.Time{}
	if taskInfo.CompletedAt == emptyTime {
		return
	}

	err = json.NewDecoder(bytes.NewReader(taskInfo.Result)).Decode(&result)

	return
}
func (ds *DomainUtilsService) PostValidateSite(payload *SiteValidationPayload) (id string, err error) {
	if payload.Link == "" {
		return "", errors.New("link is required")
	}

	url := strings.Join([]string{ds.host, ValidateSite}, "/")
	payloadAsjson, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	resp, err := ds.resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payloadAsjson).
		Post(url)
	if err != nil {
		return "", err
	}

	if resp.StatusCode() != http.StatusOK {
		return "", errors.New("request failed")
	}

	body := resp.Body()
	id = trimByteString(body, id)
	return id, nil
}

func (ds *DomainUtilsService) GetValidateSite(id string) (state TaskState, result *TaskResultPayload, taskInfo *TaskInfo, err error) {
	if id == "" {
		err = errors.New("id is required")
		return
	}

	url := strings.Join([]string{ds.host, ValidateSite, id}, "/")

	resp, err := ds.resty.R().
		SetHeader("Content-Type", "application/json").
		Get(url)
	if err != nil {
		return
	}

	if resp.StatusCode() != http.StatusOK {
		err = errors.New("request failed")
		return
	}

	taskInfo = new(TaskInfo)

	err = json.Unmarshal(resp.Body(), &taskInfo)
	if err != nil {
		return
	}

	state = taskInfo.State
	if taskInfo.CompletedAt.IsZero() {
		return
	}

	err = json.NewDecoder(bytes.NewReader(taskInfo.Result)).Decode(&result)

	return
}

func trimByteString(body []byte, id string) string {
	id = string(bytes.TrimSpace(body))
	return id
}
