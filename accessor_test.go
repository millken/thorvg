package thorvg_test

import (
	"testing"
	"unsafe"

	"github.com/millken/thorvg"
)

func TestAccessorBasic(t *testing.T) {
	// 创建 Accessor
	accessor := thorvg.NewAccessor()
	if accessor == nil {
		t.Fatal("Failed to create accessor")
	}
	defer accessor.Release()

	// 测试 GenerateID 方法
	id1 := accessor.GenerateID("test_object")
	id2 := accessor.GenerateID("test_object")
	id3 := accessor.GenerateID("different_object")

	// 相同名称应该生成相同ID
	if id1 != id2 {
		t.Errorf("GenerateID should return same ID for same name: got %d and %d", id1, id2)
	}

	// 不同名称应该生成不同ID
	if id1 == id3 {
		t.Errorf("GenerateID should return different IDs for different names: got %d and %d", id1, id3)
	}

	t.Logf("Generated ID for 'test_object': %d", id1)
	t.Logf("Generated ID for 'different_object': %d", id3)
}

func TestAccessorGenerateID(t *testing.T) {
	// 测试包级 GenerateID 函数
	id1 := thorvg.GenerateID("global_test")
	id2 := thorvg.GenerateID("global_test")
	id3 := thorvg.GenerateID("another_global_test")

	// 相同名称应该生成相同ID
	if id1 != id2 {
		t.Errorf("GenerateID should return same ID for same name: got %d and %d", id1, id2)
	}

	// 不同名称应该生成不同ID
	if id1 == id3 {
		t.Errorf("GenerateID should return different IDs for different names: got %d and %d", id1, id3)
	}
}

func TestAccessorSetNotImplemented(t *testing.T) {
	accessor := thorvg.NewAccessor()
	if accessor == nil {
		t.Fatal("Failed to create accessor")
	}
	defer accessor.Release()

	// 创建一个简单的 Paint 对象用于测试
	shape := thorvg.NewShape()
	if shape == nil {
		t.Fatal("Failed to create shape")
	}
	defer shape.Release()

	// 测试 Set 方法（应该返回错误，因为未实现）
	// 注意：由于 Shape 嵌入了 Paint，我们需要创建一个包装函数
	err := accessor.Set((*thorvg.Paint)(unsafe.Pointer(shape)), func(paint *thorvg.Paint) bool {
		return true
	})

	if err == nil {
		t.Error("Expected error for unimplemented Set method")
	}

	expectedMsg := "not implemented"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestAccessorErrors(t *testing.T) {
	// 测试空 accessor 对象
	var accessor *thorvg.Accessor

	// 测试 nil accessor 的错误处理
	if accessor != nil {
		t.Error("Expected accessor to be nil")
	}

	// 测试 nil accessor 的方法调用
	// 注意：在 Go 中，对 nil 指针调用方法会导致 panic，所以我们不测试这些
	t.Log("Nil accessor tests skipped - nil pointer dereference causes panic in Go")
}
