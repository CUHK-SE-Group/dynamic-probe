# Install Dependencies

```bash
bash init.sh
```

# Running Example

```bash
go run cmd/kb/main.go generate --tmpl "examples/kernel_c/templates/*.gtpl" --conf examples/kernel_c/conf/default.toml --header headers --output-dir examples/kernel_c/output --execute base.c.gtpl
```

