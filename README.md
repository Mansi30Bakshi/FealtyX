# **GoLang CRUD Assignment**  

This project is a basic REST API built in **Go**, designed as an assignment to demonstrate CRUD operations along with an AI-based summary generation feature. It keeps everything in memory using map data structure, making it light and ideal for testing.  

---

## **Table of Contents**  

1. [Features](#features)  
2. [Setup](#setup)  
3. [Usage](#usage)  
4. [API Endpoints](#api-endpoints)  
5. [Screenshots](#screenshots)  

---

## **Features**  

- **Basic Create, Read, Update, Delete (CRUD) operations.**  
- **Summary Generation**: AI-powered summary of item details.  
- **In-memory data storage** (no external database required).  
- **Basic validation and error handling.**  
- **Concurrently safe operations on data.**  

---

## **Setup**  

### **Prerequisites**  

- **Go 1.16+ installed.**  
- **Ollama for AI-powered summary generation**, installed and running locally (default port `11434`).  
- **Postman** (optional, for testing).  

---

## **Usage**  

You can test the API endpoints with **Postman** or by sending HTTP requests directly from your terminal. Below are descriptions of each endpoint, with example requests and responses.  

---

## **API Endpoints**  

1. **Create a New Item**  
   - **Endpoint**: `POST /students`  
   - **Description**: Adds a new student to the in-memory store.  

2. **Get All Items**  
   - **Endpoint**: `GET /students`
   - **Description**: Gets all students that are stored in memory.   

3. **Get Item by ID**  
   - **Endpoint**: `GET /students/{id}`
   - **Description**: Gets the students from the memory wit a definite id.  

4. **Update an Item**  
   - **Endpoint**: `PUT /students/{id}`
   - **Description**: Update the student with definite id  that is stored in memory. 

5. **Delete an Item**  
   - **Endpoint**: `DELETE /students/{id}`
   - **Description**: Delete the student with a definite id that is stores in memory.  

6. **Generate Summary for an Item**  
   - **Endpoint**: `GET /students/{id}/summary`
   - **Description**: Generate the summary of student with definite id with help of Ollama. 

---

## **Screenshots**  

Screenshots showcasing API functionality using **Postman** and a web browser are available in the `assets` folder within the project directory.  

---
