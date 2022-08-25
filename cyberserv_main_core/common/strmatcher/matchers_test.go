package strmatcher_test

import (
	"testing"

	"cyberservices.com/core/common"
	. "cyberservices.com/core/common/strmatcher"
)

func TestMatcher(t *testing.T) {
	cases := []struct {
		pattern string
		mType   Type
		input   string
		output  bool
	}{
		{
			pattern: "cyberservices.com",
			mType:   Domain,
			input:   "www.cyberservices.com",
			output:  true,
		},
		{
			pattern: "cyberservices.com",
			mType:   Domain,
			input:   "cyberservices.com",
			output:  true,
		},
		{
			pattern: "cyberservices.com",
			mType:   Domain,
			input:   "www.v3ray.com",
			output:  false,
		},
		{
			pattern: "cyberservices.com",
			mType:   Domain,
			input:   "2ray.com",
			output:  false,
		},
		{
			pattern: "cyberservices.com",
			mType:   Domain,
			input:   "xcyberservices.com",
			output:  false,
		},
		{
			pattern: "cyberservices.com",
			mType:   Full,
			input:   "cyberservices.com",
			output:  true,
		},
		{
			pattern: "cyberservices.com",
			mType:   Full,
			input:   "xcyberservices.com",
			output:  false,
		},
		{
			pattern: "cyberservices.com",
			mType:   Regex,
			input:   "Project CSxcom",
			output:  true,
		},
	}
	for _, test := range cases {
		matcher, err := test.mType.New(test.pattern)
		common.Must(err)
		if m := matcher.Match(test.input); m != test.output {
			t.Error("unexpected output: ", m, " for test case ", test)
		}
	}
}
