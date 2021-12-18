package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParse12(t *testing.T) {
	num := FromString("[1,2]")
	require.NotNil(t, num)
	require.Equal(t, 0, num.value)

	require.NotNil(t, num.left)
	require.Equal(t, 1, num.left.value)
	require.Nil(t, num.left.left)
	require.Nil(t, num.left.right)

	require.NotNil(t, num.right)
	require.Equal(t, 2, num.right.value)
	require.Nil(t, num.right.left)
	require.Nil(t, num.right.right)

	require.Equal(t, num, num.right.parent)
	require.Equal(t, num, num.left.parent)
}

func TestParseBigDude(t *testing.T) {
	num := FromString("[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]")
	require.Equal(t, 6, num.right.left.right.left.value)
	require.Equal(t, num, num.right.left.right.left.parent.parent.parent.parent)
}

func TestReduce(t *testing.T) {
	num := FromString("[[[[[9,8],1],2],3],4]")
	num.Reduce()
	require.Equal(t, 9, num.left.left.left.right.value)

	num = FromString("[7,[6,[5,[4,[3,2]]]]]")
	expectExplodable := FromString("[3,2]")
	explodable := num.findExplodable(0)
	require.NotNil(t, explodable, "didnt find explodable")
	require.Equal(t, expectExplodable.String(), explodable.String())
	num.Reduce() // [7,[6,[5,[7,0]]]]
	require.Equal(t, 7, num.right.right.right.left.value, num.String())

	num = FromString("[[6,[5,[4,[3,2]]]],1]")
	num.Reduce() // [[6,[5,[7,0]]],3]
	require.Equal(t, "[7,0]", num.left.right.right.String())
	require.Equal(t, 3, num.right.value, num.String())

	num = FromString("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]")
	num.Reduce() // [[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]
	require.Equal(t, "[8,0]", num.left.right.right.String())
	require.Equal(t, 9, num.right.left.value, num.String())

	num1 := FromString("[[[[4,3],4],4],[7,[[8,4],9]]]")
	num2 := FromString("[1,1]")
	num = num1.Add(num2)
	expect := FromString("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")
	require.Equal(t, expect.String(), num.String())
}

func TestExplode(t *testing.T) {
	num := FromString("[[[[[9,8],1],2],3],4]")
	explodable := num.findExplodable(0)
	require.Equal(t, num.left.left.left.left, explodable)

	expect := FromString("[[[[0,9],2],3],4]")
	explodable.Explode()
	require.Equal(t, expect.left.left.left.left.value, num.left.left.left.left.value)
	require.Equal(t, expect.left.left.left.right.value, num.left.left.left.right.value)
	require.Equal(t, expect.left.left.right.value, num.left.left.right.value)

	num = FromString("[[[[0,9],2],3],4]")
	explodable = num.findExplodable(0)
	require.Nil(t, explodable)

	num = FromString("[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]")
	explodable = num.findExplodable(0)
	require.Equal(t, "[6,7]", explodable.String())
	explodable.Explode()
	require.EqualValues(t, "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", num.String())
}

func TestFindLeft(t *testing.T) {
	num := FromString("[[[[[9,8],1],2],3],4]")

	explodable := num.findExplodable(0)
	require.Nil(t, explodable.findLeft())

	four := num.right
	require.Equal(t, 3, four.findLeft().value)

	three := num.left.right
	require.Equal(t, 2, three.findLeft().value)
}

func TestFindRight(t *testing.T) {
	num := FromString("[[[[[9,8],1],2],3],4]")

	explodable := num.findExplodable(0)
	require.Equal(t, 1, explodable.findRight().value)

	four := num.right
	require.Nil(t, four.findRight())

	three := num.left.right
	require.Equal(t, 4, three.findRight().value)

	num = FromString("[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]")
	sixSeven := num.left.right.right.right
	require.Equal(t, "[6,7]", sixSeven.String())
	one := sixSeven.findRight()
	require.NotNil(t, one)
	require.Equal(t,1, one.value)
	require.Equal(t, 1, one.parent.right.value)
}

func TestBlackMagic(t *testing.T) {
	num := FromString("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]")

	num.findExplodable(0).Explode()
	require.Equal(t, "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]", num.String())
	require.Equal(t, num, num.left.parent)

	num.findExplodable(0).Explode()
	require.Equal(t, "[[[[0,7],4],[15,[0,13]]],[1,1]]", num.String())
	require.Equal(t, num, num.left.parent)

	num.findSplittable().Split()
	require.Equal(t, "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]", num.String())
	require.Equal(t, num, num.left.parent)

	num.findSplittable().Split()
	require.Equal(t, "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]", num.String())
	require.Equal(t, num, num.left.parent)

	num.findExplodable(0).Explode()
	require.Equal(t, "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", num.String())
	require.Equal(t, num, num.left.parent)
}