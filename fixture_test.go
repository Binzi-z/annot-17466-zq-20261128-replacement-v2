package fixture

import (
    "encoding/json"
    "os"
    "testing"
)

func TestLocalSeed(t *testing.T) {
    raw, err := os.ReadFile("fixture.json")
    if err != nil { t.Fatal(err) }
    var data struct {
        Symbols []string
        GapBudget int
        Capacity int
        Remove int
        TicketSlots [][]int
        Windows []struct { Physical []int; Mode string }
        ExpectedOverlap []int
    }
    if err = json.Unmarshal(raw, &data); err != nil { t.Fatal(err) }
    switch {
    case len(data.Symbols) > 0:
        if data.GapBudget != 1 || data.Symbols[1] != "literal:open" || data.Symbols[3] != "gap" { t.Fatal(data) }
    case data.Capacity > 0:
        if data.Capacity != 2 || data.Remove != 1 || data.TicketSlots[0][0] != data.TicketSlots[1][0] { t.Fatal(data) }
    default:
        x, y := data.Windows[0], data.Windows[1]
        if x.Mode != "write" || y.Mode != "execute" || y.Physical[0] != data.ExpectedOverlap[0] ||
           x.Physical[1] != data.ExpectedOverlap[1] || x.Physical[1] <= y.Physical[0] { t.Fatal(data) }
    }
    if prompt, err := os.ReadFile("task/PROMPT.txt"); err != nil || len(prompt)==0 { t.Fatal(err) }
}
