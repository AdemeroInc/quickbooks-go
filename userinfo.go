// Copyright (c) 2018, Randy Westlund. All rights reserved.
// This code is under the BSD-2-Clause license.

package quickbooks

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// UserInfo describes a user.
type UserInfo struct {
	Sub                 string
	Email               string
	EmailVerified       bool
	GivenName           string
	FamilyName          string
	PhoneNumber         string
	PhoneNumberVerified bool
	Address             UserInfoAddress
}

// FetchUserInfo returns the QuickBooks UserInfo object.
func (c *Client) FetchUserInfo() (*UserInfo, error) {
	var u, err = url.Parse(string(c.discoveryAPI.UserinfoEndpoint))
	if err != nil {
		return nil, err
	}
	var v = url.Values{}
	v.Add("minorversion", minorVersion)
	u.RawQuery = v.Encode()
	var req *http.Request
	req, err = http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	var res *http.Response
	res, err = c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, parseFailure(res)
	}

	var userInfo UserInfo
	err = json.NewDecoder(res.Body).Decode(&userInfo)
	return &userInfo, err
}
