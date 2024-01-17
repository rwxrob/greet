# Importable Cobra command example

* Allows creation of independently managed Cobra command trees
* Prefers use of top-level package as library instead of executable
* Prefers commands under the `cmd` directory
* Does not pollute the `cmd` directory with private subcommands
* Allows subcommands to also be imported directly enabling subcommand library modules
