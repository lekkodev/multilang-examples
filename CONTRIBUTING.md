# Dev notes

This repo can be used to test compatibility of native Lekko implementations in different languages.

Here are some use cases with example commands:

## Test new version of cli or ts-tools

Install new version of Lekko CLI or update `ts/package.json` to use a new version of `@lekko/ts-transformer`.

Then run `./bisync-all.sh` to test that your change didn't break anything.

## Test how certain native language changes will be translated to config repo and other languages

- Make changes in one of the languages
- run `lekko bisync -r ..` in that language directory
- run `./gen-all` to regenerate all the other languages
- check if changes are reflected correctly in other languages
