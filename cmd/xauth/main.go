package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	authURL     = "https://twitter.com/i/oauth2/authorize"
	tokenURL    = "https://api.twitter.com/2/oauth2/token"
	redirectURI = "http://127.0.0.1:8080/callback" // Changed from localhost to 127.0.0.1
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

// generateCodeVerifier generates a random code verifier for PKCE
func generateCodeVerifier() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(bytes), nil
}

// generateCodeChallenge generates a code challenge from the verifier
func generateCodeChallenge(verifier string) string {
	hash := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(hash[:])
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: xauth <client_id> <client_secret>")
		fmt.Println("This will help you get OAuth 2.0 tokens for X API")
		fmt.Println()
		fmt.Println("IMPORTANT: Make sure your X app has:")
		fmt.Println("- OAuth 2.0 enabled")
		fmt.Println("- Callback URI set to: http://127.0.0.1:8080/callback")
		fmt.Println("- App permissions set to 'Read'")
		fmt.Println("- Type of App set to 'Web App, Automated App or Bot'")
		os.Exit(1)
	}

	clientID := os.Args[1]
	clientSecret := os.Args[2]

	fmt.Println("X OAuth 2.0 Token Helper")
	fmt.Println("=========================")
	fmt.Println()
	fmt.Printf("Client ID: %s\n", clientID)
	fmt.Printf("Redirect URI: %s\n", redirectURI)
	fmt.Println()

	// Generate PKCE parameters
	codeVerifier, err := generateCodeVerifier()
	if err != nil {
		log.Fatal("Failed to generate code verifier:", err)
	}
	codeChallenge := generateCodeChallenge(codeVerifier)

	// Step 1: Generate authorization URL with proper PKCE
	authParams := url.Values{}
	authParams.Set("response_type", "code")
	authParams.Set("client_id", clientID)
	authParams.Set("redirect_uri", redirectURI)
	authParams.Set("scope", "tweet.read users.read")
	authParams.Set("state", "state")
	authParams.Set("code_challenge", codeChallenge)
	authParams.Set("code_challenge_method", "S256")

	authURLWithParams := authURL + "?" + authParams.Encode()

	fmt.Println("Step 1: Open this URL in your browser to authorize the app:")
	fmt.Println(authURLWithParams)
	fmt.Println()

	// Step 2: Start local server to receive the callback
	fmt.Println("Step 2: Starting local server to receive authorization code...")
	fmt.Printf("Server will listen on: %s\n", redirectURI)
	fmt.Println("After authorizing, you'll be redirected to this URL")
	fmt.Println()

	var authCode string
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received callback request: %s %s\n", r.Method, r.URL.String())

		code := r.URL.Query().Get("code")
		if code != "" {
			authCode = code
			_, err := fmt.Fprintf(w, "Authorization successful! You can close this window.")
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Authorization code received successfully!")
			go func() {
				// Give the browser a moment to show the success message
				time.Sleep(2 * time.Second)
				os.Exit(0)
			}()
		} else {
			error := r.URL.Query().Get("error")
			errorDescription := r.URL.Query().Get("error_description")
			_, err := fmt.Fprintf(w, "Authorization failed! Error: %s - %s", error, errorDescription)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Authorization failed! Error: %s - %s\n", error, errorDescription)
		}
	})

	go func() {
		fmt.Printf("Starting server on %s...\n", redirectURI)
		if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
			log.Fatal("Server error:", err)
		}
	}()

	// Wait for the authorization code
	fmt.Println("Waiting for authorization...")
	for authCode == "" {
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Step 3: Exchanging authorization code for tokens...")

	// Step 3: Exchange authorization code for tokens with PKCE
	tokenParams := url.Values{}
	tokenParams.Set("grant_type", "authorization_code")
	tokenParams.Set("code", authCode)
	tokenParams.Set("redirect_uri", redirectURI)
	tokenParams.Set("client_id", clientID)
	tokenParams.Set("code_verifier", codeVerifier)

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(tokenParams.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+basicAuth(clientID, clientSecret))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Token exchange failed with status %d: %s\n", resp.StatusCode, string(body))
		os.Exit(1)
	}

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step 4: Tokens received successfully!")
	fmt.Println()
	fmt.Println("Add these to your config file (~/.config/tblogs/data.yml):")
	fmt.Println()
	fmt.Printf("app:\n")
	fmt.Printf("  x_cred:\n")
	fmt.Printf("    client_id: \"%s\"\n", clientID)
	fmt.Printf("    client_secret: \"%s\"\n", clientSecret)
	fmt.Printf("    access_token: \"%s\"\n", tokenResp.AccessToken)
	fmt.Printf("    refresh_token: \"%s\"\n", tokenResp.RefreshToken)
	fmt.Printf("    username: \"your_x_username\"\n")
	fmt.Println()
	fmt.Println("Note: Replace 'your_x_username' with your actual X username")
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
