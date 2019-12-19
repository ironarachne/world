CREATE TABLE dice (
  id INTEGER PRIMARY KEY,
  "number" INTEGER,
  sides INTEGER
);

CREATE TABLE size_categories (
  id INTEGER PRIMARY KEY,
  name TEXT,
  level INTEGER,
  base_height INTEGER,
  base_weight INTEGER
);

CREATE TABLE age_categories (
  id INTEGER PRIMARY KEY,
  name TEXT,
  age_min INTEGER,
  age_max INTEGER,
  male_height_base INTEGER,
  female_height_base INTEGER,
  height_range_dice_id INTEGER,
  male_weight_base INTEGER,
  female_weight_base INTEGER,
  weight_modifier INTEGER,
  weight_range_dice_id INTEGER,
  size_category_id INTEGER,
  commonality INTEGER
);

CREATE TABLE template_tags (
  template_id INTEGER,
  tag TEXT
);

CREATE TABLE template_values (
  id INTEGER PRIMARY KEY,
  template_id INTEGER,
  value TEXT
);

CREATE TABLE template_descriptors (
  id INTEGER PRIMARY KEY,
  template_id INTEGER,
  descriptor TEXT
);

CREATE TABLE templates (
  id INTEGER PRIMARY KEY,
  name TEXT
);

CREATE TABLE resource_tags (
  resource_id INTEGER,
  tag TEXT
);

CREATE TABLE resources (
  id INTEGER PRIMARY KEY,
  name TEXT,
  description TEXT,
  main_material TEXT,
  origin TEXT,
  commonality INTEGER,
  value INTEGER
);

CREATE TABLE quirks (
  id INTEGER PRIMARY KEY,
  slot_id INTEGER,
  quirk TEXT
);

CREATE TABLE slots (
  id INTEGER PRIMARY KEY,
  pattern_id INTEGER,
  name TEXT,
  required_tag TEXT,
  description_template TEXT
);

CREATE TABLE pattern_tags (
  pattern_id INTEGER,
  tag TEXT
);

CREATE TABLE patterns (
  id INTEGER PRIMARY KEY,
  name TEXT,
  description TEXT,
  commonality INTEGER,
  profession_id INTEGER,
  name_template TEXT,
  main_material TEXT,
  main_material_override TEXT,
  origin_override TEXT,
  value INTEGER
);

CREATE TABLE species_tags (
  species_id INTEGER,
  tag TEXT
);

CREATE TABLE species_possible_traits (
  template_id INTEGER,
  species_id INTEGER
);

CREATE TABLE species_common_traits (
  template_id INTEGER,
  species_id INTEGER
);

CREATE TABLE species_age_categories (
  species_id INTEGER,
  age_category_id INTEGER
);

CREATE TABLE species_resources (
  species_id INTEGER,
  resource_id INTEGER
);

CREATE TABLE species (
  id INTEGER PRIMARY KEY,
  name TEXT,
  plural_name TEXT,
  adjective TEXT,
  commonality INTEGER,
  humidity_min INTEGER,
  humidity_max INTEGER,
  temperature_min INTEGER,
  temperature_max INTEGER
);

CREATE TABLE profession_tags (
  profession_id INTEGER,
  tag TEXT
);

CREATE TABLE professions (
  id INTEGER PRIMARY KEY,
  name TEXT,
  description TEXT
);