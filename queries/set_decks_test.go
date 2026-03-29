package queries

import (
	"context"
	"testing"
)

func TestSetDecks_MainBoardParsedAsArray(t *testing.T) {
	conn := setupSampleDB(t)
	ctx := context.Background()
	rows, err := conn.Execute(ctx, "SELECT mainBoard FROM set_decks WHERE code = $1", "A25_DECK1")
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	board, ok := rows[0]["mainBoard"].([]any)
	if !ok {
		t.Fatalf("expected mainBoard to be []any, got %T", rows[0]["mainBoard"])
	}
	if len(board) != 2 {
		t.Fatalf("expected 2 entries in mainBoard, got %d", len(board))
	}
	entry, ok := board[0].(map[string]any)
	if !ok {
		t.Fatalf("expected board entry to be map, got %T", board[0])
	}
	if entry["uuid"] != "card-uuid-001" {
		t.Errorf("expected uuid=card-uuid-001, got %v", entry["uuid"])
	}
}

func TestSetDecks_SideBoardParsedAsArray(t *testing.T) {
	conn := setupSampleDB(t)
	ctx := context.Background()
	rows, err := conn.Execute(ctx, "SELECT sideBoard FROM set_decks WHERE code = $1", "A25_DECK1")
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	board, ok := rows[0]["sideBoard"].([]any)
	if !ok {
		t.Fatalf("expected sideBoard to be []any, got %T", rows[0]["sideBoard"])
	}
	if len(board) != 1 {
		t.Fatalf("expected 1 entry in sideBoard, got %d", len(board))
	}
}

func TestSetDecks_TokensParsedAsArray(t *testing.T) {
	conn := setupSampleDB(t)
	ctx := context.Background()
	rows, err := conn.Execute(ctx, "SELECT tokens FROM set_decks WHERE code = $1", "A25_DECK1")
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	tokens, ok := rows[0]["tokens"].([]any)
	if !ok {
		t.Fatalf("expected tokens to be []any, got %T", rows[0]["tokens"])
	}
	entry, ok := tokens[0].(map[string]any)
	if !ok {
		t.Fatalf("expected token entry to be map, got %T", tokens[0])
	}
	if entry["uuid"] != "token-uuid-001" {
		t.Errorf("expected uuid=token-uuid-001, got %v", entry["uuid"])
	}
}

func TestSetDecks_SealedProductUuidsParsedAsArray(t *testing.T) {
	conn := setupSampleDB(t)
	ctx := context.Background()
	rows, err := conn.Execute(ctx, "SELECT sealedProductUuids FROM set_decks WHERE code = $1", "A25_DECK1")
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	uuids, ok := rows[0]["sealedProductUuids"].([]any)
	if !ok {
		t.Fatalf("expected sealedProductUuids to be []any, got %T", rows[0]["sealedProductUuids"])
	}
	if uuids[0] != "sealed-uuid-001" {
		t.Errorf("expected sealed-uuid-001, got %v", uuids[0])
	}
}

func TestSetDecks_SourceSetCodesParsedAsArray(t *testing.T) {
	conn := setupSampleDB(t)
	ctx := context.Background()
	rows, err := conn.Execute(ctx, "SELECT sourceSetCodes FROM set_decks WHERE code = $1", "A25_DECK1")
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	codes, ok := rows[0]["sourceSetCodes"].([]any)
	if !ok {
		t.Fatalf("expected sourceSetCodes to be []any, got %T", rows[0]["sourceSetCodes"])
	}
	if codes[0] != "A25" {
		t.Errorf("expected A25, got %v", codes[0])
	}
}

func TestSetDecks_CommanderParsedAsArray(t *testing.T) {
	conn := setupSampleDB(t)
	ctx := context.Background()
	rows, err := conn.Execute(ctx, "SELECT commander FROM set_decks WHERE code = $1", "A25_DECK1")
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	cmdr, ok := rows[0]["commander"].([]any)
	if !ok {
		t.Fatalf("expected commander to be []any, got %T", rows[0]["commander"])
	}
	if len(cmdr) != 0 {
		t.Errorf("expected empty commander list, got %d entries", len(cmdr))
	}
}

func TestSetDecks_FilterBySetCode(t *testing.T) {
	conn := setupSampleDB(t)
	ctx := context.Background()
	rows, err := conn.Execute(ctx, "SELECT * FROM set_decks WHERE setCode = $1", "MH2")
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	if rows[0]["name"] != "Modern Horizons 2 Theme Deck" {
		t.Errorf("expected 'Modern Horizons 2 Theme Deck', got %v", rows[0]["name"])
	}
}
