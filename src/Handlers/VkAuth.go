package Handlers

import (
	"golang-auth/src/Utils"
	"net/http"
	"net/url"
)

const vkRedirect string = "https://oauth.vk.com/authorize"
const vkSecrets string = "./.vk_secrets"

/**
Редиректит на https://oauth.vk.com/authorize
Начало oauth flow vk
*/
func VkAuthRedirect(w http.ResponseWriter, r *http.Request) {
	s, err := Utils.ReadSecrets(vkSecrets)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	u, err := url.Parse(vkRedirect)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	query := u.Query()
	query.Add("client_id", s.ClientID)
	query.Add("client_secret", s.ClientSecret)
	http.Redirect(w, r, u.String(), 301)
}

/**
Callback для VK OAuth нужен для получения code
Тут сходим за токеном и запишем в базу нового юзера
*/
func VkAuthCode(w http.ResponseWriter, r *http.Request) {
	// TODO
}
