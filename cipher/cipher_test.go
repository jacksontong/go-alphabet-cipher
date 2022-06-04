package cipher

import (
	"testing"
)

func TestEncode(t *testing.T) {
	type args struct {
		keyword string
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "can encode 'meetmeontuesdayeveningatseven with a secret keyword 'vigilance'",
			args: args{
				keyword: "vigilance",
				message: "meetmeontuesdayeveningatseven",
			},
			want: "hmkbxebpxpmyllyrxiiqtoltfgzzv",
		},
		{
			name: "can encode 'meetmebythetree' with a secret keyword 'scones'",
			args: args{
				keyword: "scones",
				message: "meetmebythetree",
			},
			want: "egsgqwtahuiljgs",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.keyword, tt.args.message); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		keyword string
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "can decode 'meetmeontuesdayeveningatseven with a secret keyword 'vigilance'",
			args: args{
				keyword: "vigilance",
				message: "hmkbxebpxpmyllyrxiiqtoltfgzzv",
			},
			want: "meetmeontuesdayeveningatseven",
		},
		{
			name: "can decode 'meetmebythetree' with a secret keyword 'scones'",
			args: args{
				keyword: "scones",
				message: "egsgqwtahuiljgs",
			},
			want: "meetmebythetree",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.args.keyword, tt.args.message); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecipher(t *testing.T) {
	type args struct {
		cipher  string
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "can extract the secret keyword given an encrypted message 'opkyfipmfmwcvqoklyhxywgeecpvhelzg' and the original message 'thequickbrownfoxjumpsoveralazydog'",
			args: args{
				cipher:  "opkyfipmfmwcvqoklyhxywgeecpvhelzg",
				message: "thequickbrownfoxjumpsoveralazydog",
			},
			want: "vigilance",
		},
		{
			name: "can extract the secret keyword given an encrypted message 'hcqxqqtqljmlzhwiivgbsapaiwcenmyu' and the original message 'packmyboxwithfivedozenliquorjugs'",
			args: args{
				cipher:  "hcqxqqtqljmlzhwiivgbsapaiwcenmyu",
				message: "packmyboxwithfivedozenliquorjugs",
			},
			want: "scones",
		},
		{
			name: "can extract the secret keyword given an encrypted message 'hfnlphoontutufa' and the original message 'hellofromrussia'",
			args: args{
				cipher:  "hfnlphoontutufa",
				message: "hellofromrussia",
			},
			want: "abcabcx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decipher(tt.args.cipher, tt.args.message); got != tt.want {
				t.Errorf("Decipher() = %v, want %v", got, tt.want)
			}
		})
	}
}
