package kit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	BaseURL             = "https://leetcode.com"
	GraphQLURL          = "https://leetcode.com/graphql"
	AllProblemsURL      = "https://leetcode.com/api/problems/all"
	AllProblemsFilePath = "./problems.json"
)

var RootDir, _ = os.Getwd()

type problem struct {
	ID          int
	Title       string
	TitleSlug   string
	URL         string
	Snippet     string
	FuncName    string
	PackageName string
}

type LeetCode struct {
	all   AllData
	index map[int]*problem
}

var leetcode = new(LeetCode)

func fetchAllProblems() error {
	fmt.Print("Fetching problems...")
	resp, err := http.Get(AllProblemsURL)
	if err != nil {
		return fmt.Errorf("can't fetch problems from url: %s: %w", AllProblemsURL, err)
	}
	defer resp.Body.Close()

	file, err := os.Create(AllProblemsFilePath)
	if err != nil {
		return fmt.Errorf("can't create file: %s", AllProblemsFilePath)
	}
	if _, err = io.Copy(file, resp.Body); err != nil {
		return fmt.Errorf("can't write file: %s", AllProblemsFilePath)
	}

	return nil
}

func Init() error {
	stat, err := os.Stat(AllProblemsFilePath)
	if stat != nil && stat.IsDir() {
		return fmt.Errorf("%s is a directory", AllProblemsFilePath)
	}
	if errors.Is(err, fs.ErrNotExist) {
		if err = fetchAllProblems(); err != nil {
			return fmt.Errorf("can't fetch all problems: %w", err)
		}
	}

	file, err := os.ReadFile(AllProblemsFilePath)
	if err != nil {
		return fmt.Errorf("can't open file %s", AllProblemsFilePath)
	}

	err = json.Unmarshal(file, &leetcode.all)
	if err != nil {
		return fmt.Errorf("can't parse json: %s: %w", AllProblemsFilePath, err)
	}

	leetcode.index = make(map[int]*problem)
	for _, p := range leetcode.all.StatStatusPairs {
		leetcode.index[p.Stat.QuestionID] = &problem{
			ID:        p.Stat.QuestionID,
			Title:     p.Stat.QuestionTitle,
			TitleSlug: p.Stat.QuestionTitleSlug,
			URL:       fmt.Sprintf("%s/problems/%s", BaseURL, p.Stat.QuestionTitleSlug),
		}
	}

	return nil
}

func fetchQuestionData(problem *problem) ([]byte, error) {
	var jsonBody = []byte(fmt.Sprintf(`
{
    "operationName": "questionData",
    "variables": {
        "titleSlug": %q
    },
    "query": "query questionData($titleSlug: String!) {\nquestion(titleSlug: $titleSlug) {\nquestionId\nquestionFrontendId\ntitle\ntitleSlug\ncontent\ndifficulty\nexampleTestcases\ncodeSnippets {\nlang\nlangSlug\ncode\n}\n}}"
}
`, problem.TitleSlug))

	req, err := http.NewRequest("POST", GraphQLURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("can't request problem details: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Host", BaseURL)
	req.Header.Add("Referer", problem.URL)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't request problem details: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("bad response for problem request: %w", err)
	}

	return data, nil
}

func Get(id int) (*problem, error) {
	fmt.Printf("Getting data for problem %d\n", id)
	problem, ok := leetcode.index[id]
	if !ok {
		return nil, fmt.Errorf("no problem %d", id)
	}

	data, err := fetchQuestionData(problem)
	if err != nil {
		return nil, err
	}

	q := &QuestionData{}
	err = json.Unmarshal(data, q)
	if err != nil {
		return nil, fmt.Errorf("can't parse question data for problem %s: %w", id, err)
	}

	for _, s := range q.Data.Question.CodeSnippets {
		if s.LangSlug == "golang" {
			problem.Snippet = s.Code
			break
		}
	}

	if problem.Snippet == "" {
		return nil, fmt.Errorf("problem %n doesn't have solution in golang", id)
	}

	return problem, nil
}

var solutionTmpls = template.Must(template.ParseFiles("./kit/source.tmpl", "./kit/test.tmpl", "./kit/readme.tmpl"))

func Generate(problem *problem) {
	fmt.Println("Generating template...")
	n := problem.ID
	name := strings.Replace(problem.TitleSlug, "-", "_", -1)

	dir := fmt.Sprintf("%04d.%s", n, name)
	path := filepath.Join(RootDir, dir)
	if err := os.Mkdir(path, 0777); err != nil {
		panic(err)
	}

	generateScaffoldFile("source.tmpl", filepath.Join(path, name+".go"), problem)
	generateScaffoldFile("test.tmpl", filepath.Join(path, name+"_test.go"), problem)
	generateScaffoldFile("readme.tmpl", filepath.Join(path, "README.md"), problem)
}

func generateScaffoldFile(tmplName, fileName string, data interface{}) error {
	targetFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer targetFile.Close()
	if err = solutionTmpls.Lookup(tmplName).Execute(targetFile, data); err != nil {
		return err
	}
	return nil
}
