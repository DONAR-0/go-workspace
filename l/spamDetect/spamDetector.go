package main

import (
	"fmt"
	"log"
	"math"
	"strings"
)

type Classifier struct {
	spamWords,
	hamWords map[string]int
	totalSpamWords,
	totalHamWords,
	totalSpamDocs,
	totalHamDocs int
}

func NewClassifier() *Classifier {
	return &Classifier{
		spamWords: make(map[string]int),
		hamWords:  make(map[string]int),
	}
}

func (c *Classifier) Train(message string, isSpam bool) {
	words := tokenize(message)

	if isSpam {
		c.totalSpamDocs++
		for _, word := range words {
			c.spamWords[word]++
			c.totalSpamWords++
		}
	} else {
		c.totalHamDocs++
		for _, word := range words {
			c.hamWords[word]++
			c.totalHamWords++
		}
	}
}

// Classify predicts whether a new message is spam or not.
// It uses a Naive Bayes algorithm with log probabilities to avoid floating-point underflow.
func (c *Classifier) Classify(message string) bool {
	// A small number (alpha) is added to avoid zero probabilities. This is called Laplace smoothing.
	alpha := 1.0
	// Calculate the prior probability of a message being spam or ham.
	// We use log probabilities for numerical stability.
	spamLogProb := math.Log(float64(c.totalSpamDocs) / float64(c.totalSpamDocs+c.totalHamDocs))
	hamLogProb := math.Log(float64(c.totalHamDocs) / float64(c.totalSpamDocs+c.totalHamDocs))

	// Get the words from the new message.
	words := tokenize(message)
	totalWords := c.totalSpamWords + c.totalHamWords

	// Iterate over each word in the message and update the probabilities.
	for _, word := range words {
		// Calculate the probability of the word appearing in spam messages.
		// We use Laplace smoothing here: (word_count + alpha) / (total_spam_words + alpha * total_unique_words)
		spamWordProb := float64(c.spamWords[word]+int(alpha)) / float64(c.totalSpamWords+int(alpha)*totalWords)
		spamLogProb += math.Log(spamWordProb)

		// Calculate the probability of the word appearing in ham messages.
		hamWordProb := float64(c.hamWords[word]+int(alpha)) / float64(c.totalHamWords+int(alpha)*totalWords)
		hamLogProb += math.Log(hamWordProb)
	}

	// Compare the final log probabilities. The higher one indicates the class.
	return spamLogProb > hamLogProb
}

func tokenize(text string) []string {
	text = strings.ToLower(text)
	text = strings.NewReplacer(".", " ", ",", " ", "!", " ", "?", " ").Replace(text)
	words := strings.Fields(text)

	return words
}

func main() {
	fmt.Println("Starting Golang Spam Detector...")

	// Create a new classifier instance.
	classifier := NewClassifier()

	// --- TRAINING DATA ---
	// Define a set of training messages, both spam and ham.
	fmt.Println("\n--- Training the model ---")

	spamExamples := []string{
		"Free money now, click here for prize!",
		"Congratulations, you have won a free iPhone. Claim your prize!",
		"Urgent: You won 1 million dollars. Send your bank account details.",
		"Lose weight fast with our new diet pills!",
	}
	hamExamples := []string{
		"Hello, can we meet for lunch tomorrow?",
		"Did you see the new movie last night? It was great.",
		"The project meeting is scheduled for 10 AM.",
		"What is the plan for the weekend trip?",
	}

	// Train the classifier with the examples.
	for _, msg := range spamExamples {
		classifier.Train(msg, true) // isSpam is true
	}

	for _, msg := range hamExamples {
		classifier.Train(msg, false) // isSpam is false
	}

	fmt.Println("Model trained with example data.")

	// --- TESTING THE CLASSIFIER ---
	fmt.Println("\n--- Testing the classifier ---")

	testMessages := []string{
		"Claim your free prize now!",                    // Should be classified as spam.
		"Can you confirm the meeting time?",             // Should be classified as ham.
		"Urgent: a new email is ready for you to read.", // Should be classified as spam.
		"Do you want to get coffee this afternoon?",     // Should be classified as ham.
		"Win a new car, you've been selected!",          // Should be classified as spam.
	}

	for _, msg := range testMessages {
		isSpam := classifier.Classify(msg)
		if isSpam {
			log.Printf("Message: \"%s\" -> SPAM", msg)
		} else {
			log.Printf("Message: \"%s\" -> HAM", msg)
		}
	}
}
