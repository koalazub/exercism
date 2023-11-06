package electionday

import "fmt"

// Create a function `NewVoteCounter` that accepts the number of initial votes for a candidate and returns a pointer referring to an `int`, initialized with the given number of initial votes.

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {

	var p *int = &initialVotes
	return p
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}

	return *counter
}

// Create a function `IncrementVoteCount` that will take a counter (`*int`) as an argument and a number of votes, and will increment the counter by that number of votes. You can assume the pointer passed will never be `nil`.
func IncrementVoteCount(counter *int, increment int) {
	*counter = *counter + increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var result *ElectionResult
	result = &ElectionResult{Name: candidateName, Votes: votes}

	return result
}

// DisplayResult creates a message with the result to be displayed.
// `<candidate_name> (<votes>)`
func DisplayResult(result *ElectionResult) string {
	if result == nil {
		return ""

	}
	return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate] -= 1
}
