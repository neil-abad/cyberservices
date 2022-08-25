package strmatcher_test

import (
	"strconv"
	"testing"

	"cyberservices.com/core/common"
	. "cyberservices.com/core/common/strmatcher"
)

func BenchmarkDomainMatcherGroup(b *testing.B) {
	g := new(DomainMatcherGroup)

	for i := 1; i <= 1024; i++ {
		g.Add(strconv.Itoa(i)+".cyberservices.com", uint32(i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = g.Match("0.cyberservices.com")
	}
}

func BenchmarkFullMatcherGroup(b *testing.B) {
	g := new(FullMatcherGroup)

	for i := 1; i <= 1024; i++ {
		g.Add(strconv.Itoa(i)+".cyberservices.com", uint32(i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = g.Match("0.cyberservices.com")
	}
}

func BenchmarkMarchGroup(b *testing.B) {
	g := new(MatcherGroup)
	for i := 1; i <= 1024; i++ {
		m, err := Domain.New(strconv.Itoa(i) + ".cyberservices.com")
		common.Must(err)
		g.Add(m)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = g.Match("0.cyberservices.com")
	}
}
