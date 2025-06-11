# Integration Test: Flag Creation with Tags

This integration test verifies that flag creation correctly includes tags in the API response after the bug fix.

## Problem Solved

Previously, when creating flags (AB experiments and latches), the tags were being associated correctly in the database but were **not included in the API response**. This caused issues where:

- AB experiments didn't show the "AB" tag in the creation response
- Latches didn't show the "latch" tag in the creation response
- Frontend applications couldn't immediately see the associated tags

## Test Coverage

The integration test covers three scenarios:

### 1. AB Experiment Flag Creation
- Creates a basic AB experiment flag
- Verifies "AB" tag is included in response
- Checks default flag properties

### 2. AB Experiment with Template
- Creates an AB experiment using `simple_boolean_flag` template
- Verifies "AB" tag is included in response
- Verifies template-created segments and variants are included
- Checks that the "on" variant is created

### 3. Latch Creation
- Creates a latch flag
- Verifies "latch" tag is included in response
- Verifies latch-created segments and variants are included
- Checks that the "APPLICABLE" variant is created

## How to Run

```bash
# Run the integration test
go run integration_test_flag_creation_tags.go
```

## Expected Output

```
🧪 Integration Test: Flag Creation with Tags
============================================================

📋 Setting up test environment...
✅ Test database initialized
✅ Test environment ready

🧪 Test 1: AB Experiment Flag Creation
----------------------------------------
Creating AB experiment flag...
✅ AB Flag created successfully
   ID: 6
   Description: Integration Test AB Experiment
   Key: integration_test_ab_experiment
   Tags: AB
   ✅ AB tag correctly included in response!
   ✅ Flag created with default enabled=false
   ✅ Flag created with default dataRecordsEnabled=false

🧪 Test 2: AB Experiment Flag with Template
----------------------------------------
Creating AB experiment flag with simple_boolean_flag template...
✅ AB Flag with template created successfully
   ID: 7
   Description: Integration Test AB Experiment with Template
   Key: integration_test_ab_experiment_template
   Template: simple_boolean_flag
   Tags: AB
   ✅ AB tag correctly included in response!
   Segments: 1
   ✅ Template segment created with 100% rollout
   Variants: 1
   ✅ Template variant created with key: on

🧪 Test 3: Latch Creation
----------------------------------------
Creating latch...
✅ Latch created successfully
   ID: 8
   Description: Integration Test Latch
   Key: integration_test_latch
   Tags: latch
   ✅ Latch tag correctly included in response!
   Segments: 1
   ✅ Latch segment created with 100% rollout
   Variants: 1
   ✅ Latch variant created with key: APPLICABLE

🎉 All integration tests passed!
Flag creation now correctly includes tags in the response.
```

## Technical Details

The fix involved modifying the `mapResponseAndSaveFlagSnapShot` function in `pkg/handler/crud_flag_creation.go` to reload the flag with all relationships (segments, variants, tags) before mapping it to the API response.

**Before (Broken):**
```go
payload, err := e2rMapFlag(f)  // f.Tags is empty - not preloaded
```

**After (Fixed):**
```go
// Reload the flag with all relationships
if err := entity.PreloadSegmentsVariantsTags(getDB()).First(f, f.ID).Error; err != nil {
    // handle error
}
payload, err := e2rMapFlag(f)  // f.Tags now populated
```

This ensures that the API response includes all the associated data that was created during the flag creation process.

## Files Modified

- `pkg/handler/crud_flag_creation.go` - Fixed the response mapping
- `pkg/handler/crud_test.go` - Added unit tests
- `integration_test_flag_creation_tags.go` - This integration test

## Related Issues

This fix resolves the issue where flag creation responses were inconsistent with other flag operations that properly include all relationships in their responses.
