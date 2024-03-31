package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"testing"
)

func BenchmarkGetDistricts(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = http.Get("http://localhost:3000/cities/3401/districts")
	}
}

func BenchmarkParse(b *testing.B) {
	b.ReportAllocs()
	var body = strings.NewReader(`{"district_id":340102}`)
	for i := 0; i < b.N; i++ {
		_, _ = http.Post("http://localhost:3000/parse", echo.MIMEApplicationJSONCharsetUTF8, body)
	}
}

func BenchmarkValidate(b *testing.B) {
	b.ReportAllocs()
	var body = strings.NewReader(`{"province_id":34,"city_id":3401,"district_id":340102}`)
	for i := 0; i < b.N; i++ {
		_, _ = http.Post("http://localhost:3000/validate", echo.MIMEApplicationJSONCharsetUTF8, body)
	}
}

func BenchmarkParallelValidate(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var body = strings.NewReader(`{"province_id":34,"city_id":3401,"district_id":340102}`)
			_, _ = http.Post("http://localhost:3000/validate", echo.MIMEApplicationJSONCharsetUTF8, body)
		}
	})
}

func BenchmarkParallelParse(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var body = strings.NewReader(`{"district_id":340102}`)
			_, _ = http.Post("http://localhost:3000/parse", echo.MIMEApplicationJSONCharsetUTF8, body)
		}
	})
}

func BenchmarkParallelGetDistricts(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = http.Get("http://localhost:3000/cities/3401/districts")
		}
	})
}
