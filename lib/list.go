package lib

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/go-github/github"
)

// type List struct {
// 	sinceDate      string
// 	untilDate      string
// 	debug          bool
// 	user           string
// 	accessToken    string
// 	settingsGistID string
// }

// func (l *List) Collect() (string, error) {
// 	ctx := context.Background()
// 	client := getClient(ctx, l.accessToken)

// 	sinceTime, err := getSinceTime(l.sinceDate)
// 	if err != nil {
// 		return "", err
// 	}

// 	untilTime, err := getUntilTime(l.untilDate)
// 	if err != nil {
// 		return "", err
// 	}

// 	events, err := NewEvents(ctx, client, l.user, sinceTime, untilTime, l.debug).Collect()
// 	if err != nil {
// 		return "", err
// 	}
// 	format := NewFormat(ctx, client, l.debug)

// 	parallelNum, err := getParallelNum()
// 	if err != nil {
// 		return "", err
// 	}

// 	sem := make(chan int, parallelNum)
// 	var lines Lines
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex

// 	for i, event := range events {
// 		wg.Add(1)
// 		go func(event *github.Event, i int) {
// 			defer wg.Done()
// 			sem <- 1
// 			line := format.Line(event, i)
// 			<-sem

// 			mu.Lock()
// 			defer mu.Unlock()
// 			lines = append(lines, line)
// 		}(event, i)
// 	}
// 	wg.Wait()

// 	allLines, err := format.All(lines)
// 	if err != nil {
// 		return "", err
// 	}

// 	return allLines, nil
// }

// List outputs formated GitHub events to stdout
func List(sinceDate, untilDate string, debug bool) error {
	user, err := getUser()
	if err != nil {
		return err
	}

	accessToken, err := getAccessToken()
	if err != nil {
		return err
	}

	sinceTime, err := getSinceTime(sinceDate)
	if err != nil {
		log.Fatal(err)
	}

	untilTime, err := getUntilTime(untilDate)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	client := getClient(ctx, accessToken)

	events, err := NewEvents(ctx, client, user, sinceTime, untilTime, debug).Collect()
	if err != nil {
		return err
	}
	var settings Settings
	if err := settings.Init(ctx, client, getGistID()); err != nil {
		return err
	}

	format := NewFormat(ctx, client, settings, debug)

	parallelNum, err := getParallelNum()
	if err != nil {
		return err
	}

	sem := make(chan int, parallelNum)
	var lines Lines
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i, event := range events {
		wg.Add(1)
		go func(event *github.Event, i int) {
			defer wg.Done()
			sem <- 1
			line := format.Line(event, i)
			<-sem

			mu.Lock()
			defer mu.Unlock()
			lines = append(lines, line)
		}(event, i)
	}
	wg.Wait()

	allLines, err := format.All(lines)
	if err != nil {
		return err
	}

	fmt.Print(allLines)

	return nil
}

func getSinceTime(sinceDate string) (time.Time, error) {
	return time.Parse("20060102 15:04:05 MST", sinceDate+" 00:00:00 "+getZoneName())
}

func getUntilTime(untilDate string) (time.Time, error) {
	result, err := time.Parse("20060102 15:04:05 MST", untilDate+" 00:00:00 "+getZoneName())
	if err != nil {
		return result, err
	}

	return result.AddDate(0, 0, 1).Add(-time.Nanosecond), nil
}

func getZoneName() string {
	zone, _ := time.Now().Zone()
	return zone
}
