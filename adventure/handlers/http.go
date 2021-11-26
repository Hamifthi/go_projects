package handlers

import (
	"choose_adventure"
	"net/http"
)

type HTTPHandler struct {
	Storage map[string]adventure.Story
}

func (h *HTTPHandler) GetFirstPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("only support get request\n"))
		return
	}
	firstStory, _ := h.Storage["intro"]
	w.WriteHeader(http.StatusOK)
	tmpl := parseHtmlTemplate("./handlers/template/story.html")
	tmpl.Execute(w, firstStory)
}

func (h *HTTPHandler) GetStory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("only support get request\n"))
		return
	}
	storyArc := r.URL.Query().Get("story_arc")
	if storyArc == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("you must send the story_arc\n"))
		return
	}
	story, ok := h.Storage[storyArc]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in getting story with provided story_arc\n"))
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl := parseHtmlTemplate("./handlers/template/story.html")
	tmpl.Execute(w, story)
}
