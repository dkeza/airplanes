// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Jets", testJets)
	t.Run("Languages", testLanguages)
	t.Run("Pilots", testPilots)
}

func TestDelete(t *testing.T) {
	t.Run("Jets", testJetsDelete)
	t.Run("Languages", testLanguagesDelete)
	t.Run("Pilots", testPilotsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Jets", testJetsQueryDeleteAll)
	t.Run("Languages", testLanguagesQueryDeleteAll)
	t.Run("Pilots", testPilotsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Jets", testJetsSliceDeleteAll)
	t.Run("Languages", testLanguagesSliceDeleteAll)
	t.Run("Pilots", testPilotsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Jets", testJetsExists)
	t.Run("Languages", testLanguagesExists)
	t.Run("Pilots", testPilotsExists)
}

func TestFind(t *testing.T) {
	t.Run("Jets", testJetsFind)
	t.Run("Languages", testLanguagesFind)
	t.Run("Pilots", testPilotsFind)
}

func TestBind(t *testing.T) {
	t.Run("Jets", testJetsBind)
	t.Run("Languages", testLanguagesBind)
	t.Run("Pilots", testPilotsBind)
}

func TestOne(t *testing.T) {
	t.Run("Jets", testJetsOne)
	t.Run("Languages", testLanguagesOne)
	t.Run("Pilots", testPilotsOne)
}

func TestAll(t *testing.T) {
	t.Run("Jets", testJetsAll)
	t.Run("Languages", testLanguagesAll)
	t.Run("Pilots", testPilotsAll)
}

func TestCount(t *testing.T) {
	t.Run("Jets", testJetsCount)
	t.Run("Languages", testLanguagesCount)
	t.Run("Pilots", testPilotsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Jets", testJetsHooks)
	t.Run("Languages", testLanguagesHooks)
	t.Run("Pilots", testPilotsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Jets", testJetsInsert)
	t.Run("Jets", testJetsInsertWhitelist)
	t.Run("Languages", testLanguagesInsert)
	t.Run("Languages", testLanguagesInsertWhitelist)
	t.Run("Pilots", testPilotsInsert)
	t.Run("Pilots", testPilotsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("JetToPilotUsingPilot", testJetToOnePilotUsingPilot)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("LanguageToPilots", testLanguageToManyPilots)
	t.Run("PilotToJets", testPilotToManyJets)
	t.Run("PilotToLanguages", testPilotToManyLanguages)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("JetToPilotUsingJets", testJetToOneSetOpPilotUsingPilot)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("LanguageToPilots", testLanguageToManyAddOpPilots)
	t.Run("PilotToJets", testPilotToManyAddOpJets)
	t.Run("PilotToLanguages", testPilotToManyAddOpLanguages)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("LanguageToPilots", testLanguageToManySetOpPilots)
	t.Run("PilotToLanguages", testPilotToManySetOpLanguages)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("LanguageToPilots", testLanguageToManyRemoveOpPilots)
	t.Run("PilotToLanguages", testPilotToManyRemoveOpLanguages)
}

func TestReload(t *testing.T) {
	t.Run("Jets", testJetsReload)
	t.Run("Languages", testLanguagesReload)
	t.Run("Pilots", testPilotsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Jets", testJetsReloadAll)
	t.Run("Languages", testLanguagesReloadAll)
	t.Run("Pilots", testPilotsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Jets", testJetsSelect)
	t.Run("Languages", testLanguagesSelect)
	t.Run("Pilots", testPilotsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Jets", testJetsUpdate)
	t.Run("Languages", testLanguagesUpdate)
	t.Run("Pilots", testPilotsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Jets", testJetsSliceUpdateAll)
	t.Run("Languages", testLanguagesSliceUpdateAll)
	t.Run("Pilots", testPilotsSliceUpdateAll)
}
