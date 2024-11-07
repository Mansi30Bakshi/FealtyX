package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"regexp"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"gte=18,lte=100"`
	Email string `json:"email" validate:"required,email"`
}

var (
	students  = make(map[int]Student)
	idCounter = 1
	validate  = validator.New()
)


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/students", CreateStudent).Methods("POST")
	router.HandleFunc("/students", GetAllStudents).Methods("GET")
	router.HandleFunc("/students/{id}", GetStudentByID).Methods("GET")
	router.HandleFunc("/students/{id}", UpdateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", DeleteStudent).Methods("DELETE")
	router.HandleFunc("/students/{id}/summary", GetStudentSummary).Methods("GET")
	fmt.Println("Server running at http://localhost:3031")
	log.Fatal(http.ListenAndServe(":3031", router))
}


func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil || validate.Struct(student) != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	student.ID = idCounter
	students[idCounter] = student
	idCounter++
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}
func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	studentList := make([]Student, 0, len(students))
	for _, student := range students {
		studentList = append(studentList, student)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studentList)
}
func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr) 
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	student, exists := students[id]
	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}


func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr) 
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var student Student
	err = json.NewDecoder(r.Body).Decode(&student)
	if err != nil || validate.Struct(student) != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	existingStudent, exists := students[id]
	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	existingStudent.Name = student.Name
	existingStudent.Age = student.Age
	existingStudent.Email = student.Email
	students[id] = existingStudent

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingStudent)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr) // Convert string to int
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	_, exists := students[id]
	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	delete(students, id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Student deleted successfully"})
}
func GetStudentSummary(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	student, exists := students[id]
	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	summary, err := GenerateOllamaSummary(student)
	if err != nil {
		http.Error(w, "Failed to generate summary", http.StatusInternalServerError)
		return
	}
	cleanSummary := removeANSISequences(summary)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"summary": cleanSummary})
}
func GenerateOllamaSummary(student Student) (string, error) {
	os.Setenv("HOME", "/Users/mansibakshi")
	prompt := fmt.Sprintf("Generate a detailed profile summary for a student named %s, age %d, and email %s.", student.Name, student.Age, student.Email)
	cmd := exec.Command("ollama", "run", "llama3:latest", prompt)
	cmd.Env = append(os.Environ(), "OLLAMA_API_URL=http://localhost:11434", "HOME=/Users/mansibakshi")

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error running Ollama command: %v\nOutput: %s", err, string(out))
		return "", fmt.Errorf("Ollama command failed: %v", err)
	}
	summary := strings.TrimSpace(string(out))
	
	if summary == "" {
		return "", fmt.Errorf("empty response from Ollama")
	}
	summary = removeANSISequences(summary)

	return summary, nil
}

func removeANSISequences(input string) string {
	re := regexp.MustCompile(`\x1B(?:[@-Z\\-_]|\[[0-?]*[ -/]*[@-~])`)
	input = re.ReplaceAllString(input, "")
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\r", "")
	input = strings.ReplaceAll(input, "*", "")

	return input
}



