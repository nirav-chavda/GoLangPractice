package twitter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/oauth2"
)

// Client s
type Client struct {
	CosumerKey     string
	CosumerSercret string
	// unexported fields
	once   sync.Once
	client *http.Client
}

func (c *Client) clientSingleton() *http.Client {
	c.once.Do(func() {
		c.client, _ = getClient(c.CosumerKey, c.CosumerSercret)
	})
	return c.client
}

func getClient(key, secret string) (*http.Client, error) {
	// Authentication [https://developer.twitter.com/en/docs/authentication/api-reference/token]
	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token", strings.NewReader("grant_type=client_credentials"))
	if err != nil {
		return nil, fmt.Errorf("%s\n%s", "Request creation failed", err)
	}
	req.SetBasicAuth(key, secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s\n%s", "Authentication Failed", err)
	}
	defer res.Body.Close()

	var token oauth2.Token
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&token); err != nil {
		panic(fmt.Sprintf("%s\n%s", "Response decoding failed", err))
	}

	config := &oauth2.Config{}
	return config.Client(context.Background(), &token), nil
}

// GetRetweeters f
func (c *Client) GetRetweeters(tweetID string) ([]string, error) {
	// Get status of retweets of single tweet [https://developer.twitter.com/en/docs/twitter-api/v1/tweets/post-and-engage/api-reference/get-statuses-retweets-id]
	// [https://api.twitter.com/1.1/statuses/retweets/:id.json] id - tweet id // 1362379742410731520
	url := fmt.Sprintf("https://api.twitter.com/1.1/statuses/retweets/%s.json", tweetID)
	res, err := c.clientSingleton().Get(url)
	if err != nil {
		return nil, err
	}

	var retweets []struct {
		User struct {
			ScreenName string `json:"screen_name"`
		} `json:"user"`
	}
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&retweets)

	usernames := make([]string, 0, len(retweets))
	for _, tweet := range retweets {
		usernames = append(usernames, tweet.User.ScreenName)
	}
	return usernames, nil
}
