package common

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteSuccessResponse(ctx context.Context, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println("Failed to write success response:", err)
	}
}
