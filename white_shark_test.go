package hunt_test

import (
	hunt "testdoubles"
	"testing"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, false, 20.0)
		tuna := hunt.NewTuna("Atum", 15.0)

		err := shark.Hunt(tuna)

		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}
		if shark.Hungry {
			t.Error("expected shark to no longer be hungry, but it is still hungry")
		}
		if !shark.Tired {
			t.Error("expected shark to be tired after hunting, but it is not tired")
		}
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		shark := hunt.NewWhiteShark(false, false, 20.0)
		tuna := hunt.NewTuna("Atum", 15.0)

		err := shark.Hunt(tuna)

		if err != hunt.ErrSharkIsNotHungry {
			t.Errorf("expected error %v, got: %v", hunt.ErrSharkIsNotHungry, err)
		}
		if err == nil {
			t.Errorf("expected an error when shark is not hungry, but got none")
		}
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, true, 20.0)
		tuna := hunt.NewTuna("Atum", 15.0)

		err := shark.Hunt(tuna)

		if err != hunt.ErrSharkIsTired {
			t.Errorf("expected error %v, got: %v", hunt.ErrSharkIsTired, err)
		}
		if err == nil {
			t.Errorf("expected an error when shark is tired, but got none")
		}
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, false, 20.0)
		tuna := hunt.NewTuna("Atum", 30.0)

		err := shark.Hunt(tuna)

		if err != hunt.ErrSharkIsSlower {
			t.Errorf("expected error %v, got: %v", hunt.ErrSharkIsSlower, err)
		}
		if err == nil {
			t.Errorf("expected an error when shark is slower than tuna, but got none")
		}
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, false, 20.0)
		var tuna *hunt.Tuna

		err := shark.Hunt(tuna)
		if err == nil {
			t.Error("expected an error when tuna is nil, but got none")
		}

	})
}
