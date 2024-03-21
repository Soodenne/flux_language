// Code generated from Flux.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Flux
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type Flux struct {
	*antlr.BaseParser
}

var FluxParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func fluxParserInit() {
	staticData := &FluxParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'text'", "", "'boolean'", "", "'num'", "'na'", "", "", "",
		"'ipv4'", "", "'array'", "'loop'", "'if'", "'else'", "'fun'", "'return'",
		"'->'", "'{'", "'}'", "'('", "')'", "'['", "']'", "'.'", "':'", "';'",
		"','", "'@'", "'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'='", "'!='",
		"'<'", "'>'", "'<='", "'>='", "'and'", "'or'", "'xor'", "'not'", "'++'",
		"'--'",
	}
	staticData.SymbolicNames = []string{
		"", "TEXT", "TEXT_TYPE", "BOOLEAN", "BOOLEAN_TYPE", "NUMBER", "NUMBER_TYPE",
		"NULL", "DIGIT", "OCTET", "IPV4", "IPV4_TYPE", "ARRAY", "ARRAY_TYPE",
		"LOOP", "IF", "ELSE", "FUNC", "RETURN", "RETURN_TYPE", "L_BLOCK", "R_BLOCK",
		"L_PAREN", "R_PAREN", "L_SQUARE", "R_SQUARE", "DOT", "COLON", "SEMICOLON",
		"COMMA", "AT", "OP_PLUS", "OP_MINUS", "OP_MULTIPLY", "OP_DIVIDE", "OP_MOD",
		"OP_POWER", "OP_EQUAL", "OP_NOT_EQUAL", "OP_LESS", "OP_GREATER", "OP_LESS_EQUAL",
		"OP_GREATER_EQUAL", "OP_AND", "OP_OR", "OP_XOR", "OP_NOT", "OP_INCREMENT",
		"OP_DECREMENT", "VAR_IDENTIFIER", "COMMON_IDENTIFIER", "NEWLINE", "WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "expression", "block", "loop_statement", "if_statement",
		"return_statement", "data_type", "func_declaration", "var_type", "var_name",
		"var_value", "op_one_expression", "op_one_declaration", "string_var_declaration",
		"number_var_declaration", "boolean_var_declaration", "single_var_declaration",
		"array_var_declaration", "var_assignment", "op_level1", "op_level2",
		"op_level3", "op_level4", "op_level5", "numberic_expression", "logical_expression",
		"comparative_expression", "get_variable", "math_expression", "get_array_element",
		"get_child", "function_call",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 52, 522, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 1, 0, 5, 0, 68, 8, 0, 10, 0, 12, 0, 71, 9, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 81, 8, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 88, 8, 2, 1, 3, 1, 3, 5, 3, 92, 8, 3, 10, 3, 12, 3, 95,
		9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 3, 5, 111, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 126, 8, 8, 10, 8, 12, 8,
		129, 9, 8, 3, 8, 131, 8, 8, 1, 8, 1, 8, 1, 8, 3, 8, 136, 8, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 166, 8, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 5, 14, 172, 8, 14, 10, 14, 12, 14, 175, 9, 14, 1, 14, 1,
		14, 5, 14, 179, 8, 14, 10, 14, 12, 14, 182, 9, 14, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 15, 1, 15, 5, 15, 190, 8, 15, 10, 15, 12, 15, 193, 9, 15, 1,
		15, 1, 15, 5, 15, 197, 8, 15, 10, 15, 12, 15, 200, 9, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 208, 8, 15, 10, 15, 12, 15, 211, 9,
		15, 1, 15, 1, 15, 5, 15, 215, 8, 15, 10, 15, 12, 15, 218, 9, 15, 1, 15,
		1, 15, 3, 15, 222, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 228, 8, 16,
		10, 16, 12, 16, 231, 9, 16, 1, 16, 1, 16, 5, 16, 235, 8, 16, 10, 16, 12,
		16, 238, 9, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 246, 8,
		16, 10, 16, 12, 16, 249, 9, 16, 1, 16, 1, 16, 5, 16, 253, 8, 16, 10, 16,
		12, 16, 256, 9, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 264,
		8, 16, 10, 16, 12, 16, 267, 9, 16, 1, 16, 1, 16, 5, 16, 271, 8, 16, 10,
		16, 12, 16, 274, 9, 16, 1, 16, 1, 16, 3, 16, 278, 8, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 287, 8, 17, 10, 17, 12, 17, 290,
		9, 17, 1, 17, 1, 17, 3, 17, 294, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 5,
		18, 300, 8, 18, 10, 18, 12, 18, 303, 9, 18, 1, 18, 1, 18, 1, 18, 5, 18,
		308, 8, 18, 10, 18, 12, 18, 311, 9, 18, 1, 18, 5, 18, 314, 8, 18, 10, 18,
		12, 18, 317, 9, 18, 1, 18, 5, 18, 320, 8, 18, 10, 18, 12, 18, 323, 9, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 331, 8, 18, 10, 18, 12,
		18, 334, 9, 18, 1, 18, 1, 18, 1, 18, 5, 18, 339, 8, 18, 10, 18, 12, 18,
		342, 9, 18, 1, 18, 5, 18, 345, 8, 18, 10, 18, 12, 18, 348, 9, 18, 1, 18,
		1, 18, 1, 18, 3, 18, 353, 8, 18, 1, 19, 1, 19, 1, 19, 5, 19, 358, 8, 19,
		10, 19, 12, 19, 361, 9, 19, 1, 19, 1, 19, 5, 19, 365, 8, 19, 10, 19, 12,
		19, 368, 9, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 375, 8, 19, 10,
		19, 12, 19, 378, 9, 19, 1, 19, 1, 19, 5, 19, 382, 8, 19, 10, 19, 12, 19,
		385, 9, 19, 1, 19, 1, 19, 3, 19, 389, 8, 19, 1, 20, 1, 20, 1, 21, 1, 21,
		1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 411, 8, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 421, 8, 25, 10, 25, 12,
		25, 424, 9, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 3, 26, 436, 8, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 5, 26, 446, 8, 26, 10, 26, 12, 26, 449, 9, 26, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 3, 27, 465, 8, 27, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27,
		471, 8, 27, 10, 27, 12, 27, 474, 9, 27, 1, 28, 1, 28, 1, 28, 3, 28, 479,
		8, 28, 1, 29, 1, 29, 1, 29, 3, 29, 484, 8, 29, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 506, 8, 31, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 5, 32, 513, 8, 32, 10, 32, 12, 32, 516, 9, 32,
		3, 32, 518, 8, 32, 1, 32, 1, 32, 1, 32, 0, 3, 50, 52, 54, 33, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 0, 8, 4, 0, 2, 2, 4, 4, 6,
		6, 50, 50, 4, 0, 2, 2, 4, 4, 6, 6, 11, 11, 4, 0, 1, 1, 3, 3, 5, 5, 10,
		10, 1, 0, 47, 48, 1, 0, 33, 36, 1, 0, 31, 32, 1, 0, 43, 45, 1, 0, 37, 42,
		566, 0, 69, 1, 0, 0, 0, 2, 80, 1, 0, 0, 0, 4, 87, 1, 0, 0, 0, 6, 89, 1,
		0, 0, 0, 8, 98, 1, 0, 0, 0, 10, 103, 1, 0, 0, 0, 12, 112, 1, 0, 0, 0, 14,
		116, 1, 0, 0, 0, 16, 118, 1, 0, 0, 0, 18, 139, 1, 0, 0, 0, 20, 141, 1,
		0, 0, 0, 22, 143, 1, 0, 0, 0, 24, 145, 1, 0, 0, 0, 26, 165, 1, 0, 0, 0,
		28, 167, 1, 0, 0, 0, 30, 221, 1, 0, 0, 0, 32, 277, 1, 0, 0, 0, 34, 293,
		1, 0, 0, 0, 36, 352, 1, 0, 0, 0, 38, 388, 1, 0, 0, 0, 40, 390, 1, 0, 0,
		0, 42, 392, 1, 0, 0, 0, 44, 394, 1, 0, 0, 0, 46, 396, 1, 0, 0, 0, 48, 398,
		1, 0, 0, 0, 50, 410, 1, 0, 0, 0, 52, 435, 1, 0, 0, 0, 54, 464, 1, 0, 0,
		0, 56, 478, 1, 0, 0, 0, 58, 483, 1, 0, 0, 0, 60, 485, 1, 0, 0, 0, 62, 505,
		1, 0, 0, 0, 64, 507, 1, 0, 0, 0, 66, 68, 3, 2, 1, 0, 67, 66, 1, 0, 0, 0,
		68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 72, 1,
		0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 73, 5, 0, 0, 1, 73, 1, 1, 0, 0, 0, 74,
		81, 3, 4, 2, 0, 75, 81, 3, 8, 4, 0, 76, 81, 3, 10, 5, 0, 77, 81, 3, 16,
		8, 0, 78, 81, 3, 12, 6, 0, 79, 81, 5, 51, 0, 0, 80, 74, 1, 0, 0, 0, 80,
		75, 1, 0, 0, 0, 80, 76, 1, 0, 0, 0, 80, 77, 1, 0, 0, 0, 80, 78, 1, 0, 0,
		0, 80, 79, 1, 0, 0, 0, 81, 3, 1, 0, 0, 0, 82, 88, 3, 64, 32, 0, 83, 88,
		3, 38, 19, 0, 84, 88, 3, 34, 17, 0, 85, 88, 3, 36, 18, 0, 86, 88, 3, 56,
		28, 0, 87, 82, 1, 0, 0, 0, 87, 83, 1, 0, 0, 0, 87, 84, 1, 0, 0, 0, 87,
		85, 1, 0, 0, 0, 87, 86, 1, 0, 0, 0, 88, 5, 1, 0, 0, 0, 89, 93, 5, 20, 0,
		0, 90, 92, 3, 2, 1, 0, 91, 90, 1, 0, 0, 0, 92, 95, 1, 0, 0, 0, 93, 91,
		1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 96, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0,
		96, 97, 5, 21, 0, 0, 97, 7, 1, 0, 0, 0, 98, 99, 5, 30, 0, 0, 99, 100, 5,
		14, 0, 0, 100, 101, 3, 54, 27, 0, 101, 102, 3, 6, 3, 0, 102, 9, 1, 0, 0,
		0, 103, 104, 5, 30, 0, 0, 104, 105, 5, 15, 0, 0, 105, 106, 3, 54, 27, 0,
		106, 110, 3, 6, 3, 0, 107, 108, 5, 30, 0, 0, 108, 109, 5, 16, 0, 0, 109,
		111, 3, 6, 3, 0, 110, 107, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 11, 1,
		0, 0, 0, 112, 113, 5, 30, 0, 0, 113, 114, 5, 18, 0, 0, 114, 115, 3, 4,
		2, 0, 115, 13, 1, 0, 0, 0, 116, 117, 7, 0, 0, 0, 117, 15, 1, 0, 0, 0, 118,
		119, 5, 30, 0, 0, 119, 120, 5, 17, 0, 0, 120, 121, 5, 49, 0, 0, 121, 130,
		5, 22, 0, 0, 122, 127, 5, 49, 0, 0, 123, 124, 5, 29, 0, 0, 124, 126, 5,
		49, 0, 0, 125, 123, 1, 0, 0, 0, 126, 129, 1, 0, 0, 0, 127, 125, 1, 0, 0,
		0, 127, 128, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 130,
		122, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 135,
		5, 23, 0, 0, 133, 134, 5, 19, 0, 0, 134, 136, 3, 14, 7, 0, 135, 133, 1,
		0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 138, 3, 6, 3,
		0, 138, 17, 1, 0, 0, 0, 139, 140, 7, 1, 0, 0, 140, 19, 1, 0, 0, 0, 141,
		142, 5, 49, 0, 0, 142, 21, 1, 0, 0, 0, 143, 144, 7, 2, 0, 0, 144, 23, 1,
		0, 0, 0, 145, 146, 7, 3, 0, 0, 146, 25, 1, 0, 0, 0, 147, 148, 3, 20, 10,
		0, 148, 149, 5, 47, 0, 0, 149, 166, 1, 0, 0, 0, 150, 151, 3, 20, 10, 0,
		151, 152, 5, 48, 0, 0, 152, 166, 1, 0, 0, 0, 153, 154, 5, 47, 0, 0, 154,
		166, 3, 20, 10, 0, 155, 156, 5, 48, 0, 0, 156, 166, 3, 20, 10, 0, 157,
		158, 5, 6, 0, 0, 158, 159, 3, 20, 10, 0, 159, 160, 5, 47, 0, 0, 160, 166,
		1, 0, 0, 0, 161, 162, 5, 6, 0, 0, 162, 163, 3, 20, 10, 0, 163, 164, 5,
		48, 0, 0, 164, 166, 1, 0, 0, 0, 165, 147, 1, 0, 0, 0, 165, 150, 1, 0, 0,
		0, 165, 153, 1, 0, 0, 0, 165, 155, 1, 0, 0, 0, 165, 157, 1, 0, 0, 0, 165,
		161, 1, 0, 0, 0, 166, 27, 1, 0, 0, 0, 167, 168, 5, 2, 0, 0, 168, 169, 3,
		20, 10, 0, 169, 173, 5, 20, 0, 0, 170, 172, 5, 51, 0, 0, 171, 170, 1, 0,
		0, 0, 172, 175, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0,
		174, 176, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 176, 180, 5, 1, 0, 0, 177,
		179, 5, 51, 0, 0, 178, 177, 1, 0, 0, 0, 179, 182, 1, 0, 0, 0, 180, 178,
		1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 183, 1, 0, 0, 0, 182, 180, 1, 0,
		0, 0, 183, 184, 5, 21, 0, 0, 184, 29, 1, 0, 0, 0, 185, 186, 5, 6, 0, 0,
		186, 187, 3, 20, 10, 0, 187, 191, 5, 20, 0, 0, 188, 190, 5, 51, 0, 0, 189,
		188, 1, 0, 0, 0, 190, 193, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 191, 192,
		1, 0, 0, 0, 192, 194, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 194, 198, 5, 5,
		0, 0, 195, 197, 5, 51, 0, 0, 196, 195, 1, 0, 0, 0, 197, 200, 1, 0, 0, 0,
		198, 196, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 201, 1, 0, 0, 0, 200,
		198, 1, 0, 0, 0, 201, 202, 5, 21, 0, 0, 202, 222, 1, 0, 0, 0, 203, 204,
		5, 6, 0, 0, 204, 205, 3, 20, 10, 0, 205, 209, 5, 20, 0, 0, 206, 208, 5,
		51, 0, 0, 207, 206, 1, 0, 0, 0, 208, 211, 1, 0, 0, 0, 209, 207, 1, 0, 0,
		0, 209, 210, 1, 0, 0, 0, 210, 212, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 212,
		216, 3, 58, 29, 0, 213, 215, 5, 51, 0, 0, 214, 213, 1, 0, 0, 0, 215, 218,
		1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 219, 1, 0,
		0, 0, 218, 216, 1, 0, 0, 0, 219, 220, 5, 21, 0, 0, 220, 222, 1, 0, 0, 0,
		221, 185, 1, 0, 0, 0, 221, 203, 1, 0, 0, 0, 222, 31, 1, 0, 0, 0, 223, 224,
		5, 4, 0, 0, 224, 225, 3, 20, 10, 0, 225, 229, 5, 20, 0, 0, 226, 228, 5,
		51, 0, 0, 227, 226, 1, 0, 0, 0, 228, 231, 1, 0, 0, 0, 229, 227, 1, 0, 0,
		0, 229, 230, 1, 0, 0, 0, 230, 232, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 232,
		236, 5, 3, 0, 0, 233, 235, 5, 51, 0, 0, 234, 233, 1, 0, 0, 0, 235, 238,
		1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 239, 1, 0,
		0, 0, 238, 236, 1, 0, 0, 0, 239, 240, 5, 21, 0, 0, 240, 278, 1, 0, 0, 0,
		241, 242, 5, 4, 0, 0, 242, 243, 3, 20, 10, 0, 243, 247, 5, 20, 0, 0, 244,
		246, 5, 51, 0, 0, 245, 244, 1, 0, 0, 0, 246, 249, 1, 0, 0, 0, 247, 245,
		1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 250, 1, 0, 0, 0, 249, 247, 1, 0,
		0, 0, 250, 254, 3, 52, 26, 0, 251, 253, 5, 51, 0, 0, 252, 251, 1, 0, 0,
		0, 253, 256, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255,
		257, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0, 257, 258, 5, 21, 0, 0, 258, 278,
		1, 0, 0, 0, 259, 260, 5, 4, 0, 0, 260, 261, 3, 20, 10, 0, 261, 265, 5,
		20, 0, 0, 262, 264, 5, 51, 0, 0, 263, 262, 1, 0, 0, 0, 264, 267, 1, 0,
		0, 0, 265, 263, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 268, 1, 0, 0, 0,
		267, 265, 1, 0, 0, 0, 268, 272, 3, 54, 27, 0, 269, 271, 5, 51, 0, 0, 270,
		269, 1, 0, 0, 0, 271, 274, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 272, 273,
		1, 0, 0, 0, 273, 275, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 275, 276, 5, 21,
		0, 0, 276, 278, 1, 0, 0, 0, 277, 223, 1, 0, 0, 0, 277, 241, 1, 0, 0, 0,
		277, 259, 1, 0, 0, 0, 278, 33, 1, 0, 0, 0, 279, 294, 3, 28, 14, 0, 280,
		294, 3, 30, 15, 0, 281, 294, 3, 32, 16, 0, 282, 283, 3, 18, 9, 0, 283,
		284, 3, 20, 10, 0, 284, 288, 5, 20, 0, 0, 285, 287, 5, 51, 0, 0, 286, 285,
		1, 0, 0, 0, 287, 290, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 289, 1, 0,
		0, 0, 289, 291, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 291, 292, 5, 21, 0, 0,
		292, 294, 1, 0, 0, 0, 293, 279, 1, 0, 0, 0, 293, 280, 1, 0, 0, 0, 293,
		281, 1, 0, 0, 0, 293, 282, 1, 0, 0, 0, 294, 35, 1, 0, 0, 0, 295, 296, 3,
		18, 9, 0, 296, 297, 3, 20, 10, 0, 297, 301, 5, 20, 0, 0, 298, 300, 5, 51,
		0, 0, 299, 298, 1, 0, 0, 0, 300, 303, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0,
		301, 302, 1, 0, 0, 0, 302, 304, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 304,
		315, 3, 22, 11, 0, 305, 309, 5, 29, 0, 0, 306, 308, 5, 51, 0, 0, 307, 306,
		1, 0, 0, 0, 308, 311, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 309, 310, 1, 0,
		0, 0, 310, 312, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 312, 314, 3, 22, 11,
		0, 313, 305, 1, 0, 0, 0, 314, 317, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 315,
		316, 1, 0, 0, 0, 316, 321, 1, 0, 0, 0, 317, 315, 1, 0, 0, 0, 318, 320,
		5, 51, 0, 0, 319, 318, 1, 0, 0, 0, 320, 323, 1, 0, 0, 0, 321, 319, 1, 0,
		0, 0, 321, 322, 1, 0, 0, 0, 322, 324, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0,
		324, 325, 5, 21, 0, 0, 325, 353, 1, 0, 0, 0, 326, 327, 3, 18, 9, 0, 327,
		328, 3, 20, 10, 0, 328, 332, 5, 20, 0, 0, 329, 331, 5, 51, 0, 0, 330, 329,
		1, 0, 0, 0, 331, 334, 1, 0, 0, 0, 332, 330, 1, 0, 0, 0, 332, 333, 1, 0,
		0, 0, 333, 335, 1, 0, 0, 0, 334, 332, 1, 0, 0, 0, 335, 346, 3, 58, 29,
		0, 336, 340, 5, 29, 0, 0, 337, 339, 5, 51, 0, 0, 338, 337, 1, 0, 0, 0,
		339, 342, 1, 0, 0, 0, 340, 338, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341,
		343, 1, 0, 0, 0, 342, 340, 1, 0, 0, 0, 343, 345, 3, 58, 29, 0, 344, 336,
		1, 0, 0, 0, 345, 348, 1, 0, 0, 0, 346, 344, 1, 0, 0, 0, 346, 347, 1, 0,
		0, 0, 347, 349, 1, 0, 0, 0, 348, 346, 1, 0, 0, 0, 349, 350, 5, 51, 0, 0,
		350, 351, 5, 21, 0, 0, 351, 353, 1, 0, 0, 0, 352, 295, 1, 0, 0, 0, 352,
		326, 1, 0, 0, 0, 353, 37, 1, 0, 0, 0, 354, 355, 3, 20, 10, 0, 355, 359,
		5, 20, 0, 0, 356, 358, 5, 51, 0, 0, 357, 356, 1, 0, 0, 0, 358, 361, 1,
		0, 0, 0, 359, 357, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0, 360, 362, 1, 0, 0,
		0, 361, 359, 1, 0, 0, 0, 362, 366, 3, 22, 11, 0, 363, 365, 5, 51, 0, 0,
		364, 363, 1, 0, 0, 0, 365, 368, 1, 0, 0, 0, 366, 364, 1, 0, 0, 0, 366,
		367, 1, 0, 0, 0, 367, 369, 1, 0, 0, 0, 368, 366, 1, 0, 0, 0, 369, 370,
		5, 21, 0, 0, 370, 389, 1, 0, 0, 0, 371, 372, 3, 20, 10, 0, 372, 376, 5,
		20, 0, 0, 373, 375, 5, 51, 0, 0, 374, 373, 1, 0, 0, 0, 375, 378, 1, 0,
		0, 0, 376, 374, 1, 0, 0, 0, 376, 377, 1, 0, 0, 0, 377, 379, 1, 0, 0, 0,
		378, 376, 1, 0, 0, 0, 379, 383, 3, 58, 29, 0, 380, 382, 5, 51, 0, 0, 381,
		380, 1, 0, 0, 0, 382, 385, 1, 0, 0, 0, 383, 381, 1, 0, 0, 0, 383, 384,
		1, 0, 0, 0, 384, 386, 1, 0, 0, 0, 385, 383, 1, 0, 0, 0, 386, 387, 5, 21,
		0, 0, 387, 389, 1, 0, 0, 0, 388, 354, 1, 0, 0, 0, 388, 371, 1, 0, 0, 0,
		389, 39, 1, 0, 0, 0, 390, 391, 7, 4, 0, 0, 391, 41, 1, 0, 0, 0, 392, 393,
		7, 5, 0, 0, 393, 43, 1, 0, 0, 0, 394, 395, 7, 6, 0, 0, 395, 45, 1, 0, 0,
		0, 396, 397, 7, 7, 0, 0, 397, 47, 1, 0, 0, 0, 398, 399, 5, 46, 0, 0, 399,
		49, 1, 0, 0, 0, 400, 401, 6, 25, -1, 0, 401, 402, 5, 22, 0, 0, 402, 403,
		3, 50, 25, 0, 403, 404, 5, 23, 0, 0, 404, 411, 1, 0, 0, 0, 405, 411, 5,
		5, 0, 0, 406, 411, 3, 56, 28, 0, 407, 411, 3, 64, 32, 0, 408, 409, 5, 32,
		0, 0, 409, 411, 3, 50, 25, 1, 410, 400, 1, 0, 0, 0, 410, 405, 1, 0, 0,
		0, 410, 406, 1, 0, 0, 0, 410, 407, 1, 0, 0, 0, 410, 408, 1, 0, 0, 0, 411,
		422, 1, 0, 0, 0, 412, 413, 10, 6, 0, 0, 413, 414, 3, 40, 20, 0, 414, 415,
		3, 50, 25, 7, 415, 421, 1, 0, 0, 0, 416, 417, 10, 5, 0, 0, 417, 418, 3,
		42, 21, 0, 418, 419, 3, 50, 25, 6, 419, 421, 1, 0, 0, 0, 420, 412, 1, 0,
		0, 0, 420, 416, 1, 0, 0, 0, 421, 424, 1, 0, 0, 0, 422, 420, 1, 0, 0, 0,
		422, 423, 1, 0, 0, 0, 423, 51, 1, 0, 0, 0, 424, 422, 1, 0, 0, 0, 425, 426,
		6, 26, -1, 0, 426, 427, 5, 22, 0, 0, 427, 428, 3, 52, 26, 0, 428, 429,
		5, 23, 0, 0, 429, 436, 1, 0, 0, 0, 430, 431, 5, 46, 0, 0, 431, 436, 3,
		52, 26, 4, 432, 436, 5, 3, 0, 0, 433, 436, 3, 56, 28, 0, 434, 436, 3, 64,
		32, 0, 435, 425, 1, 0, 0, 0, 435, 430, 1, 0, 0, 0, 435, 432, 1, 0, 0, 0,
		435, 433, 1, 0, 0, 0, 435, 434, 1, 0, 0, 0, 436, 447, 1, 0, 0, 0, 437,
		438, 10, 6, 0, 0, 438, 439, 3, 44, 22, 0, 439, 440, 3, 52, 26, 7, 440,
		446, 1, 0, 0, 0, 441, 442, 10, 5, 0, 0, 442, 443, 3, 46, 23, 0, 443, 444,
		3, 52, 26, 6, 444, 446, 1, 0, 0, 0, 445, 437, 1, 0, 0, 0, 445, 441, 1,
		0, 0, 0, 446, 449, 1, 0, 0, 0, 447, 445, 1, 0, 0, 0, 447, 448, 1, 0, 0,
		0, 448, 53, 1, 0, 0, 0, 449, 447, 1, 0, 0, 0, 450, 451, 6, 27, -1, 0, 451,
		452, 3, 50, 25, 0, 452, 453, 3, 46, 23, 0, 453, 454, 3, 50, 25, 0, 454,
		465, 1, 0, 0, 0, 455, 456, 3, 52, 26, 0, 456, 457, 3, 44, 22, 0, 457, 458,
		3, 52, 26, 0, 458, 465, 1, 0, 0, 0, 459, 460, 3, 48, 24, 0, 460, 461, 5,
		22, 0, 0, 461, 462, 3, 54, 27, 0, 462, 463, 5, 23, 0, 0, 463, 465, 1, 0,
		0, 0, 464, 450, 1, 0, 0, 0, 464, 455, 1, 0, 0, 0, 464, 459, 1, 0, 0, 0,
		465, 472, 1, 0, 0, 0, 466, 467, 10, 2, 0, 0, 467, 468, 3, 44, 22, 0, 468,
		469, 3, 52, 26, 0, 469, 471, 1, 0, 0, 0, 470, 466, 1, 0, 0, 0, 471, 474,
		1, 0, 0, 0, 472, 470, 1, 0, 0, 0, 472, 473, 1, 0, 0, 0, 473, 55, 1, 0,
		0, 0, 474, 472, 1, 0, 0, 0, 475, 479, 5, 49, 0, 0, 476, 479, 3, 60, 30,
		0, 477, 479, 3, 62, 31, 0, 478, 475, 1, 0, 0, 0, 478, 476, 1, 0, 0, 0,
		478, 477, 1, 0, 0, 0, 479, 57, 1, 0, 0, 0, 480, 484, 3, 56, 28, 0, 481,
		484, 3, 50, 25, 0, 482, 484, 3, 52, 26, 0, 483, 480, 1, 0, 0, 0, 483, 481,
		1, 0, 0, 0, 483, 482, 1, 0, 0, 0, 484, 59, 1, 0, 0, 0, 485, 486, 5, 49,
		0, 0, 486, 487, 5, 24, 0, 0, 487, 488, 3, 50, 25, 0, 488, 489, 5, 25, 0,
		0, 489, 61, 1, 0, 0, 0, 490, 491, 5, 49, 0, 0, 491, 492, 5, 26, 0, 0, 492,
		506, 5, 49, 0, 0, 493, 494, 5, 49, 0, 0, 494, 495, 5, 26, 0, 0, 495, 506,
		3, 60, 30, 0, 496, 497, 5, 49, 0, 0, 497, 498, 5, 26, 0, 0, 498, 506, 3,
		62, 31, 0, 499, 500, 3, 60, 30, 0, 500, 501, 5, 26, 0, 0, 501, 502, 3,
		62, 31, 0, 502, 506, 1, 0, 0, 0, 503, 506, 3, 60, 30, 0, 504, 506, 5, 49,
		0, 0, 505, 490, 1, 0, 0, 0, 505, 493, 1, 0, 0, 0, 505, 496, 1, 0, 0, 0,
		505, 499, 1, 0, 0, 0, 505, 503, 1, 0, 0, 0, 505, 504, 1, 0, 0, 0, 506,
		63, 1, 0, 0, 0, 507, 508, 5, 49, 0, 0, 508, 517, 5, 22, 0, 0, 509, 514,
		3, 58, 29, 0, 510, 511, 5, 29, 0, 0, 511, 513, 3, 58, 29, 0, 512, 510,
		1, 0, 0, 0, 513, 516, 1, 0, 0, 0, 514, 512, 1, 0, 0, 0, 514, 515, 1, 0,
		0, 0, 515, 518, 1, 0, 0, 0, 516, 514, 1, 0, 0, 0, 517, 509, 1, 0, 0, 0,
		517, 518, 1, 0, 0, 0, 518, 519, 1, 0, 0, 0, 519, 520, 5, 23, 0, 0, 520,
		65, 1, 0, 0, 0, 51, 69, 80, 87, 93, 110, 127, 130, 135, 165, 173, 180,
		191, 198, 209, 216, 221, 229, 236, 247, 254, 265, 272, 277, 288, 293, 301,
		309, 315, 321, 332, 340, 346, 352, 359, 366, 376, 383, 388, 410, 420, 422,
		435, 445, 447, 464, 472, 478, 483, 505, 514, 517,
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

// FluxInit initializes any static state used to implement Flux. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFlux(). You can call this function if you wish to initialize the static state ahead
// of time.
func FluxInit() {
	staticData := &FluxParserStaticData
	staticData.once.Do(fluxParserInit)
}

// NewFlux produces a new parser instance for the optional input antlr.TokenStream.
func NewFlux(input antlr.TokenStream) *Flux {
	FluxInit()
	this := new(Flux)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &FluxParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Flux.g4"

	return this
}

// Flux tokens.
const (
	FluxEOF               = antlr.TokenEOF
	FluxTEXT              = 1
	FluxTEXT_TYPE         = 2
	FluxBOOLEAN           = 3
	FluxBOOLEAN_TYPE      = 4
	FluxNUMBER            = 5
	FluxNUMBER_TYPE       = 6
	FluxNULL              = 7
	FluxDIGIT             = 8
	FluxOCTET             = 9
	FluxIPV4              = 10
	FluxIPV4_TYPE         = 11
	FluxARRAY             = 12
	FluxARRAY_TYPE        = 13
	FluxLOOP              = 14
	FluxIF                = 15
	FluxELSE              = 16
	FluxFUNC              = 17
	FluxRETURN            = 18
	FluxRETURN_TYPE       = 19
	FluxL_BLOCK           = 20
	FluxR_BLOCK           = 21
	FluxL_PAREN           = 22
	FluxR_PAREN           = 23
	FluxL_SQUARE          = 24
	FluxR_SQUARE          = 25
	FluxDOT               = 26
	FluxCOLON             = 27
	FluxSEMICOLON         = 28
	FluxCOMMA             = 29
	FluxAT                = 30
	FluxOP_PLUS           = 31
	FluxOP_MINUS          = 32
	FluxOP_MULTIPLY       = 33
	FluxOP_DIVIDE         = 34
	FluxOP_MOD            = 35
	FluxOP_POWER          = 36
	FluxOP_EQUAL          = 37
	FluxOP_NOT_EQUAL      = 38
	FluxOP_LESS           = 39
	FluxOP_GREATER        = 40
	FluxOP_LESS_EQUAL     = 41
	FluxOP_GREATER_EQUAL  = 42
	FluxOP_AND            = 43
	FluxOP_OR             = 44
	FluxOP_XOR            = 45
	FluxOP_NOT            = 46
	FluxOP_INCREMENT      = 47
	FluxOP_DECREMENT      = 48
	FluxVAR_IDENTIFIER    = 49
	FluxCOMMON_IDENTIFIER = 50
	FluxNEWLINE           = 51
	FluxWS                = 52
)

// Flux rules.
const (
	FluxRULE_program                 = 0
	FluxRULE_statement               = 1
	FluxRULE_expression              = 2
	FluxRULE_block                   = 3
	FluxRULE_loop_statement          = 4
	FluxRULE_if_statement            = 5
	FluxRULE_return_statement        = 6
	FluxRULE_data_type               = 7
	FluxRULE_func_declaration        = 8
	FluxRULE_var_type                = 9
	FluxRULE_var_name                = 10
	FluxRULE_var_value               = 11
	FluxRULE_op_one_expression       = 12
	FluxRULE_op_one_declaration      = 13
	FluxRULE_string_var_declaration  = 14
	FluxRULE_number_var_declaration  = 15
	FluxRULE_boolean_var_declaration = 16
	FluxRULE_single_var_declaration  = 17
	FluxRULE_array_var_declaration   = 18
	FluxRULE_var_assignment          = 19
	FluxRULE_op_level1               = 20
	FluxRULE_op_level2               = 21
	FluxRULE_op_level3               = 22
	FluxRULE_op_level4               = 23
	FluxRULE_op_level5               = 24
	FluxRULE_numberic_expression     = 25
	FluxRULE_logical_expression      = 26
	FluxRULE_comparative_expression  = 27
	FluxRULE_get_variable            = 28
	FluxRULE_math_expression         = 29
	FluxRULE_get_array_element       = 30
	FluxRULE_get_child               = 31
	FluxRULE_function_call           = 32
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(FluxEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FluxRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2814750840850516) != 0 {
		{
			p.SetState(66)
			p.Statement()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(72)
		p.Match(FluxEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Loop_statement() ILoop_statementContext
	If_statement() IIf_statementContext
	Func_declaration() IFunc_declarationContext
	Return_statement() IReturn_statementContext
	NEWLINE() antlr.TerminalNode

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) Loop_statement() ILoop_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoop_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoop_statementContext)
}

func (s *StatementContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *StatementContext) Func_declaration() IFunc_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_declarationContext)
}

func (s *StatementContext) Return_statement() IReturn_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturn_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturn_statementContext)
}

func (s *StatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(FluxNEWLINE, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FluxRULE_statement)
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.Loop_statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(76)
			p.If_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(77)
			p.Func_declaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(78)
			p.Return_statement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(79)
			p.Match(FluxNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function_call() IFunction_callContext
	Var_assignment() IVar_assignmentContext
	Single_var_declaration() ISingle_var_declarationContext
	Array_var_declaration() IArray_var_declarationContext
	Get_variable() IGet_variableContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *ExpressionContext) Var_assignment() IVar_assignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_assignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_assignmentContext)
}

func (s *ExpressionContext) Single_var_declaration() ISingle_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingle_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingle_var_declarationContext)
}

func (s *ExpressionContext) Array_var_declaration() IArray_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_var_declarationContext)
}

func (s *ExpressionContext) Get_variable() IGet_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_variableContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FluxRULE_expression)
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Function_call()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Var_assignment()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(84)
			p.Single_var_declaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(85)
			p.Array_var_declaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(86)
			p.Get_variable()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_BLOCK() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxL_BLOCK, 0)
}

func (s *BlockContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxR_BLOCK, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FluxRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(FluxL_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2814750840850516) != 0 {
		{
			p.SetState(90)
			p.Statement()
		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(96)
		p.Match(FluxR_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoop_statementContext is an interface to support dynamic dispatch.
type ILoop_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	LOOP() antlr.TerminalNode
	Comparative_expression() IComparative_expressionContext
	Block() IBlockContext

	// IsLoop_statementContext differentiates from other interfaces.
	IsLoop_statementContext()
}

type Loop_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoop_statementContext() *Loop_statementContext {
	var p = new(Loop_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_loop_statement
	return p
}

func InitEmptyLoop_statementContext(p *Loop_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_loop_statement
}

func (*Loop_statementContext) IsLoop_statementContext() {}

func NewLoop_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Loop_statementContext {
	var p = new(Loop_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_loop_statement

	return p
}

func (s *Loop_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Loop_statementContext) AT() antlr.TerminalNode {
	return s.GetToken(FluxAT, 0)
}

func (s *Loop_statementContext) LOOP() antlr.TerminalNode {
	return s.GetToken(FluxLOOP, 0)
}

func (s *Loop_statementContext) Comparative_expression() IComparative_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparative_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparative_expressionContext)
}

func (s *Loop_statementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Loop_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Loop_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Loop_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterLoop_statement(s)
	}
}

func (s *Loop_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitLoop_statement(s)
	}
}

func (s *Loop_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitLoop_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Loop_statement() (localctx ILoop_statementContext) {
	localctx = NewLoop_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FluxRULE_loop_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(FluxAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)
		p.Match(FluxLOOP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.comparative_expression(0)
	}
	{
		p.SetState(101)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_statementContext is an interface to support dynamic dispatch.
type IIf_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAT() []antlr.TerminalNode
	AT(i int) antlr.TerminalNode
	IF() antlr.TerminalNode
	Comparative_expression() IComparative_expressionContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	ELSE() antlr.TerminalNode

	// IsIf_statementContext differentiates from other interfaces.
	IsIf_statementContext()
}

type If_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_statementContext() *If_statementContext {
	var p = new(If_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_if_statement
	return p
}

func InitEmptyIf_statementContext(p *If_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_if_statement
}

func (*If_statementContext) IsIf_statementContext() {}

func NewIf_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statementContext {
	var p = new(If_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_if_statement

	return p
}

func (s *If_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statementContext) AllAT() []antlr.TerminalNode {
	return s.GetTokens(FluxAT)
}

func (s *If_statementContext) AT(i int) antlr.TerminalNode {
	return s.GetToken(FluxAT, i)
}

func (s *If_statementContext) IF() antlr.TerminalNode {
	return s.GetToken(FluxIF, 0)
}

func (s *If_statementContext) Comparative_expression() IComparative_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparative_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparative_expressionContext)
}

func (s *If_statementContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *If_statementContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *If_statementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(FluxELSE, 0)
}

func (s *If_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterIf_statement(s)
	}
}

func (s *If_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitIf_statement(s)
	}
}

func (s *If_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitIf_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) If_statement() (localctx IIf_statementContext) {
	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FluxRULE_if_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(FluxAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.Match(FluxIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.comparative_expression(0)
	}
	{
		p.SetState(106)
		p.Block()
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(107)
			p.Match(FluxAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.Match(FluxELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Block()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturn_statementContext is an interface to support dynamic dispatch.
type IReturn_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	RETURN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturn_statementContext differentiates from other interfaces.
	IsReturn_statementContext()
}

type Return_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_statementContext() *Return_statementContext {
	var p = new(Return_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_return_statement
	return p
}

func InitEmptyReturn_statementContext(p *Return_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_return_statement
}

func (*Return_statementContext) IsReturn_statementContext() {}

func NewReturn_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_statementContext {
	var p = new(Return_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_return_statement

	return p
}

func (s *Return_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_statementContext) AT() antlr.TerminalNode {
	return s.GetToken(FluxAT, 0)
}

func (s *Return_statementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(FluxRETURN, 0)
}

func (s *Return_statementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Return_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterReturn_statement(s)
	}
}

func (s *Return_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitReturn_statement(s)
	}
}

func (s *Return_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitReturn_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Return_statement() (localctx IReturn_statementContext) {
	localctx = NewReturn_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FluxRULE_return_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(FluxAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Match(FluxRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IData_typeContext is an interface to support dynamic dispatch.
type IData_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT_TYPE() antlr.TerminalNode
	NUMBER_TYPE() antlr.TerminalNode
	BOOLEAN_TYPE() antlr.TerminalNode
	COMMON_IDENTIFIER() antlr.TerminalNode

	// IsData_typeContext differentiates from other interfaces.
	IsData_typeContext()
}

type Data_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyData_typeContext() *Data_typeContext {
	var p = new(Data_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_data_type
	return p
}

func InitEmptyData_typeContext(p *Data_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_data_type
}

func (*Data_typeContext) IsData_typeContext() {}

func NewData_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_typeContext {
	var p = new(Data_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_data_type

	return p
}

func (s *Data_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_typeContext) TEXT_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxTEXT_TYPE, 0)
}

func (s *Data_typeContext) NUMBER_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxNUMBER_TYPE, 0)
}

func (s *Data_typeContext) BOOLEAN_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxBOOLEAN_TYPE, 0)
}

func (s *Data_typeContext) COMMON_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FluxCOMMON_IDENTIFIER, 0)
}

func (s *Data_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterData_type(s)
	}
}

func (s *Data_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitData_type(s)
	}
}

func (s *Data_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitData_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Data_type() (localctx IData_typeContext) {
	localctx = NewData_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FluxRULE_data_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1125899906842708) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_declarationContext is an interface to support dynamic dispatch.
type IFunc_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	FUNC() antlr.TerminalNode
	AllVAR_IDENTIFIER() []antlr.TerminalNode
	VAR_IDENTIFIER(i int) antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	Block() IBlockContext
	RETURN_TYPE() antlr.TerminalNode
	Data_type() IData_typeContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunc_declarationContext differentiates from other interfaces.
	IsFunc_declarationContext()
}

type Func_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_declarationContext() *Func_declarationContext {
	var p = new(Func_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_func_declaration
	return p
}

func InitEmptyFunc_declarationContext(p *Func_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_func_declaration
}

func (*Func_declarationContext) IsFunc_declarationContext() {}

func NewFunc_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_declarationContext {
	var p = new(Func_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_func_declaration

	return p
}

func (s *Func_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_declarationContext) AT() antlr.TerminalNode {
	return s.GetToken(FluxAT, 0)
}

func (s *Func_declarationContext) FUNC() antlr.TerminalNode {
	return s.GetToken(FluxFUNC, 0)
}

func (s *Func_declarationContext) AllVAR_IDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(FluxVAR_IDENTIFIER)
}

func (s *Func_declarationContext) VAR_IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(FluxVAR_IDENTIFIER, i)
}

func (s *Func_declarationContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxL_PAREN, 0)
}

func (s *Func_declarationContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxR_PAREN, 0)
}

func (s *Func_declarationContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Func_declarationContext) RETURN_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxRETURN_TYPE, 0)
}

func (s *Func_declarationContext) Data_type() IData_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IData_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IData_typeContext)
}

func (s *Func_declarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FluxCOMMA)
}

func (s *Func_declarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FluxCOMMA, i)
}

func (s *Func_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterFunc_declaration(s)
	}
}

func (s *Func_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitFunc_declaration(s)
	}
}

func (s *Func_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitFunc_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Func_declaration() (localctx IFunc_declarationContext) {
	localctx = NewFunc_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FluxRULE_func_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(FluxAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Match(FluxFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		p.Match(FluxVAR_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Match(FluxL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FluxVAR_IDENTIFIER {
		{
			p.SetState(122)
			p.Match(FluxVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxCOMMA {
			{
				p.SetState(123)
				p.Match(FluxCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(124)
				p.Match(FluxVAR_IDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(132)
		p.Match(FluxR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FluxRETURN_TYPE {
		{
			p.SetState(133)
			p.Match(FluxRETURN_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.Data_type()
		}

	}
	{
		p.SetState(137)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_typeContext is an interface to support dynamic dispatch.
type IVar_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT_TYPE() antlr.TerminalNode
	NUMBER_TYPE() antlr.TerminalNode
	BOOLEAN_TYPE() antlr.TerminalNode
	IPV4_TYPE() antlr.TerminalNode

	// IsVar_typeContext differentiates from other interfaces.
	IsVar_typeContext()
}

type Var_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_typeContext() *Var_typeContext {
	var p = new(Var_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_var_type
	return p
}

func InitEmptyVar_typeContext(p *Var_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_var_type
}

func (*Var_typeContext) IsVar_typeContext() {}

func NewVar_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_typeContext {
	var p = new(Var_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_var_type

	return p
}

func (s *Var_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_typeContext) TEXT_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxTEXT_TYPE, 0)
}

func (s *Var_typeContext) NUMBER_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxNUMBER_TYPE, 0)
}

func (s *Var_typeContext) BOOLEAN_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxBOOLEAN_TYPE, 0)
}

func (s *Var_typeContext) IPV4_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxIPV4_TYPE, 0)
}

func (s *Var_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterVar_type(s)
	}
}

func (s *Var_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitVar_type(s)
	}
}

func (s *Var_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitVar_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Var_type() (localctx IVar_typeContext) {
	localctx = NewVar_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FluxRULE_var_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2132) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_nameContext is an interface to support dynamic dispatch.
type IVar_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_IDENTIFIER() antlr.TerminalNode

	// IsVar_nameContext differentiates from other interfaces.
	IsVar_nameContext()
}

type Var_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_nameContext() *Var_nameContext {
	var p = new(Var_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_var_name
	return p
}

func InitEmptyVar_nameContext(p *Var_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_var_name
}

func (*Var_nameContext) IsVar_nameContext() {}

func NewVar_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_nameContext {
	var p = new(Var_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_var_name

	return p
}

func (s *Var_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_nameContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FluxVAR_IDENTIFIER, 0)
}

func (s *Var_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterVar_name(s)
	}
}

func (s *Var_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitVar_name(s)
	}
}

func (s *Var_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitVar_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Var_name() (localctx IVar_nameContext) {
	localctx = NewVar_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FluxRULE_var_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(FluxVAR_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_valueContext is an interface to support dynamic dispatch.
type IVar_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IPV4() antlr.TerminalNode
	TEXT() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode

	// IsVar_valueContext differentiates from other interfaces.
	IsVar_valueContext()
}

type Var_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_valueContext() *Var_valueContext {
	var p = new(Var_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_var_value
	return p
}

func InitEmptyVar_valueContext(p *Var_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_var_value
}

func (*Var_valueContext) IsVar_valueContext() {}

func NewVar_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_valueContext {
	var p = new(Var_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_var_value

	return p
}

func (s *Var_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_valueContext) IPV4() antlr.TerminalNode {
	return s.GetToken(FluxIPV4, 0)
}

func (s *Var_valueContext) TEXT() antlr.TerminalNode {
	return s.GetToken(FluxTEXT, 0)
}

func (s *Var_valueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FluxNUMBER, 0)
}

func (s *Var_valueContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(FluxBOOLEAN, 0)
}

func (s *Var_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterVar_value(s)
	}
}

func (s *Var_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitVar_value(s)
	}
}

func (s *Var_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitVar_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Var_value() (localctx IVar_valueContext) {
	localctx = NewVar_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FluxRULE_var_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1066) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_one_expressionContext is an interface to support dynamic dispatch.
type IOp_one_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_INCREMENT() antlr.TerminalNode
	OP_DECREMENT() antlr.TerminalNode

	// IsOp_one_expressionContext differentiates from other interfaces.
	IsOp_one_expressionContext()
}

type Op_one_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_one_expressionContext() *Op_one_expressionContext {
	var p = new(Op_one_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_one_expression
	return p
}

func InitEmptyOp_one_expressionContext(p *Op_one_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_one_expression
}

func (*Op_one_expressionContext) IsOp_one_expressionContext() {}

func NewOp_one_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_one_expressionContext {
	var p = new(Op_one_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_op_one_expression

	return p
}

func (s *Op_one_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Op_one_expressionContext) OP_INCREMENT() antlr.TerminalNode {
	return s.GetToken(FluxOP_INCREMENT, 0)
}

func (s *Op_one_expressionContext) OP_DECREMENT() antlr.TerminalNode {
	return s.GetToken(FluxOP_DECREMENT, 0)
}

func (s *Op_one_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_one_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_one_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterOp_one_expression(s)
	}
}

func (s *Op_one_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitOp_one_expression(s)
	}
}

func (s *Op_one_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitOp_one_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Op_one_expression() (localctx IOp_one_expressionContext) {
	localctx = NewOp_one_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FluxRULE_op_one_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FluxOP_INCREMENT || _la == FluxOP_DECREMENT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_one_declarationContext is an interface to support dynamic dispatch.
type IOp_one_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_name() IVar_nameContext
	OP_INCREMENT() antlr.TerminalNode
	OP_DECREMENT() antlr.TerminalNode
	NUMBER_TYPE() antlr.TerminalNode

	// IsOp_one_declarationContext differentiates from other interfaces.
	IsOp_one_declarationContext()
}

type Op_one_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_one_declarationContext() *Op_one_declarationContext {
	var p = new(Op_one_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_one_declaration
	return p
}

func InitEmptyOp_one_declarationContext(p *Op_one_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_one_declaration
}

func (*Op_one_declarationContext) IsOp_one_declarationContext() {}

func NewOp_one_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_one_declarationContext {
	var p = new(Op_one_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_op_one_declaration

	return p
}

func (s *Op_one_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Op_one_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Op_one_declarationContext) OP_INCREMENT() antlr.TerminalNode {
	return s.GetToken(FluxOP_INCREMENT, 0)
}

func (s *Op_one_declarationContext) OP_DECREMENT() antlr.TerminalNode {
	return s.GetToken(FluxOP_DECREMENT, 0)
}

func (s *Op_one_declarationContext) NUMBER_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxNUMBER_TYPE, 0)
}

func (s *Op_one_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_one_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_one_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterOp_one_declaration(s)
	}
}

func (s *Op_one_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitOp_one_declaration(s)
	}
}

func (s *Op_one_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitOp_one_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Op_one_declaration() (localctx IOp_one_declarationContext) {
	localctx = NewOp_one_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FluxRULE_op_one_declaration)
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(147)
			p.Var_name()
		}
		{
			p.SetState(148)
			p.Match(FluxOP_INCREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.Var_name()
		}
		{
			p.SetState(151)
			p.Match(FluxOP_DECREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(153)
			p.Match(FluxOP_INCREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Var_name()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(155)
			p.Match(FluxOP_DECREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.Var_name()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(157)
			p.Match(FluxNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Var_name()
		}
		{
			p.SetState(159)
			p.Match(FluxOP_INCREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(161)
			p.Match(FluxNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Var_name()
		}
		{
			p.SetState(163)
			p.Match(FluxOP_DECREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IString_var_declarationContext is an interface to support dynamic dispatch.
type IString_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT_TYPE() antlr.TerminalNode
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	TEXT() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsString_var_declarationContext differentiates from other interfaces.
	IsString_var_declarationContext()
}

type String_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_var_declarationContext() *String_var_declarationContext {
	var p = new(String_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_string_var_declaration
	return p
}

func InitEmptyString_var_declarationContext(p *String_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_string_var_declaration
}

func (*String_var_declarationContext) IsString_var_declarationContext() {}

func NewString_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_var_declarationContext {
	var p = new(String_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_string_var_declaration

	return p
}

func (s *String_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *String_var_declarationContext) TEXT_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxTEXT_TYPE, 0)
}

func (s *String_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *String_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxL_BLOCK, 0)
}

func (s *String_var_declarationContext) TEXT() antlr.TerminalNode {
	return s.GetToken(FluxTEXT, 0)
}

func (s *String_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxR_BLOCK, 0)
}

func (s *String_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FluxNEWLINE)
}

func (s *String_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FluxNEWLINE, i)
}

func (s *String_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterString_var_declaration(s)
	}
}

func (s *String_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitString_var_declaration(s)
	}
}

func (s *String_var_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitString_var_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) String_var_declaration() (localctx IString_var_declarationContext) {
	localctx = NewString_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FluxRULE_string_var_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(FluxTEXT_TYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Var_name()
	}
	{
		p.SetState(169)
		p.Match(FluxL_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FluxNEWLINE {
		{
			p.SetState(170)
			p.Match(FluxNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(176)
		p.Match(FluxTEXT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FluxNEWLINE {
		{
			p.SetState(177)
			p.Match(FluxNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(183)
		p.Match(FluxR_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumber_var_declarationContext is an interface to support dynamic dispatch.
type INumber_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER_TYPE() antlr.TerminalNode
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	Math_expression() IMath_expressionContext

	// IsNumber_var_declarationContext differentiates from other interfaces.
	IsNumber_var_declarationContext()
}

type Number_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumber_var_declarationContext() *Number_var_declarationContext {
	var p = new(Number_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_number_var_declaration
	return p
}

func InitEmptyNumber_var_declarationContext(p *Number_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_number_var_declaration
}

func (*Number_var_declarationContext) IsNumber_var_declarationContext() {}

func NewNumber_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Number_var_declarationContext {
	var p = new(Number_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_number_var_declaration

	return p
}

func (s *Number_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Number_var_declarationContext) NUMBER_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxNUMBER_TYPE, 0)
}

func (s *Number_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Number_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxL_BLOCK, 0)
}

func (s *Number_var_declarationContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FluxNUMBER, 0)
}

func (s *Number_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxR_BLOCK, 0)
}

func (s *Number_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FluxNEWLINE)
}

func (s *Number_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FluxNEWLINE, i)
}

func (s *Number_var_declarationContext) Math_expression() IMath_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMath_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMath_expressionContext)
}

func (s *Number_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Number_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Number_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterNumber_var_declaration(s)
	}
}

func (s *Number_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitNumber_var_declaration(s)
	}
}

func (s *Number_var_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitNumber_var_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Number_var_declaration() (localctx INumber_var_declarationContext) {
	localctx = NewNumber_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FluxRULE_number_var_declaration)
	var _la int

	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.Match(FluxNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Var_name()
		}
		{
			p.SetState(187)
			p.Match(FluxL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(188)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(193)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(194)
			p.Match(FluxNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(195)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(200)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(201)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(203)
			p.Match(FluxNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.Var_name()
		}
		{
			p.SetState(205)
			p.Match(FluxL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(206)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(211)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(212)
			p.Math_expression()
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(213)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(218)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(219)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoolean_var_declarationContext is an interface to support dynamic dispatch.
type IBoolean_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOLEAN_TYPE() antlr.TerminalNode
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	Logical_expression() ILogical_expressionContext
	Comparative_expression() IComparative_expressionContext

	// IsBoolean_var_declarationContext differentiates from other interfaces.
	IsBoolean_var_declarationContext()
}

type Boolean_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolean_var_declarationContext() *Boolean_var_declarationContext {
	var p = new(Boolean_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_boolean_var_declaration
	return p
}

func InitEmptyBoolean_var_declarationContext(p *Boolean_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_boolean_var_declaration
}

func (*Boolean_var_declarationContext) IsBoolean_var_declarationContext() {}

func NewBoolean_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_var_declarationContext {
	var p = new(Boolean_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_boolean_var_declaration

	return p
}

func (s *Boolean_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_var_declarationContext) BOOLEAN_TYPE() antlr.TerminalNode {
	return s.GetToken(FluxBOOLEAN_TYPE, 0)
}

func (s *Boolean_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Boolean_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxL_BLOCK, 0)
}

func (s *Boolean_var_declarationContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(FluxBOOLEAN, 0)
}

func (s *Boolean_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxR_BLOCK, 0)
}

func (s *Boolean_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FluxNEWLINE)
}

func (s *Boolean_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FluxNEWLINE, i)
}

func (s *Boolean_var_declarationContext) Logical_expression() ILogical_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogical_expressionContext)
}

func (s *Boolean_var_declarationContext) Comparative_expression() IComparative_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparative_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparative_expressionContext)
}

func (s *Boolean_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterBoolean_var_declaration(s)
	}
}

func (s *Boolean_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitBoolean_var_declaration(s)
	}
}

func (s *Boolean_var_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitBoolean_var_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Boolean_var_declaration() (localctx IBoolean_var_declarationContext) {
	localctx = NewBoolean_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FluxRULE_boolean_var_declaration)
	var _la int

	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(223)
			p.Match(FluxBOOLEAN_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.Var_name()
		}
		{
			p.SetState(225)
			p.Match(FluxL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(229)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(226)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(231)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(232)
			p.Match(FluxBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(233)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(238)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(239)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(241)
			p.Match(FluxBOOLEAN_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Var_name()
		}
		{
			p.SetState(243)
			p.Match(FluxL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(244)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(249)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(250)
			p.logical_expression(0)
		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(251)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(256)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(257)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(259)
			p.Match(FluxBOOLEAN_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)
			p.Var_name()
		}
		{
			p.SetState(261)
			p.Match(FluxL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(262)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(267)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(268)
			p.comparative_expression(0)
		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(269)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(274)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(275)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingle_var_declarationContext is an interface to support dynamic dispatch.
type ISingle_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	String_var_declaration() IString_var_declarationContext
	Number_var_declaration() INumber_var_declarationContext
	Boolean_var_declaration() IBoolean_var_declarationContext
	Var_type() IVar_typeContext
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsSingle_var_declarationContext differentiates from other interfaces.
	IsSingle_var_declarationContext()
}

type Single_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingle_var_declarationContext() *Single_var_declarationContext {
	var p = new(Single_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_single_var_declaration
	return p
}

func InitEmptySingle_var_declarationContext(p *Single_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_single_var_declaration
}

func (*Single_var_declarationContext) IsSingle_var_declarationContext() {}

func NewSingle_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Single_var_declarationContext {
	var p = new(Single_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_single_var_declaration

	return p
}

func (s *Single_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Single_var_declarationContext) String_var_declaration() IString_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_var_declarationContext)
}

func (s *Single_var_declarationContext) Number_var_declaration() INumber_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumber_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumber_var_declarationContext)
}

func (s *Single_var_declarationContext) Boolean_var_declaration() IBoolean_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_var_declarationContext)
}

func (s *Single_var_declarationContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *Single_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Single_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxL_BLOCK, 0)
}

func (s *Single_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxR_BLOCK, 0)
}

func (s *Single_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FluxNEWLINE)
}

func (s *Single_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FluxNEWLINE, i)
}

func (s *Single_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Single_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Single_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterSingle_var_declaration(s)
	}
}

func (s *Single_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitSingle_var_declaration(s)
	}
}

func (s *Single_var_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitSingle_var_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Single_var_declaration() (localctx ISingle_var_declarationContext) {
	localctx = NewSingle_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FluxRULE_single_var_declaration)
	var _la int

	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(279)
			p.String_var_declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(280)
			p.Number_var_declaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(281)
			p.Boolean_var_declaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(282)
			p.Var_type()
		}
		{
			p.SetState(283)
			p.Var_name()
		}
		{
			p.SetState(284)
			p.Match(FluxL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(285)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(290)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(291)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArray_var_declarationContext is an interface to support dynamic dispatch.
type IArray_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_type() IVar_typeContext
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	AllVar_value() []IVar_valueContext
	Var_value(i int) IVar_valueContext
	R_BLOCK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllMath_expression() []IMath_expressionContext
	Math_expression(i int) IMath_expressionContext

	// IsArray_var_declarationContext differentiates from other interfaces.
	IsArray_var_declarationContext()
}

type Array_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_var_declarationContext() *Array_var_declarationContext {
	var p = new(Array_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_array_var_declaration
	return p
}

func InitEmptyArray_var_declarationContext(p *Array_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_array_var_declaration
}

func (*Array_var_declarationContext) IsArray_var_declarationContext() {}

func NewArray_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_var_declarationContext {
	var p = new(Array_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_array_var_declaration

	return p
}

func (s *Array_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_var_declarationContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *Array_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Array_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxL_BLOCK, 0)
}

func (s *Array_var_declarationContext) AllVar_value() []IVar_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_valueContext); ok {
			len++
		}
	}

	tst := make([]IVar_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_valueContext); ok {
			tst[i] = t.(IVar_valueContext)
			i++
		}
	}

	return tst
}

func (s *Array_var_declarationContext) Var_value(i int) IVar_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_valueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_valueContext)
}

func (s *Array_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxR_BLOCK, 0)
}

func (s *Array_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FluxNEWLINE)
}

func (s *Array_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FluxNEWLINE, i)
}

func (s *Array_var_declarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FluxCOMMA)
}

func (s *Array_var_declarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FluxCOMMA, i)
}

func (s *Array_var_declarationContext) AllMath_expression() []IMath_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMath_expressionContext); ok {
			len++
		}
	}

	tst := make([]IMath_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMath_expressionContext); ok {
			tst[i] = t.(IMath_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Array_var_declarationContext) Math_expression(i int) IMath_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMath_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMath_expressionContext)
}

func (s *Array_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterArray_var_declaration(s)
	}
}

func (s *Array_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitArray_var_declaration(s)
	}
}

func (s *Array_var_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitArray_var_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Array_var_declaration() (localctx IArray_var_declarationContext) {
	localctx = NewArray_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, FluxRULE_array_var_declaration)
	var _la int

	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(295)
			p.Var_type()
		}
		{
			p.SetState(296)
			p.Var_name()
		}
		{
			p.SetState(297)
			p.Match(FluxL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(298)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(303)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(304)
			p.Var_value()
		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxCOMMA {
			{
				p.SetState(305)
				p.Match(FluxCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(309)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == FluxNEWLINE {
				{
					p.SetState(306)
					p.Match(FluxNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(311)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(312)
				p.Var_value()
			}

			p.SetState(317)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(318)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(323)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(324)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(326)
			p.Var_type()
		}
		{
			p.SetState(327)
			p.Var_name()
		}
		{
			p.SetState(328)
			p.Match(FluxL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(332)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(329)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(334)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(335)
			p.Math_expression()
		}
		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxCOMMA {
			{
				p.SetState(336)
				p.Match(FluxCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(340)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == FluxNEWLINE {
				{
					p.SetState(337)
					p.Match(FluxNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(342)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(343)
				p.Math_expression()
			}

			p.SetState(348)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(349)
			p.Match(FluxNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(350)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_assignmentContext is an interface to support dynamic dispatch.
type IVar_assignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	Var_value() IVar_valueContext
	R_BLOCK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	Math_expression() IMath_expressionContext

	// IsVar_assignmentContext differentiates from other interfaces.
	IsVar_assignmentContext()
}

type Var_assignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_assignmentContext() *Var_assignmentContext {
	var p = new(Var_assignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_var_assignment
	return p
}

func InitEmptyVar_assignmentContext(p *Var_assignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_var_assignment
}

func (*Var_assignmentContext) IsVar_assignmentContext() {}

func NewVar_assignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_assignmentContext {
	var p = new(Var_assignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_var_assignment

	return p
}

func (s *Var_assignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_assignmentContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Var_assignmentContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxL_BLOCK, 0)
}

func (s *Var_assignmentContext) Var_value() IVar_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_valueContext)
}

func (s *Var_assignmentContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(FluxR_BLOCK, 0)
}

func (s *Var_assignmentContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FluxNEWLINE)
}

func (s *Var_assignmentContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FluxNEWLINE, i)
}

func (s *Var_assignmentContext) Math_expression() IMath_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMath_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMath_expressionContext)
}

func (s *Var_assignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_assignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_assignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterVar_assignment(s)
	}
}

func (s *Var_assignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitVar_assignment(s)
	}
}

func (s *Var_assignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitVar_assignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Var_assignment() (localctx IVar_assignmentContext) {
	localctx = NewVar_assignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, FluxRULE_var_assignment)
	var _la int

	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(354)
			p.Var_name()
		}
		{
			p.SetState(355)
			p.Match(FluxL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(356)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(361)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(362)
			p.Var_value()
		}
		p.SetState(366)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(363)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(368)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(369)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(371)
			p.Var_name()
		}
		{
			p.SetState(372)
			p.Match(FluxL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(376)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(373)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(378)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(379)
			p.Math_expression()
		}
		p.SetState(383)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxNEWLINE {
			{
				p.SetState(380)
				p.Match(FluxNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(385)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(386)
			p.Match(FluxR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_level1Context is an interface to support dynamic dispatch.
type IOp_level1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_MULTIPLY() antlr.TerminalNode
	OP_DIVIDE() antlr.TerminalNode
	OP_MOD() antlr.TerminalNode
	OP_POWER() antlr.TerminalNode

	// IsOp_level1Context differentiates from other interfaces.
	IsOp_level1Context()
}

type Op_level1Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level1Context() *Op_level1Context {
	var p = new(Op_level1Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level1
	return p
}

func InitEmptyOp_level1Context(p *Op_level1Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level1
}

func (*Op_level1Context) IsOp_level1Context() {}

func NewOp_level1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level1Context {
	var p = new(Op_level1Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_op_level1

	return p
}

func (s *Op_level1Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level1Context) OP_MULTIPLY() antlr.TerminalNode {
	return s.GetToken(FluxOP_MULTIPLY, 0)
}

func (s *Op_level1Context) OP_DIVIDE() antlr.TerminalNode {
	return s.GetToken(FluxOP_DIVIDE, 0)
}

func (s *Op_level1Context) OP_MOD() antlr.TerminalNode {
	return s.GetToken(FluxOP_MOD, 0)
}

func (s *Op_level1Context) OP_POWER() antlr.TerminalNode {
	return s.GetToken(FluxOP_POWER, 0)
}

func (s *Op_level1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterOp_level1(s)
	}
}

func (s *Op_level1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitOp_level1(s)
	}
}

func (s *Op_level1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitOp_level1(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Op_level1() (localctx IOp_level1Context) {
	localctx = NewOp_level1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, FluxRULE_op_level1)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&128849018880) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_level2Context is an interface to support dynamic dispatch.
type IOp_level2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_PLUS() antlr.TerminalNode
	OP_MINUS() antlr.TerminalNode

	// IsOp_level2Context differentiates from other interfaces.
	IsOp_level2Context()
}

type Op_level2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level2Context() *Op_level2Context {
	var p = new(Op_level2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level2
	return p
}

func InitEmptyOp_level2Context(p *Op_level2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level2
}

func (*Op_level2Context) IsOp_level2Context() {}

func NewOp_level2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level2Context {
	var p = new(Op_level2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_op_level2

	return p
}

func (s *Op_level2Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level2Context) OP_PLUS() antlr.TerminalNode {
	return s.GetToken(FluxOP_PLUS, 0)
}

func (s *Op_level2Context) OP_MINUS() antlr.TerminalNode {
	return s.GetToken(FluxOP_MINUS, 0)
}

func (s *Op_level2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterOp_level2(s)
	}
}

func (s *Op_level2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitOp_level2(s)
	}
}

func (s *Op_level2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitOp_level2(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Op_level2() (localctx IOp_level2Context) {
	localctx = NewOp_level2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, FluxRULE_op_level2)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(392)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FluxOP_PLUS || _la == FluxOP_MINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_level3Context is an interface to support dynamic dispatch.
type IOp_level3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_AND() antlr.TerminalNode
	OP_OR() antlr.TerminalNode
	OP_XOR() antlr.TerminalNode

	// IsOp_level3Context differentiates from other interfaces.
	IsOp_level3Context()
}

type Op_level3Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level3Context() *Op_level3Context {
	var p = new(Op_level3Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level3
	return p
}

func InitEmptyOp_level3Context(p *Op_level3Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level3
}

func (*Op_level3Context) IsOp_level3Context() {}

func NewOp_level3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level3Context {
	var p = new(Op_level3Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_op_level3

	return p
}

func (s *Op_level3Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level3Context) OP_AND() antlr.TerminalNode {
	return s.GetToken(FluxOP_AND, 0)
}

func (s *Op_level3Context) OP_OR() antlr.TerminalNode {
	return s.GetToken(FluxOP_OR, 0)
}

func (s *Op_level3Context) OP_XOR() antlr.TerminalNode {
	return s.GetToken(FluxOP_XOR, 0)
}

func (s *Op_level3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterOp_level3(s)
	}
}

func (s *Op_level3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitOp_level3(s)
	}
}

func (s *Op_level3Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitOp_level3(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Op_level3() (localctx IOp_level3Context) {
	localctx = NewOp_level3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, FluxRULE_op_level3)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(394)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&61572651155456) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_level4Context is an interface to support dynamic dispatch.
type IOp_level4Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_EQUAL() antlr.TerminalNode
	OP_NOT_EQUAL() antlr.TerminalNode
	OP_LESS() antlr.TerminalNode
	OP_LESS_EQUAL() antlr.TerminalNode
	OP_GREATER() antlr.TerminalNode
	OP_GREATER_EQUAL() antlr.TerminalNode

	// IsOp_level4Context differentiates from other interfaces.
	IsOp_level4Context()
}

type Op_level4Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level4Context() *Op_level4Context {
	var p = new(Op_level4Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level4
	return p
}

func InitEmptyOp_level4Context(p *Op_level4Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level4
}

func (*Op_level4Context) IsOp_level4Context() {}

func NewOp_level4Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level4Context {
	var p = new(Op_level4Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_op_level4

	return p
}

func (s *Op_level4Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level4Context) OP_EQUAL() antlr.TerminalNode {
	return s.GetToken(FluxOP_EQUAL, 0)
}

func (s *Op_level4Context) OP_NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(FluxOP_NOT_EQUAL, 0)
}

func (s *Op_level4Context) OP_LESS() antlr.TerminalNode {
	return s.GetToken(FluxOP_LESS, 0)
}

func (s *Op_level4Context) OP_LESS_EQUAL() antlr.TerminalNode {
	return s.GetToken(FluxOP_LESS_EQUAL, 0)
}

func (s *Op_level4Context) OP_GREATER() antlr.TerminalNode {
	return s.GetToken(FluxOP_GREATER, 0)
}

func (s *Op_level4Context) OP_GREATER_EQUAL() antlr.TerminalNode {
	return s.GetToken(FluxOP_GREATER_EQUAL, 0)
}

func (s *Op_level4Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level4Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level4Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterOp_level4(s)
	}
}

func (s *Op_level4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitOp_level4(s)
	}
}

func (s *Op_level4Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitOp_level4(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Op_level4() (localctx IOp_level4Context) {
	localctx = NewOp_level4Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, FluxRULE_op_level4)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(396)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8658654068736) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_level5Context is an interface to support dynamic dispatch.
type IOp_level5Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_NOT() antlr.TerminalNode

	// IsOp_level5Context differentiates from other interfaces.
	IsOp_level5Context()
}

type Op_level5Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level5Context() *Op_level5Context {
	var p = new(Op_level5Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level5
	return p
}

func InitEmptyOp_level5Context(p *Op_level5Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_op_level5
}

func (*Op_level5Context) IsOp_level5Context() {}

func NewOp_level5Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level5Context {
	var p = new(Op_level5Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_op_level5

	return p
}

func (s *Op_level5Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level5Context) OP_NOT() antlr.TerminalNode {
	return s.GetToken(FluxOP_NOT, 0)
}

func (s *Op_level5Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level5Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level5Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterOp_level5(s)
	}
}

func (s *Op_level5Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitOp_level5(s)
	}
}

func (s *Op_level5Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitOp_level5(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Op_level5() (localctx IOp_level5Context) {
	localctx = NewOp_level5Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, FluxRULE_op_level5)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(398)
		p.Match(FluxOP_NOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberic_expressionContext is an interface to support dynamic dispatch.
type INumberic_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	AllNumberic_expression() []INumberic_expressionContext
	Numberic_expression(i int) INumberic_expressionContext
	R_PAREN() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	Get_variable() IGet_variableContext
	Function_call() IFunction_callContext
	OP_MINUS() antlr.TerminalNode
	Op_level1() IOp_level1Context
	Op_level2() IOp_level2Context

	// IsNumberic_expressionContext differentiates from other interfaces.
	IsNumberic_expressionContext()
}

type Numberic_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberic_expressionContext() *Numberic_expressionContext {
	var p = new(Numberic_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_numberic_expression
	return p
}

func InitEmptyNumberic_expressionContext(p *Numberic_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_numberic_expression
}

func (*Numberic_expressionContext) IsNumberic_expressionContext() {}

func NewNumberic_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Numberic_expressionContext {
	var p = new(Numberic_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_numberic_expression

	return p
}

func (s *Numberic_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Numberic_expressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxL_PAREN, 0)
}

func (s *Numberic_expressionContext) AllNumberic_expression() []INumberic_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumberic_expressionContext); ok {
			len++
		}
	}

	tst := make([]INumberic_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumberic_expressionContext); ok {
			tst[i] = t.(INumberic_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Numberic_expressionContext) Numberic_expression(i int) INumberic_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberic_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberic_expressionContext)
}

func (s *Numberic_expressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxR_PAREN, 0)
}

func (s *Numberic_expressionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FluxNUMBER, 0)
}

func (s *Numberic_expressionContext) Get_variable() IGet_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_variableContext)
}

func (s *Numberic_expressionContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Numberic_expressionContext) OP_MINUS() antlr.TerminalNode {
	return s.GetToken(FluxOP_MINUS, 0)
}

func (s *Numberic_expressionContext) Op_level1() IOp_level1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level1Context)
}

func (s *Numberic_expressionContext) Op_level2() IOp_level2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level2Context)
}

func (s *Numberic_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Numberic_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Numberic_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterNumberic_expression(s)
	}
}

func (s *Numberic_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitNumberic_expression(s)
	}
}

func (s *Numberic_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitNumberic_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Numberic_expression() (localctx INumberic_expressionContext) {
	return p.numberic_expression(0)
}

func (p *Flux) numberic_expression(_p int) (localctx INumberic_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewNumberic_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx INumberic_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 50
	p.EnterRecursionRule(localctx, 50, FluxRULE_numberic_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(401)
			p.Match(FluxL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(402)
			p.numberic_expression(0)
		}
		{
			p.SetState(403)
			p.Match(FluxR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(405)
			p.Match(FluxNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(406)
			p.Get_variable()
		}

	case 4:
		{
			p.SetState(407)
			p.Function_call()
		}

	case 5:
		{
			p.SetState(408)
			p.Match(FluxOP_MINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)
			p.numberic_expression(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(420)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
			case 1:
				localctx = NewNumberic_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FluxRULE_numberic_expression)
				p.SetState(412)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(413)
					p.Op_level1()
				}
				{
					p.SetState(414)
					p.numberic_expression(7)
				}

			case 2:
				localctx = NewNumberic_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FluxRULE_numberic_expression)
				p.SetState(416)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(417)
					p.Op_level2()
				}
				{
					p.SetState(418)
					p.numberic_expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(424)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogical_expressionContext is an interface to support dynamic dispatch.
type ILogical_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	AllLogical_expression() []ILogical_expressionContext
	Logical_expression(i int) ILogical_expressionContext
	R_PAREN() antlr.TerminalNode
	OP_NOT() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	Get_variable() IGet_variableContext
	Function_call() IFunction_callContext
	Op_level3() IOp_level3Context
	Op_level4() IOp_level4Context

	// IsLogical_expressionContext differentiates from other interfaces.
	IsLogical_expressionContext()
}

type Logical_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogical_expressionContext() *Logical_expressionContext {
	var p = new(Logical_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_logical_expression
	return p
}

func InitEmptyLogical_expressionContext(p *Logical_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_logical_expression
}

func (*Logical_expressionContext) IsLogical_expressionContext() {}

func NewLogical_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Logical_expressionContext {
	var p = new(Logical_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_logical_expression

	return p
}

func (s *Logical_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Logical_expressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxL_PAREN, 0)
}

func (s *Logical_expressionContext) AllLogical_expression() []ILogical_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogical_expressionContext); ok {
			len++
		}
	}

	tst := make([]ILogical_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogical_expressionContext); ok {
			tst[i] = t.(ILogical_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Logical_expressionContext) Logical_expression(i int) ILogical_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogical_expressionContext)
}

func (s *Logical_expressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxR_PAREN, 0)
}

func (s *Logical_expressionContext) OP_NOT() antlr.TerminalNode {
	return s.GetToken(FluxOP_NOT, 0)
}

func (s *Logical_expressionContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(FluxBOOLEAN, 0)
}

func (s *Logical_expressionContext) Get_variable() IGet_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_variableContext)
}

func (s *Logical_expressionContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Logical_expressionContext) Op_level3() IOp_level3Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level3Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level3Context)
}

func (s *Logical_expressionContext) Op_level4() IOp_level4Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level4Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level4Context)
}

func (s *Logical_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logical_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Logical_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterLogical_expression(s)
	}
}

func (s *Logical_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitLogical_expression(s)
	}
}

func (s *Logical_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitLogical_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Logical_expression() (localctx ILogical_expressionContext) {
	return p.logical_expression(0)
}

func (p *Flux) logical_expression(_p int) (localctx ILogical_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewLogical_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILogical_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 52
	p.EnterRecursionRule(localctx, 52, FluxRULE_logical_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(426)
			p.Match(FluxL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(427)
			p.logical_expression(0)
		}
		{
			p.SetState(428)
			p.Match(FluxR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(430)
			p.Match(FluxOP_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(431)
			p.logical_expression(4)
		}

	case 3:
		{
			p.SetState(432)
			p.Match(FluxBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(433)
			p.Get_variable()
		}

	case 5:
		{
			p.SetState(434)
			p.Function_call()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(445)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
			case 1:
				localctx = NewLogical_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FluxRULE_logical_expression)
				p.SetState(437)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(438)
					p.Op_level3()
				}
				{
					p.SetState(439)
					p.logical_expression(7)
				}

			case 2:
				localctx = NewLogical_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FluxRULE_logical_expression)
				p.SetState(441)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(442)
					p.Op_level4()
				}
				{
					p.SetState(443)
					p.logical_expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(449)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparative_expressionContext is an interface to support dynamic dispatch.
type IComparative_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNumberic_expression() []INumberic_expressionContext
	Numberic_expression(i int) INumberic_expressionContext
	Op_level4() IOp_level4Context
	AllLogical_expression() []ILogical_expressionContext
	Logical_expression(i int) ILogical_expressionContext
	Op_level3() IOp_level3Context
	Op_level5() IOp_level5Context
	L_PAREN() antlr.TerminalNode
	Comparative_expression() IComparative_expressionContext
	R_PAREN() antlr.TerminalNode

	// IsComparative_expressionContext differentiates from other interfaces.
	IsComparative_expressionContext()
}

type Comparative_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparative_expressionContext() *Comparative_expressionContext {
	var p = new(Comparative_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_comparative_expression
	return p
}

func InitEmptyComparative_expressionContext(p *Comparative_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_comparative_expression
}

func (*Comparative_expressionContext) IsComparative_expressionContext() {}

func NewComparative_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comparative_expressionContext {
	var p = new(Comparative_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_comparative_expression

	return p
}

func (s *Comparative_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Comparative_expressionContext) AllNumberic_expression() []INumberic_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumberic_expressionContext); ok {
			len++
		}
	}

	tst := make([]INumberic_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumberic_expressionContext); ok {
			tst[i] = t.(INumberic_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Comparative_expressionContext) Numberic_expression(i int) INumberic_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberic_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberic_expressionContext)
}

func (s *Comparative_expressionContext) Op_level4() IOp_level4Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level4Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level4Context)
}

func (s *Comparative_expressionContext) AllLogical_expression() []ILogical_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogical_expressionContext); ok {
			len++
		}
	}

	tst := make([]ILogical_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogical_expressionContext); ok {
			tst[i] = t.(ILogical_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Comparative_expressionContext) Logical_expression(i int) ILogical_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogical_expressionContext)
}

func (s *Comparative_expressionContext) Op_level3() IOp_level3Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level3Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level3Context)
}

func (s *Comparative_expressionContext) Op_level5() IOp_level5Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level5Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level5Context)
}

func (s *Comparative_expressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxL_PAREN, 0)
}

func (s *Comparative_expressionContext) Comparative_expression() IComparative_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparative_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparative_expressionContext)
}

func (s *Comparative_expressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxR_PAREN, 0)
}

func (s *Comparative_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comparative_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comparative_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterComparative_expression(s)
	}
}

func (s *Comparative_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitComparative_expression(s)
	}
}

func (s *Comparative_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitComparative_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Comparative_expression() (localctx IComparative_expressionContext) {
	return p.comparative_expression(0)
}

func (p *Flux) comparative_expression(_p int) (localctx IComparative_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewComparative_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IComparative_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 54
	p.EnterRecursionRule(localctx, 54, FluxRULE_comparative_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(451)
			p.numberic_expression(0)
		}
		{
			p.SetState(452)
			p.Op_level4()
		}
		{
			p.SetState(453)
			p.numberic_expression(0)
		}

	case 2:
		{
			p.SetState(455)
			p.logical_expression(0)
		}
		{
			p.SetState(456)
			p.Op_level3()
		}
		{
			p.SetState(457)
			p.logical_expression(0)
		}

	case 3:
		{
			p.SetState(459)
			p.Op_level5()
		}
		{
			p.SetState(460)
			p.Match(FluxL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(461)
			p.comparative_expression(0)
		}
		{
			p.SetState(462)
			p.Match(FluxR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewComparative_expressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, FluxRULE_comparative_expression)
			p.SetState(466)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(467)
				p.Op_level3()
			}
			{
				p.SetState(468)
				p.logical_expression(0)
			}

		}
		p.SetState(474)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGet_variableContext is an interface to support dynamic dispatch.
type IGet_variableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_IDENTIFIER() antlr.TerminalNode
	Get_array_element() IGet_array_elementContext
	Get_child() IGet_childContext

	// IsGet_variableContext differentiates from other interfaces.
	IsGet_variableContext()
}

type Get_variableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGet_variableContext() *Get_variableContext {
	var p = new(Get_variableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_get_variable
	return p
}

func InitEmptyGet_variableContext(p *Get_variableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_get_variable
}

func (*Get_variableContext) IsGet_variableContext() {}

func NewGet_variableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Get_variableContext {
	var p = new(Get_variableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_get_variable

	return p
}

func (s *Get_variableContext) GetParser() antlr.Parser { return s.parser }

func (s *Get_variableContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FluxVAR_IDENTIFIER, 0)
}

func (s *Get_variableContext) Get_array_element() IGet_array_elementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_array_elementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_array_elementContext)
}

func (s *Get_variableContext) Get_child() IGet_childContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_childContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_childContext)
}

func (s *Get_variableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Get_variableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Get_variableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterGet_variable(s)
	}
}

func (s *Get_variableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitGet_variable(s)
	}
}

func (s *Get_variableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitGet_variable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Get_variable() (localctx IGet_variableContext) {
	localctx = NewGet_variableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, FluxRULE_get_variable)
	p.SetState(478)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(475)
			p.Match(FluxVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(476)
			p.Get_array_element()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(477)
			p.Get_child()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMath_expressionContext is an interface to support dynamic dispatch.
type IMath_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Get_variable() IGet_variableContext
	Numberic_expression() INumberic_expressionContext
	Logical_expression() ILogical_expressionContext

	// IsMath_expressionContext differentiates from other interfaces.
	IsMath_expressionContext()
}

type Math_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMath_expressionContext() *Math_expressionContext {
	var p = new(Math_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_math_expression
	return p
}

func InitEmptyMath_expressionContext(p *Math_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_math_expression
}

func (*Math_expressionContext) IsMath_expressionContext() {}

func NewMath_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Math_expressionContext {
	var p = new(Math_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_math_expression

	return p
}

func (s *Math_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Math_expressionContext) Get_variable() IGet_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_variableContext)
}

func (s *Math_expressionContext) Numberic_expression() INumberic_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberic_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberic_expressionContext)
}

func (s *Math_expressionContext) Logical_expression() ILogical_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogical_expressionContext)
}

func (s *Math_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Math_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Math_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterMath_expression(s)
	}
}

func (s *Math_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitMath_expression(s)
	}
}

func (s *Math_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitMath_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Math_expression() (localctx IMath_expressionContext) {
	localctx = NewMath_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, FluxRULE_math_expression)
	p.SetState(483)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(480)
			p.Get_variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(481)
			p.numberic_expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(482)
			p.logical_expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGet_array_elementContext is an interface to support dynamic dispatch.
type IGet_array_elementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_IDENTIFIER() antlr.TerminalNode
	L_SQUARE() antlr.TerminalNode
	Numberic_expression() INumberic_expressionContext
	R_SQUARE() antlr.TerminalNode

	// IsGet_array_elementContext differentiates from other interfaces.
	IsGet_array_elementContext()
}

type Get_array_elementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGet_array_elementContext() *Get_array_elementContext {
	var p = new(Get_array_elementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_get_array_element
	return p
}

func InitEmptyGet_array_elementContext(p *Get_array_elementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_get_array_element
}

func (*Get_array_elementContext) IsGet_array_elementContext() {}

func NewGet_array_elementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Get_array_elementContext {
	var p = new(Get_array_elementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_get_array_element

	return p
}

func (s *Get_array_elementContext) GetParser() antlr.Parser { return s.parser }

func (s *Get_array_elementContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FluxVAR_IDENTIFIER, 0)
}

func (s *Get_array_elementContext) L_SQUARE() antlr.TerminalNode {
	return s.GetToken(FluxL_SQUARE, 0)
}

func (s *Get_array_elementContext) Numberic_expression() INumberic_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberic_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberic_expressionContext)
}

func (s *Get_array_elementContext) R_SQUARE() antlr.TerminalNode {
	return s.GetToken(FluxR_SQUARE, 0)
}

func (s *Get_array_elementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Get_array_elementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Get_array_elementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterGet_array_element(s)
	}
}

func (s *Get_array_elementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitGet_array_element(s)
	}
}

func (s *Get_array_elementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitGet_array_element(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Get_array_element() (localctx IGet_array_elementContext) {
	localctx = NewGet_array_elementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, FluxRULE_get_array_element)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(485)
		p.Match(FluxVAR_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(486)
		p.Match(FluxL_SQUARE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(487)
		p.numberic_expression(0)
	}
	{
		p.SetState(488)
		p.Match(FluxR_SQUARE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGet_childContext is an interface to support dynamic dispatch.
type IGet_childContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVAR_IDENTIFIER() []antlr.TerminalNode
	VAR_IDENTIFIER(i int) antlr.TerminalNode
	DOT() antlr.TerminalNode
	Get_array_element() IGet_array_elementContext
	Get_child() IGet_childContext

	// IsGet_childContext differentiates from other interfaces.
	IsGet_childContext()
}

type Get_childContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGet_childContext() *Get_childContext {
	var p = new(Get_childContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_get_child
	return p
}

func InitEmptyGet_childContext(p *Get_childContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_get_child
}

func (*Get_childContext) IsGet_childContext() {}

func NewGet_childContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Get_childContext {
	var p = new(Get_childContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_get_child

	return p
}

func (s *Get_childContext) GetParser() antlr.Parser { return s.parser }

func (s *Get_childContext) AllVAR_IDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(FluxVAR_IDENTIFIER)
}

func (s *Get_childContext) VAR_IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(FluxVAR_IDENTIFIER, i)
}

func (s *Get_childContext) DOT() antlr.TerminalNode {
	return s.GetToken(FluxDOT, 0)
}

func (s *Get_childContext) Get_array_element() IGet_array_elementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_array_elementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_array_elementContext)
}

func (s *Get_childContext) Get_child() IGet_childContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_childContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_childContext)
}

func (s *Get_childContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Get_childContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Get_childContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterGet_child(s)
	}
}

func (s *Get_childContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitGet_child(s)
	}
}

func (s *Get_childContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitGet_child(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Get_child() (localctx IGet_childContext) {
	localctx = NewGet_childContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, FluxRULE_get_child)
	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(490)
			p.Match(FluxVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(491)
			p.Match(FluxDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(492)
			p.Match(FluxVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(493)
			p.Match(FluxVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(494)
			p.Match(FluxDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(495)
			p.Get_array_element()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(496)
			p.Match(FluxVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(497)
			p.Match(FluxDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(498)
			p.Get_child()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(499)
			p.Get_array_element()
		}
		{
			p.SetState(500)
			p.Match(FluxDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(501)
			p.Get_child()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(503)
			p.Get_array_element()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(504)
			p.Match(FluxVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_IDENTIFIER() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	AllMath_expression() []IMath_expressionContext
	Math_expression(i int) IMath_expressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunction_callContext differentiates from other interfaces.
	IsFunction_callContext()
}

type Function_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_function_call
	return p
}

func InitEmptyFunction_callContext(p *Function_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FluxRULE_function_call
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FluxRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FluxVAR_IDENTIFIER, 0)
}

func (s *Function_callContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxL_PAREN, 0)
}

func (s *Function_callContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(FluxR_PAREN, 0)
}

func (s *Function_callContext) AllMath_expression() []IMath_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMath_expressionContext); ok {
			len++
		}
	}

	tst := make([]IMath_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMath_expressionContext); ok {
			tst[i] = t.(IMath_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Function_callContext) Math_expression(i int) IMath_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMath_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMath_expressionContext)
}

func (s *Function_callContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FluxCOMMA)
}

func (s *Function_callContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FluxCOMMA, i)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FluxListener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (s *Function_callContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FluxVisitor:
		return t.VisitFunction_call(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Flux) Function_call() (localctx IFunction_callContext) {
	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, FluxRULE_function_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(507)
		p.Match(FluxVAR_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(508)
		p.Match(FluxL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(517)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&633322996760616) != 0 {
		{
			p.SetState(509)
			p.Math_expression()
		}
		p.SetState(514)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FluxCOMMA {
			{
				p.SetState(510)
				p.Match(FluxCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(511)
				p.Math_expression()
			}

			p.SetState(516)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(519)
		p.Match(FluxR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *Flux) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 25:
		var t *Numberic_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Numberic_expressionContext)
		}
		return p.Numberic_expression_Sempred(t, predIndex)

	case 26:
		var t *Logical_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Logical_expressionContext)
		}
		return p.Logical_expression_Sempred(t, predIndex)

	case 27:
		var t *Comparative_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Comparative_expressionContext)
		}
		return p.Comparative_expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Flux) Numberic_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Flux) Logical_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Flux) Comparative_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
