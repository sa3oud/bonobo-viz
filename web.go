package main

import (
    "fmt"
    "net/http"
    "os/exec"
    "strings"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        
        // Run the visualization
        cmd := exec.Command("go", "run", "main.go")
        output, err := cmd.Output()
        if err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        
        // Convert to HTML with proper formatting
        html := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>BonoboLab Community Visualization</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <style>
        body {
            background: #050508;
            color: #f0f0f2;
            font-family: 'Courier New', monospace;
            line-height: 1.5;
            margin: 0;
            padding: 2rem;
            display: flex;
            justify-content: center;
            min-height: 100vh;
        }
        pre {
            background: #0a0a0f;
            padding: 2rem;
            border-radius: 12px;
            border: 1px solid rgba(240,240,242,0.06);
            box-shadow: 0 10px 15px rgba(0,0,0,0.15);
            font-size: 14px;
            max-width: 800px;
            margin: 0 auto;
        }
        /* Glyph colors */
        .coral { color: #e07a5f; }
        .sage { color: #7ba38f; }
        .purple { color: #8b7ec8; }
        .blue { color: #5d9cec; }
        .gold { color: #d4a574; }
        .light { color: #f0f0f2; }
        .muted { color: rgba(240,240,242,0.4); }
        .border { color: rgba(240,240,242,0.06); }
    </style>
</head>
<body>
    <pre>%s</pre>
</body>
</html>`, strings.ReplaceAll(string(output), "\n", "<br>"))
        
        fmt.Fprint(w, html)
    })
    
    fmt.Println("✦ BonoboLab Visualization Server running at http://localhost:8080")
    fmt.Println("◈ Press Ctrl+C to stop")
    http.ListenAndServe(":8080", nil)
}
