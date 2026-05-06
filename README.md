# meshgate

`meshgate` is a compact Go repository for platform engineering, centered on this goal: Model service-mesh routes, retries, traffic splits, and budgets.

## Why This Exists

This is intentionally local and self-contained so it can be inspected without credentials, services, or seeded history.

## Meshgate Review Notes

Start with `route drift` and `secret scope`. Those cases create the widest score spread in this repo, so they are the best quick check when the model changes.

## Capabilities

- `fixtures/domain_review.csv` adds cases for rollout width and quota pressure.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/meshgate-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `route drift` and `secret scope`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Implementation Shape

The repository has two validation layers: the original compact policy fixture and the domain review fixture. They are separate so one can change without hiding failures in the other.

The added Go path is deliberately direct, with fixtures doing most of the explaining.

## Local Usage

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Verification

The check exercises the source code and the review fixture. `edge` is the high score at 224; `recovery` is the low score at 175.

## Roadmap

No external service is required. A deeper version would add more negative cases and a clearer boundary around invalid input.
