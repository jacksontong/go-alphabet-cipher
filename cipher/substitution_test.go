package cipher

import (
	"testing"
)

func Test_findSubstitution(t *testing.T) {
	type args struct {
		column rune
		row    rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		// column a
		{
			name: "column a row a",
			args: args{column: 'a', row: 'a'},
			want: 'a',
		},
		{
			name: "column a row b",
			args: args{column: 'a', row: 'b'},
			want: 'b',
		},
		{
			name: "column a row c",
			args: args{column: 'a', row: 'c'},
			want: 'c',
		},
		{
			name: "column a row d",
			args: args{column: 'a', row: 'd'},
			want: 'd',
		},
		{
			name: "column a row z",
			args: args{column: 'a', row: 'z'},
			want: 'z',
		},
		// column b
		{
			name: "column b row a",
			args: args{column: 'b', row: 'a'},
			want: 'b',
		},
		{
			name: "column b row b",
			args: args{column: 'b', row: 'b'},
			want: 'c',
		},
		{
			name: "column b row c",
			args: args{column: 'b', row: 'c'},
			want: 'd',
		},
		{
			name: "column b row d",
			args: args{column: 'b', row: 'd'},
			want: 'e',
		},
		{
			name: "column b row z",
			args: args{column: 'b', row: 'z'},
			want: 'a',
		},
		// column c
		{
			name: "column c row a",
			args: args{column: 'c', row: 'a'},
			want: 'c',
		},
		{
			name: "column c row b",
			args: args{column: 'c', row: 'b'},
			want: 'd',
		},
		{
			name: "column c row c",
			args: args{column: 'c', row: 'c'},
			want: 'e',
		},
		{
			name: "column c row d",
			args: args{column: 'c', row: 'd'},
			want: 'f',
		},
		{
			name: "column c row z",
			args: args{column: 'c', row: 'z'},
			want: 'b',
		},
		// column z
		{
			name: "column z row a",
			args: args{column: 'z', row: 'a'},
			want: 'z',
		},
		{
			name: "column z row b",
			args: args{column: 'z', row: 'b'},
			want: 'a',
		},
		{
			name: "column z row c",
			args: args{column: 'z', row: 'c'},
			want: 'b',
		},
		{
			name: "column z row d",
			args: args{column: 'z', row: 'd'},
			want: 'c',
		},
		{
			name: "column z row z",
			args: args{column: 'z', row: 'z'},
			want: 'y',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstitution(tt.args.column, tt.args.row); got != tt.want {
				t.Errorf("findSubstitution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repeatedSecret(t *testing.T) {
	type args struct {
		secret  string
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "secret vigilance message meetmeontuesdayeveningatseven",
			args: args{
				secret:  "vigilance",
				message: "meetmeontuesdayeveningatseven",
			},
			want: "vigilancevigilancevigilancevi",
		},
		{
			name: "secret scones message meetmebythetree",
			args: args{
				secret:  "scones",
				message: "meetmebythetree",
			},
			want: "sconessconessco",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSecret(tt.args.secret, tt.args.message); got != tt.want {
				t.Errorf("repeatedSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_revertSubstitution(t *testing.T) {
	type args struct {
		substitution rune
		column       rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			name: "substitution e column s",
			args: args{
				substitution: 'e',
				column:       's',
			},
			want: 'm',
		},
		{
			name: "substitution g column c",
			args: args{
				substitution: 'g',
				column:       'c',
			},
			want: 'e',
		},
		{
			name: "substitution s column o",
			args: args{
				substitution: 's',
				column:       'o',
			},
			want: 'e',
		},
		{
			name: "substitution g column n",
			args: args{
				substitution: 'g',
				column:       'n',
			},
			want: 't',
		},
		{
			name: "substitution q column e",
			args: args{
				substitution: 'q',
				column:       'e',
			},
			want: 'm',
		},
		{
			name: "substitution y column z",
			args: args{
				substitution: 'y',
				column:       'z',
			},
			want: 'z',
		},
		{
			name: "substitution a column a",
			args: args{
				substitution: 'a',
				column:       'a',
			},
			want: 'a',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := revertSubstitution(tt.args.substitution, tt.args.column); got != tt.want {
				t.Errorf("revertSubstitution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRepeatedSecret(t *testing.T) {
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
			name: "scones",
			args: args{
				cipher:  "egsgqwtahuiljgs",
				message: "meetmebythetree",
			},
			want: "sconessconessco",
		},
		{
			name: "vigilance",
			args: args{
				cipher:  "hmkbxebpxpmyllyrxiiqtoltfgzzv",
				message: "meetmeontuesdayeveningatseven",
			},
			want: "vigilancevigilancevigilancevi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatedSecret(tt.args.cipher, tt.args.message); got != tt.want {
				t.Errorf("findRepeatedSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}
