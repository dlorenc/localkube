// Package androidpublisher provides access to the Google Play Developer API.
//
// See https://developers.google.com/android-publisher
//
// Usage example:
//
//   import "google.golang.org/api/androidpublisher/v1"
//   ...
//   androidpublisherService, err := androidpublisher.New(oauthHttpClient)
package androidpublisher // import "google.golang.org/api/androidpublisher/v1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "androidpublisher:v1"
const apiName = "androidpublisher"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/androidpublisher/v1/applications/"

// OAuth2 scopes used by this API.
const (
	// View and manage your Google Play Developer account
	AndroidpublisherScope = "https://www.googleapis.com/auth/androidpublisher"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Purchases = NewPurchasesService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Purchases *PurchasesService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewPurchasesService(s *Service) *PurchasesService {
	rs := &PurchasesService{s: s}
	return rs
}

type PurchasesService struct {
	s *Service
}

// SubscriptionPurchase: A SubscriptionPurchase resource indicates the
// status of a user's subscription purchase.
type SubscriptionPurchase struct {
	// AutoRenewing: Whether the subscription will automatically be renewed
	// when it reaches its current expiry time.
	AutoRenewing bool `json:"autoRenewing,omitempty"`

	// InitiationTimestampMsec: Time at which the subscription was granted,
	// in milliseconds since Epoch.
	InitiationTimestampMsec int64 `json:"initiationTimestampMsec,omitempty,string"`

	// Kind: This kind represents a subscriptionPurchase object in the
	// androidpublisher service.
	Kind string `json:"kind,omitempty"`

	// ValidUntilTimestampMsec: Time at which the subscription will expire,
	// in milliseconds since Epoch.
	ValidUntilTimestampMsec int64 `json:"validUntilTimestampMsec,omitempty,string"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AutoRenewing") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *SubscriptionPurchase) MarshalJSON() ([]byte, error) {
	type noMethod SubscriptionPurchase
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// method id "androidpublisher.purchases.cancel":

type PurchasesCancelCall struct {
	s              *Service
	packageName    string
	subscriptionId string
	token          string
	urlParams_     gensupport.URLParams
	ctx_           context.Context
}

// Cancel: Cancels a user's subscription purchase. The subscription
// remains valid until its expiration time.
func (r *PurchasesService) Cancel(packageName string, subscriptionId string, token string) *PurchasesCancelCall {
	c := &PurchasesCancelCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.packageName = packageName
	c.subscriptionId = subscriptionId
	c.token = token
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *PurchasesCancelCall) QuotaUser(quotaUser string) *PurchasesCancelCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *PurchasesCancelCall) UserIP(userIP string) *PurchasesCancelCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PurchasesCancelCall) Fields(s ...googleapi.Field) *PurchasesCancelCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PurchasesCancelCall) Context(ctx context.Context) *PurchasesCancelCall {
	c.ctx_ = ctx
	return c
}

func (c *PurchasesCancelCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{packageName}/subscriptions/{subscriptionId}/purchases/{token}/cancel")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"packageName":    c.packageName,
		"subscriptionId": c.subscriptionId,
		"token":          c.token,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "androidpublisher.purchases.cancel" call.
func (c *PurchasesCancelCall) Do() error {
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Cancels a user's subscription purchase. The subscription remains valid until its expiration time.",
	//   "httpMethod": "POST",
	//   "id": "androidpublisher.purchases.cancel",
	//   "parameterOrder": [
	//     "packageName",
	//     "subscriptionId",
	//     "token"
	//   ],
	//   "parameters": {
	//     "packageName": {
	//       "description": "The package name of the application for which this subscription was purchased (for example, 'com.some.thing').",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "The purchased subscription ID (for example, 'monthly001').",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "token": {
	//       "description": "The token provided to the user's device when the subscription was purchased.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{packageName}/subscriptions/{subscriptionId}/purchases/{token}/cancel",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "androidpublisher.purchases.get":

type PurchasesGetCall struct {
	s              *Service
	packageName    string
	subscriptionId string
	token          string
	urlParams_     gensupport.URLParams
	ifNoneMatch_   string
	ctx_           context.Context
}

// Get: Checks whether a user's subscription purchase is valid and
// returns its expiry time.
func (r *PurchasesService) Get(packageName string, subscriptionId string, token string) *PurchasesGetCall {
	c := &PurchasesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.packageName = packageName
	c.subscriptionId = subscriptionId
	c.token = token
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *PurchasesGetCall) QuotaUser(quotaUser string) *PurchasesGetCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *PurchasesGetCall) UserIP(userIP string) *PurchasesGetCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PurchasesGetCall) Fields(s ...googleapi.Field) *PurchasesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *PurchasesGetCall) IfNoneMatch(entityTag string) *PurchasesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PurchasesGetCall) Context(ctx context.Context) *PurchasesGetCall {
	c.ctx_ = ctx
	return c
}

func (c *PurchasesGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{packageName}/subscriptions/{subscriptionId}/purchases/{token}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"packageName":    c.packageName,
		"subscriptionId": c.subscriptionId,
		"token":          c.token,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "androidpublisher.purchases.get" call.
// Exactly one of *SubscriptionPurchase or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *SubscriptionPurchase.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *PurchasesGetCall) Do() (*SubscriptionPurchase, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SubscriptionPurchase{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Checks whether a user's subscription purchase is valid and returns its expiry time.",
	//   "httpMethod": "GET",
	//   "id": "androidpublisher.purchases.get",
	//   "parameterOrder": [
	//     "packageName",
	//     "subscriptionId",
	//     "token"
	//   ],
	//   "parameters": {
	//     "packageName": {
	//       "description": "The package name of the application for which this subscription was purchased (for example, 'com.some.thing').",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "The purchased subscription ID (for example, 'monthly001').",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "token": {
	//       "description": "The token provided to the user's device when the subscription was purchased.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{packageName}/subscriptions/{subscriptionId}/purchases/{token}",
	//   "response": {
	//     "$ref": "SubscriptionPurchase"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}
