package redshift

import (
	"errors"
	"strings"
	"testing"
)

func verifyParser(q string, t *testing.T) error {
	q = strings.TrimSpace(q)
	q = strings.Replace(q, "json", "myjson", -1)

	ok := ParseCopy(q)
	if !ok {
		return errors.New("...")
	}
	return nil
}

func TestCopyFromParser_1(t *testing.T) {
	t.Logf("query: %v\n\n", q1)
	err := verifyParser(q1, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_2(t *testing.T) {
	t.Logf("query: %v\n\n", q2)
	err := verifyParser(q2, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_3(t *testing.T) {
	t.Logf("query: %v\n\n", q3)
	err := verifyParser(q3, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_4(t *testing.T) {
	t.Logf("query: %v\n\n", q4)
	err := verifyParser(q4, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_5(t *testing.T) {
	t.Logf("query: %v\n\n", q5)
	err := verifyParser(q5, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_6(t *testing.T) {
	t.Logf("query: %v\n\n", q6)
	err := verifyParser(q6, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_7(t *testing.T) {
	t.Logf("query: %v\n\n", q7)
	err := verifyParser(q7, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_8(t *testing.T) {
	t.Logf("query: %v\n\n", q8)
	err := verifyParser(q8, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_9(t *testing.T) {
	t.Logf("query: %v\n\n", q9)
	err := verifyParser(q9, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_10(t *testing.T) {
	t.Logf("query: %v\n\n", q10)
	err := verifyParser(q10, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_11(t *testing.T) {
	t.Logf("query: %v\n\n", q11)
	err := verifyParser(q11, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_12(t *testing.T) {
	t.Logf("query: %v\n\n", q12)
	err := verifyParser(q12, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_13(t *testing.T) {
	t.Logf("query: %v\n\n", q13)
	err := verifyParser(q13, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_14(t *testing.T) {
	t.Logf("query: %v\n\n", q14)
	err := verifyParser(q14, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_15(t *testing.T) {
	t.Logf("query: %v\n\n", q15)
	err := verifyParser(q15, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_16(t *testing.T) {
	t.Logf("query: %v\n\n", q16)
	err := verifyParser(q16, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_17(t *testing.T) {
	t.Logf("query: %v\n\n", q17)
	err := verifyParser(q17, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_18(t *testing.T) {
	t.Logf("query: %v\n\n", q18)
	err := verifyParser(q18, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_19(t *testing.T) {
	t.Logf("query: %v\n\n", q19)
	err := verifyParser(q19, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_20(t *testing.T) {
	t.Logf("query: %v\n\n", q20)
	err := verifyParser(q20, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_21(t *testing.T) {
	t.Logf("query: %v\n\n", q21)
	err := verifyParser(q21, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_22(t *testing.T) {
	t.Logf("query: %v\n\n", q22)
	err := verifyParser(q22, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_23(t *testing.T) {
	t.Logf("query: %v\n\n", q23)
	err := verifyParser(q23, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_24(t *testing.T) {
	t.Logf("query: %v\n\n", q24)
	err := verifyParser(q24, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_25(t *testing.T) {
	t.Logf("query: %v\n\n", q25)
	err := verifyParser(q25, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_26(t *testing.T) {
	t.Logf("query: %v\n\n", q26)
	err := verifyParser(q26, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_27(t *testing.T) {
	t.Logf("query: %v\n\n", q27)
	err := verifyParser(q27, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

func TestCopyFromParser_28(t *testing.T) {
	t.Logf("query: %v\n\n", q28)
	err := verifyParser(q28, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}
