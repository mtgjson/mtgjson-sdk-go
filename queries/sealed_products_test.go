package queries

import (
	"context"
	"testing"
)

func TestSealedProducts_ContentsParsedAsObject(t *testing.T) {
	conn := setupSampleDB(t)
	ctx := context.Background()
	rows, err := conn.Execute(ctx, "SELECT contents FROM sealed_products WHERE uuid = $1", "sealed-uuid-001")
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	contents, ok := rows[0]["contents"].(map[string]any)
	if !ok {
		t.Fatalf("expected contents to be map, got %T", rows[0]["contents"])
	}
	if _, ok := contents["pack"]; !ok {
		t.Error("expected contents to have 'pack' key")
	}
}

func TestSealedProducts_IdentifiersParsedAsObject(t *testing.T) {
	conn := setupSampleDB(t)
	ctx := context.Background()
	rows, err := conn.Execute(ctx, "SELECT identifiers FROM sealed_products WHERE uuid = $1", "sealed-uuid-001")
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	ids, ok := rows[0]["identifiers"].(map[string]any)
	if !ok {
		t.Fatalf("expected identifiers to be map, got %T", rows[0]["identifiers"])
	}
	if ids["tcgplayerProductId"] != "162583" {
		t.Errorf("expected tcgplayerProductId=162583, got %v", ids["tcgplayerProductId"])
	}
}

func TestSealedProducts_PurchaseUrlsParsedAsObject(t *testing.T) {
	conn := setupSampleDB(t)
	ctx := context.Background()
	rows, err := conn.Execute(ctx, "SELECT purchaseUrls FROM sealed_products WHERE uuid = $1", "sealed-uuid-001")
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	urls, ok := rows[0]["purchaseUrls"].(map[string]any)
	if !ok {
		t.Fatalf("expected purchaseUrls to be map, got %T", rows[0]["purchaseUrls"])
	}
	if _, ok := urls["tcgplayer"]; !ok {
		t.Error("expected purchaseUrls to have 'tcgplayer' key")
	}
}

func TestSealedProducts_FilterBySetCode(t *testing.T) {
	conn := setupSampleDB(t)
	ctx := context.Background()
	rows, err := conn.Execute(ctx, "SELECT * FROM sealed_products WHERE setCode = $1", "A25")
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(rows))
	}
}

func TestSealedProducts_FilterByCategory(t *testing.T) {
	conn := setupSampleDB(t)
	ctx := context.Background()
	rows, err := conn.Execute(ctx, "SELECT * FROM sealed_products WHERE category = $1", "booster_box")
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(rows))
	}
}
