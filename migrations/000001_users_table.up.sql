create table users(
    id int generated always as identity primary key,
    username varchar(255) unique not null,
    email varchar(255) not null,
    password text not null,
    points int default 0,
    created_at timestamp not null default now(),
    updated_at timestamp
);

create table tasks (
    id int generated always as identity primary key,
    name varchar(255) not null,
    description text,
    points_reward int not null,
    created_at timestamp not null default now()
);

create table user_tasks (
  user_id int references users(id),
  task_id int references tasks(id),
  completed_at timestamp not null default now(),
  primary key (user_id, task_id)
);

create table referrals (
  referrer_id int references users(id),
  referred_id int references users(id),
  created_at timestamp not null default now(),
  primary key (referrer_id, referred_id)
);

insert into tasks(name, description, points_reward) values
('join_telegram', 'Join our Telegram channel', 100),
('follow_twitter', 'Follow us on Twitter', 50),
('refer_friend', 'Refer a friend', 200);