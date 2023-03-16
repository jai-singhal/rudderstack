-- create database IF NOT EXISTS rudderstack ;
-- use rudderstack;

CREATE TABLE IF NOT EXISTS tracking_plans (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    display_name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS event_rules (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  tracking_plan_id BIGINT NOT NULL,
  name VARCHAR(255) NOT NULL UNIQUE,
  description TEXT,
  rules JSON NOT NULL,

  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (tracking_plan_id) REFERENCES tracking_plans(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS events (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    properties JSON NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS event_tracking (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    event_id BIGINT REFERENCES events(id) ON DELETE CASCADE,
    event_rule_id BIGINT REFERENCES event_rules(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (event_id, event_rule_id)
);