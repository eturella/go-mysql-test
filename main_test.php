<?php
$link = mysqli_connect('localhost:3306','root','') or die ('ERR: mysqli_connect');
$res = mysqli_query($link, "show variables") or die('ERR: mysqli_query');
$row = mysqli_fetch_assoc($res);
print_r($row);

$res = mysqli_query($link, "SELECT @@global.max_allowed_packet") or die('ERR: mysqli_query');
$row = mysqli_fetch_assoc($res);
print_r($row);

$res = mysqli_query($link, "select * from dual") or die('ERR: mysqli_query');
$row = mysqli_fetch_assoc($res);
print_r($row);
