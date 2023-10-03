create table todo ( 
  id int auto_increment primary key,
  note varchar(255),
  completed bool,
  position int,
  lastChanged timestamp
);

create trigger set_position_on_insert
    before insert on todo for each row
begin
    declare max_order int;
    select MAX(position) into max_order from todo;
    set new.position = IFNULL(max_order, 0) + 1;
end;
