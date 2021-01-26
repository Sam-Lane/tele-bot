package stackoverflow

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type StackOverResult struct {
	Tags  []string `json:"tags"`
	Owner struct {
		Reputation   int    `json:"reputation"`
		UserID       int    `json:"user_id"`
		UserType     string `json:"user_type"`
		AcceptRate   int    `json:"accept_rate"`
		ProfileImage string `json:"profile_image"`
		DisplayName  string `json:"display_name"`
		Link         string `json:"link"`
	} `json:"owner"`
	IsAnswered       bool   `json:"is_answered"`
	ViewCount        int    `json:"view_count"`
	ProtectedDate    int    `json:"protected_date"`
	AcceptedAnswerID int    `json:"accepted_answer_id"`
	AnswerCount      int    `json:"answer_count"`
	Score            int    `json:"score"`
	LastActivityDate int    `json:"last_activity_date"`
	CreationDate     int    `json:"creation_date"`
	LastEditDate     int    `json:"last_edit_date"`
	QuestionID       int    `json:"question_id"`
	ContentLicense   string `json:"content_license"`
	Link             string `json:"link"`
	Title            string `json:"title"`
}

type StackOverResults struct {
	Items []StackOverResult
}

func StackOverFlowQuery(msg *tgbotapi.Message, bot *tgbotapi.BotAPI, reply *tgbotapi.MessageConfig) {
	res, err := query(msg.CommandArguments())
	if err != nil {
		reply.Text = err.Error()
		bot.Send(reply)
		return
	}
	msg.Text = res.Items[0].Link
	bot.Send(reply)
}

func query(query string) (StackOverResults, error) {
	sRes := StackOverResults{}
	safe := url.QueryEscape(query)
	url := fmt.Sprintf("https://api.stackexchange.com/2.2/search?order=desc&sort=votes&intitle=%s&site=stackoverflow", safe)
	resp, err := http.Get(url)

	if err != nil {
		log.Printf("failed to retrieve query from StackOverFlow: %s", err.Error())
		return sRes, fmt.Errorf("Failed to retrieve query")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("failed reading body from request: %s", err.Error())
		return sRes, fmt.Errorf("failed to retrieve query")
	}

	err = json.Unmarshal(body, &sRes)
	if err != nil {
		log.Printf("failed to unmarshal body to struct: %s", err.Error())
	}
	if len(sRes.Items) == 0 {
		return sRes, fmt.Errorf("No results for your query 🙃")
	}
	return sRes, nil
}
