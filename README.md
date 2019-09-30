# World Generator

This is a tool and library to procedurally generate fantasy worlds.

The API runs on port 7531.

## API Routes

The API supports the following routes. Where `123` is specified, it refers to an alphanumeric seed.

**Character, random**: `/character`

**Character, specific seed**: `/character/123`

**Climate, random**: `/climate`

**Climate, specific seed**: `/climate/123`

**Country, random**: `/country`

**Country, specific seed**: `/country/123`

**Culture, random**: `/culture`

**Culture, specific seed**: `/culture/123`

**Heraldry, random**: `/heraldry`

**Heraldry, specific seed**: `/heraldry/123`

**Language, random**: `/language`

**Language, specific seed**: `/language/123`

**Merchant, random**: `/merchant`

**Merchant, specific seed**: `/merchant/123`

**Monster, random**: `/monster`

**Monster, specific seed**: `/monster/123`

**Organization, random**: `/organization`

**Organization, specific seed**: `/organization/123`

**Pantheon, random**: `/pantheon`

**Pantheon, specific seed**: `/pantheon/123`

**Region, random**: `/region`

**Region, specific seed**: `/region/123`

**Town, random**: `/town`

**Town, specific seed**: `/town/123`

## Building

Just run `./build.sh`. The binary will be `build/worldapi`.
