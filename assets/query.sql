select id_control_unit_data as ID_TRIP, ship_description, id_ship, state_name AS STATE, id_state, ts_main_event_field_val
from trips_logs
inner join control_unit_data 
on fk_control_unit_data = id_control_unit_data 
inner join ships
on control_unit_data.fk_ship = id_ship
inner join states
on trips_logs.fk_state = id_state
and fk_state != 10
where control_unit_data.fk_portinformer = 28
and ts_main_event_field_val > '2018-01-01 00:00'
GROUP BY id_trip, ts_main_event_field_val, ship_description, id_ship, state, id_state
ORDER BY id_control_unit_data, state