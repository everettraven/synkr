package github

type RepoItemType string

const (
	RepoItemTypeIssue       RepoItemType = "Issue"
	RepoItemTypePullRequest RepoItemType = "PullRequest"
)

// TODO: Expand here to capture more things that
// may be important to filter on
type RepoItem struct {
	ID        int64        `json:"id"`
	URL       string       `json:"url"`
	Author    string       `json:"author"`
	Labels    []string     `json:"labels"`
	Type      RepoItemType `json:"type"`
	Assignees []string     `json:"assignees"`
	Title     string       `json:"title"`
	Body      string       `json:"body"`
	State     string       `json:"state"`
	Priority  int          `json:"priority"`
	Status    string       `json:"status"`
	Created   string       `json:"created"`
	Updated   string       `json:"updated"`
	Comments  int          `json:"comments"`
	Milestone string       `json:"milestone"`
	Mentions  []string     `json:"mentions"`

	// Only populated on PullRequests
	RequestedReviewers []string `json:"requestedReviewers"`
	Draft              bool     `json:"draft"`
}
