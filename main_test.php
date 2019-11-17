<?php
$link = mysqli_connect('localhost:3306','root','') or die ('ERR: mysqli_connect');
$res = mysqli_query($link, "select * from dual") or die('ERR: mysqli_query');
$row = mysqli_fetch_assoc($res);
print_r($row);
