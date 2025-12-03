package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTurnDial(t *testing.T) {
	d := NewDial()
	d.dir = "L"
	d.Turn(10)
	assert.Equal(t, int32(0), d.count, "they should be equal")
}

func TestTurnDial_L_CrossZero(t *testing.T) {
	d := NewDial()
	d.dir = "L"
	d.Turn(80)
	assert.Equal(t, int32(1), d.count, "they should be equal")
}

func TestTurnDial_R_CrossZero(t *testing.T) {
	d := NewDial()
	d.dir = "R"
	d.Turn(80)
	assert.Equal(t, int32(1), d.count, "they should be equal")
}

func TestTurnDial_R_CrossZero_Multiple(t *testing.T) {
	d := NewDial()
	d.dir = "R"
	d.Turn(280)
	assert.Equal(t, int32(3), d.count, "they should be equal")
}

func TestTurnDial_L_CrossZero_Multiple(t *testing.T) {
	d := NewDial()
	d.dir = "L"
	d.Turn(280)
	assert.Equal(t, int32(3), d.count, "they should be equal")
}

func TestTurnDial_BackAndForth_CrossZero(t *testing.T) {
	d := NewDial()
	d.value = 1
	d.dir = "R"
	d.Turn(100)
	assert.Equal(t, int32(1), d.count, "they should be equal")

	d.dir = "L"
	d.Turn(1)
	assert.Equal(t, int32(2), d.count, "they should be equal")
}

func TestTurnDial_BackAndForth_CrossZero_Multiple(t *testing.T) {
	d := NewDial()
	d.value = 1
	d.dir = "R"
	d.Turn(100) //lands on 1, crosses zero once
	assert.Equal(t, int32(1), d.count, "they should be equal")

	d.dir = "L"
	d.Turn(1) //lands back on zero
	assert.Equal(t, int32(2), d.count, "they should be equal")

	d.dir = "R"
	d.Turn(150) // lands on 50, crosses zero once
	assert.Equal(t, int32(3), d.count, "they should be equal")
}

func TestTurnDial_DontDoubleCountZero(t *testing.T) {
	d := NewDial()
	d.dir = "R"
	d.Turn(50) //lands on 0
	assert.Equal(t, int32(1), d.count, "they should be equal")

	d.dir = "R"
	d.Turn(101) //lands back on zero
	assert.Equal(t, int32(2), d.count, "they should be equal")
}

func Test_RL_No_Pass(t *testing.T) {
	d := NewDial()
	d.dir = "R"
	d.Turn(39) //lands on 89
	assert.Equal(t, int32(0), d.count, "they should be equal")

	d.dir = "L"
	d.Turn(20) //lands back on 69
	assert.Equal(t, int32(0), d.count, "they should be equal")
}

func Test_LR_No_Pass(t *testing.T) {
	d := NewDial()
	d.dir = "L"
	d.Turn(39) //lands on 89
	assert.Equal(t, int32(0), d.count, "they should be equal")

	d.dir = "R"
	d.Turn(20) //lands back on 69
	assert.Equal(t, int32(0), d.count, "they should be equal")
}

// func Test_Normalize_Positive(t *testing.T) {
// 	d := NewDial()
// 	d.value = d.normalize()
// 	assert.Equal(t, int32(50), d.value, "they should be equal")
// }

// func Test_Normalize_Negative(t *testing.T) {
// 	d := NewDial()
// 	d.value = -30
// 	d.value = d.normalize()
// 	assert.Equal(t, int32(70), d.value, "they should be equal")
// }

func Test_2(t *testing.T) {
	d := NewDial()
	d.value = 32
	d.dir = "R"
	d.Turn(68) //lands on 0
	assert.Equal(t, int32(1), d.count, "they should be equal")

	d.dir = "R"
	d.Turn(97) //lands back on 97
	assert.Equal(t, int32(1), d.count, "they should be equal")

	d.dir = "R"
	d.Turn(3) //lands back on 0
	assert.Equal(t, int32(2), d.count, "they should be equal")

	d.dir = "L"
	d.Turn(259) //lands back on 41
	assert.Equal(t, int32(4), d.count, "they should be equal")
}
