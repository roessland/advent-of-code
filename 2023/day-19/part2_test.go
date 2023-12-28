package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestN(t *testing.T) {
	// {
	// 	wfs := map[string]Workflow{
	// 		"asdf": {
	// 			Name: "asdf",
	// 			Rules: []Rule{
	// 				{
	// 					Set:     NewSet(),
	// 					Cat:     "x",
	// 					OutName: "A",
	// 				},
	// 			},
	// 		},
	// 	}
	// 	require.Equal(t, 4000*4000*4000*4000, N(wfs, "asdf", NewSets()))
	// }
	//
	// {
	// 	wfs := map[string]Workflow{
	// 		"asdf": {
	// 			Name: "asdf",
	// 			Rules: []Rule{
	// 				MakeRule(RuleInput{
	// 					Cat:       "x",
	// 					Output:    "R",
	// 					Op:        "<",
	// 					Threshold: 2001,
	// 				}),
	// 				{
	// 					Set:     NewSet(),
	// 					Cat:     "x",
	// 					OutName: "A",
	// 				},
	// 			},
	// 		},
	// 	}
	// 	require.Equal(t, 2000*4000*4000*4000, N(wfs, "asdf", NewSets()))
	// }
	//
	{
		wfs := map[string]Workflow{
			"aaa": {
				Name: "aaa",
				Rules: []Rule{
					MakeRule(RuleInput{
						Cat:       "m",
						Output:    "bbb",
						Op:        "<",
						Threshold: 2001,
					}),
					{
						Set:     NewSet(),
						OutName: "R",
					},
				},
			},

			"bbb": {
				Name: "bbb",
				Rules: []Rule{
					MakeRule(RuleInput{
						Cat:       "m",
						Output:    "R",
						Op:        ">",
						Threshold: 1000,
					}),
					{
						Set:     NewSet(),
						OutName: "A",
					},
				},
			},
		}
		// < 2001 && <= 1000
		require.Equal(t, 4000*4000*4000*4000, N(wfs, "A", NewSets()))
		require.Equal(t, 0, N(wfs, "R", NewSets()))
		require.Equal(t, 1000*4000*4000*4000, N(wfs, "bbb", NewSets()))
		require.Equal(t, 1000*4000*4000*4000, N(wfs, "aaa", NewSets()))
	}
}
