// Code generated by "core generate ./testdata"; DO NOT EDIT.

package testdata

import (
	"cogentcore.org/core/tree"
	"cogentcore.org/core/types"
)

// TestNodeType is the [types.Type] for [TestNode]
var TestNodeType = types.AddType(&types.Type{Name: "cogentcore.org/core/tree/testdata.TestNode", IDName: "test-node", Embeds: []types.Field{{Name: "NodeBase"}}, Instance: &TestNode{}})

// NewTestNode adds a new [TestNode] with the given name to the given parent:
func NewTestNode(parent tree.Node, name ...string) *TestNode {
	return parent.NewChild(TestNodeType, name...).(*TestNode)
}

// NodeType returns the [*types.Type] of [TestNode]
func (t *TestNode) NodeType() *types.Type { return TestNodeType }

// New returns a new [*TestNode] value
func (t *TestNode) New() tree.Node { return &TestNode{} }

// NodeEmbedType is the [types.Type] for [NodeEmbed]
var NodeEmbedType = types.AddType(&types.Type{Name: "cogentcore.org/core/tree/testdata.NodeEmbed", IDName: "node-embed", Doc: "NodeEmbed embeds tree.Node and adds a couple of fields.\nAlso has a directive processed by typegen.", Directives: []types.Directive{{Tool: "direct", Directive: "value"}}, Embeds: []types.Field{{Name: "NodeBase"}}, Fields: []types.Field{{Name: "Mbr1"}, {Name: "Mbr2"}}, Instance: &NodeEmbed{}})

// NewNodeEmbed adds a new [NodeEmbed] with the given name to the given parent:
// NodeEmbed embeds tree.Node and adds a couple of fields.
// Also has a directive processed by typegen.
func NewNodeEmbed(parent tree.Node, name ...string) *NodeEmbed {
	return parent.NewChild(NodeEmbedType, name...).(*NodeEmbed)
}

// NodeType returns the [*types.Type] of [NodeEmbed]
func (t *NodeEmbed) NodeType() *types.Type { return NodeEmbedType }

// New returns a new [*NodeEmbed] value
func (t *NodeEmbed) New() tree.Node { return &NodeEmbed{} }

// SetMbr1 sets the [NodeEmbed.Mbr1]
func (t *NodeEmbed) SetMbr1(v string) *NodeEmbed { t.Mbr1 = v; return t }

// SetMbr2 sets the [NodeEmbed.Mbr2]
func (t *NodeEmbed) SetMbr2(v int) *NodeEmbed { t.Mbr2 = v; return t }

// NodeFieldType is the [types.Type] for [NodeField]
var NodeFieldType = types.AddType(&types.Type{Name: "cogentcore.org/core/tree/testdata.NodeField", IDName: "node-field", Embeds: []types.Field{{Name: "NodeEmbed"}}, Fields: []types.Field{{Name: "Field1"}}, Instance: &NodeField{}})

// NewNodeField adds a new [NodeField] with the given name to the given parent:
func NewNodeField(parent tree.Node, name ...string) *NodeField {
	return parent.NewChild(NodeFieldType, name...).(*NodeField)
}

// NodeType returns the [*types.Type] of [NodeField]
func (t *NodeField) NodeType() *types.Type { return NodeFieldType }

// New returns a new [*NodeField] value
func (t *NodeField) New() tree.Node { return &NodeField{} }

// SetField1 sets the [NodeField.Field1]
func (t *NodeField) SetField1(v NodeEmbed) *NodeField { t.Field1 = v; return t }

// SetMbr1 sets the [NodeField.Mbr1]
func (t *NodeField) SetMbr1(v string) *NodeField { t.Mbr1 = v; return t }

// SetMbr2 sets the [NodeField.Mbr2]
func (t *NodeField) SetMbr2(v int) *NodeField { t.Mbr2 = v; return t }

// NodeField2Type is the [types.Type] for [NodeField2]
var NodeField2Type = types.AddType(&types.Type{Name: "cogentcore.org/core/tree/testdata.NodeField2", IDName: "node-field2", Embeds: []types.Field{{Name: "NodeField"}}, Fields: []types.Field{{Name: "Field2"}}, Instance: &NodeField2{}})

// NewNodeField2 adds a new [NodeField2] with the given name to the given parent:
func NewNodeField2(parent tree.Node, name ...string) *NodeField2 {
	return parent.NewChild(NodeField2Type, name...).(*NodeField2)
}

// NodeType returns the [*types.Type] of [NodeField2]
func (t *NodeField2) NodeType() *types.Type { return NodeField2Type }

// New returns a new [*NodeField2] value
func (t *NodeField2) New() tree.Node { return &NodeField2{} }

// SetField2 sets the [NodeField2.Field2]
func (t *NodeField2) SetField2(v NodeEmbed) *NodeField2 { t.Field2 = v; return t }

// SetMbr1 sets the [NodeField2.Mbr1]
func (t *NodeField2) SetMbr1(v string) *NodeField2 { t.Mbr1 = v; return t }

// SetMbr2 sets the [NodeField2.Mbr2]
func (t *NodeField2) SetMbr2(v int) *NodeField2 { t.Mbr2 = v; return t }

// SetField1 sets the [NodeField2.Field1]
func (t *NodeField2) SetField1(v NodeEmbed) *NodeField2 { t.Field1 = v; return t }