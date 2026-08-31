package gin

type Param struct {
	Key   string
	Value string
}

type Params []Param

func (ps Params) Get(name string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (ps Params) ByName(name string) (va string) { _ = "STUB: not implemented"; return "" }

type methodTree struct {
	method string
	root   *node
}

type methodTrees []methodTree

func (trees methodTrees) get(method string) *node { _ = "STUB: not implemented"; return nil }

func longestCommonPrefix(a, b string) int { _ = "STUB: not implemented"; return 0 }

func (n *node) addChild(child *node) { _ = "STUB: not implemented"; return }

func countParams(path string) uint16 { _ = "STUB: not implemented"; return 0 }

func countSections(path string) uint16 { _ = "STUB: not implemented"; return 0 }

type nodeType uint8

const (
	static nodeType = iota
	root
	param
	catchAll
)

type node struct {
	path      string
	indices   string
	wildChild bool
	nType     nodeType
	priority  uint32
	children  []*node
	handlers  HandlersChain
	fullPath  string
}

func (n *node) incrementChildPrio(pos int) int { _ = "STUB: not implemented"; return 0 }

func (n *node) addRoute(path string, handlers HandlersChain) { _ = "STUB: not implemented"; return }

func findWildcard(path string) (wildcard string, i int, valid bool) {
	_ = "STUB: not implemented"
	return "", 0, false
}

func (n *node) insertChild(path string, fullPath string, handlers HandlersChain) {
	_ = "STUB: not implemented"
	return
}

type nodeValue struct {
	handlers HandlersChain
	params   *Params
	tsr      bool
	fullPath string
}

type skippedNode struct {
	path        string
	node        *node
	paramsCount int16
}

func (n *node) getValue(path string, params *Params, skippedNodes *[]skippedNode, unescape bool) (value nodeValue) {
	_ = "STUB: not implemented"
	return *new(nodeValue)
}

func (n *node) findCaseInsensitivePath(path string, fixTrailingSlash bool) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func shiftNRuneBytes(rb [4]byte, n int) [4]byte { _ = "STUB: not implemented"; return [4]byte{} }

func (n *node) findCaseInsensitivePathRec(path string, ciPath []byte, rb [4]byte, fixTrailingSlash bool) []byte {
	_ = "STUB: not implemented"
	return nil
}
