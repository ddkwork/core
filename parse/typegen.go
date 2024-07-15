// Code generated by "core generate -add-types"; DO NOT EDIT.

package parse

import (
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/parse.FileState", IDName: "file-state", Doc: "FileState contains the full lexing and parsing state information for a given file.\nIt is the master state record for everything that happens in parse.  One of these\nshould be maintained for each file; texteditor.Buf has one as ParseState field.\n\nSeparate State structs are maintained for each stage (Lexing, PassTwo, Parsing) and\nthe final output of Parsing goes into the AST and Syms fields.\n\nThe Src lexer.File field maintains all the info about the source file, and the basic\ntokenized version of the source produced initially by lexing and updated by the\nremaining passes.  It has everything that is maintained at a line-by-line level.", Fields: []types.Field{{Name: "Src", Doc: "the source to be parsed -- also holds the full lexed tokens"}, {Name: "LexState", Doc: "state for lexing"}, {Name: "TwoState", Doc: "state for second pass nesting depth and EOS matching"}, {Name: "ParseState", Doc: "state for parsing"}, {Name: "AST", Doc: "ast output tree from parsing"}, {Name: "Syms", Doc: "symbols contained within this file -- initialized at start of parsing and created by AddSymbol or PushNewScope actions.  These are then processed after parsing by the language-specific code, via Lang interface."}, {Name: "ExtSyms", Doc: "External symbols that are entirely maintained in a language-specific way by the Lang interface code.  These are only here as a convenience and are not accessed in any way by the language-general parse code."}, {Name: "SymsMu", Doc: "mutex protecting updates / reading of Syms symbols"}, {Name: "WaitGp", Doc: "waitgroup for coordinating processing of other items"}, {Name: "AnonCtr", Doc: "anonymous counter -- counts up"}, {Name: "PathMap", Doc: "path mapping cache -- for other files referred to by this file, this stores the full path associated with a logical path (e.g., in go, the logical import path -> local path with actual files) -- protected for access from any thread"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/parse.FileStates", IDName: "file-states", Doc: "FileStates contains two FileState's: one is being processed while the\nother is being used externally.  The FileStates maintains\na common set of file information set in each of the FileState items when\nthey are used.", Fields: []types.Field{{Name: "Filename", Doc: "the filename"}, {Name: "Sup", Doc: "the known file type, if known (typically only known files are processed)"}, {Name: "BasePath", Doc: "base path for reporting file names -- this must be set externally e.g., by gide for the project root path"}, {Name: "DoneIndex", Doc: "index of the state that is done"}, {Name: "FsA", Doc: "one filestate"}, {Name: "FsB", Doc: "one filestate"}, {Name: "SwitchMu", Doc: "mutex locking the switching of Done vs. Proc states"}, {Name: "ProcMu", Doc: "mutex locking the parsing of Proc state -- reading states can happen fine with this locked, but no switching"}, {Name: "Meta", Doc: "extra meta data associated with this FileStates"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/parse.Language", IDName: "language", Doc: "Language provides a general interface for language-specific management\nof the lexing, parsing, and symbol lookup process.\nThe parse lexer and parser machinery is entirely language-general\nbut specific languages may need specific ways of managing these\nprocesses, and processing their outputs, to best support the\nfeatures of those languages.  That is what this interface provides.\n\nEach language defines a type supporting this interface, which is\nin turn registered with the StdLangProperties map.  Each supported\nlanguage has its own .go file in this parse package that defines its\nown implementation of the interface and any other associated\nfunctionality.\n\nThe Language is responsible for accessing the appropriate [Parser] for this\nlanguage (initialized and managed via LangSupport.OpenStandard() etc)\nand the [FileState] structure contains all the input and output\nstate information for a given file.\n\nThis interface is likely to evolve as we expand the range of supported\nlanguages.", Methods: []types.Method{{Name: "Parser", Doc: "Parser returns the [Parser] for this language", Returns: []string{"Parser"}}, {Name: "ParseFile", Doc: "ParseFile does the complete processing of a given single file, given by txt bytes,\nas appropriate for the language -- e.g., runs the lexer followed by the parser, and\nmanages any symbol output from parsing as appropriate for the language / format.\nThis is to be used for files of \"primary interest\" -- it does full type inference\nand symbol resolution etc.  The Proc() FileState is locked during parsing,\nand Switch is called after, so Done() will contain the processed info after this call.\nIf txt is nil then any existing source in fs is used.", Args: []string{"fs", "txt"}}, {Name: "HighlightLine", Doc: "HighlightLine does the lexing and potentially parsing of a given line of the file,\nfor purposes of syntax highlighting -- uses Done() FileState of existing context\nif available from prior lexing / parsing. Line is in 0-indexed \"internal\" line indexes,\nand provides relevant context for the overall parsing, which is performed\non the given line of text runes, and also updates corresponding source in FileState\n(via a copy).  If txt is nil then any existing source in fs is used.", Args: []string{"fs", "line", "txt"}, Returns: []string{"Line"}}, {Name: "CompleteLine", Doc: "CompleteLine provides the list of relevant completions for given text\nwhich is at given position within the file.\nTypically the language will call ParseLine on that line, and use the AST\nto guide the selection of relevant symbols that can complete the code at\nthe given point.", Args: []string{"fs", "text", "pos"}, Returns: []string{"Matches"}}, {Name: "CompleteEdit", Doc: "CompleteEdit returns the completion edit data for integrating the\nselected completion into the source", Args: []string{"fs", "text", "cp", "comp", "seed"}, Returns: []string{"ed"}}, {Name: "Lookup", Doc: "Lookup returns lookup results for given text which is at given position\nwithin the file.  This can either be a file and position in file to\nopen and view, or direct text to show.", Args: []string{"fs", "text", "pos"}, Returns: []string{"Lookup"}}, {Name: "IndentLine", Doc: "IndentLine returns the indentation level for given line based on\nprevious line's indentation level, and any delta change based on\ne.g., brackets starting or ending the previous or current line, or\nother language-specific keywords.  See lexer.BracketIndentLine for example.\nIndent level is in increments of tabSz for spaces, and tabs for tabs.\nOperates on rune source with markup lex tags per line.", Args: []string{"fs", "src", "tags", "ln", "tabSz"}, Returns: []string{"pInd", "delInd", "pLn", "ichr"}}, {Name: "AutoBracket", Doc: "AutoBracket returns what to do when a user types a starting bracket character\n(bracket, brace, paren) while typing.\npos = position where bra will be inserted, and curLn is the current line\nmatch = insert the matching ket, and newLine = insert a new line.", Args: []string{"fs", "bra", "pos", "curLn"}, Returns: []string{"match", "newLine"}}, {Name: "ParseDir", Doc: "ParseDir does the complete processing of a given directory, optionally including\nsubdirectories, and optionally forcing the re-processing of the directory(s),\ninstead of using cached symbols.  Typically the cache will be used unless files\nhave a more recent modification date than the cache file.  This returns the\nlanguage-appropriate set of symbols for the directory(s), which could then provide\nthe symbols for a given package, library, or module at that path.", Args: []string{"fs", "path", "opts"}, Returns: []string{"Symbol"}}, {Name: "LexLine", Doc: "LexLine is a lower-level call (mostly used internally to the language) that\ndoes just the lexing of a given line of the file, using existing context\nif available from prior lexing / parsing.\nLine is in 0-indexed \"internal\" line indexes.\nThe rune source is updated from the given text if non-nil.", Args: []string{"fs", "line", "txt"}, Returns: []string{"Line"}}, {Name: "ParseLine", Doc: "ParseLine is a lower-level call (mostly used internally to the language) that\ndoes complete parser processing of a single line from given file, and returns\nthe FileState for just that line.  Line is in 0-indexed \"internal\" line indexes.\nThe rune source information is assumed to have already been updated in FileState\nExisting context information from full-file parsing is used as appropriate, but\nthe results will NOT be used to update any existing full-file AST representation --\nshould call ParseFile to update that as appropriate.", Args: []string{"fs", "line"}, Returns: []string{"FileState"}}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/parse.LanguageDirOptions", IDName: "language-dir-options", Doc: "LanguageDirOptions provides options for the [Language.ParseDir] method", Fields: []types.Field{{Name: "Subdirs", Doc: "process subdirectories -- otherwise not"}, {Name: "Rebuild", Doc: "rebuild the symbols by reprocessing from scratch instead of using cache"}, {Name: "Nocache", Doc: "do not update the cache with results from processing"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/parse.LanguageFlags", IDName: "language-flags", Doc: "LanguageFlags are special properties of a given language"})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/parse.LanguageProperties", IDName: "language-properties", Doc: "LanguageProperties contains properties of languages supported by the parser\nframework", Fields: []types.Field{{Name: "Known", Doc: "known language -- must be a supported one from Known list"}, {Name: "CommentLn", Doc: "character(s) that start a single-line comment -- if empty then multi-line comment syntax will be used"}, {Name: "CommentSt", Doc: "character(s) that start a multi-line comment or one that requires both start and end"}, {Name: "CommentEd", Doc: "character(s) that end a multi-line comment or one that requires both start and end"}, {Name: "Flags", Doc: "special properties for this language -- as an explicit list of options to make them easier to see and set in defaults"}, {Name: "Lang", Doc: "Lang interface for this language"}, {Name: "Parser", Doc: "parser for this language -- initialized in OpenStandard"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/parse.LanguageSupporter", IDName: "language-supporter", Doc: "LanguageSupporter provides general support for supported languages.\ne.g., looking up lexers and parsers by name.\nAlso implements the lexer.LangLexer interface to provide access to other\nGuest Lexers"})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/parse.Parser", IDName: "parser", Doc: "Parser is the overall parser for managing the parsing", Fields: []types.Field{{Name: "Lexer", Doc: "lexer rules for first pass of lexing file"}, {Name: "PassTwo", Doc: "second pass after lexing -- computes nesting depth and EOS finding"}, {Name: "Parser", Doc: "parser rules for parsing lexed tokens"}, {Name: "Filename", Doc: "file name for overall parser (not file being parsed!)"}, {Name: "ReportErrs", Doc: "if true, reports errors after parsing, to stdout"}, {Name: "ModTime", Doc: "when loaded from file, this is the modification time of the parser -- re-processes cache if parser is newer than cached files"}}})
