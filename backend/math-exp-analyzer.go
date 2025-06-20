package main

import (
	//Read files line by line
	// Reads arguments from the command-line
	"fmt" // Prints on screen
	"log" // Shows errors and ends program if an error happens
	"net/http"
	"os" // read and write on files
	"strings"

	"github.com/gin-contrib/cors" //para qagregar soporte de cors
	"github.com/gin-gonic/gin"

	// Para manipular cadenas de texto
	"gopkg.in/yaml.v2" // read yaml files
)

// Struct with the PDA formal definition established in the yaml
type PDA struct {
	InitialState  string     `yaml:"S"`
	States        []string   `yaml:"K"` // All the states
	Alphabet      []string   `yaml:"E"` // Alphabet of the readen expressions
	Transitions   [][]string `yaml:"T"`
	FinalStates   []string   `yaml:"F"`
	AlphabetStack []string   `yaml:"P"`
}

// Struct to save 1 transition (this to have the values well identified)
type Transition struct {
	input     string // char that should read from the expression
	pop       string // char popped from the stack
	nextState string
	push      string // char pushed to the stack
}

// This struct was done by copilot and it is an aid for the function printTransitionLog, to visualize the transitions undergone
type TransitionLog struct {
	fromState string
	toState   string
	input     string
	pop       string
	push      string
	stack     []string
}

// Express is an attribute in which +if a json arrives, it will look for the atribute expression and save it in Express attribute as a string, the binding is to tell that it is obligatory that the json has it, if not, it will return an error
type Expression struct {
	Express string `json:"expression" binding:"required"`
}

func status(c *gin.Context) { //function that checks the server connection
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

// this function is the one that recieves the expression from the front and returns a json with the accepted state and the expression
func validateExpression(c *gin.Context) {
	var express Expression
	err := c.ShouldBindJSON(&express) //reads  the json sent by react and saves it in the variable express
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //checks if the json has the correct format
		return
	}

	mathExpression := strings.ReplaceAll(express.Express, " ", "") //deletes all white spaces bacause the pda doesn't detect them
	valid := processExpression(mathExpression)                     //validates the expression, if valid returns true , if not is false. This answer is sent in the json

	c.JSON(http.StatusOK, gin.H{ //sends the answer to react in a json
		"valid": valid,
	})

}

// Function that reads the pda saves it in a struct
func buildPDA() {
	yamlFile, err := os.ReadFile("pda.yaml") // Reads the YAML file
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err) // Finishes program if it couldn't read the yaml
	}

	err = yaml.Unmarshal(yamlFile, &pda) // Converts the PDA in the yaml to the pda struct
	if err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}
	mappedTransitions = mapTransitions(pda)

}

// Transform the list  of transitions into a map with slices of Transition structs
func mapTransitions(pda PDA) map[string][]Transition {
	transitions := make(map[string][]Transition) //we create a map of Transitions, so that the key is the actual state and the value is a slice of Transitions struct correspoding to that state, each struct  saves the symbol it is reading, the symbol that will be popped from the stack, the next state it will arrive to after choosing that transition and the symbol that will be pushed into the stack

	for _, currentTransition := range pda.Transitions {
		currentState, readSymbol, popSymbol, nextState, pushSymbol := currentTransition[0], currentTransition[1], currentTransition[2], currentTransition[3], currentTransition[4]

		//Adds a new transtion to the corresponding starting state or current state
		transitions[currentState] = append(transitions[currentState], Transition{
			input:     readSymbol,
			pop:       popSymbol,
			nextState: nextState,
			push:      pushSymbol,
		})
	}
	return transitions
}

// This function is to print the transitions it passes through, it was done by copilot
func printTransitionLog(logs []TransitionLog) {
	fmt.Println("\nTransitions Path:")
	fmt.Printf("%-10s | %-10s | %-8s | %-8s | %-8s | %-20s\n",
		"From", "To", "Input", "Pop", "Push", "Stack")
	fmt.Println("----------|-----------|---------|---------|---------|-------------------")
	for _, log := range logs {
		fmt.Printf("%-10s | %-10s | %-8s | %-8s | %-8s | %v\n",
			log.fromState,
			log.toState,
			log.input,
			log.pop,
			log.push,
			log.stack)
	}
}

// Process the expression using the PDA
func processExpression(expression string) bool {

	currentState := pda.InitialState // Always start in initial state
	stack := []string{}
	currentPosition := 0 // position of the char you are reading from the string
	expressionValid := false
	transitionLogs := []TransitionLog{} // Stores the path of the transitions (done by copilot)

	// Reads the each character of 1 expression
	for currentPosition < len(expression) {
		foundTransition := false
		var possibleTransitions []Transition

		// Caracter actual de la expresión
		currentChar := string(expression[currentPosition])

		// Stores in Transition slice the transitions that start in the current state and have the char you are reading
		for _, transition := range mappedTransitions[currentState] {
			if transition.input == currentChar {
				possibleTransitions = append(possibleTransitions, transition)
			}
		}

		//If there aren't transitions with the current char, search for empty transitions
		if len(possibleTransitions) == 0 {
			for _, transition := range mappedTransitions[currentState] {
				if transition.input == "ε" {
					possibleTransitions = append(possibleTransitions, transition)
				}
			}
		}

		// Processing the possible transitions
		for _, transition := range possibleTransitions {

			stackAvailable := false

			//can continue validating that transition if the transition doesn't need to pop something or if it needs to pop something, checks that the stack isn't empty and that the char to pop defined in that transition is equal to the last element of the stack
			if transition.pop == "ε" || (len(stack) > 0 && stack[len(stack)-1] == transition.pop) {
				stackAvailable = true
			}

			// This doesn't allow ([]) by checking if the element to push defined by that transition and the top element in the stack is a (, it stops checking that transition and continues checking the next transition
			if transition.push == "[" && len(stack) > 0 && stack[len(stack)-1] == "(" {
				continue
			}

			if stackAvailable {
				//with this it reads **, after arriving to q8 by reading one *, the char of the transition is  *, if you aren't at the end of the expression it checks the next char and sees if it is another *. If it is, it looks for the transition that accepts this second * and takes you to the next state
				if currentState == "q8" && transition.input == "*" && currentPosition < len(expression)-1 && expression[currentPosition+1] == '*' {

					//looks for the transition in q8 as starting state that has the other * and saves it in transition
					for _, asterisk := range mappedTransitions["q8"] {
						if asterisk.input == "*" {
							transition = asterisk
							break
						}
					}
				}

				// Saves the stackstate before the transition in the stack that will be read to print it in screen
				stackCopy := make([]string, len(stack))
				copy(stackCopy, stack)

				// If a pop has to be done,it takes out the element from the stack
				if transition.pop != "ε" {
					stack = stack[:len(stack)-1] //deletes the last element in a slice
				}

				// If a push  has to be done,it inserts the element in the stack
				if transition.push != "ε" {
					stack = append(stack, transition.push)
				}

				// Save the transitions in the struct that will be used to print the transitions path for each expression
				transitionLogs = append(transitionLogs, TransitionLog{
					fromState: currentState,
					toState:   transition.nextState,
					input:     transition.input,
					pop:       transition.pop,
					push:      transition.push,
					stack:     stackCopy,
				})

				currentState = transition.nextState // changes to next state dictated by the transition taken

				//if a char from the expression was consumed, it mvoes to the next char
				if transition.input != "ε" {
					currentPosition++
				}

				foundTransition = true
				break // Takes the 1st valid transition
			}
		}
		//If it didn't find a valid transition, it stops looking for chars in this expression and goes to evaluating the next expression
		if !foundTransition {
			break
		}
	}

	// Checks that it ends in a final state
	endsFinalState := false
	for _, state := range pda.FinalStates {
		if state == currentState {
			endsFinalState = true
		}
	}

	//Checks if it ended reading the expression, ends in final state and the stack is empty, if all valid, expression is valid
	if currentPosition == len(expression) && endsFinalState && len(stack) == 0 {
		expressionValid = true
	}

	printTransitionLog(transitionLogs)

	return expressionValid
}

var pda PDA
var mappedTransitions map[string][]Transition

func main() {

	buildPDA() //Reads tha automata from the yaml

	server := gin.Default()
	server.Use(cors.Default())
	server.GET("/status", status)
	server.POST("/validate", validateExpression)

	server.Run()

}
