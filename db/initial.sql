drop database rudderstack;
create database rudderstack;
use rudderstack;

CREATE TABLE tracking_plans (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    display_name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE event_rules (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  tracking_plan_id BIGINT NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  rules JSON NOT NULL,

  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE (tracking_plan_id, name),
  FOREIGN KEY (tracking_plan_id) REFERENCES tracking_plans(id) ON DELETE CASCADE
);

CREATE TABLE events (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    properties JSON NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE event_tracking (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    event_id BIGINT REFERENCES events(id) ON DELETE CASCADE,
    event_rule_id BIGINT REFERENCES event_rules(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (event_id, event_rule_id)
);