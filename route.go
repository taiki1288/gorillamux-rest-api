package main

import (
	"encoding/json"
	"net/http"

	"./repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func getPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err  != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error getting the posts"}`)) 
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func addPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err  != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling the request"}`))
		return 
	}
	post.Id = len(posts) + 1
	posts = append(posts, post)
	resp.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	resp.Write(result)
}