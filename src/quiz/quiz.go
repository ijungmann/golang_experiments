package main

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {
	err := LoadConfig()
	if err != nil {
		return
	}
	fmt.Println("Welcome to the quiz!")
	fmt.Printf("You have %s!\n", viper.Get("quiz.time"))
	fmt.Println("Timer starts when you hit enter!")
	duration, err := time.ParseDuration(viper.Get("quiz.time").(string))
	if err != nil {
		panic(err)
	}
	timer := time.NewTimer(duration)
	go func() {
		<-timer.C
		fmt.Println("Times Up!")
		os.Exit(0)
	}()
	_, err = fmt.Scanln()
	if err != nil {
		return
	}
	quiz, err := os.Open("resources/quiz.csv")
	defer quiz.Close()
	reader := csv.NewReader(quiz)
	if err != nil {
		panic(err)
	}
	var missedQuestions []string
	var questionCount int
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		questionCount++
		var userAnswer int
		fmt.Println(line[0])
		realAnswer, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}
		_, err = fmt.Scanln(&userAnswer)
		if err != nil {
			panic(err)
		}
		if userAnswer != realAnswer {
			missedQuestions = append(missedQuestions, line[0])
		}
	}
	fmt.Printf("Score: %d out of %d\n", questionCount-len(missedQuestions), questionCount)
	fmt.Println("Missed Questions:")
	for _, question := range missedQuestions {
		fmt.Printf("%s\n", question)
	}

}
