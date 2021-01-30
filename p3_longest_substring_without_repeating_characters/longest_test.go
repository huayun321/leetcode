package p3_longest_substring_without_repeating_characters

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
	}{
		{"lohe", args{"hellohelloh"}, 4},
		{"abc", args{"abcabcbb"}, 3},
		{"wke", args{"pwwkew"}, 3},
		{"empty", args{""}, 0},
		{"blank", args{" "}, 1},
		{"vdf", args{"dvdf"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLen := lengthOfLongestSubstring(tt.args.s); gotLen != tt.wantLen {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", gotLen, tt.wantLen)
			}
		})
	}
}
