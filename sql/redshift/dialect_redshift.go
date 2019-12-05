package redshift

import (
	"github.com/araddon/qlbridge/lex"
)

var (
	// TokenCopy is COPY
	TokenCopy lex.TokenType = 250 // TURELLA
	// TokenCIamRole is "iam_role" for authorization
	TokenCIamRole lex.TokenType = 350 // iam_role // TURELLA
	// TokenCDelimiter is "delimiter"
	TokenCDelimiter lex.TokenType = 351 // delimiter // TURELLA
	// TokenCEscape is "escape"
	TokenCEscape lex.TokenType = 352 // escape // TURELLA
	// TokenCManifest is "manifest"
	TokenCManifest lex.TokenType = 353 // manifest // TURELLA
	// TokenCQuoteAs is "quote-as"
	TokenCQuoteAs lex.TokenType = 354 // quote-as // TURELLA
	// TokenCExplicitIds is "explicit_ids"
	TokenCExplicitIds lex.TokenType = 355 // explicit_ids // TURELLA
	// TokenCRemovequotes is "removequotes"
	TokenCRemovequotes lex.TokenType = 356 // removequotes // TURELLA
	// TokenCEmptyasnull is "emptyasnull"
	TokenCEmptyasnull lex.TokenType = 357 // emptyasnull // TURELLA
	// TokenCBlanksasnull is "blanksasnull"
	TokenCBlanksasnull lex.TokenType = 358 // blanksasnull // TURELLA
	// TokenCMaxerror is "maxerror"
	TokenCMaxerror lex.TokenType = 359 // maxerror // TURELLA
	// TokenCTimeformat is "timeformat"
	TokenCTimeformat lex.TokenType = 360 // timeformat // TURELLA
	// TokenCFormatAs is "format-as"
	TokenCFormatAs lex.TokenType = 361 // format-as // TURELLA
	// TokenCFormatParquet is "parquet"
	TokenCFormatParquet lex.TokenType = 370 // parquet // TURELLA
	// TokenCFormatAvro is "avro"
	TokenCFormatAvro lex.TokenType = 371 // avro // TURELLA
	// TokenCGzip is "gzip"
	TokenCGzip lex.TokenType = 372 // gzip // TURELLA
	// TokenCFixedwidth is "fixedwidth"
	TokenCFixedwidth lex.TokenType = 373 // fixedwidth // TURELLA
	// TokenCCsv is "csv"
	TokenCCsv lex.TokenType = 374 // csv // TURELLA
	// TokenCJson is "JSON"
	TokenCJson lex.TokenType = 375 // JSON // TURELLA
	// TokenCLzop is "lzop"
	TokenCLzop lex.TokenType = 376 // lzop // TURELLA
	// TokenCReadratio is "readratio"
	TokenCReadratio lex.TokenType = 377 // readratio // TURELLA
	// TokenCEncripted is "encripted"
	TokenCEncripted lex.TokenType = 378 // encripted // TURELLA
	// TokenCRegion is "region"
	TokenCRegion lex.TokenType = 379 // region // TURELLA
	// TokenCSSH is "ssh"
	TokenCSSH lex.TokenType = 380 // ssh // TURELLA

	// RedshiftDialect is a dialect to parse the COPY FROM command
	RedshiftDialect *lex.Dialect = &lex.Dialect{
		Statements: []*lex.Clause{
			//{Token: TokenUpdate, Clauses: SqlUpdateR},
			{Token: TokenCopy, Clauses: RedshiftCopy},
		},
		IdentityQuoting: lex.IdentityQuoting,
	}

	// RedshiftCopy copy statement
	RedshiftCopy = []*lex.Clause{
		{Token: TokenCopy, Lexer: lex.LexIdentifierOfType(lex.TokenTable)},
		{Token: lex.TokenLeftParenthesis, Lexer: lex.LexColumnNames, Optional: true},
		{Token: lex.TokenFrom, Lexer: lex.LexTableReferences, Clauses: sourceClauses, Name: "copyFrom.source"},
		//{Token: TokenCIamRole, Lexer: lex.LexValue},
		{Token: TokenCRemovequotes, Lexer: lex.LexEmpty, Optional: true},
		{Token: TokenCEmptyasnull, Lexer: lex.LexEmpty, Optional: true},
		{Token: TokenCBlanksasnull, Lexer: lex.LexEmpty, Optional: true},
		{Token: TokenCMaxerror, Lexer: lex.LexValue, Optional: true},
		{Token: TokenCGzip, Lexer: lex.LexEmpty, Optional: true},
		{Token: TokenCEscape, Lexer: lex.LexEmpty, Optional: true},
		{Token: TokenCDelimiter, Lexer: lex.LexValue, Optional: true, Clauses: delimiterClauses},
		{Token: TokenCLzop, Lexer: lex.LexEmpty, Optional: true},
		{Token: TokenCEscape, Lexer: lex.LexEmpty, Optional: true},
		{Token: TokenCExplicitIds, Lexer: lex.LexEmpty, Optional: true},
		{Token: TokenCJson, Lexer: lex.LexValue, Optional: true},
		{Token: TokenCReadratio, Lexer: lex.LexValue, Optional: true},
		// {Token: TokenCManifest, Lexer: lex.LexEmpty, Optional: true},
		{Token: TokenCTimeformat, Lexer: lex.LexValue, Optional: true},
		{Token: TokenCFixedwidth, Lexer: lex.LexValue, Optional: true},
		{Token: TokenCCsv, Lexer: lex.LexEmpty, Optional: true},
		//{Token: TokenCQuoteAs, Lexer: lex.LexEmpty, Optional: true},
		//{Token: lex.TokenAs, Lexer: lex.LexValue, Optional: true},
		{Token: TokenCQuoteAs, Lexer: lex.LexEmpty, Optional: true, Clauses: quoteClauses},
		// {Token: TokenCFormatAs, Lexer: SkipAs, Optional: true},
		// {Token: TokenCFormatParquet, Lexer: lex.LexEmpty, Optional: true},
		// {Token: TokenCFormatAvro, Lexer: lex.LexIdentityOrValue, Optional: true},
		{Token: TokenCFormatAs, Lexer: lex.LexEmpty, Optional: true, Clauses: formatClauses},
		{Token: lex.TokenEofOrEos, Lexer: lex.LexEmpty},
	}

	sourceClauses = []*lex.Clause{
		{Token: TokenCIamRole, Lexer: lex.LexValue},
		{Token: lex.TokenValue, Lexer: lex.LexValue},
		{Token: TokenCManifest, Lexer: lex.LexEmpty, Optional: true},
		//{Token: TokenCRegion, Lexer: lex.LexEmpty, Optional: true},
		//{Token: lex.TokenAs, Lexer: lex.LexValue, Optional: true},
		{Token: TokenCRegion, Lexer: lex.LexValue, Optional: true, Clauses: regionClauses},
		{Token: lex.TokenValue, Lexer: lex.LexValue, Optional: true},
		{Token: TokenCSSH, Lexer: lex.LexEmpty, Optional: true},
	}

	delimiterClauses = []*lex.Clause{
		{Token: lex.TokenAs, Lexer: lex.LexEmpty, Optional: true},
		{Token: lex.TokenValue, Lexer: lex.LexValue},
	}

	regionClauses = []*lex.Clause{
		{Token: lex.TokenAs, Lexer: lex.LexValue, Optional: true},
		{Token: lex.TokenValue, Lexer: lex.LexIdentityOrValue},
	}

	quoteClauses = []*lex.Clause{
		{Token: lex.TokenAs, Lexer: lex.LexEmpty, Optional: true},
		{Token: lex.TokenValue, Lexer: lex.LexIdentityOrValue},
	}

	formatClauses = []*lex.Clause{
		{Token: lex.TokenAs, Lexer: lex.LexEmpty, Optional: true},
		{Token: TokenCFormatParquet, Lexer: lex.LexEmpty, Optional: true},
		{Token: TokenCFormatAvro, Lexer: lex.LexIdentityOrValue, Optional: true},
		{Token: lex.TokenValue, Lexer: lex.LexIdentityOrValue},
	}
)

func init() {
	// inject any new tokens into QLBridge.Lex describing the custom tokens we created
	lex.TokenNameMap[TokenCopy] = &lex.TokenInfo{Description: "copy"}                         // TURELLA
	lex.TokenNameMap[TokenCIamRole] = &lex.TokenInfo{Kw: "iam_role", Description: "iam_role"} // TURELLA
	lex.TokenNameMap[TokenCDelimiter] = &lex.TokenInfo{Description: "delimiter"}              // TURELLA
	lex.TokenNameMap[TokenCEscape] = &lex.TokenInfo{Description: "escape"}                    // TURELLA
	lex.TokenNameMap[TokenCManifest] = &lex.TokenInfo{Description: "manifest"}                // TURELLA
	lex.TokenNameMap[TokenCQuoteAs] = &lex.TokenInfo{Description: "quote"}                    // TURELLA
	lex.TokenNameMap[TokenCExplicitIds] = &lex.TokenInfo{Description: "explicit_ids"}         // TURELLA
	lex.TokenNameMap[TokenCRemovequotes] = &lex.TokenInfo{Description: "removequotes"}        // TURELLA
	lex.TokenNameMap[TokenCEmptyasnull] = &lex.TokenInfo{Description: "emptyasnull"}          // TURELLA
	lex.TokenNameMap[TokenCBlanksasnull] = &lex.TokenInfo{Description: "blanksasnull"}        // TURELLA
	lex.TokenNameMap[TokenCMaxerror] = &lex.TokenInfo{Description: "maxerror"}                // TURELLA
	lex.TokenNameMap[TokenCTimeformat] = &lex.TokenInfo{Description: "timeformat"}            // TURELLA
	lex.TokenNameMap[TokenCFormatAs] = &lex.TokenInfo{Description: "format"}                  // TURELLA
	lex.TokenNameMap[TokenCFormatParquet] = &lex.TokenInfo{Description: "parquet"}            // TURELLA
	lex.TokenNameMap[TokenCFormatAvro] = &lex.TokenInfo{Description: "avro"}                  // TURELLA
	lex.TokenNameMap[TokenCGzip] = &lex.TokenInfo{Description: "gzip"}                        // TURELLA
	lex.TokenNameMap[TokenCFixedwidth] = &lex.TokenInfo{Description: "fixedwidth"}            // TURELLA
	lex.TokenNameMap[TokenCCsv] = &lex.TokenInfo{Description: "csv"}                          // TURELLA
	lex.TokenNameMap[TokenCJson] = &lex.TokenInfo{Description: "myjson"}                      // TURELLA
	lex.TokenNameMap[TokenCLzop] = &lex.TokenInfo{Description: "lzop"}                        // TURELLA
	lex.TokenNameMap[TokenCReadratio] = &lex.TokenInfo{Description: "readratio"}              // TURELLA
	lex.TokenNameMap[TokenCEncripted] = &lex.TokenInfo{Description: "encripted"}              // TURELLA
	lex.TokenNameMap[TokenCRegion] = &lex.TokenInfo{Description: "region"}                    // TURELLA
	lex.TokenNameMap[TokenCSSH] = &lex.TokenInfo{Description: "ssh"}                          // TURELLA

	// OverRide the Identity Characters in lexer to allow a dash in identity
	//lex.IDENTITY_CHARS = "_.-"
	lex.LoadTokenInfo()

	RedshiftDialect.Init()
}

// NewRedshiftLexer creates a new lexer for the input string using RedshiftDialect
// this is sql(ish) compatible parser.
func NewRedshiftLexer(input string) *lex.Lexer {
	return lex.NewLexer(input, RedshiftDialect)
}
