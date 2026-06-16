# cofiswarm-mode-flat

Mode plugin service (flat) — HTTP execute endpoint for `cofiswarm-dispatch`.

- SDK: [cofiswarm-mode-sdk](../cofiswarm-mode-sdk)
- Legacy C++: `legacy/cpp/`
- Config gate: `dispatch_url` + `slot_manager_url` in `test/standalone/etc/cofiswarm/mode-flat/mode-flat.yaml`

## Run

```bash
make build
./bin/cofiswarm-mode-flat -config test/standalone/etc/cofiswarm/mode-flat/mode-flat.yaml
```

Default listen: `:8021`
