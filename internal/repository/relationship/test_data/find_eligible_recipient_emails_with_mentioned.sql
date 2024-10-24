INSERT INTO "users" ("id", "email") 
VALUES 
    (1, 'user1@example.com'),
    (2, 'user2@example.com'),
    (3, 'user3@example.com'),
   	(4, 'user4@example.com'),
    (5, 'user5@example.com');

INSERT INTO "relationships" ("id", "requester_id", "target_id", "type") 
VALUES 
    (1, 1, 2, 'friend'),
    (2, 1, 3, 'friend'),
    (3, 4, 2, 'friend'),
    (4, 4, 3, 'friend'),
    (5, 5, 2, 'friend'),
    (6, 5, 3, 'friend'),
    (7, 1, 5, 'friend'),
   	(8, 4, 1, 'subscribe'),
	(9, 2, 1, 'block');
