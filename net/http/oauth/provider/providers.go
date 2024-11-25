/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package provider

import (
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/amazon"
	"github.com/markbates/goth/providers/auth0"
	"github.com/markbates/goth/providers/dropbox"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/gitlab"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/microsoftonline"
	"github.com/markbates/goth/providers/onedrive"
	"github.com/markbates/goth/providers/openidConnect"
	"github.com/markbates/goth/providers/twitter"
)

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

func NewAuth() *ProviderIndex {
	goth.UseProviders(
		twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost/api/auth/twitter/callback"),
		// If you'd like to use authenticate instead of authorize in Twitter provider, use this instead.
		// twitter.NewAuthenticate(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost/api/auth/twitter/callback"),
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "http://localhost/api/auth/google/callback"),
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:8070/auth/github/callback"),
		dropbox.New(os.Getenv("DROPBOX_KEY"), os.Getenv("DROPBOX_SECRET"), "http://localhost/api/auth/dropbox/callback"),
		amazon.New(os.Getenv("AMAZON_KEY"), os.Getenv("AMAZON_SECRET"), "http://localhost/api/auth/amazon/callback"),
		onedrive.New(os.Getenv("ONEDRIVE_KEY"), os.Getenv("ONEDRIVE_SECRET"), "http://localhost/api/auth/onedrive/callback"),
		microsoftonline.New(os.Getenv("MICROSOFTONLINE_KEY"), os.Getenv("MICROSOFTONLINE_SECRET"), "http://localhost/api/auth/microsoftonline/callback"),
		gitlab.New(os.Getenv("GITLAB_KEY"), os.Getenv("GITLAB_SECRET"), "http://localhost/api/auth/gitlab/callback"),

		//Auth0 allocates domain per customer, a domain must be provided for auth0 to work
		auth0.New(os.Getenv("AUTH0_KEY"), os.Getenv("AUTH0_SECRET"), "http://localhost/api/auth/auth0/callback", os.Getenv("AUTH0_DOMAIN")),
		New(os.Getenv("HOPER_KEY"), os.Getenv("HOPER_SECRET"), "http://localhost/auth/hoper/callback"),
	)

	// OpenID Connect is based on OpenID Connect Auto Discovery URL (https://openid.net/specs/openid-connect-discovery-1_0-17.html)
	// because the OpenID Connect provider initialize it self in the New(), it can return an error which should be handled or ignored
	// ignore the error for now
	openidConnect, _ := openidConnect.New(os.Getenv("OPENID_CONNECT_KEY"), os.Getenv("OPENID_CONNECT_SECRET"), "http://localhost/api/auth/openid-connect/callback", os.Getenv("OPENID_CONNECT_DISCOVERY_URL"))
	if openidConnect != nil {
		goth.UseProviders(openidConnect)
	}

	m := make(map[string]string)
	m["amazon"] = "Amazon"
	m["bitbucket"] = "Bitbucket"
	m["dropbox"] = "Dropbox"
	m["facebook"] = "Facebook"
	m["github"] = "Github"
	m["gitlab"] = "Gitlab"
	m["google"] = "Google"
	m["gplus"] = "Google Plus"
	m["steam"] = "Steam"
	m["twitch"] = "Twitch"
	m["uber"] = "Uber"
	m["wepay"] = "Wepay"
	m["instagram"] = "Instagram"
	m["onedrive"] = "Onedrive"
	m["azuread"] = "Azure AD"
	m["microsoftonline"] = "Microsoft Online"
	m["paypal"] = "Paypal"
	m["twitter"] = "Twitter"
	m["auth0"] = "Auth0"
	m["hoper"] = "hoper"

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	return &ProviderIndex{Providers: keys, ProvidersMap: m}

}
