package part2

import (
	"testing"
)

func TestMoveDial(t *testing.T) {
	// Define a table of test cases
	tests := []struct {
		name              string
		initialDial       int
		initialZeroCount  int
		instruction       Instruction
		expectedDial      int
		expectedZeroCount int
	}{
		// --- Basic Movement Cases (No Zero Crossing) ---
		{
			name:              "Right_Partial_NoZero_From50",
			initialDial:       50,
			initialZeroCount:  0,
			instruction:       Instruction{direction: right, steps: 25},
			expectedDial:      75,
			expectedZeroCount: 0,
		},
		{
			name:              "Left_Partial_NoZero_From50",
			initialDial:       50,
			initialZeroCount:  0,
			instruction:       Instruction{direction: left, steps: 25},
			expectedDial:      25,
			expectedZeroCount: 0,
		},
		// --- Zero Crossing Cases (Within 100 steps) ---
		{
			name:              "Right_Partial_ZeroCrossing",
			initialDial:       90, // 10 steps to 0
			initialZeroCount:  0,
			instruction:       Instruction{direction: right, steps: 15},
			expectedDial:      5, // (90 + 15) % 100 = 5
			expectedZeroCount: 1, // Crossed 100-90=10 (15 >= 10)
		},
		{
			name:              "Left_Partial_ZeroCrossing",
			initialDial:       10, // 10 steps to 0
			initialZeroCount:  0,
			instruction:       Instruction{direction: left, steps: 15},
			expectedDial:      95, // (10 - 15 + 100) % 100 = 95
			expectedZeroCount: 1,  // Crossed 10 (15 >= 10)
		},
		// --- Exactly One Zero Crossing Cases ---
		{
			name:              "Right_ExactZeroCrossing",
			initialDial:       95,
			initialZeroCount:  0,
			instruction:       Instruction{direction: right, steps: 5},
			expectedDial:      0,
			expectedZeroCount: 1, // Crossed (5 >= 100-95=5)
		},
		{
			name:              "Left_TwoFullRevolutions_NoPartialZero",
			initialDial:       50,
			initialZeroCount:  10,
			instruction:       Instruction{direction: left, steps: 200},
			expectedDial:      50, // (50 - 200 + 100) % 100 = 50
			expectedZeroCount: 12, // totalRevolutions = 200/100 = 2, partialSteps=0
		},
		// --- Combined Revolutions and Partial Zero Crossing ---
		{
			name:              "Right_Revolutions_And_PartialZero",
			initialDial:       90,
			initialZeroCount:  0,
			instruction:       Instruction{direction: right, steps: 115}, // 1 full rev + 15 steps
			expectedDial:      5,                                         // (90 + 15) % 100 = 5
			expectedZeroCount: 2,                                         // 1 from totalRevolutions, 1 from partial (15 >= 10)
		},
		{
			name:              "Left_Revolutions_And_PartialZero",
			initialDial:       10,
			initialZeroCount:  0,
			instruction:       Instruction{direction: left, steps: 315}, // 3 full revs + 15 steps
			expectedDial:      95,                                       // (10 - 15 + 100) % 100 = 95
			expectedZeroCount: 4,                                        // 3 from totalRevolutions, 1 from partial (15 >= 10)
		},
		// --- Edge Case: Steps exactly 100 (1 full revolution, partialSteps=0) ---
		{
			name:              "Right_Exact100Steps",
			initialDial:       50,
			initialZeroCount:  0,
			instruction:       Instruction{direction: right, steps: 100},
			expectedDial:      50,
			expectedZeroCount: 1, // totalRevolutions=1, partialSteps=0. The partial logic is skipped as 0 is not >= any positive number.
		},
		{
			name:              "Right_Exact100Steps",
			initialDial:       0,
			initialZeroCount:  0,
			instruction:       Instruction{direction: right, steps: 150},
			expectedDial:      50,
			expectedZeroCount: 1, // totalRevolutions=1, partialSteps=0. The partial logic is skipped as 0 is not >= any positive number.
		},
		{
			name:              "New_R90_NoZeroCrossing",
			initialDial:       50,
			initialZeroCount:  0,
			instruction:       Instruction{direction: right, steps: 90},
			expectedDial:      40, // (50 + 90) % 100 = 140 % 100 = 40
			expectedZeroCount: 1,  // 90 >= (100 - 50) is true (90 >= 50)
		},
		{
			name:              "New_R46_NoZeroCrossing",
			initialDial:       50,
			initialZeroCount:  0,
			instruction:       Instruction{direction: right, steps: 46},
			expectedDial:      96, // (50 + 46) % 100 = 96
			expectedZeroCount: 0,  // 46 >= 50 is false
		},
		{
			name:              "New_L36_NoZeroCrossing",
			initialDial:       50,
			initialZeroCount:  0,
			instruction:       Instruction{direction: left, steps: 36},
			expectedDial:      14, // (50 - 36 + 100) % 100 = 14
			expectedZeroCount: 0,  // 36 >= 50 is false
		},
		{
			name:              "New_R85_ZeroCrossing",
			initialDial:       50,
			initialZeroCount:  0,
			instruction:       Instruction{direction: right, steps: 85},
			expectedDial:      35, // (50 + 85) % 100 = 135 % 100 = 35
			expectedZeroCount: 1,  // 85 >= 50 is true
		},
		{
			name:              "New_L967_NineRevsAndZeroCrossing",
			initialDial:       50,
			initialZeroCount:  0,
			instruction:       Instruction{direction: left, steps: 967}, // 9 full revs + 67 steps left
			expectedDial:      83,                                       // (50 - 67 + 100) % 100 = 83
			expectedZeroCount: 10,                                       // 9 from totalRevolutions (967/100 = 9). Partial zero count is skipped with my fix.
		},
		{
			name:              "New_L618_SixRevsAndNoZeroCrossing",
			initialDial:       50,
			initialZeroCount:  0,
			instruction:       Instruction{direction: left, steps: 618}, // 6 full revs + 18 steps left
			expectedDial:      32,                                       // (50 - 18 + 100) % 100 = 32
			expectedZeroCount: 6,                                        // 6 from totalRevolutions (618/100 = 6). Partial zero count is skipped.
		},
		{
			name:              "New_R1_NoChange",
			initialDial:       50,
			initialZeroCount:  0,
			instruction:       Instruction{direction: right, steps: 1},
			expectedDial:      51,
			expectedZeroCount: 0,
		},
		{
			name:              "New_R99_ZeroCrossing",
			initialDial:       50,
			initialZeroCount:  0,
			instruction:       Instruction{direction: right, steps: 99},
			expectedDial:      49, // (50 + 99) % 100 = 149 % 100 = 49
			expectedZeroCount: 1,  // 99 >= 50 is true
		},
		{
			name:              "New_R54_ZeroCrossing",
			initialDial:       50,
			initialZeroCount:  0,
			instruction:       Instruction{direction: right, steps: 54},
			expectedDial:      4, // (50 + 54) % 100 = 104 % 100 = 4
			expectedZeroCount: 1, // 54 >= 50 is true
		},
		{
			name:              "New_L28_NoZeroCrossing",
			initialDial:       50,
			initialZeroCount:  0,
			instruction:       Instruction{direction: left, steps: 28},
			expectedDial:      22, // (50 - 28 + 100) % 100 = 22
			expectedZeroCount: 0,  // 28 >= 50 is false
		},
		{
			name:              "New_R74_ZeroCrossing",
			initialDial:       50,
			initialZeroCount:  0,
			instruction:       Instruction{direction: right, steps: 74},
			expectedDial:      24, // (50 + 74) % 100 = 124 % 100 = 24
			expectedZeroCount: 1,  // 74 >= 50 is true
		},
		{
			name:              "New_R74_ZeroCrossing",
			initialDial:       0,
			initialZeroCount:  0,
			instruction:       Instruction{direction: right, steps: 100},
			expectedDial:      0,
			expectedZeroCount: 1,
		},
		{
			name:              "New_R74_ZeroCrossing",
			initialDial:       0,
			initialZeroCount:  0,
			instruction:       Instruction{direction: left, steps: 100},
			expectedDial:      0,
			expectedZeroCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup: Set the global variables to the initial state for this test
			dial = tt.initialDial
			zeroCount = tt.initialZeroCount

			// Action: Call the function under test
			moveDial(tt.instruction)

			// Assert: Check the resulting values
			if dial != tt.expectedDial {
				t.Errorf("Dial after moveDial(%+v) was incorrect.\n\tGot Dial: %d, Expected Dial: %d", tt.instruction, dial, tt.expectedDial)
			}
			if zeroCount != tt.expectedZeroCount {
				t.Errorf("ZeroCount after moveDial(%+v) was incorrect.\n\tGot ZeroCount: %d, Expected ZeroCount: %d", tt.instruction, zeroCount, tt.expectedZeroCount)
			}
		})
	}
}
