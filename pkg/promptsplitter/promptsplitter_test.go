package promptsplitter

import (
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Length 5", args{5}, false},
		{"Length 10", args{10}, false},
		{"Length 100", args{100}, false},
		{"Length 0", args{0}, false},
		{"Length -5", args{-5}, true}, // this should now return an error
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GenerateRandomString(tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateRandomString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSplitPrompt(t *testing.T) {
	type args struct {
		text        string
		splitLength int
	}
	tests := []struct {
		name    string
		args    args
		want    []FileData
		wantErr bool
	}{
		{
			name: "case 1: text length is less than split length",
			args: args{
				text:        "Hello, world!",
				splitLength: 20,
			},
			want: []FileData{
				{
					Name:       "split_001_of_001.txt",
					Content:    "[START PART 1/1]\nHello, world!\n[END PART 1/1]\nALL PARTS SENT. Now you can continue processing the request.",
					PartNumber: 1,
					TotalParts: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "case 2: text length is greater than split length",
			args: args{
				text:        "Hello, world!",
				splitLength: 5,
			},
			want: []FileData{
				{
					Name:       "split_001_of_003.txt",
					Content:    "Do not answer yet. This is just another part of the text I want to send you. Just receive and acknowledge as \"Part 1/3 received\" and wait for the next part.\n[START PART 1/3]\nHello\n[END PART 1/3]\nRemember not answering yet. Just acknowledge you received this part with the message \"Part 1/3 received\" and wait for the next part.",
					PartNumber: 1,
					TotalParts: 3,
				},
				{
					Name:       "split_002_of_003.txt",
					Content:    "Do not answer yet. This is just another part of the text I want to send you. Just receive and acknowledge as \"Part 2/3 received\" and wait for the next part.\n[START PART 2/3]\n, wor\n[END PART 2/3]\nRemember not answering yet. Just acknowledge you received this part with the message \"Part 2/3 received\" and wait for the next part.",
					PartNumber: 2,
					TotalParts: 3,
				},
				{
					Name:       "split_003_of_003.txt",
					Content:    "[START PART 3/3]\nld!\n[END PART 3/3]\nALL PARTS SENT. Now you can continue processing the request.",
					PartNumber: 3,
					TotalParts: 3,
				},
			},
			wantErr: false,
		},
		{
			name: "case 3: split length is zero or less",
			args: args{
				text:        "Hello, world!",
				splitLength: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tt.wantErr == false {
						t.Errorf("SplitPrompt() panicked unexpectedly: %v", r)
					}
				}
			}()
			got := SplitPrompt(tt.args.text, tt.args.splitLength)
			if (len(got) == 0 || got[0].PartHash == "") && !tt.wantErr {
				t.Errorf("SplitPrompt() = %v, want non-empty PartHash", got)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1: a is less than b",
			args: args{
				a: 5,
				b: 10,
			},
			want: 5,
		},
		{
			name: "case 2: a is greater than b",
			args: args{
				a: 10,
				b: 5,
			},
			want: 5,
		},
		{
			name: "case 3: a equals b",
			args: args{
				a: 5,
				b: 5,
			},
			want: 5,
		},
		{
			name: "case 4: a and b are negative and a is less than b",
			args: args{
				a: -10,
				b: -5,
			},
			want: -10,
		},
		{
			name: "case 5: a and b are negative and a is greater than b",
			args: args{
				a: -5,
				b: -10,
			},
			want: -10,
		},
		{
			name: "case 6: a is negative and b is positive",
			args: args{
				a: -5,
				b: 10,
			},
			want: -5,
		},
		{
			name: "case 7: a is positive and b is negative",
			args: args{
				a: 10,
				b: -5,
			},
			want: -5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}
