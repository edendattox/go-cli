package todo

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{Task: task, Done: false, CreatedAt: time.Now(), CompletedAt: time.Now()}

	*t = append(*t, todo)

}

func (t *Todos) Complete(index int) error {
	ls := *t

	if index <= 0 || index > len(ls) {
		return errors.New("Invalid index")
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil

}

func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("Invalid Index")
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil

}

func (t *Todos) Load(filename string) error {

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)

	if err != nil {
		return err
	}

	return nil

}

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)

}

func (t *Todos) Print() {
	for i, item := range *t {
		fmt.Printf("%d - %s\n", i, item.Task)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("empty todo is not allowed")
	}

	return scanner.Text(), nil

}
