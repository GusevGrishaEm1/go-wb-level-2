package main

import "testing"

func TestUncompressString(t *testing.T) {
	tests := []struct {
		name string
		errorMsg string
		input string
		output string
	} {
		{
            name: "test#1",
            errorMsg: "",
            input: "a4bc2d5e",
            output: "aaaabccddddde",
        },
		{
            name: "test#2",
            errorMsg: "",
            input: `qwe\4\5`,
            output: "qwe45",
        },
		{
            name: "test#3",
            errorMsg: "некорректная строка",
            input: "45",
            output: "",
        },
		{
            name: "test#4",
            errorMsg: "",
            input: "",
            output: "",
        },
	}
	for _, tt := range tests {
		t.Run(tt.name , func(t *testing.T) {
			result, err := uncompressString(tt.input)
			if err!= nil {
				if tt.errorMsg != "" && err.Error() == tt.errorMsg {
					return
				}
				t.Error(err)
			}
			if result != tt.output {
                t.Errorf("uncompressString() = %v, want %v", result, tt.output)
            }
		})
	}
}