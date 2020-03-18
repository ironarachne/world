# Change Log

## 0.11.1 (2020-03-18)

- Add JSON hinting for music

## 0.11.0 (2020-03-16)

- Add geographic areas, which include biomes, climates, and geographic regions
- Replace previous "climate" package with the new geographic areas
- Rework music style package to reflect musical theory instead of instruments
- Change how plants and animals are derived from geography
- Remove climate from culture to make it independent, despite originating a culture in a geography
- Fix deity relationships and add relationships to deity descriptions
- Update to use Go 1.14

## 0.10.1 (2020-01-25)

- Fix bug where descriptors were dropped for blazon descriptions of singular charges

## 0.10.0 (2020-01-25)

- Refactor heraldry package into several packages
- Add about 130 new charges to heraldry generator
- Add NumberToWord function

## 0.9.2 (2020-01-15)

- Use simplified devices instead of devices
- Add error handling for API endpoints
- Add sample phrase and sample phrase translation to full language

## 0.9.1 (2020-01-15)

- Add JSON hinting to alcoholic drinks
- Add description to Deity
- Improve error messages for heraldry generation

## 0.9.0 (2020-01-10)

- Switch to using JSON files for data instead of hard-coded values
- Add insects and fish to animal data for climates
- Fix bug in ordering coordinates by distance
- Change where heraldry images are saved
- Switch to using Make for development
- Added /data endpoints to the API

## 0.8.19 (2019-12-03)

- Add ability to generate regions from a culture passed in via POST

## 0.8.18 (2019-12-03)

- Add drinking culture style to culture
- Add json hinting to a bunch more packages

## 0.8.17 (2019-12-03)

- Add ability to generate new words for a given language
- Add json hinting to all structs required for generating languages
- Return full language data from `/language` endpoints, not a simplified version

## 0.8.16 (2019-12-02)

- Add ability to generate cultures from a climate passed in via POST

## 0.8.15 (2019-12-02)

- Fix bug in tree generation that prevented trees being generated for climates with low humidity
- Add json hinting to all structs required for generating climates fully
- Make the `/climate` endpoints return a full climate instead of a simplified climate

## 0.8.14 (2019-11-25)

- Add remaining package description comments

## 0.8.13 (2019-11-21)

- Make file save target configurable
- Add ability to save PNG files to disk instead of Digital Ocean Spaces
- Add `run.sh` script to make running the API easier for local development
- Add package description comments to about half of the packages in this project

## 0.8.12 (2019-11-20)

- Expand organization member ranks to have multiple possible age categories
- Refactor character hobbies to use multiple potential age categories instead of adult/child booleans
- Fix a bug in hobby generation when there was only one potential hobby
- Fix a minor wording problem in character descriptions regarding motivation

## 0.8.11 (2019-11-08)

- Refactor climate generation

## 0.8.10 (2019-11-07)

- Convert fish to Species struct
- Convert insects to Species struct
- Convert monsters to Species struct
- Convert plants to Species struct
- Convert trees to Species struct

## 0.8.9 (2019-11-07)

- Convert animals to Species struct
- Remove Animal struct

## 0.8.8 (2019-11-06)

- Split language package into conlang, language, and writing packages
- Add support for premade languages
- Add the Common premade language
- Remove "practice suffix"
- Add basic religion name generation

## 0.8.7 (2019-11-03)

- Add several new weapon patterns
- Add hamlets and boroughs to town generation
- Add several utility functions to species

## 0.8.6 (2019-10-05)

- Add tiefling to the race package

## 0.8.5 (2019-10-01)

Make HTML index page with links to generators

### Change List

- Make HTML index page with links to generators

## 0.8.4 (2019-09-27)

Fix a bug in clothing style generation.

### Change List

- Add prefix and suffix options to kilts
- Add trews

## 0.8.3 (2019-09-27)

Fix a bug in deity holy symbol generation.

### Change List

- Add a check for no domain holy symbols in deity holy symbol generation

## 0.8.2 (2019-09-26)

Allow for configuring the domain name static assets are saved to.

### Change List

- Add configurable static asset domain via `STATIC_DOMAIN_NAME` environment variable

## 0.8.1 (2019-09-23)

Make region rulers noble families instead of individuals.

### Change List

- Make region rulers organizations
- Make region rulers noble families
- Add noble families
- Fix a bug in organization size calculation
- Remove a debugging line from cloud save
- Fix a bug in age calculation
- Add heraldry to organizations as an option

## 0.8.0 (2019-09-21)

Refactor races as species.

### Change List

- Add "species" package
- Replace race object with species
- Refactor height and weight to be dependent on age category
- Include age categories in species definition
- Add wood elves, high elves, mountain dwarves, hill dwarves, and gnomes
- Add trait system
- Tweak weighting for races
- Replace old character appearance system with new trait system
- Give dwarves beards and halflings hairy feet

## 0.7.2 (2019-09-17)

Improve organizations.

### Change List

- Add wizard societies
- Add age calculation for organization members based on rank
- Add size limit variation to organizations

## 0.7.1 (2019-09-16)

Improve how organizations are generated.

### Change List

- Add wizardry schools to organizations
- Add ranks to organizations
- Make the number of organizations in a region dependent on its size
- Add profession name to its list of tags

## 0.7.0 (2019-09-16)

Improve the pantheon package and relationships.

### Change List

- Split domains and deities into their own packages
- Split relationships into their own package
- Added more logic to determining relationship validity
- Added more descriptors to relationships

## 0.6.5 (2019-09-16)

Fix bug in tinctures

### Change List

- Fix a bug that prevented tinctures from loading

## 0.6.4 (2019-09-13)

Install ca-certificates in the Docker image

### Change List

- Install ca-certificates in the Docker image

## 0.6.3 (2019-09-13)

Change the base image used for building the Docker image.

### Change List

- Change the base image from `scratch` to `ubuntu:18.04` for the Docker image

## 0.6.2 (2019-09-13)

Make the data path configurable.

### Change List

- Make data path configurable via `WORLDAPI_DATA_PATH` environment variable

## 0.6.1 (2019-09-13)

Add debugging in the form of logging errors for the API.

### Change List

- Add error logging to the API

## 0.6.0 (2019-09-13)

This version has a completely rewritten heraldry package. It now renders PNG images instead of SVG images and saves
them to Digital Ocean Spaces.

### Change List

- Rewrite heraldry package from scratch
- Update error handling

## 0.5.0 (2019-09-06)

### Change List

- Add tags to trees
- Change deciduous/coniferous logic to use tags
- Add cartwrights and land vehicles
- Add warlock, necromancer, and philosopher professions
- Change how resources and trade goods are named and described
- Clean up climate plant generation
- Clean up climate animal generation
- Change how climates are generated
- Fix obscure bug in climates for regions and towns
- Move fish to their own package
- Add tags to fish

## 0.4.14 (2019-09-05)

Fix bad random logic.

### Change List

- Fix bad random slice item logic throughout the code

## 0.4.13 (2019-09-05)

Add real error handling.

### Change List

- Add error handling to most areas
- Improve the word list for cultures

## 0.4.12 (2019-09-03)

Improve stone and religion just a bit.

### Change List

- Add stone types
- Modify generation of stone resources
- Add virtues to religion
- Add names to religion

## 0.4.11 (2019-08-30)

Improve character description variety.

### Change List

- Replace static character description format with a new random template system

## 0.4.10 (2019-09-29)

Add insects, honey, and beekeepers.

### Change List

- Add insects
- Add honey
- Add beekeepers
- Fix some bugs with plants

## 0.4.9 (2019-09-29)

Add cactii, move trees to their own package, recategorize "fruits" into different groups.

### Change List

- Move trees into their own package
- Move "fruits" into bushes and melons
- Move avocados to trees
- Add cactii

## 0.4.8 (2019-08-28)

Add travelling merchants.

### Change List

- Add merchant endpoint
- Add values to resources
- Add price calculation to trade goods
- Improve a number of basic and refined resources
- Add a few new resources

## 0.4.7 (2019-08-28)

Add several new resources and the ability to override resource origins.

### Change List

- Add distilled beverages
- Add metal alloys
- Add velveteen as a fabric
- Support overriding the origin of a resource

## 0.4.6 (2019-08-27)

Add Sentry error tracking.

### Change List

- Add mandatory Sentry error tracking

## 0.4.5 (2019-08-01)

Add ability to save to Digital Ocean Spaces.

### Change List

- Add ability to save PNG and SVG images to Digital Ocean Spaces

## 0.4.4 (2019-07-28)

This is a big refactor of the heraldry package. This work enables a later addition of generating heraldry with specific charge types.

### Change List

- Refactor heraldry package to improve adherence to Go best practices
- Add heraldry utility functions for finding charges by tag
- Add more tags to charges

## 0.4.3 (2019-07-25)

Add tags to animals

### Change List

- Add tags to animals
- Fix bug in climate generation that sometimes prevented clothing styles from being created

## 0.4.2 (2019-07-24)

Add resource tiers

### Change List

- Implements tiered resources
- Summarizes resources as trade goods
- Adds a number of resources and patterns

## 0.4.1 (2019-07-21)

Improved resources, professions, and organizations.

### Change List

- Fix bug in organization generation where it didn't recognize all types of organization
- Add crafting guilds back to organization generation
- Add blacksmith, bowyer, and farrier goods
- Add many more patterns, especially to armor and weapons
- Add ability to turn patterns into resources
- Add a few new professions
- Add tags to professions that were missing them

## 0.4.0 (2019-07-20)

Massive rework of resources and institution of professions.

### Change List

- Add professions
- Use professions everywhere, including for resource generation and for characters
- Add resource system to replace existing "trade goods" resources

## 0.3.2 (2019-07-19)

Changes to food.

### Change List

- Add basic desserts
- Remove fruit from main dishes as a base

## 0.3.1 (2019-07-19)

Small bug fixes.

### Change List

- Fix a bug where weapons as trade goods would cause a crash
- Fix a bug where a climate wouldn't always include at least one fabric plant, one fruit, and one grain

## 0.3.0 (2019-07-19)

Big upgrade to how resources are dealt with.

### Change List

- Move resources to their own package
- Declare resources with their owning item instead of elsewhere
- Move drinks to their own package
- Add more detail to drinks and make them ingredient-dependent

## 0.2.7 (2019-07-18)

Integrates size package with races.

### Change List

- Replace race size with race size category
- Add unique traits to each race
- Add unique traits to race description

## 0.2.6 (2019-07-18)

Adds basic monsters.

### Change List

- Add basic monsters and a monster API
- Add size package and use it with monsters and animals

## 0.2.5 (2019-07-14)

Fixes a bug in deity simplification.

### Change List

- Fix a bug in generating deities where sometimes a blank simplified deity would be created

## 0.2.4 (2019-07-01)

Adds translation to languages and population of word lists based
on the culture's origin climate.

### Change List

- Update word list based on culture climate
- Add translation functions
- Add capitalize first function

## 0.2.3 (2019-07-01)

Big update to languages this time around. This is foundational work
for real phrasebooks and better name generation.

### Change List

- Added generation of coherent word lists to languages
- Added a "Rosetta stone" phrase to language output
- Added basic verb conjugation support

## 0.2.2 (2019-07-01)

Fixed two crash bugs.

### Change List

- Fixed a crash when holy items weren't being calculated correctly
- Fixed a crash when no fish were generated for a climate

## 0.2.1 (2019-06-30)

Building styles and clothing styles are the big winners in this release.
Also quite a few bug fixes made it in.

### Change List

- Clothing styles are much improved, with outfits making more sense
- Building styles got a little more variety
- Refactoring was done for animals and plants to make changes easier in the future
- Bug fixes, primarily in materials

## 0.2.0 (2019-06-25)

This version moves to using manual semantic versioning. The previous
scheme used 0.1.X versioning, where X was the CircleCI build number.

The version number 0.2.0 was chosen to differentiate it from that
process.

### Change List

- Tag the current version of the World API as v0.2.0
- Update the CI/CD configuration for CircleCI to reflect the new scheme
- Add the git commit hash as an additional Docker image tag
