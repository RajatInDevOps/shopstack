package main

import "testing"

func TestANewStoreIsSeeded(t *testing.T) {
	if got := NewStore().Size(); got != 3 {
		t.Errorf("Size() = %d, want 3", got)
	}
}

func TestReserveTakesStockOut(t *testing.T) {
	store := NewStore()

	item, ok := store.Reserve("SKU-1001", 5)
	if !ok {
		t.Fatal("a reservation within stock was refused")
	}
	if item.Quantity != 20 {
		t.Errorf("Quantity = %d, want 20", item.Quantity)
	}
}

func TestReserveRefusesMoreThanThereIs(t *testing.T) {
	store := NewStore()

	if _, ok := store.Reserve("SKU-1003", 1); ok {
		t.Error("reserved an item with no stock")
	}
	if _, ok := store.Reserve("SKU-1002", 99); ok {
		t.Error("reserved more than the warehouse holds")
	}
}

func TestReserveRefusesANonsenseAmount(t *testing.T) {
	store := NewStore()

	for _, amount := range []int{0, -1, -100} {
		if _, ok := store.Reserve("SKU-1001", amount); ok {
			t.Errorf("Reserve(%d) was allowed", amount)
		}
	}
}

func TestAnUnknownSKUIsNotFound(t *testing.T) {
	store := NewStore()

	if _, ok := store.Find("NOPE"); ok {
		t.Error("found an sku that was never added")
	}
	if _, ok := store.Reserve("NOPE", 1); ok {
		t.Error("reserved against an sku that was never added")
	}
}

func TestAddingAnExistingSKUTopsItUp(t *testing.T) {
	store := NewStore()
	store.Add("SKU-1003", "Doohickey", 7)

	item, _ := store.Find("SKU-1003")
	if item.Quantity != 7 {
		t.Errorf("Quantity = %d, want 7", item.Quantity)
	}
	if store.Size() != 3 {
		t.Errorf("a duplicate sku created a second row: Size() = %d", store.Size())
	}
}

func TestLowStockUsesTheThreshold(t *testing.T) {
	store := NewStore()

	if got := len(store.LowStock(1)); got != 1 {
		t.Errorf("LowStock(1) = %d items, want 1", got)
	}
	if got := len(store.LowStock(10)); got != 2 {
		t.Errorf("LowStock(10) = %d items, want 2", got)
	}
}

func TestAllIsStableInInsertionOrder(t *testing.T) {
	items := NewStore().All()

	want := []string{"SKU-1001", "SKU-1002", "SKU-1003"}
	for i, sku := range want {
		if items[i].SKU != sku {
			t.Errorf("item %d = %s, want %s", i, items[i].SKU, sku)
		}
	}
}
