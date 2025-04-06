package inmemoryfs

import (
	"testing"
)

func TestInMemoryFS_Create(t *testing.T) {
	ifs := New()
	_ = ifs.Create("fruits/apple")
	t.Run("Create 2 nested directory's", func(t *testing.T) {

		var childMap = ifs.GetChildren().(map[string]*Directory)
		if _, exists := childMap["fruits"]; !exists {
			t.Errorf("Create() error. Expected to find fruits ")
		}
		if _, exists := childMap["fruits"].children["apple"]; !exists {
			t.Errorf("Create() error. Expected to find apple ")
		}
	})

	_ = ifs.Create("food")
	t.Run("Create 1 child directory", func(t *testing.T) {

		var childMap = ifs.GetChildren().(map[string]*Directory)
		if _, exists := childMap["food"]; !exists {
			t.Errorf("Create() error. Expected to find food ")
		}
	})
}

func TestInMemoryFS_Delete(t *testing.T) {

	t.Run("Delete nested child directory", func(t *testing.T) {

		ifs := New()
		_ = ifs.Create("fruits/apple")
		_ = ifs.Delete("fruits/apple")

		var childMap = ifs.GetChildren().(map[string]*Directory)
		if _, exists := childMap["fruits"]; !exists {
			t.Errorf("Create() error. Expected to find fruits ")
		}
		if _, exists := childMap["fruits"].children["apple"]; exists {
			t.Errorf("Delete() error. Expected to NOT find apple ")
		}
	})
}

func TestInMemoryFS_Move(t *testing.T) {
	t.Run("Move", func(t *testing.T) {

		ifs := New()
		_ = ifs.Create("fruits/apple")
		_ = ifs.Create("food")
		_ = ifs.Move("fruits", "food")

		var childMap = ifs.GetChildren().(map[string]*Directory)
		if _, exists := childMap["food"]; !exists {
			t.Errorf("Move() error. Expected to find food ")
		}
		if _, exists := childMap["food"].children["fruits"]; !exists {
			t.Errorf("Move() error. Expected to find fruits ")
		}
		if _, exists := childMap["food"].children["fruits"].children["apple"]; !exists {
			t.Errorf("Move() error. Expected to find apple ")
		}
	})
}
