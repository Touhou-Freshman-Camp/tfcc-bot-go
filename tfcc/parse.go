package tfcc

import "strings"

type trieNode struct {
	child map[rune]*trieNode
	value func(*ParseResult)
}

func newTrieNode() *trieNode {
	return &trieNode{child: make(map[rune]*trieNode)}
}

type trie struct {
	root *trieNode
}

func (t *trie) putIfAbsent(key string, value func(*ParseResult)) bool {
	if value == nil {
		panic("cannot put a nil value")
	}
	if t.root == nil {
		t.root = newTrieNode()
	}
	node := t.root
	for _, c := range key {
		n, ok := node.child[c]
		if ok {
			node = n
		} else {
			newNode := newTrieNode()
			node.child[c] = newNode
			node = newNode
		}
	}
	if node.value != nil {
		return false
	}
	node.value = value
	return true
}

func (t *trie) get(key string) func(*ParseResult) {
	if t.root == nil {
		return nil
	}
	node := t.root
	for _, c := range key {
		n, ok := node.child[c]
		if ok {
			node = n
		} else {
			return nil
		}
	}
	return node.value
}

type ParseResult struct {
	Work, Rank              string
	Route, Character, CType map[string]struct{}
	AllSpell                bool
}

func newParseResult() *ParseResult {
	return &ParseResult{
		Route:     make(map[string]struct{}),
		Character: make(map[string]struct{}),
		CType:     make(map[string]struct{}),
	}
}

var workDict = &trie{}
var otherDict = &trie{}

func addWorkMap(result string, represent ...string) {
	for _, s := range represent {
		if !workDict.putIfAbsent(strings.ToLower(s), func(res *ParseResult) { parseWork(res, result) }) {
			panic("repeated trie keys: " + s)
		}
	}
}

func addOtherMap(f func(*ParseResult, string), result string, represent ...string) {
	for _, s := range represent {
		if !otherDict.putIfAbsent(strings.ToLower(s), func(res *ParseResult) { f(res, result) }) {
			panic("repeated trie keys: " + s)
		}
	}
}

func parseWork(res *ParseResult, result string) {
	if len(res.Work) == 0 {
		res.Work = result
	}
}

func parseRank(res *ParseResult, result string) {
	if len(res.Rank) == 0 {
		res.Rank = result
	}
}

func parseRoute(res *ParseResult, result string) {
	res.Route[result] = struct{}{}
}

func parseCharacter(res *ParseResult, result string) {
	res.Character[result] = struct{}{}
}

func parseCType(res *ParseResult, result string) {
	res.CType[result] = struct{}{}
}

func parseCharacterCType(res *ParseResult, result string) {
	switch result[len(result)-2:] {
	case "SA":
		res.CType["Spring"] = struct{}{}
		parseCharacter(res, result[:len(result)-2])
	case "SB":
		res.CType["Summer"] = struct{}{}
		parseCharacter(res, result[:len(result)-2])
	case "SC":
		res.CType["Autumn"] = struct{}{}
		parseCharacter(res, result[:len(result)-2])
	case "SD":
		res.CType["Winter"] = struct{}{}
		parseCharacter(res, result[:len(result)-2])
	default:
		switch result[len(result)-1] {
		case 'W':
			res.CType["Wolf"] = struct{}{}
			parseCharacter(res, result[:len(result)-1])
		case 'O':
			res.CType["Otter"] = struct{}{}
			parseCharacter(res, result[:len(result)-1])
		case 'E':
			res.CType["Eagle"] = struct{}{}
			parseCharacter(res, result[:len(result)-1])
		}
	}
}

func init() {
	addWorkMap("6", "???", "?????????", "hmx", "th6", "th06", "EoSD")
	addWorkMap("7", "???", "?????????", "yym", "th7", "th07", "PCB")
	addWorkMap("8", "???", "?????????", "yyc", "th8", "th08", "IN")
	addWorkMap("9", "???", "?????????", "hyz", "th9", "th09", "PoFV")
	addWorkMap("10", "???", "?????????", "fsl", "th10", "MoF")
	addWorkMap("11", "???", "???", "?????????", "dld", "th11", "SA")
	addWorkMap("12", "???", "???", "?????????", "xlc", "th12", "UFO")
	addWorkMap("128", "???", "?????????", "dzz", "th128", "128")
	addWorkMap("13", "???", "???", "?????????", "slm", "th13", "TD")
	addWorkMap("14", "???", "???", "?????????", "hzc", "th14", "DDC")
	addWorkMap("15", "???", "?????????", "gzz", "th15", "LoLK")
	addWorkMap("16", "???", "???", "?????????", "tkz", "th16", "HSiFS")
	addWorkMap("17", "???", "?????????", "gxs", "th17", "WBaWC")
	addWorkMap("18", "???", "???", "?????????", "hld", "th18", "UM")
	addOtherMap(parseRank, "Easy", "e")
	addOtherMap(parseRank, "Normal", "n")
	addOtherMap(parseRank, "Hard", "h")
	addOtherMap(parseRank, "Lunatic", "l")
	addOtherMap(parseRank, "Extra", "ex", "Phantasm", "ph")
	addOtherMap(parseCharacter, "Reimu", "???", "???", "??????", "?????????", "??????", "??????", "Reimu", "????????????")
	addOtherMap(parseCharacter, "Marisa", "???", "?????????", "m", "??????", "????????????", "Marisa")
	addOtherMap(parseCharacter, "Sakuya", "???", "??????", "s", "16", "??????", "?????????", "???16", "Sakuya")
	addOtherMap(parseCharacter, "Sanae", "???", "??????", "Sanae")
	addOtherMap(parseCharacter, "Youmu", "???", "??????", "??????", "?????????", "Youmu")
	addOtherMap(parseCharacter, "RY", "?????????", "RY", "????????????")
	addOtherMap(parseCharacter, "MA", "?????????", "MA", "?????????????????????")
	addOtherMap(parseCharacter, "SR", "?????????", "SR", "?????????????????????")
	addOtherMap(parseCharacter, "YY", "?????????", "YY", "??????????????????")
	addOtherMap(parseCharacter, "Yukari", "???", "?????????", "??????", "Yukari")
	addOtherMap(parseCharacter, "Alice", "??????", "?????????", "?????????", "??????", "????????????", "Alice", "???????????????")
	addOtherMap(parseCharacter, "Remilia", "??????", "????????????", "?????????", "???????????????", "Remilia", "??????????????????")
	addOtherMap(parseCharacter, "Yuyuko", "?????????", "uuz", "????????????", "Yuyuko", "??????????????????")
	addOtherMap(parseCharacter, "Reisen", "??????", "??????", "??????", "Reisen")
	addOtherMap(parseCharacter, "Cirno", "?????????", "???", "Cirno")
	addOtherMap(parseCharacter, "Aya", "???", "??????", "????????????", "Aya")
	addOtherMap(parseCType, "A", "A")
	addOtherMap(parseCType, "B", "B")
	addOtherMap(parseCType, "C", "C")
	addOtherMap(parseCType, "Spring", "???")
	addOtherMap(parseCType, "Summer", "???")
	addOtherMap(parseCType, "Autumn", "???")
	addOtherMap(parseCType, "Winter", "???")
	addOtherMap(parseCType, "Wolf", "???")
	addOtherMap(parseCType, "Otter", "???")
	addOtherMap(parseCType, "Eagle", "???")
	addOtherMap(parseRoute, "6A", "6A")
	addOtherMap(parseRoute, "6B", "6B")
	for _, cType := range []string{"SA", "SB", "SC", "SD"} {
		for _, ch := range []string{"Reimu", "Marisa", "Cirno", "Aya"} {
			addOtherMap(parseCharacterCType, ch+cType, ch+cType)
		}
	}
	for _, cType := range []string{"W", "O", "E"} {
		for _, ch := range []string{"Reimu", "Marisa", "Youmu"} {
			addOtherMap(parseCharacterCType, ch+cType, ch+cType)
		}
	}
}

func tryParse(res *ParseResult, t *trie, s *string, nLen int) bool {
	ref := []rune(*s)
	length := len(ref)
	for i := range ref {
		m := nLen
		if i+m > length {
			m = length - i
		}
		for n := m; n > 0; n-- {
			f := t.get(string(ref[i : i+n]))
			if f != nil {
				f(res)
				*s = string(append(ref[:i], ref[i+n:]...))
				return true
			}
		}
	}
	return false
}

func ParseMsg(s string) *ParseResult {
	if len(s) == 0 {
		panic("s is empty")
	}
	length := len(s)
	s = strings.ToLower(s)
	res := newParseResult()
	tryParse(res, workDict, &s, length)
	for i := 0; i < 10; i++ {
		if !tryParse(res, otherDict, &s, length) {
			break
		}
	}
	for i := 0; i < 10; i++ {
		if !tryParse(res, workDict, &s, length) {
			break
		}
	}
	if strings.Contains(s, "??????") {
		res.AllSpell = true
		s = strings.Replace(s, "??????", "", 1)
	}
	return res
}
