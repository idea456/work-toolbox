package notify

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ttacon/chalk"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

// retrieves the token from local if saved already, if not from web, and returns the HTTP client
func getClient(config *oauth2.Config) *http.Client {
        // The file token.json stores the user's access and refresh tokens, and is
        // created automatically when the authorization flow completes for the first
        // time.
        tokFile := "token.json"
        tok, err := tokenFromFile(tokFile)
        if err != nil {
                tok = getTokenFromWeb(config)
                saveToken(tokFile, tok)
        }
        return config.Client(context.Background(), tok)
}
func tokenFromFile(file string) (*oauth2.Token, error) {
        f, err := os.Open(file)
        if err != nil {
                return nil, err
        }
        defer f.Close()
        tok := &oauth2.Token{}
        err = json.NewDecoder(f).Decode(tok)
        return tok, err
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			state := r.URL.Query()["state"][0]
			if state != "state-token" {
				log.Fatalf("Error in authorizing!")
			}

			code := r.URL.Query()["code"][0]
			fmt.Fprintf(w, code)
			
		})
		log.Fatal(http.ListenAndServe(":4000", nil))
	}()
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to this link and key in the authorization code: ",);
	fmt.Println(chalk.Underline.TextStyle(authURL))
	fmt.Printf("%s => ", chalk.Magenta.Color("Authorization code"))

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err);
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve the token from web: %v", err)
	}
	return tok
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credentials to path %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to save credentials: %v", err)
	}
	
	defer f.Close()

	json.NewEncoder(f).Encode(token)
}

func Authenticate() {
	ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read credentials file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse credentials: %v", err)
	}
	client := getClient(config)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	user := "me"
	r, err := srv.Users.Threads.List(user).Q("from:notifications@github.com,label:inbox").Do();
	if err != nil {
		log.Fatalf("Unable to retrieve lables: %v", err)
	}

	for _, msg := range r.Threads {
		// msgReq, _ := srv.Users.Messages.Get(user, msg.Id).Do();
		fmt.Printf("- %s\n", msg.Id)
	}

	msgReq, _ := srv.Users.Threads.Get(user, "1818e9f50076c315").Do();

	fmt.Println(msgReq.Messages[len(msgReq.Messages)- 1].Id)

	mr, _ := srv.Users.Messages.Get(user, "1818e9f50076c315").Do();

	decoded, _ := base64.URLEncoding.DecodeString(mr.Payload.Parts[0].Body.Data);
	fmt.Println(string(decoded));
}