package main

type Scanner struct {
	source   string
	tokens   []Token
	start    int
	current  int
	line     int
	reporter *ErrorHandler
}

func NewScanner(source string, reporter *ErrorHandler) Scanner {
	return Scanner{source: source, reporter: reporter}
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, NewToken(EOF, "", s.line, nil))
	return s.tokens
}

/*
PRIVATE METHODS
TODO: REFACTOR LATER
*/

func (s *Scanner) scanToken() {
	char := s.advance()
	switch char {
    case '(':
        s.addToken(LEFT_PAREN)
    case ')':
        s.addToken(RIGHT_PAREN)
    case '{':
        s.addToken(LEFT_BRACE)
    case '}':
        s.addToken(RIGHT_BRACE)
    case ',':
        s.addToken(COMMA)
    case '.':
        s.addToken(DOT)
    case '-':
        s.addToken(MINUS)
    case '+':
        s.addToken(PLUS)
    case ';':
        s.addToken(SEMICOLON)
    case '*':
        s.addToken(STAR)
    default:
        s.reporter.error(s.line, "Unexpected character.")
    }
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() byte {
    c := s.source[s.current]
    s.current++
    return c
}

func (s* Scanner) addToken(tokenType TokenType){
	s.addTokenLiteral(tokenType, nil)
}

func (s *Scanner) addTokenLiteral(tokenType TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, NewToken(tokenType, text, s.line, literal))
}