package main

import (
	"fmt"
)

type ElectionResult struct {
	Name  string
	Votes int
}

// NewVoteCounter returns a new vote counter with
// a given number of inital votes.
func NewVoteCounter(initialVotes int) *int {
	votes := initialVotes
	return &votes

}

// IncrementVoteCount increments the value in a vote counter
func VoteCount(counter *int) int {
	votes := counter
	if counter == nil {
		return 0
	}
	return *votes
}

// NewElectionResult creates a new election result
func IncrementVoteCount(counter *int, increment int) {
	*counter = *counter + increment
	fmt.Println(*counter)
}

// NewElectionResult creates a new election result
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var result *ElectionResult
	result = &ElectionResult{Name: candidateName, Votes: votes}
	return result

}

// DisplayResult creates a message with the result to be displayed
func DisplayResult(result *ElectionResult) string {
	name := result.Name
	votes := result.Votes
	final := name + " " + "(" + fmt.Sprint(votes) + ")"
	return final
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate] -= 1
	fmt.Println(results[candidate])

}

func main() {
	initialVotes := 3
	votecounter := &initialVotes
	var finalResults = map[string]int{
		"Mary": 10,
		"John": 51,
	}
	fmt.Println(NewVoteCounter(initialVotes))
	fmt.Println(VoteCount(votecounter))
	IncrementVoteCount(votecounter, 2)
	fmt.Println(NewElectionResult("Peter", 3))
	(DecrementVotesOfCandidate(finalResults, "John"))

}
