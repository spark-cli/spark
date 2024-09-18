<div align="center">
  
  # Spark
  :sparkles: Ignite your code with a spark.
</div>

> [!WARNING]  
> This project is WIP. Advertised features may not exist yet.

Spark generates and syncs code from a language agnostic code templates.

## Usage

### Creating a project

This creates a `.spark.toml` and `.spark-lock.toml` file.
```sh
spark new
```

### Editing config

Updating the config will also run it.
```sh
spark source add spark github.com/spark-cli/spark-
spark add spark/nodejs -- --name my-project --author bricked
spark add spark/astro
spark source add component spark/astro-component dev
spark add component -- --name button
```

### Config file

```toml
[sources.spark]
source = "github.com/spark-cli/spark-"

[sources.nodejs]
source = "spark/nodejs"

[sources.astro]
source = "spark/astro"

[sources.component]
source = "spark/astro-component"
branch = "dev"

[sparks.nodejs]
name = "my-project"
author = "bricked"

[sparks.component.button]
name = "button"
```

### Lock file

```toml
[sparks.nodejs]
commit = "9121b541434dc7e8b73329f2c3e6b59a2b762745"

[sparks.astro]
commit = "350cd8123f831d5f92301fb9dce1f0c36c3b67ee"

[sparks.component.button]
commit = "bdd00c4e5961d691cb7c5ca8b5607d3b775b0a87"
```
