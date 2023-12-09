with
    input_file as (
     select pg_read_file as f from pg_read_file('/Users/aros/proj/advent-of-code/2023/day-05/input.txt')
    ),
    all_sections as (
        select regexp_split_to_table(input_file.f, E'\n\n') as s from input_file
    ),
    maps_sections as (
        select trim(s, E'\n') as s from all_sections offset 1
    ),
    seeds_line as (
        select split_part(input_file.f, E'\n\n', 1) as s from input_file
    ),
    seeds1 as (
        select regexp_split_to_table(seeds_line.s, E' ') as seed_id from seeds_line
    ),
    seeds as (
        select seed_id::int  from seeds1 where seeds1.seed_id != 'seeds:'
    ),
    map_sections1 as (
        select regexp_split_to_table(s, E'\n\n') as s from maps_sections
    ),
    map_sections2 as (
        select regexp_split_to_array(s, E'\n') as s from map_sections1
    ),
    map_sections3 as (
        select s[1] as id, s[2:] as map from map_sections2
    ),
    map_sections4 as (
        select regexp_replace(map_sections3.id, ' .*', '') as id, map_sections3.map from map_sections3
    ),
    map_sections5 as (
        select regexp_split_to_array(map_sections4.id, '-to-') as id, map_sections4.map from map_sections4
    ),
    map_sections6 as (
        select id[1] as from_type, id[2] as to_type, map_sections5.map from map_sections5
    ),
    map_sections7 as (
        select map_sections6.from_type, map_sections6.to_type,
               map_sections6.map[i] as map
        from map_sections6, generate_subscripts(map_sections6.map, 1) as i
    ),
    map_sections8 as (
        select from_type, to_type, regexp_split_to_array(map, ' ') as map from map_sections7
    ),
    map_sections9 as (
        select
            from_type,
            to_type,
            CAST(map[1] as int) as to_start,
            CAST(map[2] as int) as from_start,
            CAST(map[3] as int) as count
        from map_sections8
    ),
    map_sections10 as (
        select 
            from_type,
            to_type, 
            count, 
            generate_series(from_start, from_start+count-1) as from_id,
            generate_series(to_start, to_start+count-1) as to_id
        from map_sections9
    ),
    mappings1 as (
        select from_type, to_type, from_id, to_id from map_sections10
    ),
    type_mappings as (
        select from_type, to_type from mappings1 group by from_type, to_type
    ),
    go1 as (
        select
            'seed' as from_type,
            seeds.seed_id as from_id,
            coalesce(mappings1.to_type, tm.to_type) as to_type,
            coalesce(mappings1.to_id, seeds.seed_id) as to_id
        from seeds
            join type_mappings as tm
                  on tm.from_type = 'seed'
            left outer join mappings1
                on seeds.seed_id = mappings1.from_id
                       and mappings1.from_type = 'seed'

    ),
    go2 as (
        select
            go1.to_type as from_type,
            go1.to_id as from_id,
            coalesce(mappings1.to_type, tm.to_type) as to_type,
            coalesce(mappings1.to_id, go1.to_id) as to_id
        from go1
                 join type_mappings as tm
                      on tm.from_type = go1.to_type
                 left outer join mappings1
                                 on go1.to_id = mappings1.from_id
                                     and mappings1.from_type = go1.to_type
    ),
    go3 as (
        select
            go2.to_type as from_type,
            go2.to_id as from_id,
            coalesce(mappings1.to_type, tm.to_type) as to_type,
            coalesce(mappings1.to_id, go2.to_id) as to_id,
            go2.to_id as go2_to_id
        from go2
                 join type_mappings as tm
                      on tm.from_type = go2.to_type
                 left outer join mappings1
                                 on go2.to_id = mappings1.from_id
                                     and mappings1.from_type = go2.to_type

    ),
    go4 as (
        select
            go3.to_type as from_type,
            go3.from_id as from_id,
            coalesce(mappings1.to_type, tm.to_type) as to_type,
            coalesce(mappings1.to_id, go3.to_id) as to_id
        from go3
                 join type_mappings as tm
                      on tm.from_type = go3.to_type

                 left outer join mappings1
                                 on go3.to_id = mappings1.from_id
                                     and mappings1.from_type = go3.to_type

    ),
    go5 as (
        select
            go4.to_type as from_type,
            go4.from_id as from_id,
            coalesce(mappings1.to_type, tm.to_type) as to_type,
            coalesce(mappings1.to_id, go4.to_id) as to_id

        from go4
                 join type_mappings as tm
                      on tm.from_type = go4.to_type
                 left outer join mappings1
                                 on go4.to_id = mappings1.from_id
                                     and mappings1.from_type = go4.to_type

    ),
    go6 as (
        select
            go5.to_type as from_type,
            go5.from_id as from_id,
            coalesce(mappings1.to_type, tm.to_type) as to_type,
            coalesce(mappings1.to_id, go5.to_id) as to_id

        from go5
                 join type_mappings as tm
                      on tm.from_type = go5.to_type
                 left outer join mappings1
                                 on go5.to_id = mappings1.from_id
                                     and mappings1.from_type = go5.to_type

    ),
    go7 as (
        select
            go6.to_type as from_type,
            go6.from_id as from_id,
            coalesce(mappings1.to_type, tm.to_type) as to_type,
            coalesce(mappings1.to_id, go6.to_id) as to_id
        from go6
                 join type_mappings as tm
                      on tm.from_type = go6.to_type
                 left outer join mappings1
                                 on go6.to_id = mappings1.from_id
                                     and mappings1.from_type = go6.to_type

    ),
    go8 as (
        select
            go7.to_type as from_type,
            go7.from_id as from_id,
            coalesce(mappings1.to_type, tm.to_type) as to_type,
            coalesce(mappings1.to_id, go7.to_id) as to_id
        from go7
                 join type_mappings as tm
                      on tm.from_type = go7.to_type
                 left outer join mappings1
                                 on go7.to_id = mappings1.from_id
                                     and mappings1.from_type = go7.to_type

    ),
    asdf as (select 1 as x)
    select min(to_id) from go7;

--     go2 as (
--         select * from go1
--                  join mappings1
--                  on go1.to_type = where from_type = go1.to_type
--     )
--     select * from mappings1 as a, mappings1 as b
--     where a.to_type = b.from_type;
;
