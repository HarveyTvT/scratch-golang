package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type EventType string

const (
	CommitCommentEvent           EventType = "CommitCommentEvent"
	CreateEvent                  EventType = "CreateEvent"
	DeleteEvent                  EventType = "DeleteEvent"
	ForkEvent                    EventType = "ForkEvent"
	GollumEvent                  EventType = "GollumEvent"
	IssueCommentEvent            EventType = "IssueCommentEvent"
	IssueEvent                   EventType = "IssueEvent"
	MemberEvent                  EventType = "MemberEvent"
	PublicEvent                  EventType = "PublicEvent"
	PullRequestEvent             EventType = "PullRequestEvent"
	PullRequestReviewEvent       EventType = "PullRequestReviewEvent"
	PullRequestReviewThreadEvent EventType = "PullRequestReviewThreadEvent"
	PushEvent                    EventType = "PushEvent"
	ReleaseEvent                 EventType = "ReleaseEvent"
	SponsorshipEvent             EventType = "SponsorshipEvent"
	WatchEvent                   EventType = "WatchEvent"
)

type Actor struct {
	Id           uint64 `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarId   string `json:"gravatar_id"`
	Url          string `json:"url"`
	AvatarUrl    string `json:"avatar_url"`
}

type Repo struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Payload struct {
	Action string `json:"action"`
	// CreateEvent
	RefType      string `json:"ref_type"`
	MasterBranch string `json:"master_branch"`
	Description  string `json:"description"`
	PusherType   string `json:"pusher_type"`
	// ForkEvent
	Forkee Repo `json:"forkee"`
}

type Event struct {
	Id        string    `json:"id"`
	Type      EventType `json:"type"`
	Actor     Actor     `json:"actor"`
	Repo      Repo      `json:"repo"`
	Payload   Payload   `json:"payload"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
}

func (t *Event) String() string {
	var result string
	switch t.Type {
	case CommitCommentEvent:
		action := cases.Title(language.English).String(t.Payload.Action)
		result = fmt.Sprintf("%s commit comment in %s", action, t.Repo.Name)
	case CreateEvent:
		switch t.Payload.RefType {
		case "repository":
			result = fmt.Sprintf("Create a repository %s", t.Repo.Name)
		case "branch", "tag":
			result = fmt.Sprintf("Create a %s in %s", t.Payload.RefType, t.Repo.Name)
		}
	case DeleteEvent:
		result = fmt.Sprintf("Delete a %s in %s", t.Payload.RefType, t.Repo.Name)
	case ForkEvent:
		result = fmt.Sprintf("Fork a repository %s", t.Repo.Name)
	case GollumEvent:
		result = fmt.Sprintf("Create/Update a wiki page in %s", t.Repo.Name)
	case IssueCommentEvent:
		action := cases.Title(language.English).String(t.Payload.Action)
		result = fmt.Sprintf("%s comment in %s", action, t.Repo.Name)
	case IssueEvent:
		action := cases.Title(language.English).String(t.Payload.Action)
		result = fmt.Sprintf("%s issue in %s", action, t.Repo.Name)
	case MemberEvent:
		action := cases.Title(language.English).String(t.Payload.Action)
		result = fmt.Sprintf("%s a member in %s", action, t.Repo.Name)
	case PublicEvent:
		result = fmt.Sprintf("Open source a repository %s", t.Repo.Name)
	case PullRequestEvent:
		action := cases.Title(language.English).String(t.Payload.Action)
		result = fmt.Sprintf("%s pull request in %s", action, t.Repo.Name)
	case PullRequestReviewEvent:
		action := cases.Title(language.English).String(t.Payload.Action)
		result = fmt.Sprintf("%s pull request review in %s", action, t.Repo.Name)
	case PullRequestReviewThreadEvent:
		action := cases.Title(language.English).String(t.Payload.Action)
		result = fmt.Sprintf("%s pull request review thread in %s", action, t.Repo.Name)
	case PushEvent:
		result = fmt.Sprintf("Push a commit in %s", t.Repo.Name)
	case ReleaseEvent:
		result = fmt.Sprintf("Publish a new release in %s", t.Repo.Name)
	case SponsorshipEvent:
		result = fmt.Sprintf("Sponsor a user in %s", t.Repo.Name)
	case WatchEvent:
		result = fmt.Sprintf("Watch a repository %s", t.Repo.Name)
	default:
		return "Unknown event"
	}

	return result
}

func GetActivity(username string) ([]*Event, error) {
	apiUrl := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to get activity: %s", resp.Status)
	}

	var events []*Event
	err = json.NewDecoder(resp.Body).Decode(&events)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func main() {
	var username string
	if len(os.Args) > 1 {
		username = os.Args[1]
	}
	events, err := GetActivity(username)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(events) == 0 {
		fmt.Println("No activity found")
		return
	}

	fmt.Println("output: ")

	for _, event := range events {
		fmt.Printf("- %s\n", event.String())
	}
}
