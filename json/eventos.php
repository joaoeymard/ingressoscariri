<?php
$json = file_get_contents("eventos.json");

print(json_encode($json));

die();
