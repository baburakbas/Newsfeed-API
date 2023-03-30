package newsfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	if len(feed.Items) != 1 {
		t.Errorf("No items were added")
	}
}
func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	result := feed.GetAll()

	if len(result) != 1 {
		t.Errorf("No items were added")
	}
}
