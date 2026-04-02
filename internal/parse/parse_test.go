package parse

import (
	"reflect"
	"testing"
)

func TestProblemIDs(t *testing.T) {
	got, err := ProblemIDs("1")
	if err != nil || !reflect.DeepEqual(got, []int{1}) {
		t.Fatalf("ProblemIDs(1) = %v, %v", got, err)
	}

	got, err = ProblemIDs("1, 2, 3")
	if err != nil || !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("ProblemIDs(1,2,3) = %v, %v", got, err)
	}

	got, err = ProblemIDs("2-4")
	if err != nil || !reflect.DeepEqual(got, []int{2, 3, 4}) {
		t.Fatalf("ProblemIDs(2-4) = %v, %v", got, err)
	}

	if _, err := ProblemIDs(""); err == nil {
		t.Fatal("expected error for empty")
	}
	if _, err := ProblemIDs("1-2-3"); err == nil {
		t.Fatal("expected error for bad range")
	}
	if _, err := ProblemIDs("5-3"); err == nil {
		t.Fatal("expected error for inverted range")
	}
}

func TestLanguages(t *testing.T) {
	if got := Languages(""); !reflect.DeepEqual(got, []string{"python3"}) {
		t.Fatalf("Languages(\"\") = %v", got)
	}
	if got := Languages("go, rust"); !reflect.DeepEqual(got, []string{"golang", "rust"}) {
		t.Fatalf("Languages(go,rust) = %v", got)
	}
}
