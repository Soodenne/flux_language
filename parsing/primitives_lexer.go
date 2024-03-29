// Code generated from Primitives.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type Primitives struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PrimitivesLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func primitivesLexerInit() {
	staticData := &PrimitivesLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "'text'", "", "'boolean'", "", "'num'", "'na'", "", "", "",
		"'ipv4'", "'loop'", "'if'", "'else'", "'fun'", "'return'", "'->'", "'{'",
		"'}'", "'('", "')'", "'['", "']'", "'.'", "':'", "';'", "','", "'@'",
		"'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'='", "'!='", "'<'", "'>'",
		"'<='", "'>='", "'and'", "'or'", "'xor'", "'not'", "'++'", "'--'",
	}
	staticData.SymbolicNames = []string{
		"", "TEXT", "TEXT_TYPE", "BOOLEAN", "BOOLEAN_TYPE", "NUMBER", "NUMBER_TYPE",
		"NULL", "DIGIT", "OCTET", "IPV4", "IPV4_TYPE", "LOOP", "IF", "ELSE",
		"FUNC", "RETURN", "RETURN_TYPE", "L_BLOCK", "R_BLOCK", "L_PAREN", "R_PAREN",
		"L_SQUARE", "R_SQUARE", "DOT", "COLON", "SEMICOLON", "COMMA", "AT",
		"OP_PLUS", "OP_MINUS", "OP_MULTIPLY", "OP_DIVIDE", "OP_MOD", "OP_POWER",
		"OP_EQUAL", "OP_NOT_EQUAL", "OP_LESS", "OP_GREATER", "OP_LESS_EQUAL",
		"OP_GREATER_EQUAL", "OP_AND", "OP_OR", "OP_XOR", "OP_NOT", "OP_INCREMENT",
		"OP_DECREMENT", "VAR_IDENTIFIER", "COMMON_IDENTIFIER", "NEWLINE", "WS",
	}
	staticData.RuleNames = []string{
		"TEXT", "TEXT_TYPE", "BOOLEAN", "BOOLEAN_TYPE", "NUMBER", "NUMBER_TYPE",
		"NULL", "DIGIT", "OCTET", "IPV4", "IPV4_TYPE", "LOOP", "IF", "ELSE",
		"FUNC", "RETURN", "RETURN_TYPE", "L_BLOCK", "R_BLOCK", "L_PAREN", "R_PAREN",
		"L_SQUARE", "R_SQUARE", "DOT", "COLON", "SEMICOLON", "COMMA", "AT",
		"OP_PLUS", "OP_MINUS", "OP_MULTIPLY", "OP_DIVIDE", "OP_MOD", "OP_POWER",
		"OP_EQUAL", "OP_NOT_EQUAL", "OP_LESS", "OP_GREATER", "OP_LESS_EQUAL",
		"OP_GREATER_EQUAL", "OP_AND", "OP_OR", "OP_XOR", "OP_NOT", "OP_INCREMENT",
		"OP_DECREMENT", "VAR_IDENTIFIER", "COMMON_IDENTIFIER", "NEWLINE", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 50, 323, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 1, 0, 1, 0, 5, 0, 104, 8, 0,
		10, 0, 12, 0, 107, 9, 0, 1, 0, 1, 0, 1, 0, 5, 0, 112, 8, 0, 10, 0, 12,
		0, 115, 9, 0, 1, 0, 3, 0, 118, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 134, 8, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 3, 4, 145, 8, 4, 1, 4,
		4, 4, 148, 8, 4, 11, 4, 12, 4, 149, 1, 4, 1, 4, 5, 4, 154, 8, 4, 10, 4,
		12, 4, 157, 9, 4, 3, 4, 159, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 3, 8, 173, 8, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 186, 8, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29,
		1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1,
		34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38,
		1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1,
		42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44,
		1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 5, 46, 300, 8, 46, 10, 46, 12, 46, 303,
		9, 46, 1, 47, 1, 47, 5, 47, 307, 8, 47, 10, 47, 12, 47, 310, 9, 47, 1,
		48, 3, 48, 313, 8, 48, 1, 48, 1, 48, 1, 49, 4, 49, 318, 8, 49, 11, 49,
		12, 49, 319, 1, 49, 1, 49, 1, 113, 0, 50, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29,
		15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47,
		24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65,
		33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83,
		42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 1,
		0, 9, 1, 0, 34, 34, 1, 0, 48, 57, 1, 0, 49, 57, 1, 0, 48, 52, 1, 0, 48,
		53, 1, 0, 97, 122, 3, 0, 48, 57, 65, 90, 97, 122, 1, 0, 65, 90, 3, 0, 9,
		10, 13, 13, 32, 32, 339, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0,
		0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0,
		0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1,
		0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67,
		1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0,
		75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0,
		0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0,
		0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0,
		0, 0, 0, 99, 1, 0, 0, 0, 1, 117, 1, 0, 0, 0, 3, 119, 1, 0, 0, 0, 5, 133,
		1, 0, 0, 0, 7, 135, 1, 0, 0, 0, 9, 144, 1, 0, 0, 0, 11, 160, 1, 0, 0, 0,
		13, 164, 1, 0, 0, 0, 15, 167, 1, 0, 0, 0, 17, 185, 1, 0, 0, 0, 19, 187,
		1, 0, 0, 0, 21, 195, 1, 0, 0, 0, 23, 200, 1, 0, 0, 0, 25, 205, 1, 0, 0,
		0, 27, 208, 1, 0, 0, 0, 29, 213, 1, 0, 0, 0, 31, 217, 1, 0, 0, 0, 33, 224,
		1, 0, 0, 0, 35, 227, 1, 0, 0, 0, 37, 229, 1, 0, 0, 0, 39, 231, 1, 0, 0,
		0, 41, 233, 1, 0, 0, 0, 43, 235, 1, 0, 0, 0, 45, 237, 1, 0, 0, 0, 47, 239,
		1, 0, 0, 0, 49, 241, 1, 0, 0, 0, 51, 243, 1, 0, 0, 0, 53, 245, 1, 0, 0,
		0, 55, 247, 1, 0, 0, 0, 57, 249, 1, 0, 0, 0, 59, 251, 1, 0, 0, 0, 61, 253,
		1, 0, 0, 0, 63, 255, 1, 0, 0, 0, 65, 257, 1, 0, 0, 0, 67, 259, 1, 0, 0,
		0, 69, 261, 1, 0, 0, 0, 71, 263, 1, 0, 0, 0, 73, 266, 1, 0, 0, 0, 75, 268,
		1, 0, 0, 0, 77, 270, 1, 0, 0, 0, 79, 273, 1, 0, 0, 0, 81, 276, 1, 0, 0,
		0, 83, 280, 1, 0, 0, 0, 85, 283, 1, 0, 0, 0, 87, 287, 1, 0, 0, 0, 89, 291,
		1, 0, 0, 0, 91, 294, 1, 0, 0, 0, 93, 297, 1, 0, 0, 0, 95, 304, 1, 0, 0,
		0, 97, 312, 1, 0, 0, 0, 99, 317, 1, 0, 0, 0, 101, 105, 5, 34, 0, 0, 102,
		104, 8, 0, 0, 0, 103, 102, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103,
		1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 108, 1, 0, 0, 0, 107, 105, 1, 0,
		0, 0, 108, 118, 5, 34, 0, 0, 109, 113, 5, 39, 0, 0, 110, 112, 9, 0, 0,
		0, 111, 110, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 113,
		111, 1, 0, 0, 0, 114, 116, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 118,
		5, 39, 0, 0, 117, 101, 1, 0, 0, 0, 117, 109, 1, 0, 0, 0, 118, 2, 1, 0,
		0, 0, 119, 120, 5, 116, 0, 0, 120, 121, 5, 101, 0, 0, 121, 122, 5, 120,
		0, 0, 122, 123, 5, 116, 0, 0, 123, 4, 1, 0, 0, 0, 124, 125, 5, 116, 0,
		0, 125, 126, 5, 114, 0, 0, 126, 127, 5, 117, 0, 0, 127, 134, 5, 101, 0,
		0, 128, 129, 5, 102, 0, 0, 129, 130, 5, 97, 0, 0, 130, 131, 5, 108, 0,
		0, 131, 132, 5, 115, 0, 0, 132, 134, 5, 101, 0, 0, 133, 124, 1, 0, 0, 0,
		133, 128, 1, 0, 0, 0, 134, 6, 1, 0, 0, 0, 135, 136, 5, 98, 0, 0, 136, 137,
		5, 111, 0, 0, 137, 138, 5, 111, 0, 0, 138, 139, 5, 108, 0, 0, 139, 140,
		5, 101, 0, 0, 140, 141, 5, 97, 0, 0, 141, 142, 5, 110, 0, 0, 142, 8, 1,
		0, 0, 0, 143, 145, 5, 45, 0, 0, 144, 143, 1, 0, 0, 0, 144, 145, 1, 0, 0,
		0, 145, 147, 1, 0, 0, 0, 146, 148, 7, 1, 0, 0, 147, 146, 1, 0, 0, 0, 148,
		149, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 158,
		1, 0, 0, 0, 151, 155, 5, 46, 0, 0, 152, 154, 7, 1, 0, 0, 153, 152, 1, 0,
		0, 0, 154, 157, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0,
		156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 158, 151, 1, 0, 0, 0, 158,
		159, 1, 0, 0, 0, 159, 10, 1, 0, 0, 0, 160, 161, 5, 110, 0, 0, 161, 162,
		5, 117, 0, 0, 162, 163, 5, 109, 0, 0, 163, 12, 1, 0, 0, 0, 164, 165, 5,
		110, 0, 0, 165, 166, 5, 97, 0, 0, 166, 14, 1, 0, 0, 0, 167, 168, 7, 1,
		0, 0, 168, 16, 1, 0, 0, 0, 169, 186, 5, 48, 0, 0, 170, 172, 7, 2, 0, 0,
		171, 173, 3, 15, 7, 0, 172, 171, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173,
		186, 1, 0, 0, 0, 174, 175, 5, 49, 0, 0, 175, 176, 3, 15, 7, 0, 176, 177,
		3, 15, 7, 0, 177, 186, 1, 0, 0, 0, 178, 179, 5, 50, 0, 0, 179, 180, 7,
		3, 0, 0, 180, 186, 3, 15, 7, 0, 181, 182, 5, 50, 0, 0, 182, 183, 5, 53,
		0, 0, 183, 184, 1, 0, 0, 0, 184, 186, 7, 4, 0, 0, 185, 169, 1, 0, 0, 0,
		185, 170, 1, 0, 0, 0, 185, 174, 1, 0, 0, 0, 185, 178, 1, 0, 0, 0, 185,
		181, 1, 0, 0, 0, 186, 18, 1, 0, 0, 0, 187, 188, 3, 17, 8, 0, 188, 189,
		5, 46, 0, 0, 189, 190, 3, 17, 8, 0, 190, 191, 5, 46, 0, 0, 191, 192, 3,
		17, 8, 0, 192, 193, 5, 46, 0, 0, 193, 194, 3, 17, 8, 0, 194, 20, 1, 0,
		0, 0, 195, 196, 5, 105, 0, 0, 196, 197, 5, 112, 0, 0, 197, 198, 5, 118,
		0, 0, 198, 199, 5, 52, 0, 0, 199, 22, 1, 0, 0, 0, 200, 201, 5, 108, 0,
		0, 201, 202, 5, 111, 0, 0, 202, 203, 5, 111, 0, 0, 203, 204, 5, 112, 0,
		0, 204, 24, 1, 0, 0, 0, 205, 206, 5, 105, 0, 0, 206, 207, 5, 102, 0, 0,
		207, 26, 1, 0, 0, 0, 208, 209, 5, 101, 0, 0, 209, 210, 5, 108, 0, 0, 210,
		211, 5, 115, 0, 0, 211, 212, 5, 101, 0, 0, 212, 28, 1, 0, 0, 0, 213, 214,
		5, 102, 0, 0, 214, 215, 5, 117, 0, 0, 215, 216, 5, 110, 0, 0, 216, 30,
		1, 0, 0, 0, 217, 218, 5, 114, 0, 0, 218, 219, 5, 101, 0, 0, 219, 220, 5,
		116, 0, 0, 220, 221, 5, 117, 0, 0, 221, 222, 5, 114, 0, 0, 222, 223, 5,
		110, 0, 0, 223, 32, 1, 0, 0, 0, 224, 225, 5, 45, 0, 0, 225, 226, 5, 62,
		0, 0, 226, 34, 1, 0, 0, 0, 227, 228, 5, 123, 0, 0, 228, 36, 1, 0, 0, 0,
		229, 230, 5, 125, 0, 0, 230, 38, 1, 0, 0, 0, 231, 232, 5, 40, 0, 0, 232,
		40, 1, 0, 0, 0, 233, 234, 5, 41, 0, 0, 234, 42, 1, 0, 0, 0, 235, 236, 5,
		91, 0, 0, 236, 44, 1, 0, 0, 0, 237, 238, 5, 93, 0, 0, 238, 46, 1, 0, 0,
		0, 239, 240, 5, 46, 0, 0, 240, 48, 1, 0, 0, 0, 241, 242, 5, 58, 0, 0, 242,
		50, 1, 0, 0, 0, 243, 244, 5, 59, 0, 0, 244, 52, 1, 0, 0, 0, 245, 246, 5,
		44, 0, 0, 246, 54, 1, 0, 0, 0, 247, 248, 5, 64, 0, 0, 248, 56, 1, 0, 0,
		0, 249, 250, 5, 43, 0, 0, 250, 58, 1, 0, 0, 0, 251, 252, 5, 45, 0, 0, 252,
		60, 1, 0, 0, 0, 253, 254, 5, 42, 0, 0, 254, 62, 1, 0, 0, 0, 255, 256, 5,
		47, 0, 0, 256, 64, 1, 0, 0, 0, 257, 258, 5, 37, 0, 0, 258, 66, 1, 0, 0,
		0, 259, 260, 5, 94, 0, 0, 260, 68, 1, 0, 0, 0, 261, 262, 5, 61, 0, 0, 262,
		70, 1, 0, 0, 0, 263, 264, 5, 33, 0, 0, 264, 265, 5, 61, 0, 0, 265, 72,
		1, 0, 0, 0, 266, 267, 5, 60, 0, 0, 267, 74, 1, 0, 0, 0, 268, 269, 5, 62,
		0, 0, 269, 76, 1, 0, 0, 0, 270, 271, 5, 60, 0, 0, 271, 272, 5, 61, 0, 0,
		272, 78, 1, 0, 0, 0, 273, 274, 5, 62, 0, 0, 274, 275, 5, 61, 0, 0, 275,
		80, 1, 0, 0, 0, 276, 277, 5, 97, 0, 0, 277, 278, 5, 110, 0, 0, 278, 279,
		5, 100, 0, 0, 279, 82, 1, 0, 0, 0, 280, 281, 5, 111, 0, 0, 281, 282, 5,
		114, 0, 0, 282, 84, 1, 0, 0, 0, 283, 284, 5, 120, 0, 0, 284, 285, 5, 111,
		0, 0, 285, 286, 5, 114, 0, 0, 286, 86, 1, 0, 0, 0, 287, 288, 5, 110, 0,
		0, 288, 289, 5, 111, 0, 0, 289, 290, 5, 116, 0, 0, 290, 88, 1, 0, 0, 0,
		291, 292, 5, 43, 0, 0, 292, 293, 5, 43, 0, 0, 293, 90, 1, 0, 0, 0, 294,
		295, 5, 45, 0, 0, 295, 296, 5, 45, 0, 0, 296, 92, 1, 0, 0, 0, 297, 301,
		7, 5, 0, 0, 298, 300, 7, 6, 0, 0, 299, 298, 1, 0, 0, 0, 300, 303, 1, 0,
		0, 0, 301, 299, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 94, 1, 0, 0, 0,
		303, 301, 1, 0, 0, 0, 304, 308, 7, 7, 0, 0, 305, 307, 7, 6, 0, 0, 306,
		305, 1, 0, 0, 0, 307, 310, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 308, 309,
		1, 0, 0, 0, 309, 96, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 311, 313, 5, 13,
		0, 0, 312, 311, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0,
		314, 315, 5, 10, 0, 0, 315, 98, 1, 0, 0, 0, 316, 318, 7, 8, 0, 0, 317,
		316, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 319, 320,
		1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321, 322, 6, 49, 0, 0, 322, 100, 1, 0,
		0, 0, 15, 0, 105, 113, 117, 133, 144, 149, 155, 158, 172, 185, 301, 308,
		312, 319, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// PrimitivesInit initializes any static state used to implement Primitives. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewPrimitives(). You can call this function if you wish to initialize the static state ahead
// of time.
func PrimitivesInit() {
	staticData := &PrimitivesLexerStaticData
	staticData.once.Do(primitivesLexerInit)
}

// NewPrimitives produces a new lexer instance for the optional input antlr.CharStream.
func NewPrimitives(input antlr.CharStream) *Primitives {
	PrimitivesInit()
	l := new(Primitives)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PrimitivesLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Primitives.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Primitives tokens.
const (
	PrimitivesTEXT              = 1
	PrimitivesTEXT_TYPE         = 2
	PrimitivesBOOLEAN           = 3
	PrimitivesBOOLEAN_TYPE      = 4
	PrimitivesNUMBER            = 5
	PrimitivesNUMBER_TYPE       = 6
	PrimitivesNULL              = 7
	PrimitivesDIGIT             = 8
	PrimitivesOCTET             = 9
	PrimitivesIPV4              = 10
	PrimitivesIPV4_TYPE         = 11
	PrimitivesLOOP              = 12
	PrimitivesIF                = 13
	PrimitivesELSE              = 14
	PrimitivesFUNC              = 15
	PrimitivesRETURN            = 16
	PrimitivesRETURN_TYPE       = 17
	PrimitivesL_BLOCK           = 18
	PrimitivesR_BLOCK           = 19
	PrimitivesL_PAREN           = 20
	PrimitivesR_PAREN           = 21
	PrimitivesL_SQUARE          = 22
	PrimitivesR_SQUARE          = 23
	PrimitivesDOT               = 24
	PrimitivesCOLON             = 25
	PrimitivesSEMICOLON         = 26
	PrimitivesCOMMA             = 27
	PrimitivesAT                = 28
	PrimitivesOP_PLUS           = 29
	PrimitivesOP_MINUS          = 30
	PrimitivesOP_MULTIPLY       = 31
	PrimitivesOP_DIVIDE         = 32
	PrimitivesOP_MOD            = 33
	PrimitivesOP_POWER          = 34
	PrimitivesOP_EQUAL          = 35
	PrimitivesOP_NOT_EQUAL      = 36
	PrimitivesOP_LESS           = 37
	PrimitivesOP_GREATER        = 38
	PrimitivesOP_LESS_EQUAL     = 39
	PrimitivesOP_GREATER_EQUAL  = 40
	PrimitivesOP_AND            = 41
	PrimitivesOP_OR             = 42
	PrimitivesOP_XOR            = 43
	PrimitivesOP_NOT            = 44
	PrimitivesOP_INCREMENT      = 45
	PrimitivesOP_DECREMENT      = 46
	PrimitivesVAR_IDENTIFIER    = 47
	PrimitivesCOMMON_IDENTIFIER = 48
	PrimitivesNEWLINE           = 49
	PrimitivesWS                = 50
)
