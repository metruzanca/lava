# Lava

> 🚧 Lava is currently in Development 🚧

Lava is a minimalist static site generator for Obsidian Vaults.

Heavily inspired by [Quartz](https://quartz.jzhao.xyz/) but not happy configuration how its used. But unlike Quartz, Lava will be heavily opinionated (still configurable) and simpler to use and update.

## Usage

Lava will only require two things:

- The cli installed
- The `lava.toml` file in the root of your obsidian vault.

---

## FaQ

> Why Golang and no xyz?

A couple reasons:

1. I need to practice go for work. I need to become somewhat of an expert.
2. I like the UX of `go install github.com/metruzanca/lava`, especially for development. Its very convenient.
3. It's fast enough. JS with svelte/solid/astro would have been faster to develop and more convenient (due to html/css in js) but I didn't want to make something that was like quartz but worse (because quartz is truly amazing on the frontend side). Also I don't want to write JS in side projects if I can help it.
4. The gopher mascot is cute and I have two of them [gophers.jpeg](.github/assets/gophers.jpeg).
