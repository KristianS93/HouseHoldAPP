CREATE TABLE meals { id serial not null unique,
mealname varchar(64),
householdid varchar(65),
items text,
primary key(id) };
insert INTO meals(mealname, householdid, items)
values (
        'Pasta kødsovs',
        "testhouseid",
        "[{fuld kek},{merekek}]"
    ),
    (
        'Drøm mig væk',
        "testhouseid",
        "[{fuld drøm},{mere væk}]"
    );