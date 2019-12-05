package redshift

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bisegni/go-c-interface-test/rscopy"

	"github.com/araddon/qlbridge/lex"
)

// ParseCopy parses the COPY staatement, using the redshift dialect
func ParseCopy(sql string) (*rscopy.RsCopy, bool) {
	sql = strings.TrimSpace(sql)
	sql = strings.Replace(sql, "format as ", "format ", 1)
	sql = strings.Replace(sql, "delimiter as ", "delimiter ", 1)
	sql = strings.Replace(sql, "quote as ", "quote ", 1)
	sql = strings.Replace(sql, "region as ", "region ", 1)
	sql = strings.Replace(sql, "\r", " ", 1)
	sql = strings.Replace(sql, "\n", " ", 1)
	if len(sql) <= 0 {
		return nil, false
	}
	l := NewRedshiftLexer(sql)
	res := rscopy.RsCopy{}
	tok := l.NextToken()
	for tok.T != lex.TokenEOF {
		fmt.Printf("got:%v  \n", tok)
		if tok.T == lex.TokenError {
			fmt.Printf("  --> errore: %v \n", tok.Err(l))
		}
		if tok.T == TokenCopy {
			// Gestione della testata
			fmt.Printf(" -> Statement = %s \n", tok.V)
			tok = l.NextToken()
			if tok.T != lex.TokenTable {
				return nil, false
			}
			fmt.Printf(" -> Tabella dati = %s \n", tok.V)
			res.Table = tok.V
			// Riconoscimento della lista dei campi
			l.SkipWhiteSpacesNewLine()
			w := l.PeekWord()
			// fmt.Printf(" 1. %s\n", w)
			if w == "(" {
				tok = l.NextToken()
				// fmt.Printf(" 2. %s\n", tok.V)
				fl := []string{}
				l.SkipWhiteSpacesNewLine()
				w = l.PeekWord()
				// fmt.Printf(" 3. %s\n", w)
				i := 0
				for w != ")" && tok.T != lex.TokenRightParenthesis && tok.T != lex.TokenEofOrEos && i < 10 {
					// fmt.Printf("w = %s - tok = %s\n", w, tok.V)
					tok = l.NextToken() // nome del campo
					// fmt.Printf(" 4. %s\n", tok.V)
					fl = append(fl, tok.V)
					tok = l.NextToken() // separatore virgola
					l.SkipWhiteSpacesNewLine()
					w = l.PeekWord()
					i++
				}
				fmt.Printf(" -> Elenco campi = %v \n", fl)
				res.Fields = fl
			}
		} else if tok.T == lex.TokenFrom {
			// Riconoscimento del tipo FROM-TO e parametro
			fmt.Printf(" -> Direzione copia = %s \n", tok.V)
			tok = l.NextToken()
			if tok.T != lex.TokenValue {
				return nil, false
			}
			fmt.Printf(" -> Sorg/Dest esterna = %s \n", tok.V)
			res.Source.SourceURI = tok.V
		} else if tok.T == TokenCIamRole || (tok.T == lex.TokenIdentity && tok.V == "iam_role") {
			// Riconoscimento del parametro di autenticazione
			fmt.Printf(" -> Autenticazione tipo = %s \n", tok.V)
			tok = l.NextToken()
			if tok.T != lex.TokenValue {
				return nil, false
			}
			fmt.Printf(" -> Autenticazione valore = %s \n", tok.V)
			//a := rscopy.Auth{IamRole: tok.V}
			res.Source.ExternalAuth.IamRole = tok.V
		} else if tok.T == TokenCRegion || (tok.T == lex.TokenIdentity && tok.V == "region") {
			// Riconoscimento del parametro "region"
			fmt.Printf(" -> Origine = %s \n", tok.V)
			l.SkipWhiteSpacesNewLine()
			w := l.PeekWord()
			if w == "as" {
				tok = l.NextToken()
			}
			tok = l.NextToken()
			fmt.Printf(" -> Origine parametro = %s \n", tok.V)
			res.Source.Region = tok.V
		} else if tok.T == TokenCFormatAs {
			// Riconoscimento del parametro "format"
			fmt.Printf(" -> Formato = %s \n", tok.V)
			l.SkipWhiteSpacesNewLine()
			w := l.PeekWord()
			if w == "as" {
				tok = l.NextToken()
			}
			tok = l.NextToken()
			//if tok.T != TokenCFormatParquet && tok.T != TokenCFormatAvro {
			//	return false
			//}
			switch tok.T {
			case TokenCFormatParquet:
				res.Format.Type = rscopy.PARQUET
			case TokenCFormatAvro:
				res.Format.Type = rscopy.AVRO
			default:
				return nil, false
			}
			fmt.Printf(" -> Formato tipologia = %s \n", tok.V)
			opt := "auto"
			l.SkipWhiteSpacesNewLine()
			w = l.PeekWord()
			if w == "as" {
				l.NextToken()
				l.SkipWhiteSpacesNewLine()
				w = l.PeekWord()
			}
			// fmt.Printf(" %s = %d \n", w, len(w))
			if w == "auto" || strings.HasPrefix(w, "'s3") || strings.HasPrefix(w, "'auto") {
				tok = l.NextToken()
				opt = tok.V
			}
			fmt.Printf(" -> Formato tipologia_option = %s \n", opt)
			res.Format.Option = opt
		} else if tok.T == lex.TokenIdentity && tok.V == "json" {
			// Riconoscimento del formato "json"
			fmt.Printf(" -> Formato = %s \n", tok.V)
			res.Format.Type = rscopy.JSON
			opt := "auto"
			l.SkipWhiteSpacesNewLine()
			w := l.PeekWord()
			if w == "as" {
				l.NextToken()
				l.SkipWhiteSpacesNewLine()
				w = l.PeekWord()
			}
			// fmt.Printf(" %s = %d \n", w, len(w))
			if w == "auto" || strings.HasPrefix(w, "'s3") || strings.HasPrefix(w, "'auto") {
				tok = l.NextToken()
				opt = tok.V
			}
			fmt.Printf(" -> Formato json_option = %s \n", opt)
			res.Format.Option = opt
		} else if tok.T == TokenCCsv || tok.T == lex.TokenIdentity && tok.V == "csv" {
			// Riconoscimento del formato "csv"
			fmt.Printf(" -> Formato = %s \n", tok.V)
			res.Format.Type = rscopy.CSV
		} else if tok.T == TokenCFixedwidth {
			// Riconoscimento del formato "fixedwidth"
			fmt.Printf(" -> Formato = %s \n", tok.V)
			res.Format.Type = rscopy.FIXEDWIDTH
			tok = l.NextToken()
			if tok.T != lex.TokenValue {
				return nil, false
			}
			fmt.Printf(" -> Formato fixedwidth_spec = %s \n", tok.V)
			res.Format.Option = tok.V
			// tmp := strings.Split(tok.V, ",")
			// m := make(map[string]int, len(tmp))
			// for _, t := range tmp {
			// 	if strings.Contains(t, ":") {
			// 		tmp := strings.Split(t, ":")
			// 		v, err := strconv.Atoi(tmp[1])
			// 		if err != nil {
			// 			return false
			// 		}
			// 		m[tmp[0]] = v
			// 	}
			// }
			// fmt.Printf(" -> Formato fixedwidth_spec = %+v \n", m)
		} else if tok.T == TokenCSSH || (tok.T == lex.TokenIdentity && tok.V == "ssh") {
			// Riconoscimento del data-source "ssh"
			fmt.Printf(" -> Origine file = %s \n", tok.V)
			res.Source.SSH = true
		} else if tok.T == TokenCReadratio {
			// Riconoscimento del parametro "readratio" (dynamoDB)
			fmt.Printf(" -> Opzione = %s \n", tok.V)
			tok = l.NextToken()
			if tok.T != lex.TokenInteger {
				return nil, false
			}
			fmt.Printf(" -> Opzione parametro = %s \n", tok.V)
			r, err := strconv.Atoi(tok.V)
			if err != nil {
				return nil, false
			}
			res.Format.Readratio = r
		} else if tok.T == TokenCQuoteAs {
			// Riconoscimento del parametro "quote"
			fmt.Printf(" -> Opzione = %s \n", tok.V)
			tok = l.NextToken()
			if tok.T == lex.TokenAs {
				tok = l.NextToken()
			}
			if tok.T != lex.TokenValue {
				return nil, false
			}
			fmt.Printf(" -> Opzione parametro = %s \n", tok.V)
			res.Format.Quote = tok.V
		} else if tok.T == TokenCDelimiter {
			// Riconoscimento del parametro "delimiter"
			fmt.Printf(" -> Opzione = %s \n", tok.V)
			tok = l.NextToken()
			// fmt.Printf(" 1. %s\n", tok.V)
			if tok.T == lex.TokenAs {
				tok = l.NextToken()
			}
			// fmt.Printf(" 2. %s\n", tok.V)
			if tok.T != lex.TokenValue {
				return nil, false
			}
			fmt.Printf(" -> Opzione parametro = %s \n", tok.V)
			res.Format.Delimiter = tok.V
		} else if tok.T == TokenCTimeformat {
			// Riconoscimento del parametro "timeformat"
			fmt.Printf(" -> Opzione = %s \n", tok.V)
			tok = l.NextToken()
			if tok.T == lex.TokenAs {
				tok = l.NextToken()
			}
			if tok.T != lex.TokenValue {
				return nil, false
			}
			fmt.Printf(" -> Opzione parametro = %s \n", tok.V)
			res.Format.Timeformat = tok.V
		} else if tok.T == TokenCLzop {
			// Riconoscimento dei formati di compressione
			fmt.Printf(" -> Compressione file = %s \n", tok.V)
			res.Format.Compression = rscopy.LZOP
		} else if tok.T == TokenCLzop || tok.T == TokenCGzip {
			// Riconoscimento dei formati di compressione
			fmt.Printf(" -> Compressione file = %s \n", tok.V)
			res.Format.Compression = rscopy.GZIP
		} else if tok.T == TokenCManifest || (tok.T == lex.TokenIdentity && tok.V == "manifest") {
			// Riconoscimento del parametro "manifest"
			fmt.Printf(" -> Opzione = %s \n", tok.V)
			res.Source.Manifest = true
		} else if tok.T == TokenCEncripted || (tok.T == lex.TokenIdentity && tok.V == "encrypted") {
			// Riconoscimento del parametro "encrypted"
			fmt.Printf(" -> Opzione = %s \n", tok.V)
			res.Source.Encrypted = true
		} else if tok.T == TokenCExplicitIds || (tok.T == lex.TokenIdentity && tok.V == "explicit_ids") {
			// Riconoscimento del parametro "explicit_ids"
			fmt.Printf(" -> Opzione = %s \n", tok.V)
			res.Format.ExplicitIds = true
		} else if tok.T == TokenCEscape || (tok.T == lex.TokenIdentity && tok.V == "escape") {
			// Riconoscimento del parametro "escape"
			fmt.Printf(" -> Opzione = %s \n", tok.V)
			res.Format.Escape = true
		} else if tok.T == lex.TokenEOS {
			// Riconoscimento del punto e virgola finale
			break
		} else {
			// Casistiche non gestite
			fmt.Printf(" ** non riconosciuto = %s \n", tok.V)
			return nil, false
		}
		tok = l.NextToken()
	}
	return &res, true
}
