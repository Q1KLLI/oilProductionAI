select
    s1.well_id,
     w.name,
    s1.event
  from state s1
  join (
        select
               max(dt) dt,
               well_id
          from
               state
         group by
            well_id
  ) s2
    on s1.well_id = s2.well_id
   and s1.dt = s2.dt
  join well w on w.id = s1.well_id
