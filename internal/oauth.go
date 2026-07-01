package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/zoho/zoho-payments-go-sdk/auth"
	"github.com/zoho/zoho-payments-go-sdk/edition"
	"github.com/zoho/zoho-payments-go-sdk/exception"
)

func GenerateAccessToken(refreshToken, clientID, clientSecret, redirectURI string, ed edition.Edition) (*auth.OAuthToken, error) {
	if !ed.Valid() {
		return nil, exception.NewValidationError("edition", "a valid edition is required")
	}
	if refreshToken == "" {
		return nil, exception.NewValidationError("refreshToken", "is required")
	}
	if clientID == "" {
		return nil, exception.NewValidationError("clientId", "is required")
	}
	if clientSecret == "" {
		return nil, exception.NewValidationError("clientSecret", "is required")
	}
	if redirectURI == "" {
		return nil, exception.NewValidationError("redirectUri", "is required")
	}

	form := url.Values{}
	form.Set("refresh_token", refreshToken)
	form.Set("client_id", clientID)
	form.Set("client_secret", clientSecret)
	form.Set("redirect_uri", redirectURI)
	form.Set("grant_type", "refresh_token")

	endpoint := ed.AccountsURL() + "/oauth/v2/token"
	httpClient := &http.Client{Timeout: DefaultRequestTimeout}

	req, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, exception.NewConnectionError("failed to build token request", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", UserAgent())
	req.Header.Set("Accept", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, exception.NewConnectionError("token request failed", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, exception.NewConnectionError("failed to read token response", err)
	}

	var parsed map[string]any
	if len(strings.TrimSpace(string(data))) > 0 {
		if err := json.Unmarshal(data, &parsed); err != nil {
			return nil, exception.NewConnectionError("failed to parse token response", err)
		}
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		errCode := stringField(parsed, "error")
		if errCode == "" {
			errCode = "token_refresh_failed"
		}
		return nil, &exception.AuthenticationError{APIError: exception.NewAPIError(resp.StatusCode, errCode, fmt.Sprintf("token refresh failed with HTTP %d", resp.StatusCode))}
	}

	if errCode := stringField(parsed, "error"); errCode != "" {
		return nil, &exception.AuthenticationError{APIError: exception.NewAPIError(resp.StatusCode, errCode, "OAuth token refresh failed")}
	}

	accessToken := stringField(parsed, "access_token")
	if accessToken == "" {
		return nil, &exception.AuthenticationError{APIError: exception.NewAPIError(resp.StatusCode, "invalid_token_response", "access_token missing from token response")}
	}

	expiresIn := int64(3600)
	if value, ok := parsed["expires_in_sec"]; ok {
		expiresIn = toInt64(value)
	} else if value, ok := parsed["expires_in"]; ok {
		expiresIn = toInt64(value)
	}

	return &auth.OAuthToken{AccessToken: accessToken, ExpiresIn: expiresIn}, nil
}
