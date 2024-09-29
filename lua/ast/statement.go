package ast

import (
	"github.com/raiguard/luapls/lua/token"
)

type Statement interface {
	Node
	statementNode()
}

type AssignmentStatement struct {
	Vars   Punctuated[Expression]
	Assign Unit
	Exps   Punctuated[Expression]
}

func (as *AssignmentStatement) statementNode() {}
func (as *AssignmentStatement) Pos() token.Pos {
	return as.Vars.Pos()
}
func (as *AssignmentStatement) End() token.Pos {
	return as.Exps.End()
}
func (as *AssignmentStatement) Leaves() []Node {
	return []Node{&as.Vars, &as.Exps}
}

type BreakStatement Unit

func (bs *BreakStatement) statementNode() {}
func (bs *BreakStatement) Pos() token.Pos {
	return bs.Token.Pos
}
func (bs *BreakStatement) End() token.Pos {
	return bs.Token.End()
}
func (bs *BreakStatement) Leaves() (n []Node) {
	return n
}

type DoStatement struct {
	DoTok  Unit
	Body   Block
	EndTok Unit
}

func (ds *DoStatement) statementNode() {}
func (ds *DoStatement) Pos() token.Pos {
	return ds.DoTok.Pos()
}
func (ds *DoStatement) End() token.Pos {
	return ds.EndTok.End()
}
func (bs *DoStatement) Leaves() (n []Node) {
	return n
}

type ForStatement struct {
	ForTok    Unit
	Name      *Identifier
	AssignTok Unit
	Start     Pair[Expression]
	Finish    Pair[Expression]
	Step      *Pair[Expression]
	DoTok     Unit
	Body      Block
	EndTok    Unit
}

func (fs *ForStatement) statementNode() {}
func (fs *ForStatement) Pos() token.Pos {
	return fs.ForTok.Pos()
}
func (fs *ForStatement) End() token.Pos {
	return fs.EndTok.End()
}
func (bs *ForStatement) Leaves() (n []Node) {
	return n
}

type ForInStatement struct {
	ForTok Unit
	Names  Punctuated[*Identifier]
	InTok  Unit
	Exps   Punctuated[Expression]
	DoTok  Unit
	Body   Block
	EndTok Unit
}

func (fs *ForInStatement) statementNode() {}
func (fs *ForInStatement) Pos() token.Pos {
	return fs.ForTok.Pos()
}
func (fs *ForInStatement) End() token.Pos {
	return fs.EndTok.End()
}

func (fs *ForInStatement) Leaves() []Node {
	return []Node{&fs.Names, &fs.Exps, &fs.Body}
}

type FunctionStatement struct {
	LocalTok   *Unit
	FuncTok    Unit
	Name       Expression
	LeftParen  Unit
	Params     Punctuated[*Identifier]
	Vararg     *Unit
	RightParen Unit
	Body       Block
	EndTok     Unit
}

func (fs *FunctionStatement) statementNode() {}
func (fs *FunctionStatement) Pos() token.Pos {
	if fs.LocalTok != nil {
		return fs.LocalTok.Pos()
	}
	return fs.Name.Pos()
}
func (fs *FunctionStatement) End() token.Pos {
	return fs.EndTok.End()
}
func (fs *FunctionStatement) Leaves() []Node {
	return []Node{fs.Name, &fs.Params, &fs.Body}
}

type GotoStatement struct {
	GotoTok Unit
	Name    *Identifier
}

func (gs *GotoStatement) statementNode() {}
func (gs *GotoStatement) Pos() token.Pos {
	return gs.GotoTok.Pos()
}
func (gs *GotoStatement) End() token.Pos {
	return gs.Name.End()
}
func (gs *GotoStatement) Leaves() []Node {
	return []Node{gs.Name}
}

type IfStatement struct {
	IfTok   Unit // TODO: Remove this in favor of Clauses[0].LeadingTok, since it will always exist
	Clauses []*IfClause
	EndTok  Unit
}

func (is *IfStatement) statementNode() {}
func (is *IfStatement) Pos() token.Pos {
	return is.IfTok.Pos()
}
func (is *IfStatement) End() token.Pos {
	return is.EndTok.Pos()
}
func (is *IfStatement) Leaves() (n []Node) {
	for _, i := range is.Clauses {
		n = append(n, i.Leaves()...)
	}
	return n
}

type IfClause struct {
	LeadingTok Unit
	Condition  Expression
	ThenTok    *Unit
	Body       Block
}

func (ic *IfClause) statementNode() {}
func (ic *IfClause) Pos() token.Pos {
	return ic.LeadingTok.Pos()
}
func (ic *IfClause) End() token.Pos {
	return ic.Body.End()
}
func (ic *IfClause) Leaves() []Node {
	return []Node{ic.Condition, &ic.Body}
}

type LabelStatement struct {
	LeadingLabelTok  Unit
	Name             *Identifier
	TrailingLabelTok Unit
}

func (ls *LabelStatement) statementNode() {}
func (ls *LabelStatement) Pos() token.Pos {
	return ls.LeadingLabelTok.Pos()
}
func (ls *LabelStatement) End() token.Pos {
	return ls.TrailingLabelTok.End()
}
func (ls *LabelStatement) Leaves() (n []Node) {
	return
}
func (ls *LabelStatement) leaf() {}

type LocalStatement struct {
	LocalTok  Unit
	Names     Punctuated[*Identifier]
	AssignTok *Unit
	Exps      *Punctuated[Expression]
}

func (ls *LocalStatement) statementNode() {}
func (ls *LocalStatement) Pos() token.Pos {
	return ls.LocalTok.Pos()
}
func (ls *LocalStatement) End() token.Pos {
	if ls.Exps != nil {
		return ls.Exps.End()
	}
	if ls.AssignTok != nil {
		return ls.AssignTok.End()
	}
	return ls.Names.End()
}
func (ls *LocalStatement) Leaves() []Node {
	return []Node{&ls.Names, ls.Exps}
}

type RepeatStatement struct {
	RepeatTok Unit
	Body      Block
	UntilTok  Unit
	Condition Expression
}

func (rs *RepeatStatement) statementNode() {}
func (rs *RepeatStatement) Pos() token.Pos {
	return rs.RepeatTok.Pos()
}
func (rs *RepeatStatement) End() token.Pos {
	return rs.Condition.End()
}
func (rs *RepeatStatement) Leaves() []Node {
	return []Node{&rs.Body, rs.Condition}
}

type ReturnStatement struct {
	ReturnTok Unit
	Exps      *Punctuated[Expression]
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) Pos() token.Pos {
	return rs.ReturnTok.Pos()
}
func (rs *ReturnStatement) End() token.Pos {
	if rs.Exps != nil {
		return rs.Exps.End()
	}
	return rs.ReturnTok.End()
}
func (rs *ReturnStatement) Leaves() []Node {
	return []Node{rs.Exps}
}

type SemicolonStatement Unit

func (ss *SemicolonStatement) statementNode() {}
func (ss *SemicolonStatement) Pos() token.Pos {
	return ss.Token.Pos
}
func (ss *SemicolonStatement) End() token.Pos {
	return ss.Token.End()
}
func (ss *SemicolonStatement) Leaves() (n []Node) {
	return
}

type WhileStatement struct {
	WhileTok  Unit
	Condition Expression
	DoTok     Unit
	Body      Block
	EndTok    Unit
}

func (ws *WhileStatement) statementNode() {}
func (ws *WhileStatement) Pos() token.Pos {
	return ws.WhileTok.Pos()
}
func (ws *WhileStatement) End() token.Pos {
	return ws.EndTok.End()
}
func (ws *WhileStatement) Leaves() []Node {
	return []Node{ws.Condition, &ws.Body}
}
