package redshift

import (
	"errors"
	"strings"
	"testing"

	"github.com/araddon/qlbridge/lex"
)

// func tv(t lex.TokenType, v string) lex.Token {
// 	return lex.Token{T: t, V: v}
// }

// func verifyTokens(t *testing.T, sql string, tokens []lex.Token) {
// 	l := lex.NewLexer(sql, InfluxQlDialect)
// 	for _, goodToken := range tokens {
// 		tok := l.NextToken()
// 		u.Debugf("%#v  %#v", tok, goodToken)
// 		assert.Equal(t, tok.V, goodToken.V, "has='%v' want='%v'", tok.V, goodToken.V)
// 		assert.Equal(t, tok.T, goodToken.T, "has='%v' want='%v'", tok.V, goodToken.V)
// 	}
// }

func verifyLex(q string, t *testing.T) error {
	q = strings.TrimSpace(q)
	q = strings.Replace(q, "json", "myjson", -1)

	l := NewRedshiftLexer(q)
	tok := l.NextToken()
	for tok.T != lex.TokenEOF {
		t.Logf("got:%v  \n", tok)
		if tok.T == lex.TokenError {
			t.Logf("Token ERROR: %s", q)
			return errors.New(tok.V)
		}
		tok = l.NextToken()
	}
	return nil
}

var q1 = `copy listing
	from 's3://mybucket/data/listing/' 
	iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole';`

func TestCopyFromLex_1(t *testing.T) {
	t.Logf("query: %v\n\n", q1)
	err := verifyLex(q1, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q2 = `copy favoritemovies from 'dynamodb://Movies'
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole' 
readratio 50;`

func TestCopyFromLex_2(t *testing.T) {
	t.Logf("query: %v\n\n", q2)
	err := verifyLex(q2, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q3 = `copy sales
from 'emr://j-SAMPLE2B500FC/myoutput/part-*' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
delimiter '\t' lzop;`

func TestCopyFromLex_3(t *testing.T) {
	t.Logf("query: %v\n\n", q3)
	err := verifyLex(q3, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q4 = `copy sales
from 'emr://j-SAMPLE2B500FC/myoutput/json/' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
json 's3://mybucket/jsonpaths.txt';`

func TestCopyFromLex_4(t *testing.T) {
	t.Logf("query: %v\n\n", q4)
	err := verifyLex(q4, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q5 = `copy category
from 's3://mybucket/custdata' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole';`

func TestCopyFromLex_5(t *testing.T) {
	t.Logf("query: %v\n\n", q5)
	err := verifyLex(q5, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q6 = `copy customer
from 's3://mybucket/cust.manifest' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
ssh;`

func TestCopyFromLex_6(t *testing.T) {
	t.Logf("query: %v\n\n", q6)
	err := verifyLex(q6, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q7 = `copy customer
from 's3://mybucket/cust.manifest' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
manifest
encrypted
region as 'awr-region';`

func TestCopyFromLex_7(t *testing.T) {
	t.Logf("query: %v\n\n", q7)
	err := verifyLex(q7, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q8 = `copy listing 
from 's3://mybucket/data/listings_pipe.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole';`

func TestCopyFromLex_8(t *testing.T) {
	t.Logf("query: %v\n\n", q8)
	err := verifyLex(q8, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q9 = `copy listing 
from 's3://mybucket/data/listings/parquet/' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
format as parquet;`

func TestCopyFromLex_9(t *testing.T) {
	t.Logf("query: %v\n\n", q9)
	err := verifyLex(q9, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q10 = `copy event
from 's3://mybucket/data/allevents_pipe.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole' 
removequotes
emptyasnull
blanksasnull
maxerror 5
delimiter '|'
timeformat 'YYYY-MM-DD HH:MI:SS';`

func TestCopyFromLex_10(t *testing.T) {
	t.Logf("query: %v\n\n", q10)
	err := verifyLex(q10, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q11 = `copy venue
	from 's3://mybucket/data/venue_fw.txt' 
	iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
	fixedwidth 'venueid:3,venuename:25,venuecity:12,venuestate:2,venueseats:6';`

func TestCopyFromLex_11(t *testing.T) {
	t.Logf("query: %v\n\n", q11)
	err := verifyLex(q11, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q12 = `copy category
from 's3://mybucket/data/category_csv.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole' 
csv;`

func TestCopyFromLex_12(t *testing.T) {
	t.Logf("query: %v\n\n", q12)
	err := verifyLex(q12, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q13 = `copy category
from 's3://mybucket/data/category_csv.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole' 
csv quote as '%';`

func TestCopyFromLex_13(t *testing.T) {
	t.Logf("query: %v\n\n", q13)
	err := verifyLex(q13, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q14 = `copy venue
from 's3://mybucket/data/venue.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
explicit_ids;`

func TestCopyFromLex_14(t *testing.T) {
	t.Logf("query: %v\n\n", q14)
	err := verifyLex(q14, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q15 = `copy time
from 's3://mybucket/data/timerows.gz' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
gzip
delimiter '|';`

func TestCopyFromLex_15(t *testing.T) {
	t.Logf("query: %v\n\n", q15)
	err := verifyLex(q15, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q16 = `copy timestamp1 
	from 's3://mybucket/data/time.txt' 
	iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
	timeformat 'YYYY-MM-DD HH:MI:SS';`

func TestCopyFromLex_16(t *testing.T) {
	t.Logf("query: %v\n\n", q16)
	err := verifyLex(q16, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q17 = `copy venue_new(venueid, venuename, venuecity, venuestate) 
from 's3://mybucket/data/venue_noseats.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
delimiter '|';`

func TestCopyFromLex_17(t *testing.T) {
	t.Logf("query: %v\n\n", q17)
	err := verifyLex(q17, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q18 = `copy venue(venueid, venuecity, venuestate) 
	from 's3://mybucket/data/venue_pipe.txt' 
	iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
	delimiter '|';`

func TestCopyFromLex_18(t *testing.T) {
	t.Logf("query: %v\n\n", q18)
	err := verifyLex(q18, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q19 = `copy venue(venueid, venuename, venuecity, venuestate) 
from 's3://mybucket/data/venue_pipe.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
delimiter '|' explicit_ids;`

func TestCopyFromLex_19(t *testing.T) {
	t.Logf("query: %v\n\n", q19)
	err := verifyLex(q19, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q20 = `copy venue(venuename, venuecity, venuestate) 
from 's3://mybucket/data/venue_pipe.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
delimiter '|' explicit_ids;`

func TestCopyFromLex_20(t *testing.T) {
	t.Logf("query: %v\n\n", q20)
	err := verifyLex(q20, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q21 = `copy venue(venueid, venuename, venuecity, venuestate)
from 's3://mybucket/data/venue_pipe.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
delimiter '|';`

func TestCopyFromLex_21(t *testing.T) {
	t.Logf("query: %v\n\n", q21)
	err := verifyLex(q21, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q22 = `copy redshiftinfo from 's3://mybucket/data/redshiftinfo.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole' 
delimiter '|' escape;`

func TestCopyFromLex_22(t *testing.T) {
	t.Logf("query: %v\n\n", q22)
	err := verifyLex(q22, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q23 = `copy category
from 's3://mybucket/category_object_auto.json'
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole' 
json 'auto';`

func TestCopyFromLex_23(t *testing.T) {
	t.Logf("query: %v\n\n", q23)
	err := verifyLex(q23, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q24 = `copy category
from 's3://mybucket/category_object_paths.json'
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole' 
json 's3://mybucket/category_jsonpath.json';`

func TestCopyFromLex_24(t *testing.T) {
	t.Logf("query: %v\n\n", q24)
	err := verifyLex(q24, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q25 = `copy category
from 's3://mybucket/category_array_data.json'
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole' 
json 's3://mybucket/category_array_jsonpath.json';`

func TestCopyFromLex_25(t *testing.T) {
	t.Logf("query: %v\n\n", q25)
	err := verifyLex(q25, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q26 = `copy category
from 's3://mybucket/category_auto.avro'
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
format as avro 'auto';`

func TestCopyFromLex_26(t *testing.T) {
	t.Logf("query: %v\n\n", q26)
	err := verifyLex(q26, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q27 = `copy category
from 's3://mybucket/category_object_paths.avro'
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole' 
format avro 's3://mybucket/category_path.avropath ';`

func TestCopyFromLex_27(t *testing.T) {
	t.Logf("query: %v\n\n", q27)
	err := verifyLex(q27, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}

var q28 = `copy t2 from 's3://mybucket/data/nlTest2.txt' 
	iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'  
	escape
	delimiter as '|';`

func TestCopyFromLex_28(t *testing.T) {
	t.Logf("query: %v\n\n", q28)
	err := verifyLex(q28, t)
	if err != nil {
		t.Logf("got:%v  \n", err)
		t.FailNow()
	}
}
