# Implementation Plan: Extended Composer Exit Actions

## Phase 1: Configuration Extension [checkpoint: 5a539cc]
- [x] Task: Add `EnableEnhancedComposerExit` boolean to `Config`, `secureDiskConfig`, and `diskConfig` structs in `config/config.go`. d7c1f92
- [x] Task: Update `SaveConfig` and `LoadConfig` to serialize/deserialize the new field. d7c1f92
- [x] Task: Write a regression test in `config/config_test.go` to ensure `EnableEnhancedComposerExit` persists correctly. d7c1f92
- [x] Task: Conductor - User Manual Verification 'Config Extension' (Protocol in workflow.md) d7c1f92

## Phase 2: Settings UI Integration [checkpoint: a316071]
- [x] Task: Add localization keys for `settings_general.enable_enhanced_composer_exit` in `i18n/locales/en.json` and `es.json`. 0107555
- [x] Task: Add a new toggle to `buildGeneralOptions()` in `tui/settings_general.go`. 0107555
- [x] Task: Update `updateGeneral()` in `tui/settings_general.go` to handle the setting index. 0107555
- [x] Task: Conductor - User Manual Verification 'Settings UI' (Protocol in workflow.md) 0107555

## Phase 3: Composer UI Refactor [checkpoint: ed7e94f]
- [x] Task: Modify `Composer` model in `tui/composer.go` to handle the new keyboard states in the confirmation overlay. bd85392
- [x] Task: Update the `View()` logic in `tui/composer.go` to render the new extended prompt when `EnableEnhancedComposerExit` is true. bd85392
- [x] Task: Implement the `s` (Send), `a` (Abort/Delete), `d` (Save), and `c` (Cancel) logic within the composer's `Update` loop. bd85392
- [x] Task: Conductor - User Manual Verification 'Composer Dialog' (Protocol in workflow.md) bd85392

## Phase 4: Final Verification
- [x] Task: Run `make build` and perform end-to-end manual verification. 6890de8
- [x] Task: Verify that existing "y/n" functionality is unchanged when the setting is OFF. 6890de8
- [x] Task: Conductor - User Manual Verification 'Final Quality Gate' (Protocol in workflow.md) 6890de8
