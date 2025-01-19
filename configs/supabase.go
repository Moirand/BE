package configs

import (
    "github.com/supabase-community/supabase-go"
)

func InitSupabase() (*supabase.Client, error) {
    client := supabase.CreateClient(
        "https://lotpaaakejmnmbrmgqcb.supabase.co",
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImxvdHBhYWFrZWptbm1icm1ncWNiIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MzcxMTAyMDksImV4cCI6MjA1MjY4NjIwOX0.JHnb9xfAIz9CfU0qGqXKJBLFNVhWYeDG1-8fBxI_U_I",
    )
    return client, nil
}