drop table if exists person;

CREATE TABLE person (
    id UUID PRIMARY KEY,
    email VARCHAR(255),
    phone VARCHAR(20),
    first_name VARCHAR(50),
    last_name VARCHAR(50)
);

INSERT INTO person (id, email, phone, first_name, last_name)
    VALUES ('a93f84a9-078a-49ad-a042-b281d2d0dbc9', 'test@gmail.com', '79005002030', 'Alex', 'Smith'),
    ('8c2ee53f-ae6a-4db3-9597-316a2f30c619', 'test@gmail.com', '79005002030', 'Alex', 'Smith'),
    ('6b985ff6-ef9a-4042-ab40-28d1ecd3ad1a', 'test@gmail.com', '79005002030', 'Alex', 'Smith'),
    ('df93d8bb-091a-4f47-966b-f4dcfdf0585f', 'test@gmail.com', '79005002030', 'Alex', 'Smith'),
    ('e95ca454-82fb-461e-8a34-251a0a3c7215', 'test@gmail.com', '79005002030', 'Alex', 'Smith'),
    ('9843b4b8-6c55-44a1-8d89-d581105988b9', 'test@gmail.com', '79005002030', 'Alex', 'Smith'),
    ('c5a53ddc-8afe-41d3-b9ee-062df721b4dd', 'test@gmail.com', '79005002030', 'Alex', 'Smith'),
    ('3cf11415-d57a-4876-98e9-a13369932232', 'test@gmail.com', '79005002030', 'Alex', 'Smith'),
    ('be0e5c1e-4e63-4b80-97d1-1b439281e3f3', 'test@gmail.com', '79005002030', 'Alex', 'Smith'),
    ('499defa0-a926-4359-b664-94fb0387335e', 'test@gmail.com', '79005002030', 'Alex', 'Smith');
         