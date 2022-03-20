-- organizations
INSERT INTO organizations (org_id) VALUES (1);
INSERT INTO organizations (org_id) VALUES (2);

-- users
INSERT INTO users (user_id, org_id) VALUES (1, 1);
INSERT INTO users (user_id, org_id) VALUES (2, 1);
INSERT INTO users (user_id, org_id) VALUES (3, 2);

-- services and versions (org 1)
INSERT INTO services (service_id, title, summary, org_id) VALUES (1, 'Test 1', 'Test 1 summary', 1);
INSERT INTO versions (service_id, version_id, summary) VALUES (1, 1, 'Test Version 1');

INSERT INTO services (service_id, title, summary, org_id) VALUES (2, 'Test 2', 'Test 2 summary', 1);
INSERT INTO versions (service_id, version_id, summary) VALUES (2, 1, 'Test Version 2');

INSERT INTO services (service_id, title, summary, org_id) VALUES (3, 'Test 3', 'Test 3 summary', 1);
INSERT INTO versions (service_id, version_id, summary) VALUES (3, 1, 'Test Version 3');

INSERT INTO services (service_id, title, summary, org_id) VALUES (4, 'Test 4', 'Test 4 summary', 1);
INSERT INTO versions (service_id, version_id, summary) VALUES (4, 1, 'Test Version 4');

-- services and versions (org 2)
INSERT INTO services (service_id, title, summary, org_id) VALUES (5, 'Test 5', 'Test 5 summary', 2);
INSERT INTO versions (service_id, version_id, summary) VALUES (5, 1, 'Test Version 5');
