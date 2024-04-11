# ranked-choice

Run:
```bash
docker compose run rp-api go run main.go
```

## UI Requirements
Reactive and usable on both desktop and mobile.

### Elements

- Navigation bar
  - Always visible
  - Logo
  - New survey creation button
  - Current user chip
- Survey create view
  - Prompt (optional)
  - Create survey options
  - Set response limit (optional)
- Survey response view (open survey)
  - Rank options. Experiment with UX
    - Drag and drop
    - Number input
    - Arrows
  - Submission limit, e.g. "3/5"
    - Percentage status bar
