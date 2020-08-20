package main

import "testing"

func TestLoopExercisesDir(t *testing.T) {
	ret := LoopExercisesDir(exercisesDir)
	t.Log(ret)
}
