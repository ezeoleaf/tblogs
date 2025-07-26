package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/pkg/browser"
	"github.com/rivo/tview"
)

// XPost represents a post from the X API
type XPost struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	AuthorID  string `json:"author_id"`
	CreatedAt string `json:"created_at"`
}

// XUser represents a user from the X API
type XUser struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

// XTimelineResponse represents the response from the X timeline API
type XTimelineResponse struct {
	Data     []XPost `json:"data"`
	Includes struct {
		Users []XUser `json:"users"`
	} `json:"includes"`
	Meta struct {
		ResultCount int    `json:"result_count"`
		NextToken   string `json:"next_token"`
	} `json:"meta"`
}

// XUserResponse represents the response from the X user lookup API
type XUserResponse struct {
	Data XUser `json:"data"`
}

// XOAuthResponse represents the OAuth 2.0 token response
type XOAuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

// getUserID fetches the user ID from username using X API with OAuth 2.0
func (a *App) getUserID(accessToken, username string) (string, error) {
	url := fmt.Sprintf("https://api.twitter.com/2/users/by/username/%s", username)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var result XUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Data.ID, nil
}

// fetchTimeline fetches the user's home timeline from X API with OAuth 2.0
func (a *App) fetchTimeline(accessToken, userID string) (*XTimelineResponse, error) {
	url := fmt.Sprintf("https://api.twitter.com/2/users/%s/timelines/reverse_chronological", userID)

	// Add query parameters for better results
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("max_results", "20") // Limit to 20 posts
	q.Add("tweet.fields", "created_at,author_id")
	q.Add("user.fields", "name,username")
	q.Add("expansions", "author_id")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var result XTimelineResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// findUserByID finds a user in the includes section by ID
func (response *XTimelineResponse) findUserByID(userID string) (XUser, bool) {
	for _, user := range response.Includes.Users {
		if user.ID == userID {
			return user, true
		}
	}
	return XUser{}, false
}

func (a *App) generateX(listX *tview.List) {
	appCfg := a.Config.App
	listX.Clear()

	// Check if OAuth credentials are configured
	if appCfg.XCred.ClientID == "" || appCfg.XCred.ClientSecret == "" {
		listX.AddItem("X OAuth credentials not configured", "Add your X OAuth credentials to the config file", ' ', nil)
		listX.AddItem("Press Ctrl+R to reload", "", ' ', nil)
		return
	}

	// Check if username is configured
	if appCfg.XCred.Username == "" {
		listX.Clear()
		listX.AddItem("X username not configured", "Add your X username to the config file", ' ', nil)
		listX.AddItem("Press Ctrl+R to retry", "", ' ', nil)
		return
	}

	// Show loading message
	listX.AddItem("Loading X timeline...", "", ' ', nil)

	// Get access token
	var accessToken string
	if appCfg.XCred.AccessToken == "" {
		// No access token available
		listX.Clear()
		listX.AddItem("No access token available", "Run the OAuth flow to get a new access token", ' ', nil)
		listX.AddItem("Press Ctrl+R to retry", "", ' ', nil)
		return
	} else {
		accessToken = appCfg.XCred.AccessToken
	}

	username := appCfg.XCred.Username

	// Get user ID first
	userID, err := a.getUserID(accessToken, username)
	if err != nil {
		listX.Clear()
		listX.AddItem("Failed to get user ID", err.Error(), ' ', nil)
		listX.AddItem("Press Ctrl+R to retry", "", ' ', nil)
		return
	}

	// Fetch timeline
	timeline, err := a.fetchTimeline(accessToken, userID)
	if err != nil {
		listX.Clear()
		listX.AddItem("Failed to fetch timeline", err.Error(), ' ', nil)
		listX.AddItem("Press Ctrl+R to retry", "", ' ', nil)
		return
	}

	// Clear loading message and display posts
	listX.Clear()

	if len(timeline.Data) == 0 {
		listX.AddItem("No posts found", "Your timeline appears to be empty", ' ', nil)
		listX.AddItem("Press Ctrl+R to reload", "", ' ', nil)
		return
	}

	// Display posts
	for _, post := range timeline.Data {
		// Find the author information
		author, found := timeline.findUserByID(post.AuthorID)
		authorName := "Unknown"
		if found {
			authorName = author.Name
		}

		// Parse and format the creation time
		createdTime := "Unknown time"
		if post.CreatedAt != "" {
			if t, err := time.Parse(time.RFC3339, post.CreatedAt); err == nil {
				createdTime = t.Format("2006-01-02 15:04")
			}
		}

		// Truncate text if too long
		displayText := post.Text
		if len(displayText) > 100 {
			displayText = displayText[:97] + "..."
		}

		// Create click handler to open the post in browser
		postURL := fmt.Sprintf("https://twitter.com/i/status/%s", post.ID)
		listX.AddItem(displayText, fmt.Sprintf("by %s • %s", authorName, createdTime), 'x', func() {
			err := browser.OpenURL(postURL)
			if err != nil {
				log.Printf("failed to open URL: %v", err)
			}
		})
	}

	// Add reload option
	listX.AddItem("Press Ctrl+R to reload timeline", "", ' ', nil)

	// Set up input capture for reload functionality
	listX.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlR {
			a.generateX(listX)
			return nil
		}
		return event
	})
}

func (a *App) xPage(nextSlide func()) (title string, content tview.Primitive) {
	listX := a.viewsList["x"]
	if listX == nil {
		listX = getList()
		a.viewsList["x"] = listX
		a.generateX(listX)
	}

	return xSection, tview.NewFlex().
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(listX, 0, 1, true), 0, 1, true)
}
