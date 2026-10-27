package main

import "sync"

// Item is one line in the warehouse.
type Item struct {
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

// Store holds the warehouse. It is safe for concurrent use.
type Store struct {
	mu    sync.RWMutex
	items map[string]*Item
	order []string
}

func NewStore() *Store {
	store := &Store{items: map[string]*Item{}}
	store.Add("SKU-1001", "Widget", 25)
	store.Add("SKU-1002", "Gadget", 8)
	store.Add("SKU-1003", "Doohickey", 0)
	return store
}

func (s *Store) Add(sku, name string, quantity int) *Item {
	s.mu.Lock()
	defer s.mu.Unlock()

	if quantity < 0 {
		quantity = 0
	}
	if existing, ok := s.items[sku]; ok {
		existing.Quantity += quantity
		return existing
	}
	item := &Item{SKU: sku, Name: name, Quantity: quantity}
	s.items[sku] = item
	s.order = append(s.order, sku)
	return item
}

func (s *Store) All() []Item {
	s.mu.RLock()
	defer s.mu.RUnlock()

	found := make([]Item, 0, len(s.order))
	for _, sku := range s.order {
		found = append(found, *s.items[sku])
	}
	return found
}

func (s *Store) LowStock(threshold int) []Item {
	var found []Item
	for _, item := range s.All() {
		if item.Quantity < threshold {
			found = append(found, item)
		}
	}
	if found == nil {
		found = []Item{}
	}
	return found
}

func (s *Store) Find(sku string) (Item, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item, ok := s.items[sku]
	if !ok {
		return Item{}, false
	}
	return *item, true
}

// Reserve takes stock out, reporting whether it could.
func (s *Store) Reserve(sku string, amount int) (Item, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if amount < 1 {
		return Item{}, false
	}
	item, ok := s.items[sku]
	if !ok || item.Quantity < amount {
		return Item{}, false
	}
	item.Quantity -= amount
	return *item, true
}

func (s *Store) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.items)
}
