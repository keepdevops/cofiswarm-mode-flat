# cofiswarm-mode-flat

Cofiswarm component: `mode-flat`.

- Layout: [REPO-STANDARD-LAYOUT](https://github.com/keepdevops/cofiswarmdev/blob/main/docs/REPO-STANDARD-LAYOUT.md)
- Migration: [MIGRATION-SPRINTS](https://github.com/keepdevops/cofiswarmdev/blob/main/docs/MIGRATION-SPRINTS.md)

## FHS paths

| Path | Purpose |
|------|---------|
| `/etc/cofiswarm/mode-flat/` | config |
| `/var/lib/cofiswarm/mode-flat/` | state |
| `/var/log/cofiswarm/mode-flat/` | logs |

## Test

```bash
./test/scripts/assert-layout.sh mode-flat
```
