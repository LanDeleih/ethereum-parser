package provider_test

import (
	"testing"

	provider "github.com/landeleih/ethereum-parser/json-rpc"
)

func TestConvertHexToNumber(t *testing.T) {
	t.Parallel()

	expected := 20178487
	result, err := provider.ConvertHexToNumber("0x133e637")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if result != expected {
		t.Log("Wrong result", "got", result, "expected", expected)
		t.Fail()
	}
}

func TestConvertNumberToHex(t *testing.T) {
	t.Parallel()

	expected := "0x133e637"

	result := provider.ConvertNumberToHex(20178487)
	if result != expected {
		t.Log("Wrong result", "got", result, "expected", expected)
		t.Fail()
	}
}
