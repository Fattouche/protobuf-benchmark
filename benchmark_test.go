package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/protobuf-benchmark/pb"
)

var (
	small = pb.Comp{
		Str:  "small",
		I:    10,
		F:    10,
		B:    true,
		Byte: []byte(`This is a small test`),
	}

	medium = pb.Comp{
		Str:  "xj3jJd9A8sK31D5R25UWy8OzMRI3Ok022aE8W1dmRKycZHe2zf7bzU4Qvfd",
		I:    10,
		F:    10,
		B:    true,
		Byte: []byte(`3U6CsMB4D8yPlH3cje0KHEX7QyZaFSbfuRMDzEZaPmFjwLiXamAXee2YIiBX3UaWBikJWAmUGaj87dqTUSps1kcwOAbWpAaWoJzAfTCtrsGErq69cCarneVAajfyAkYlZvXGRLIifqxRZnrOjfW5oAj7mwBkCYXo43i6KnRll3iTNtUSwKMYwK3qdG04LPIjvGIzKapB`),
	}

	large = pb.Comp{
		Str: "xj3jJd9A8sK31D5R25UWy8OzMRI3Ok022aE8W1dmRKycZHe2zf7bzU4QvfdAlFQcDgXYHIG42JldotgmVp6uIyJMtMqmJ1PzQyEgvTNUUWy3HjL3eTRh78rxuUmCXB2XXpt1CEl9VJpFRshGSkN3pZ",
		I:   1051251251,
		F:   10888512.2,
		B:   true,
		Byte: []byte(`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`),
	}
)

func BenchmarkMarshalJSON(b *testing.B) {
	s, m, l := small, medium, large

	b.ResetTimer()

	b.Run("Small", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			json.Marshal(&s)
		}
	})
	b.Run("Medium", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			json.Marshal(&m)
		}
	})
	b.Run("Large", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			json.Marshal(&l)
		}
	})
	fmt.Print("\n")
}

func BenchmarkMarshalProto(b *testing.B) {
	s, m, l := small, medium, large

	b.ResetTimer()

	b.Run("Small", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			proto.Marshal(&s)
		}
	})
	b.Run("Medium", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			proto.Marshal(&m)
		}
	})
	b.Run("Large", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			proto.Marshal(&l)
		}
	})
	fmt.Print("\n")
}

func BenchmarkUnMarshalJSON(b *testing.B) {
	s, m, l := small, medium, large
	sb, _ := json.Marshal(s)
	mb, _ := json.Marshal(m)
	lb, _ := json.Marshal(l)

	var smallShell, mediumShell, largeShell pb.Comp

	b.ResetTimer()

	b.Run("Small", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			json.Unmarshal(sb, &smallShell)
		}
	})
	b.Run("Medium", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			json.Unmarshal(mb, &mediumShell)
		}
	})
	b.Run("Large", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			json.Unmarshal(lb, &largeShell)
		}
	})
	fmt.Print("\n")
}

func BenchmarkUnMarshalProto(b *testing.B) {
	s, m, l := small, medium, large
	sb, _ := json.Marshal(s)
	mb, _ := json.Marshal(m)
	lb, _ := json.Marshal(l)

	var smallShell, mediumShell, largeShell pb.Comp

	b.ResetTimer()

	b.Run("Small", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			proto.Unmarshal(sb, &smallShell)
		}
	})
	b.Run("Medium", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			proto.Unmarshal(mb, &mediumShell)
		}
	})
	b.Run("Large", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			proto.Unmarshal(lb, &largeShell)
		}
	})
	fmt.Print("\n")
}
