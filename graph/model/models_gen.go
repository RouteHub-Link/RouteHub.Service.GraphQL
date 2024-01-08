// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"github.com/google/uuid"
)

type AccountPhoneInput struct {
	Number      string `json:"number"`
	CountryCode string `json:"countryCode"`
}

type AnalyticReport struct {
	Link         *database_models.Link   `json:"link"`
	Domain       *database_models.Domain `json:"domain"`
	TotalHits    int                     `json:"totalHits"`
	TotalSuccess int                     `json:"totalSuccess"`
	TotalFailed  int                     `json:"totalFailed"`
	Referers     []*MetricAnalytics      `json:"referers"`
	Locations    []*MetricAnalytics      `json:"locations"`
	Useragents   []*MetricAnalytics      `json:"useragents"`
}

type AnalyticReports struct {
	TodayObservations     []*ObservationAnalytic `json:"todayObservations"`
	YesterdayObservations []*ObservationAnalytic `json:"yesterdayObservations"`
	LastWeekObservations  string                 `json:"lastWeekObservations"`
	LastMonthObservations string                 `json:"lastMonthObservations"`
	LastYearObservations  string                 `json:"lastYearObservations"`
	AllTimeObservations   string                 `json:"allTimeObservations"`
}

type BooleanFilter struct {
	Equal    *bool `json:"equal,omitempty"`
	NotEqual *bool `json:"notEqual,omitempty"`
	IsNull   *bool `json:"isNull,omitempty"`
}

type ClientInformationInput struct {
	Useragent string `json:"useragent"`
	IP        string `json:"ip"`
}

type CrawlRequestInput struct {
	LinkID uuid.UUID `json:"linkId"`
}

type DomainCreateInput struct {
	OrganizationID    uuid.UUID               `json:"organizationId"`
	Name              string                  `json:"name"`
	URL               string                  `json:"url"`
	ClientInformation *ClientInformationInput `json:"clientInformation"`
}

type DomainVerification struct {
	ID        uuid.UUID                `json:"id"`
	Domain    *database_models.Domain  `json:"domain"`
	Status    database_enums.DNSStatus `json:"status"`
	Logs      []*database_types.Log    `json:"logs"`
	CreatedAt time.Time                `json:"createdAt"`
	UpdatedAt *time.Time               `json:"updatedAt,omitempty"`
	DeletedAt *time.Time               `json:"deletedAt,omitempty"`
}

type FloatFilter struct {
	Equal    *float64  `json:"equal,omitempty"`
	NotEqual *float64  `json:"notEqual,omitempty"`
	In       []float64 `json:"in,omitempty"`
	NotIn    []float64 `json:"notIn,omitempty"`
	Gt       *float64  `json:"gt,omitempty"`
	Gte      *float64  `json:"gte,omitempty"`
	Lt       *float64  `json:"lt,omitempty"`
	Lte      *float64  `json:"lte,omitempty"`
	IsNull   *bool     `json:"isNull,omitempty"`
}

type IDFilter struct {
	Equal        *string  `json:"equal,omitempty"`
	EqualFold    *string  `json:"equalFold,omitempty"`
	NotEqual     *string  `json:"notEqual,omitempty"`
	In           []string `json:"in,omitempty"`
	NotIn        []string `json:"notIn,omitempty"`
	Contains     *string  `json:"contains,omitempty"`
	ContainsFold *string  `json:"containsFold,omitempty"`
	Gt           *string  `json:"gt,omitempty"`
	Gte          *string  `json:"gte,omitempty"`
	Lt           *string  `json:"lt,omitempty"`
	Lte          *string  `json:"lte,omitempty"`
	HasPrefix    *string  `json:"hasPrefix,omitempty"`
	HasSuffix    *string  `json:"hasSuffix,omitempty"`
	IsNull       *bool    `json:"isNull,omitempty"`
}

type IntFilter struct {
	Equal    *int  `json:"equal,omitempty"`
	NotEqual *int  `json:"notEqual,omitempty"`
	In       []int `json:"in,omitempty"`
	NotIn    []int `json:"notIn,omitempty"`
	Gt       *int  `json:"gt,omitempty"`
	Gte      *int  `json:"gte,omitempty"`
	Lt       *int  `json:"lt,omitempty"`
	Lte      *int  `json:"lte,omitempty"`
	IsNull   *bool `json:"isNull,omitempty"`
}

type LinkCreateInput struct {
	Target             string                             `json:"target"`
	Path               *string                            `json:"path,omitempty"`
	PlatformID         uuid.UUID                          `json:"platformId"`
	RedirectionOptions *database_enums.RedirectionOptions `json:"redirectionOptions,omitempty"`
	OpenGraph          *database_types.OpenGraph          `json:"openGraph"`
}

type LinkFilter struct {
	Platform *UUIDFilter `json:"platform,omitempty"`
	Domain   *UUIDFilter `json:"domain,omitempty"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Token string `json:"token"`
}

type MapEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MapFilter struct {
	Equal        *MapEntry   `json:"equal,omitempty"`
	EqualFold    *MapEntry   `json:"equalFold,omitempty"`
	NotEqual     *MapEntry   `json:"notEqual,omitempty"`
	In           *MapInEntry `json:"in,omitempty"`
	NotIn        *MapInEntry `json:"notIn,omitempty"`
	Contains     *MapEntry   `json:"contains,omitempty"`
	ContainsFold *MapEntry   `json:"containsFold,omitempty"`
	Gt           *MapEntry   `json:"gt,omitempty"`
	Gte          *MapEntry   `json:"gte,omitempty"`
	Lt           *MapEntry   `json:"lt,omitempty"`
	Lte          *MapEntry   `json:"lte,omitempty"`
	HasPrefix    *MapEntry   `json:"hasPrefix,omitempty"`
	HasSuffix    *MapEntry   `json:"hasSuffix,omitempty"`
	IsNull       *bool       `json:"isNull,omitempty"`
}

type MapInEntry struct {
	Key   string   `json:"key"`
	Value []string `json:"value,omitempty"`
}

type MetricAnalytics struct {
	Feeder       string    `json:"feeder"`
	TotalHits    int       `json:"totalHits"`
	TotalSuccess int       `json:"totalSuccess"`
	TotalFailed  int       `json:"totalFailed"`
	StartAt      time.Time `json:"startAt"`
	EndAt        time.Time `json:"endAt"`
}

type ObservationAnalytic struct {
	ID                uuid.UUID                         `json:"id"`
	Link              *database_models.Link             `json:"link"`
	Domain            *database_models.Domain           `json:"domain"`
	Platform          *database_models.Platform         `json:"platform"`
	Useragent         string                            `json:"useragent"`
	IP                string                            `json:"ip"`
	Referrer          string                            `json:"referrer"`
	Location          string                            `json:"location"`
	RedirectionChoice database_enums.RedirectionOptions `json:"redirectionChoice"`
	Success           bool                              `json:"success"`
	CreatedAt         time.Time                         `json:"createdAt"`
	UpdatedAt         *time.Time                        `json:"updatedAt,omitempty"`
}

type ObservationInput struct {
	LinkID            uuid.UUID                         `json:"linkId"`
	DomainID          uuid.UUID                         `json:"domainId"`
	PlatformID        uuid.UUID                         `json:"platformId"`
	Useragent         string                            `json:"useragent"`
	IP                string                            `json:"ip"`
	Referrer          string                            `json:"referrer"`
	Location          string                            `json:"location"`
	RedirectionChoice database_enums.RedirectionOptions `json:"redirectionChoice"`
	Success           bool                              `json:"success"`
}

type OrganizationInput struct {
	Name         string              `json:"name"`
	Website      string              `json:"website"`
	Description  string              `json:"description"`
	Location     string              `json:"location"`
	SocialMedias []*SocialMediaInput `json:"socialMedias"`
}

type PasswordReset struct {
	ID        uuid.UUID             `json:"id"`
	User      *database_models.User `json:"user"`
	Token     string                `json:"token"`
	CreatedAt time.Time             `json:"createdAt"`
	ExpiresAt time.Time             `json:"expiresAt"`
	UpdatedAt *time.Time            `json:"updatedAt,omitempty"`
	DeletedAt *time.Time            `json:"deletedAt,omitempty"`
}

type PasswordResetCreateInput struct {
	Email             string                  `json:"email"`
	ClientInformation *ClientInformationInput `json:"clientInformation"`
}

type PasswordResetUpdateInput struct {
	Token             string                  `json:"token"`
	Password          string                  `json:"password"`
	ConfirmPassword   string                  `json:"confirmPassword"`
	ClientInformation *ClientInformationInput `json:"clientInformation"`
}

type Permission struct {
	ID            uuid.UUID                       `json:"id"`
	Name          string                          `json:"name"`
	Description   string                          `json:"description"`
	Organizations []*database_models.Organization `json:"organizations"`
	Domains       []*database_models.Domain       `json:"domains"`
	Platforms     []*database_models.Platform     `json:"platforms"`
}

type PlatformDeployment struct {
	ID        uuid.UUID                 `json:"id"`
	Platform  *database_models.Platform `json:"platform"`
	Domain    *database_models.Domain   `json:"domain"`
	Status    DeploymentStatus          `json:"status"`
	Logs      []*database_types.Log     `json:"logs"`
	CreatedAt time.Time                 `json:"createdAt"`
	UpdatedAt *time.Time                `json:"updatedAt,omitempty"`
	DeletedAt *time.Time                `json:"deletedAt,omitempty"`
}

type SocialMediaInput struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Icon string `json:"icon"`
}

type StringFilter struct {
	Equal        *string  `json:"equal,omitempty"`
	EqualFold    *string  `json:"equalFold,omitempty"`
	NotEqual     *string  `json:"notEqual,omitempty"`
	In           []string `json:"in,omitempty"`
	NotIn        []string `json:"notIn,omitempty"`
	Contains     *string  `json:"contains,omitempty"`
	ContainsFold *string  `json:"containsFold,omitempty"`
	Gt           *string  `json:"gt,omitempty"`
	Gte          *string  `json:"gte,omitempty"`
	Lt           *string  `json:"lt,omitempty"`
	Lte          *string  `json:"lte,omitempty"`
	HasPrefix    *string  `json:"hasPrefix,omitempty"`
	HasSuffix    *string  `json:"hasSuffix,omitempty"`
	IsNull       *bool    `json:"isNull,omitempty"`
}

type Template struct {
	ID                uuid.UUID                         `json:"id"`
	Name              string                            `json:"name"`
	Platform          *database_models.Platform         `json:"platform"`
	OpenGraph         *database_types.OpenGraph         `json:"openGraph"`
	RedirectionChoice database_enums.RedirectionOptions `json:"redirectionChoice"`
	State             database_enums.StatusState        `json:"state"`
	CreatedBy         *database_models.User             `json:"createdBy"`
	EditedBy          *database_models.User             `json:"editedBy"`
	CreatedAt         time.Time                         `json:"createdAt"`
	UpdatedAt         *time.Time                        `json:"updatedAt,omitempty"`
	DeletedAt         *time.Time                        `json:"deletedAt,omitempty"`
}

type TemplateInput struct {
	Name              string                            `json:"name"`
	OpenGraph         *database_types.OpenGraph         `json:"openGraph"`
	RedirectionChoice database_enums.RedirectionOptions `json:"redirectionChoice"`
	State             database_enums.StatusState        `json:"state"`
}

type TimeFilter struct {
	Equal    *time.Time   `json:"equal,omitempty"`
	NotEqual *time.Time   `json:"notEqual,omitempty"`
	In       []*time.Time `json:"in,omitempty"`
	NotIn    []*time.Time `json:"notIn,omitempty"`
	Gt       *time.Time   `json:"gt,omitempty"`
	Gte      *time.Time   `json:"gte,omitempty"`
	Lt       *time.Time   `json:"lt,omitempty"`
	Lte      *time.Time   `json:"lte,omitempty"`
	IsNull   *bool        `json:"isNull,omitempty"`
}

type UUIDFilter struct {
	Equal    *uuid.UUID  `json:"equal,omitempty"`
	NotEqual *uuid.UUID  `json:"notEqual,omitempty"`
	In       []uuid.UUID `json:"in,omitempty"`
	NotIn    []uuid.UUID `json:"notIn,omitempty"`
	Gt       *uuid.UUID  `json:"gt,omitempty"`
	Gte      *uuid.UUID  `json:"gte,omitempty"`
	Lt       *uuid.UUID  `json:"lt,omitempty"`
	Lte      *uuid.UUID  `json:"lte,omitempty"`
	IsNull   *bool       `json:"isNull,omitempty"`
}

type UpdateUserInviteInput struct {
	Code   string                          `json:"code"`
	Status database_enums.InvitationStatus `json:"status"`
	User   *UserInput                      `json:"user"`
}

type UserInput struct {
	Email             string                  `json:"email"`
	Password          string                  `json:"password"`
	ConfirmPassword   string                  `json:"confirmPassword"`
	Fullname          string                  `json:"fullname"`
	Phone             *AccountPhoneInput      `json:"phone"`
	ClientInformation *ClientInformationInput `json:"clientInformation"`
}

type UserUpdatePasswordInput struct {
	Password          string                  `json:"password"`
	ConfirmPassword   string                  `json:"confirmPassword"`
	ClientInformation *ClientInformationInput `json:"clientInformation"`
}

type DeploymentStatus string

const (
	DeploymentStatusWorking DeploymentStatus = "WORKING"
	DeploymentStatusFailed  DeploymentStatus = "FAILED"
	DeploymentStatusSuccess DeploymentStatus = "SUCCESS"
)

var AllDeploymentStatus = []DeploymentStatus{
	DeploymentStatusWorking,
	DeploymentStatusFailed,
	DeploymentStatusSuccess,
}

func (e DeploymentStatus) IsValid() bool {
	switch e {
	case DeploymentStatusWorking, DeploymentStatusFailed, DeploymentStatusSuccess:
		return true
	}
	return false
}

func (e DeploymentStatus) String() string {
	return string(e)
}

func (e *DeploymentStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DeploymentStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DeploymentStatus", str)
	}
	return nil
}

func (e DeploymentStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
