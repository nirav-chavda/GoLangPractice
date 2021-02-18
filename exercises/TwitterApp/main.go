package main

import (
	"encoding/csv"
	"encoding/json"
	"os"

	"github.com/nirav-chavda/practice/exercises/TwitterApp/twitter"
)

const (
	keyFile  string = "keys.json"
	tweetID  string = "1362379742410731520"
	userFile string = "users.csv"
)

func main() {
	key, secret, err := getKeys(keyFile)
	if err != nil {
		panic(err)
	}
	existingUsers := existing(userFile)
	client := twitter.Client{
		CosumerKey:     key,
		CosumerSercret: secret,
	}
	newUsers, err := client.GetRetweeters(tweetID)
	if err != nil {
		panic(err)
	}
	if err := saveToFile(userFile, merge(existingUsers, newUsers)); err != nil {
		panic(err)
	}
}

func getKeys(filename string) (string, string, error) {
	// one time need
	// so rather than creating type we've created variable
	var keys struct {
		Key    string `json:"key"`
		Secret string `json:"secret"`
	}
	file, err := os.Open(filename)
	if err != nil {
		return "", "", err
	}
	defer file.Close()
	d := json.NewDecoder(file)
	d.Decode(&keys)
	return keys.Key, keys.Secret, nil
}

func existing(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}
	}
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return []string{}
	}
	users := make([]string, 0, len(lines))
	for _, line := range lines {
		users = append(users, line[0])
	}
	return users
}

func merge(existingUsers, newUsers []string) []string {
	uniqueUsers := make(map[string]struct{}, 0) // we need only keys that's why value is empty struct
	for _, user := range existingUsers {
		uniqueUsers[user] = struct{}{}
	}
	for _, user := range newUsers {
		uniqueUsers[user] = struct{}{}
	}
	allUsers := make([]string, 0, len(uniqueUsers))
	for key := range uniqueUsers {
		allUsers = append(allUsers, key)
	}
	return allUsers
}

func saveToFile(filename string, content []string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	w := csv.NewWriter(file)
	for _, user := range content {
		w.Write([]string{user})
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}
	return nil
}
