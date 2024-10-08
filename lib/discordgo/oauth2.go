// Discordgo - Discord bindings for Go
// Available at https://github.com/bwmarrin/discordgo

// Copyright 2015-2016 Bruce Marriner <bruce@sqls.net>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains functions related to Discord OAuth2 endpoints

package discordgo

// ------------------------------------------------------------------------------------------------
// Code specific to Discord OAuth2 Applications
// ------------------------------------------------------------------------------------------------

// An Application struct stores values for a Discord OAuth2 Application
type Application struct {
	ID                  int64     `json:"id,string,omitempty"`
	Name                string    `json:"name"`
	Description         string    `json:"description,omitempty"`
	Icon                string    `json:"icon,omitempty"`
	Secret              string    `json:"secret,omitempty"`
	RedirectURIs        *[]string `json:"redirect_uris,omitempty"`
	BotRequireCodeGrant bool      `json:"bot_require_code_grant,omitempty"`
	BotPublic           bool      `json:"bot_public,omitempty"`
	RPCApplicationState int       `json:"rpc_application_state,omitempty"`
	Flags               int       `json:"flags,omitempty"`
	Owner               *User     `json:"owner"`
	Bot                 *User     `json:"bot"`
}

// Application returns an Application structure of a specific Application
//   appID : The ID of an Application
func (s *Session) Application(appID int64) (st *Application, err error) {

	body, err := s.RequestWithBucketID("GET", EndpointApplication(appID), nil, nil, EndpointApplication(0))
	if err != nil {
		return
	}

	err = unmarshal(body, &st)
	return
}

// Application returns an Application structure of the current bot
func (s *Session) ApplicationMe() (st *Application, err error) {

	body, err := s.RequestWithBucketID("GET", EndpointApplicationMe, nil, nil, EndpointApplicationMe)
	if err != nil {
		return
	}

	err = unmarshal(body, &st)
	return
}

// Applications returns all applications for the authenticated user
func (s *Session) Applications() (st []*Application, err error) {

	body, err := s.RequestWithBucketID("GET", EndpointApplications, nil, nil, EndpointApplications)
	if err != nil {
		return
	}

	err = unmarshal(body, &st)
	return
}

// ApplicationCreate creates a new Application
//    name : Name of Application / Bot
//    uris : Redirect URIs (Not required)
func (s *Session) ApplicationCreate(ap *Application) (st *Application, err error) {

	data := struct {
		Name         string    `json:"name"`
		Description  string    `json:"description"`
		RedirectURIs *[]string `json:"redirect_uris,omitempty"`
	}{ap.Name, ap.Description, ap.RedirectURIs}

	body, err := s.RequestWithBucketID("POST", EndpointApplications, data, nil, EndpointApplications)
	if err != nil {
		return
	}

	err = unmarshal(body, &st)
	return
}

// ApplicationUpdate updates an existing Application
//   var : desc
func (s *Session) ApplicationUpdate(appID int64, ap *Application) (st *Application, err error) {

	data := struct {
		Name         string    `json:"name"`
		Description  string    `json:"description"`
		RedirectURIs *[]string `json:"redirect_uris,omitempty"`
	}{ap.Name, ap.Description, ap.RedirectURIs}

	body, err := s.RequestWithBucketID("PUT", EndpointApplication(appID), data, nil, EndpointApplication(0))
	if err != nil {
		return
	}

	err = unmarshal(body, &st)
	return
}

// ApplicationDelete deletes an existing Application
//   appID : The ID of an Application
func (s *Session) ApplicationDelete(appID int64) (err error) {

	_, err = s.RequestWithBucketID("DELETE", EndpointApplication(appID), nil, nil, EndpointApplication(0))
	if err != nil {
		return
	}

	return
}

// ------------------------------------------------------------------------------------------------
// Code specific to Discord OAuth2 Application Bots
// ------------------------------------------------------------------------------------------------

// ApplicationBotCreate creates an Application Bot Account
//
//   appID : The ID of an Application
//
// NOTE: func name may change, if I can think up something better.
func (s *Session) ApplicationBotCreate(appID int64) (st *User, err error) {

	body, err := s.RequestWithBucketID("POST", EndpointApplicationsBot(appID), nil, nil, EndpointApplicationsBot(0))
	if err != nil {
		return
	}

	err = unmarshal(body, &st)
	return
}
