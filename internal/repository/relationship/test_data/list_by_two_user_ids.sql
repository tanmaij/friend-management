INSERT INTO "users" ("id", "email") 
VALUES 
    (1, 'user1@example.com'),
    (2, 'user2@example.com'),
    (3, 'user3@example.com');

INSERT INTO "relationships" ("id", "requester_id", "target_id", "type") 
VALUES 
    (1, 1, 2, 'friend'),
    (2, 1, 2, 'block');
