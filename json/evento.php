<?php
$json = file_get_contents("eventos.json");

$list_eventos = json_encode($json);
$id = $_POST["id"];

foreach ($list_eventos as $evento) {
  if($evento->id == $id){
    print($evento);
    break;
  }
}

die();
