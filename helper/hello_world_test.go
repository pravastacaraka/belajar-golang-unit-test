package helper

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/**
 * Unit Testing
 *
 * Name: Test{function_name}
 * Example:
 * $ go test -v -run {function_name}
**/

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Raka")
	if result != "Hello Raka" {
		t.Error("Expected Hello Raka but got", result)
	}
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Raka")
	assert.Equal(t, "Hello Raka", result, "Result must be 'Hello Raka'")
	println("TestHelloWorld with Assert is done.")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Raka")
	require.Equal(t, "Hello Raka", result, "Result must be 'Hello Raka'")
	println("TestHelloWorld with Require is done.")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping test on Windows")
	}
	result := HelloWorld("Raka")
	require.Equal(t, "Hello Raka", result, "Result must be 'Hello Raka'")
}

func TestMain(m *testing.M) {
	// before
	println("BEFORE UNIT TEST")
	
	m.Run()

	// after
	println("AFTER UNIT TEST")
}

func TestSubTest(t *testing.T) {
	t.Run("World", func(t *testing.T) {
		result := HelloWorld("World")
		require.Equal(t, "Hello World", result, "Result must be 'Hello World'")
	})
	t.Run("Raka", func(t *testing.T) {
		result := HelloWorld("Raka")
		require.Equal(t, "Hello Raka", result, "Result must be 'Hello Raka'")
	})
}

func TestTableHelloWorld(t *testing.T) {
	var tests = []struct {
		name string
		request string
		expected string
	}{
		{"HelloWorld('Raka')", "Raka", "Hello Raka"},
		{"HelloWorld('World')", "World", "Hello World"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result, "Result must be '" + test.expected + "'")
		})
	}
}

/** 
 * Benchmarking
 * 
 * Mekanisme menghitung kecepatan performa kode program aplikasi
 * dilakukan dengan cara otomatis melakukan iterasi sampai waktu tertentu
 * di Go sudah diatur menggunakan testing.B bawaan dari package testing 
 *
 * Name: Benchmark{function_name}
 * Example:
 * $ go test -v -bench=.
 * $ go test -v -bench={function_name}
 **/

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Raka")
	}
}

func BenchmarkSubHelloWorld(b *testing.B) {
	b.Run("Raka", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Raka")
		}
	})
	b.Run("World", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("World")
		}
	})
}

func BenchmarkTableHelloWorld(b *testing.B) {
	var benchmarks = []struct {
		name string
		request string
	}{
		{"HelloWorld('Raka')", "Raka"},
		{"HelloWorld('World')", "World"},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
