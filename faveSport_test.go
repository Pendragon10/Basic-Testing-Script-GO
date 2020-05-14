package main

import "testing"

func TestFavororiteSport(t *testing.T) {

	//Testing for valid argument
	fSport := FavoriteSport("Football")

	if fSport != "Football: 100/10" {
		t.Errorf("Failed: Expected %v, got %v", "Football: 100/10", fSport)
	} else {
		t.Logf("Success: Expected %v, got %v", "Football: 100/10", fSport)
	}

	//Testing for invalid argument
	fSportInvalid := FavoriteSport("Ew Sports")

	if fSportInvalid != "Invalid Sport or Argument" {
		t.Errorf("Failed: Expected %v, got %v", "Invalid Sport or Argument", fSportInvalid)
	} else {
		t.Logf("Success: Expected %v, got %v", "Invalid Sport or Argument", fSportInvalid)
	}
}
