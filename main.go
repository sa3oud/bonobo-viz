package main

import (
    "fmt"
    "strings"
)

// ANSI color codes
var (
    coral   = "\033[38;5;209m"
    sage    = "\033[38;5;108m"
    purple  = "\033[38;5;140m"
    blue    = "\033[38;5;75m"
    gold    = "\033[38;5;180m"
    light   = "\033[38;5;255m"
    muted   = "\033[38;5;244m"
    subtle  = "\033[38;5;236m"
    border  = "\033[38;5;235m"
    reset   = "\033[0m"
)

// Glyphs from your design system
var glyphs = map[string]string{
    "signal":    "◈",
    "hex":       "⬡",
    "triangle":  "△",
    "link":      "⟁",
    "chat":      "◇",
    "dashboard": "⬢",
    "discover":  "◎",
    "profile":   "◉",
    "skills":    "◈",
    "showcases": "⬡",
    "vouches":   "◆",
    "network":   "⧉",
    "match":     "◈",
    "pending":   "◇",
    "completed": "◆",
    "key":       "⚷",
    "star":      "✦",
    "lock":      "◧",
    "verified":  "◉",
    "featured":  "✧",
    "premium":   "◈",
}

type Connection struct {
    ID       string
    Type     string
    User1    string
    User2    string
    Strength float64
    Verified bool
    Featured bool
    TimeAgo  string
}

type Community struct {
    Name        string
    Connections []Connection
    Members     []string
}

func main() {
    // Clear screen
    fmt.Print("\033[2J\033[H")
    
    community := Community{
        Name:    "BonoboLab Creative Collective",
        Members: []string{"Alex", "Jordan", "Casey", "Riley", "Morgan", "Taylor", "Quinn", "Sage"},
        Connections: []Connection{
            {"1", "romantic", "Alex", "Jordan", 0.95, true, true, "2h ago"},
            {"2", "collab", "Jordan", "Casey", 0.88, true, false, "1d ago"},
            {"3", "tandem", "Casey", "Riley", 0.92, false, true, "3h ago"},
            {"4", "both", "Riley", "Morgan", 0.78, true, false, "5h ago"},
            {"5", "collab", "Morgan", "Taylor", 0.85, false, false, "1d ago"},
            {"6", "romantic", "Taylor", "Quinn", 0.71, true, true, "2d ago"},
            {"7", "tandem", "Quinn", "Sage", 0.94, true, false, "4h ago"},
            {"8", "collab", "Sage", "Alex", 0.63, false, false, "30m ago"},
            {"9", "both", "Jordan", "Riley", 0.89, true, true, "6h ago"},
            {"10", "collab", "Casey", "Morgan", 0.77, false, false, "12h ago"},
        },
    }
    
    displayDashboard(community)
    fmt.Print(reset)
}

func displayDashboard(c Community) {
    // Header
    fmt.Println(border + strings.Repeat("▀", 64) + reset)
    fmt.Println(coral + centerText(" ✦  B O N O B O L A B  C O M M U N I T Y  ✦ ", 64) + reset)
    fmt.Println(border + strings.Repeat("▄", 64) + reset)
    
    // Community Stats
    fmt.Println()
    fmt.Printf("  %s%s%s %s%s%s\n", 
        gold, glyphs["dashboard"], reset,
        light, "COMMUNITY DASHBOARD", reset)
    fmt.Println(border + strings.Repeat("─", 40) + reset)
    
    fmt.Printf("  %s%s%s %-12s %s%s%s\n",
        coral, glyphs["profile"], reset,
        "Members:",
        light, fmt.Sprintf("%d", len(c.Members)), reset)
    
    fmt.Printf("  %s%s%s %-12s %s%s%s\n",
        sage, glyphs["network"], reset,
        "Connections:",
        light, fmt.Sprintf("%d", len(c.Connections)), reset)
    
    fmt.Printf("  %s%s%s %-12s %s%s%s\n",
        purple, glyphs["star"], reset,
        "Verified:",
        light, fmt.Sprintf("%d", countVerified(c.Connections)), reset)
    
    // Connection Types
    fmt.Println()
    fmt.Printf("  %s%s%s %s\n", gold, glyphs["link"], reset, "CONNECTION TYPES")
    fmt.Println(border + strings.Repeat("─", 40) + reset)
    
    types := map[string]int{"romantic": 0, "collab": 0, "tandem": 0, "both": 0}
    for _, conn := range c.Connections {
        types[conn.Type]++
    }
    
    // Romantic
    fmt.Printf("  %s%s%s %-8s %s%s %s %s\n",
        coral, glyphs["signal"], reset,
        "Romantic",
        light, strings.Repeat(glyphs["signal"], types["romantic"]),
        fmt.Sprintf("(%d)", types["romantic"]),
        reset)
    
    // Collaboration
    fmt.Printf("  %s%s%s %-8s %s%s %s %s\n",
        purple, glyphs["hex"], reset,
        "Collab",
        light, strings.Repeat(glyphs["hex"], types["collab"]),
        fmt.Sprintf("(%d)", types["collab"]),
        reset)
    
    // Tandem
    fmt.Printf("  %s%s%s %-8s %s%s %s %s\n",
        sage, glyphs["triangle"], reset,
        "Tandem",
        light, strings.Repeat(glyphs["triangle"], types["tandem"]),
        fmt.Sprintf("(%d)", types["tandem"]),
        reset)
    
    // Both
    fmt.Printf("  %s%s%s %-8s %s%s %s %s\n",
        blue, glyphs["link"], reset,
        "Multi",
        light, strings.Repeat(glyphs["link"], types["both"]),
        fmt.Sprintf("(%d)", types["both"]),
        reset)
    
    // Recent Activity
    fmt.Println()
    fmt.Printf("  %s%s%s %s\n", gold, glyphs["chat"], reset, "RECENT ACTIVITY")
    fmt.Println(border + strings.Repeat("─", 40) + reset)
    
    for i, conn := range c.Connections[:5] {
        // Get color and glyph based on type
        color := muted
        glyph := glyphs["pending"]
        
        switch conn.Type {
        case "romantic":
            color = coral
            glyph = glyphs["signal"]
        case "collab":
            color = purple
            glyph = glyphs["hex"]
        case "tandem":
            color = sage
            glyph = glyphs["triangle"]
        case "both":
            color = blue
            glyph = glyphs["link"]
        }
        
        if conn.Verified {
            glyph = glyphs["verified"]
            color = gold
        }
        
        fmt.Printf("  %s•%s %s%s%s %s%-8s%s → %s%-8s%s %s(%s)%s\n",
            muted, reset,
            color, glyph, reset,
            light, conn.User1, reset,
            light, conn.User2, reset,
            muted, conn.TimeAgo, reset)
        
        if i < 4 {
            fmt.Printf("  %s|%s\n", muted, reset)
        }
    }
    
    // Footer
    fmt.Println()
    fmt.Println(border + strings.Repeat("▀", 64) + reset)
    fmt.Printf("%s%s%s %s %s %s %s %s %s\n",
        muted, strings.Repeat(" ", 20),
        gold, glyphs["star"],
        light, "BonoboLab",
        gold, glyphs["verified"],
        reset)
    fmt.Println(border + strings.Repeat("▄", 64) + reset)
}

func centerText(text string, width int) string {
    if len(text) >= width {
        return text
    }
    padding := width - len(text)
    left := padding / 2
    right := padding - left
    return strings.Repeat(" ", left) + text + strings.Repeat(" ", right)
}

func countVerified(connections []Connection) int {
    count := 0
    for _, conn := range connections {
        if conn.Verified {
            count++
        }
    }
    return count
}
